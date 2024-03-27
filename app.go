package main

import (
	"context"
	"fmt"
	"os/exec"
	"tectone/tectone-desktop/internal/model"
	"tectone/tectone-desktop/internal/service/zoho"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context

	r []registerer
}

// NewApp creates a new App application struct
func NewApp(r ...registerer) *App {
	return &App{
		r: r,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	for _, r := range a.r {
		r.register(ctx)
	}

}

// shutdown is called when the app shuts down.
// We can do any cleanup required here
func (a *App) shutdown(ctx context.Context) {
	return
}

// GetAppInfo returns application information
// consisting of version, commit hash, and build date.
func (a *App) GetAppInfo() model.AppInfo {
	return model.AppInfo{
		Version:    VERSION,
		CommitHash: COMMIT_HASH,
		BuildDate:  BUILD_DATE,
	}
}

// OpenDir triggers a dialog to open in the UI for the user to select a directory and
// returns the path of the directory selected by the user in the UI dialog
func (a *App) OpenDir() string {
	options := runtime.OpenDialogOptions{}

	path, err := runtime.OpenDirectoryDialog(a.ctx, options)
	if err != nil {
		runtime.LogErrorf(a.ctx, "could not open directory dialog: %s\n", err)
		return ""
	}

	runtime.LogDebugf(a.ctx, "opened the directory: %s\n", path)

	return path
}

func (a *App) OpenFile() string {
	options := runtime.OpenDialogOptions{}
	path, err := runtime.OpenFileDialog(a.ctx, options)
	if err != nil {
		estr := fmt.Sprintf("could not open file dialog: %s", err)
		runtime.LogError(a.ctx, estr)
		runtime.EventsEmit(a.ctx, "error", estr)
		return ""
	}

	runtime.LogDebugf(a.ctx, "opened the file: %s", path)
	return path
}

// CheckDependencies checks if the git and docker are found in the PATH
func (a *App) CheckDependencies() {
	_, err := exec.LookPath("git")
	if err != nil {
		runtime.LogErrorf(a.ctx, "could not find git in path: %s\n", err)
		runtime.EventsEmit(a.ctx, "dependency-not-found", "git")
	}

	_, err = exec.LookPath("docker")
	if err != nil {
		runtime.LogErrorf(a.ctx, "could not find docker in path: %s\n", err)
		runtime.EventsEmit(a.ctx, "dependency-not-found", "docker")
	}

	return
}

// SubmitFeedback submits feedback to Zoho ticketing system
func (a *App) SubmitFeedback(f model.Feedback) bool {

	z := zoho.New()
	if err := z.SubmitFeedback(&f); err != nil {
		estr := fmt.Sprintf("could not submit feedback: %s", err)
		runtime.LogError(a.ctx, estr)
		runtime.EventsEmit(a.ctx, "error", estr)
		return false
	}

	istr := "sucessfully submitted feedback!"
	runtime.LogDebug(a.ctx, istr)
	runtime.EventsEmit(a.ctx, "info", istr)

	return true
}
