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

package v1alpha1

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NotebookInstanceV2Identity{}
	_ identity.Resource   = &NotebookInstanceV2{}
)

var NotebookInstanceV2IdentityFormat = gcpurls.Template[NotebookInstanceV2Identity]("notebooks.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// NotebookInstanceV2Identity is the identity of a GCP NotebookInstanceV2 resource.
// +k8s:deepcopy-gen=false
type NotebookInstanceV2Identity struct {
	Project  string
	Location string
	Instance string
}

func (i *NotebookInstanceV2Identity) String() string {
	return NotebookInstanceV2IdentityFormat.ToString(*i)
}

func (i *NotebookInstanceV2Identity) FromExternal(ref string) error {
	parsed, match, err := NotebookInstanceV2IdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NotebookInstanceV2 external=%q was not known (use %s): %w", ref, NotebookInstanceV2IdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NotebookInstanceV2 external=%q was not known (use %s)", ref, NotebookInstanceV2IdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NotebookInstanceV2Identity) Host() string {
	return NotebookInstanceV2IdentityFormat.Host()
}

func (i *NotebookInstanceV2Identity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromNotebookInstanceV2Spec(ctx context.Context, reader client.Reader, obj *NotebookInstanceV2) (*NotebookInstanceV2Identity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == nil || *location == "" {
		return nil, fmt.Errorf("cannot resolve location (spec.location is empty or nil)")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NotebookInstanceV2Identity{
		Project:  projectID,
		Location: *location,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *NotebookInstanceV2) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNotebookInstanceV2Spec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NotebookInstanceV2Identity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NotebookInstanceV2 identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
