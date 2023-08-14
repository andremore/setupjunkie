package main

import (
	"fmt"
	"strings"
)

func progressBar(percentage float64, width int) string {
	done := int(percentage * float64(width) / 100)
	pending := width - done
	return fmt.Sprintf("[%s%s]", strings.Repeat("=", done), strings.Repeat(" ", pending))
}
