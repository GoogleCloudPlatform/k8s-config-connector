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
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"sigs.k8s.io/yaml"
)

func TestSerializeInstance(t *testing.T) {
	testSerialize(t, "testdata/instance.terraform.instancestate", "google_compute_instance", "testdata/instance.golden.tf")
}

func TestSerializeCluster(t *testing.T) {
	testSerialize(t, "testdata/cluster.terraform.instancestate", "google_container_cluster", "testdata/cluster.golden.tf")
}

func testSerialize(t *testing.T, instanceStateFile, tfType, goldenFile string) {
	provider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	b := test.MustReadFile(t, instanceStateFile)
	is := &terraform.InstanceState{}
	if err := yaml.Unmarshal(b, &is); err != nil {
		t.Fatalf("parsing yaml: %v", err)
	}
	info := &terraform.InstanceInfo{
		Id:   "foo",
		Type: tfType,
	}
	hcl, err := InstanceStateToHCL(is, info, provider)
	if err != nil {
		t.Fatalf("failed to create HCL: %v", err)
	}

	test.CompareGoldenFile(t, goldenFile, hcl, test.IgnoreLeadingComments)
}
