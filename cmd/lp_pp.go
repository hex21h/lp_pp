package main

import (
	"fmt"
	"os"

	"github.com/hex21h/lp_pp/internal/color_cli"
)

func main() {
	colorGreen, err := color_cli.GetColor("green")
	if err != nil {
		os.Exit(0)
	}
	fmt.Println("LP(learning project) price parsing start..............................[", colorGreen, "ok", color_cli.ResetColor(), "]")
}
