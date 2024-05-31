package workspace

import (
	"context"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/form"
	"github.com/eljamo/bajira/internal/strings"
)

func NewDeleteWorkspaceForm(ctx context.Context) (*huh.Form, error) {
	group, err := generateWorkspaceListFormGroup(ctx)
	if err != nil {
		return nil, err
	}

	return form.New(ctx, group)
}

func DeleteWorkspace(ctx context.Context, workspaceId string) (string, error) {
	path, err := getWorkspacePath(ctx, workspaceId)
	if err != nil {
		return "", err
	}

	err = os.RemoveAll(path)
	if err != nil {
		return "", errorconc.LocalizedError(err, "failed to delete workspace")
	}

	return strings.WorkspaceDeleted, nil
}
