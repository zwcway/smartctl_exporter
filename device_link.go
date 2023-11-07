package main

import (
	"os"
	"path/filepath"
)

func readDevicesLink(list []string) (ds []string) {
	devices := []string{}
	for _, f := range list {
		_, err := os.Stat(f)
		if err != nil {
			if files, err := filepath.Glob(f); err == nil {
				devices = append(devices, files...)
			}
			continue
		}
		ds = append(ds, f)
	}
	for _, f := range devices {
		f, err := filepath.EvalSymlinks(f)
		if err != nil {
			continue
		}
		ds = append(ds, f)
	}

	return
}
