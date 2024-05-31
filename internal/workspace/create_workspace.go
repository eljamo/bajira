package workspace

import (
	"context"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/consts"
	"github.com/eljamo/bajira/internal/directory"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/form"
	"github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/toml"
)

// NewCreateWorkspaceForm creates a new form for creating a workspace. The context is used to get the configuration.
// Which is used to determine if the form should be in accessible mode or not.
func NewCreateWorkspaceForm(ctx context.Context) (*huh.Form, error) {
	return form.New(ctx, workspaceNameAndKeyFormGroup)
}

// CreateWorkspace creates a new workspace with the given name in the data directory.
// If the workspace already exists, a directory with a name appended with a number will be created.
func CreateWorkspace(ctx context.Context, name string, customKey string) (string, error) {
	workspaceId, err := generateWorkspaceId(ctx, name, customKey)
	if err != nil {
		return "", err
	}

	dataDirPath, err := directory.GetDataDirectory()
	if err != nil {
		return "", err
	}

	workspaceDirPath, err := directory.CreateWorkspaceDirectory(dataDirPath, workspaceId)
	if err != nil {
		cerr := errorconc.LocalizedError(err, "failed to create workspace", name)
		return "", cerr
	}

	workspaceConfigFilePath := filepath.Join(workspaceDirPath, consts.BajiraFileNameConfig)
	err = toml.EncodeToFile(&WorkspaceConfig{Id: workspaceId, Name: name}, workspaceConfigFilePath)
	if err != nil {
		cerr := errorconc.LocalizedError(err, "failed to create workspace config file", name)
		return "", cerr
	}

	return strings.WorkspaceCreated, nil
}
