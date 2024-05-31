package form

import (
	"context"
	"sort"

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

func orderedMapKeys[T comparable](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func NewSelect[T comparable](title string, options map[string]T, val *T) *huh.Select[T] {
	sortedKeys := orderedMapKeys(options)
	opts := make([]huh.Option[T], 0, len(options))
	for _, key := range sortedKeys {
		opt := huh.NewOption(key, options[key])
		opts = append(opts, opt)
	}

	return huh.NewSelect[T]().
		Title(title).
		Options(opts...).
		Value(val)
}
