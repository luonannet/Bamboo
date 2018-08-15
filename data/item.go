package data

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"os"
	"path"
)

var itemDataFile *os.File

func init() {
	itemDataFile = createData("item.data")
	gob.Register(new(ItemData))

}
func createData(fileName string) (file *os.File) {
	dir := "store"
	dataName := path.Join(dir, fileName)
	_, fileerr := os.Stat(dataName)
	os.MkdirAll(dir, os.ModeDir|os.ModePerm)
	if os.IsExist(fileerr) {
		file, fileerr = os.OpenFile(dataName, os.O_CREATE|os.O_APPEND, os.ModeDevice|os.ModePerm)
	} else {
		file, fileerr = os.Create(dataName)
	}
	if fileerr != nil {
		fmt.Println(fmt.Sprintf("createLibrary  Error : %v ", fileerr))
		return
	}
	return
}

//ItemData 节点数据
type ItemData struct {
	IP         string
	StartIndex uint64
	EndIndex   uint64
}

//BuildRouteBytes 分解出route的bytes
func (d *ItemData) BuildRouteBytes() []byte {
	nodeDataBf := bytes.NewBuffer(nil)
	nodeEncode := gob.NewEncoder(nodeDataBf)
	nodeEncode.Encode(d)
	return nodeDataBf.Bytes()
}

//UnBuildRouteBytes UnBuildRouteBytes
func (d *ItemData) UnBuildRouteBytes(oribytes *[]byte) {
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
	routeDataFile.Write(resultBytes)
}

//BuildItemDataBytes 分解出itemData的bytes
func (d *ItemData) BuildItemDataBytes() {

}

// //SaveRouteBytes 分解出route的bytes
// func SaveRouteBytes(databytes []byte) {
// 	dataFile.Write(databytes)
// }
