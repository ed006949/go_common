package io_fs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func Abs(path string) string {
	switch outbound, err := filepath.Abs(path); {
	case err != nil:
		log.Fatalf("filepath.Abs error: %v. ACTION: ERROR.", err)
		return ""
	default:
		return outbound
	}
}

func ReadLink(name string) string {
	switch outbound, err := os.Readlink(name); {
	case err != nil:
		log.Fatalf("os.Readlink error: %v. ACTION: ERROR.", err)
		return ""
	default:
		return outbound
	}
}

func ReadFile(name string) []byte {
	switch outbound, err := os.ReadFile(name); {
	case err != nil:
		log.Fatalf("os.ReadFile error: %v. ACTION: ERROR.", err)
		return nil
	default:
		return outbound
	}
}

func WriteFile(name string, data []byte, perm fs.FileMode) {
	switch err := os.WriteFile(name, data, perm); {
	case err != nil:
		log.Fatalf("os.WriteFile error: %v. ACTION: ERROR.", err)
		return
	default:
		return
	}
}

func MkdirAll(path string, perm fs.FileMode) {
	switch err := os.MkdirAll(path, perm); {
	case err != nil:
		log.Fatalf("os.MkdirAll error: %v. ACTION: ERROR.", err)
		return
	default:
		return
	}
}

func IsExist(path string) bool {
	return !IsNotExist(path)
}

func IsNotExist(path string) bool {
	switch _, err := os.Stat(path); {
	case errors.Is(err, fs.ErrNotExist):
		return true
	case err != nil:
		log.Fatalf("os.Stat error: %v. ACTION: ERROR.", err)
		return false
	default:
		return false
	}
}

func IsSymlink(path string) bool {
	switch stat, err := os.Lstat(path); {
	case err != nil:
		log.Fatalf("os.Stat error: %v. ACTION: ERROR.", err)
		return false
	default:
		return stat.Mode().Type() == fs.ModeSymlink
	}
}

func Symlink(oldname string, newname string) {
	switch err := os.Symlink(oldname, newname); {
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
	case err != nil:
		log.Fatalf("os.Symlink error: %v. ACTION: ERROR.", err)
		return
	default:
		return
	}
}
