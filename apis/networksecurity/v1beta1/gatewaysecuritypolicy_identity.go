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
	"strings"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type GatewaySecurityPolicyIdentity struct {
	Parent string
	ID     string
}

func (i *GatewaySecurityPolicyIdentity) String() string {
	return i.Parent + "/gatewaySecurityPolicies/" + i.ID
}

// GetIdentity resolves the identity of the GatewaySecurityPolicy resource.
func (obj *NetworkSecurityGatewaySecurityPolicy) GetIdentity(ctx context.Context, reader client.Reader) (*GatewaySecurityPolicyIdentity, error) {
	if obj.Spec.ProjectRef.Name == "" && obj.Spec.ProjectRef.External == "" {
		return nil, fmt.Errorf("projectRef must be specified")
	}

	// Ensure ProjectRef is resolved
	project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	parent := fmt.Sprintf("projects/%s/locations/%s", project.ProjectID, obj.Spec.Location)

	id := obj.GetName()
	if obj.Spec.ResourceID != nil {
		id = *obj.Spec.ResourceID
	}
	if obj.Status.ExternalRef != nil {
		id = parseGatewaySecurityPolicyID(*obj.Status.ExternalRef)
	}

	return &GatewaySecurityPolicyIdentity{
		Parent: parent,
		ID:     id,
	}, nil
}

func parseGatewaySecurityPolicyID(externalRef string) string {
	parts := strings.Split(externalRef, "/")
	return parts[len(parts)-1]
}
