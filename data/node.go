package data

import (
	"Bamboo/utils"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"os"
	"path"
)

var dataFile *os.File

func init() {
	dir := "store"
	dataName := path.Join(dir, "bamboo.data")
	_, fileerr := os.Stat(dataName)
	os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if os.IsExist(fileerr) {
		dataFile, fileerr = os.OpenFile(dataName, os.O_CREATE|os.O_APPEND, os.ModeDevice|os.ModePerm)
	} else {
		dataFile, fileerr = os.Create(dataName)
	}
	if fileerr != nil {
		utils.Error(fmt.Sprintf("createLibrary  Error : %v ", fileerr))
		return
	}
	gob.Register(new(NodeData))

}

//NodeData 节点数据
type NodeData struct {
	IP         string
	StartIndex uint64
	EndIndex   uint64
}

//BuildRouteBytes 分解出route的bytes
func (d *NodeData) BuildRouteBytes() []byte {
	nodeDataBf := bytes.NewBuffer(nil)
	nodeEncode := gob.NewEncoder(nodeDataBf)
	nodeEncode.Encode(d)
	return nodeDataBf.Bytes()
}

//UnBuildRouteBytes UnBuildRouteBytes
func (d *NodeData) UnBuildRouteBytes(oribytes *[]byte) {
	nodeDataBf := bytes.NewBuffer(*oribytes)
	nodeDecode := gob.NewDecoder(nodeDataBf)
	nodeDecode.Decode(d)
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
	dataFile.Write(resultBytes)
}

//BuildItemDataBytes 分解出itemData的bytes
func (d *NodeData) BuildItemDataBytes() {

}

// //SaveRouteBytes 分解出route的bytes
// func SaveRouteBytes(databytes []byte) {
// 	dataFile.Write(databytes)
// }
