package requests

import "context"

type Op struct {
	ctx          context.Context
	body         []byte
	header       map[string]string
	param        map[string]interface{}
	jsonBody     interface{}
	respJsonBody interface{}
	proto        string
}

type Option func(*Op)

func (o *Op) applyOpts(opts []Option) {
	for i := range opts {
		opts[i](o)
	}
}

func WithCtx(ctx context.Context) Option {
	return func(op *Op) {
		op.ctx = ctx
	}
}

func WithBody(body []byte) Option {
	return func(op *Op) {
		op.body = body
	}
}

func WithBodyJson(body interface{}) Option {
	return func(op *Op) {
		op.jsonBody = body
	}
}

func WithRespBodyJson(resp interface{}) Option {
	return func(op *Op) {
		op.respJsonBody = resp
	}
}

func WithHeader(header map[string]string) Option {
	return func(op *Op) {
		op.header = header
	}
}

func WithParam(param map[string]interface{}) Option {
	return func(op *Op) {
		op.param = param
	}
}

func WithProto(proto string) Option {
	return func(op *Op) {
		op.proto = proto
	}
}
