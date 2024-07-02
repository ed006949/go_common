package io_fs

import (
	"io/fs"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func (receiver Content) Read() {
	var (
		err error
	)
	for _, b := range receiver {
		b.Abs = Abs(b.Abs)
		var (
			walkDirFunc = func(name string, dirEntry fs.DirEntry, err error) error {
				Fatalf("filepath.WalkDirFunc error: %v.", err)

				// switch {
				// case err != nil:
				// 	log.Fatalf("filepath.WalkDirFunc error: %v. ACTION: ERROR.", err)
				// }

				switch dirEntry.Type() {
				case fs.ModeSymlink:
					b.Entries[name] = &Entry{
						DirEntry:  dirEntry,
						Content:   ReadFile(name),
						Target:    ReadLink(name),
						IsChanged: false,
					}
				case 0:
					b.Entries[name] = &Entry{
						DirEntry:  dirEntry,
						Content:   ReadFile(name),
						Target:    "",
						IsChanged: false,
					}
				default:
				}

				return nil
			}
		)

		err = filepath.WalkDir(b.Abs, walkDirFunc)
		Fatalf("filepath.WalkDir error: %v.", err)

		// switch err = filepath.WalkDir(b.Abs, walkDirFunc); {
		// case err != nil:
		// 	log.Fatalf("filepath.WalkDir error: %v. ACTION: ERROR.", err)
		// }
	}
}

// func (receiver Content) IOFileWrite() {
// 	var (
// 		err error
// 	)
// 	for _, b := range receiver {
// 		switch _, err = os.Stat(b.Abs); {
// 		case err != nil:
// 			log.Fatalf("os.Stat error: %v. ACTION: ERROR.", err)
// 		case os.IsNotExist(err):
// 			switch f := os.MkdirAll(b.Abs, 0700); {
// 			case f != nil:
// 				log.Fatalf("os.MkdirAll error: %v. ACTION: ERROR.", f)
// 			}
// 		}
// 	}
// 	return
// }

func (receiver Content) WriteTemp() {
	var (
		// dirTemp, err = os.MkdirTemp(Abs("./tmp/"), "")
		dirTemp = Abs("./tmp/")
	)

	// switch {
	// case err != nil:
	// 	log.Fatalf("os.MkdirTemp error: %v. ACTION: ERROR.", err)
	// }

	// var (
	// 	dirTemp string
	// 	err     error
	// )
	// switch dirTemp, err = os.MkdirTemp("", ""); {
	// switch dirTemp, err = os.MkdirTemp(Abs("./tmp/"), ""); {
	// case err != nil:
	// 	log.Fatalf("os.MkdirTemp error: %v. ACTION: ERROR.", err)
	// }

	for _, b := range receiver {
		for c, d := range b.Entries {
			switch {
			case d.DirEntry.IsDir():
				continue
			}

			var (
				path = filepath.Join(dirTemp, filepath.Dir(c))
				name = filepath.Join(path, d.DirEntry.Name())
			)

			switch {
			case IsNotExist(path):
				MkdirAll(path, 0755)
			}

			switch d.DirEntry.Type() {
			case fs.ModeDir:
				continue
			case fs.ModeSymlink:
				Symlink(d.Target, name)
			case 0:
			}

			switch {
			case d.IsChanged:
				WriteFile(name, d.Content, 0644)
			}

		}
	}

	log.Infof("%s", dirTemp)
}
