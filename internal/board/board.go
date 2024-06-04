package board

import (
	"context"
	"fmt"
	"path/filepath"
	"slices"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/consts"
	"github.com/eljamo/bajira/internal/directory"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/form"
	"github.com/eljamo/bajira/internal/key"
	bajiraStrings "github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/toml"
	"github.com/eljamo/bajira/internal/workspace"
)

var (
	BoardId   string
	BoardName string
)

func NewBoardIdAndNameFormGroup(id, name string) *huh.Group {
	if !bajiraStrings.StringIsEmpty(id) {
		BoardId = id
	}

	if !bajiraStrings.StringIsEmpty(name) {
		BoardName = name
	}

	return huh.NewGroup(
		huh.NewInput().
			Title(bajiraStrings.IdUpper).
			Description(bajiraStrings.WorkspaceIdDescription).
			Value(&BoardId).
			Validate(func(str string) error {
				if len(str) >= 1 && bajiraStrings.StringIsEmpty(str) {
					return errorconc.LocalizedError(nil, "id cannot be empty")
				}
				return nil
			}),
		huh.NewInput().
			Title(bajiraStrings.NameUpper).
			Value(&BoardName).
			Validate(func(str string) error {
				if bajiraStrings.StringIsEmpty(str) {
					return errorconc.LocalizedError(nil, "name cannot be empty")
				}
				return nil
			}),
	)
}

const (
	DefaultBoardStatusBacklog string = "backlog"
	DefaultBoardStatusToDo    string = "todo"
	DefaultBoardStatusInProg  string = "inprog"
	DefaultBoardStatusDone    string = "done"
	DefaultBoardStatusClosed  string = "closed"
	DefaultBoardStatusInvalid string = "invalid"
)

func GetDefaultBoardStatusString(status string) (string, error) {
	switch status {
	case DefaultBoardStatusBacklog:
		return bajiraStrings.BacklogUpper, nil
	case DefaultBoardStatusToDo:
		return bajiraStrings.ToDoUpper, nil
	case DefaultBoardStatusInProg:
		return bajiraStrings.InProgressUpper, nil
	case DefaultBoardStatusDone:
		return bajiraStrings.DoneUpper, nil
	case DefaultBoardStatusClosed:
		return bajiraStrings.ClosedUpper, nil
	}

	return bajiraStrings.InvalidUpper, errorconc.LocalizedError(nil, "invalid default board status", status)
}

var DefaultStatuses = []string{
	DefaultBoardStatusBacklog,
	DefaultBoardStatusToDo,
	DefaultBoardStatusInProg,
	DefaultBoardStatusDone,
	DefaultBoardStatusClosed,
	DefaultBoardStatusInvalid,
}

var DefaultKanbanStatuses = []string{
	DefaultBoardStatusToDo,
	DefaultBoardStatusInProg,
	DefaultBoardStatusDone,
}

// BoardConfig holds the configuration for a workspace.
type BoardConfig struct {
	Id             string
	Name           string
	Archived       bool
	Statuses       []string
	KanbanStatuses []string
}

type BoardData struct {
	BoardConfig
	Path string
}

func getBoardsPath(ctx context.Context, workspaceId string) (string, error) {
	workspaceDirPath, err := workspace.GetWorkspacePath(ctx, workspaceId)
	if err != nil {
		return "", err
	}

	return filepath.Join(workspaceDirPath, consts.BajiraDirectoryNameBoards), nil
}

func getWorkspaceBoardsPaths(ctx context.Context, workspaceId string) ([]string, error) {
	boardsPath, err := getBoardsPath(ctx, workspaceId)
	if err != nil {
		return nil, err
	}

	return directory.GetSubdirectoryPaths(boardsPath)
}

func getBoardData(dir string) (*BoardConfig, error) {
	configPath := filepath.Join(dir, consts.BajiraFileNameConfig)
	var bcfg BoardConfig
	err := toml.DecodeFromFile(configPath, &bcfg)
	if err != nil {
		return nil, errorconc.LocalizedError(err, "failed to decode workspace config file")
	}

	return &bcfg, nil
}

// getUsedBoardIds returns a slice of all workspace ids in use.
func getUsedBoardIds(workspaceDirPath string) ([]string, error) {
	path := filepath.Join(workspaceDirPath, consts.BajiraDirectoryNameBoards)
	allDirs, err := directory.GetSubdirectoryPaths(path)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(allDirs))
	for _, dir := range allDirs {
		bcfg, err := getBoardData(dir)
		if err != nil {
			return nil, err
		}

		ids = append(ids, bcfg.Id)
	}

	return ids, nil
}

// generateUniqueId generates a unique id based on the given name and list of used ids.
func generateUniqueId(name string, usedBoardIds []string) string {
	baseKey := key.GenerateKey(name)
	idStr := baseKey
	counter := 1

	for slices.Contains(usedBoardIds, idStr) {
		counter++
		idStr = fmt.Sprintf("%s%d", baseKey, counter)
	}

	return idStr
}

// generateBoardId generates a workspace id, ensuring it doesn't already exist if a custom id is provided.
func generateBoardId(workspaceDirPath string, name string, customKey string) (string, error) {
	usedBoardIds, err := getUsedBoardIds(workspaceDirPath)
	if err != nil {
		return "", errorconc.LocalizedError(err, "failed to get used boards ids")
	}

	if !bajiraStrings.StringIsEmpty(customKey) {
		customKey = key.GenerateKey(customKey)
		if slices.Contains(usedBoardIds, customKey) {
			return "", errorconc.LocalizedError(nil, "board id already exists", customKey)
		}
		return customKey, nil
	}

	return generateUniqueId(name, usedBoardIds), nil
}

func GetBoardPath(ctx context.Context, workspaceId string, boardId string) (string, error) {
	allDirs, err := getWorkspaceBoardsPaths(ctx, workspaceId)
	if err != nil {
		return "", err
	}

	for _, dir := range allDirs {
		bcfg, err := getBoardData(dir)
		if err != nil {
			return "", err
		}

		if bcfg.Id == boardId {
			return dir, nil
		}
	}

	return "", errorconc.LocalizedError(nil, "board not found", boardId)
}

// gets all board data from the boards config files and it's path.
func getBoards(ctx context.Context, workspaceId string, all bool, archived bool) ([]BoardData, error) {
	allDirs, err := getWorkspaceBoardsPaths(ctx, workspaceId)
	if err != nil {
		return nil, err
	}

	boards := make([]BoardData, 0)
	for _, dir := range allDirs {
		workspaceData, err := getBoardData(dir)
		if err != nil {
			return nil, err
		}

		// Skip non-archived workspaces if only archived workspaces are requested.
		if archived && !workspaceData.Archived {
			continue
		}

		// Skip archived workspaces if not all workspaces are requested.
		if !all && !archived && workspaceData.Archived {
			continue
		}

		boards = append(boards, BoardData{*workspaceData, dir})
	}

	return boards, nil
}

func getBoardsTableData(ctx context.Context, workspaceId string, all bool, archived bool) ([][]string, error) {
	data := make([][]string, 0)
	boards, err := getBoards(ctx, workspaceId, all, archived)
	if err != nil {
		return nil, err
	}

	for _, board := range boards {
		if all {
			data = append(data, []string{board.Id, board.Name, board.Path, bajiraStrings.FormatBoolCapitalized(board.Archived, true)})
		} else {
			data = append(data, []string{board.Id, board.Name, board.Path})
		}
	}

	return data, nil
}

func generateBoardListFormGroup(ctx context.Context, workspaceId string, all bool, archived bool) (*huh.Group, error) {
	boards, err := getBoards(ctx, workspaceId, all, archived)
	if err != nil {
		return nil, err
	}

	bMap := make(map[string]string, len(boards))
	for _, board := range boards {
		key := fmt.Sprintf("%s - %s", board.Id, board.Name)
		bMap[key] = board.Id
	}

	return huh.NewGroup(form.NewSelect(bajiraStrings.SelectAWorkspace, bMap, &BoardId)), nil
}

func NewSelectBoardForm(ctx context.Context, workspaceId string, all bool, archived bool) (*huh.Form, error) {
	group, err := generateBoardListFormGroup(ctx, workspaceId, all, archived)
	if err != nil {
		return nil, err
	}

	return form.New(ctx, group)
}
