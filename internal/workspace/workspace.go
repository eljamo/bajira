package workspace

import (
	"errors"
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/file"
	"github.com/eljamo/bajira/internal/key"
	"github.com/eljamo/bajira/internal/toml"
)

// WorkspaceConfig holds the configuration for a workspace.
type WorkspaceConfig struct {
	Key      string
	Name     string
	Archived bool
}

var (
	CreateWorkspaceName string
	CreateWorkspaceKey  string
)

func checkIfStringIsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

var workspaceNameAndKeyFormGroup = huh.NewGroup(
	huh.NewInput().
		Title("Workspace Name").
		Value(&CreateWorkspaceName).
		Validate(func(str string) error {
			if checkIfStringIsEmpty(str) {
				return errors.New("name cannot be empty")
			}
			return nil
		}),
	huh.NewInput().
		Title("Workspace Key").
		Description(`Key for the workspace, if not provided a key will be generated from the name. 
If provided, the key will be formatted to be all uppercase and remove any special characters.
		`).
		Value(&CreateWorkspaceKey).
		Validate(func(str string) error {
			if len(str) >= 1 && checkIfStringIsEmpty(str) {
				return errors.New("key cannot be empty")
			}
			return nil
		}),
)

// getUsedWorkspaceKeys returns a slice of all workspace keys in use.
func getUsedWorkspaceKeys() ([]string, error) {
	err := file.CreateWorkspaceRootDirectory()
	if err != nil {
		return nil, err
	}

	allDirs, err := file.GetAllWorkspaceDirectories()
	if err != nil {
		return nil, err
	}

	keys := make([]string, len(allDirs))
	for _, dir := range allDirs {
		configPath := filepath.Join(dir, config.BajiraFileNameConfig)
		var wsConfig WorkspaceConfig
		err := toml.DecodeFromFile(configPath, &wsConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to decode workspace config file: %w", err)
		}

		keys = append(keys, wsConfig.Key)
	}

	return keys, nil
}

// generateWorkspaceKey generates a workspace key, ensuring it doesn't already exist if a custom key is provided.
func generateWorkspaceKey(name, customKey string) (string, error) {
	usedWorkspaceKeys, err := getUsedWorkspaceKeys()
	if err != nil {
		return "", fmt.Errorf("failed to get used workspace keys: %w", err)
	}

	if customKey != "" && !checkIfStringIsEmpty(customKey) {
		customKey = key.GenerateKey(customKey)
		if slices.Contains(usedWorkspaceKeys, customKey) {
			return "", fmt.Errorf("workspace key %s already exists", customKey)
		}
		return customKey, nil
	}

	return generateUniqueKey(name, usedWorkspaceKeys), nil
}

// generateUniqueKey generates a unique key based on the given name and list of used keys.
func generateUniqueKey(name string, usedWorkspaceKeys []string) string {
	baseKey := key.GenerateKey(name)
	keyStr := baseKey
	counter := 1

	for slices.Contains(usedWorkspaceKeys, keyStr) {
		counter++
		keyStr = fmt.Sprintf("%s%d", baseKey, counter)
	}

	return keyStr
}

func getAllWorkspacesData() ([][]string, error) {
	allDirs, err := file.GetAllWorkspaceDirectories()
	if err != nil {
		return nil, err
	}

	var keysNamesPaths [][]string
	for _, dir := range allDirs {
		configPath := filepath.Join(dir, config.BajiraFileNameConfig)
		var wsConfig WorkspaceConfig
		err := toml.DecodeFromFile(configPath, &wsConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to decode workspace config file: %w", err)
		}

		archivedStr := "No"
		if wsConfig.Archived {
			archivedStr = "Yes"
		}

		keysNamesPaths = append(keysNamesPaths, []string{wsConfig.Key, wsConfig.Name, dir, archivedStr})
	}

	return keysNamesPaths, nil
}
