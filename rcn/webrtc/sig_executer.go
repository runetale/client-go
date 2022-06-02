package webrtc

import (
	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/pion/ice/v2"
)

// this package provides the functions needed for udp hole punching using webrtc
// dependent on signal client
//

type SigExecuter struct {
	signalClient grpc.SignalClientImpl
	dstmk        string
	srcmk        string

	dotlog *dotlog.DotLog
}

func NewSigExecuter(
	signalClient grpc.SignalClientImpl,
	dstmk string,
	srcmk string,
	dotlog *dotlog.DotLog,
) *SigExecuter {
	return &SigExecuter{
		signalClient: signalClient,
		dstmk:        dstmk,
		srcmk:        srcmk,

		dotlog: dotlog,
	}
}

func (s *SigExecuter) Candidate(
	candidate ice.Candidate,
) {
	if candidate == nil {
		s.dotlog.Logger.Errorf("ice candidate failed")
		return
	}

	go func() {
		err := s.signalClient.Candidate(s.dstmk, s.srcmk, candidate)
		if err != nil {
			s.dotlog.Logger.Errorf("failed to candidate against signal server, becasuse %s", err.Error())
			return
		}
	}()
}

func (s *SigExecuter) Offer(
	uFlag string,
	pwd string,
) error {
	return s.signalClient.Offer(s.dstmk, s.srcmk, uFlag, pwd)
}

func (s *SigExecuter) Answer(
	uFlag string,
	pwd string,
) error {
	return s.signalClient.Answer(s.dstmk, s.srcmk, uFlag, pwd)
}
