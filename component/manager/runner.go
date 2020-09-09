package manager

import (
	"go-stress/component/model"
	"net/http"
	"time"
)

type Runner struct {
	Client    *http.Client
}

func NewRunner() *Runner{
	runner := &Runner{
		Client:  http.DefaultClient,
	}
	return runner
}

func (runner *Runner) Do(requests []*model.RequestConfig) (reports []*model.ReportConfig){
	reports = make([]*model.ReportConfig, 0, len(requests))
	for _, req := range requests {
		report := runner.do(req)
		reports = append(reports, report)
	}
	return
}

func (runner *Runner) do(req *model.RequestConfig)*model.ReportConfig {
	resp := &model.ReportConfig{
		ID:         req.ID,
		StartAt:    0,
		CostTime:   0,
		StatusCode: 0,
	}
	start := time.Now().UnixNano()
	request, err := req.NewHttpRequest()
	if err != nil {
		return resp
	}
	response, err := runner.Client.Do(request)
	end := time.Now().UnixNano()
	resp.StartAt = start
	resp.EndAt = end
	resp.CostTime = end - start
	if err != nil {
		return resp
	}
	status := response.StatusCode
	defer response.Body.Close()
	resp.StatusCode = status
	return resp
}

