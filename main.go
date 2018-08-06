package main

import (
	"Bamboo/utils"
	"fmt"
	"log"
	"strconv"
	"time"

	tp "github.com/henrylee2cn/teleport"
)

//Machine 本机
type Machine struct {
	tp.Peer
	neighbor []*tp.Session

	caller *callerStruct
	sender *senderStruct
}
type callerStruct struct {
	tp.CallCtx
}
type senderStruct struct {
	tp.PushCtx
}

var machine Machine

func main() {
	machine.Peer = tp.NewPeer(tp.PeerConfig{
		CountTime:  true,
		ListenPort: uint16(utils.Config.Port),
	})

	machine.pingNeighbor()
	machine.caller = new(callerStruct)
	machine.sender = new(senderStruct)
	machine.Peer.RouteCall(machine.caller)
	machine.Peer.RoutePush(machine.sender)
	machine.Peer.ListenAndServe()
}

//pingNeighbor 发现邻居
func (m *Machine) pingNeighbor() {
	for _, add := range utils.Config.NeighborAddrs {
		sess, err := machine.Peer.Dial(add + ":" + strconv.Itoa(utils.Config.Port))
		if err != nil {
			log.Printf("pingNeighbor  Error : %v", err)
		}
		machine.neighbor = append(machine.neighbor, &sess)
	}
}

//createLibrary 创建自己的仓库
func (m *Machine) createLibrary() {

}

//AddNeighbor 添加邻居,
func (m *Machine) AddNeighbor(arg *[]int) (int, *tp.Rerror) {
	if m.caller.Query().Get("addNeighbor") == "yes" {
		m.caller.Session().Push(
			"/sender/status",
			fmt.Sprintf("%d numbers are being added...", len(*arg)),
		)
		time.Sleep(time.Millisecond * 10)
	}
	var r int
	for _, a := range *arg {
		r += a
	}
	return r, nil
}

//BoardCast BoardCast
func (m *Machine) BoardCast(arg *[]int) (int, *tp.Rerror) {
	if m.caller.Query().Get("push_status") == "yes" {
		m.caller.Session().Push(
			"/sender/status",
			fmt.Sprintf("%d numbers are being added...", len(*arg)),
		)
		time.Sleep(time.Millisecond * 10)
	}
	var r int
	for _, a := range *arg {
		r += a
	}
	return r, nil
}
func (s *Machine) Status(arg *string) *tp.Rerror {
	tp.Printf("server status: %s", *arg)
	return nil
}
