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
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/cli/krmtohcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testservicemappingloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemappingloader"
	tfprovider "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/tf/provider"
)

func TestUnstructuredToHCL(t *testing.T) {
	smLoader := testservicemappingloader.New(t)
	tfProvider := tfprovider.NewOrLogFatal(tfprovider.UnitTestConfig())
	testDir := "testdata"

	testCases := FindTestCases(t, testDir, ".golden.tf")
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			ctx := context.TODO()

			krmFile := testCase + ".yaml"
			goldenHCLFile := testCase + ".golden.tf"

			var u unstructured.Unstructured
			b := test.MustReadFile(t, krmFile)
			if err := yaml.Unmarshal(b, &u); err != nil {
				t.Fatalf("parsing yaml from %q: %v", krmFile, err)
			}

			// the managed-by-cnrm is removed to test Terraform export for resources which are not managed by KCC.
			labels := u.GetLabels()
			delete(labels, "managed-by-cnrm")
			u.SetLabels(labels)
			hcl, err := krmtohcl.UnstructuredToHCL(ctx, &u, smLoader, tfProvider)
			if err != nil {
				t.Fatalf("error converting unstructured to HCL: %v", err)
			}

			test.CompareGoldenFile(t, goldenHCLFile, hcl, test.IgnoreLeadingComments)
		})
	}
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
