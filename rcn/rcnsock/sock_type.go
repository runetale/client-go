package rcnsock

import (
	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/machine"
)

// TODO: (shinta) is this safe?
// appropriate permission and feel it would be better to
// have a process that creates a file
//
const sockaddr = "/tmp/rcn.sock"

type RSocketMessageType int

const (
	Peer RSocketMessageType = 0
)

type RcnDialSock struct {
	MessageType RSocketMessageType

	PeerSock *PeerSock
}

// psock
type PSockCommandType int

const (
	SyncRemotePeerConnecting PSockCommandType = 0
	SetupRemotePeersConn     PSockCommandType = 1
)

type PeerSock struct {
	Commands PSockCommandType

	RemotePeers []*machine.RemotePeer
	Ip          string
	Cidr        string
}
