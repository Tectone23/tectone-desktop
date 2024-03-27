package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"sync"
	"tectone/tectone-desktop/internal/event"
	"tectone/tectone-desktop/internal/model"
	"tectone/tectone-desktop/internal/service/docker"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Docker ...
type Docker struct {
	ctx context.Context
}

// StartContainer starts an existing container by the given name
func (d *Docker) StartContainer(ctx context.Context, containerId string) bool {
	if err := docker.StartContainer(d.ctx, containerId); err != nil {
		estr := fmt.Sprintf("could not start container of id %s: %s", containerId, err)
		runtime.LogError(d.ctx, estr)
		runtime.EventsEmit(d.ctx, "error", estr)
		return false
	}

	runtime.EventsEmit(d.ctx, "info", fmt.Sprintf("Successfully started container %s", containerId))
	return true
}

// StopContainer stops an existing container with the given name
func (d *Docker) StopContainer(ctx context.Context, containerId string) bool {
	if err := docker.StopContainer(d.ctx, containerId); err != nil {
		estr := fmt.Sprintf("could not stop container of id %s: %s", containerId, err)
		runtime.LogError(d.ctx, estr)
		runtime.EventsEmit(d.ctx, "error", estr)
		return false
	}

	runtime.EventsEmit(d.ctx, "info", fmt.Sprintf("Successfully stopped container %s", containerId))
	return true
}

// DockerComposeUp executes `docker-compose up` command in the given path
func (d *Docker) DockerComposeUp(project Project) bool {
	runtime.LogDebugf(d.ctx, "docker compose up project: %s", project.Name)
	eventName := fmt.Sprintf("compose-up-%s", project.ID)

	if err := docker.ComposeUp(d.ctx, project.SandboxPath, project.KeyHash, eventName); err != nil {
		estr := fmt.Sprintf("could not execute docker compose up: %s", err)
		runtime.LogError(d.ctx, estr)
		runtime.EventsEmit(d.ctx, "error", estr)
		return false
	}

	runtime.EventsEmit(d.ctx, "info", fmt.Sprintf("Successfully started all containers for project: %s", project.Name))
	return true
}

// DockerComposeDown executes `docker-compose down` command in the given path
func (d *Docker) DockerComposeDown(project Project) bool {
	runtime.LogDebugf(d.ctx, "docker compose down project: %s", project.Name)
	eventName := fmt.Sprintf("sandbox-down-%s", project.ID)

	if err := docker.ComposeDown(d.ctx, project.SandboxPath, eventName); err != nil {
		estr := fmt.Sprintf("could not execute docker compose down: %s", err)
		runtime.LogError(d.ctx, estr)
		runtime.EventsEmit(d.ctx, "error", estr)
		return false
	}

	runtime.EventsEmit(d.ctx, "info", fmt.Sprintf("Successfully stopped all containers for project: %s", project.Name))
	return true
}

// StreamContainerLogs streams the logs of all the containers of a the project to the GUI
// the logs are streamed on event name 'project-container-logs'
func (d *Docker) StreamContainerLogs(p Project) {

	runtime.LogDebug(d.ctx, "inside the stream containr logs")
	names := docker.GenerateContainerNamesFromHash(p.KeyHash)

	containers, err := docker.GetContainersByNames(d.ctx, names)
	if err != nil {
		runtime.LogErrorf(d.ctx, "could not get containers by their names(%s): %s", names, err)
		return
	}

	runtime.LogDebug(d.ctx, "Looping over containers")
	var wg sync.WaitGroup
	for _, c := range containers {
		wg.Add(1)

		go func(wg *sync.WaitGroup, d *Docker, c model.Container) {
			rc, err := docker.GetContainerLogs(d.ctx, c.ID)
			if err != nil {
				runtime.LogErrorf(d.ctx, "could not get container logs: %s", err)
				return
			}
			defer rc.Close()

			s := bufio.NewScanner(rc)

			s.Split(bufio.ScanLines)
			for s.Scan() {
				line := s.Text()

				runtime.EventsEmit(d.ctx, "project-container-logs", line)
				runtime.LogDebugf(d.ctx, "%s", line)
			}

			wg.Done()

		}(&wg, d, c)

		wg.Wait()
	}
	runtime.LogDebug(d.ctx, "finished going over container logs")
}

func (d *Docker) StreamComposeLogs(p Project) {
	cmd := exec.Command("docker", "compose", "logs")
	cmd.Env = os.Environ()
	cmd.Dir = p.SandboxPath

	if err := event.Stream(d.ctx, cmd, "project-container-logs"); err != nil {
		runtime.LogErrorf(d.ctx, "could not stream compose logs: %s", err)
	}
}

// ListContainers lists the containers available
func (d *Docker) ListContainers() []model.Container {

	list, err := docker.GetContainerList(d.ctx)
	if err != nil {
		runtime.LogErrorf(d.ctx, "could not get list of all containers: %s", err)
		return nil
	}

	return list
}

// GetContainerStatus retrieves the container status
func (d *Docker) GetContainerStatus(containerId string) string {

	status, err := docker.GetContainerStatus(d.ctx, containerId)
	if err != nil {
		runtime.LogErrorf(d.ctx, "could not get container status for container with id (%s): %s", containerId, err)
		return ""
	}

	return status
}

// register registers the app context.Context
func (d *Docker) register(ctx context.Context) {
	d.ctx = ctx
}

// newDocker returns a new Docker
func newDocker() *Docker {
	return &Docker{}
}
