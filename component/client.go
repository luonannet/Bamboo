package component

import (
	"Bamboo/data"
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

//JoinedIn 节点加入net之后
func (c *Client) JoinedIn(arg *[]string) *tp.Rerror {
	utils.Debug("Receive:" + (*arg)[0])
	return nil
}

//ReportData 携带本机数据加入网络
func ReportData(ip string) *data.RouteData {
	thisData := new(data.RouteData)
	thisData.StartIndex = utils.Config.StartIndex
	thisData.EndIndex = utils.Config.EndIndex
	thisData.IP = ip
	return thisData
}
