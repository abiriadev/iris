package main

import (
	// "fmt"
	"fmt"
	"math"
)

func main() {
	for i := 0; i <= math.MaxUint8; i++ {
		i := uint8(i)
		fmt.Printf("%sForeground %d%s\t%sBackground %d%s\n", Ansi256Fg(i), i, Reset, Ansi256Bg(i), i, Reset)
	}
}
