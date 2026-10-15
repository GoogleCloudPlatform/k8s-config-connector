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

package identity_test

import (
	"fmt"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
)

func TestAssertConformance(t *testing.T) {
	identity.AssertConformance(t, identity.ConformanceCase[fakeWidgetIdentity]{
		SampleExternal: "projects/my-project/locations/us-central1/widgets/my-widget",
		WantHost:       "widget.googleapis.com",
		WantParent:     "projects/my-project/locations/us-central1",
		WantIdentity: &fakeWidgetIdentity{
			Project:  "my-project",
			Location: "us-central1",
			Widget:   "my-widget",
		},
		ExpectServerGenerated: true,
		Ref:                   &fakeWidgetRef{},
	})
}

var fakeWidgetFormat = gcpurls.Template[fakeWidgetIdentity](
	"widget.googleapis.com",
	"projects/{project}/locations/{location}/widgets/{widget}",
)

type fakeWidgetIdentity struct {
	Project  string
	Location string
	Widget   string
}

func (i *fakeWidgetIdentity) String() string {
	return fakeWidgetFormat.ToString(*i)
}

func (i *fakeWidgetIdentity) FromExternal(ref string) error {
	parsed, match, err := fakeWidgetFormat.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("no match for %q", ref)
	}
	*i = *parsed
	return nil
}

func (i *fakeWidgetIdentity) Host() string {
	return fakeWidgetFormat.Host()
}

func (i *fakeWidgetIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func (i *fakeWidgetIdentity) HasIdentitySpecified() bool {
	return i != nil && i.Widget != ""
}

type fakeWidgetRef struct {
	External string
}

func (r *fakeWidgetRef) SetExternal(ref string) {
	r.External = ref
}

func (r *fakeWidgetRef) ValidateExternal(ref string) error {
	id := &fakeWidgetIdentity{}
	return id.FromExternal(ref)
}

func (r *fakeWidgetRef) ParseExternalToIdentity() (identity.Identity, error) {
	id := &fakeWidgetIdentity{}
	if err := id.FromExternal(r.External); err != nil {
		return nil, err
	}
	return id, nil
}
