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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
	"k8s.io/klog/v2"
)

func GetComputeForwardingRuleResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ComputeForwardingRule",
	}
	ro.Overrides = append(ro.Overrides, noLabelsOnCreate())
	return ro
}

func noLabelsOnCreate() ResourceOverride {
	o := ResourceOverride{}

	o.PreTerraformApply = func(ctx context.Context, op *operations.PreTerraformApply) error {
		// There's some unexpected validation in forwardingRules, only when targeting serviceAttachments (PSC).
		// We can't specify labels in the create operation.  Terraform gets this wrong: https://github.com/hashicorp/terraform-provider-google/issues/16255
		// If we want the create to succeed, we cannot pass the labels.
		// This does mean that the labels won't be applied on first reconciliation, but we don't have many options here.
		// We do expect the labels will be applied next-time round.
		// This is a shorter-term fix, we should investigate fixing terraform or possibly replacing terraform with something we can fix directly.
		if op.LiveState.Empty() {
			target, ok := op.TerraformConfig.Config["target"].(string)
			if ok && strings.Contains(target, "/serviceAttachments/") {
				klog.Infof("removing labels before creating forwardingRule with serviceAttachment target")
				delete(op.TerraformConfig.Config, "labels")
			}
		}

		return nil
	}

	return o
}
