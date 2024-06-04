package board

import (
	"context"
	"os"

	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/strings"
)

func DeleteBoard(ctx context.Context, workspaceId string, boardId string) (string, error) {
	path, err := GetBoardPath(ctx, workspaceId, boardId)
	if err != nil {
		return "", err
	}

	err = os.RemoveAll(path)
	if err != nil {
		return "", errorconc.LocalizedError(err, "failed to delete board")
	}

	return strings.BoardDeleted, nil
}
