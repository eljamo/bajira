package locale

import (
	"fmt"

	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/file"
	"github.com/jeandeaual/go-locale"
	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/language"
)

func Get() (string, error) {
	cfg, err := file.GetBajiraConfig()
	if err != nil {
		return "", fmt.Errorf("error getting config: %w", err)
	}

	if cfg != nil && cfg.Locale != "" {
		return cfg.Locale, nil
	}

	userLocale, err := locale.GetLocale()
	if err != nil {
		return "", fmt.Errorf("error guessing user locale: %w", err)
	}

	return userLocale, nil
}

func GetLanguageTag() language.Tag {
	locale, err := Get()
	if err != nil {
		return language.English
	}

	tag, err := language.Parse(locale)
	if err != nil {
		return language.English
	}

	return tag
}

func Set() error {
	locale, err := Get()
	if err != nil {
		return fmt.Errorf("error getting locale: %w", err)
	}

	gotext.Configure(
		config.BajiraPortableObjectDirectoryName,
		locale,
		config.BajiraPortableObjectFileName,
	)

	return nil
}
