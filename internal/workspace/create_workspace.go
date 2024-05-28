package workspace

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/file"
	"github.com/eljamo/bajira/internal/toml"
)

// CreateWorkspaceForm is a form for creating a new workspace. Used if no arguments are provided.
var CreateWorkspaceForm = huh.NewForm(workspaceNameAndKeyFormGroup)

// CreateWorkspace creates a new workspace with the given name in the data directory.
// If the workspace already exists, a directory with a name appended with a number will be created.
func CreateWorkspace(name string, customKey string) (string, error) {
	workspaceKey, err := generateWorkspaceKey(name, customKey)
	if err != nil {
		return "", err
	}

	dataDirPath, err := file.GetDataDirectory()
	if err != nil {
		return "", err
	}

	workspaceDirPath, err := file.CreateWorkspaceDirectory(dataDirPath, workspaceKey)
	if err != nil {
		return "", fmt.Errorf("failed to create workspace %s: %w", name, err)
	}

	workspaceConfigFilePath := filepath.Join(workspaceDirPath, config.BajiraFileNameConfig)
	err = toml.EncodeToFile(&WorkspaceConfig{Key: workspaceKey, Name: name}, workspaceConfigFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to create workspace config file: %w", err)
	}

	msg := fmt.Sprintf(`Workspace "%s" created at "%s"`, name, workspaceDirPath)

	return msg, nil
}
