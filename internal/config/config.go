package config

import (
	"context"
	"path/filepath"

	"github.com/eljamo/bajira/internal/consts"
	"github.com/eljamo/bajira/internal/directory"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/toml"
	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/language"
)

type ConfigContextKey string

type ApplicationConfig struct {
	DataDirectory      string
	DefaultWorkspaceId string
	Locale             language.Tag
}

type ApplicationConfigFile struct {
	DataDirectory      string
	DefaultWorkspaceId string
	Locale             string
}

func GetApplicationConfig() (*ApplicationConfig, error) {
	// Get default application directories
	dataDir, configDir, _, err := directory.GetApplicationDirectories()
	if err != nil {
		return nil, err
	}

	// Parse the default language
	locale, err := language.Parse(consts.BajiraDefaultLanguage)
	if err != nil {
		return nil, errorconc.LocalizedError(err, "failed to parse default language")
	}

	// first set up the default values
	cfg := &ApplicationConfig{
		Locale:        locale,
		DataDirectory: dataDir,
	}

	// read config file to override default values
	var cfgFile ApplicationConfigFile
	err = toml.DecodeFromFile(filepath.Join(configDir, consts.BajiraFileNameConfig), &cfgFile)
	if err != nil {
		return nil, err
	}

	// override default values
	if cfgFile.DataDirectory != "" {
		cfg.DataDirectory = cfgFile.DataDirectory
	}

	if cfgFile.DefaultWorkspaceId != "" {
		cfg.DefaultWorkspaceId = cfgFile.DefaultWorkspaceId
	}

	if cfgFile.Locale != "" {
		locale, err := language.Parse(cfgFile.Locale)
		if err != nil {
			return nil, errorconc.LocalizedError(err, "failed to parse config language")
		}
		cfg.Locale = locale
	}

	// make sure the data directory exists
	err = directory.CreateAllDirectories(cfg.DataDirectory)
	if err != nil {
		return nil, err
	}

	// configure the locale
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
