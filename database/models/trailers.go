package models

type Trailers struct {
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}