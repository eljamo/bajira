package workspace

import (
	"github.com/charmbracelet/lipgloss/table"
	bajiraTable "github.com/eljamo/bajira/internal/table"
)

var WorkspaceListHeaders = []string{"Key", "Name", "Path", "Archived"}

func GenerateWorkspaceList() (*table.Table, error) {
	data, err := getAllWorkspacesData()
	if err != nil {
		return nil, err
	}

	return bajiraTable.Generate(WorkspaceListHeaders, data), nil
}
