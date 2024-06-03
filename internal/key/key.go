package key

import "github.com/eljamo/bajira/internal/strings"

var maxKeyLength = 10

// GenerateKey generates a project key from the given str string. It normalizes
// the str by removing non-alphabetic characters, and truncating to a maximum
// of 10 characters.
func GenerateKey(str string) string {
	str = strings.SanitizeString(str)
	if len(str) > maxKeyLength {
		str = str[:maxKeyLength]
	}
	return str
}
