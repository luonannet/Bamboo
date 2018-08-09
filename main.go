package main

import (
	"fmt"
	// "Bamboo"
	"Bamboo/hand"
	"Bamboo/utils"
	"net"
	"os"
	"path"
	"strconv"

	tp "github.com/henrylee2cn/teleport"
)

//Machine 本机
type Machine struct {
	neighbor []*tp.Session
	ip       string
	client   *hand.Client
	server   *hand.Server
	dataFile *os.File
}

var machine Machine

func main() {
	machine.ip = getMachineIP()
	machine.createLibrary()
	machine.client = hand.NewClient()
	machine.server = hand.NewServer()
	machine.pingNeighbor()
	go machine.joinNet()
	machine.server.Peer.ListenAndServe()
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
		utils.Error(fmt.Sprintf("createLibrary  Error : %v ", fileerr))
		return
	}
}

//pingNeighbor 发现邻居
func (m *Machine) pingNeighbor() {
	for _, add := range utils.Config.NeighborAddrs {
		sess, err := machine.client.Peer.Dial(add + ":" + strconv.Itoa(utils.Config.Port))
		if err != nil {
			utils.Debug(fmt.Sprintf("pingNeighbor  Error : %v ", err))
			continue
		}
		machine.neighbor = append(machine.neighbor, &sess)
	}
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

var result interface{}

//joinNet 向预设的节点发出请求通知，报告自己的ip，获取自己节点的相关路由
func (m *Machine) joinNet() {
	myip := []string{m.ip, "30000", "40000"}
	for _, sess := range m.neighbor {
		(*sess).Call("/server/queryroute", &myip, &result)
	}
}
