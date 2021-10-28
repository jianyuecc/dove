package store

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Separation byte = '♠' //用来分隔一列上的多个数据
	Null = '♦' // 表示空格
)


type DoveWriter struct {
	File *os.File
	Writer *bufio.Writer
	Offset int
}

// 将p写入文件,并重置文件的偏移指针
func (d *DoveWriter) Write(p []byte) error {

	nn, err := d.Writer.Write(p)
	if err != nil {
		return err
	}
	d.Offset += nn
	d.File.Seek(int64(nn), 0)
	d.Writer.Flush()
	return nil
}

// offset 文件的偏移量
// size 读取的字节数
func (d *DoveReader) Read(offset int64, size int64) ([]byte, error) {
	res := make([]byte, size)
	_, err := d.File.ReadAt(res, offset)
	return res, err
}

type DoveReader struct {
	File *os.File
	Offset int
}

func (d DoveWriter) Open(name string) error {
	var err error
	d.File, err = os.OpenFile(name, os.O_CREATE | os.O_RDWR, 0600)
	d.Writer = bufio.NewWriter(d.File)

	return err
}

type Footer struct {

}

type MetaData struct {
	uint32

	Schema map[string]TYPE
	len uint16
}






type RowGroup struct {
	len uint16
	table string
	Max int64
	min int64
	size uint32
	offset int64
}

type Column struct {
	len uint16
	TYPE
	Max int64
	min int64
	offset int64 //column chunk 开始位置
	size uint32	// column chunk 大小
}

type Page struct {
	len uint16
	check uint32
	value []byte // n字节
}

var len uint64 = 0 //记录Footer开始的位置

type TYPE uint8

const (
	INTEGER TYPE = 0
	STRING	TYPE = 1
	BOOL	TYPE = 2
	FLOAT   TYPE = 3
)

func  name()  {
	var a uint16 = 2
	var b uint32 = 4
	fmt.Println(uint32(a) + b)
}