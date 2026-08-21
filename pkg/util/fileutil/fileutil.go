// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// DirExists determines if the given path points to a directory that exists.
func DirExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// HasSubdirs determines if a given path contains at least one subdirectory.
func HasSubdirs(path string) (bool, error) {
	subdirs, err := SubdirsIn(path)
	if err != nil {
		return false, err
	}
	return len(subdirs) > 0, nil
}

// SubdirsIn gets the names of subdirectories found in a given path.
func SubdirsIn(path string) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, fmt.Errorf("error reading directory '%v': %w", path, err)
	}
	subdirNames := make([]string, 0)
	for _, fi := range fileInfos {
		if fi.IsDir() {
			subdirNames = append(subdirNames, fi.Name())
		}
	}
	return subdirNames, nil
}

// FileNamesWithSuffixInDir gets all the filenames in the directory at the
// given path which end with the given suffix.
func FileNamesWithSuffixInDir(path, suffix string) (names []string, err error) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, fmt.Errorf("error reading directory '%v': %w", path, err)
	}
	names = make([]string, 0)
	for _, fi := range fileInfos {
		if fi.IsDir() {
			continue
		}
		if strings.HasSuffix(fi.Name(), suffix) {
			names = append(names, fi.Name())
		}
	}
	return names, nil
}

// NewEmptyFile creates an empty file at the given path.
func NewEmptyFile(path string) (*os.File, error) {
	if err := ensureParentDirectoryExists(path); err != nil {
		return nil, fmt.Errorf("error ensuring parent directory of %v exists: %w", path, err)
	}
	return os.Create(path)
}

func ensureParentDirectoryExists(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("error creating directory %v and its parents: %w", dir, err)
	}
	return nil
}
