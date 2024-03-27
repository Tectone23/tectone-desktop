package config

import (
	"context"
	"encoding/json"
	"tectone/tectone-desktop/internal/filesystem"
	"tectone/tectone-desktop/internal/model"
)

// Config represents a configruation for the SDK
//
// TODO: (abdisamad)
type Config struct {
	RootPath string `json:"root_path"`
	ProjectsConfig
	// ContainersConfig
}

// ProjectsConfig stories the configurations configurations for each project
type ProjectsConfig struct {
	Projects []model.Project `json:"projects"`
	// Containers []model.Container `json:"containers"`
}

// SaveProjectConfig saves the project configuration file to disk
func SaveProjectConfig(ctx context.Context, cfg *ProjectsConfig, project model.Project) (*ProjectsConfig, error) {

	for i, c := range cfg.Projects {
		if c.ID == project.ID {
			cfg.Projects[i] = project
			break
		}
	}

	b, err := json.Marshal(cfg)
	if err != nil {
		return cfg, err
	}

	_, err = filesystem.Save(ctx, ConfigFileName, b)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
