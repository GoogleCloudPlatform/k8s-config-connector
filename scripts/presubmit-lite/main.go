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

// This package gets & joins the intended presubmit-lite tests into a usable
// regex string for the scripts/presubmit-lite-test.sh

package main

import (
	"fmt"

	testconstants "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/constants"
)

func main() {
	testList := testconstants.GetPresubmitLiteRegexStringArray()
	testRegexString := testconstants.JoinTestNamesWithRegexFormat(testList)

	fmt.Println(testRegexString)
}
