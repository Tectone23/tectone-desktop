package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"tectone/tectone-desktop/internal/config"
	"tectone/tectone-desktop/internal/filesystem"
	"tectone/tectone-desktop/internal/genesis"
	"tectone/tectone-desktop/internal/model"
	"tectone/tectone-desktop/internal/service/docker"
	tgit "tectone/tectone-desktop/internal/service/git"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

const SDKGitURL = "https://github.com/algorand/go-algorand-sdk.git"
const SandboxGitURL = "https://github.com/algorand/sandbox.git"

type Project struct {
	ctx        context.Context
	configFile *os.File
	config     *config.ProjectsConfig
	model.Project
}

// New creates a new project
func (p *Project) New(name, dir string) string {

	path := filepath.Join(dir, strings.Replace(name, " ", "-", -1))
	if err := os.Mkdir(path, 0755); err != nil {
		estr := fmt.Sprintf("could not create dir (%s): %s\n", path, err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return ""
	}

	return path
}

func (p *Project) GetAllProjects() []model.Project {
	return p.config.Projects
}

func (p *Project) GetProjectByID(id string) model.Project {
	var proj model.Project

	for _, i := range p.config.Projects {
		if i.ID == id {
			return i
		}
	}

	return proj

}

func (p *Project) DeleteProjectByID(id string) bool {
	// delete project
	proj := p.GetProjectByID(id)

	// delete containers  and images
	if err := docker.ComposeRemove(p.ctx, proj.SandboxPath); err != nil {
		runtime.LogErrorf(p.ctx, "could not delete project containers and images: %s", err)
		return false
	}

	// delete project folders
	if err := filesystem.RemoveDir(proj.Path); err != nil {
		runtime.LogErrorf(p.ctx, "could not delete project folder: %s", err)
		return false
	}

	pc := config.ProjectsConfig{}

	for _, pr := range p.config.Projects {
		if pr.ID != id {
			pc.Projects = append(pc.Projects, pr)
		}
	}

	p.config = &pc

	b, err := json.Marshal(p.config)
	if err != nil {
		estr := fmt.Sprintf("could not marshal project configuration to json: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	n, err := filesystem.Save(p.ctx, config.ConfigFileName, b)
	if err != nil {
		estr := fmt.Sprintf("could not save project config to disk: %s", err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	runtime.EventsEmit(p.ctx, "info", fmt.Sprintf("Successfully deleted project %s", proj.Name))
	runtime.LogDebugf(p.ctx, "written configuration file to disk successfully, written %d bytes", n)

	return true
}

func (p *Project) SetupProject(name, wd string) bool {
	var entry model.Project

	runtime.LogDebug(p.ctx, "setting up the project")

	// Fetch Algorand SDK
	sdkWd := path.Join(wd, "algorand-sdk")
	if !p.FetchSdk(sdkWd) {
		return false
	}
	runtime.LogDebug(p.ctx, "got the algorand sdk")

	// Fetch Sandbox
	hash := randomHash(10)
	entry.KeyHash = hash
	sandboxWd := path.Join(wd, strings.Join([]string{"algorand-sandbox", hash}, "-"))
	if !p.FetchSandbox(sandboxWd) {
		return false
	}
	runtime.LogDebug(p.ctx, "got the algorand sandbox")

	entry.ID = uuid.New().String()
	entry.Name = name
	entry.Path = wd
	entry.SDKPath = sdkWd
	entry.SandboxPath = sandboxWd
	entry.Containers = []model.Container{}
	entry.CreatedAt = time.Now().UnixMilli()
	entry.TestNet.Genesis = genesis.GetDefaultTestNetGenesis(p.ctx)
	entry.DevNet.Genesis = genesis.GetDefaultDevNetGenesis(p.ctx)
	entry.MainNet.Genesis = genesis.GetDefaultMainNetGenesis(p.ctx)

	if err := docker.InitCompose(p.ctx, entry.SandboxPath, entry.ID, entry.KeyHash); err != nil {
		estr := fmt.Sprintf("could not initialise containers for project (%s): %s", entry.Name, err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	// projectName := filepath.Base(entry.SandboxPath)
	// ListContainersByProjectName(p.ctx, entry.SandboxPath, projectName, "new-project-stdout")

	names := docker.GenerateContainerNamesFromHash(entry.KeyHash)

	containers, err := docker.GetContainersByNames(p.ctx, names)
	if err != nil {
		estr := fmt.Sprintf("could not get containers: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	runtime.LogDebugf(p.ctx, "containers found: %v", containers)
	entry.Containers = containers

	// TODO: (abdisamad) this all requires error handling
	p.config.Projects = append(p.config.Projects, entry)

	b, err := json.Marshal(p.config)
	if err != nil {
		estr := fmt.Sprintf("could not marshal configuration file to json: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	n, err := filesystem.Save(p.ctx, config.ConfigFileName, b)
	if err != nil {
		estr := fmt.Sprintf("could not save configuration to file: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}
	runtime.LogDebugf(p.ctx, "written configuration file to disk, %d bytes written", n)
	runtime.EventsEmit(p.ctx, "info", "Successfully created project!")

	return true

}

func stdoutToGUI(ctx context.Context, r *io.PipeReader) {
	runtime.LogDebug(ctx, "inside go routine")
	s := bufio.NewScanner(r)

	for s.Scan() {
		runtime.EventsEmit(ctx, "new-project-stdout", string(s.Bytes()))
	}
}

func (p *Project) FetchSandbox(wd string) bool {
	r1, w1 := io.Pipe()
	defer r1.Close()
	defer w1.Close()

	go stdoutToGUI(p.ctx, r1)

	r2, w2 := io.Pipe()
	defer r2.Close()
	defer w2.Close()

	go stdoutToGUI(p.ctx, r2)

	if err := tgit.Clone(wd, SandboxGitURL, w1, w2); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		estr := fmt.Sprintf("could not fetch sandbox repository: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	return true

}

// FetchAlgorandSdk clones Algorand Go SDK into the working directory (wd)
func (p *Project) FetchSdk(wd string) bool {
	r1, w1 := io.Pipe()
	defer r1.Close()
	defer w1.Close()

	go stdoutToGUI(p.ctx, r1)

	r2, w2 := io.Pipe()
	defer r2.Close()
	defer w2.Close()

	go stdoutToGUI(p.ctx, r2)

	if err := tgit.Clone(wd, SDKGitURL, w1, w2); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		estr := fmt.Sprintf("could not fetch sdk repository: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	return true
}

func (p *Project) register(ctx context.Context) {
	p.ctx = ctx
}

// LoadConfig loads the configuration from the config file on disk
func (p *Project) LoadConfig() bool {

	if err := os.Chdir(binaryPath); err != nil {
		estr := fmt.Sprintf("could not change directory: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	_, err := os.Stat(config.ConfigFileName)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		estr := fmt.Sprintf("could not check if file exists: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)

		return false
	}

	var (
		b []byte
	)

	cfg := &config.ProjectsConfig{
		Projects: []model.Project{},
	}
	if err != nil && errors.Is(err, os.ErrNotExist) {

		b, err = json.Marshal(cfg)
		if err != nil {
			estr := fmt.Sprintf("could not setup default configuration file: %s", err)
			runtime.LogError(p.ctx, estr)
			runtime.EventsEmit(p.ctx, "error", estr)
			return false
		}

		f, err := os.Create(config.ConfigFileName)
		if err != nil {
			estr := fmt.Sprintf("could not create configuration file: %s", err)
			runtime.LogError(p.ctx, estr)
			runtime.EventsEmit(p.ctx, "error", estr)
			return false
		}
		defer f.Close()

		_, err = f.Write(b)
		if err != nil {
			estr := fmt.Sprintf("could not write default configuration to file: %s", err)
			runtime.LogError(p.ctx, estr)
			runtime.EventsEmit(p.ctx, "error", estr)
			return false
		}

	} else {
		file, err := os.Open(config.ConfigFileName)
		if err != nil {
			estr := fmt.Sprintf("could not open configuration file: %s", err)
			runtime.LogError(p.ctx, estr)
			runtime.EventsEmit(p.ctx, "error", estr)
			return false
		}
		defer file.Close()

		b, err = io.ReadAll(file)
		if err != nil {
			estr := fmt.Sprintf("could not read configuration from file: %s", err)
			runtime.LogError(p.ctx, estr)
			runtime.EventsEmit(p.ctx, "error", estr)
			return false
		}

	}

	runtime.LogDebugf(p.ctx, "the bytes from the file: %s", string(b))
	if err := json.Unmarshal(b, cfg); err != nil {
		estr := fmt.Sprintf("could not parse configuration json from file: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return false
	}

	if len(cfg.Projects) > 0 {
		for i, proj := range cfg.Projects {
			names := docker.GenerateContainerNamesFromHash(proj.KeyHash)
			containers, err := docker.GetContainersByNames(p.ctx, names)
			if err != nil {
				estr := fmt.Sprintf("could not get containers by their names: %s", err)
				runtime.LogError(p.ctx, estr)
				runtime.EventsEmit(p.ctx, "error", estr)
				return false
			}
			cfg.Projects[i].Containers = containers
		}
	}

	p.config = cfg

	runtime.LogDebugf(p.ctx, "successfully loaded projects config from file")
	return true
}

func (p *Project) GetDefaultTestNetGenesis() model.Genesis {
	return genesis.GetDefaultTestNetGenesis(p.ctx)
}

func (p *Project) GetDefaultMainNetGenesis() model.Genesis {
	return genesis.GetDefaultMainNetGenesis(p.ctx)
}

func (p *Project) GetDefaultDevNetGenesis() model.Genesis {
	return genesis.GetDefaultDevNetGenesis(p.ctx)
}

func (p *Project) SaveProjectConfig(project model.Project) model.Project {
	var pc config.ProjectsConfig
	pc = *p.config

	for i, c := range pc.Projects {
		if c.ID == project.ID {
			pc.Projects[i] = project
			runtime.LogDebugf(p.ctx, "Found match of c(%s) and p(%s) \n", c.ID, p.ID)
			break
		}
	}

	p.config = &pc

	b, err := json.Marshal(p.config)
	if err != nil {
		estr := fmt.Sprintf("could not marshal project configuration to json: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return project
	}

	n, err := filesystem.Save(p.ctx, config.ConfigFileName, b)
	if err != nil {
		estr := fmt.Sprintf("could not save project config to disk: %s", err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return project
	}
	runtime.LogDebugf(p.ctx, "written configuration file to disk successfully, written %d bytes", n)

	return project
}

func (p *Project) SaveTestNetGenesisConfig(project model.Project, gen model.Genesis) model.Genesis {

	var pc config.ProjectsConfig
	pc = *p.config

	for i, c := range pc.Projects {
		if c.ID == project.ID {
			pc.Projects[i] = project
			pc.Projects[i].TestNet.Genesis = gen
			runtime.LogDebugf(p.ctx, "Found match of c(%s) and p(%s) \n", c.ID, p.ID)
			break
		}
	}

	p.config = &pc
	b, err := json.Marshal(p.config)
	if err != nil {
		estr := fmt.Sprintf("could not marshal testnet configuration to json: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return gen
	}

	n, err := filesystem.Save(p.ctx, config.ConfigFileName, b)
	if err != nil {
		estr := fmt.Sprintf("could not write testnet configuration to disk: %s", err)
		runtime.LogError(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return gen
	}

	runtime.LogDebugf(p.ctx, "written successfully testnet genesis configuration to disk, written %d bytes", n)

	runtime.EventsEmit(p.ctx, "info", "Sucessfully saved testnet configuration!")

	return gen
}

func (p *Project) SaveDevNetGenesisConfig(project model.Project, gen model.Genesis) model.Genesis {

	var pc config.ProjectsConfig
	pc = *p.config

	for i, c := range pc.Projects {
		if c.ID == project.ID {
			pc.Projects[i] = project
			pc.Projects[i].DevNet.Genesis = gen
			break
		}

	}

	p.config = &pc
	b, err := json.Marshal(p.config)
	if err != nil {
		estr := fmt.Sprintf("could not marshal devnet genesis config to json: %s", err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return gen
	}

	n, err := filesystem.Save(p.ctx, config.ConfigFileName, b)
	if err != nil {
		estr := fmt.Sprintf("could not write devnet configuration to disk: %s\n", err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, estr)
		return gen
	}

	runtime.LogDebugf(p.ctx, "successfully written devnet genesis configuration to disk, %d bytes written", n)

	runtime.EventsEmit(p.ctx, "info", "Successfully saved devnet configuration!")
	return gen
}

func (p *Project) SaveMainNetGenesisConfig(project model.Project, gen model.Genesis) model.Genesis {

	var pc config.ProjectsConfig
	pc = *p.config

	for i, c := range pc.Projects {
		if c.ID == project.ID {
			pc.Projects[i] = project
			pc.Projects[i].MainNet.Genesis = gen
			break
		}

	}

	p.config = &pc
	b, err := json.Marshal(p.config)
	if err != nil {
		estr := fmt.Sprintf("could not marshal mainnet genesis configuration to json: %s", err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return gen
	}

	n, err := filesystem.Save(p.ctx, config.ConfigFileName, b)
	if err != nil {
		estr := fmt.Sprintf("could not write mainnet genesis configuration to disk: %s", err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return gen
	}
	runtime.LogDebugf(p.ctx, "successfully written mainnnet genesis configuration to disk, %d bytes written to disk", n)
	runtime.EventsEmit(p.ctx, "info", "Successfully saved mainnet configuration!")

	return gen
}

func (p *Project) SaveNewAllocationToConfig(project model.Project, gen model.Genesis) model.Genesis {
	var pc config.ProjectsConfig
	pc = *p.config

	for i, c := range pc.Projects {
		if c.ID == project.ID {
			pc.Projects[i] = project
			pc.Projects[i].MainNet.Genesis = gen
			break
		}

	}

	p.config = &pc
	b, err := json.Marshal(p.config)
	if err != nil {
		estr := fmt.Sprintf("could not marshal new allocation to json: %s", err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return gen
	}

	n, err := filesystem.Save(p.ctx, config.ConfigFileName, b)
	if err != nil {
		estr := fmt.Sprintf("could not new allocation to configuration file on disk: %s", err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)
		return gen
	}
	runtime.LogDebugf(p.ctx, "written successfully new allocation to configuration file on disk, written %d bytes", n)
	runtime.EventsEmit(p.ctx, "info", "Successfully saved new allocation!")

	return gen

}

func (p *Project) CreateProjectFromExampleProject(name string) {

	proj := p.config.Projects[0]

	src := filepath.Join(proj.SDKPath, "examples", name)
	dest := proj.Path
	if err := filesystem.CopyDirectory(src, dest); err != nil {
		estr := fmt.Sprintf("could not copy contents from %s to %s: %s", src, dest, err)
		runtime.LogErrorf(p.ctx, estr)
		runtime.EventsEmit(p.ctx, "error", estr)

		return
	}

	istr := fmt.Sprintf("Successfully created Project from example. Files copied to: %s", dest)
	runtime.EventsEmit(p.ctx, "info", istr)

	return
}

func newProject() *Project {
	return &Project{}
}
