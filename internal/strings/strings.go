package strings

import (
	"strconv"
	"strings"

	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CheckIfStringIsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func Capitalize(lang language.Tag, str string) string {
	return cases.Title(lang).String(str)
}

// FormatBool formats a boolean value to localized "true"/"false" and optionally converts to uppercase.
func FormatBool(b bool, upper bool) string {
	result := gotext.Get(strconv.FormatBool(b))
	if upper {
		return strings.ToUpper(result)
	}
	return result
}

// FormatBoolYN formats a boolean value to localized "yes"/"no" and optionally converts to uppercase.
func FormatBoolYN(b bool, upper bool) string {
	result := gotext.Get("no")
	if b {
		result = gotext.Get("yes")
	}
	if upper {
		return strings.ToUpper(result)
	}
	return result
}

// FormatBoolCustom formats a boolean value either as localized "true"/"false" or "yes"/"no", capitalizing the result.
func FormatBoolCapitalized(b bool, yesNo bool) string {
	if yesNo {
		return Capitalize(language.English, FormatBoolYN(b, false))
	}
	return Capitalize(language.English, FormatBool(b, false))
}

// Words
var (
	Archived      = gotext.Get("archived")
	ArchivedUpper = strings.ToUpper(Archived)
	Id            = gotext.Get("id")
	IdUpper       = strings.ToUpper(Id)
	Name          = gotext.Get("name")
	NameUpper     = strings.ToUpper(Name)
	Path          = gotext.Get("path")
	PathUpper     = strings.ToUpper(Path)
)

// Messages and Descriptions
var (
	ArchiveDescription                = gotext.Get("Archive a workspace, board, or task")
	ArchiveWorkspaceDescription       = gotext.Get("Archive a workspace")
	BajiraApplicationDescription      = gotext.Get("A bug tracker, issue tracker, and project management tool")
	CreateDescription                 = gotext.Get("Create a workspace, board, or task")
	CreateWorkspaceDescription        = gotext.Get("Create a new workspace")
	DeleteDescription                 = gotext.Get("Delete a workspace, board, task, or subtask")
	DeleteWorkspaceDescription        = gotext.Get("Delete a workspace")
	ListDescription                   = gotext.Get("List workspaces, boards, or tasks")
	ListWorkspaceDescription          = gotext.Get("List workspaces")
	ListAllWorkspacesDescription      = gotext.Get("List all workspaces")
	ListArchivedWorkspacesDescription = gotext.Get("List archived workspaces")
	SelectAWorkspace                  = gotext.Get("Select a workspace")
	UnarchiveWorkspaceDescription     = gotext.Get("Unarchive a workspace")
	UnarchiveDescription              = gotext.Get("Unarchive a workspace, board, or task")
	UpdateDescription                 = gotext.Get("Update a workspace, board, or task")
	UpdateWorkspaceDescription        = gotext.Get("Update a workspace")
	WorkspaceNameDescription          = gotext.Get("Name of the workspace")
	WorkspaceIdDescription            = gotext.Get("Id for the workspace")
	WorkspaceCreated                  = gotext.Get("Workspace created")
	WorkspaceDeleted                  = gotext.Get("Workspace deleted")
	WorkspaceArchived                 = gotext.Get("Workspace archived")
	WorkspaceUnarchived               = gotext.Get("Workspace unarchived")
)
