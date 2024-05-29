package config

const (
	BajiraApplicationName             string = "bajira"
	BajiraDirectoryNameWorkspace      string = "workspace"
	BajiraDirectoryNameBoard          string = "board"
	BajiraFileNameConfig              string = "config.toml"
	BajiraDefaultLanguage             string = "en"
	BajiraPortableObjectDirectoryName string = "po"
	BajiraPortableObjectFileName      string = "translations"
)

type BajiraConfig struct {
	Locale string
}
