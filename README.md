# 🐝 go := "Adblock filter counter"

🐲 A dead simple Go module that counts Adblock filter rules.🦘

> Note: it only operates on a string-level, it doesn't check the validity of the provided filter rules.

<br>
<br>

Documentation: https://godocs.io/github.com/igorskyflyer/go-adblock-filter-counter

See the [releases](https://github.com/igorskyflyer/go-adblock-filter-counter/releases) page for a changelog.

<br>

This library requires Go 1.20 or newer; add it to your go.mod with:

```shell
go get github.com/igorskyflyer/go-adblock-filter-counter@latest
```

and import it like

```go
...

import (
	abcounter "github.com/igorskyflyer/go-adblock-filter-counter"
)

...
```

<br>

## 🤹🏼 Examples

### `CountRules(source string) int`

```go
source string = `
	[Adblock Plus 2.0]
	||hello.world^
	||hello.world^
	||hello.world^
	! Comment
	||another.test^
	`

abcounter.CountRules(source) // returns 4

```

<br>

### `CountFileRules(filename string) (int, error)`

```go
filePath string = "./data/AdVoid.Core.txt"

abcounter.CountFileRules(filePath) // returns (2495, nil)

```
