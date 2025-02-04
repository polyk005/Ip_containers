package service

import (
	"github.com/polyk005/Ip_containers/back/internal/model"
	"github.com/polyk005/Ip_containers/back/pkg/api/repository"
)

type Authorization interface {
}

type ContainerService interface {
	GetContainers() ([]model.Container, error)
	AddContainer(ipAddress, pingTime, lastSuccessfulPing string) error
}

type Service struct {
	Authorization
	Container ContainerService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Container:     NewContainerService(repos.Container),
	}
}
