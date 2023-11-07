package main

import (
	"os"
	"path/filepath"
)

func readDevicesLink(list []string) (ds []string) {
	devices := []string{}

	for _, f := range list {
		fs, err := os.Stat(f)
		if err != nil {
			if files, err := filepath.Glob(f); err == nil {
				devices = append(devices, files...)
			}
			continue
		}
		if fs.Mode()&os.ModeDevice != 0 {
			ds = append(ds, f)
		}
	}

	for _, f := range devices {
		fs, err := os.Stat(f)
		if err != nil {
			continue
		}

		if fs.Mode()&os.ModeSymlink != 0 {
			if f, err = os.Readlink(f); err != nil {
				continue
			}
			ds = append(ds, f)
		} else if fs.Mode()&os.ModeDevice != 0 {
			ds = append(ds, f)
		}
	}
	return
}
