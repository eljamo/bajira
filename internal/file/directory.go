package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/eljamo/bajira/internal/config"
	gap "github.com/muesli/go-app-paths"
)

var (
	ErrNoDataDirectory   = errors.New("no data directory found")
	ErrNoConfigDirectory = errors.New("no config directory found")
)

var scope = gap.NewScope(gap.User, config.BajiraApplicationName)

var (
	dirPermissionsVal = 0o755
	dirPermissions    = os.FileMode(dirPermissionsVal)
)

// createDir creates all the directories of a given path if they don't exist.
func createAllDirectories(path string) error {
	// Check if the directory already exists
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// Directory already exists
		return nil
	}

	// Create the directory
	err := os.MkdirAll(path, dirPermissions)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Return nil if the directory was created successfully
	return nil
}

// GetDataDirectory retrieves the first available data directory.
func GetDataDirectory() (string, error) {
	dirs, err := scope.DataDirs()
	if err != nil {
		return "", fmt.Errorf("failed to get data directory: %w", err)
	}
	if len(dirs) == 0 {
		return "", ErrNoDataDirectory
	}

	err = createAllDirectories(dirs[0])
	if err != nil {
		return "", err
	}

	return dirs[0], nil
}

// GetConfigDirectory retrieves the first available config directory.
func GetConfigDirectory() (string, error) {
	dirs, err := scope.ConfigDirs()
	if err != nil {
		return "", fmt.Errorf("failed to get config directory: %w", err)
	}
	if len(dirs) == 0 {
		return "", ErrNoConfigDirectory
	}

	err = createAllDirectories(dirs[0])
	if err != nil {
		return "", err
	}

	return dirs[0], nil
}

// GetCacheDirectory retrieves the cache directory.
func GetCacheDirectory() (string, error) {
	dir, err := scope.CacheDir()
	if err != nil {
		return "", fmt.Errorf("failed to get cache directory: %w", err)
	}

	err = createAllDirectories(dir)
	if err != nil {
		return "", err
	}

	return dir, nil
}

// CreateWorkspaceRootDirectories creates the root directories for workspaces.
func CreateWorkspaceRootDirectory() error {
	dir, err := GetDataDirectory()
	if err != nil {
		return err
	}

	// create workspace directory if it doesn't exist
	err = createAllDirectories(filepath.Join(dir, config.BajiraDirectoryNameWorkspace))
	if err != nil {
		return err
	}

	return nil
}

// SanitizeDirectoryName removes invalid characters from a directory name.
func sanitizeDirectoryName(input string) string {
	// Replace invalid characters with an underscore
	re := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)
	sanitized := re.ReplaceAllString(input, "")

	// Trim leading and trailing underscores or whitespace
	sanitized = strings.TrimFunc(sanitized, func(r rune) bool {
		return unicode.IsSpace(r) || r == '_'
	})

	// Return the sanitized directory name
	return sanitized
}

// CreateDirectory creates a directory and handles duplicates by appending a suffix.
func createSingleDirectory(basePath, dirName, duplicateDirectoryNameFormat string) (string, error) {
	finalDirName := sanitizeDirectoryName(dirName)
	fullPath := filepath.Join(basePath, finalDirName)
	counter := 1

	for {
		err := os.Mkdir(fullPath, dirPermissions)
		if err == nil {
			// Directory created successfully
			return fullPath, nil
		}
		if !os.IsExist(err) {
			// An error other than "directory already exists" occurred
			return "", fmt.Errorf("failed to create directory: %w", err)
		}

		// Directory already exists, so increment the counter and try a new name
		counter++
		finalDirName = fmt.Sprintf(duplicateDirectoryNameFormat, dirName, counter)
		fullPath = filepath.Join(basePath, finalDirName)
	}
}

// CreateWorkspaceDirectory creates a workspace directory and handles duplicates by appending a suffix.
func CreateWorkspaceDirectory(basePath, dirName string) (string, error) {
	basePath = filepath.Join(basePath, config.BajiraDirectoryNameWorkspace)

	err := createAllDirectories(basePath)
	if err != nil {
		return "", err
	}

	return createSingleDirectory(basePath, dirName, "%s (%d)")
}

// Give a directory path and this will return a slice of directory names at that path.
func getSubdirectoryPaths(path string) ([]string, error) {
	var paths []string

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, filepath.Join(path, file.Name()))
		}
	}

	return paths, nil
}

// GetAllWorkspaceDirectories returns a slice of workspace directories.
func GetAllWorkspaceDirectories() ([]string, error) {
	dataDir, err := GetDataDirectory()
	if err != nil {
		return nil, err
	}

	workspaceDirPath := filepath.Join(dataDir, config.BajiraDirectoryNameWorkspace)
	workspaceDirs, err := getSubdirectoryPaths(workspaceDirPath)
	if err != nil {
		return nil, err
	}

	return workspaceDirs, nil
}
