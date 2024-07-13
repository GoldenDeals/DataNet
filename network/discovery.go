package network

import (
	"context"
	"time"

	noise "github.com/GoldenDeals/noise-edits"
	"github.com/GoldenDeals/noise-edits/kademlia"
	"github.com/spf13/viper"
)

func (n *Node) SetupDiscovery() {
	n.Ka = kademlia.New(
		kademlia.WithProtocolLogger(n.Log.Desugar()),
		kademlia.WithProtocolEvents(kademlia.Events{
			OnPeerAdmitted: n.onPeerAdded,
			OnPeerEvicted:  n.onPeerEvicted,
		}),
		kademlia.WithProtocolPingTimeout(viper.GetDuration("node.timeouts.ping")*time.Second),
	)
}

func (n *Node) RunDiscovery(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			n.Log.Info("Discovery stopped")
			return
		case <-time.NewTicker(10 * time.Second).C:
			peers := n.Ka.Discover()

			c := 0
			for _, p := range peers {
				if n.isInTable(p) {
					continue
				}
				c++
				_, err := n.N.Ping(ctx, p.Address)
				if err != nil {
					n.Log.Infof("Failed to ping new peer. Err: %s", err)
				}
			}

			if c > 0 {
				n.Log.Infof("Discovered %d new peers", c)
			}
		}
	}
}

func (n *Node) isInTable(p noise.ID) bool {
	for _, pi := range n.Ka.Table().Entries() {
		if pi.ID == p.ID {
			return true
		}
	}
	return false
}

func (n *Node) onPeerAdded(id noise.ID) {
	n.Log.Debugf("Peer with addr: %s Added!", id.Address)
}

func (n *Node) onPeerEvicted(id noise.ID) {
	n.Log.Debugf("Peer with addr: %s Evicetd!", id.Address)
}
