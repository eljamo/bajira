package board

import (
	"context"
	"slices"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/form"
	"github.com/eljamo/bajira/internal/workspace"
)

// NewCreateBoardForm creates a new form for creating a workspace. The context is used to get the configuration.
// Which is used to determine if the form should be in accessible mode or not.
func NewUpdateBoardForm(ctx context.Context, boardId string, boardName string) (*huh.Form, error) {
	return form.New(ctx, NewBoardIdAndNameFormGroup(boardId, boardName))
}

func UpdateBoard(ctx context.Context, workspaceId string, boardId string, newBoardId string, newBoardName string) error {
	workspaceDirPath, err := workspace.GetWorkspacePath(ctx, workspaceId)
	if err != nil {
		return err
	}

	cfg, err := GetBoardConfig(ctx, workspaceId, boardId)
	if err != nil {
		return err
	}

	if cfg.Id != newBoardId {
		usedBoardIds, err := getUsedBoardIds(workspaceDirPath)
		if err != nil {
			return err
		}

		if slices.Contains(usedBoardIds, newBoardId) {
			return errorconc.LocalizedError(nil, "board_id already in use")
		}

		cfg.Id = newBoardId
	}

	if cfg.Name != newBoardName {
		cfg.Name = newBoardName
	}

	return updateBoardConfig(ctx, workspaceId, boardId, cfg)
}
