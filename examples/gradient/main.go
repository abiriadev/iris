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

	monochrome, _ :=
		colorgrad.NewGradient().Build()

	w1, _ := colorgrad.NewGradient().
		HtmlColors("#C41189", "#00BFFF", "#FFD700").
		Build()
	w2, _ := colorgrad.NewGradient().
		HtmlColors("gold", "hotpink", "darkturquoise").
		Build()
	w3, _ := colorgrad.NewGradient().
		HtmlColors("#7d6edd", "#e673f7", "#bbc6f7").
		Build()

	w4, _ := colorgrad.NewGradient().
		HtmlColors("deeppink", "gold", "seagreen").
		Build()

	for _, g := range []colorgrad.Gradient{
		monochrome,
		colorgrad.Sinebow(),
		colorgrad.Rainbow(),
		w3,
		colorgrad.Viridis(),
		colorgrad.Plasma(),
		colorgrad.Spectral(),
		colorgrad.Warm(),
		colorgrad.Cool(),
		colorgrad.Magma(),
		w1, w2, w4,
	} {
		showGrad(g, col)
	}
}
