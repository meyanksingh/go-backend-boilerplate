package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/meyanksingh/vlink-backend/internal/app/models"
	db "github.com/meyanksingh/vlink-backend/internal/db"
	"golang.org/x/crypto/bcrypt"
)

func CheckEmailExists(email string) (bool, error) {
	var existingUser models.User
	if err := db.DB.Where("email = ?", email).First(&existingUser).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil
		}
		return false, err
	} else {
		return true, nil
	}

}

func CreateUser(firstName, lastName, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("could not hash password")
	}

	user := models.User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  string(hashedPassword),
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return nil, errors.New("could not create user")
	}

	return &user, nil
}

func AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
