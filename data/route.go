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
	var resultBytes []byte
	ipbf := bytes.NewBufferString(d.IP)
	blankLen := 46 - len(ipbf.Bytes()) //ipv6的长度46
	blankBytes := make([]byte, blankLen, blankLen)
	ipbf.Write(blankBytes) //空白补齐
	resultBytes = ipbf.Bytes()

	startbf := bytes.NewBuffer(nil)
	binary.Write(startbf, binary.BigEndian, d.StartIndex)
	resultBytes = append(resultBytes, startbf.Bytes()...)

	endbf := bytes.NewBuffer(nil)
	binary.Write(endbf, binary.BigEndian, d.EndIndex)
	resultBytes = append(resultBytes, endbf.Bytes()...)
	routeDataFile.Write(resultBytes)
	var testreadRoutedata RouteData
	testreadRoutedata.readFromFile(0)
}

//BuildItemDataBytes 分解出itemData的bytes
func (d *RouteData) BuildItemDataBytes() {

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
