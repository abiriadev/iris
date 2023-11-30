package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/abiriadev/iris"
	"github.com/mazznoer/colorgrad"
)

func main() {
	flag.Parse()
	f := flag.Arg(0)

	var in io.Reader

	if f == "" {
		in = os.Stdin
	} else {
		var err error
		in, err = os.Open(f)

		fmt.Println(err)

		if err != nil {
			panic(err)
		}
	}

	r := bufio.NewReader(in)
	w := bufio.NewWriter(os.Stdout)
	p := colorgrad.Sinebow()

	for i, line := 0.0, 0.0; ; i += 0.01 {
		ch, _, err := r.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			return
		}

		if ch == '\n' {
			line++
			i = line / 100
			w.WriteString(iris.Reset)
			w.WriteRune(ch)
			w.Flush()
		} else {
			w.WriteString(iris.ColorFg(p.At(i)))
			w.WriteRune(ch)
		}
	}
}
