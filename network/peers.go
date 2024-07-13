package network

import (
	"context"
	"encoding/hex"

	"github.com/GoldenDeals/noise-edits"
	"github.com/spf13/viper"
)

func (n *Node) pingInitialPeers() error {
	initPeersId := viper.GetStringSlice("node.initialPeers")
	for _, peerId := range initPeersId {
		byt, err := hex.DecodeString(peerId)
		if err != nil {
			return err
		}
		id, err := noise.UnmarshalID(byt)
		if err != nil {
			return err
		}

		_, err = n.N.Ping(context.TODO(), id.Address)
		if err != nil {
			return err
		}
	}
	return nil
}
