package form

import (
	"context"

	"github.com/charmbracelet/huh"
	"github.com/eljamo/bajira/internal/config"
)

// New creates a new form with the given group. The context is used to get the configuration.
// Which is used to determine if the form should be in accessible mode or not.
func New(ctx context.Context, group *huh.Group) (*huh.Form, error) {
	cfg, err := config.GetConfigFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return huh.NewForm(group).WithAccessible(cfg.AccessibleMode), nil
}
