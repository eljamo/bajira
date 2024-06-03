package workspace

import (
	"context"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/form"
)

// NewCreateWorkspaceForm creates a new form for creating a workspace. The context is used to get the configuration.
// Which is used to determine if the form should be in accessible mode or not.
func NewUpdateWorkspaceForm(ctx context.Context, workspaceId string, workspaceName string) (*huh.Form, error) {
	return form.New(ctx, NewWorkspaceIdAndNameFormGroup(workspaceId, workspaceName))
}

func UpdateWorkspace(ctx context.Context, workspaceId string, newWorkspaceId string, newWorkspaceName string) error {
	cfg, err := GetWorkspaceConfig(ctx, workspaceId)
	if err != nil {
		return err
	}

	if cfg.Id != newWorkspaceId {
		cfg.Id = newWorkspaceId
	}

	if cfg.Name != newWorkspaceName {
		cfg.Name = newWorkspaceName
	}

	return updateWorkspaceConfig(ctx, workspaceId, cfg)
}
