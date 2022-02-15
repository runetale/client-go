package engine

import (
	"fmt"
	"sync"

	grpc_client "github.com/Notch-Technologies/wizy/cmd/server/grpc_client"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/negotiation"
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

func (e *Engine) Start() {
}

func (e *Engine) ReceiveClient() {
	go func() {
		err := e.client.Receive(e.stream, func(msg *negotiation.StreamMessage) error {
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

	fmt.Println("connecting signal server")
}
