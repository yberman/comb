# comb
Combinatorics in go!

Most of the time, output a channel of int slices.


## Example

```go
package main

import (
	"fmt"

	"github.com/yberman/comb"
)

func main() {
	for c := range comb.Combinations(5, 3) {
		fmt.Println(c)
	}
}
```

Will output

```
[0 1 2]
[0 1 3]
[0 1 4]
[0 2 3]
[0 2 4]
[0 3 4]
[1 2 3]
[1 2 4]
[1 3 4]
[2 3 4]
```
