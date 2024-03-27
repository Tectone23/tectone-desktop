package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

var binaryPath string

var (
	VERSION     = "vx.x.x-dev"
	COMMIT_HASH = "commit-xxxxxxx"
	BUILD_DATE  = time.Now().Format(time.DateOnly)
)

func main() {
	var err error

	binaryPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("can not determin the path of the binary: %s", err)
	}

	// Create an instance of the app structure
	sdk := newSDK()
	docker := newDocker()
	project := newProject()
	app := NewApp(project, docker, sdk)

	// Create application with options
	err = wails.Run(&options.App{
		Title: "Tectone SDK",
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  true,
				// HideTitleBar               bool
				// FullSizeContent            bool
				// UseToolbar                 bool
				// HideToolbarSeparator
			},
		},
		Windows:   &windows.Options{},
		Linux:     &linux.Options{},
		MinWidth:  1200,
		MinHeight: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
			project,
			docker,
			sdk,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
