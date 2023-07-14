package adblockfiltercounter

import (
	"testing"
)

const (
	source string = `
	||hello.world^
	||hello.world^
	||hello.world^
	! Comment
	||another.test^
	`

	filePath string = "./data/AdVoid.Core.txt"
)

func TestCountRules(t *testing.T) {
	t.Run("CountRules() - empty string", func(t *testing.T) {
		const expected int = 0
		var actual int = countRules("")

		if actual != expected {
			t.Error("Should have returned a 0.")
		}
	})

	t.Run("CountRules() - local string", func(t *testing.T) {
		const expected int = 4
		var actual int = countRules(source)

		if actual != expected {
			t.Error("Should have returned 4.")
		}
	})
}

func TestCountFileRules(t *testing.T) {
	t.Run("CountFileRules() - non-valid file", func(t *testing.T) {
		count, error := countFileRules("")

		if count > -1 || error != nil {
			t.Error("Should have errored with a non-valid file, error: ", error)
		}
	})

	t.Run("CountFileRules() - valid file", func(t *testing.T) {
		const expected int = 2495
		actual, error := countFileRules(filePath)

		if actual != expected || error != nil {
			t.Errorf("Should have returned %d but returned %d", expected, actual)
		}
	})
}
