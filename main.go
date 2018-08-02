package main

import (
	"fmt"
	"time"

	tp "github.com/henrylee2cn/teleport"
)

type Machine struct {
	tp.Peer
	neighbor     []*tp.Session
	neighborAddr []string
}

var machine Machine

func main() {
	machine.Peer = tp.NewPeer(tp.PeerConfig{
		CountTime:  true,
		ListenPort: 9527,
	})

	findNeighbor()
	machine.Peer.RouteCall(new(caller))
	machine.Peer.RoutePush(new(sender))
	machine.Peer.ListenAndServe()
}

type caller struct {
	tp.CallCtx
}

type sender struct {
	tp.PushCtx
}

//findNeighbor 发现邻居
func findNeighbor() {
	time.Sleep(time.Second * 1)
	machine.neighborAddr = []string{":9528", "9529"}
	for _, add := range machine.neighborAddr {
		sess, err := machine.Peer.Dial(add)
		if err != nil {
			tp.Fatalf("%v", err)
		}
		machine.neighbor = append(machine.neighbor, &sess)
	}
}

//AddNeighbor 添加邻居
func (m *caller) AddNeighbor(arg *[]int) (int, *tp.Rerror) {
	if m.Query().Get("addNeighbor") == "yes" {
		m.Session().Push(
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
func (m *caller) BoardCast(arg *[]int) (int, *tp.Rerror) {
	if m.Query().Get("push_status") == "yes" {
		m.Session().Push(
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
func (s *sender) Status(arg *string) *tp.Rerror {
	tp.Printf("server status: %s", *arg)
	return nil
}
