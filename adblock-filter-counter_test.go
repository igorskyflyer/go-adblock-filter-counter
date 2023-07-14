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
	t.Run("CountRules()", func(t *testing.T) {
		var count int = countRules(source)

		if count != 3 {
			t.Fail()
		}
	})
}

func TestCountFileRules(t *testing.T) {
	t.Run("CountFileRules()", func(t *testing.T) {
		count, error := countFileRules("")

		if count > -1 || error != nil {
			t.Fail()
		}
	})
}
