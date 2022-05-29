package cmd

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/url"
	"strconv"
	"sync"
	"time"

	grpc_client "github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/daemon"
	dd "github.com/Notch-Technologies/dotshake/daemon/dotshaker"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/dotsignal"
	"github.com/Notch-Technologies/dotshake/paths"
	"github.com/Notch-Technologies/dotshake/polymer/conn"
	"github.com/Notch-Technologies/dotshake/store"
	"github.com/Notch-Technologies/dotshake/types/flagtype"
	"github.com/peterbourgon/ff/v2/ffcli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var startArgs struct {
	clientPath string
	signalHost string
	signalPort int64
	logFile    string
	logLevel   string
	dev        bool
	daemon     bool
}

var upCmd = &ffcli.Command{
	Name:       "up",
	ShortUsage: "up [flags]",
	ShortHelp:  "command to start dotshaker.",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("up", flag.ExitOnError)
		fs.StringVar(&startArgs.clientPath, "path", paths.DefaultClientConfigFile(), "client default config file")
		fs.StringVar(&startArgs.signalHost, "signal-host", "http://127.0.0.1", "signaling server host url")
		fs.Int64Var(&startArgs.signalPort, "signal-port", flagtype.DefaultSignalingServerPort, "signaling server host port")
		fs.StringVar(&startArgs.logFile, "logfile", paths.DefaultDotShakerLogFile(), "set logfile path")
		fs.StringVar(&startArgs.logLevel, "loglevel", dotlog.DebugLevelStr, "set log level")
		fs.BoolVar(&startArgs.dev, "dev", true, "is dev")
		fs.BoolVar(&startArgs.daemon, "daemon", false, "whether to run the daemon process")
		return fs
	})(),
	Exec: execUp,
}

func execUp(ctx context.Context, args []string) error {
	err := dotlog.InitDotLog(startArgs.logLevel, startArgs.logFile, startArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger. because %v", err)
	}

	dotlog := dotlog.NewDotLog("dotshaker")

	// configure file store
	//
	cfs, err := store.NewFileStore(paths.DefaultDotshakeClientStateFile(), dotlog)
	if err != nil {
		dotlog.Logger.Fatalf("failed to create clietnt state. because %v", err)
	}

	// configure client store
	//
	cs := store.NewClientStore(cfs, dotlog)
	err = cs.WritePrivateKey()
	if err != nil {
		dotlog.Logger.Fatalf("failed to write client state private key. because %v", err)
	}

	dotlog.Logger.Debugf("client machine key: %s\n", cs.GetPublicKey())

	signalURL := startArgs.signalHost + ":" + strconv.Itoa(int(startArgs.signalPort))
	signalHostURL, err := url.Parse(signalURL)
	if err != nil {
		dotlog.Logger.Fatalf("failed to parsing signal host => [%s:%d]. because %v", err)
		return err
	}

	clientCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	option := grpc.WithTransportCredentials(insecure.NewCredentials())

	gconn, err := grpc.DialContext(
		clientCtx,
		signalHostURL.Host,
		option,
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second,
			Timeout: 10 * time.Second,
		}),
	)
	if err != nil {
		dotlog.Logger.Fatalf("failed to connect server client. because %v", err)
	}

	connState := conn.NewConnectedState()

	signalClient := grpc_client.NewSignalClient(ctx, gconn, connState)

	ch := make(chan struct{})
	mu := &sync.Mutex{}

	ds := dotsignal.NewDotSignal(signalClient, cs.GetPublicKey(), ch, mu, dotlog)

	if startArgs.daemon {
		dotlog.Logger.Debugf("starting dotshaker daemon...\n")
		d := daemon.NewDaemon(dd.BinPath, dd.ServiceName, dd.DaemonFilePath, dd.SystemConfig, dotlog)
		err = d.Install()
		if err != nil {
			dotlog.Logger.Errorf("failed to install dotshaker. %v", err)
			return err
		}
		dotlog.Logger.Debugf("start dotshaker daemon.\n")
		return nil
	}

	ds.ConnectDotSignal()

	select {
	case <-ch:
	case <-ctx.Done():
	}

	return errors.New("terminated the dotshaker")
}
