package model

type ReportConfig struct {
	ID         int64  `json:"id"`
	StartAt    int64  `json:"start_at"`
	EndAt int64 `json:"end_at"`
	CostTime   int64  `json:"cost"`
	StatusCode int    `json:"status"`
	Data       string `json:"data"`
}