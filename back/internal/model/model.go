package model

type Container struct {
	ID                 int    `db:"id"`
	IPAddress          string `db:"ip_address"`
	PingTime           string `db:"ping_time"`
	LastSuccessfulPing string `db:"last_successful_ping"`
}
