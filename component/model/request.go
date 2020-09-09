package model

import (
	"bytes"
	"net/http"
)

type Request struct {
	ID     int64  `json:"id"`
	Url    string `json:"url"`
	Method string `json:"method"`
}

type Header struct {
	ID        int64             `json:"id"`
	Header    map[string]string `json:"header"`
}

type Data struct {
	ID       int64  `json:"id"`
	HeaderID int64  `json:"header_id"`
	Data     string `json:"data"`
}

type RequestConfig struct{
	ID int64
	Url string
	Method string
	Header map[string]string
	Data string
}

func (r *RequestConfig) NewHttpRequest() (*http.Request, error) {
	req, err := http.NewRequest(r.Method, r.Url, bytes.NewBufferString(r.Data))
	if err != nil {
		return nil, err
	}
	for k, v := range r.Header {
		req.Header.Set(k, v)
	}
	return req, nil
}