package model

import "time"

type Model struct {
	Id       int64 `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
	CreateBy int64
	UpdateBy int64
	Status   int16
	Sort     int32
	IsDelete bool
	Remark   string
}
