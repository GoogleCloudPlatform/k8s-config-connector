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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
)

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
		PreTerraformExport: h.PreTerraformExport,
	}
	return o
}

type IAMCustomRole struct {
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
