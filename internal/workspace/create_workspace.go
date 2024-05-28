package workspace

import (
	"fmt"
	"path/filepath"
	"slices"

	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/file"
	"github.com/eljamo/bajira/internal/key"
	"github.com/eljamo/bajira/internal/toml"
)

// CreateWorkspace creates a new workspace with the given name in the data directory.
// If the workspace already exists, a directory with a name appended with a number will be created.
func CreateWorkspace(name string, customKey string) (string, error) {
	workspaceKey, err := generateWorkspaceKey(name, customKey)
	if err != nil {
		return "", err
	}

	dataDirPath, err := file.GetDataDirectory()
	if err != nil {
		return "", err
	}

	workspaceDirPath, err := file.CreateWorkspaceDirectory(dataDirPath, workspaceKey)
	if err != nil {
		return "", fmt.Errorf("failed to create workspace %s: %w", name, err)
	}

	workspaceConfigFilePath := filepath.Join(workspaceDirPath, config.BajiraFileNameConfig)
	err = toml.EncodeToFile(&WorkspaceConfig{Key: workspaceKey, Name: name}, workspaceConfigFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to create workspace config file: %w", err)
	}

	msg := fmt.Sprintf(`Workspace "%s" created at "%s"`, name, workspaceDirPath)

	return msg, nil
}

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

	if customKey != "" {
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
