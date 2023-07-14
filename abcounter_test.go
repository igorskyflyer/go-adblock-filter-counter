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
	"testing"
)

const (
	source string = `
	[Adblock Plus 2.0]
	||hello.world^
	||hello.world^
	||hello.world^
	! Comment
	||another.test^
	`

	filePathAdVoid string = "./data/AdVoid.Core.txt"
	filePathEmpty  string = "./data/empty.txt"
)

func TestCountRules(t *testing.T) {
	t.Run("CountRules() - empty string", func(t *testing.T) {
		const expected int = 0
		var actual int = CountRules("")

		if actual != expected {
			t.Error("Should have returned a 0.")
		}
	})

	t.Run("CountRules() - local string", func(t *testing.T) {
		const expected int = 4
		var actual int = CountRules(source)

		if actual != expected {
			t.Error("Should have returned 4.")
		}
	})
}

func TestCountFileRules(t *testing.T) {
	t.Run("CountFileRules() - non-valid file", func(t *testing.T) {
		count, error := CountFileRules("")

		if count > -1 || error == nil {
			t.Error("Should have errored with a non-valid file.")
		}
	})

	t.Run("CountFileRules() - valid file (non-empty)", func(t *testing.T) {
		const expected int = 2495
		actual, error := CountFileRules(filePathAdVoid)

		if actual != expected || error != nil {
			t.Errorf("Should have returned %d but returned %d", expected, actual)
		}
	})

	t.Run("CountFileRules() - valid file (empty)", func(t *testing.T) {
		const expected int = 0
		actual, error := CountFileRules(filePathEmpty)

		if actual != expected || error != nil {
			t.Errorf("Should have returned %d but returned %d", expected, actual)
		}
	})
}
