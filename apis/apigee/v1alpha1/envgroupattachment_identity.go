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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	apigeev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigee/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type EnvgroupAttachmentIdentity struct {
	parent *EnvgroupAttachmentParent
	id     string
}

func (i *EnvgroupAttachmentIdentity) String() string {
	return fmt.Sprintf("%s/attachments/%s", i.parent, i.id)
}

func (i *EnvgroupAttachmentIdentity) ID() string {
	return i.id
}

func (i *EnvgroupAttachmentIdentity) Parent() *EnvgroupAttachmentParent {
	return i.parent
}

type EnvgroupAttachmentParent struct {
	Organization string
	Envgroup     string
}

func (p *EnvgroupAttachmentParent) String() string {
	return "organizations/" + p.Organization + "/envgroups/" + p.Envgroup
}

func NewEnvgroupAttachmentIdentity(ctx context.Context, reader client.Reader, obj *ApigeeEnvgroupAttachment) (*EnvgroupAttachmentIdentity, error) {
	if obj == nil {
		return nil, fmt.Errorf("object cannot be nil")
	}

	// Get Parent
	orgRef := obj.Spec.OrganizationRef
	if orgRef == nil {
		return nil, fmt.Errorf("no parent organization")
	}
	orgExternal, err := orgRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve organization: %w", err)
	}

	org, err := apigeev1beta1.ParseOrganizationExternal(orgExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse external organization: %w", err)
	}

	envgroupRef := obj.Spec.EnvgroupRef
	if envgroupRef == nil {
		return nil, fmt.Errorf("no envgroup reference")
	}
	envgroupExternal, err := envgroupRef.NormalizedExternal(ctx, reader, obj.Namespace)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve envgroup: %w", err)
	}
	_, envgroupID, err := ParseEnvironmentGroupExternal(envgroupExternal)
	if err != nil {
		return nil, fmt.Errorf("cannot parse envgroup external: %w", err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		actualParent, actualResourceID, err := ParseEnvgroupAttachmentExternalRef(externalRef)
		if err != nil {
			return nil, err
		}

		if actualParent.Organization != org {
			return nil, fmt.Errorf("spec.organizationRef changed, expect %s, got %s", actualParent.Organization, org)
		}
		if actualParent.Envgroup != envgroupID {
			return nil, fmt.Errorf("spec.envgroup changed, expect %s, got %s", actualParent.Envgroup, obj.Spec.EnvgroupRef)
		}
		if actualResourceID != actualResourceID {
			return nil, fmt.Errorf("spec.resourceID changed, expect %s, got %s", actualResourceID, resourceID)
		}
	}

	return &EnvgroupAttachmentIdentity{
		parent: &EnvgroupAttachmentParent{
			Organization: org,
			Envgroup:     envgroupID,
		},
		id: resourceID,
	}, nil
}

func ParseEnvgroupAttachmentExternalRef(external string) (parent *EnvgroupAttachmentParent, resourceID string, err error) {
	tokens := strings.Split(external, "/")
	if len(tokens) != 6 {
		return nil, "", fmt.Errorf("invalid external format: %s, expecting organizations/{{organizationID}}/envgroups/{{envgroupID}}/attachments/{{attachmentID}} ", external)
	}

	parent = &EnvgroupAttachmentParent{
		Organization: tokens[1],
		Envgroup:     tokens[3],
	}
	resourceID = tokens[5]
	return parent, resourceID, nil
}
