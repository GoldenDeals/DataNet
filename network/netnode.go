package network

import (
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"net"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	noise "github.com/GoldenDeals/noise-edits"
	"github.com/GoldenDeals/noise-edits/kademlia"
)

type Node struct {
	N   *noise.Node
	Log *zap.SugaredLogger
	Ka  *kademlia.Protocol
}

func SetupNode(log *zap.SugaredLogger) (*Node, error) {
	ip := net.IP{}
	err := ip.UnmarshalText([]byte(viper.GetString("node.listen")))
	if err != nil {
		return nil, err
	}

	nod := new(Node)
	nod.Log = log

	nod.N, err = noise.NewNode(
		noise.WithNodeLogger(log.Desugar()),
		noise.WithNodeAddress(viper.GetString("node.addr")),
		noise.WithNodeBindHost(ip),
		noise.WithNodeBindPort(viper.GetUint16("node.port")),
		noise.WithNodeIdleTimeout(viper.GetDuration("node.timeouts.idle")),
		noise.WithNodeMaxDialAttempts(viper.GetUint("node.max.dialAttempts")),
		noise.WithNodeMaxInboundConnections(viper.GetUint("node.max.inConnections")),
		noise.WithNodeMaxOutboundConnections(viper.GetUint("node.max.outConnections")),
		noise.WithNodeMaxRecvMessageSize(viper.GetUint32("node.max.messageSize")),
		noise.WithNodeNumWorkers(viper.GetUint("node.max.workers")),
		noise.WithNodePrivateKey(getPrivateKey(viper.GetString("node.privateKeySeed"))),
	)
	if err != nil {
		return nil, err
	}

	nod.N.RegisterMessage(ChatMessage{}, UnmarshalChatMessage)
	nod.N.Handle(nod.HandlerF)

	nod.SetupDiscovery()
	nod.N.Bind(nod.Ka.Protocol())

	err = nod.N.Listen()
	if err != nil {
		return nil, err
	}

	log.Infof("I'am %s", hex.EncodeToString(nod.N.ID().Marshal()))
	err = nod.pingInitialPeers()
	if err != nil {
		log.Warn("Failed to connect to peers! Running in standalone mode")
	}

	log.Debugf("Discovered %d peers", len(nod.Ka.Discover(kademlia.WithIteratorLogger(nod.Log.Desugar()))))
	go nod.RunSender(context.TODO())
	go nod.RunDiscovery(context.TODO())

	return nod, nil
}

func getPrivateKey(s string) noise.PrivateKey {
	if s == "gen" {
		_, key, err := noise.GenerateKeys(nil)
		if err != nil {
			panic(err)
		}

		return key
	}

	dec, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	privkey := ed25519.NewKeyFromSeed(dec)
	var key noise.PrivateKey
	if len(key) != len(privkey) {
		panic("keys length mismatch")
	}

	for i := 0; i < len(key); i++ {
		key[i] = privkey[i]
	}

	return key
}
