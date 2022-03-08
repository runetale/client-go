package service

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/Notch-Technologies/wizy/cmd/signaling/key"
	"github.com/Notch-Technologies/wizy/cmd/signaling/pb/negotiation"
	"github.com/Notch-Technologies/wizy/cmd/signaling/registry"
	"google.golang.org/grpc/metadata"
)

type NegotiationServerServiceCaller interface {
	ConnectStream(stream negotiation.Negotiation_ConnectStreamServer) error
	Send(ctx context.Context, msg *negotiation.Body) (*negotiation.Body, error)
}

type NegotiationServerService struct {
	registry *registry.Registry

	negotiation.UnimplementedNegotiationServer
}

func NewNegotiationServerService() *NegotiationServerService {
	return &NegotiationServerService{
		registry: registry.NewRegistry(),
	}
}

func (n *NegotiationServerService) AuthFuncOverride(
	ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (n *NegotiationServerService) registerPeer(stream negotiation.Negotiation_ConnectStreamServer) (*registry.Peer, error) {
	if meta, hasMeta := metadata.FromIncomingContext(stream.Context()); hasMeta {
		if wgPubKey, found := meta[key.WgPubKey]; found {
			p := registry.NewPeer(wgPubKey[0], stream)
			n.registry.Register(p)
			return p, nil
		} else {
			return nil, errors.New("missing connection header")
		}
	} else {
		return nil, errors.New("missing stream data")
	}
}

func (n *NegotiationServerService) ConnectStream(stream negotiation.Negotiation_ConnectStreamServer) error {
	// register peer with registry struct
	p, err := n.registerPeer(stream)
	if err != nil {
		return err
	}

	defer func() {
		fmt.Printf("peer disconnected [%s]\n", p.ClientMachineKey)
		n.registry.Deregister(p)
	}()

	header := metadata.Pairs(key.HeaderRegisterd, "1")
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
		if dstPeer, found := n.registry.Get(msg.GetRemotekey()); found {
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
	}
	<-stream.Context().Done()
	return stream.Context().Err()
}

func (n *NegotiationServerService) Send(ctx context.Context, msg *negotiation.Body) (*negotiation.Body, error) {
	if !n.registry.IsPeerRegistered(msg.Key) {
		return nil, fmt.Errorf("peer %s is not registered", msg.Key)
	}

	fmt.Printf("remote client machineKey: %s from %s\n", msg.GetRemotekey(), msg.GetClientMachineKey())
	if dstPeer, found := n.registry.Get(msg.GetRemotekey()); found {
		err := dstPeer.Stream.Send(msg)
		if err != nil {
			fmt.Printf("error while forwarding message from peer [%s] to peer [%s] %v\n", msg.Key, msg.GetRemotekey(), err)
		}
	} else {
		fmt.Println("Negotiation Send Error")
		fmt.Printf("message from peer [%s] can't be forwarded to peer [%s] because destination peer is not connected\n", msg.Key, msg.GetRemotekey())
	}

	return &negotiation.Body{}, nil
}
