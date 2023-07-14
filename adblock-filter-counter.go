package adblockfiltercounter

import (
	"os"
	"regexp"
	"strings"
)

// Counts Adblock filter rules found in the provided string.
//
// It returns the number of filter rules.
func countRules(source string) int {
	var pattern string = `\r?\n`
	var regex *regexp.Regexp = regexp.MustCompile(pattern)
	var lines []string = regex.Split(source, -1)
	var count int = 0

	for _, line := range lines {
		var trimmedLine string = strings.TrimSpace(line)

		if trimmedLine != "" && !strings.HasPrefix(trimmedLine, "!") {
			count++
		}
	}

	return count
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func countFileRules(filename string) (int, error) {
	exists, err := fileExists(filename)

	if !exists {
		return -1, err
	}

	source, err := os.ReadFile(filename)

	if err != nil {
		return -1, err
	}

	return countRules(string(source)), nil
}
