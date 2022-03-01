package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/negotiation"
	"google.golang.org/grpc/metadata"

	"sync"
)

type Peer struct {
	ClientMachineKey string

	Stream negotiation.Negotiation_ConnectStreamServer
}

// NewPeer creates a new instance of a connected Peer
func NewPeer(key string, stream negotiation.Negotiation_ConnectStreamServer) *Peer {
	return &Peer{
		ClientMachineKey: key,
		Stream:           stream,
	}
}

// Registry registry that holds all currently connected Peers
type Registry struct {
	// Peer.key -> Peer
	Peers sync.Map
}

// NewRegistry creates a new connected Peer registry
func NewRegistry() *Registry {
	return &Registry{}
}

// Get gets a peer from the registry
func (registry *Registry) Get(peerId string) (*Peer, bool) {
	if load, ok := registry.Peers.Load(peerId); ok {
		return load.(*Peer), ok
	}
	return nil, false

}

func (registry *Registry) IsPeerRegistered(peerId string) bool {
	if _, ok := registry.Peers.Load(peerId); ok {
		return ok
	}
	return false
}

func (registry *Registry) Register(peer *Peer) {
	// can be that peer already exists but it is fine (e.g. reconnect)
	// todo investigate what happens to the old peer (especially Peer.Stream) when we override it
	registry.Peers.Store(peer.ClientMachineKey, peer)
	fmt.Println("** Founded dstPeer from ConnectStream")
	fmt.Printf("peer registered [%s]\n", peer.ClientMachineKey)
}

func (registry *Registry) Deregister(peer *Peer) {
	_, loaded := registry.Peers.LoadAndDelete(peer.ClientMachineKey)
	if loaded {
		fmt.Printf("peer deregistered [%s]\n", peer.ClientMachineKey)
	} else {
		fmt.Printf("attempted to remove non-existent peer [%s]", peer.ClientMachineKey)
	}

}

type NegotiationServiceServer struct {
	db       *database.Sqlite
	registry *Registry

	negotiation.UnimplementedNegotiationServer
}

func NewNegotiationServiceServer(db *database.Sqlite) *NegotiationServiceServer {
	return &NegotiationServiceServer{
		db:       db,
		registry: NewRegistry(),
	}
}

func (nss *NegotiationServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

const HeaderRegisterd = "registered"

func (nss *NegotiationServiceServer) ConnectStream(stream negotiation.Negotiation_ConnectStreamServer) error {

	// register peer with registry struct
	p, err := nss.registerPeer(stream)
	if err != nil {
		return err
	}

	defer func() {
		fmt.Printf("peer disconnected [%s]\n", p.ClientMachineKey)
		nss.registry.Deregister(p)
	}()

	header := metadata.Pairs(HeaderRegisterd, "1")
	err = stream.SendHeader(header)
	if err != nil {
		return err
	}

	fmt.Printf("peer connected [%s]\n", p.ClientMachineKey)

	for {
		msg, err := stream.Recv()
		fmt.Printf("recv connect stream from %s\n", msg.GetClientMachineKey())
		fmt.Println(msg)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		} else if err != nil {
			fmt.Println("connect stream err")
			fmt.Println(err)
			return err
		}
		if dstPeer, found := nss.registry.Get(msg.GetRemotekey()); found {
			fmt.Println("** Founded dstPeer from ConnectStream")
			//forward the message to the target peer
			err := dstPeer.Stream.Send(msg)
			if err != nil {
				fmt.Printf("error while forwarding message from peer [%s] to peer [%s] %v\n", p.ClientMachineKey, msg.GetRemotekey(), err)
			}
		} else {
			fmt.Println("Connect Stream Error")
			fmt.Printf("message from peer [%s] can't be forwarded to peer [%s] because destination peer is not connected\n", p.ClientMachineKey, msg.GetClientMachineKey())
		}
		fmt.Println("nothing")
	}
	<-stream.Context().Done()
	return stream.Context().Err()
}

const WgPubKey = "client-machine-key"

func (nss *NegotiationServiceServer) registerPeer(stream negotiation.Negotiation_ConnectStreamServer) (*Peer, error) {
	if meta, hasMeta := metadata.FromIncomingContext(stream.Context()); hasMeta {
		if wgPubKey, found := meta[WgPubKey]; found {
			p := NewPeer(wgPubKey[0], stream)
			nss.registry.Register(p)
			return p, nil
		} else {
			return nil, errors.New("missing connection header")
		}
	} else {
		return nil, errors.New("missing stream data")
	}
}

// Remote Peerに送る
func (nss *NegotiationServiceServer) Send(ctx context.Context, msg *negotiation.Body) (*negotiation.Body, error) {
	if !nss.registry.IsPeerRegistered(msg.Key) {
		return nil, fmt.Errorf("peer %s is not registered\n", msg.Key)
	}

	fmt.Printf("remote client machineKey: %s from %s\n", msg.GetRemotekey(), msg.GetClientMachineKey())
	if dstPeer, found := nss.registry.Get(msg.GetRemotekey()); found {
		err := dstPeer.Stream.Send(msg)
		if err != nil {
			fmt.Printf("error while forwarding message from peer [%s] to peer [%s] %v\n", msg.Key, msg.GetRemotekey(), err)
		}
		fmt.Println()
	} else {
		fmt.Println("Negotiation Send Error")
		fmt.Printf("message from peer [%s] can't be forwarded to peer [%s] because destination peer is not connected\n", msg.Key, msg.GetRemotekey())
	}

	return &negotiation.Body{}, nil
}
