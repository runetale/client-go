package rcnsock

// a package that communicates using rcn and unix sockets
//

import (
	"encoding/gob"
	"io"
	"net"
	"os"

	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/rcn/controlplane"
)

type RcnSock struct {
	scp *controlplane.ControlPlane

	dotlog *dotlog.DotLog

	ch chan struct{}
}

// if scp is nil when making this function call, just listen
//
func NewRcnSock(
	dotlog *dotlog.DotLog,
	ch chan struct{},
	scp *controlplane.ControlPlane,
) *RcnSock {
	return &RcnSock{
		scp: scp,

		dotlog: dotlog,

		ch: ch,
	}
}

func (s *RcnSock) cleanup() error {
	if _, err := os.Stat(sockaddr); err == nil {
		if err := os.RemoveAll(sockaddr); err != nil {
			return err
		}
	}
	return nil
}

func (s *RcnSock) listen(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for {
		m := &RcnDialSock{}
		err := decoder.Decode(m)
		if err != nil {
			if err == io.EOF {
				s.dotlog.Logger.Debugf("close the rcn socket client")
				break
			}
			break
		}

		switch m.MessageType {
		case Peer:
			s.dotlog.Logger.Debugf("read from peer message => %d", m.PeerSock.Commands)
			switch m.PeerSock.Commands {
			case SyncRemotePeerConnecting:
				if m.PeerSock.RemotePeers == nil {
					s.dotlog.Logger.Debugf("received remove peers in peer unix sock, but remove peers is nil. => %d", m.PeerSock.Commands)
					break
				}

				err := s.scp.SyncRemotePeerConnecting(m.PeerSock.RemotePeers)
				if err != nil {
					// TODO: retry or do somethings
					s.dotlog.Logger.Errorf("failed to delete remote peer", m.PeerSock.Commands)
					break
				}
			case SetupRemotePeersConn:
				if m.PeerSock.RemotePeers == nil {
					s.dotlog.Logger.Debugf("received conn peers in peer unix sock, but conn peers is nil. => %d", m.PeerSock.Commands)
					break
				}

				err := s.scp.SetupRemotePeerConn(m.PeerSock.RemotePeers, m.PeerSock.Ip, m.PeerSock.Cidr)
				if err != nil {
					// TODO: retry
					s.dotlog.Logger.Errorf("failed to connection remote peer", m.PeerSock.Commands)
					break
				}
				break
			}
			break
		}

		err = encoder.Encode(m)
		if err != nil {
			s.dotlog.Logger.Errorf("failed to encode puncher signal. %s", err.Error())
			break
		}
	}
}

func (s *RcnSock) DialPeerSock(peers *PeerSock) error {
	conn, err := net.Dial("unix", sockaddr)
	defer conn.Close()
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	m := &RcnDialSock{
		MessageType: Peer,
		PeerSock:    peers,
	}

	err = encoder.Encode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("write dial peersock, called for [%s/%s]", m.PeerSock.Ip, m.PeerSock.Cidr)

	err = decoder.Decode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("read dial peersock, called for [%s/%s]", m.PeerSock.Ip, m.PeerSock.Cidr)

	return nil
}

func (s *RcnSock) Connect() error {
	err := s.cleanup()
	if err != nil {
		return err
	}

	listener, err := net.Listen("unix", sockaddr)
	if err != nil {
		return err
	}

	go func() {
		<-s.ch
		s.dotlog.Logger.Debugf("close the rcn socket")
		s.cleanup()
	}()

	s.dotlog.Logger.Debugf("starting rcn socket")
	for {
		conn, err := listener.Accept()
		if err != nil {
			s.dotlog.Logger.Errorf("failed to accept rcn socket. %s", err.Error())
		}

		s.dotlog.Logger.Debugf("accepted rcn sock => [%s]", conn.RemoteAddr().Network())

		go s.listen(conn)
	}
}
