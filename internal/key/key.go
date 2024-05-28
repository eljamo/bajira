package key

import (
	"regexp"
	"strings"
)

var maxKeyLength = 10

// GenerateJIRAKey generates a JIRA project key from the given input string.
// It normalizes the input by removing non-alphabetic characters, converting to uppercase,
// and truncating to a maximum of 10 characters.
func GenerateKey(input string) string {
	// Define a regular expression to match non-alphabetic characters
	re := regexp.MustCompile("[^a-zA-Z]")

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
