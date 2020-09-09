package tests

import (
	"go-stress/component/manager"
	"runtime"
)

func Demo(file string, core, freq int) {
	runtime.GOMAXPROCS(core)

	loader := &manager.Loader{}
	if err := loader.Init(file); err != nil {
		panic(err)
	}
	requests := loader.MakeRequest(freq)

	mgr := &manager.Manager{}
	mgr.Init(core, freq)
	reports := mgr.Start(requests)

	reporter := &manager.Reporter{}
	reporter.Init(core, reports)
	reporter.Report()
}
