package hand

import tp "github.com/henrylee2cn/teleport"

//Server Server
type Server struct {
	tp.CallCtx
}

//queryroute 从指定的节点获取自己的路由
func (s *Server) queryroute(arg *string) (int, *tp.Rerror) {
	s.Session().Push("/client/receive", "---------"+(*arg))
	return 2, nil
}
