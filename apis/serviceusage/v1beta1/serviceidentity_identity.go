// Copyright 2025 Google LLC
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

// +tool:krm-identity
// proto.service: google.api.serviceusage.v1beta1.ServiceUsage
// proto.message: google.api.serviceusage.v1beta1.ServiceIdentity
// crd.kind: ServiceIdentity
// crd.version: v1beta1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ identity.Identity = &ServiceIdentityIdentity{}

type ServiceIdentityIdentity struct {
	ParentID *refs.Project
	Service  string
}

func (i *ServiceIdentityIdentity) String() string {
	return i.ParentID.String() + "/services/" + i.Service
}

func (i *ServiceIdentityIdentity) FromExternal(ref string) error {
	tokens := strings.Split(ref, "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "services" {

		parentID := &refs.Project{}
		if err := parentID.FromExternal(strings.Join(tokens[:len(tokens)-2], "/")); err != nil {
			return fmt.Errorf("format of ServiceIdentity ref=%q was not known (use %q)", ref, "projects/<project>/services/<service>")
		}

		service := tokens[len(tokens)-1]

		i.ParentID = parentID
		i.Service = service

		return nil
	}

	return fmt.Errorf("format of ServiceIdentity ref=%q was not known (use %q)", ref, "projects/<project>/services/<service>")

}

var _ identity.Resource = &ServiceIdentity{}

func (obj *ServiceIdentity) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get parent ID
	parentID, err := obj.GetParentIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	id := &ServiceIdentityIdentity{
		ParentID: parentID.(*refs.Project),
		Service:  resourceID,
	}

	// status.externalRef does not exist yet
	// // Attempt to ensure ID is immutable, by verifying against previously-set `status.externalRef`.
	// externalRef := common.ValueOf(obj.Status.ExternalRef)
	// if externalRef != "" {
	// 	previousID := &ServiceIdentityIdentity{}
	// 	if err := previousID.FromExternal(externalRef); err != nil {
	// 		return nil, err
	// 	}
	// 	if id.String() != previousID.String() {
	// 		return nil, fmt.Errorf("cannot update ServiceIdentity identity (old=%q, new=%q): identity is immutable", previousID.String(), id.String())
	// 	}
	// }

	return id, nil
}

func (obj *ServiceIdentity) GetParentIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Normalize parent reference
	projectRef := *obj.Spec.ProjectRef
	if err := projectRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	// Get parent identity
	parentID := &refs.Project{}
	if err := parentID.FromExternal(obj.Spec.ProjectRef.External); err != nil {
		return nil, err
	}
	return parentID, nil
}
