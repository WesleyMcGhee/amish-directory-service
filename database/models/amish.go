package models

type Amish struct {
	ID      string `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode uint   `json:"zip_code"`
	NoFly   bool   `json:"no_fly"`
}