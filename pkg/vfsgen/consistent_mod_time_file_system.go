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

package vfsgen

import (
	"net/http"
	"os"
	"time"
)

// consistentModTimeFileSystem is a wrapper for http.FileSystem. The value returned by the ModTime() function is the
// 'empty' time value. In doing so, it ensures that the output of vfsgen.Generate(...) is consistent with regards to
// the time at which Generate(...) is run.
type ConsistentModTimeFileSystem struct {
	HTTPFS http.FileSystem
}

type ConsistentModTimeFile struct {
	http.File
}

type ConsistentModTimeFileInfo struct {
	os.FileInfo
}

func (fs ConsistentModTimeFileSystem) Open(name string) (http.File, error) {
	file, err := fs.HTTPFS.Open(name)
	if err != nil {
		return nil, err
	}
	return ConsistentModTimeFile{File: file}, nil
}

func (f ConsistentModTimeFile) Stat() (os.FileInfo, error) {
	fileInfo, err := f.File.Stat()
	if err != nil {
		return nil, err
	}
	return ConsistentModTimeFileInfo{FileInfo: fileInfo}, nil
}

func (f ConsistentModTimeFileInfo) ModTime() time.Time {
	return time.Time{}
}
