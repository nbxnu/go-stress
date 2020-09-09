package manager

import (
	"fmt"
	"go-stress/component/model"
	"sort"
)

type Reporter struct {
	Core    int
	Reports []*model.ReportConfig
}

func (reporter *Reporter) Init(core int, reports []*model.ReportConfig){
	reporter.Core = core
	reporter.Reports = reports
	sort.Slice(reporter.Reports, func(i, j int)bool{
		return reporter.Reports[i].StartAt < reporter.Reports[j].StartAt
	})
}

type Statistics struct {
	RequestPerSecond   int
	RequestProcessTime int64
}

func (reporter *Reporter) Report() {
	global := &Statistics{
		RequestPerSecond:   0,
		RequestProcessTime: 0,
	}
	recorders := map[int64]*Statistics{}
	fmt.Printf("%11s|%5s|%10s|%10s|\n", "time", "qps", "sum", "avg")
	for _, r := range reporter.Reports {
		global.RequestPerSecond += 1
		global.RequestProcessTime += r.CostTime

		idx := r.StartAt / 1e9
		s := recorders[idx]
		if s == nil{
			recorders[idx] = &Statistics{}
		}
		recorders[idx].RequestPerSecond += 1
		recorders[idx].RequestProcessTime += r.CostTime
	}


	last := int64(0)
	for _, r := range reporter.Reports{
		idx := r.StartAt/1e9
		if idx == last{
			continue
		}
		last = idx
		reporter.fmtStatistics(idx, recorders[idx])
	}
	reporter.fmtStatistics(0, global)
}

func (reporter *Reporter) fmtStatistics(idx int64, s *Statistics) {
	sum := float64(s.RequestProcessTime) / float64(reporter.Core) / 1e6
	avg := float64(0)
	if s.RequestPerSecond != 0 {
		avg = sum / float64(s.RequestPerSecond)
	}
	fmt.Printf("%11d|%5d|%10.4f|%10.4f|\n",
		idx,
		s.RequestPerSecond,
		sum,
		avg,
	)
}
