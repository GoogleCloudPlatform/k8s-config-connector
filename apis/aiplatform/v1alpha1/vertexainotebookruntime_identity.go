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
	_ identity.IdentityV2 = &VertexAINotebookRuntimeIdentity{}
	_ identity.Resource   = &VertexAINotebookRuntime{}
)

var VertexAINotebookRuntimeIdentityFormat = gcpurls.Template[VertexAINotebookRuntimeIdentity]("aiplatform.googleapis.com", "projects/{project}/locations/{location}/notebookRuntimes/{notebookRuntime}")

// VertexAINotebookRuntimeIdentity is the identity of a GCP VertexAINotebookRuntime resource.
// +k8s:deepcopy-gen=false
type VertexAINotebookRuntimeIdentity struct {
	Project         string
	Location        string
	NotebookRuntime string
}

func (i *VertexAINotebookRuntimeIdentity) String() string {
	return VertexAINotebookRuntimeIdentityFormat.ToString(*i)
}

func (i *VertexAINotebookRuntimeIdentity) FromExternal(ref string) error {
	parsed, match, err := VertexAINotebookRuntimeIdentityFormat.Parse(ref)
	if err != nil {
		return fmt.Errorf("format of VertexAINotebookRuntime external=%q was not known (use %s): %w", ref, VertexAINotebookRuntimeIdentityFormat.CanonicalForm(), err)
	}
	if !match {
		return fmt.Errorf("format of VertexAINotebookRuntime external=%q was not known (use %s)", ref, VertexAINotebookRuntimeIdentityFormat.CanonicalForm())
	}

	*i = *parsed
	return nil
}

func (i *VertexAINotebookRuntimeIdentity) Host() string {
	return VertexAINotebookRuntimeIdentityFormat.Host()
}

func (i *VertexAINotebookRuntimeIdentity) ParentString() string {
	return fmt.Sprintf("projects/%s/locations/%s", i.Project, i.Location)
}

func getIdentityFromVertexAINotebookRuntimeSpec(ctx context.Context, reader client.Reader, obj *VertexAINotebookRuntime) (*VertexAINotebookRuntimeIdentity, error) {
	resourceID, err := refs.GetResourceID(obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectID, err := refs.ResolveProjectID(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve project")
	}

	identity := &VertexAINotebookRuntimeIdentity{
		Project:         projectID,
		Location:        location,
		NotebookRuntime: resourceID,
	}
	return identity, nil
}

func (obj *VertexAINotebookRuntime) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	specIdentity, err := getIdentityFromVertexAINotebookRuntimeSpec(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		statusIdentity := &VertexAINotebookRuntimeIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}

		if statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("cannot change VertexAINotebookRuntime identity (old=%q, new=%q)", statusIdentity.String(), specIdentity.String())
		}
	}

	return specIdentity, nil
}
