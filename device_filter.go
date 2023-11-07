// Copyright 2022 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

type deviceFilter struct {
	ignorePattern *regexp.Regexp
	acceptPattern *regexp.Regexp

	ignores []string
	accepts []string
}

func newDeviceFilter(ignoredPattern, acceptPattern string) (f deviceFilter) {
	if ignoredPattern != "" {
		f.ignores = readLinkInDevDir(ignoredPattern)
		if len(f.ignores) == 0 {
			f.ignorePattern = regexp.MustCompile(ignoredPattern)
		}
	}

	if acceptPattern != "" {
		f.accepts = readLinkInDevDir(ignoredPattern)
		if len(f.accepts) == 0 {
			f.acceptPattern = regexp.MustCompile(acceptPattern)
		}
	}

	return
}

func readLinkInDevDir(pattern string) (list []string) {
	if !strings.HasPrefix(pattern, "/dev/") {
		return
	}

	files, err := filepath.Glob("/dev/**/*")
	if err != nil {
		return
	}

	reg := regexp.MustCompile(pattern)

	for _, file := range files {
		if !reg.MatchString(file) {
			continue
		}
		fs, err := os.Stat(file)
		if err != nil || fs.Mode()&os.ModeSymlink == 0 {
			continue
		}
		file, err = os.Readlink(file)
		if err != nil {
			continue
		}
		list = append(list, file)
	}
	return
}

// ignored returns whether the device should be ignored
func (f *deviceFilter) ignored(name string) bool {
	return ((f.ignorePattern != nil && f.ignorePattern.MatchString(name)) ||
		(f.acceptPattern != nil && !f.acceptPattern.MatchString(name))) ||
		(f.ignores != nil && slices.Index(f.ignores, name) >= 0) ||
		(f.accepts != nil && slices.Index(f.accepts, name) >= 0)
}
