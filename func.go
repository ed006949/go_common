package io_fs

import (
	"io/fs"
	"os"

	log "github.com/sirupsen/logrus"
)

func walk(name string, entry fs.DirEntry, err error) error {
	switch {
	case err != nil:
		log.Fatalf("filepath.WalkDirFunc error: %v. ACTION: ERROR.", err)
	}

	Content[name] = &IO_FS_type{
		Entry: &entry,
		Content: func() []byte {
			switch {
			case !entry.IsDir():
				switch a, b := os.ReadFile(name); {
				case b != nil:
					log.Fatalf("os.ReadFile error: %v. ACTION: ERROR.", b)
				default:
					return a
				}
			}
			return nil
		}(),
	}

	return nil
}
