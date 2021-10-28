package store

import (
	"encoding/binary"
	"github.com/golang/snappy"
)

type Line struct {
	database string
	table string
	
}

func RangeTimestamp()  {
}
// insert system,cpu region=Shanghai,host=127.0.0.1 usage=56.454 idle=23656
// insert system,cpu region=Beijing, host=127.0.0.1 usage=56.251 idle=6623

func Receive()  {
	var meta Meta
	var rowGroup RowGroup
	var column Column
	meta.Schema = make(map[string]TYPE)
	meta.Schema["region"] = STRING
	meta.Schema["host"] = STRING
	meta.Schema["usage"] = FLOAT
	meta.Schema["idle"] = FLOAT

	binary.Write()
}

