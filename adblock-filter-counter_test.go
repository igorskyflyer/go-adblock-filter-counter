package adblockfiltercounter

import (
	"testing"
)

const (
	source string = `
	||hello.world^
	||hello.world^
	||hello.world^
	`
)

func TestCountRules(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		var count int = countRules(source)

		if count != 3 {
			t.Fail()
		}
	})

	// Additional tests...
}
