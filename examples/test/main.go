package main

import (
	"fmt"
	"math"

	"github.com/abiriadev/iris"
)

func main() {
	for i := 0; i <= math.MaxUint8; i++ {
		i := uint8(i)
		fmt.Printf("%sForeground %d%s\t%sBackground %d%s\n", iris.Ansi256Fg(i), i, iris.Reset, iris.Ansi256Bg(i), i, iris.Reset)
	}
}
