package cmd

import (
	"flag"
	"strings"

	"github.com/peterbourgon/ff/ffcli"
)

func Run(args []string) error {
	if len(args) == 1 && (args[0] == "-V" || args[0] == "--version" || args[0] == "-v") {
		args = []string{"version"}
	}

	fs := flag.NewFlagSet("wissy", flag.ExitOnError)

	cmd := &ffcli.Command{
		Name:      "wissy",
		Usage:     "wissy <subcommands> [command flags]",
		ShortHelp: "Use WireGuard for easy and secure private connections.",
		LongHelp: strings.TrimSpace(`
All flags can use a single or double hyphen.

For help on subcommands, prefix with -help.

Flags and options are subject to change.
`),
		Subcommands: []*ffcli.Command{
			upCmd,
			loginCmd,
			// uninstall,
			// installSystemDaemon,
			versionCmd,
		},
		FlagSet: fs,
		Exec:    func([]string) error { return flag.ErrHelp },
	}

	if err := cmd.Run(args); err != nil {
		if err == flag.ErrHelp {
			return nil
		}
		return err
	}

	return nil
}
