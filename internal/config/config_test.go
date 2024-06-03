package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/eljamo/bajira/internal/consts"
	"golang.org/x/text/language"
)

func TestGuessLocale(t *testing.T) {
	t.Parallel()

	err := guessLocale()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

const (
	testDataDir   = "/data"
	testConfigDir = "/config"
	testCacheDir  = "/cache"
	testUserName  = "user"
)

func TestValidateDirectories(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		dataDir   string
		configDir string
		cacheDir  string
		wantErr   bool
	}{
		{
			name:      "All different",
			dataDir:   testDataDir,
			configDir: testConfigDir,
			cacheDir:  testCacheDir,
			wantErr:   false,
		},
		{
			name:      "Data and config same",
			dataDir:   "/same",
			configDir: "/same",
			cacheDir:  testCacheDir,
			wantErr:   true,
		},
		{
			name:      "Data and cache same",
			dataDir:   "/same",
			configDir: testConfigDir,
			cacheDir:  "/same",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := validateDirectories(tt.dataDir, tt.configDir, tt.cacheDir)
			if (err != nil) != tt.wantErr {
				t.Fatalf("expected error: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}

func TestOverwriteConfig(t *testing.T) {
	t.Parallel()

	cfg := &ApplicationConfig{}
	cfgFile := &ApplicationConfigFile{
		DataDirectory:      testDataDir,
		DefaultWorkspaceId: "workspace123",
		Locale:             "en-US",
		Assignee: AssigneeConfig{
			Default: DefaultAssigneeConfig{
				Name: testUserName,
			},
			Workspace: []WorkspaceAssigneeConfig{
				{
					Name:        "workspaceUser",
					WorkspaceID: "workspace123",
				},
			},
		},
		AccessibleMode: true,
	}

	overwriteConfig(cfg, cfgFile)

	if cfg.DataDirectory != testDataDir {
		t.Fatalf("expected data directory to be /data, got %v", cfg.DataDirectory)
	}
	if cfg.DefaultWorkspaceId != "workspace123" {
		t.Fatalf("expected default workspace id to be workspace123, got %v", cfg.DefaultWorkspaceId)
	}
	if cfg.Locale.String() != "en-US" {
		t.Fatalf("expected locale to be en-US, got %v", cfg.Locale)
	}
	if cfg.Assignee.Default.Name != testUserName {
		t.Fatalf("expected default assignee name to be user, got %v", cfg.Assignee.Default.Name)
	}
	if cfg.AccessibleMode != true {
		t.Fatalf("expected accessible mode to be true, got %v", cfg.AccessibleMode)
	}
}

func TestGetConfigFromContext(t *testing.T) {
	t.Parallel()

	cfg := &ApplicationConfig{
		CacheDirectory:  testCacheDir,
		ConfigDirectory: testConfigDir,
		DataDirectory:   testDataDir,
		Locale:          language.AmericanEnglish,
		Assignee: AssigneeConfig{
			Default: DefaultAssigneeConfig{
				Name: testUserName,
			},
		},
		AccessibleMode: true,
	}

	ctx := context.WithValue(context.Background(), ConfigContextKey(consts.Config), cfg)

	retrievedCfg, err := GetConfigFromContext(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if retrievedCfg.CacheDirectory != testCacheDir {
		t.Fatalf("expected cache directory to be /cache, got %v", retrievedCfg.CacheDirectory)
	}
	if retrievedCfg.ConfigDirectory != testConfigDir {
		t.Fatalf("expected config directory to be /config, got %v", retrievedCfg.ConfigDirectory)
	}
	if retrievedCfg.DataDirectory != testDataDir {
		t.Fatalf("expected data directory to be /data, got %v", retrievedCfg.DataDirectory)
	}
	if retrievedCfg.Locale != language.AmericanEnglish {
		t.Fatalf("expected locale to be AmericanEnglish, got %v", retrievedCfg.Locale)
	}
	if retrievedCfg.Assignee.Default.Name != testUserName {
		t.Fatalf("expected default assignee name to be user, got %v", retrievedCfg.Assignee.Default.Name)
	}
	if retrievedCfg.AccessibleMode != true {
		t.Fatalf("expected accessible mode to be true, got %v", retrievedCfg.AccessibleMode)
	}
}

func TestGetApplicationConfig(t *testing.T) {
	t.Parallel()

	// Set up test environment
	testDir := t.TempDir()
	testDataDir := filepath.Join(testDir, "data")
	testConfigDir := filepath.Join(testDir, "config")
	testCacheDir := filepath.Join(testDir, "cache")

	err := os.MkdirAll(testDataDir, 0o755)
	if err != nil {
		t.Fatalf("failed to create test data directory: %v", err)
	}
	err = os.MkdirAll(testConfigDir, 0o755)
	if err != nil {
		t.Fatalf("failed to create test config directory: %v", err)
	}
	err = os.MkdirAll(testCacheDir, 0o755)
	if err != nil {
		t.Fatalf("failed to create test cache directory: %v", err)
	}

	// Create a dummy config file
	cfgFilePath := filepath.Join(testConfigDir, consts.BajiraFileNameConfig)
	cfgContent := fmt.Sprintf(`
		accessible_mode = true
		data_directory = "%s"
		default_workspace_id = "workspace123"
		locale = "en-US"
		[assignee.default]
		name = "user"
		[[assignee.workspace]]
		name = "workspaceUser"
		workspace_id = "workspace123"
	`, testDataDir)
	err = os.WriteFile(cfgFilePath, []byte(cfgContent), 0o644)
	if err != nil {
		t.Fatalf("failed to write test config file: %v", err)
	}

	// Mock getAppDirsFunc
	mockGetAppDirsFunc := func() (string, string, string, error) {
		return testDataDir, testConfigDir, testCacheDir, nil
	}

	cfg, err := getApplicationConfig(mockGetAppDirsFunc)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if cfg.DataDirectory != testDataDir {
		t.Fatalf("expected data directory to be /data, got %v", cfg.DataDirectory)
	}
	if cfg.DefaultWorkspaceId != "workspace123" {
		t.Fatalf("expected default workspace id to be workspace123, got %v", cfg.DefaultWorkspaceId)
	}
	if cfg.Locale.String() != "en-US" {
		t.Fatalf("expected locale to be en-US, got %v", cfg.Locale)
	}
	if cfg.Assignee.Default.Name != testUserName {
		t.Fatalf("expected default assignee name to be user, got %v", cfg.Assignee.Default.Name)
	}
	if cfg.AccessibleMode != true {
		t.Fatalf("expected accessible mode to be true, got %v", cfg.AccessibleMode)
	}
}
