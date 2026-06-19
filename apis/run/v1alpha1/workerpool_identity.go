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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpurls"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.IdentityV2 = &CloudRunWorkerPoolIdentity{}

var workerPoolURL = gcpurls.Template[CloudRunWorkerPoolIdentity](
	"run.googleapis.com",
	"projects/{project}/locations/{location}/workerPools/{workerpool}",
)

// CloudRunWorkerPoolIdentity defines the resource reference to CloudRunWorkerPool, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type CloudRunWorkerPoolIdentity struct {
	Project    string `json:"project"`
	Location   string `json:"location"`
	WorkerPool string `json:"worker_pool"`
}

func (i *CloudRunWorkerPoolIdentity) FromExternal(ref string) error {
	out, match, err := workerPoolURL.Parse(ref)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("format of CloudRunWorkerPool external=%q was not known (use %s)", ref, workerPoolURL.CanonicalForm())
	}
	*i = *out
	return nil
}

func (i *CloudRunWorkerPoolIdentity) String() string {
	return workerPoolURL.ToString(*i)
}

func (i *CloudRunWorkerPoolIdentity) ID() string {
	return i.WorkerPool
}

func (i *CloudRunWorkerPoolIdentity) Host() string {
	return workerPoolURL.Host()
}

// New builds a CloudRunWorkerPoolIdentity from the Config Connector CloudRunWorkerPool object.
func NewCloudRunWorkerPoolIdentity(ctx context.Context, reader client.Reader, obj *CloudRunWorkerPool) (*CloudRunWorkerPoolIdentity, error) {

	// Get Parent
	projectRef, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	location := common.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	// Get desired ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	// Use approved External
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		// Validate desired with actual
		actualIdentity := &CloudRunWorkerPoolIdentity{}
		if err := actualIdentity.FromExternal(externalRef); err != nil {
			return nil, err
		}
		if actualIdentity.Project != projectID {
			return nil, fmt.Errorf("spec.projectRef changed, expect %s, got %s", actualIdentity.Project, projectID)
		}
		if actualIdentity.Location != location {
			return nil, fmt.Errorf("spec.location changed, expect %s, got %s", actualIdentity.Location, location)
		}
		if actualIdentity.WorkerPool != resourceID {
			return nil, fmt.Errorf("cannot reset `metadata.name` or `spec.resourceID` to %s, since it has already assigned to %s",
				resourceID, actualIdentity.WorkerPool)
		}
	}
	return &CloudRunWorkerPoolIdentity{
		Project:    projectID,
		Location:   location,
		WorkerPool: resourceID,
	}, nil
}

// GetIdentity implements the identity.Resource interface.
func (o *CloudRunWorkerPool) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	return NewCloudRunWorkerPoolIdentity(ctx, reader, o)
}

// ExternalIdentifier implements the identity.ExternalIdentifier interface.
func (o *CloudRunWorkerPool) ExternalIdentifier() *string {
	if o.Status.ExternalRef != nil {
		return o.Status.ExternalRef
	}
	// Fallback to the object's identity, this might not be accurate if the resource is not yet created.
	// We recommend controllers use the selfLink from the API response instead.
	return nil
}
