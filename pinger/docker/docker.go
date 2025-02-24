package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetContainers(cli *client.Client) ([]types.Container, error) {
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}
	return containers, nil
}
