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

import (
	"context"
	"fmt"
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
)

// AccessPolicyIdentity defines the resource reference to AccessContextManagerAccessPolicy, which "External" field
// holds the GCP identifier for the KRM object.
// +k8s:deepcopy-gen=false
type AccessPolicyIdentity struct {
	resourceID string
	title      string
}

func (i *AccessPolicyIdentity) String() string {
	// return "projects/" + i.ProjectID + "/accesspolicys/" + i.resourceID
	return "/accesspolicys/" + i.resourceID
}

func (i *AccessPolicyIdentity) ResourceID() string {
	return i.resourceID
}

func (i *AccessPolicyIdentity) Title() string {
	return i.title
}

// New builds a AccessPolicyIdentity from the Config Connector AccessPolicy object.
func NewAccessPolicyIdentity(ctx context.Context, reader client.Reader, obj *AccessContextManagerAccessPolicy) (*AccessPolicyIdentity, error) {
	return &AccessPolicyIdentity{
		resourceID: *obj.Spec.ResourceID,
		title:      *obj.Spec.Title,
	}, nil
}

func ParseAccessPolicyExternal(external string) (resourceID string, err error) {
	// pattern: "accessPolicies/{access_policy}"
	tokens := strings.Split(external, "/")
	if len(tokens) != 2 || tokens[0] != "accesspolicys" {
		return "", fmt.Errorf("format of AccessContextManagerAccessPolicy external=%q was not known (use accessPolicies/{{access_policy}})", external)
	}
	resourceID = tokens[1]
	return resourceID, nil
}

var _ identity.Identity = &AccessPolicyIdentity{}
