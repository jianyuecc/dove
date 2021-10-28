package meta

import (
	"time"
)


type Database struct {
	Tables
	TimeToLive int64
	Replication int8

}

type Tables []Table
type Levels []Level
type Fields []Field
type Table struct {
	Timestamp
	Levels
	Fields
}

type Timestamp time.Timer

type Field struct {
	DataType
}

type Level struct {
	DataType
}

type DataType byte

const (
	INT	   DataType = 0
	FLOAT  DataType = 1
	STRING DataType = 2
	BOOL   DataType = 3
)