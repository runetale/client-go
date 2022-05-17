package cmd

import (
	"fmt"
	"log"

	"github.com/Notch-Technologies/dotshake/version"
	"github.com/peterbourgon/ff/v2/ffcli"
	"golang.org/x/net/context"
)

var versionCmd = &ffcli.Command{
	Name:       "version",
	ShortUsage: "version",
	ShortHelp:  "Show dotshake Version",
	Exec:       execVersion,
}

func execVersion(ctx context.Context, args []string) error {
	if len(args) > 0 {
		log.Fatalf("too many arugments: %q", args)
	}

	fmt.Println(version.String())

	return nil
}
