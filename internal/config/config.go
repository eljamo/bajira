package config

const (
	BajiraApplicationName        string = "bajira"
	BajiraDirectoryNameWorkspace string = "workspace"
	BajiraDirectoryNameBoard     string = "board"
	BajiraFileNameConfig         string = "config.toml"
)

// ConfigWorkspace holds the configuration for a workspace.
type ConfigWorkspace struct {
	Key      string
	Name     string
	Archived bool
}

// ConfigBoard holds the configuration for a board.
type ConfigBoard struct {
	Key      string
	Name     string
	Archived bool
	Status   []string
}

// ConfigTask holds the configuration for a task.
type ConfigTask struct {
	Title       string
	Description string
	Status      string
	AssignedTo  []string
	WatchedBy   []string
	Flagged     bool
}
