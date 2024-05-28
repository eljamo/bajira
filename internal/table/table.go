package table

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/eljamo/bajira/internal/styles"
)

var widthPadding = 8

func Generate(headers []string, rows [][]string) *table.Table {
	// Calculate column widths
	colWidth := make([]int, len(headers))
	for i, header := range headers {
		colWidth[i] = len(header) + widthPadding
	}
	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > colWidth[i] {
				colWidth[i] = len(cell) + widthPadding
			}
		}
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(styles.BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == 0:
				return styles.HeaderStyle
			case row%2 == 0:
				style = styles.EvenRowStyle
			default:
				style = styles.OddRowStyle
			}

			for i, width := range colWidth {
				if col == i {
					style = style.Width(width)
				}
			}

			return style
		}).
		Headers(headers...).
		Rows(rows...)

	return t
}
