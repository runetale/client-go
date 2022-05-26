package dotsignal

import (
	"fmt"
	"sync"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/negotiation"
	"github.com/Notch-Technologies/dotshake/client/grpc"
)

type DotSignal struct {
	signalClient grpc.SignalClientImpl

	mk string

	mu *sync.Mutex
	ch chan struct{}
}

func NewDotSignal(
	signalClient grpc.SignalClientImpl,
	mk string,
	ch chan struct{},
) *DotSignal {
	return &DotSignal{
		signalClient: signalClient,

		mk: mk,

		mu: &sync.Mutex{},
		ch: ch,
	}
}

func (s *DotSignal) ConnectDotSignal() {
	go func() {
		err := s.signalClient.StartConnect(s.mk, func(msg *negotiation.NegotiationResponse) error {
			s.mu.Lock()
			defer s.mu.Unlock()

			fmt.Println(msg)

			return nil
		})
		if err != nil {
			close(s.ch)
			return
		}
	}()
	s.signalClient.WaitStartConnect()
}
