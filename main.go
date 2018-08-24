package main

import (
	"Bamboo/component"
)

func main() {
	machine := component.NewMachine()
	machine.PingNeighbor()
	go machine.JoinNet(component.ReportData(machine.IP))
	machine.Server.Peer.ListenAndServe()
}
