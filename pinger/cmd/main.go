package main

import (
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/polyk005/Ip_containers/pinger/docker"
	"github.com/polyk005/Ip_containers/pinger/pkg/api"
	"github.com/polyk005/Ip_containers/pinger/pkg/ping"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}

	for {
		containers, err := docker.GetContainers(cli)
		if err != nil {
			log.Printf("Failed to fetch containers: %v", err)
			time.Sleep(10 * time.Second)
			continue
		}

		for _, container := range containers {
			ip := getContainerIPAddress(container)
			if ip == "" {
				continue
			}

			pingTime := time.Now().Format(time.RFC3339)
			lastSuccessfulPing := ""

			if ping.Ping(ip) {
				lastSuccessfulPing = pingTime
			}

			err := api.SendData(ip, pingTime, lastSuccessfulPing)
			if err != nil {
				log.Printf("Failed to send data to backend: %v", err)
			}
		}

		time.Sleep(15 * time.Second)
	}
}

func getContainerIPAddress(container types.Container) string {
	if networkSettings := container.NetworkSettings; networkSettings != nil {
		if network, exists := networkSettings.Networks["bridge"]; exists {
			return network.IPAddress
		}
	}
	return ""
}
