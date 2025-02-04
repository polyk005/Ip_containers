package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/polyk005/Ip_containers/back/internal/model"
)

type containerPostgres struct {
	db *sqlx.DB
}

func NewContainerPostgres(db *sqlx.DB) *containerPostgres {
	return &containerPostgres{db: db}
}

func (r *containerPostgres) GetContainers() ([]model.Container, error) {
	var containers []model.Container
	query := `SELECT id, ip_address, ping_time, last_successful_ping FROM containers`
	err := r.db.Select(&containers, query)
	if err != nil {
		return nil, err
	}
	return containers, nil
}

func (r *containerPostgres) AddContainer(ipAddress, pingTime, lastSuccessfulPing string) error {
	query := `INSERT INTO containers (ip_address, ping_time, last_successful_ping) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, ipAddress, pingTime, lastSuccessfulPing)
	return err
}
