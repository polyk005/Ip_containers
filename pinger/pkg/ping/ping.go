package ping

import (
	"time"

	"github.com/go-ping/ping"
)

func Ping(ip string) bool {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return false
	}

	pinger.Count = 1
	pinger.Timeout = time.Second * 2
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil {
		return false
	}

	return pinger.Statistics().PacketsRecv > 0
}
