package component

import (
	"Bamboo/data"
	"Bamboo/utils"
	"fmt"
	"net"
	"strconv"
	"strings"

	tp "github.com/henrylee2cn/teleport"
)

//Machine 本机
type Machine struct {
	neighbor []*tp.Session
	IP       string
	Client   *Client
	Server   *Server
}

//NewMachine NewMachine
func NewMachine() *Machine {
	machine := new(Machine)
	machine.IP = getMachineIP()
	machine.Client = NewClient()
	machine.Server = NewServer()
	return machine
}
func getMachineIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}

//PingNeighbor 发现邻居
func (m *Machine) PingNeighbor() {

	for _, add := range utils.Config.NeighborAddrs {
		if strings.EqualFold(add, m.IP) {
			//如果是本機
			continue
		}
		sess, err := m.Client.Peer.Dial(add + ":" + strconv.Itoa(utils.Config.Port))
		if err != nil {
			fmt.Println(fmt.Sprintf("pingNeighbor  Error : %v ", err))
			continue
		}
		m.neighbor = append(m.neighbor, &sess)
	}
}

var result interface{}

//JoinNet 向预设的节点发出请求通知，报告自己的ip，获取自己节点的相关路由
func (m *Machine) JoinNet(data *data.RouteData) {

	for _, sess := range m.neighbor {
		(*sess).Call("/server/nodejoin", data, &result)
	}
}
