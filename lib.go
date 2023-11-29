package iris

import (
	"fmt"
	"image/color"
)

const Reset = "\x1b[0m"

const ResetFg = "\x1b[39m"
const ResetBg = "\x1b[49m"

const (
	BlackFg = "\x1b[30m"
	BlackBg = "\x1b[40m"

	RedFg = "\x1b[31m"
	RedBg = "\x1b[41m"

	GreenFg = "\x1b[32m"
	GreenBg = "\x1b[42m"

	YellowFg = "\x1b[33m"
	YellowBg = "\x1b[43m"

	BlueFg = "\x1b[34m"
	BlueBg = "\x1b[44m"

	MagentaFg = "\x1b[35m"
	MagentaBg = "\x1b[45m"

	CyanFg = "\x1b[36m"
	CyanBg = "\x1b[46m"

	WhiteFg = "\x1b[37m"
	WhiteBg = "\x1b[47m"
)

func AnsiFgUnchecked(f uint8) string {
	return fmt.Sprintf("\x1b[%dm", f+30)
}

func AnsiBgUnchecked(f uint8) string {
	return fmt.Sprintf("\x1b[%dm", f+40)
}

func AnsiFgBgUnchecked(f uint8, b uint8) string {
	return fmt.Sprintf("\x1b[%d;%dm", f+30, b+40)
}

func Ansi256Fg(f uint8) string {
	return fmt.Sprintf("\x1b[38;5;%dm", f)
}

func Ansi256Bg(b uint8) string {
	return fmt.Sprintf("\x1b[48;5;%dm", b)
}

func Ansi256FgBg(f uint8, b uint8) string {
	return fmt.Sprintf("\x1b[38;5;%dm\x1b[48;5;%dm", f, b)
}

func RgbFg(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r, g, b)
}

func ColorFg(color color.Color) string {
	r, g, b, _ := color.RGBA()
	return RgbFg(uint8(r/0xff), uint8(g/0xff), uint8(b/0xff))
}

func RgbBg(r uint8, g uint8, b uint8) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", r, g, b)
}

func ColorBg(color color.Color) string {
	r, g, b, _ := color.RGBA()
	return RgbBg(uint8(r/0xff), uint8(g/0xff), uint8(b/0xff))
}

func RgbFgBg(fr uint8, fg uint8, fb uint8, br uint8, bg uint8, bb uint8) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm", fr, fg, fb, br, bg, bb)
}

func ColorFgBg(fcolor color.Color, bcolor color.Color) string {
	fr, fg, fb, _ := fcolor.RGBA()
	br, bg, bb, _ := bcolor.RGBA()
	return RgbFgBg(uint8(fr/0xff), uint8(fg/0xff), uint8(fb/0xff), uint8(br/0xff), uint8(bg/0xff), uint8(bb/0xff))
}
