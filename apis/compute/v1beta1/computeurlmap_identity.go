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

package v1beta1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &ComputeURLMapIdentity{}
	_ identity.Resource   = &ComputeURLMap{}
)

var ComputeGlobalURLMapIdentityFormat = gcpurls.Template[ComputeURLMapIdentity]("compute.googleapis.com", "projects/{project}/global/urlMaps/{urlmap}")
var ComputeRegionalURLMapIdentityFormat = gcpurls.Template[ComputeURLMapIdentity]("compute.googleapis.com", "projects/{project}/regions/{region}/urlMaps/{urlmap}")

// ComputeURLMapIdentity is the identity of a GCP ComputeURLMap resource.
// +k8s:deepcopy-gen=false
type ComputeURLMapIdentity struct {
	Project string
	Region  string
	Urlmap  string
}

func (i *ComputeURLMapIdentity) IsGlobal() bool {
	return i.Region == "" || i.Region == "global"
}

func (i *ComputeURLMapIdentity) String() string {
	if !i.IsGlobal() {
		return ComputeRegionalURLMapIdentityFormat.ToString(*i)
	}
	return ComputeGlobalURLMapIdentityFormat.ToString(*i)
}

func (i *ComputeURLMapIdentity) FromExternal(ref string) error {
	ref = apirefs.TrimComputeURIPrefix(ref)

	if parsed, match, _ := ComputeGlobalURLMapIdentityFormat.Parse(ref); match {
		*i = *parsed
		i.Region = "global"
		return nil
	}
	if parsed, match, _ := ComputeRegionalURLMapIdentityFormat.Parse(ref); match {
		*i = *parsed
		return nil
	}
	return fmt.Errorf("format of ComputeURLMap external=%q was not known (use %s or %s)", ref, ComputeGlobalURLMapIdentityFormat.CanonicalForm(), ComputeRegionalURLMapIdentityFormat.CanonicalForm())
}

func (i *ComputeURLMapIdentity) Host() string {
	return ComputeGlobalURLMapIdentityFormat.Host()
}

func (i *ComputeURLMapIdentity) ParentString() string {
	if !i.IsGlobal() {
		return fmt.Sprintf("projects/%s/regions/%s", i.Project, i.Region)
	}
	return fmt.Sprintf("projects/%s/global", i.Project)
}

func ParseComputeURLMapExternal(external string) (*ComputeURLMapIdentity, error) {
	if external == "" {
		return nil, fmt.Errorf("empty ComputeURLMap external value")
	}
	id := &ComputeURLMapIdentity{}
	if err := id.FromExternal(external); err != nil {
		return nil, err
	}
	return id, nil
}

func getIdentityFromComputeURLMapSpec(ctx context.Context, reader client.Reader, obj *ComputeURLMap) (*ComputeURLMapIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	location := "global"
	if obj.Spec.Location != "" {
		location = obj.Spec.Location
	}

	identity := &ComputeURLMapIdentity{
		Project: projectID,
		Region:  location,
		Urlmap:  resourceID,
	}
	return identity, nil
}

func (obj *ComputeURLMap) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromComputeURLMapSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.selfLink, if present.
	selfLink := common.ValueOf(obj.Status.SelfLink)
	if selfLink != "" {
		statusIdentity := &ComputeURLMapIdentity{}
		if err := statusIdentity.FromExternal(selfLink); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change ComputeURLMap identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
