package key

import (
	"regexp"
	"strings"
)

var maxKeyLength = 10

// GenerateKey generates a project key from the given input string. It normalizes
// the input by removing non-alphabetic characters, converting to uppercase, and
// truncating to a maximum of 10 characters.
func GenerateKey(input string) string {
	re := regexp.MustCompile(`[^\p{L}\p{N}]`)

	// Remove non-alphabetic characters
	cleaned := re.ReplaceAllString(input, "")

	// Convert to uppercase
	uppercased := strings.ToUpper(cleaned)

	// Truncate to a maximum of 10 characters
	if len(uppercased) > maxKeyLength {
		uppercased = uppercased[:maxKeyLength]
	}

	return uppercased
}
