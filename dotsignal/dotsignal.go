package dotsignal

import (
	"sync"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/negotiation"
	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
)

type DotSignal struct {
	signalClient grpc.SignalClientImpl

	mk string

	mu *sync.Mutex
	ch chan struct{}

	dotlog *dotlog.DotLog
}

func NewDotSignal(
	signalClient grpc.SignalClientImpl,
	mk string,
	ch chan struct{},
	mu *sync.Mutex,
	dotlog *dotlog.DotLog,
) *DotSignal {
	return &DotSignal{
		signalClient: signalClient,

		mk: mk,

		mu: mu,
		ch: ch,

		dotlog: dotlog,
	}
}

func (s *DotSignal) ConnectDotSignal() {
	go func() {
		err := s.signalClient.StartConnect(s.mk, func(msg *negotiation.NegotiationResponse) error {
			s.mu.Lock()
			defer s.mu.Unlock()

			return nil
		})
		if err != nil {
			close(s.ch)
			return
		}
	}()
	s.signalClient.WaitStartConnect()
}
