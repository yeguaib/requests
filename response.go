package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	HttpResponse *http.Response
	BodyData     []byte
	Err          error
}

func (r *Response) ToBytes() ([]byte, error) {
	if r.BodyData != nil {
		return r.BodyData, nil
	}
	defer r.HttpResponse.Body.Close()
	respBody, err := ioutil.ReadAll(r.HttpResponse.Body)
	if err != nil {
		return nil, err
	}
	r.BodyData = respBody
	return respBody, nil
}

func (r *Response) ToJSON(v interface{}) error {
	data, err := r.ToBytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (r *Response) StatusCode() int {
	return r.HttpResponse.StatusCode
}
