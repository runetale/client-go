package registry

import (
	"fmt"
	"sync"
)

type Registry struct {
	// Peer.key -> Peer
	Peers sync.Map
}

func NewRegistry() *Registry {
	return &Registry{}
}

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
