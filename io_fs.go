package io_fs

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// func Initialize() {
// }

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
