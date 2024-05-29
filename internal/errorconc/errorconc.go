package errorconc

import (
	"fmt"
	"strings"

	"github.com/leonelquinteros/gotext"
)

// LocalizedError creates an error with a concatenated localized message and wraps the provided error.
func LocalizedError(wrapErr error, strs ...string) error {
	if len(strs) == 0 {
		return wrapErr
	}

	var builder strings.Builder
	for _, str := range strs {
		msg := gotext.Get(str)
		_, err := builder.WriteString(msg)
		if err != nil {
			return fmt.Errorf("error creating localized error message: %w", err)
		}
		_, err = builder.WriteString(": ")
		if err != nil {
			return fmt.Errorf("error creating localized error message: %w", err)
		}
	}

	finalMsg := strings.TrimSuffix(builder.String(), ": ")

	if wrapErr == nil {
		return fmt.Errorf(finalMsg)
	}

	return fmt.Errorf("%s: %w", finalMsg, wrapErr)
}
