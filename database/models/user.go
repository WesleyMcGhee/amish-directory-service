package models

type Users struct {
	ID           string `json:"id" gorm:"primary_key"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	RoleId       uint   `json:"role_id"`
	DriverId     uint   `json:"driver_id"`
}