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
	"strings"

	"github.com/golang/glog"
)

const (
	operatorSrcRoot = "github.com/GoogleCloudPlatform/k8s-config-connector/operator"
)

func GetOperatorSrcRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting working directory: %v", err)
	}
	if idx := strings.Index(dir, operatorSrcRoot); idx != -1 {
		return dir[0 : idx+len(operatorSrcRoot)], nil
	}
	return "", fmt.Errorf("unable to locate operator source root '%v' in working directory '%v'",
		operatorSrcRoot, dir)
}

func GetOperatorCRDsPath() string {
	return filepath.Join(GetOperatorSrcRootOrLogFatal(), "config", "crd", "bases")
}

func GetOperatorSrcRootOrLogFatal() string {
	root, err := GetOperatorSrcRoot()
	if err != nil {
		glog.Fatal(err)
	}
	return root
}
