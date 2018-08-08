package main

import (
	// "Bamboo"
	"Bamboo/hand"
	"Bamboo/utils"
	"os"
	"path"
	"strconv"

	tp "github.com/henrylee2cn/teleport"
)

//Machine 本机
type Machine struct {
	tp.Peer
	neighbor []*tp.Session
	index    uint64
	client   *hand.Client
	server   *hand.Server
	dataFile *os.File
}

var machine Machine

func main() {
	machine.Peer = tp.NewPeer(tp.PeerConfig{
		CountTime:  true,
		ListenPort: uint16(utils.Config.Port),
	})
	// machine.index = 101
	machine.createLibrary()
	machine.pingNeighbor()
	machine.client = new(hand.Client)
	machine.server = new(hand.Server)
	machine.Peer.RouteCall(machine.server)
	machine.Peer.RoutePush(machine.client)
	go QueryRoute(machine.neighbor)
	machine.Peer.ListenAndServe()
}

//createLibrary 创建自己的分段仓库
func (m *Machine) createLibrary() {
	dir := "data"
	dataName := path.Join(dir, "bamboo.data")
	_, fileerr := os.Stat(dataName)
	os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if os.IsExist(fileerr) {
		m.dataFile, fileerr = os.OpenFile(dataName, os.O_RDWR, os.ModeDevice|os.ModePerm)
	} else {
		m.dataFile, fileerr = os.Create(dataName)
	}
	if fileerr != nil {
		utils.Log.Error("createLibrary  Error : %v", fileerr)
		return
	}
}

//pingNeighbor 发现邻居
func (m *Machine) pingNeighbor() {
	for _, add := range utils.Config.NeighborAddrs {
		sess, err := machine.Peer.Dial(add + ":" + strconv.Itoa(utils.Config.Port))
		if err != nil {
			utils.Log.Debug("pingNeighbor  Error : %v", err)
			continue
		}
		machine.neighbor = append(machine.neighbor, &sess)
	}
}

//AddNeighbor 添加邻居
func (m *Machine) AddNeighbor(arg *[]int) {

}

var result interface{}

//QueryRoute QueryRoute
func QueryRoute(sessList []*tp.Session) {
	for _, sess := range sessList {
		(*sess).Call("/server/queryroute", "queryItemData data", &result)
	}
}
