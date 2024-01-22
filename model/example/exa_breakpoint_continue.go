package example

import (
	"github.com/yaoyaochil/bodo-admin-server/server/global"
)

// file struct, 文件结构体
type ExaFile struct {
	global.BODO_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.BODO_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
