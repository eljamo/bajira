package strings

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func StringIsEmpty(str string) bool {
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

func SanitizeString(str string) string {
	re := regexp.MustCompile(`[^\p{L}\p{N}]`)
	return re.ReplaceAllString(str, "")
}

// Words
var (
	Archived        = gotext.Get("archived")
	ArchivedUpper   = strings.ToUpper(Archived)
	Id              = gotext.Get("id")
	IdUpper         = strings.ToUpper(Id)
	Name            = gotext.Get("name")
	NameUpper       = strings.ToUpper(Name)
	Path            = gotext.Get("path")
	PathUpper       = strings.ToUpper(Path)
	Status          = gotext.Get("status")
	StatusUpper     = strings.ToUpper(Status)
	Backlog         = gotext.Get("backlog")
	BacklogUpper    = strings.ToUpper(Backlog)
	ToDo            = gotext.Get("to do")
	ToDoUpper       = strings.ToUpper(ToDo)
	InProgress      = gotext.Get("in progress")
	InProgressUpper = strings.ToUpper(InProgress)
	Done            = gotext.Get("done")
	DoneUpper       = strings.ToUpper(Done)
	Closed          = gotext.Get("closed")
	ClosedUpper     = strings.ToUpper(Closed)
	invalid         = gotext.Get("invalid")
	InvalidUpper    = strings.ToUpper(invalid)
)

// Messages and Descriptions
var (
	ArchiveBoardDescription                    = gotext.Get("Archive a board")
	ArchiveDescription                         = gotext.Get("Archive a workspace, board, or task")
	ArchiveWorkspaceDescription                = gotext.Get("Archive a workspace")
	BajiraApplicationDescription               = gotext.Get("A bug tracker, issue tracker, and project management tool")
	BoardCreated                               = gotext.Get("Board created")
	BoardDeleted                               = gotext.Get("Board deleted")
	BoardIdDescription                         = gotext.Get("Id for the board")
	BoardNameDescription                       = gotext.Get("Name of the board")
	CreateBoardDescription                     = gotext.Get("Create a new board")
	CreateDescription                          = gotext.Get("Create a workspace, board, or task")
	CreateWorkspaceDescription                 = gotext.Get("Create a new workspace")
	DeleteBoardDescription                     = gotext.Get("Delete a board")
	DeleteDescription                          = gotext.Get("Delete a workspace, board, task, or subtask")
	DeleteWorkspaceDescription                 = gotext.Get("Delete a workspace")
	ListAllBoardsDescription                   = gotext.Get("List all boards")
	ListAllWorkspacesAndBoardsDescription      = gotext.Get("List all workspaces and boards")
	ListAllWorkspacesDescription               = gotext.Get("List all workspaces")
	ListArchivedBoardsDescription              = gotext.Get("List archived boards")
	ListArchivedWorkspacesAndBoardsDescription = gotext.Get("List archived workspaces and boards")
	ListArchivedWorkspacesDescription          = gotext.Get("List archived workspaces")
	ListBoardsDescription                      = gotext.Get("List boards")
	ListDescription                            = gotext.Get("List workspaces, boards, or tasks")
	ListWorkspaceDescription                   = gotext.Get("List workspaces")
	NewBoardIdDescription                      = gotext.Get("New id for the board")
	NewBoardNameDescription                    = gotext.Get("New name for the board")
	NewWorkspaceIdDescription                  = gotext.Get("New id for the workspace")
	NewWorkspaceNameDescription                = gotext.Get("New name for the workspace")
	SelectAWorkspace                           = gotext.Get("Select a workspace")
	UnarchiveBoardDescription                  = gotext.Get("Unarchive a board")
	UnarchiveDescription                       = gotext.Get("Unarchive a workspace, board, or task")
	UnarchiveWorkspaceDescription              = gotext.Get("Unarchive a workspace")
	UpdateBoardDescription                     = gotext.Get("Update a board")
	UpdateDescription                          = gotext.Get("Update a workspace, board, or task")
	UpdateWorkspaceDescription                 = gotext.Get("Update a workspace")
	WorkspaceArchived                          = gotext.Get("Workspace archived")
	WorkspaceCreated                           = gotext.Get("Workspace created")
	WorkspaceDeleted                           = gotext.Get("Workspace deleted")
	WorkspaceIdDescription                     = gotext.Get("Id for the workspace")
	WorkspaceNameDescription                   = gotext.Get("Name of the workspace")
	WorkspaceUnarchived                        = gotext.Get("Workspace unarchived")
)
