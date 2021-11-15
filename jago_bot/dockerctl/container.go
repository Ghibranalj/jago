package dockerctl

import (
	"context"
	"fmt"

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

// var (
// 	dPort    = 1234
// 	usedPort = map[int]bool{}
// )

// func getPort() string {
// 	port := dPort

// 	nFound := true
// 	for nFound {
// 		_, nFound = usedPort[port]
// 		// fmt.Println(port, found)
// 		port++
// 	}
// 	usedPort[port] = true

// 	return strconv.Itoa(port)
// }

func SpinUp() (string, string, error) {

	port := "1234"

	fmt.Println(port)

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

func Kill(id string) error {
	return cli.ContainerKill(ctx, id, "KILL")
}
