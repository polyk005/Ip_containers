package model

type ContainerStatus struct {
	IPAddress      string `db:"ip_address"`
	PingTime       int64  `db:"ping_time"`
	LastSuccessful string `db:"last_successful"`
}
