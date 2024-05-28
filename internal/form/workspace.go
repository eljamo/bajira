package form

import (
	"errors"
	"strings"

	"github.com/charmbracelet/huh"
)

var (
	CreateWorkspaceName string
	CreateWorkspaceKey  string
)

func checkIfStringIsEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

var CreateWorkspaceForm = huh.NewForm(
	huh.NewGroup(
		huh.NewInput().
			Title("Workspace Name").
			Value(&CreateWorkspaceName).
			Validate(func(str string) error {
				if checkIfStringIsEmpty(str) {
					return errors.New("name cannot be empty")
				}
				return nil
			}),
		huh.NewInput().
			Title("Workspace Key").
			Description(`Key for the workspace, if not provided a key will be generated from the name. 
If provided, the key will be formatted to be all uppercase and remove any special characters.
			`).
			Value(&CreateWorkspaceKey),
	),
)
