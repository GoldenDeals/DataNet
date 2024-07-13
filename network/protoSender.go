package network

import (
	"context"
	"time"
)

func (n *Node) RunSender(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			n.Log.Info("Sender stopped")
			return

		case <-time.NewTicker(2 * time.Second).C:
			for _, p := range n.Ka.Table().Entries() {
				if p.ID == n.N.ID().ID {
					continue
				}
				err := n.N.SendMessage(ctx, p.Address, ChatMessage{"hello"})
				if err != nil {
					n.Log.Infof("Failed to send message. err: %s", err)
				}
			}
		}
	}
}
