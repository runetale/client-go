package channel

import (
	"fmt"
	"sync"

	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
)

type UpdateMessage struct {
	Update *peer.SyncResponse
}

type PeersUpdateManager struct {
	peerChannels map[string]chan *UpdateMessage
	channelsMux  *sync.Mutex
}

// NewPeersUpdateManager returns a new instance of PeersUpdateManager
func NewPeersUpdateManager() *PeersUpdateManager {
	return &PeersUpdateManager{
		peerChannels: make(map[string]chan *UpdateMessage),
		channelsMux:  &sync.Mutex{},
	}
}

// SendUpdate sends update message to the peer's channel
func (p *PeersUpdateManager) SendUpdate(peer string, update *UpdateMessage) error {
	fmt.Println("Send Update")
	p.channelsMux.Lock()
	defer p.channelsMux.Unlock()
	if channel, ok := p.peerChannels[peer]; ok {
		channel <- update
		return nil
	}
	fmt.Printf("peer %s has no channel", peer)
	return nil
}

func (p *PeersUpdateManager) CreateChannel(peerKey string) chan *UpdateMessage {
	p.channelsMux.Lock()
	defer p.channelsMux.Unlock()

	if channel, ok := p.peerChannels[peerKey]; ok {
		delete(p.peerChannels, peerKey)
		close(channel)
	}
	//mbragin: todo shouldn't it be more? or configurable?
	channel := make(chan *UpdateMessage, 100)
	p.peerChannels[peerKey] = channel

	fmt.Printf("opened updates channel for a peer %s\n", peerKey)
	return channel
}

func (p *PeersUpdateManager) CloseChannel(peerKey string) {
	p.channelsMux.Lock()
	defer p.channelsMux.Unlock()
	if channel, ok := p.peerChannels[peerKey]; ok {
		delete(p.peerChannels, peerKey)
		close(channel)
	}

	fmt.Printf("closed updates channel of a peer %s", peerKey)
}
