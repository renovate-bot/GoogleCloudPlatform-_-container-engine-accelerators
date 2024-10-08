// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"
	"regexp"

	"github.com/fsnotify/fsnotify"
)

func DeviceNameFromPath(path string) (string, error) {
	gpuPathRegex := regexp.MustCompile("/dev/(nvidia[0-9]+)$")
	m := gpuPathRegex.FindStringSubmatch(path)
	if len(m) != 2 {
		return "", fmt.Errorf("path (%s) is not a valid GPU device path", path)
	}
	return m[1], nil
}

// Files creates a Watcher for the specified files.
func Files(files ...string) (*fsnotify.Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		err = watcher.Add(f)
		if err != nil {
			watcher.Close()
			return nil, err
		}
	}
	return watcher, nil
}
