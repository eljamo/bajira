package locale

import (
	"fmt"

	"github.com/eljamo/bajira/internal/config"
	"github.com/eljamo/bajira/internal/file"
	"github.com/jeandeaual/go-locale"
	"github.com/leonelquinteros/gotext"
)

func guessLocale() {
	userLocale, err := locale.GetLocale()
	if err != nil {
		// should probably never reach this point
		fmt.Println("Warning: Error getting user locale: ", err)
	}

	gotext.Configure(
		config.BajiraPortableObjectDirectoryName,
		userLocale,
		config.BajiraPortableObjectFileName,
	)
}

func Set() {
	cfg, err := file.GetBajiraConfig()
	if err != nil {
		// should probably never reach this point
		fmt.Println("Warning: Error getting config: ", err)
	}

	if cfg != nil && cfg.Locale != "" {
		gotext.Configure(
			config.BajiraPortableObjectDirectoryName,
			cfg.Locale,
			config.BajiraPortableObjectFileName,
		)
	} else {
		guessLocale()
	}
}
