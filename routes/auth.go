package routes

import (
	"amish-directory/database"
	"amish-directory/database/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword))

	return err == nil
}

func Signup(c echo.Context) error {
	password := c.FormValue("password")
	email := c.FormValue("email")
	username := c.FormValue("username")

	hashedPassword, err := HashPassword(password)

	if err != nil {
		c.Logger().Fatal("Password Hashing Failed")
	}

	user := models.Users{Username: username, PasswordHash: hashedPassword, Email: email,}

	database.DB.Create(&user)

	return c.String(http.StatusCreated, "User Created")
}

func Login(c echo.Context) error {
	var user models.Users

	username := c.FormValue("username")
	password := c.FormValue("password")

	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.String(http.StatusBadRequest, "User not found")
	}

	result := CheckPasswordHash(password, user.PasswordHash)

	if !result {
		return c.String(http.StatusForbidden, "Password or Username is Incorrect")
	}

	return c.JSON(http.StatusOK, user)
}