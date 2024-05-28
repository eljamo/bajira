package styles

import (
	"os"

	"github.com/charmbracelet/lipgloss"
)

const (
	White     = lipgloss.Color("255")
	LightGray = lipgloss.Color("248")
	Green     = lipgloss.Color("70")
)

var Renderer = lipgloss.NewRenderer(os.Stdout)
