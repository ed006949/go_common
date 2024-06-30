package io_fs

import (
	"io/fs"
)

type List_type map[string]string
type Content_type map[string]*IO_FS_type

type IO_FS_type struct {
	Entry   *fs.DirEntry
	Content []byte
}
