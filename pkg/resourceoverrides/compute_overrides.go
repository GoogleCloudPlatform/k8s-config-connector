// Copyright 2024 Google LLC
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
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
)

var (
	networkInterfacePath = []string{"networkInterface"}
	networkIPFieldPath   = []string{"networkIp"}

	oauth2ClientIDPath    = []string{"iap", "oauth2ClientId"}
	oauth2ClientIDRefPath = []string{"iap", "oauth2ClientIdRef"}

	networkIPRefFieldPath = []string{"networkIpRef"}
	supportedKinds        = []string{"ComputeAddress"}
)

func GetComputeManagedSSLCertificateResourceOverrides() ResourceOverrides {
	return ResourceOverrides{
		Kind: "ComputeManagedSSLCertificate",
		Overrides: []ResourceOverride{
			mapCertificateIdToCertificateID(),
		},
	}
}

func mapCertificateIdToCertificateID() ResourceOverride {
	return ResourceOverride{
		CRDDecorate: func(crd *apiextensions.CustomResourceDefinition) error {
			for _, version := range crd.Spec.Versions {
				schema := version.Schema.OpenAPIV3Schema
				status, ok := schema.Properties["status"]
				if !ok {
					return fmt.Errorf("status field not found for version %s", version.Name)
				}
				observedState, ok := status.Properties["observedState"]
				if !ok {
					return fmt.Errorf("observedState field not found for version %s", version.Name)
				}
				certificateId, ok := observedState.Properties["certificateId"]
				if !ok {
					return fmt.Errorf("certificateId field not found for version %s", version.Name)
				}

				// Rename certificateId to certificateID
				observedState.Properties["certificateID"] = certificateId
				delete(observedState.Properties, "certificateId")
			}

			return nil
		},
		PostActuationTransform: func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
			observedState, found := reconciled.Status["observedState"]
			if !found {
				// if there is no observedState there is nothing to do!
				return nil
			}
			observedStateM, ok := observedState.(map[string]interface{})
			if !ok {
				return fmt.Errorf("cannot parse observedState map")
			}
			certificateId, found := observedStateM["certificateId"]
			if !found {
				// field certificateId not populated
				return nil
			}

			observedStateM["certificateID"] = certificateId
			delete(observedStateM, "certificateId")

			reconciled.Status["observedState"] = observedStateM

			return nil
		},
	}
}

func GetComputeInstanceResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ComputeInstance",
	}
	ro.Overrides = append(ro.Overrides, addNetworkIPRefField())
	return ro
}

func addNetworkIPRefField() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		if err := PreserveMutuallyExclusiveNonReferenceField(crd, networkInterfacePath, networkIPRefFieldPath[0], networkIPFieldPath[0]); err != nil {
			return err
		}

		return EnsureReferenceFieldIsMultiKind(crd, networkInterfacePath, networkIPRefFieldPath[0], supportedKinds)
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := FavorReferenceFieldOverNonReferenceFieldUnderSlice(r, networkInterfacePath, networkIPFieldPath, networkIPRefFieldPath); err != nil {
			return fmt.Errorf("error handling '%v' and '%v' fields in pre-actuation transformation: %w", strings.Join(networkIPFieldPath, "."), strings.Join(networkIPRefFieldPath, "."), err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		if err := PreserveUserSpecifiedLegacyFieldUnderSlice(original, reconciled, networkInterfacePath, networkIPFieldPath); err != nil {
			return fmt.Errorf("error preserving '%v' in post-actuation transformation: %w", strings.Join(networkIPFieldPath, "."), err)
		}
		if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecifiedUnderSlice(original, reconciled, networkInterfacePath, networkIPFieldPath, networkIPRefFieldPath); err != nil {
			return fmt.Errorf("error conditionally pruning '%v' in post-actuation transformation: %w", strings.Join(networkIPRefFieldPath, "."), err)
		}
		return nil
	}
	return o
}

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

func GetComputeBackendServiceResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ComputeBackendService",
	}
	// Preserve the legacy non-reference field 'iap.oauth2ClientId' after it is changed to
	// a reference field, 'iap.oauth2ClientIdRef'.
	ro.Overrides = append(ro.Overrides, keepOauth2ClientIDField())
	return ro
}

func keepOauth2ClientIDField() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		return PreserveMutuallyExclusiveNonReferenceField(crd, []string{"iap"}, oauth2ClientIDRefPath[1], oauth2ClientIDPath[1])
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := FavorAuthoritativeFieldOverLegacyField(r, oauth2ClientIDPath, oauth2ClientIDRefPath); err != nil {
			return fmt.Errorf("error handling '%v' and '%v' fields in pre-actuation transformation: %w", strings.Join(oauth2ClientIDPath, "."), strings.Join(oauth2ClientIDRefPath, "."), err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		if err := PreserveUserSpecifiedLegacyField(original, reconciled, oauth2ClientIDPath...); err != nil {
			return fmt.Errorf("error preserving '%v' in post-actuation transformation: %w", strings.Join(oauth2ClientIDPath, "."), err)
		}
		if err := PruneDefaultedAuthoritativeFieldIfOnlyLegacyFieldSpecified(original, reconciled, oauth2ClientIDPath, oauth2ClientIDRefPath); err != nil {
			return fmt.Errorf("error conditionally pruning '%v' in post-actuation transformation: %w", strings.Join(oauth2ClientIDRefPath, "."), err)
		}
		return nil
	}
	return o
}
