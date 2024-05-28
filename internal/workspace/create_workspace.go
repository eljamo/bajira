package workspace

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/file"
	"github.com/eljamo/bajira/internal/toml"
)

var (
	CreateWorkspaceName string
	CreateWorkspaceKey  string
)

func checkIfStringIsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

// CreateWorkspaceForm is a form for creating a new workspace. Used if no arguments are provided.
var CreateWorkspaceForm = huh.NewForm(
	huh.NewGroup(
		huh.NewInput().
			Title("Workspace Name").
			Value(&CreateWorkspaceName).
			Validate(func(str string) error {
				if checkIfStringIsEmpty(str) {
					return errors.New("name cannot be empty")
				}
				return nil
			}),
		huh.NewInput().
			Title("Workspace Key").
			Description(`Key for the workspace, if not provided a key will be generated from the name. 
If provided, the key will be formatted to be all uppercase and remove any special characters.
			`).
			Value(&CreateWorkspaceKey),
	),
)

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
