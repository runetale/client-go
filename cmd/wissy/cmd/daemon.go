package cmd

import (
	"context"
	"flag"
	"log"

	"github.com/Notch-Technologies/wizy/daemon"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/peterbourgon/ff/v2/ffcli"
)

var daemonArgs struct {
	logFile  string
	logLevel string
	dev      bool
}

var daemonCmd = &ffcli.Command{
	Name:       "daemon",
	ShortUsage: "daemon <subcommand> [command flags]",
	ShortHelp:  "Install and uninstall daemons, etc",
	Exec:       func(context.Context, []string) error { return flag.ErrHelp },
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("up", flag.ExitOnError)
		fs.StringVar(&daemonArgs.logFile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&daemonArgs.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
		fs.BoolVar(&daemonArgs.dev, "dev", true, "is dev")
		return fs
	})(),
	Subcommands: []*ffcli.Command{
		installDaemonCmd,
		uninstallDaemonCmd,
		statusCmd,
	},
}

var installDaemonCmd = &ffcli.Command{
	Name:       "install",
	ShortUsage: "install",
	ShortHelp:  "install the daemon",
	Exec:       installDaemon,
}

func installDaemon(ctx context.Context, args []string) error {
	err := wislog.InitWisLog(daemonArgs.logLevel, daemonArgs.logFile, daemonArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	wislog := wislog.NewWisLog("daemon")

	d := daemon.NewDaemon(daemon.TargetPath, daemon.ServiceName, daemon.PlistName, daemon.PlistFile, wislog)
	err = d.Install()
	if err != nil {
		return err
	}
	return nil
}

var uninstallDaemonCmd = &ffcli.Command{
	Name:       "uninstall",
	ShortUsage: "uninstall",
	ShortHelp:  "uninstall the daemon",
	Exec:       uninstallDaemon,
}

func uninstallDaemon(ctx context.Context, args []string) error {
	err := wislog.InitWisLog(daemonArgs.logLevel, daemonArgs.logFile, daemonArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	wislog := wislog.NewWisLog("daemon")

	d := daemon.NewDaemon(daemon.TargetPath, daemon.ServiceName, daemon.PlistName, daemon.PlistFile, wislog)
	err = d.Uninstall()
	if err != nil {
		return err
	}
	return nil
}

var statusCmd = &ffcli.Command{
	Name:       "status",
	ShortUsage: "status",
	ShortHelp:  "status the daemon",
	Exec:       statusDaemon,
}

func statusDaemon(ctx context.Context, args []string) error {
	err := wislog.InitWisLog(daemonArgs.logLevel, daemonArgs.logFile, daemonArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	wislog := wislog.NewWisLog("daemon")

	d := daemon.NewDaemon(daemon.TargetPath, daemon.ServiceName, daemon.PlistName, daemon.PlistFile, wislog)
	err = d.Status()
	if err != nil {
		return err
	}
	return nil
}
