package directory

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestCreateAllDirectories(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		path          string
		expectedError bool
	}{
		{
			name:          "Valid Path",
			path:          filepath.Join(os.TempDir(), "testdir"),
			expectedError: false,
		},
		{
			name:          "Empty Path",
			path:          "",
			expectedError: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := CreateAllDirectories(tt.path)
			if (err != nil) != tt.expectedError {
				t.Errorf("CreateAllDirectories() error = %v, expectedError %v", err, tt.expectedError)
			}
			if !tt.expectedError {
				_, err := os.Stat(tt.path)
				if os.IsNotExist(err) {
					t.Errorf("CreateAllDirectories() failed, directory not created: %s", tt.path)
				}
			}
		})
	}
}

func TestCreateSingleDirectory(t *testing.T) {
	t.Parallel()

	basePath := os.TempDir()
	dirName := "testdir"
	duplicateFormat := "%s_%d"

	tests := []struct {
		name                  string
		initialDir            string
		duplicateDir          string
		expectedFinalDirCount int
		expectedError         bool
	}{
		{
			name:                  "Valid Directory",
			initialDir:            dirName,
			duplicateDir:          dirName,
			expectedFinalDirCount: 1,
			expectedError:         false,
		},
		{
			name:                  "Duplicate Directory",
			initialDir:            dirName,
			duplicateDir:          dirName,
			expectedFinalDirCount: 2,
			expectedError:         false,
		},
		{
			name:                  "Empty Directory Name",
			initialDir:            "",
			duplicateDir:          "",
			expectedFinalDirCount: 0,
			expectedError:         true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			finalDir, err := CreateSingleDirectory(basePath, tt.initialDir, duplicateFormat)
			if (err != nil) != tt.expectedError {
				t.Errorf("CreateSingleDirectory() error = %v, expectedError %v", err, tt.expectedError)
			}
			if !tt.expectedError {
				if _, err := os.Stat(finalDir); os.IsNotExist(err) {
					t.Errorf("CreateSingleDirectory() failed, directory not created: %s", finalDir)
				}
			}
		})
	}
}

//gocognit:ignore
func TestGetSubdirectoryPaths(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		path          string
		expectedError bool
		expectedDirs  []string
	}{
		{
			name:          "Valid Directory",
			path:          filepath.Join(os.TempDir(), "parentdir"),
			expectedError: false,
			expectedDirs:  []string{"childdir1", "childdir2"},
		},
		{
			name:          "Non-existent Directory",
			path:          filepath.Join(os.TempDir(), "nonexistentdir"),
			expectedError: false,
			expectedDirs:  []string{},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if tt.name == "Valid Directory" {
				setupTestDirectory(t, tt.path, tt.expectedDirs)
				defer teardownTestDirectory(t, tt.path)
			}

			paths, err := GetSubdirectoryPaths(tt.path)
			if (err != nil) != tt.expectedError {
				t.Errorf("GetSubdirectoryPaths() error = %v, expectedError %v", err, tt.expectedError)
			}
			if tt.expectedError {
				return
			}

			expectedPaths := []string{}
			for _, dir := range tt.expectedDirs {
				expectedPaths = append(expectedPaths, filepath.Join(tt.path, dir))
			}

			if paths == nil {
				paths = []string{}
			}
			if expectedPaths == nil {
				expectedPaths = []string{}
			}

			if !slices.Equal(paths, expectedPaths) {
				t.Errorf("GetSubdirectoryPaths() got = %v, want %v", paths, expectedPaths)
			}
		})
	}
}

func setupTestDirectory(t *testing.T, parentDir string, dirs []string) {
	t.Helper()

	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(parentDir, dir), dirPermissions)
		if err != nil {
			t.Fatalf("Setup error: %v", err)
		}
	}
}

func teardownTestDirectory(t *testing.T, parentDir string) {
	t.Helper()

	if err := os.RemoveAll(parentDir); err != nil {
		t.Fatalf("Teardown error: %v", err)
	}
}
