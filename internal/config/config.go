package config

import (
	"context"
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/eljamo/bajira/internal/consts"
	"github.com/eljamo/bajira/internal/directory"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/toml"
	"github.com/jeandeaual/go-locale"
	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/language"
)

type ConfigContextKey string

type AssigneeConfig struct {
	Default   DefaultAssigneeConfig
	Workspace []WorkspaceAssigneeConfig
}

type DefaultAssigneeConfig struct {
	Name string `toml:"name"`
}

type WorkspaceAssigneeConfig struct {
	Name        string `toml:"name"`
	WorkspaceID string `toml:"workspace_id"`
}

type ApplicationConfig struct {
	AccessibleMode     bool         `toml:"accessible_mode"`
	CacheDirectory     string       `toml:"cache_directory"`
	ConfigDirectory    string       `toml:"config_directory"`
	DataDirectory      string       `toml:"data_directory"`
	DefaultWorkspaceId string       `toml:"default_workspace_id"`
	Locale             language.Tag `toml:"locale"`
	Assignee           AssigneeConfig
}

type ApplicationConfigFile struct {
	AccessibleMode     bool   `toml:"accessible_mode"`
	DataDirectory      string `toml:"data_directory"`
	DefaultWorkspaceId string `toml:"default_workspace_id"`
	Locale             string `toml:"locale"`
	Assignee           AssigneeConfig
}

func guessLocale() error {
	userLocale, err := locale.GetLocale()
	if err != nil {
		return fmt.Errorf("failed to get user locale: %w", err)
	}

	gotext.Configure(
		consts.BajiraPortableObjectDirectoryName,
		userLocale,
		consts.BajiraPortableObjectFileName,
	)

	return nil
}

func validateDirectories(dataDir, configDir, cacheDir string) error {
	if dataDir == configDir {
		return errorconc.LocalizedError(nil, "data directory and config directory are the same")
	}
	if dataDir == cacheDir {
		return errorconc.LocalizedError(nil, "data directory and cache directory are the same")
	}
	return nil
}

func overwriteConfig(cfg *ApplicationConfig, cfgFile *ApplicationConfigFile) {
	if cfgFile.DataDirectory != "" {
		cfg.DataDirectory = cfgFile.DataDirectory
	}
	if cfgFile.DefaultWorkspaceId != "" {
		cfg.DefaultWorkspaceId = cfgFile.DefaultWorkspaceId
	}
	if cfgFile.Locale != "" {
		locale, err := language.Parse(cfgFile.Locale)
		if err != nil {
			return
		}
		cfg.Locale = locale
	}
	if cfgFile.Assignee.Default.Name != "" {
		cfg.Assignee.Default.Name = cfgFile.Assignee.Default.Name
	}
	if len(cfgFile.Assignee.Workspace) > 0 {
		cfg.Assignee.Workspace = cfgFile.Assignee.Workspace
	}
	if cfgFile.AccessibleMode {
		cfg.AccessibleMode = cfgFile.AccessibleMode
	}
}

func GetApplicationConfig() (*ApplicationConfig, error) {
	// guess locale
	err := guessLocale()
	if err != nil {
		return nil, err
	}

	dataDir, configDir, cacheDir, err := directory.GetApplicationDirectories()
	if err != nil {
		return nil, err
	}

	locale, err := language.Parse(consts.BajiraDefaultLanguage)
	if err != nil {
		return nil, errorconc.LocalizedError(err, "failed to parse default language")
	}

	currentUser, err := user.Current()
	if err != nil {
		return nil, errorconc.LocalizedError(err, "failed to get current user")
	}

	cfg := &ApplicationConfig{
		CacheDirectory:  cacheDir,
		ConfigDirectory: configDir,
		DataDirectory:   dataDir,
		Locale:          locale,
		Assignee: AssigneeConfig{
			Default: DefaultAssigneeConfig{
				Name: currentUser.Username,
			},
		},
	}

	var cfgFile ApplicationConfigFile
	err = toml.DecodeFromFile(filepath.Join(configDir, consts.BajiraFileNameConfig), &cfgFile)
	if err != nil {
		return nil, err
	}

	if err := validateDirectories(cfg.DataDirectory, cfg.ConfigDirectory, cfg.CacheDirectory); err != nil {
		return nil, err
	}

	overwriteConfig(cfg, &cfgFile)

	if err := directory.CreateAllDirectories(cfg.DataDirectory); err != nil {
		return nil, err
	}

	gotext.Configure(
		consts.BajiraPortableObjectDirectoryName,
		cfg.Locale.String(),
		consts.BajiraPortableObjectFileName,
	)

	return cfg, nil
}

func GetConfigFromContext(ctx context.Context) (*ApplicationConfig, error) {
	cfg, ok := ctx.Value(ConfigContextKey(consts.Config)).(*ApplicationConfig)
	if !ok {
		return nil, errorconc.LocalizedError(nil, "failed to get application config from context")
	}
	return cfg, nil
}
