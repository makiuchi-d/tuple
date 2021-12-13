# tuple
N-tuple structures for Go with generics

## Types

### 2-tuple

- Couple
- Pair
- Dyad

### 3-tuple

- Triple
- Triplet
- Triad

### 4-tuple

- Quadruple
- Quartet
- Quad
- Tetrad

### 5-tuple

- Quintuple
- Quint
- Pentuple
- Pentad

### 6-tuple

- Sextuple
- Hextuple
- Hexad

### 7-tuple

- Septuple
- Heptuple
- Heptad

### 8-tuple

- Octuple
- Octet
- Octad

### 9-tuple

- Nonuple
- Nonad

### 10-tuple

- Decuple
- Decad

## Example

```
package main

import (
	"fmt"
	"os"

	"github.com/makiuchi-d/tuple"
)

func main() {
	buf := make([]byte, 100)
	ch := make(chan tuple.Pair[int, error])

	go func() {
		n, err := os.Stdin.Read(buf)
		ch <- tuple.NewPair(n, err)
	}()

	r := <-ch
	fmt.Println(buf[:r.V1], r.V2)
}
```
