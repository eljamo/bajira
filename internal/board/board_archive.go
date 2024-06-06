package board

import (
	"context"

	"github.com/eljamo/bajira/internal/strings"
)

func changeBoardArchiveStatus(ctx context.Context, workspaceId string, boardId string, status bool) (string, error) {
	cfg, err := GetBoardConfig(ctx, workspaceId, boardId)
	if err != nil {
		return "", err
	}

	cfg.Archived = status

	err = updateBoardConfig(ctx, workspaceId, boardId, cfg)
	if err != nil {
		return "", err
	}

	if status {
		return strings.BoardArchived, nil
	}

	return strings.BoardUnarchived, nil
}

func ArchiveBoard(ctx context.Context, workspaceId string, boardId string) (string, error) {
	return changeBoardArchiveStatus(ctx, workspaceId, boardId, true)
}

func UnarchiveBoard(ctx context.Context, workspaceId string, boardId string) (string, error) {
	return changeBoardArchiveStatus(ctx, workspaceId, boardId, false)
}
