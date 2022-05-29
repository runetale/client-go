package unixsock

import (
	"encoding/gob"
	"io"
	"net"
	"os"

	"github.com/Notch-Technologies/dotshake/dotlog"
)

const protocol = "unix"
const sockaddr = "/tmp/polymer.sock"

type PuncherSignal struct {
	PuncherSignal string
}

type PolymerSock struct {
	dotlog *dotlog.DotLog

	ch chan struct{}
}

func NewPolyerSock(
	dotlog *dotlog.DotLog,
	ch chan struct{},
) *PolymerSock {
	return &PolymerSock{
		dotlog: dotlog,
		ch:     ch,
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

func (s *PolymerSock) puncherSignal(conn net.Conn) {
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	for {
		m := &PuncherSignal{}
		err := decoder.Decode(m)
		if err != nil {
			if err == io.EOF {
				s.dotlog.Logger.Debugf("close the polymer socket client")
				break
			}
			break
		}

		s.dotlog.Logger.Debugf("read puncher signal from client => %s", m.PuncherSignal)

		m.PuncherSignal = "hoge"

		err = encoder.Encode(m)
		if err != nil {
			s.dotlog.Logger.Errorf("failed to encode puncher signal. %s", err.Error())
			break
		}

		s.dotlog.Logger.Debugf("write puncher signal from server %s", m.PuncherSignal)
	}
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
		s.dotlog.Logger.Debugf("close the polymer socket")
		s.cleanup()
	}()

	s.dotlog.Logger.Debugf("starting polymer socket")
	for {
		conn, err := listener.Accept()
		if err != nil {
			s.dotlog.Logger.Errorf("failed to accept polymer socket. %s", err.Error())
		}

		s.dotlog.Logger.Debugf("accepted polymer sock => [%s]", conn.RemoteAddr().Network())

		go s.puncherSignal(conn)
	}

}

func (s *PolymerSock) Dial() error {
	conn, err := net.Dial(protocol, sockaddr)
	defer conn.Close()
	if err != nil {
		return err
	}

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	m := &PuncherSignal{
		PuncherSignal: "aa",
	}

	err = encoder.Encode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("write purchase signal => %s", m.PuncherSignal)

	err = decoder.Decode(m)
	if err != nil {
		return err
	}

	s.dotlog.Logger.Debugf("read purchase signal from server  => %s", m.PuncherSignal)

	return nil
}
