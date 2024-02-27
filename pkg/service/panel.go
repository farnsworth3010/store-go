package service

import (
	"store/models"
	"store/pkg/repository"
)

type PanelService struct {
	repo repository.Panel
}

func NewPanelService(repo repository.Panel) *PanelService {
	return &PanelService{repo: repo}
}

func (s *PanelService) GetAdmins() ([]models.Admin, error) {
	return s.repo.GetAdmins()
}

func (s *PanelService) GetUsers() ([]models.ShortUser, error) {
	return s.repo.GetUsers()
}

func (s *PanelService) GetBlogs() ([]models.ShortBlog, error) {
	return s.repo.GetBlogs()
}

func (s *PanelService) SetRole(ID uint, RoleID uint) error {
	return s.repo.SetRole(ID, RoleID)
}
