package workspace

import (
	"context"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/strings"
	bajiraTable "github.com/eljamo/bajira/internal/table"
)

func ListWorkspaces(ctx context.Context, all bool, archived bool) (*table.Table, error) {
	workspaceListHeaders := []string{strings.IdUpper, strings.NameUpper, strings.PathUpper}

	if all {
		workspaceListHeaders = append(workspaceListHeaders, strings.ArchivedUpper)
	}

	data, err := getWorkspaces(ctx, all, archived)
	if err != nil {
		return nil, err
	}

	table, err := bajiraTable.Generate(workspaceListHeaders, data)
	if err != nil {
		return nil, errorconc.LocalizedError(err, "failed to generate table")
	}

	return table, nil
}
