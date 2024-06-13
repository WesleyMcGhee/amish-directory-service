package models

type Roles struct {
	ID          string  `json:"id" gorm:"primary_key"`
	Name        string  `json:"name"`
	Permissions string  `json:"permissions"`
	Users       []Users `gorm:"foreignKey:RoleId"`
}