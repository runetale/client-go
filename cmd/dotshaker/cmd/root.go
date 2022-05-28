package cmd

// dotshaker commands is an always running daemon process that provides the necessary
// functionality for the dotshake command
// it is a behind-the-scenes process that assists
// udp hole punching and the relay of packets to be implemented in the future.

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

	fs := flag.NewFlagSet("dotshaker", flag.ExitOnError)
	cmd := &ffcli.Command{
		Name:       "dotshaker",
		ShortUsage: "dotshaker <subcommands> [command flags]",
		ShortHelp:  "daemon that provides various functions needed to use dotshaker with dotshake.",
		LongHelp: strings.TrimSpace(`
All flags can use a single or double hyphen.

For help on subcommands, prefix with -help.

Flags and options are subject to change.
`),
		Subcommands: []*ffcli.Command{
			daemonCmd,
			upCmd,
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
