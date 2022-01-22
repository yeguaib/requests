package requests

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"New",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewDefault()
		})
	}
}

func TestRequests_Do(t *testing.T) {
	type args struct {
		method string
		toUrl  string
		op     []Option
	}
	tests := []struct {
		name     string
		args     args
		wantResp int
		wantErr  bool
	}{
		{
			name: "POST",
			args: args{
				method: "POST",
				toUrl:  "https://www.baidu.com/",
				op:     nil,
			},
			wantResp: 200,
			wantErr:  false,
		},
		{
			name: "GET",
			args: args{
				method: "GET",
				toUrl:  "https://www.baidu.com/",
				op:     nil,
			},
			wantResp: 200,
			wantErr:  false,
		},
		{
			name: "GET请求头和参数测试",
			args: args{
				method: "GET",
				toUrl:  "https://www.baidu.com?a=a",
				op: []Option{
					WithHeader(Header{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"}),
					WithParam(Param{"a": "a"}),
				},
			},
			wantResp: 200,
			wantErr:  false,
		},
		{
			name: "POST请求头和参数测试",
			args: args{
				method: "POST",
				toUrl:  "https://www.baidu.com/",
				op: []Option{
					WithHeader(Header{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"}),
					WithParam(Param{"a": "a"}),
				},
			},
			wantResp: 200,
			wantErr:  false,
		},
		{
			name: "GET请求头和参数测试",
			args: args{
				method: "GET",
				toUrl:  "https://www.baidu.com/",
				op: []Option{
					WithHeader(Header{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36"}),
					WithBodyJson(struct {
						A string
					}{"test"}),
				},
			},
			wantResp: 200,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDefault()
			gotResp, err := r.Do(tt.args.method, tt.args.toUrl, tt.args.op...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp.HttpResponse.StatusCode, tt.wantResp) {
				t.Errorf("Do() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestRequests_Get(t *testing.T) {
	type args struct {
		url string
		v   []Option
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "POST",
			args: args{
				url: "https://www.baidu.com/",
				v:   nil,
			},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDefault()
			got, err := r.Get(tt.args.url, tt.args.v...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.HttpResponse.StatusCode, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequests_Post(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		url string
		v   []Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name:   "POST",
			fields: fields{},
			args: args{
				url: "https://www.baidu.com/",
				v:   nil,
			},
			want:    200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDefault()
			got, err := r.Post(tt.args.url, tt.args.v...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.HttpResponse.StatusCode, tt.want) {
				t.Errorf("Post() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRequests_SetTimeout(t *testing.T) {
	type args struct {
		method string
		toUrl  string
		op     []Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "超时",
			args: args{
				method: "POST",
				toUrl:  "https://www.facebook.com/",
				op:     nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDefault()
			r.SetTimeout(2 * time.Second)
			_, err := r.Do(tt.args.method, tt.args.toUrl, tt.args.op...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRequests_GetJson(t *testing.T) {
	type fields struct {
		client *http.Client
	}
	type args struct {
		url string
		v   []Option
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:   "Json响应结果",
			fields: fields{},
			args: args{
				url: "https://www.baidu.com/",
				v:   nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewDefault()
			got, err := r.Get(tt.args.url, tt.args.v...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			body, err := got.ToBytes()
			if (err != nil) != tt.wantErr && len(body) > 0 {
				t.Errorf("Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
