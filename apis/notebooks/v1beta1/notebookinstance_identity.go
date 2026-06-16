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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	_ identity.IdentityV2 = &NotebookInstanceIdentity{}
	_ identity.Resource   = &NotebookInstance{}
)

var NotebookInstanceIdentityFormat = gcpurls.Template[NotebookInstanceIdentity]("notebooks.googleapis.com", "projects/{project}/locations/{location}/instances/{instance}")

// NotebookInstanceIdentity is the identity of a Google Cloud NotebookInstance resource.
// +k8s:deepcopy-gen=false
type NotebookInstanceIdentity struct {
	Project  string
	Location string
	Instance string
}

func (i *NotebookInstanceIdentity) String() string {
	return NotebookInstanceIdentityFormat.ToString(*i)
}

func (i *NotebookInstanceIdentity) FromExternal(ref string) error {
	parsed, match, err := NotebookInstanceIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of NotebookInstance external=%q was not known (use %s): %w", ref, NotebookInstanceIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of NotebookInstance external=%q was not known (use %s)", ref, NotebookInstanceIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *NotebookInstanceIdentity) Host() string {
	return NotebookInstanceIdentityFormat.Host()
}

func (i *NotebookInstanceIdentity) ParentString() string {
	return "projects/" + i.Project + "/locations/" + i.Location
}

func getIdentityFromNotebookInstanceSpec(ctx context.Context, reader client.Reader, obj *NotebookInstance) (*NotebookInstanceIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Zone
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location (spec.zone is empty)")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &NotebookInstanceIdentity{
		Project:  projectID,
		Location: location,
		Instance: resourceID,
	}
	return identity, nil
}

func (obj *NotebookInstance) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromNotebookInstanceSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Cross-check the identity against status.externalRef, if present.
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &NotebookInstanceIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change NotebookInstance identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
