package adblockfiltercounter

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func countRules(source string) int {
	pattern := `\r?\n`
	regex := regexp.MustCompile(pattern)
	lines := regex.Split(source, -1)
	count := 0

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if trimmedLine != "" && !strings.HasPrefix(trimmedLine, "!") {
			count++
		}
	}

	return count
}

func countFileRules(filename string) int {
	source, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0
	}

	return countRules(string(source))
}
