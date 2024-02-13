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

package testos

import (
	"os"
	"testing"
)

// returns a 'mock' stdin in the form of a file
func GetStdin(t *testing.T, content string) (func(), *os.File) {
	if content == "" {
		return func() {}, os.Stdin
	}
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatalf("error creating temporary file: %v", err)
	}
	// if anything else in this function fails, remove the file
	defer func(err *error) {
		if *err != nil {
			os.Remove(tmpFile.Name())
		}
	}(&err)
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("error writing to file: %v", err)
	}
	if _, err := tmpFile.Seek(0, 0); err != nil {
		t.Fatalf("error seeking in file: :%v", err)
	}
	cleanup := func() {
		os.Remove(tmpFile.Name())
	}
	return cleanup, tmpFile
}
