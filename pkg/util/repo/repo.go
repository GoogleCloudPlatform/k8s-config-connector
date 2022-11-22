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

package repo

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/golang/glog"
)

func getGitRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting working directory: %w", err)
	}

	for {
		p := filepath.Join(dir, ".git")
		_, err := os.Stat(p)
		if err != nil {
			if !os.IsNotExist(err) {
				return "", fmt.Errorf("error getting stat of %q: %w", p, err)
			}
		}
		if err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("unable to locate repo root in working directory %q", dir)
		}
		dir = parent
	}
}

func GetRoot() (string, error) {
	return getGitRoot()
}

func GetRootOrLogFatal() string {
	root, err := GetRoot()
	if err != nil {
		glog.Fatal(err)
	}
	return root
}

func GetRootOrTestFatal(t *testing.T) string {
	t.Helper()
	root, err := GetRoot()
	if err != nil {
		t.Fatal(err)
	}
	return root
}

func GetCallerPackagePath() (string, error) {
	return getCallerPackagePath()
}

func GetCallerPackagePathOrLogFatal() string {
	path, err := getCallerPackagePath()
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func GetCallerPackagePathOrTestFatal(t *testing.T) string {
	t.Helper()
	path, err := getCallerPackagePath()
	if err != nil {
		t.Fatal(err)
	}
	return path
}

func getCallerPackagePath() (string, error) {
	_, filePath, _, ok := runtime.Caller(2)
	if !ok {
		return "", fmt.Errorf("unable to get runtimer caller")
	}
	return filepath.Dir(filePath), nil
}
