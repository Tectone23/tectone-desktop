package docker

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"tectone/tectone-desktop/internal/event"

	"github.com/docker/docker/api/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var initPorts = map[string]int{"ALGOD_PORT": 4001, "KMD_PORT": 4002, "CDT_PORT": 9392, "ALGOD_FOLLOWER_PORT": 3999, "INDEXER_PORT": 8980, "CONDUIT_PORT": 3998, "POSTGRES_PORT": 5433}

func ComposeUp(ctx context.Context, path, keyHash, eventName string) error {

	assignedPorts := []int{}

	cmd := exec.Command("./sandbox", "up", "devnet", "-v")
	cmd.Env = os.Environ()
	// container names
	cmd.Env = append(cmd.Env, fmt.Sprintf("ALGOD_CONTAINER=algod-sandbox-%s", keyHash))
	cmd.Env = append(cmd.Env, fmt.Sprintf("INDEXER_CONTAINER=indexer-sandbox-%s", keyHash))
	cmd.Env = append(cmd.Env, fmt.Sprintf("CONDUIT_CONTAINER=conduit-sandbox-%s", keyHash))
	cmd.Env = append(cmd.Env, fmt.Sprintf("POSTGRES_CONTAINER=postgres-sandbox-%s", keyHash))

	// container ports
	for k, v := range initPorts {

		for {
		top:
			ok, err := Check(ctx, v)
			if ok {
				if contains := slices.Contains(assignedPorts, v); contains {
					v = v + 1
					goto top
				}

				assignedPorts = append(assignedPorts, v)
				cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, strconv.Itoa(v)))
				runtime.LogDebugf(ctx, "Found new value ---> \t %s=%s ", k, strconv.Itoa(v))
				break
			}
			// container ports
			runtime.LogDebugf(ctx, "TAKEN VALUES ---> \t %s=%s | err: %s", k, strconv.Itoa(v), err)
			v = v + 1

		}

	}

	cmd.Dir = path

	return event.Stream(ctx, cmd, eventName)

}

func ComposeDown(ctx context.Context, path, eventName string) error {
	cmd := exec.Command("./sandbox", "down", "-v")
	cmd.Env = os.Environ()
	cmd.Dir = path

	return event.Stream(ctx, cmd, eventName)

}

func InitCompose(ctx context.Context, composePath, projectId, hash string) error {
	if err := ComposeUp(ctx, composePath, hash, "new-project-stdout"); err != nil {
		return err

	}
	return nil

}

// Check if a port is available
func Check(ctx context.Context, port int) (status bool, err error) {

	runtime.LogDebugf(ctx, "checking port %d", port)
	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return false, err
	}

	// close the server
	server.Close()

	// we successfully used and closed the port
	// so it's now available to be used again
	return true, nil

}

func ComposeRemove(ctx context.Context, path string) error {
	cli, err := getCli()
	if err != nil {
		return err
	}

	return cli.ContainerRemove(ctx, "", types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         true,
	})
}
