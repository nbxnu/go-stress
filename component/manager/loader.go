package manager

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-stress/component/model"
	"io/ioutil"
)


type RequestPkg struct {
	NN      string          `json:"type"`
	Request model.Request   `json:"request"`
	Header  []*model.Header `json:"header"`
	Data    []*model.Data   `json:"data"`
}

type Loader struct {
	RequestPkgs []*RequestPkg
}

func (loader *Loader) Init(path string) error{
	var(
		binary []byte
		err error
	)

	if binary, err = ioutil.ReadFile(path); err != nil{
		return err
	}
	if err = json.Unmarshal(binary, &loader.RequestPkgs); err != nil{
		return err
	}

	return nil
}

func (loader *Loader)MakeRequest(repeat int) []*model.RequestConfig {
	req := loader.makeRequest()
	requests := make([]*model.RequestConfig, 0, len(req)*repeat)
	for i := 0; i < repeat; i++ {
		requests = append(requests, req...)
	}
	return requests
}

func (loader *Loader) makeRequest()[]*model.RequestConfig {
	requests := make([]*model.RequestConfig, 0, loader.CalcMem())
	for _, pkg := range loader.RequestPkgs {
		switch pkg.NN {
		case "nn":
			requests = append(requests, loader.makeNN(pkg)...)
		case "1n":
			requests = append(requests, loader.make1N(pkg)...)
		}
	}
	return requests
}

func (loader *Loader) makeNN(pkg *RequestPkg)[]*model.RequestConfig{
	requests := make([]*model.RequestConfig, 0, loader.CalcMemAt(pkg))
	for _, h := range pkg.Header{
		for _, d := range pkg.Data{
			requests = append(requests, &model.RequestConfig{
				ID: int64(uuid.New().ID()),
				Url:    pkg.Request.Url,
				Method: pkg.Request.Method,
				Header: h.Header,
				Data:   d.Data,
			})
		}
	}
	return requests
}

func (loader *Loader) make1N(pkg *RequestPkg)[]*model.RequestConfig{
	requests := make([]*model.RequestConfig, 0, loader.CalcMemAt(pkg))
	for _, h := range pkg.Header{
		for _, d := range pkg.Data{
			if d.HeaderID == h.ID{
				requests = append(requests, &model.RequestConfig{
					ID: int64(uuid.New().ID()),
					Url:    pkg.Request.Url,
					Method: pkg.Request.Method,
					Header: h.Header,
					Data:   d.Data,
				})
			}
		}
	}
	return requests
}

func (loader *Loader) CalcMem() int{
	var count = 0
	for _, pkg := range loader.RequestPkgs{
		count += loader.CalcMemAt(pkg)
	}
	return count
}

func (loader *Loader) CalcMemAt(pkg *RequestPkg)int{
	switch pkg.NN {
	case "nn":
		return len(pkg.Header) * len(pkg.Data)
	case "1n":
		return len(pkg.Data)
	}
	return 0
}