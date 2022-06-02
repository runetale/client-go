package unixsock

import "github.com/Notch-Technologies/client-go/notch/dotshake/v1/machine"

const protocol = "unix"
const sockaddr = "/tmp/rcn.sock"

type SocketMessageType string

const (
	Peer   SocketMessageType = "peer"
	Signal SocketMessageType = "signal"
)

type RecvSocketMesage struct {
	MessageType SocketMessageType
	SignalSock  *SignalSock
	PeerSock    *PeerSock
}

// commands identify what kind of operation,
// the following types are available
// 0 => offer
// 1 => answer
// 2 => candidate
type SignalSock struct {
	Commands int
}

// psock
type PSockCommandType int

const (
	RemovePeers = 0
	ConnPeers   = 1
)

type PeerSock struct {
	RemovePeers map[string]struct{} // for remove peers, otherwise it is nil.

	ConnPeers []*machine.RemotePeer // for conn peers, otherwise it is nil.
	Ip        string                // for conn peers, otherwise it is nil.
	Cidr      string                // for conn peers, otherwise it is nil.
	Commands  PSockCommandType
}
