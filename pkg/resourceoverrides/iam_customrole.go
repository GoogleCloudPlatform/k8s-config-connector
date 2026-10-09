// Copyright 2022 Google LLC
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

package resourceoverrides

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var iamCustomRoleIDRegex = regexp.MustCompile(`^[a-zA-Z0-9_\.]{3,64}$`)

func GetIAMCustomRoleResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "IAMCustomRole",
	}
	ro.Overrides = append(ro.Overrides, buildIAMCustomRole())
	return ro
}

func buildIAMCustomRole() ResourceOverride {
	h := &IAMCustomRole{}
	o := ResourceOverride{
		ConfigValidate:        h.ConfigValidate,
		PreActuationTransform: h.PreActuationTransform,
		PreTerraformExport:    h.PreTerraformExport,
	}
	return o
}

type IAMCustomRole struct {
}

func (h *IAMCustomRole) ConfigValidate(u *unstructured.Unstructured) error {
	resourceID, exists, err := unstructured.NestedString(u.Object, "spec", k8s.ResourceIDFieldName)
	if err != nil {
		return fmt.Errorf("error reading spec.%s: %w", k8s.ResourceIDFieldName, err)
	}
	if exists && resourceID != "" {
		if !iamCustomRoleIDRegex.MatchString(resourceID) {
			return fmt.Errorf("invalid spec.%s %q: IAM custom role IDs must match regex %s (letters, numbers, underscores, and periods only, length 3-64; hyphens are not allowed)", k8s.ResourceIDFieldName, resourceID, iamCustomRoleIDRegex.String())
		}
	}
	return nil
}

func (h *IAMCustomRole) PreActuationTransform(r *k8s.Resource) error {
	resourceID, exists, err := unstructured.NestedString(r.Spec, k8s.ResourceIDFieldName)
	if err != nil {
		return fmt.Errorf("error reading spec.%s: %w", k8s.ResourceIDFieldName, err)
	}
	if !exists || resourceID == "" {
		// Default spec.resourceID from metadata.name with hyphens replaced by underscores,
		// because Google Cloud IAM custom role IDs do not allow hyphens (must match ^[a-zA-Z0-9_\.]{3,64}$).
		sanitized := strings.ReplaceAll(r.GetName(), "-", "_")
		if !iamCustomRoleIDRegex.MatchString(sanitized) {
			return fmt.Errorf("defaulted role ID %q from metadata.name %q is invalid: IAM custom role IDs must match regex %s (length 3-64, alphanumeric, underscores, and periods only)", sanitized, r.GetName(), iamCustomRoleIDRegex.String())
		}
		if err := unstructured.SetNestedField(r.Spec, sanitized, k8s.ResourceIDFieldName); err != nil {
			return fmt.Errorf("error defaulting sanitized spec.%s: %w", k8s.ResourceIDFieldName, err)
		}
	}
	return nil
}

func (h *IAMCustomRole) PreTerraformExport(_ context.Context, op *operations.TerraformExport) error {
	projectID := op.TerraformState.Attributes["project"]
	orgID := op.TerraformState.Attributes["org_id"]

	if projectID != "" && orgID != "" {
		return fmt.Errorf("project=%q and org_id=%q were both set", projectID, orgID)
	}

	if projectID != "" {
		op.TerraformInfo.Type = "google_project_iam_custom_role"
	} else if orgID != "" {
		op.TerraformInfo.Type = "google_organization_iam_custom_role"
	} else {
		return fmt.Errorf("unable to determine whether IAMCustomRole is org or project scoped")
	}

	return nil
}

