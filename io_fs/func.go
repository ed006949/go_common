package io_fs

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/ed006949/io_fs/e"
)

func Abs(path string) string {
	var (
		outbound, err = filepath.Abs(path)
	)
	e.Fatalf("filepath.Abs error: %v.", err)
	return outbound
}

func ReadLink(name string) string {
	var (
		outbound, err = os.Readlink(name)
	)
	e.Fatalf("os.Readlink error: %v.", err)
	return outbound
}

func ReadFile(name string) []byte {
	var (
		outbound, err = os.ReadFile(name)
	)
	e.Fatalf("os.ReadFile error: %v.", err)
	return outbound
}

func WriteFile(name string, data []byte, perm fs.FileMode) {
	var (
		err = os.WriteFile(name, data, perm)
	)
	e.Fatalf("os.WriteFile error: %v.", err)
	return
}

func MkdirAll(path string, perm fs.FileMode) {
	var (
		err = os.MkdirAll(path, perm)
	)
	e.Fatalf("os.MkdirAll error: %v.", err)
	return
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
	e.Fatalf("os.Stat error: %v.", err)
	return false
}

func IsSymlink(path string) bool {
	var (
		stat, err = os.Lstat(path)
	)
	e.Fatalf("os.Lstat error: %v.", err)
	return stat.Mode().Type() == fs.ModeSymlink
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
	e.Fatalf("os.Symlink error: %v.", err)
	return
}
