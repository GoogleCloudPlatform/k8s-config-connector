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

package iamresource_test

import (
	"flag"
	"fmt"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/cmd/bulkexport/iamresource"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
)

// add this flag to your test if you have a golden file which may sometimes need updating
var update = flag.Bool("update", false, "update .golden files")

func TestSplitPolicy(t *testing.T) {
	testSplitPolicy(t, "pubsub-topic-iampolicy.yaml")
	testSplitPolicy(t, "project-external-iampolicy.yaml")
}

func testSplitPolicy(t *testing.T, policyFilePath string) {
	iamPolicy := readIAMPolicyFromFile(t, policyFilePath)
	policyMembers := iamresource.SplitPolicy(iamPolicy)
	testPolicyFilePath := fmt.Sprintf("testdata/%v", policyFilePath)
	goldenFile := strings.Replace(testPolicyFilePath, ".yaml", ".golden.yaml", -1)
	if *update {
		testyaml.WriteValueToFile(t, policyMembers, goldenFile)
	}
	testyaml.AssertFileContentsMatchValue(t, goldenFile, policyMembers)
}

func TestConvertIAMPolicyToIAMPartialPolicy(t *testing.T) {
	testConvertIAMPolicyToIAMPartialPolicy(t, "pubsub-topic-iampolicy.yaml")
	testConvertIAMPolicyToIAMPartialPolicy(t, "project-external-iampolicy.yaml")
}

func testConvertIAMPolicyToIAMPartialPolicy(t *testing.T, policyFilePath string) {
	iamPolicy := readIAMPolicyFromFile(t, policyFilePath)
	partialPolicy := iamresource.ConvertIAMPolicyToIAMPartialPolicy(iamPolicy)
	testPolicyFilePath := fmt.Sprintf("testdata/%v", policyFilePath)
	goldenFile := strings.Replace(testPolicyFilePath, ".yaml", ".golden.partialpolicy.yaml", -1)
	testyaml.AssertFileContentsMatchValue(t, goldenFile, partialPolicy)
}

func readIAMPolicyFromFile(t *testing.T, policyFilePath string) *v1beta1.IAMPolicy {
	var iamPolicy v1beta1.IAMPolicy
	testPolicyFilePath := fmt.Sprintf("testdata/%v", policyFilePath)
	testyaml.UnmarshalFile(t, testPolicyFilePath, &iamPolicy)
	return &iamPolicy
}
