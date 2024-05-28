package styles

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

const (
	White     = lipgloss.Color("255")
	LightGray = lipgloss.Color("245")
	Green     = lipgloss.Color("70")
)

var Renderer = lipgloss.NewRenderer(os.Stdout)

var cellStyleWidth = 6

var (
	// HeaderStyle is the lipgloss style used for the table headers.
	HeaderStyle = Renderer.NewStyle().Foreground(Green).Bold(true).Align(lipgloss.Center)
	// CellStyle is the base lipgloss style used for the table rows.
	CellStyle = Renderer.NewStyle().Padding(0, 1).Width(cellStyleWidth)
	// OddRowStyle is the lipgloss style used for odd-numbered table rows.
	OddRowStyle = CellStyle.Foreground(LightGray)
	// EvenRowStyle is the lipgloss style used for even-numbered table rows.
	EvenRowStyle = CellStyle.Foreground(White)
	// BorderStyle is the lipgloss style used for the table border.
	BorderStyle = lipgloss.NewStyle().Foreground(Green)
)
