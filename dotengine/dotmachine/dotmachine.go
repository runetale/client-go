package dotmachine

import (
	"sync"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/machine"
	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
)

type DotMachine struct {
	serverClient grpc.ServerClientImpl

	mk string

	mu *sync.Mutex
	ch chan struct{}

	dotlog *dotlog.DotLog
}

func NewDotMachine(
	serverClient grpc.ServerClientImpl,
	mk string,
	ch chan struct{},
	mu *sync.Mutex,
	dotlog *dotlog.DotLog,
) *DotMachine {
	return &DotMachine{
		serverClient: serverClient,

		mk: mk,

		mu: mu,
		ch: ch,

		dotlog: dotlog,
	}
}

func (m *DotMachine) Up() {
	go func() {
		err := m.SyncMachine(func(res *machine.SyncMachinesResponse) error {
			m.dotlog.Logger.Debugf("connected sync machine")
			m.mu.Lock()
			defer m.mu.Unlock()

			return nil
		})
		if err != nil {
			close(m.ch)
			return
		}
	}()
}

func (m *DotMachine) SyncMachine(handler func(res *machine.SyncMachinesResponse) error) error {
	err := m.serverClient.SyncMachines(m.mk, handler)
	if err != nil {
		return err
	}

	return nil
}
