package main

import (
	"fmt"
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/termenv"
)

// 5800FF 0096FF 00D7FF 72FFFF
var (
	term          = termenv.EnvColorProfile()
	keyword       = makeFgStyle("#00D7FF")
	subtle        = makeFgStyle("241")
	dot           = colorFg(" • ", "236")
)

func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}

func makeRamp(colorA, colorB string, steps float64) (s []string) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(s, colorToHex(c))
	}
	return
}

func colorToHex(c colorful.Color) string {
	return fmt.Sprintf("#%s%s%s", colorFloatToHex(c.R), colorFloatToHex(c.G), colorFloatToHex(c.B))
}

func colorFloatToHex(f float64) (s string) {
	s = strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return
}

func moveCursorUp(m model) model {
	if m.Cursor > 0 {
		m.Cursor--
	}
	return m
}

func moveCursorDown(m model) model {
	if m.Cursor < len(choices)-1 {
		m.Cursor++
	}
	return m
}
