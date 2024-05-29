package workspace

import (
	"github.com/charmbracelet/lipgloss/table"
	"github.com/eljamo/bajira/internal/strings"
	bajiraTable "github.com/eljamo/bajira/internal/table"
)

var WorkspaceListHeaders = []string{strings.KeyUpper, strings.NameUpper, strings.PathUpper, strings.ArchivedUpper}

func GenerateWorkspaceList(all bool, archived bool) (*table.Table, error) {
	data, err := getWorkspacesData(all, archived)
	if err != nil {
		return nil, err
	}

	return bajiraTable.Generate(WorkspaceListHeaders, data), nil
}
