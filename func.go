package io_fs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func Fatalf(format string, err error, args ...interface{}) {
	switch {
	case err != nil:
		log.Fatalf(format+" ACTION: Exit(1). ", err, args)
	}
}

func Abs(path string) string {
	var (
		outbound, err = filepath.Abs(path)
	)
	Fatalf("filepath.Abs error: %v.", err)
	return outbound

	// switch outbound, err := filepath.Abs(path); {
	// case err != nil:
	// 	log.Fatalf("filepath.Abs error: %v.", err)
	// 	return ""
	// default:
	// 	return outbound
	// }
}

func ReadLink(name string) string {
	var (
		outbound, err = os.Readlink(name)
	)
	Fatalf("os.Readlink error: %v.", err)
	return outbound

	// switch outbound, err := os.Readlink(name); {
	// case err != nil:
	// 	log.Fatalf("os.Readlink error: %v.", err)
	// 	return ""
	// default:
	// 	return outbound
	// }
}

func ReadFile(name string) []byte {
	var (
		outbound, err = os.ReadFile(name)
	)
	Fatalf("os.ReadFile error: %v.", err)
	return outbound

	// switch outbound, err := os.ReadFile(name); {
	// case err != nil:
	// 	log.Fatalf("os.ReadFile error: %v.", err)
	// 	return nil
	// default:
	// 	return outbound
	// }
}

func WriteFile(name string, data []byte, perm fs.FileMode) {
	var (
		err = os.WriteFile(name, data, perm)
	)
	Fatalf("os.WriteFile error: %v.", err)
	return

	// switch err := os.WriteFile(name, data, perm); {
	// case err != nil:
	// 	log.Fatalf("os.WriteFile error: %v.", err)
	// 	return
	// default:
	// 	return
	// }
}

func MkdirAll(path string, perm fs.FileMode) {
	var (
		err = os.MkdirAll(path, perm)
	)
	Fatalf("os.MkdirAll error: %v.", err)
	return

	// switch err := os.MkdirAll(path, perm); {
	// case err != nil:
	// 	log.Fatalf("os.MkdirAll error: %v.", err)
	// 	return
	// default:
	// 	return
	// }
}

func IsExist(path string) bool {
	return !IsNotExist(path)
}

func IsNotExist(path string) bool {
	var (
		_, err = os.Stat(path)
	)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return true
	}
	Fatalf("os.Stat error: %v.", err)
	return false

	// switch _, err := os.Stat(path); {
	// case errors.Is(err, fs.ErrNotExist):
	// 	return true
	// case err != nil:
	// 	log.Fatalf("os.Stat error: %v.", err)
	// 	return false
	// default:
	// 	return false
	// }
}

func IsSymlink(path string) bool {
	var (
		stat, err = os.Lstat(path)
	)
	Fatalf("os.Lstat error: %v.", err)
	return stat.Mode().Type() == fs.ModeSymlink

	// switch stat, err := os.Lstat(path); {
	// case err != nil:
	// 	log.Fatalf("os.Stat error: %v.", err)
	// 	return false
	// default:
	// 	return stat.Mode().Type() == fs.ModeSymlink
	// }
}

func Symlink(oldname string, newname string) {
	var (
		err = os.Symlink(oldname, newname)
	)
	switch {
	case errors.Is(err, fs.ErrExist):
		var (
			interim *os.LinkError
			_       = errors.As(err, &interim)
		)
		switch {
		case IsSymlink(newname) && interim.Old == oldname && interim.New == newname:
			log.Warnf("os.Symlink warning: %v. ACTION: IGNORE.", err)
		}
		return
	}
	Fatalf("os.Symlink error: %v.", err)
	return

	// switch err := os.Symlink(oldname, newname); {
	// case errors.Is(err, fs.ErrExist):
	// 	var (
	// 		interim *os.LinkError
	// 		_       = errors.As(err, &interim)
	// 	)
	// 	switch {
	// 	case IsSymlink(newname) && interim.Old == oldname && interim.New == newname:
	// 		log.Warnf("os.Symlink warning: %v. ACTION: IGNORE.", err)
	// 	}
	// 	return
	// case err != nil:
	// 	log.Fatalf("os.Symlink error: %v.", err)
	// 	return
	// default:
	// 	return
	// }
}
