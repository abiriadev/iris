package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/abiriadev/iris"
	"github.com/mazznoer/colorgrad"
)

func main() {
	stdinr := bufio.NewReader(os.Stdin)
	p := colorgrad.Sinebow()

	for i := 0.0; ; i += 0.01 {
		ch, _, err := stdinr.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			return
		}

		fmt.Printf("%s%c", iris.ColorFg(p.At(i)), ch)
	}
}
