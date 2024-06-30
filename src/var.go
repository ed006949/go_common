package io_fs

var Content = make(map[string]*IO_FS_Content)
var (
	List = map[string]string{
		"etc":  "./etc/",
		"tmpl": "./tmpl/",
		"var":  "./var/",
	}
)
