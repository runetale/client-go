package main

import (
	"fmt"
	"os"

	"github.com/Notch-Technologies/dotshake/cmd/dotshaker/cmd"
)

func main() {
	if err := cmd.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
