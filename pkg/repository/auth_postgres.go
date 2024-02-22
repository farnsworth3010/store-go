package repository

import (
	"gorm.io/gorm"
	"store/models"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (uint, error) {
	var newUser models.User = models.User{Email: user.Email, Firstname: user.Firstname, Password: user.Password, PhoneNumber: user.PhoneNumber}
	res := r.db.Create(&newUser)
	// FIX INFINITE REG
	if res.Error != nil {
		return 0, res.Error
	}
	return newUser.ID, nil
}

func (r *AuthPostgres) GetUser(email string, password string) (models.User, error) {
	var user models.User
	res := r.db.Where("email=? AND password=?", email, password).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, res.Error
}
func (r *AuthPostgres) GetUserInfo(ID uint) (models.User, error) {
	var user models.User
	res := r.db.Find(&user, ID)
	if res.Error != nil {
		return user, res.Error
	}
	return user, res.Error
}

func (r *AuthPostgres) DeleteUser(ID uint) error {
	res := r.db.Delete(&models.User{}, ID)
	if res != nil {
		return res.Error
	}
	return nil
}
