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

package apikeyskey

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"cnrm.googlesource.com/cnrm/pkg/clients/generated/apis/apikeys/v1alpha1"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

func TestBuildCreate(t *testing.T) {
	ctx := context.TODO()

	b := `
displayName: "Human readable name"
projectRef:
  external: "projects/example-project"
resourceID: sample
restrictions:
  apiTargets:
  - service: "translate.googleapis.com"
    methods: [ "GET" ]
`

	u := &v1alpha1.APIKeysKeySpec{}
	if err := yaml.Unmarshal([]byte(b), u); err != nil {
		t.Fatalf("error parsing object: %v", err)
	}
	t.Logf("u: %+v", u)

	a := &adapter{}
	a.desired = u
	req, err := a.BuildCreate(ctx)
	if err != nil {
		t.Fatalf("error building create: %v", err)
	}

	j, err := protojson.Marshal(req)
	if err != nil {
		t.Fatalf("error converting proto to json: %v", err)
	}

	out := &unstructured.Unstructured{Object: make(map[string]interface{})}
	if err := json.Unmarshal(j, &out.Object); err != nil {
		t.Fatalf("error parsing json: %v", err)
	}

	outYAML, err := yaml.Marshal(out)
	if err != nil {
		t.Fatalf("error marshaling yaml: %v", err)
	}

	got := string(outYAML)

	want := `
key:
  displayName: Human readable name
  restrictions:
    apiTargets:
    - methods:
      - GET
      service: translate.googleapis.com
`

	got = strings.TrimSpace(got)
	want = strings.TrimSpace(want)

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("unexpected diff: %v", diff)
	}
}
