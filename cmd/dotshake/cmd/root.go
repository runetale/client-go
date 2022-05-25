package cmd

import (
	"context"
	"flag"
	"strings"

	"github.com/peterbourgon/ff/v2/ffcli"
)

func Run(args []string) error {
	if len(args) == 1 && (args[0] == "-V" || args[0] == "--version" || args[0] == "-v") {
		args = []string{"version"}
	}

	fs := flag.NewFlagSet("dotshake", flag.ExitOnError)

	cmd := &ffcli.Command{
		Name:       "dotshake",
		ShortUsage: "dotshake <subcommands> [command flags]",
		ShortHelp:  "Use WireGuard for easy and secure private connections.",
		LongHelp: strings.TrimSpace(`
All flags can use a single or double hyphen.

For help on subcommands, prefix with -help.

Flags and options are subject to change.
`),
		Subcommands: []*ffcli.Command{
			// upCmd,
			loginCmd,
			daemonCmd,
			versionCmd,
		},
		FlagSet: fs,
		Exec:    func(context.Context, []string) error { return flag.ErrHelp },
	}

	if err := cmd.Parse(args); err != nil {
		return err
	}

	if err := cmd.Run(context.Background()); err != nil {
		if err == flag.ErrHelp {
			return nil
		}
		return err
	}

	return nil
}
