package workspace

import (
	"context"

	"github.com/eljamo/bajira/internal/strings"
)

func changeWorkspaceArchiveStatus(ctx context.Context, workspaceId string, status bool) (string, error) {
	cfg, err := getWorkspaceConfig(ctx, workspaceId)
	if err != nil {
		return "", err
	}

	cfg.Archived = status

	err = updateWorkspaceConfig(ctx, workspaceId, cfg)
	if err != nil {
		return "", err
	}

	if status {
		return strings.WorkspaceArchived, nil
	}

	return strings.WorkspaceUnarchived, nil
}

func ArchiveWorkspace(ctx context.Context, workspaceId string) (string, error) {
	return changeWorkspaceArchiveStatus(ctx, workspaceId, true)
}

func UnarchiveWorkspace(ctx context.Context, workspaceId string) (string, error) {
	return changeWorkspaceArchiveStatus(ctx, workspaceId, false)
}
