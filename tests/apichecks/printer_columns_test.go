// Copyright 2026 Google LLC
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

package lint

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	"github.com/google/go-cmp/cmp"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestAdditionalPrinterColumns(t *testing.T) {
	crds, err := crdloader.LoadAllCRDs()
	if err != nil {
		t.Fatalf("error loading crds: %v", err)
	}

	wantColumns := []apiextensions.CustomResourceColumnDefinition{
		{
			Name:     "Age",
			Type:     "date",
			JSONPath: ".metadata.creationTimestamp",
		},
		{
			Name:        "Ready",
			Type:        "string",
			Description: "When 'True', the most recent reconcile of the resource succeeded",
			JSONPath:    ".status.conditions[?(@.type=='Ready')].status",
		},
		{
			Name:        "Status",
			Type:        "string",
			Description: "The reason for the value in 'Ready'",
			JSONPath:    ".status.conditions[?(@.type=='Ready')].reason",
		},
		{
			Name:        "Status Age",
			Type:        "date",
			Description: "The last transition time for the value in 'Status'",
			JSONPath:    ".status.conditions[?(@.type=='Ready')].lastTransitionTime",
		},
	}

	var errs []string
	for _, crd := range crds {
		for _, version := range crd.Spec.Versions {
			if diff := cmp.Diff(wantColumns, version.AdditionalPrinterColumns); diff != "" {
				errs = append(errs, fmt.Sprintf("crd=%s version=%v: additionalPrinterColumns mismatch:\n%s", crd.Name, version.Name, diff))
			}
		}
	}

	sort.Strings(errs)

	want := strings.Join(errs, "---\n")
	if want != "" {
		want += "\n"
	}

	test.CompareGoldenFile(t, "testdata/exceptions/printercolumns.txt", want)
}
