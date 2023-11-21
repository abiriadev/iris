package main

import (
	"fmt"
)

func Rgb(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}
