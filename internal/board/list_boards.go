package board

import (
	"context"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/eljamo/bajira/internal/errorconc"
	"github.com/eljamo/bajira/internal/strings"
	bajiraTable "github.com/eljamo/bajira/internal/table"
)

func ListBoards(ctx context.Context, workspaceId string, all bool, archived bool) (*table.Table, error) {
	boardListHeaders := []string{strings.IdUpper, strings.NameUpper, strings.PathUpper}

	if all {
		boardListHeaders = append(boardListHeaders, strings.ArchivedUpper)
	}

	data, err := getBoardsTableData(ctx, workspaceId, all, archived)
	if err != nil {
		return nil, err
	}

	table, err := bajiraTable.Generate(boardListHeaders, data)
	if err != nil {
		return nil, errorconc.LocalizedError(err, "failed to generate table")
	}

	return table, nil
}
