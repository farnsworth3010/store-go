package repository

import (
	"store/models"

	"gorm.io/gorm"
)

type PanelPostgres struct {
	db *gorm.DB
}

func NewPanelPostgres(db *gorm.DB) *PanelPostgres {
	return &PanelPostgres{db: db}
}

func (r *PanelPostgres) GetAdmins() ([]models.Admin, error) {
	var admins []models.Admin
	res := r.db.Table("users").Select([]string{"ID", "email", "firstname", "phone_number", "role_id"}).Where("role_id = 1").Find(&admins)
	if res.Error != nil {
		return admins, res.Error
	}
	return admins, nil
}

func (r *PanelPostgres) GetUsers() ([]models.ShortUser, error) {
	var users []models.ShortUser
	res := r.db.Table("users").Select([]string{
		"ID", "email", "firstname", "phone_number", "role_id"}).Find(&users)
	if res.Error != nil {
		return users, res.Error
	}
	return users, nil
}

func (r *PanelPostgres) GetBlogs() ([]models.ShortBlog, error) {
	var blogs []models.ShortBlog
	res := r.db.Table("blogs").Select([]string{"id", "title"}).Find(&blogs)
	if res.Error != nil {
		return blogs, res.Error
	}
	return blogs, nil
}

func (r *PanelPostgres) SetRole(ID uint, RoleID uint) error {
	res := r.db.Model(&models.User{}).Where("id = ?", ID).Update("role_id", RoleID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
