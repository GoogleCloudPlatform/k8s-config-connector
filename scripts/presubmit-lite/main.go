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

package main

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/GoogleCloudPlatform/k8s-config-connector/operator/scripts/utils"
	testconstants "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/constants"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"
)

func main() {
	testList := testconstants.GetPresubmitLiteRegexStringArray()
	testRegexString := testconstants.JoinTestNamesWithRegexFormat(testList)

	repoRoot, _ := filepath.Abs(repo.GetRootOrLogFatal())
	freshEnvCmd := exec.Command("./scripts/run-tests-fresh-environment.sh", "--target-directory",
		testconstants.DynamicTestPackagePath, "--run-tests", "\""+testRegexString+"\"")
	cmd := exec.Command("./scripts/run-command-new-env.sh", "--command", freshEnvCmd.String())
	cmd.Dir = repoRoot
	if err := utils.Execute(cmd); err != nil {
		log.Fatal(err)
	}
}
