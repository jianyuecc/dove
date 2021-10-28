package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/snappy"
	"hash/crc32"
	"log"
	"os"
)

func main() {
	//store := store.Store{}
	//store.Open("doved/engine/store/data.dove")
	////store.Write(nil)
	//	//
	//	//file, err := os.OpenFile("doved/engine/store/data.dove", os.O_CREATE | os.O_RDWR, 0600)
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	//	//
	//	//write := bufio.NewWriter(file)
	//	//
	//	//
	//	//file.Seek(0,2)
	//	//write.Write([]byte{'A','B','C','D','E','2'})
	//	//write.Flush()

	testSnappy()


}


func testBinary() {
	//序列化
	dataA := []byte("san")

	var buffer bytes.Buffer
	err1 := binary.Write(&buffer, binary.BigEndian, &dataA)
	if err1!=nil{
		log.Panic(err1)
	}
	byteA:=buffer.Bytes()
	fmt.Println("序列化后：",byteA)

	//反序列化
	var dataB uint64
	var byteB []byte=byteA
	err2:=binary.Read(bytes.NewReader(byteB),binary.BigEndian,&dataB)
	if err2!=nil{
		log.Panic(err2)
	}
	fmt.Println("反序列化后：",byteB)
}


func testCrc32() {
	var src = []byte{'A', 'B', 'C', 'D', 'A', 'B', 'C', 'D','A', 'B', 'C', 'D', 'A', 'B', 'C', 'D','A', 'B', 'C', 'D', 'A', 'B', 'C', 'D'}
	check := crc32.ChecksumIEEE(src)
	fmt.Println(check)
	src = []byte{'A', 'B', 'C', 'D', 'A', 'B', 'C', 'D','A', 'B', 'C', 'D', 'A', 'B', 'C', 'D','A', 'B', 'C', 'D', 'A', 'B', 'C', 'D'}
	if crc32.ChecksumIEEE(src) == check {
		fmt.Println("数据无损!")
	}
}

func testSnappy() {
	var src = []byte{'A', 'B', 'C', 'D', 'A', 'B', 'C', 'D','A', 'B', 'C', 'D', 'A', 'B', 'C', 'D','A', 'B', 'C', 'D', 'A', 'B', 'C', 'D'}
	dst := snappy.Encode(nil, src)

	file, err := os.OpenFile("doved/engine/store/data.dove", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0600)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(len(dst))
	file.Write(dst)

	file.Close()


}

