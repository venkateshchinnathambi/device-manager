package models

import "gorm.io/gorm"

type IndoorTemperature struct {
	gorm.Model
	DeviceID  string
	Value     string
	Scale     string
	Timestamp int64
}
