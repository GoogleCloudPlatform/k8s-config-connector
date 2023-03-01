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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	rrdatasFieldName     = "rrdatas"
	rrdatasRefsFieldName = "rrdatasRefs"
)

func GetDNSRecordSetOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "DNSRecordSet",
	}
	// Preserve the legacy non-reference field 'rrdatas' after it is changed to
	// a reference field, 'rrdatasRefs'.
	ro.Overrides = append(ro.Overrides, preserveRrdatasFieldAndEnsureRrdatasRefsFieldIsMultiKind())
	return ro
}

func preserveRrdatasFieldAndEnsureRrdatasRefsFieldIsMultiKind() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {

		if err := PreserveMutuallyExclusiveNonReferenceField(crd, nil, rrdatasRefsFieldName, rrdatasFieldName); err != nil {
			return fmt.Errorf("error preserving '%v' field in DNSRecordSet: %w", rrdatasFieldName, err)
		}
		if err := EnsureReferenceFieldIsMultiKind(crd, nil, rrdatasRefsFieldName, []string{"ComputeAddress"}); err != nil {
			return fmt.Errorf("error ensuring '%v' field in DNSRecordSet is a multi-kind reference field: %w", rrdatasRefsFieldName, err)
		}
		return nil
	}
	o.PreActuationTransform = func(r *k8s.Resource) error {
		if err := FavorReferenceArrayFieldOverNonReferenceArrayField(r, []string{rrdatasFieldName}, []string{rrdatasRefsFieldName}); err != nil {
			return fmt.Errorf("error handling '%v' and '%v' fields in pre-actuation transformation: %w", rrdatasFieldName, rrdatasRefsFieldName, err)
		}
		return nil
	}
	o.PostActuationTransform = func(original, reconciled *k8s.Resource, tfState *terraform.InstanceState, dclState *unstructured.Unstructured) error {
		if err := PreserveUserSpecifiedLegacyArrayField(original, reconciled, []string{rrdatasFieldName}...); err != nil {
			return fmt.Errorf("error preserving user-specified '%v' in post-actuation transformation: %w", rrdatasFieldName, err)
		}
		if err := PruneDefaultedAuthoritativeArrayFieldIfOnlyLegacyArrayFieldSpecified(original, reconciled, []string{rrdatasFieldName}, []string{rrdatasRefsFieldName}); err != nil {
			return fmt.Errorf("error conditionally pruning defaulted '%v' in post-actuation transformation: %w", rrdatasRefsFieldName, err)
		}
		return nil
	}
	return o
}
