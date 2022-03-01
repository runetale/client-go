package cmd

import (
	"flag"

	"github.com/peterbourgon/ff/ffcli"
)

var upCmd = &ffcli.Command{
	Name:      "up",
	Usage:     "up",
	ShortHelp: "launch a logged-in peer",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("up", flag.ExitOnError)
		return fs
	})(),
	Exec: execUp,
}

func execUp(args []string) error {

	return nil
}
