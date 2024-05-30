package workspace

import (
	"context"
	"os"

	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/strings"
)

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
