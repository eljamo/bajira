package workspace

import (
	"context"
	"fmt"
	"path/filepath"
	"slices"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/consts"
	"github.com/eljamo/bajira/internal/directory"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/form"
	"github.com/eljamo/bajira/internal/key"
	bajiraStrings "github.com/eljamo/bajira/internal/strings"
	"github.com/eljamo/bajira/internal/toml"
)

// WorkspaceConfig holds the configuration for a workspace.
type WorkspaceConfig struct {
	Id       string
	Name     string
	Archived bool
}

var (
	WorkspaceId   string
	WorkspaceName string
)

func NewWorkspaceIdAndNameFormGroup(id, name string) *huh.Group {
	if !bajiraStrings.StringIsEmpty(id) {
		WorkspaceId = id
	}

	if !bajiraStrings.StringIsEmpty(name) {
		WorkspaceName = name
	}

	return huh.NewGroup(
		huh.NewInput().
			Title(bajiraStrings.IdUpper).
			Description(bajiraStrings.WorkspaceIdDescription).
			Value(&WorkspaceId).
			Validate(func(str string) error {
				if len(str) >= 1 && bajiraStrings.StringIsEmpty(str) {
					return errorconc.LocalizedError(nil, "id cannot be empty")
				}
				return nil
			}),
		huh.NewInput().
			Title(bajiraStrings.NameUpper).
			Value(&WorkspaceName).
			Validate(func(str string) error {
				if bajiraStrings.StringIsEmpty(str) {
					return errorconc.LocalizedError(nil, "name cannot be empty")
				}
				return nil
			}),
	)
}

func generateWorkspaceListFormGroup(ctx context.Context, all bool, archived bool) (*huh.Group, error) {
	workspaceIdsNamesPaths, err := getWorkspaces(ctx, all, archived)
	if err != nil {
		return nil, err
	}

	wkspaceMap := make(map[string]string, len(workspaceIdsNamesPaths))
	for _, wkspace := range workspaceIdsNamesPaths {
		key := fmt.Sprintf("%s - %s", wkspace[0], wkspace[1])
		wkspaceMap[key] = wkspace[0]
	}

	return huh.NewGroup(form.NewSelect(bajiraStrings.SelectAWorkspace, wkspaceMap, &WorkspaceId)), nil
}

// getUsedWorkspaceIds returns a slice of all workspace ids in use.
func getUsedWorkspaceIds(ctx context.Context) ([]string, error) {
	path, err := getWorkspaceDirectoryPath(ctx)
	if err != nil {
		return nil, err
	}

	allDirs, err := directory.GetSubdirectoryPaths(path)
	if err != nil {
		return nil, err
	}

	ids := make([]string, len(allDirs))
	for _, dir := range allDirs {
		wsConfig, err := getWorkspaceData(dir)
		if err != nil {
			return nil, err
		}

		ids = append(ids, wsConfig.Id)
	}

	return ids, nil
}

// generateWorkspaceId generates a workspace id, ensuring it doesn't already exist if a custom id is provided.
func generateWorkspaceId(ctx context.Context, name string, customKey string) (string, error) {
	usedWorkspaceIds, err := getUsedWorkspaceIds(ctx)
	if err != nil {
		return "", errorconc.LocalizedError(err, "failed to get used workspace ids")
	}

	if !bajiraStrings.StringIsEmpty(customKey) {
		customKey = key.GenerateKey(customKey)
		if slices.Contains(usedWorkspaceIds, customKey) {
			return "", errorconc.LocalizedError(nil, "workspace id already exists", customKey)
		}
		return customKey, nil
	}

	return generateUniqueId(name, usedWorkspaceIds), nil
}

// generateUniqueId generates a unique id based on the given name and list of used ids.
func generateUniqueId(name string, usedWorkspaceIds []string) string {
	baseKey := key.GenerateKey(name)
	idStr := baseKey
	counter := 1

	for slices.Contains(usedWorkspaceIds, idStr) {
		counter++
		idStr = fmt.Sprintf("%s%d", baseKey, counter)
	}

	return idStr
}

// gets all workspace data from the workspace config files and it's path.
func getWorkspaces(ctx context.Context, all bool, archived bool) ([][]string, error) {
	path, err := getWorkspaceDirectoryPath(ctx)
	if err != nil {
		return nil, err
	}

	allDirs, err := directory.GetSubdirectoryPaths(path)
	if err != nil {
		return nil, err
	}

	idsNamesPaths := make([][]string, 0)
	for _, dir := range allDirs {
		workspaceData, err := getWorkspaceData(dir)
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

		idsNamesPaths = append(idsNamesPaths, getWorkspaceInfo(workspaceData, dir, all))
	}

	return idsNamesPaths, nil
}

func getWorkspaceInfo(workspaceData *WorkspaceConfig, dir string, all bool) []string {
	if all {
		return []string{workspaceData.Id, workspaceData.Name, dir, bajiraStrings.FormatBoolCapitalized(workspaceData.Archived, true)}
	} else {
		return []string{workspaceData.Id, workspaceData.Name, dir}
	}
}

func getWorkspaceData(dir string) (*WorkspaceConfig, error) {
	configPath := filepath.Join(dir, consts.BajiraFileNameConfig)
	var wsConfig WorkspaceConfig
	err := toml.DecodeFromFile(configPath, &wsConfig)
	if err != nil {
		return nil, errorconc.LocalizedError(err, "failed to decode workspace config file")
	}

	return &wsConfig, nil
}

func getWorkspacePath(ctx context.Context, workspaceId string) (string, error) {
	path, err := getWorkspaceDirectoryPath(ctx)
	if err != nil {
		return "", err
	}
	allDirs, err := directory.GetSubdirectoryPaths(path)
	if err != nil {
		return "", err
	}

	for _, dir := range allDirs {
		wsConfig, err := getWorkspaceData(dir)
		if err != nil {
			return "", err
		}

		if wsConfig.Id == workspaceId {
			return dir, nil
		}
	}

	return "", errorconc.LocalizedError(nil, "workspace not found", workspaceId)
}

func GetWorkspaceConfig(ctx context.Context, workspaceId string) (*WorkspaceConfig, error) {
	path, err := getWorkspacePath(ctx, workspaceId)
	if err != nil {
		return nil, err
	}

	return getWorkspaceData(path)
}

func updateWorkspaceConfig(ctx context.Context, workspaceId string, wsConfig *WorkspaceConfig) error {
	path, err := getWorkspacePath(ctx, workspaceId)
	if err != nil {
		return err
	}

	configPath := filepath.Join(path, consts.BajiraFileNameConfig)
	err = toml.EncodeToFile(wsConfig, configPath)
	if err != nil {
		return errorconc.LocalizedError(err, "failed to encode workspace config file")
	}

	return nil
}

func getWorkspaceDirectoryPath(ctx context.Context) (string, error) {
	cfg, err := config.GetConfigFromContext(ctx)
	if err != nil {
		return "", err
	}

	return filepath.Join(cfg.DataDirectory, consts.BajiraDirectoryNameWorkspace), nil
}

func NewSelectWorkspaceForm(ctx context.Context, all bool, archived bool) (*huh.Form, error) {
	group, err := generateWorkspaceListFormGroup(ctx, all, archived)
	if err != nil {
		return nil, err
	}

	return form.New(ctx, group)
}
