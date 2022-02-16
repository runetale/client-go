package engine

import (
	"fmt"
	"sync"

	grpc_client "github.com/Notch-Technologies/wizy/cmd/server/grpc_client"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/negotiation"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/wislog"
)

type Engine struct {
	client *grpc_client.GrpcClient
	stream negotiation.Negotiation_ConnectStreamClient

	syncMsgMux *sync.Mutex

	wislog *wislog.WisLog
}

func NewEngine(
	log *wislog.WisLog,
	client *grpc_client.GrpcClient,
	stream negotiation.Negotiation_ConnectStreamClient,
) *Engine {
	return &Engine{
		client: client,
		stream: stream,

		syncMsgMux: &sync.Mutex{},

		wislog: log,
	}
}

// TODO:
// 1. create Engine
// 2. send stream message
// 3. create management json
// 4. return to stun and sturn
// 5. send to stun and turn udp request
// 6. send to signal offer
// 7. connection peer to peer test
// 8. management peers and how save the peer connectivity state? maybe sync mutex??
func (e *Engine) Start(publicMachineKey string) {
	e.syncMsgMux.Lock()
	defer e.syncMsgMux.Unlock()

	e.receiveClient(publicMachineKey)
	e.syncClient(publicMachineKey)
}

func (e *Engine) receiveClient(machineKey string) {
	go func() {
		err := e.client.Receive(machineKey, func(msg *negotiation.StreamMessage) error {
			e.syncMsgMux.Lock()
			defer e.syncMsgMux.Unlock()

			fmt.Println(msg.GetPrivateKey())
			fmt.Println(msg.GetClientMachineKey())

			return nil
		})
		if err != nil {
			return
		}
	}()

	e.client.WaitStreamConnected()
}

func (e *Engine) syncClient(machineKey string) {
	go func() {
		err := e.client.Sync(machineKey, func(update *peer.SyncResponse) error {
			e.syncMsgMux.Lock()
			defer e.syncMsgMux.Unlock()

			fmt.Println("sync")
			fmt.Println(update)

			// TODO: (shintard) send signal offer

			return nil
		})
		if err != nil {
			fmt.Println("stopping recive management server")
			return
		}
	}()
	fmt.Println("connecting management server")
}
