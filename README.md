# Google Translate for Go.

## Install

```
go get github.com/notiku/gotranslate
```

## Example

```go
package main

import (
	"fmt"

	gt "github.com/notiku/gotranslate"
)

var text string = "Hello, World!"

func main() {
	result, _ := gt.Translate(text, "en", "es")
	fmt.Println(result.Translated)
}
```