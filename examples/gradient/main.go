package main

import (
	"fmt"

	"github.com/abiriadev/iris"
	"github.com/mazznoer/colorgrad"
	"golang.org/x/term"
)

func showGrad(grad colorgrad.Gradient, col int) {
	for i := 0; i < col; i++ {
		fmt.Printf("%s ", iris.ColorBg(grad.At(float64(i)/float64(col))))
	}

	fmt.Printf("%s\n\n", iris.Reset)

}

func main() {
	col, _, _ := term.GetSize(0)

	grad, _ := colorgrad.NewGradient().Build()

	showGrad(grad, col)

	grad2 := colorgrad.Sinebow()

	showGrad(grad2, col)

	grad3 := colorgrad.Rainbow()

	showGrad(grad3, col)
}
