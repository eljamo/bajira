package directory

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/eljamo/bajira/internal/consts"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/strings"
	gap "github.com/muesli/go-app-paths"
)

var scope = gap.NewScope(gap.User, consts.BajiraApplicationName)

var (
	dirPermissionsVal = 0o755
	dirPermissions    = os.FileMode(dirPermissionsVal)
)

// createDir creates all the directories of a given path if they don't exist.
func CreateAllDirectories(path string) error {
	// Check if the directory already exists
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// Directory already exists
		return nil
	}

	// Create the directory
	err := os.MkdirAll(path, dirPermissions)
	if err != nil {
		return errorconc.LocalizedError(err, "failed to create directory")
	}

	// Return nil if the directory was created successfully
	return nil
}

// GetDataDirectory retrieves the first available data directory.
func GetDataDirectory() (string, error) {
	dirs, err := scope.DataDirs()
	if err != nil {
		cerr := errorconc.LocalizedError(err, "failed to get data directory")
		return "", cerr
	}
	if len(dirs) == 0 {
		return "", errorconc.LocalizedError(nil, "no data directory found")
	}

	return dirs[0], nil
}

// GetConfigDirectory retrieves the first available config directory.
func GetConfigDirectory() (string, error) {
	dirs, err := scope.ConfigDirs()
	if err != nil {
		cerr := errorconc.LocalizedError(err, "failed to get config directory")
		return "", cerr
	}
	if len(dirs) == 0 {
		return "", errorconc.LocalizedError(nil, "no config directory found")
	}

	err = CreateAllDirectories(dirs[0])
	if err != nil {
		return "", err
	}

	return dirs[0], nil
}

// GetCacheDirectory retrieves the cache directory.
func GetCacheDirectory() (string, error) {
	dir, err := scope.CacheDir()
	if err != nil {
		cerr := errorconc.LocalizedError(err, "failed to get cache directory")
		return "", cerr
	}

	err = CreateAllDirectories(dir)
	if err != nil {
		return "", err
	}

	return dir, nil
}

// GetApplicationDirectoriesFunc is a type of function that returns the application directories.
type GetApplicationDirectoriesFunc func() (dataDir string, configDir string, cacheDir string, funcErr error)

// GetApplicationDirectories returns the application directories.
func GetApplicationDirectories() (dataDir string, configDir string, cacheDir string, funcErr error) {
	dataDir, err := GetDataDirectory()
	if err != nil {
		return "", "", "", err
	}

	configDir, err = GetConfigDirectory()
	if err != nil {
		return "", "", "", err
	}

	cacheDir, err = GetCacheDirectory()
	if err != nil {
		return "", "", "", err
	}

	return dataDir, configDir, cacheDir, nil
}

// CreateDirectory creates a directory and handles duplicates by appending a suffix.
func CreateSingleDirectory(basePath, dirName, duplicateDirectoryNameFormat string) (string, error) {
	finalDirName := strings.SanitizeString(dirName)
	if strings.StringIsEmpty(finalDirName) {
		return "", errorconc.LocalizedError(nil, "directory name is empty")
	}
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
			cerr := errorconc.LocalizedError(err, "failed to create directory")
			return "", cerr
		}

		// Directory already exists, so increment the counter and try a new name
		counter++
		finalDirName = fmt.Sprintf(duplicateDirectoryNameFormat, dirName, counter)
		fullPath = filepath.Join(basePath, finalDirName)
	}
}

// Give a directory path and this will return a slice of directory names at that path.
func GetSubdirectoryPaths(path string) ([]string, error) {
	var paths []string

	// create path
	err := CreateAllDirectories(path)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(path)
	if err != nil {
		cerr := errorconc.LocalizedError(err, "failed to read directory")
		return nil, cerr
	}

	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, filepath.Join(path, file.Name()))
		}
	}

	return paths, nil
}
