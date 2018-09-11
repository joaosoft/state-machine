package main

import (
	"color"
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, fmt.Sprintf("%s joao", color.WithColor("hello", color.FormatBold, color.ForegroundRed, color.BackgroundCyan)))
}
