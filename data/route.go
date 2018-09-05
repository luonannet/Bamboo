package data

import (
	"Bamboo/utils"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"os"
)

var routeDataFile *os.File

func init() {
	routeDataFile = createData("route.data")
	gob.Register(new(RouteData))
}

const (
	//NodeDataLength 每一個nodeData所佔用的字節數
	NodeDataLength int64 = 62
)

//RouteData 节点数据
type RouteData struct {
	IP         string
	StartIndex uint64
	EndIndex   uint64
}

//routeList
type routeListStruct []*RouteData

var routeList routeListStruct

//BuildRouteBytes 分解出route的bytes
func (d *RouteData) BuildRouteBytes() []byte {
	nodeDataBf := bytes.NewBuffer(nil)
	nodeEncode := gob.NewEncoder(nodeDataBf)
	nodeEncode.Encode(d)
	return nodeDataBf.Bytes()
}

//SaveRoute SaveRoute
func (d *RouteData) SaveRoute(oribytes *[]byte) {
	nodeDataBf := bytes.NewBuffer(*oribytes)
	nodeDecode := gob.NewDecoder(nodeDataBf)
	nodeDecode.Decode(d)
	d.writeToFile()
}
func (d *RouteData) writeToFile() {
	routeDataFile.Write(d.toBytes())

}

func (d *RouteData) toBytes() []byte {
	ipbf := bytes.NewBufferString(d.IP)
	resultBF := bytes.NewBuffer(nil)
	resultBF.Truncate(0)
	resultBF.Write(ipbf.Bytes())
	blankBytes := make([]byte, 46-len(ipbf.Bytes()))
	resultBF.Write(blankBytes)
	binary.Write(resultBF, binary.BigEndian, d.StartIndex)
	binary.Write(resultBF, binary.BigEndian, d.EndIndex)
	return resultBF.Bytes()
}

//BuildItemDataBytes 分解出itemData的bytes
func (rl *routeListStruct) readFromFile(index int64) {
	var place int64
	var i int64
	fileinfo, _ := routeDataFile.Stat()
	routeList = routeList[0:0]
	for i = 0; i <= index; i++ {
		place = i * NodeDataLength
		if place >= fileinfo.Size() {
			continue
		}
		data := make([]byte, 62)
		_, err := routeDataFile.ReadAt(data, place)
		if err != nil || len(data) == 0 {
			utils.Debug("readFromFile error:", err)
			return
		}
		var d RouteData
		d.IP = string(data[0:46])
		startReader := bytes.NewReader(data[46:54])
		err = binary.Read(startReader, binary.BigEndian, &d.StartIndex)
		if err != nil {
			utils.Debug(" StartIndex ", err.Error())
			return
		}
		endReader := bytes.NewReader(data[54:62])
		err = binary.Read(endReader, binary.BigEndian, &d.EndIndex)
		if err != nil {
			utils.Debug(" EndIndex ", err.Error())
			return
		}
		routeList = append(routeList, &d)
	}
}
func findRoute(index uint64) (nextIP string) {
	//先逐个查找，将来再改成折半查找等更优化的办法
	for _, route := range routeList {
		if route.StartIndex >= index && index < route.EndIndex {
			return route.IP
		}
	}
	return ""
}
