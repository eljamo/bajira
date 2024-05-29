package strings

import (
	"strings"
	"sync"

	"github.com/eljamo/bajira/internal/locale"
	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/cases"
)

// Setup a pool of casers to be used for capitalisation
var caserPool = &sync.Pool{
	New: func() any {
		t := locale.GetLanguageTag()
		c := cases.Title(t)
		return &c
	},
}

// Capitalize capitalizes the first letter of a string.
func Capitalize(str string) string {
	c := caserPool.Get().(*cases.Caser)
	defer caserPool.Put(c)
	return c.String(str)
}

var (
	Archived       = gotext.Get("archived")
	ArchivedUpper  = strings.ToUpper(Archived)
	Key            = gotext.Get("key")
	KeyUpper       = strings.ToUpper(Key)
	Name           = gotext.Get("name")
	NameUpper      = strings.ToUpper(Name)
	No             = gotext.Get("no")
	NoCapitalized  = Capitalize(No)
	NoUpper        = strings.ToUpper(No)
	Path           = gotext.Get("path")
	PathUpper      = strings.ToUpper(Path)
	Yes            = gotext.Get("yes")
	YesCapitalized = Capitalize(Yes)
	YesUpper       = strings.ToUpper(Yes)
)

var (
	ArchiveDescription                = gotext.Get("Archive a workspace, board, or task")
	ArchiveWorkspaceDescription       = gotext.Get("Archive a workspace")
	BajiraApplicationDescription      = gotext.Get("A bug tracker, issue tracker, and project management tool")
	CreateDescription                 = gotext.Get("Create a workspace, board, or task")
	CreateWorkspaceDescription        = gotext.Get("Create a new workspace")
	DeleteDescription                 = gotext.Get("Delete a workspace, board, task, or subtask")
	DeleteWorkspaceDescription        = gotext.Get("Delete a workspace")
	HelpDescription                   = gotext.Get("Help about any command")
	ListDescription                   = gotext.Get("List workspaces, boards, or tasks")
	ListWorkspaceDescription          = gotext.Get("List workspaces")
	ListAllWorkspacesDescription      = gotext.Get("List all workspaces")
	ListArchivedWorkspacesDescription = gotext.Get("List archived workspaces")
	NoWorkspacesFound                 = gotext.Get("No workspaces found")
	UnarchiveWorkspaceDescription     = gotext.Get("Unarchive a workspace")
	UnarchiveDescription              = gotext.Get("Unarchive a workspace, board, or task")
	UpdateDescription                 = gotext.Get("Update a workspace, board, or task")
	UpdateWorkspaceDescription        = gotext.Get("Update a workspace")
	WorkspaceNameDescription          = gotext.Get("Name of the workspace")
	WorkspaceKeyDescription           = gotext.Get("Key for the workspace")
	WorkspaceCreated                  = gotext.Get("Workspace created")
)
