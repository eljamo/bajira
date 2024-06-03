package key

import (
	"regexp"
)

var maxKeyLength = 10

// GenerateKey generates a project key from the given input string. It normalizes
// the input by removing non-alphabetic characters, and truncating to a maximum
// of 10 characters.
func GenerateKey(input string) string {
	re := regexp.MustCompile(`[^\p{L}\p{N}]`)
	cleaned := re.ReplaceAllString(input, "")
	if len(cleaned) > maxKeyLength {
		cleaned = cleaned[:maxKeyLength]
	}

	return cleaned
}
