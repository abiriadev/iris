package main

import (
	"fmt"
	"math"

	"github.com/abiriadev/iris"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/term"
)

func fitSize() int {
	col, row, _ := term.GetSize(0)

	m := min(col/2, row)

	if m == 0 {
		return m
	}

	return 1 << math.Ilogb(float64(m))
}

func main() {
	s := fitSize()

	for y := 0; y < s; y++ {
		for x := 0; x < s; x++ {
			fmt.Printf("%s  ", iris.ColorBg(colorful.Hsl(float64(x^y)*360/float64(s), 1, 0.5)))
		}

		fmt.Printf("%s\n", iris.Reset)
	}
}
