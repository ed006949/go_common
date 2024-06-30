package io_fs

import (
	"io/fs"
)

type IO_FS_Content struct {
	Entry   *fs.DirEntry
	Content []byte
}
