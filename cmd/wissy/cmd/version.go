package cmd

import (
	"fmt"
	"log"

	"github.com/Notch-Technologies/wizy/version"
	"github.com/peterbourgon/ff/ffcli"
)

var versionCmd = &ffcli.Command{
	Name:      "version",
	Usage:     "version",
	ShortHelp: "Show Wizy Version",
	Exec:      execVersion,
}

func execVersion(args []string) error {
	if len(args) > 0 {
		log.Fatalf("too many arugments: %q", args)
	}

	fmt.Println(version.String())

	return nil
}
