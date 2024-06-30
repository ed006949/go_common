package io_fs

import (
	"io/fs"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func IO_file_read(list List_type, content Content_type) {
	for a, b := range list {
		switch c, d := filepath.Abs(b); {
		case d != nil:
			log.Fatalf("filepath.Abs error: %v. ACTION: ERROR.", d)
		default:
			list[a] = c
			switch f := filepath.WalkDir(c, func(name string, entry fs.DirEntry, err error) error {
				switch {
				case err != nil:
					log.Fatalf("filepath.WalkDirFunc error: %v. ACTION: ERROR.", err)
				}

				content[name] = &IO_FS_type{
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

func IO_file_write() {}
