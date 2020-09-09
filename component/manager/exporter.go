package manager

import (
	"encoding/json"
	"go-stress/component/model"
	"os"
)

type Exporter struct{
	File *os.File
}

func (exporter *Exporter) Init(path string)(err error){
	exporter.File, err = os.Create(path)
	if err != nil{
		return err
	}
	return nil
}

func (exporter *Exporter) MakeReport(reports []*model.ReportConfig) error{
	encoder := json.NewEncoder(exporter.File)
	if err := encoder.Encode(reports); err != nil{
		return err
	}
	return nil
}