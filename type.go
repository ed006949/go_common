package io_fs

import (
	"io/fs"
)

type Content map[string]*Folder

type EntryList map[string]*Entry

type Folder struct {
	Abs     string
	Entries EntryList
}

type Entry struct {
	DirEntry fs.DirEntry
	Content  []byte
}
