package entry

type Meta struct {
	MetaData
	DataPages []DataPage
	size uint64
	Footer uint64
}

type MetaData struct {
	Schema []SchemaElement // n bytes
	RowGroups []RowGroup // n bytes
	size uint32 //4 bytes
}


type SchemaElement struct {
	size uint32 // 4 bytes 自己的大小
	name string // n bytes
	Type TYPE //1 bytes

}



type RowGroup struct {
	size uint32	// 4 bytes 自己的大小
	Columns []ColumnChunk // n bytes
	NumRows uint64 // 8 bytes
	min uint64
	max uint64
}
type ColumnChunk struct {
	size uint32 //自己的大小
	name string
	Type TYPE
	DataPageOffset uint64 // 8 bytes, DataPage的位置
	min uint64
	max uint64
}

type DataPage struct {
	offset uint64
	min uint64
	max uint64
}


const (
	PageSize = 8 * 1024
)

type TYPE uint8

const (
	INTEGER TYPE = 0
	STRING	TYPE = 1
	BOOL	TYPE = 2
	FLOAT   TYPE = 3
)



