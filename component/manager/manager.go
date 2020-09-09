package manager

import (
	"go-stress/component/model"
	"math/rand"
)

type Manager struct {
	Core      int
	Frequency int
	Runners   []*Runner
}

func (mgr *Manager) Init(core, freq int){
	mgr.Core = core
	mgr.Frequency = freq
	mgr.Runners = make([]*Runner, mgr.Core)
	for idx, _ := range mgr.Runners {
		mgr.Runners[idx] = NewRunner()
	}
}

func (mgr *Manager) Start(requests []*model.RequestConfig) []*model.ReportConfig {
	reportsChn := make(chan []*model.ReportConfig, mgr.Core)
	mgr.dispatch(requests,reportsChn)
	return mgr.collection( reportsChn)
}

func (mgr *Manager) dispatch(requests []*model.RequestConfig, chn chan<- []*model.ReportConfig) {
	rand.Shuffle(len(requests), func(i, j int) {
		requests[i], requests[j] = requests[j], requests[i]
	})
	part := len(requests) / mgr.Core
	for i, runner := range mgr.Runners {
		sli := requests[i*part : (i+1)*part]
		go func(runner *Runner, requests []*model.RequestConfig) {
			reports := runner.Do(requests)
			chn <- reports
		}(runner, sli)
	}
}

func (mgr *Manager) collection(chn <-chan []*model.ReportConfig) []*model.ReportConfig {
	reports := make([]*model.ReportConfig, 0)
	for i := 0; i < mgr.Core; i++ {
		rep := <-chn
		reports = append(reports, rep...)
	}
	return reports
}