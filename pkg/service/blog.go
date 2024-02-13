package service

import (
	"store/models"
	"store/pkg/repository"
)

type BlogService struct {
	repo repository.Blog
}

func NewBlogService(repo repository.Blog) *BlogService {
	return &BlogService{repo: repo}
}

func (s *BlogService) Create(blog models.CreateBlogParams) (uint, error) {
	return s.repo.Create(blog)
}

func (s *BlogService) Get(page int, limit int) ([]models.Blog, int64) {
	return s.repo.Get(page, limit)
}

func (s *BlogService) Delete(ID uint) {
	s.repo.Delete(ID)
}
