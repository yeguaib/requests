# requests

简单易用的网络请求封装，类似于python的requests

## 安装

`go get xxx`

## 使用

### 创建Client

可以通过以下两种方式创建一个client

首先需要使用new创建一个redis客户端实例,并传入redis主机,端口等配置参数

```go
// 通过默认配置创建client
r := requests.NewDefault()
// 通过配置创建client
r = requests.NewFormConfig(requests.Config{})
```

也可以不初始化直接使用，此时的client将使用默认配置，同`NewDefault`创建的client

```go
var r requests.Requests
r.Get("http://127.0.0.1:8000")
```

### 自定义http client

```GO
var (
    r requests.Requests
    myClient http.Client
)
r.SetClient(&myClient)
```

### 拦截器

可以对请求添加前置和后置拦截器, 拦截器将根据添加顺序执行

```go
var r requests.Requests
// 前置拦截
r.AddBeforeHandler(func(r *http.Request){})
// 后置拦截
r.AddAfterHandler(func(r *http.Response){})
```

### 发送请求

```go
var r requests.Requests
// 前置拦截
r.Get("http://127.0.0.1:8000")
// 使用默认的clientt发送请求
requests.Get("http://127.0.0.1:8000")
requests.Post("http://127.0.0.1:8000")
```

### 可选操作

更多操作见godoc

```python
requests.Get("http://127.0.0.1:8000", requests.WithHeader(map[string]string{
		"trace-id": "123",
	}), requests.WithParam(map[string]interface{}{
		"id": 1,
	}), requests.WithCtx(context.Background()),
		requests.WithProto("HTTP/2.0"),
	)
```

