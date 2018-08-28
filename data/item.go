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
var startIndex int64 //存储在本节点的 开始index
var endIndex int64   //存储在本节点的 结束index
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
	Index uint64
	Data  []byte
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
	routeDataFile.Write(d.toBytes())
}
func (d *ItemData) toBytes() []byte {
	resultBytes := make([]byte, 50, 50)
	resultBF := bytes.NewBuffer(resultBytes)
	resultBF.Truncate(0)
	binary.Write(resultBF, binary.BigEndian, d.Index)
	resultBF.Write(d.Data)
	return resultBF.Bytes()
}
