package model

import "time"

type Model struct {
	CreateTime time.Time
	UpdateTime time.Time
	CreateBy   string
	UpdateBy   string
	Status     string
	Sort       int32
	DelFlag    string
	Remark     string
}
