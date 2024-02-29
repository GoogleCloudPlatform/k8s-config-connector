// Copyright 2024 Google LLC
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

package apikeys

import (
	"encoding/json"
	"strings"
	"testing"

	pb "cloud.google.com/go/apikeys/apiv2/apikeyspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apikeys/v1alpha1"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
	"sigs.k8s.io/yaml"
)

func TestBuildCreate(t *testing.T) {
	b := `
apiVersion: apikeys.cnrm.cloud.google.com/v1alpha1
kind: APIKeysKey
spec:
  displayName: "Human readable name"
  projectRef:
    external: "projects/example-project"
  resourceID: sample
  restrictions:
    serverKeyRestrictions:
      allowedIps:
      - 1.2.3.4
      - 5.6.7.8
`

	u := &v1alpha1.APIKeysKey{}
	if err := yaml.Unmarshal([]byte(b), u); err != nil {
		t.Fatalf("error parsing object: %v", err)
	}
	t.Logf("u: %+v", u)

	a := &adapter{}
	a.desired = u
	a.projectID = "example-project"
	a.location = "global"
	req, err := a.buildCreateRequest()
	if err != nil {
		t.Fatalf("error building create: %v", err)
	}

	j, err := protojson.Marshal(req)
	if err != nil {
		t.Fatalf("error converting proto to json: %v", err)
	}

	out := make(map[string]any)
	if err := json.Unmarshal(j, &out); err != nil {
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
    serverKeyRestrictions:
      allowedIps:
      - 1.2.3.4
      - 5.6.7.8
parent: projects/example-project/locations/global
`

	got = strings.TrimSpace(got)
	want = strings.TrimSpace(want)

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("unexpected diff: %v", diff)
	}
}

func TestMapping(t *testing.T) {
	originalKRM := &v1alpha1.APIKeysKey{}
	originalKRM.Spec.Restrictions = &v1alpha1.KeyRestrictions{
		AndroidKeyRestrictions: &v1alpha1.KeyAndroidKeyRestrictions{
			AllowedApplications: []v1alpha1.KeyAllowedApplications{
				{
					Sha1Fingerprint: "sha1-fingerprint",
					PackageName:     "package-name",
				},
				{
					Sha1Fingerprint: "sha1-fingerprint-2",
					PackageName:     "package-name-2",
				},
			},
		},
	}

	gotProto := &pb.Key{}
	if err := keyMapping.Map(originalKRM, gotProto); err != nil {
		t.Fatalf("error mapping: %v", err)
	}

	wantProto := &pb.Key{
		Restrictions: &pb.Restrictions{
			ClientRestrictions: &pb.Restrictions_AndroidKeyRestrictions{
				AndroidKeyRestrictions: &pb.AndroidKeyRestrictions{
					AllowedApplications: []*pb.AndroidApplication{
						{
							Sha1Fingerprint: "sha1-fingerprint",
							PackageName:     "package-name",
						},
						{
							Sha1Fingerprint: "sha1-fingerprint-2",
							PackageName:     "package-name-2",
						},
					},
				},
			},
		},
	}
	if diff := cmp.Diff(gotProto, wantProto, protocmp.Transform()); diff != "" {
		t.Errorf("unexpected diff: %v", diff)
	}

	gotKRM := &v1alpha1.APIKeysKey{}
	if err := keyMapping.Map(gotProto, gotKRM); err != nil {
		t.Fatalf("error mapping: %v", err)
	}

	if diff := cmp.Diff(gotKRM, originalKRM); diff != "" {
		t.Errorf("unexpected diff: %v", diff)
	}
}
