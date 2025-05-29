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

package paths

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/klog/v2"
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

func GetOperatorSrcRoot() (string, error) {
	gitRoot, err := getGitRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(gitRoot, "operator"), nil
}

func GetOperatorCRDsPaths() []string {
	return []string{
		filepath.Join(GetOperatorSrcRootOrLogFatal(), "config", "crd", "base", "bases"),             // core kcc operator CRDs
		filepath.Join(GetOperatorSrcRootOrLogFatal(), "config", "crd", "overlays", "full", "bases"), // all customization CRDs
	}
}

func GetOperatorSrcRootOrLogFatal() string {
	root, err := GetOperatorSrcRoot()
	if err != nil {
		klog.Fatal(err)
	}
	return root
}
