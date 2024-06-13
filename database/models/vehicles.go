package models

type Vehicles struct {
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}