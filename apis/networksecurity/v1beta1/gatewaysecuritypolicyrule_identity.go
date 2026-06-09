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

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type GatewaySecurityPolicyRuleIdentity struct {
	Parent string
	ID     string
}

func (i *GatewaySecurityPolicyRuleIdentity) String() string {
	return i.Parent + "/rules/" + i.ID
}

// GetIdentity resolves the identity of the GatewaySecurityPolicyRule resource.
func (obj *NetworkSecurityGatewaySecurityPolicyRule) GetIdentity(ctx context.Context, reader client.Reader) (*GatewaySecurityPolicyRuleIdentity, error) {
	if obj.Spec.GatewaySecurityPolicyRef.External == "" {
		return nil, fmt.Errorf("GatewaySecurityPolicyRef.External must be resolved")
	}

	parent := obj.Spec.GatewaySecurityPolicyRef.External

	id := obj.GetName()
	if obj.Spec.ResourceID != nil {
		id = *obj.Spec.ResourceID
	}
	if obj.Status.ExternalRef != nil {
		id = parseGatewaySecurityPolicyRuleID(*obj.Status.ExternalRef)
	}

	return &GatewaySecurityPolicyRuleIdentity{
		Parent: parent,
		ID:     id,
	}, nil
}

func parseGatewaySecurityPolicyRuleID(externalRef string) string {
	parts := strings.Split(externalRef, "/")
	return parts[len(parts)-1]
}
