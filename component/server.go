package component

import (
	"Bamboo/data"
	"Bamboo/utils"

	tp "github.com/henrylee2cn/teleport"
)

//Server Server
type Server struct {
	Peer tp.Peer
	tp.CallCtx
}

//NewServer NewServer
func NewServer() *Server {
	servers := new(Server)
	servers.Peer = tp.NewPeer(tp.PeerConfig{
		CountTime:  true,
		ListenPort: uint16(utils.Config.Port),
	})
	servers.Peer.RouteCall(servers)
	return servers
}

//Nodejoin 新节点汇报自己的存储信息
func (s *Server) Nodejoin(nodeData *data.RouteData) (int, *tp.Rerror) {
	routeBytes := nodeData.BuildRouteBytes()
	var readNodeData data.RouteData
	readNodeData.UnBuildRouteBytes(&routeBytes)
	updateRoute(&readNodeData)
	s.Session().Push("/client/receive", *nodeData)
	return 2, nil
}

//updateRoute 更新路由
func updateRoute(nodeData *data.RouteData) {

}

//UpdateItem 更新数据
func (s *Server) UpdateItem(nodeData *data.RouteData) (int, *tp.Rerror) {

	s.Session().Push("/client/receive", *nodeData)
	return 2, nil
}
