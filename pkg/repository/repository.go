package repository

import (
	"github.com/jmoiron/sqlx"
	"store/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Blog interface {
	Get(page int, limit int) ([]models.Blog, int, error)
	Create(blog models.CreateBlogParams) (int, error)
}

type Repository struct {
	Authorization
	Blog
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db), Blog: NewBlogPostgres(db)}
}
