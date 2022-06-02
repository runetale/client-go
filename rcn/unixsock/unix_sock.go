package unixsock

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

type PolymerSock struct {
	scp *controlplane.ControlPlane

	dotlog *dotlog.DotLog

	ch chan struct{}
}

func NewPolymerSock(
	dotlog *dotlog.DotLog,
	ch chan struct{},
	scp *controlplane.ControlPlane,
) *PolymerSock {
	return &PolymerSock{
		scp: scp,

		dotlog: dotlog,

		ch: ch,
	}
}

func (s *PolymerSock) cleanup() error {
	if _, err := os.Stat(sockaddr); err == nil {
		if err := os.RemoveAll(sockaddr); err != nil {
			return err
		}
	}
	return nil
}

func (s *PolymerSock) listen(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for {
		m := &RecvSocketMesage{}
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
			case RemovePeers:
				if m.PeerSock.RemovePeers == nil {
					s.dotlog.Logger.Debugf("received remove peers in peer unix sock, but remove peers is nil. => %d", m.PeerSock.Commands)
					break
				}

				err := s.scp.Delete(m.PeerSock.RemovePeers)
				if err != nil {
					// TODO: retry or do somethings
					s.dotlog.Logger.Errorf("failed to delete remote peer", m.PeerSock.Commands)
					break
				}
			case ConnPeers:
				if m.PeerSock.ConnPeers == nil {
					s.dotlog.Logger.Debugf("received conn peers in peer unix sock, but conn peers is nil. => %d", m.PeerSock.Commands)
					break
				}

				err := s.scp.SetupRemotePeerConn(m.PeerSock.ConnPeers, m.PeerSock.Ip, m.PeerSock.Cidr)
				if err != nil {
					// TODO: retry
					s.dotlog.Logger.Errorf("failed to connection remote peer", m.PeerSock.Commands)
					break
				}
				break
			}
			break
		case Signal:
			s.dotlog.Logger.Debugf("read from signal message => %d", m.SignalSock.Commands)
			break
		}

		err = encoder.Encode(m)
		if err != nil {
			s.dotlog.Logger.Errorf("failed to encode puncher signal. %s", err.Error())
			break
		}
	}
}

func (s *PolymerSock) DialPuncherSignal(signal *SignalSock) error {
	conn, err := net.Dial(protocol, sockaddr)
	defer conn.Close()
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	m := &RecvSocketMesage{
		MessageType: Signal,
		SignalSock:  signal,
	}

	err = encoder.Encode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("write purchase signal => %s", m.MessageType)

	err = decoder.Decode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("read purchase signal from server  => %d", signal.Commands)

	return nil
}

func (s *PolymerSock) DialPeerSock(peers *PeerSock) error {
	conn, err := net.Dial(protocol, sockaddr)
	defer conn.Close()
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	m := &RecvSocketMesage{
		MessageType: Peer,
		PeerSock:    peers,
	}

	err = encoder.Encode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("write remote peer => %s", m.MessageType)

	err = decoder.Decode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("read remote peer from server => %d", peers.Commands)

	return nil
}

func (s *PolymerSock) Connect() error {
	err := s.cleanup()
	if err != nil {
		return err
	}

	listener, err := net.Listen(protocol, sockaddr)
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
