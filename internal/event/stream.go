package event

import (
	"bufio"
	"context"
	"io"
	"os/exec"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func Stream(ctx context.Context, cmd *exec.Cmd, eventName string) error {

	// execute and get a pipe
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	s := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for s.Scan() {
		runtime.EventsEmit(ctx, eventName, string(s.Bytes()))
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil

}
