package dockerctl

import (
	"context"
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

var (
	ctx context.Context
	cli *client.Client
)

func Init() {

	ctx = context.Background()

	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

}

var usedPort = map[int]bool{}

func getPort() string {
	port := 1234

	i := 0

	for {
		_, ok := usedPort[port+i]
		if !ok {
			usedPort[port+i] = true
			return strconv.Itoa(port + i)
		}
		i++
	}
}

func SpinUp() (string, string, error) {

	port := getPort()
	fmt.Println(usedPort)
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "jago_compiler",
		ExposedPorts: nat.PortSet{
			"1234/tcp": struct{}{},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"1234/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
	}, nil, nil, "")

	if err != nil {
		return "", "", err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", "", err
	}

	return port, resp.ID, nil
}

func Kill(port, id string) error {
	iPort, _ := strconv.Atoi(port)
	delete(usedPort, iPort)
	return cli.ContainerKill(ctx, id, "KILL")
}
