package service

import (
	"github.com/polyk005/Ip_containers/back/internal/model"
	"github.com/polyk005/Ip_containers/back/pkg/api/repository"
)

type containerService struct {
	repo repository.ContainerRepository
}

func NewContainerService(repo repository.ContainerRepository) ContainerService { // Изменено на ContainerService
	return &containerService{repo: repo}
}

func (s *containerService) GetContainers() ([]model.Container, error) { // Изменено на model.Container
	return s.repo.GetContainers()
}

func (s *containerService) AddContainer(ipAddress, pingTime, lastSuccessfulPing string) error {
	return s.repo.AddContainer(ipAddress, pingTime, lastSuccessfulPing)
}
