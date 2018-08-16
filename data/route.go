package data

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
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

//BuildRouteBytes 分解出route的bytes
func (d *RouteData) BuildRouteBytes() []byte {
	nodeDataBf := bytes.NewBuffer(nil)
	nodeEncode := gob.NewEncoder(nodeDataBf)
	nodeEncode.Encode(d)
	return nodeDataBf.Bytes()
}

//UnBuildRouteBytes UnBuildRouteBytes
func (d *RouteData) UnBuildRouteBytes(oribytes *[]byte) {
	nodeDataBf := bytes.NewBuffer(*oribytes)
	nodeDecode := gob.NewDecoder(nodeDataBf)
	nodeDecode.Decode(d)
	d.writeToFile()
}
func (d *RouteData) writeToFile() {
	routeDataFile.Write(d.toBytes())
	var testreadRoutedata RouteData
	testreadRoutedata.readFromFile(0)
}

func (d *RouteData) toBytes() []byte {
	resultBytes := make([]byte, 62, 62)
	ipbf := bytes.NewBufferString(d.IP)
	resultBF := bytes.NewBuffer(resultBytes)
	resultBF.Truncate(0)
	resultBF.Write(ipbf.Bytes())
	binary.Write(resultBF, binary.BigEndian, d.StartIndex)
	binary.Write(resultBF, binary.BigEndian, d.EndIndex)
	return resultBF.Bytes()
}

//BuildItemDataBytes 分解出itemData的bytes
func (d *RouteData) readFromFile(index int64) {
	place := index * NodeDataLength
	data := make([]byte, 62)
	_, err := routeDataFile.ReadAt(data, place)
	if err != nil || len(data) == 0 {
		fmt.Println("readFromFile error:", err)
		return
	}
	d.IP = string(data[0:46])
	startReader := bytes.NewReader(data[46:54])
	err = binary.Read(startReader, binary.BigEndian, &d.StartIndex)
	if err != nil {
		fmt.Println(" StartIndex ", err.Error())
		return
	}
	endReader := bytes.NewReader(data[54:62])
	err = binary.Read(endReader, binary.BigEndian, &d.EndIndex)
	if err != nil {
		fmt.Println(" EndIndex ", err.Error())
		return
	}
}
func findRoute(index string) (nextIP []string) {
	fileInfo, _ := routeDataFile.Stat()
	totalSize := fileInfo.Size()
	if totalSize == 0 {
		return nil
	}
	if totalSize == int64(NodeDataLength) {
		_ = (totalSize - NodeDataLength) / 2
	}
	return
}
