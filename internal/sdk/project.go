package sdk

import (
	"tectone/tectone-desktop/internal/filesystem"
)

// GetExampleProjects gets the projects in the examples folder of the SDK
func GetAllExampleProjectNames(path string) ([]string, error) {
	return filesystem.GetDirNames((path))
}
