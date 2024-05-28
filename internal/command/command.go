package command

import "github.com/eljamo/bajira/internal/config"

const (
	CommandApplicationName string = config.BajiraApplicationName
	CommandArchive         string = "archive"
	CommandBoard           string = "board"
	CommandCreate          string = "create"
	CommandDelete          string = "delete"
	CommandKanban          string = "kanban"
	CommandList            string = "list"
	CommandRename          string = "rename"
	CommandStart           string = "start"
	CommandStop            string = "stop"
	CommandSubtask         string = "subtask"
	CommandTable           string = "table"
	CommandTasks           string = "tasks"
	CommandTask            string = "task"
	CommandTemplate        string = "template"
	CommandTime            string = "time"
	CommandUnarchive       string = "unarchive"
	CommandUpdate          string = "update"
	CommandWorkspace       string = "workspace"
	CommandWorkspaces      string = "workspaces"
)
