package strings

import (
	"strings"

	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Capitalize capitalizes the first letter of a string.
func Capitalize(lang language.Tag, str string) string {
	return cases.Title(lang).String(str)
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

// Descriptions
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
	UnarchiveWorkspaceDescription     = gotext.Get("Unarchive a workspace")
	UnarchiveDescription              = gotext.Get("Unarchive a workspace, board, or task")
	UpdateDescription                 = gotext.Get("Update a workspace, board, or task")
	UpdateWorkspaceDescription        = gotext.Get("Update a workspace")
	WorkspaceNameDescription          = gotext.Get("Name of the workspace")
	WorkspaceIdDescription            = gotext.Get("Id for the workspace")
)

// Messages
var (
	WorkspaceCreated = gotext.Get("Workspace created")
	WorkspaceDeleted = gotext.Get("Workspace deleted")
)
