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

package serialization

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"

	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"gopkg.in/yaml.v2"
)

func TestSerializeInstance(t *testing.T) {
	testSerialize(t, "testdata/instance.terraform.instancestate", "google_compute_instance", "testdata/instance.golden.tf")
}

func TestSerializeCluster(t *testing.T) {
	testSerialize(t, "testdata/cluster.terraform.instancestate", "google_container_cluster", "testdata/cluster.golden.tf")
}

func testSerialize(t *testing.T, instanceStateFile, tfType, goldenFile string) string {
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	b, err := ioutil.ReadFile(instanceStateFile)
	if err != nil {
		t.Fatalf("failed to load instance state for instance, %s", err.Error())
	}
	is := &terraform.InstanceState{}
	if err := yaml.Unmarshal(b, &is); err != nil {
		t.Fatal(err)
	}
	info := &terraform.InstanceInfo{
		Id:   "foo",
		Type: tfType,
	}
	out, err := InstanceStateToHCL(is, info, provider)
	if err != nil {
		t.Fatalf("failed to create HCL: %s", err.Error())
	}

	fmt.Println(out)
	if err != nil {
		t.Fatal(err)
	}
	golden, err := ioutil.ReadFile(goldenFile)
	if err != nil {
		t.Fatal(err)
	}
	goldenLines := strings.Split(string(golden), "\n")
	for linenum, line := range strings.Split(out, "\n") {
		if !stringContains(goldenLines, line) {
			t.Fatalf("golden value didn't match provided value: line %d, %q, not in golden.", linenum, line)
		}
	}
	return out
}

func stringContains(ss []string, s string) bool {
	re := regexp.MustCompile(`\s*=\s*`)
	for _, sss := range ss {
		if re.ReplaceAllString(sss, " = ") == re.ReplaceAllString(s, " = ") {
			return true
		}
	}
	return false
}
