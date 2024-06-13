package models

import "time"

type Events struct {
	ID     string `json:"id" gorm:"primary_key"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	Notes string `json:"notes"`
	Driver Drivers `gorm:"foreignKey:ID"`
	Amish Amish `gorm:"foreignKey:ID"`
	Vehicle Vehicles `gorm:"foreignKey:ID"`
	Trailer  Trailers `gorm:"foreignKey:ID"`
}