package hand

import (
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
	server := new(Server)
	server.Peer = tp.NewPeer(tp.PeerConfig{
		CountTime:  true,
		ListenPort: uint16(utils.Config.Port),
	})
	server.Peer.RouteCall(server)
	return server
}

//queryroute 从指定的节点获取自己的路由
func (s *Server) Queryroute(arg *[]string) (int, *tp.Rerror) {
	s.Session().Push("/client/receive", *arg)
	return 2, nil
}
