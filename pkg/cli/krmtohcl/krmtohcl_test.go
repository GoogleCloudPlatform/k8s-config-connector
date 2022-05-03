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

package krmtohcl_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/krmtohcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testcmp "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/cmp"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	testyaml "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/yaml"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var update = flag.Bool("update", false, "update .golden files")

// when adding a new test case or when a resource schema changes run this test with '-update' parameter to update the
// 'golden files'. The HCL output is not deterministic so running an update will almost modify all the HCL files.
func TestUnstructuredToHCL(t *testing.T) {
	smLoader := testservicemappingloader.New(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.DefaultConfig)
	testCases := []struct {
		Name          string
		KRMFile       string
		GoldenHCLFile string
	}{
		{
			Name:          "ComputeInstance",
			KRMFile:       "computeinstance.yaml",
			GoldenHCLFile: "computeinstance.golden.tf",
		},
		{
			Name:          "ContainerCluster",
			KRMFile:       "containercluster.yaml",
			GoldenHCLFile: "containercluster.golden.tf",
		},
		{
			Name:          "PubSubSubscription",
			KRMFile:       "pubsubsubscription.yaml",
			GoldenHCLFile: "pubsubsubscription.golden.tf",
		},
		{
			Name:          "StorageBucket",
			KRMFile:       "storagebucket.yaml",
			GoldenHCLFile: "storagebucket.golden.tf",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			krmToHCL(t, tc.KRMFile, tc.GoldenHCLFile, smLoader, tfProvider)
		})
	}
}

func krmToHCL(t *testing.T, krmFile, hclFile string, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider) {
	var u unstructured.Unstructured
	testyaml.UnmarshalFile(t, fmt.Sprintf("testdata/%v", krmFile), &u)
	// the managed-by-cnrm is removed to test Terraform export for resources which are not managed by KCC.
	labels := u.GetLabels()
	delete(labels, "managed-by-cnrm")
	u.SetLabels(labels)
	hcl, err := krmtohcl.UnstructuredToHCL(&u, smLoader, tfProvider)
	if err != nil {
		t.Fatalf("error converting unstructured to HCL: %v", err)
	}
	goldenHCLFile := fmt.Sprintf("testdata/%v", hclFile)
	if *update {
		if err := ioutil.WriteFile(goldenHCLFile, []byte(hcl), 0644); err != nil {
			t.Fatalf("error writing file '%v': %v", goldenHCLFile, err)
		}
	}
	bytes, err := ioutil.ReadFile(goldenHCLFile)
	if err != nil {
		t.Fatalf("error reading file '%v': %v", goldenHCLFile, err)
	}
	goldenHCL := test.TrimLicenseHeaderFromTF(string(bytes))
	// HCL output is not stable so we do a line by line comparison
	testcmp.UnorderedLineByLineComparisonIgnoreBlankLines(t, goldenHCL, hcl)
}
