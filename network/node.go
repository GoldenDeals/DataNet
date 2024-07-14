package network

import (
	"github.com/GoldenDeals/DataNet/utils"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	peerstore "github.com/libp2p/go-libp2p/core/peer"
	"github.com/rs/zerolog/log"
)

var lg utils.LoggerGroup

type Node struct {
	host.Host
}

func (n *Node) Init() (*Node, error) {
	lg = utils.NewLoggerGroup("net")
	var err error
	n.Host, err = libp2p.New()
	if err != nil {
		return nil, err
	}

	adds, err := n.getDescriptor()
	if err != nil {
		return nil, err
	}
	lg(log.Debug()).Msgf("I am: %s", adds)

	return n, nil
}
func (n *Node) Shutdown() {
	err := n.Close()
	if err != nil {
		lg(log.Error()).Err(err).Msg("Error while closing host")
	}
}

func (n *Node) getDescriptor() (string, error) {
	peerInfo := peerstore.AddrInfo{
		ID:    n.ID(),
		Addrs: n.Addrs(),
	}
	addrs, err := peerstore.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		return "", err
	}

	return addrs[0].String(), nil
}
