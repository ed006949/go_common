package io_fs

import (
	"io/fs"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func Initialize() {
	IO_file_read()
}

func IO_file_read() {
	for a, b := range List {
		switch c, d := filepath.Abs(b); {
		case d != nil:
			log.Fatalf("filepath.Abs error: %v. ACTION: ERROR.", d)
		default:
			List[a] = c
			switch f := filepath.WalkDir(c, walk); {
			case f != nil:
				log.Fatalf("filepath.WalkDir error: %v. ACTION: ERROR.", f)
			}
		}
	}
	// for _, b := range List {
	// 	switch d := filepath.WalkDir(b, walk); {
	// 	case d != nil:
	// 		log.Fatalf("filepath.WalkDir error: %v. ACTION: ERROR.", d)
	// 	}
	// }
}
func IO_file_write() {}

func walk(name string, entry fs.DirEntry, err error) error {
	switch {
	case err != nil:
		log.Fatalf("filepath.WalkDirFunc error: %v. ACTION: ERROR.", err)
	}

	Content[name] = &IO_FS_Content{
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
