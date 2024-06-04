package board

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
	"github.com/eljamo/bajira/internal/workspace"
)

// NewCreateBoardForm creates a new form for creating a board. The context is used to get the configuration.
// Which is used to determine if the form should be in accessible mode or not.
func NewCreateBoardForm(ctx context.Context, boardId string, boardName string) (*huh.Form, error) {
	return form.New(ctx, NewBoardIdAndNameFormGroup(boardId, boardName))
}

// createBoardDirectory creates a board directory with the given name in the base path.
// If the directory already exists, a directory with a name appended with a number will be created.
func createBoardDirectory(basePath, dirName string) (string, error) {
	basePath = filepath.Join(basePath, consts.BajiraDirectoryNameBoards)

	err := directory.CreateAllDirectories(basePath)
	if err != nil {
		return "", err
	}

	return directory.CreateSingleDirectory(basePath, dirName, "%s (%d)")
}

// CreateBoard creates a new board with the given name in the boards directory of a workspace.
// If the board directory name is taken already, a directory with a name appended with a number will be created.
func CreateBoard(ctx context.Context, workspaceId string, boardCustomId string, boardName string) (string, error) {
	workspaceDirPath, err := workspace.GetWorkspacePath(ctx, workspaceId)
	if err != nil {
		return "", err
	}

	boardId, err := generateBoardId(workspaceDirPath, boardName, boardCustomId)
	if err != nil {
		return "", err
	}

	boardDirPath, err := createBoardDirectory(workspaceDirPath, boardId)
	if err != nil {
		cerr := errorconc.LocalizedError(err, "failed to create board", boardName)
		return "", cerr
	}

	boardConfigFilePath := filepath.Join(boardDirPath, consts.BajiraFileNameConfig)
	err = toml.EncodeToFile(&BoardConfig{Id: boardId, Name: boardName}, boardConfigFilePath)
	if err != nil {
		cerr := errorconc.LocalizedError(err, "failed to create board config file", boardName)
		return "", cerr
	}

	return strings.BoardCreated, nil
}
