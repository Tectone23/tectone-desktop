package docker

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"tectone/tectone-desktop/internal/model"

	"github.com/docker/docker/api/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const listByProjectNameQuery = "'label=com.docker.compose.project=%s'"

// RemoveContainer removes the container by the given ID
func RemoveContainer(ctx context.Context, containerId string) error {
	cli, err := getCli()
	if err != nil {
		return err
	}
	defer cli.Close()

	options := types.ContainerRemoveOptions{RemoveVolumes: true, RemoveLinks: true, Force: true}
	return cli.ContainerRemove(ctx, containerId, options)
}

// StartContainer starts an existing container by the given ID
func StartContainer(ctx context.Context, containerId string) error {
	cli, err := getCli()
	if err != nil {
		return err
	}
	defer cli.Close()

	return cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{})
}

// StopContainer stops an existing container with the given ID
func StopContainer(ctx context.Context, containerId string) error {
	cli, err := getCli()
	if err != nil {
		return err
	}
	defer cli.Close()

	return cli.ContainerStart(ctx, containerId, types.ContainerStartOptions{})

}

func StreamContainerLogs(ctx context.Context, p model.Project) {

	names := GenerateContainerNamesFromHash(p.KeyHash)

	containers, err := GetContainersByNames(ctx, names)
	if err != nil {
		runtime.LogErrorf(ctx, "could not get containers by their names(%s): %s", names, err)
		return
	}

	runtime.LogDebug(ctx, "Looping over containers")
	var wg sync.WaitGroup
	for _, c := range containers {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ctx context.Context, c model.Container) {
			rc, err := GetContainerLogs(ctx, c.ID)
			if err != nil {
				runtime.LogErrorf(ctx, "could not get container logs: %s", err)
				return
			}
			defer rc.Close()

			s := bufio.NewScanner(rc)

			s.Split(bufio.ScanLines)
			for s.Scan() {
				line := s.Text()

				runtime.EventsEmit(ctx, "project-container-logs", line)
				runtime.LogDebugf(ctx, "%s", line)
			}

			wg.Done()

		}(&wg, ctx, c)

		wg.Wait()
	}
	runtime.LogDebug(ctx, "finished going over container logs")

}

// GetContainerLogs gets return a io.ReadCloser from which the logs for the container can be read
func GetContainerLogs(ctx context.Context, containerId string) (io.ReadCloser, error) {
	cli, err := getCli()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	options := types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true, Details: true}

	return cli.ContainerLogs(ctx, containerId, options)
}

// GetContainersByName gets the containers that match the names in the slice
func GetContainersByNames(ctx context.Context, names []string) ([]model.Container, error) {
	cli, err := getCli()
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	conts := []model.Container{}

	runtime.LogDebug(ctx, "getting containers list....")
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	runtime.LogDebugf(ctx, "found %d containers .... ", len(containers))
	for _, c := range containers {
		for _, n := range names {
			runtime.LogDebugf(ctx, "container name (%s) and arg name (%s)", c.Names[0], n)
			name := strings.TrimPrefix(c.Names[0], string(os.PathSeparator))
			if n == name {
				ctr := model.Container{
					ID:     c.ID,
					Name:   name,
					Size:   c.SizeRw,
					Status: c.State,
				}

				conts = append(conts, ctr)
				runtime.LogDebugf(ctx, "container matched:\n\nID: %s\nName: %v\nSize: %d\nStatus: %s", ctr.ID, ctr.Name, ctr.Size, ctr.Status)
			}
		}

	}

	runtime.LogDebug(ctx, "finished getting containers by names")

	return conts, nil
}

func ListContainersByProjectName(ctx context.Context, path, projectName, eventName string) {
	projectName = strings.ToLower(projectName)

	runtime.LogDebugf(ctx, "Fetching containers by Project Name \n\n\n path: %s\n projectName: %s\n eventName: %s", path, projectName, eventName)

	filter := fmt.Sprintf(listByProjectNameQuery, projectName)

	cmd := exec.Command("docker", "ps", "-a", "--format", `'{"ID":"{{ .ID }}", "Image": "{{ .Image }}", "Names":"{{ .Names }}"}'`, "--filter", filter)
	runtime.LogDebugf(ctx, "docker command: %s", cmd.Path)
	runtime.LogDebugf(ctx, "docker command args %v:", cmd.Args)
	cmd.Env = os.Environ()
	cmd.Dir = path

	r, err := cmd.StdoutPipe()
	if err != nil {
		runtime.LogErrorf(ctx, "could not get stdoutPipe: %s", err)
		return
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	if err := cmd.Start(); err != nil {
		runtime.LogErrorf(ctx, "could not start command 'docker ps -a': %s", err)
		return
	}

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line == `\r\n` {
			continue
		}
		runtime.LogDebugf(ctx, "ListContainersByName: %s", line)
		runtime.EventsEmit(ctx, eventName, line)
	}

	if err := cmd.Wait(); err != nil {
		runtime.LogErrorf(ctx, "failed waiting for command 'docker ps -a': %s", err)
		return
	}

	runtime.LogDebug(ctx, "completed listing containers")

}
func GenerateContainerNamesFromHash(hash string) []string {

	algod := fmt.Sprintf("algod-sandbox-%s", hash)
	conduit := fmt.Sprintf("conduit-sandbox-%s", hash)
	indexer := fmt.Sprintf("indexer-sandbox-%s", hash)
	postgres := fmt.Sprintf("postgres-sandbox-%s", hash)
	names := []string{algod, conduit, indexer, postgres}

	return names
}

// GetContainerStatus retrieves the status of the container who matches the ID given
func GetContainerStatus(ctx context.Context, containerId string) (string, error) {

	cli, err := getCli()
	if err != nil {
		return "", err

	}
	defer cli.Close()

	resp, err := cli.ContainerInspect(ctx, containerId)
	if err != nil {
		return "", err
	}

	return resp.State.Status, nil
}

func GetContainerList(ctx context.Context) ([]model.Container, error) {
	list := []model.Container{}
	cli, err := getCli()
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}
	for _, container := range containers {
		cont := model.Container{
			ID:     container.ID,
			Name:   strings.Join(container.Names, " "),
			Size:   container.SizeRw,
			Status: container.Status,
		}
		list = append(list, cont)
	}

	return list, nil
}
