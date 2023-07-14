// MIT License
//
// Copyright (c) 2023 Igor Dimitrijević (igorskyflyer)
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package abcounter

import (
	"os"
	"regexp"
	"strings"
)

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false, err
	} else {
		return true, err
	}
}

// Counts Adblock filter rules found in the provided string.
//
// It returns the number of filter rules.
func CountRules(source string) int {
	var pattern string = `\r?\n`
	var regex *regexp.Regexp = regexp.MustCompile(pattern)
	var lines []string = regex.Split(source, -1)
	var count int = 0

	for _, line := range lines {
		var trimmedLine string = strings.TrimSpace(line)

		if trimmedLine != "" && !strings.HasPrefix(trimmedLine, "!") && !strings.HasPrefix(trimmedLine, "[") {
			count++
		}
	}

	return count
}

// Counts Adblock filter rules found in the provided file.
//
// It returns the number of filter rules or -1 in case of an error.
func CountFileRules(filename string) (int, error) {
	exists, err := fileExists(filename)

	if !exists {
		return -1, err
	}

	source, err := os.ReadFile(filename)

	if err != nil {
		return -1, err
	}

	return CountRules(string(source)), nil
}
