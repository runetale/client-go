package wonderwall

import (
	"encoding/gob"
	"io"
	"net"
	"os"

	"github.com/Notch-Technologies/dotshake/dotengine/proxy"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/pion/ice/v2"
)

// TODO: (shinta) is this safe?
// appropriate permission and feel it would be better to
// have a process that creates a file
//
const sockaddr = "/tmp/wonderwall.sock"

type WSocketMessageType int

const (
	// called after SetupRemotePeersConn in rcn is called
	StartWonderWall WSocketMessageType = 0
	StopWonderWall  WSocketMessageType = 1
)

type WDialSock struct {
	MessageType WSocketMessageType

	StartWonderWallSock *StartWonderWallSock
	CloseWonderWallSock *CloseWonderWallSock
}

type StartWonderWallSock struct {
	Uname string
	Pwd   string

	Agent     *ice.Agent
	WireProxy *proxy.WireProxy

	RemoteWgPubKey string
	WgPubKey       string

	DisconnectFunc func() error
	ConnectFunc    func() error

	Dotlog *dotlog.DotLog
}

type CloseWonderWallSock struct {
	CloseCh chan struct{}
}

type WonderWallSock struct {
	wonderWall *WonderWall

	agent     *ice.Agent
	wireProxy *proxy.WireProxy

	dotlog *dotlog.DotLog

	ch chan struct{}
}

// called dotengine with read sock or rcn with write sock
//
func NewWonderWallSock(
	wonderwall *WonderWall,
	dotlog *dotlog.DotLog,
) *WonderWallSock {
	return &WonderWallSock{
		wonderWall: wonderwall,
		dotlog:     dotlog,
	}
}

func (w *WonderWallSock) cleanup() error {
	if _, err := os.Stat(sockaddr); err == nil {
		if err := os.RemoveAll(sockaddr); err != nil {
			return err
		}
	}
	return nil
}

// called dotengine
//
func (w *WonderWallSock) listen(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for {
		m := &WDialSock{}
		err := decoder.Decode(m)
		if err != nil {
			if err == io.EOF {
				w.dotlog.Logger.Debugf("close the wonderwall socket client")
				break
			}
			break
		}

		switch m.MessageType {
		case StartWonderWall:
			w.wonderWall.setup(
				m.StartWonderWallSock.Agent,
				m.StartWonderWallSock.WireProxy,
				m.StartWonderWallSock.RemoteWgPubKey,
				m.StartWonderWallSock.WgPubKey,
				m.StartWonderWallSock.DisconnectFunc,
				m.StartWonderWallSock.ConnectFunc,
			)

			err := w.wonderWall.Start(m.StartWonderWallSock.Uname, m.StartWonderWallSock.Pwd)
			if err != nil {
				// TODO: retry
				w.dotlog.Logger.Errorf("failed to start wonderwall sock")
				break
			}
			break
		case StopWonderWall:
			err := w.wonderWall.Close()
			if err != nil {
				// TODO: retry
				w.dotlog.Logger.Errorf("failed to stop wonderwall sock")
				break
			}
			break
		}

		err = encoder.Encode(m)
		if err != nil {
			w.dotlog.Logger.Errorf("failed to encode wondersock. %s", err.Error())
			break
		}
	}
}

// called rcn
//
func (s *WonderWallSock) DialStartWonderWall(sock *StartWonderWallSock) error {
	conn, err := net.Dial("unix", sockaddr)
	defer conn.Close()
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	m := &WDialSock{
		MessageType:         StartWonderWall,
		StartWonderWallSock: sock,
	}

	err = encoder.Encode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("write startwonderwall from rcn => %s", m.MessageType)

	err = decoder.Decode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("read startwonderwall from rcn => [%s:%s]", sock.Uname, sock.Pwd)

	return nil
}

// called rcn
//
func (s *WonderWallSock) DialStopWonderWall(sock *CloseWonderWallSock) error {
	conn, err := net.Dial("unix", sockaddr)
	defer conn.Close()
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	m := &WDialSock{
		MessageType:         StartWonderWall,
		CloseWonderWallSock: sock,
	}

	err = encoder.Encode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("write stopwonderwall => %s", m.MessageType)

	err = decoder.Decode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("read stopwonderwall from rcn => [%s:%s]")

	return nil
}

// called dotengine
//
func (w *WonderWallSock) Connect() error {
	err := w.cleanup()
	if err != nil {
		return err
	}

	listener, err := net.Listen("unix", sockaddr)
	if err != nil {
		return err
	}

	go func() {
		<-w.ch
		w.dotlog.Logger.Debugf("close the wonderwall socket")
		w.cleanup()
	}()

	w.dotlog.Logger.Debugf("starting wonderwall socket")
	for {
		conn, err := listener.Accept()
		if err != nil {
			w.dotlog.Logger.Errorf("failed to accept wonderwall socket. %s", err.Error())
		}

		w.dotlog.Logger.Debugf("accepted wonderwall sock => [%s]", conn.RemoteAddr().Network())

		go w.listen(conn)
	}
}
