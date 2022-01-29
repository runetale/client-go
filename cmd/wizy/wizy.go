package main

import (
	"fmt"
	"os"

	"github.com/Notch-Technologies/wizy/cmd/wizy/cmd"
)

func main() {
	if err := cmd.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
