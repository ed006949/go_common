package io_fs

import (
	"io/fs"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func (receiver Content) IOFileRead() {
	for _, b := range receiver {
		switch c, d := filepath.Abs(b.Abs); {
		case d != nil:
			log.Fatalf("filepath.Abs error: %v. ACTION: ERROR.", d)
		default:
			b.Abs = c
			switch f := filepath.WalkDir(c, func(name string, dirEntry fs.DirEntry, err error) error {
				switch {
				case err != nil:
					log.Fatalf("filepath.WalkDirFunc error: %v. ACTION: ERROR.", err)
				}

				b.Entries[name] = &Entry{
					DirEntry: dirEntry,
					Content: func() []byte {
						switch {
						case !dirEntry.IsDir():
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
			}); {
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

func (receiver Content) IOFileWrite() {
	return
}
