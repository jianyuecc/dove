package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/acaee/dove/doved/example/meta"
	"github.com/golang/snappy"
	"log"
)

const FileName = "doved\\example\\meta\\meta.db"

func main() {
	meta := meta.Meta{FileName: FileName}
	meta.CreateDatabase("system")
	meta.CreateTable("system", "cpu")
	meta.CreateTable("system", "disk")
	meta.CreateTable("system", "mem")
	err, databases := meta.ShowDatabases()
	if err != nil {
		log.Fatalln(err)
	}

	for _, database := range databases {
		fmt.Println(database)
	}

	meta.Put([]byte("Duration"), []byte("30d"))
	meta.Put([]byte("Replication"), []byte("2"))
	meta.Put([]byte("Sduration"), []byte("1d"))


	meta.PutCPU([]byte("Duration"), []byte("15d"))
	meta.PutCPU([]byte("Replication"), []byte("1"))
	meta.PutCPU([]byte("Sduration"), []byte("0d"))

	_, databases = meta.ShowTables("system")
	for _, database := range databases {
		fmt.Println(database)
	}
	fmt.Println("-----------------------------")
	var b = []byte{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,}
	fmt.Println(b)
	fmt.Println(b[0:4])
	fmt.Println(b[4:8])
	fmt.Println("------------------binary----------")
	res := make([]byte,2)
	binary.BigEndian.PutUint16(res, 65535)
	res2 := make([]byte,4)
	binary.BigEndian.PutUint16(res2, 65535)
	fmt.Println(res)
	var buffer bytes.Buffer
	buffer.Write(res)
	buffer.Write(res2)
	re := snappy.Encode(nil, buffer.Bytes())
	fmt.Println(re)



}
