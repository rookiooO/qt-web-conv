package model

import "time"

type Model struct {
	CreateTime time.Time
	UpdateTime time.Time
	CreateBy   int64
	UpdateBy   int64
	Status     int16
	Sort       int32
	IsDelete   bool
	Remark     string
}
