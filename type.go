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
	// Symlink   *SymlinkEntry
	Target    string
	IsChanged bool
}

// type SymlinkEntry struct {
// 	Target string
// 	Abs    string
// }

// type FSEntry map[string]*FS
//
// type FS struct {
// 	// meta cache
// 	// TO DO replace using methods?
// 	Abs string
//
// 	// sys
// 	DirEntry fs.DirEntry
//
// 	// file
// 	Content []byte
//
// 	// symlink
// 	Target    string
// 	TargetAbs string
//
// 	// nested
// 	FSEntry
// } // TO DO process scope vfs
//
// type FSDir struct {
// 	DirEntry   fs.DirEntry
// 	Upstream   *FSDir
// 	Downstream map[string]*interface{}
// }
// type FSFile struct {
// 	DirEntry  fs.DirEntry
// 	Content   []byte
// 	IsChanged bool
// }
// type FSSymlink struct {
// 	DirEntry  fs.DirEntry
// 	Target    string
// 	TargetAbs string
// }
