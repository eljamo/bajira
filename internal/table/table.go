package table

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/eljamo/bajira/internal/styles"
)

var (
	characterPadding = 2
	cellStyleWidth   = 6
)

var (
	HeaderStyle  = styles.Renderer.NewStyle().Foreground(styles.Green).Bold(true).Align(lipgloss.Center)
	CellStyle    = styles.Renderer.NewStyle().Padding(0, 1).Width(cellStyleWidth)
	OddRowStyle  = CellStyle.Foreground(styles.White)
	EvenRowStyle = CellStyle.Foreground(styles.LightGray)
	BorderStyle  = lipgloss.NewStyle().Foreground(styles.Green)
)

func getColWidths(headers []string, rows [][]string) []int {
	colWidth := make([]int, len(headers))
	for i, header := range headers {
		colWidth[i] = len(header) + characterPadding
	}
	for _, row := range rows {
		for i, cell := range row {
			val := len(cell) + characterPadding
			if val > colWidth[i] {
				colWidth[i] = len(cell) + characterPadding
			}
		}
	}
	return colWidth
}

func Generate(headers []string, rows [][]string) *table.Table {
	if len(headers) == 0 || len(rows) == 0 {
		return nil
	}

	colWidth := getColWidths(headers, rows)

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
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
