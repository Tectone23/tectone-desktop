package docker

import (
	"github.com/docker/docker/client"
)

// getcli creates a new client and returns pointer to that client
func getCli() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	return cli, nil
}
