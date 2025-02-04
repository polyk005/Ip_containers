package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/polyk005/Ip_containers/back/internal/model"
)

type Authorization interface {
}

type ContainerRepository interface {
	GetContainers() ([]model.Container, error)
	AddContainer(ipAddress, pingTime, LastSuccessfulPing string) error
}

type Repository struct {
	Authorization
	Container ContainerRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Container:     NewContainerPostgres(db),
	}
}
