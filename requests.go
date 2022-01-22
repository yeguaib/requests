package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

//go:generate mockgen -source=requests.go -destination=mock/requests.go -package=mock . IRequests
type IRequests interface {
	Do(method, toUrl string, opts ...Option) (resp *Response, err error)
	Post(url string, opts ...Option) (*Response, error)
	Put(url string, opts ...Option) (*Response, error)
	Delete(url string, opts ...Option) (*Response, error)
	Get(url string, opts ...Option) (*Response, error)
	Head(url string, opts ...Option) (*Response, error)
	Options(url string, opts ...Option) (*Response, error)
	Patch(url string, opts ...Option) (*Response, error)

	SetTimeout(t time.Duration)
	SetProxyUrl(proxyUrl string) error
	SetClient(client *http.Client)
}

type BeforeHandler func(r *http.Request)
type AfterHandler func(r *http.Response)

type Config struct {
	Timeout    time.Duration
	Transport  http.RoundTripper
	Headers    http.Header
	HttpClient *http.Client
}

type Requests struct {
	once          sync.Once
	client        *http.Client
	globalHeaders http.Header
	beforeHandler []BeforeHandler
	afterHandler  []AfterHandler
}

func NewDefault() *Requests {
	return &Requests{client: newDefaultClient()}
}

func NewFormHttpClient(client *http.Client) *Requests {
	return &Requests{client: client}
}

// NewFormConfig 若设置了http.client, 则transport, timeout的设置不会生效, 将以httpClient的设置为准
func NewFormConfig(cfg Config) *Requests {
	requests := &Requests{client: cfg.HttpClient}
	if requests.client == nil {
		requests.client = &http.Client{
			Transport: cfg.Transport,
			Timeout:   5 * cfg.Timeout,
		}
	}
	return requests
}

// create a default client
func newDefaultClient() *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}
}

func (r *Requests) Client() *http.Client {
	if r.client == nil {
		r.once.Do(func() {
			r.client = newDefaultClient()
		})
	}
	return r.client
}

func (r *Requests) SetTimeout(t time.Duration) {
	r.Client().Timeout = t
}

func (r *Requests) SetClient(client *http.Client) {
	r.client = client
}

// SetProxyUrl 设置http代理
func (r *Requests) SetProxyUrl(proxyUrl string) error {
	trans, _ := r.Client().Transport.(*http.Transport)
	if trans == nil {
		return ErrTransport
	}
	u, err := url.Parse(proxyUrl)
	if err != nil {
		return err
	}
	trans.Proxy = http.ProxyURL(u)
	return nil
}

// AddBeforeHandler 添加前置拦截器, 仅在请求处理前执行, 请求处理前发生error将不会执行, 多个将按顺序执行 todo 上下文
func (r *Requests) AddBeforeHandler(f BeforeHandler) {
	r.beforeHandler = append(r.beforeHandler, f)
}

// AddAfterHandler 添加后置拦截器, 仅在请求成功时执行, 多个将按顺序执行 todo 上下文
func (r *Requests) AddAfterHandler(f AfterHandler) {
	r.afterHandler = append(r.afterHandler, f)
}

// Do todo 链路追踪 requests id
func (r *Requests) Do(method, toUrl string, opts ...Option) (resp *Response, err error) {
	var (
		param        param
		responseData interface{}
		op           Op
	)
	// 应用配置
	op.applyOpts(opts)

	req := &http.Request{
		Method:     method,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}
	// 处理前置操作
	if op.ctx != nil {
		req = req.WithContext(op.ctx)
	}
	if op.proto != "" {
		req.Proto = op.proto
	}
	if op.header != nil {
		for key, value := range op.header {
			req.Header.Add(key, value)
		}
	}
	// todo param和form是否拆离
	if op.param != nil {
		if strings.ToUpper(method) == GetMethod || strings.ToUpper(method) == HeadMethod {
			param.Adds(op.param)
			paramStr := param.Encode()
			if strings.IndexByte(toUrl, '?') == -1 {
				toUrl = toUrl + "?" + paramStr
			} else {
				toUrl = toUrl + "&" + paramStr
			}
		} else {
			param.Adds(op.param)
			data := []byte(param.Encode())
			req.Body = ioutil.NopCloser(bytes.NewReader(data))
			req.ContentLength = int64(len(data))
		}
	}
	// 同时存在多个body时, 优先使用body -> json body -> xxx
	if op.body != nil {
		req.Body = ioutil.NopCloser(bytes.NewReader(op.body))
		req.ContentLength = int64(len(op.body))
	} else if op.jsonBody != nil {
		req.Header.Add("Content-Type", "application/json")
		data, err := json.Marshal(op.jsonBody)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(data))
		req.ContentLength = int64(len(data))
	}
	// 处理url
	u, err := url.Parse(toUrl)
	if err != nil {
		return nil, err
	}
	req.URL = u

	//if host := req.Header.Get("Host"); host != "" {
	//	req.Host = host
	//}
	// 执行前置拦截器
	for i := range r.beforeHandler {
		r.beforeHandler[i](req)
	}
	response, err := r.Client().Do(req)
	if err != nil {
		return nil, err
	}
	for i := range r.afterHandler {
		r.afterHandler[i](response)
	}
	// 执行前置拦截器
	result := &Response{HttpResponse: response}
	// 处理后置操作
	if op.respJsonBody != nil {
		// todo 此处error处理不对劲
		err = result.ToJSON(responseData)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (r *Requests) Post(url string, opts ...Option) (*Response, error) {
	return r.Do(PostMethod, url, opts...)
}

func (r *Requests) Get(url string, opts ...Option) (*Response, error) {
	return r.Do(GetMethod, url, opts...)
}

func (r *Requests) Put(url string, opts ...Option) (*Response, error) {
	return r.Do(PutMethod, url, opts...)
}

func (r *Requests) Delete(url string, opts ...Option) (*Response, error) {
	return r.Do(DeleteMethod, url, opts...)
}

func (r *Requests) Patch(url string, opts ...Option) (*Response, error) {
	return r.Do(PatchMethod, url, opts...)
}

func (r *Requests) Options(url string, opts ...Option) (*Response, error) {
	return r.Do(OptionsMethod, url, opts...)
}

func (r *Requests) Head(url string, v ...Option) (*Response, error) {
	return r.Do(HeadMethod, url, v...)
}
