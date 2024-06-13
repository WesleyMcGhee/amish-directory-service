package models

type Drivers struct {
	ID     string `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Status string `json:"status"`
	User   Users  `gorm:"foreignKey:DriverId"`
}