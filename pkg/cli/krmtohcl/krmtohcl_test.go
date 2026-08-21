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
	"context"
	"flag"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
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
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	testDir := "testdata"

	testCases := FindTestCases(t, testDir, ".golden.tf")
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			krmFile := testCase + ".yaml"
			goldenHCLFile := testCase + ".golden.tf"

			testUnstructuredToHCL(t, krmFile, goldenHCLFile, smLoader, tfProvider)
		})
	}
}

func testUnstructuredToHCL(t *testing.T, krmFile, goldenHCLFile string, smLoader *servicemappingloader.ServiceMappingLoader, tfProvider *schema.Provider) {
	ctx := context.TODO()

	var u unstructured.Unstructured
	testyaml.UnmarshalFile(t, krmFile, &u)
	// the managed-by-cnrm is removed to test Terraform export for resources which are not managed by KCC.
	labels := u.GetLabels()
	delete(labels, "managed-by-cnrm")
	u.SetLabels(labels)
	hcl, err := krmtohcl.UnstructuredToHCL(ctx, &u, smLoader, tfProvider)
	if err != nil {
		t.Fatalf("error converting unstructured to HCL: %v", err)
	}
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

// FindTestCases returns all the test cases under basedir.
// It only returns ones which match the suffix, and strips the suffix.
func FindTestCases(t *testing.T, basedir string, suffix string) []string {
	var cases []string
	if err := filepath.Walk(basedir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		p := path
		if strings.HasSuffix(p, suffix) {
			cases = append(cases, strings.TrimSuffix(p, suffix))
		}
		return nil
	}); err != nil {
		t.Fatalf("error from filepath.Walk(%q): %v", basedir, err)
	}

	return cases
}
