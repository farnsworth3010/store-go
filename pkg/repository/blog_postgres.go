package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"store/models"
	"strconv"
)

type BlogPostgres struct {
	db *sqlx.DB
}

func NewBlogPostgres(db *sqlx.DB) *BlogPostgres {
	return &BlogPostgres{db: db}
}

func (r *BlogPostgres) Create(blog models.CreateBlogParams) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, text) VALUES ($1, $2) RETURNING id", blogTable)
	row := r.db.QueryRow(query, blog.Title, blog.Text)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BlogPostgres) Get(page int, limit int) ([]models.Blog, int, error) {
	var blog []models.Blog
	total := CountRows(blogTable, r.db)

	query := fmt.Sprintf("SELECT * FROM %s ORDER BY creation_date DESC LIMIT "+strconv.Itoa(limit)+" OFFSET "+strconv.Itoa(limit*page), blogTable)
	err := r.db.Select(&blog, query)

	return blog, total, err
}
