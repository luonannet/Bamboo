package hand

import (
	"Bamboo/utils"

	tp "github.com/henrylee2cn/teleport"
)

//Client Client
type Client struct {
	Peer tp.Peer
	tp.PushCtx
}

//NewClient NewClient
func NewClient() *Client {
	client := new(Client)
	client.Peer = tp.NewPeer(tp.PeerConfig{})
	client.Peer.RoutePush(client)
	return client
}

//Receive Receive
func (c *Client) Receive(arg *[]string) *tp.Rerror {
	utils.Debug("Receive:" + (*arg)[0])
	return nil
}
