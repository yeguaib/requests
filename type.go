package requests

import (
	"fmt"
	"net/url"
)

const (
	GetMethod     = "GET"
	PostMethod    = "POST"
	HeadMethod    = "HEAD"
	OptionsMethod = "OPTIONS"
	PutMethod     = "PUT"
	PatchMethod   = "PATCH"
	DeleteMethod  = "DELETE"
	TraceMethod   = "TRACE"
	ConnectMethod = "CONNECT"
)

type Header map[string]string

type Param map[string]interface{}

type param struct {
	url.Values
}

func (p *param) getValues() url.Values {
	if p.Values == nil {
		p.Values = make(url.Values)
	}
	return p.Values
}

func (p *param) Adds(m map[string]interface{}) {
	if len(m) == 0 {
		return
	}
	vs := p.getValues()
	for k, v := range m {
		vs.Add(k, fmt.Sprint(v))
	}
}
