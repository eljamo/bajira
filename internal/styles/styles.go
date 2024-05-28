package styles

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

const (
	Purple    = lipgloss.Color("99")
	Gray      = lipgloss.Color("245")
	LightGray = lipgloss.Color("241")
)

var Renderer = lipgloss.NewRenderer(os.Stdout)

var (
	// HeaderStyle is the lipgloss style used for the table headers.
	HeaderStyle = Renderer.NewStyle().Foreground(Purple).Bold(true).Align(lipgloss.Center)
	// CellStyle is the base lipgloss style used for the table rows.
	CellStyle = Renderer.NewStyle().Padding(0, 1).Width(14)
	// OddRowStyle is the lipgloss style used for odd-numbered table rows.
	OddRowStyle = CellStyle.Foreground(Gray)
	// EvenRowStyle is the lipgloss style used for even-numbered table rows.
	EvenRowStyle = CellStyle.Foreground(LightGray)
	// BorderStyle is the lipgloss style used for the table border.
	BorderStyle = lipgloss.NewStyle().Foreground(Purple)
)
