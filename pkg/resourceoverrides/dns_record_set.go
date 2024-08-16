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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/crdutil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	rrdatasFieldName       = "rrdatas"
	rrdatasRefsFieldName   = "rrdatasRefs"
	routingPolicyFieldName = "routingPolicy"
	nameFieldName          = "name"

	rrdatasRefKinds = []MultiKindRef{
		{
			Kind:        "ComputeAddress",
			TargetField: "address",
		},
		{
			Kind:        "CertificateManagerDnsAuthorization",
			TargetField: "dnsResourceRecord.data",
		},
	}

	routingPolicyRefKinds = []MultiKindRef{
		{
			Kind: "ComputeAddress",
		},
	}
)

func GetDNSRecordSetOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "DNSRecordSet",
	}
	// Preserve the legacy non-reference field 'rrdatas' after it is changed to
	// a reference field, 'rrdatasRefs'.
	ro.Overrides = append(ro.Overrides, preserveRrdatasFieldAndEnsureRrdatasRefsFieldIsMultiKind())
	// Configure the top-level OneOf to make 'routingPolicy', 'rrdatas', 'rrdatasRef
	// and 'dnsAuthorizationsRef' mutually exclusive.
	ro.Overrides = append(ro.Overrides, enforceMutuallyExclusiveRrdatasAndRoutingPolicy())
	// Configure rrdatasRefs fields under routingPolicy to be MultiKind.
	ro.Overrides = append(ro.Overrides, ensureRoutingPoliciesRrDatasRefsFieldsAreMultiKind())
	// Configure rrdata
	return ro
}

func preserveRrdatasFieldAndEnsureRrdatasRefsFieldIsMultiKind() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {

		if err := PreserveMutuallyExclusiveNonReferenceField(crd, nil, rrdatasRefsFieldName, rrdatasFieldName); err != nil {
			return fmt.Errorf("error preserving '%v' field in DNSRecordSet: %w", rrdatasFieldName, err)
		}
		if err := EnsureReferenceFieldIsMultiKind(crd, nil, rrdatasRefsFieldName, rrdatasRefKinds); err != nil {
			return fmt.Errorf("error ensuring '%v' field in DNSRecordSet is a multi-kind reference field: %w", rrdatasRefsFieldName, err)
		}
		// PreserveMutuallyExclusiveNonReferenceField adds a `not` condition to
		// prevent rrdatas and rrdatasRefs from being set together. This is
		// redundant due to the enforceMutuallyExclusiveRrdatasAndRoutingPolicy
		// override, so we will manually remove it.
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		spec := schema.Properties["spec"]
		if err := crdutil.SetNotRuleForObjectOrArray(&spec, nil); err != nil {
			return err
		}

		return crdutil.SetSchemaForFieldUnderObjectOrArray("spec", schema, &spec)
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

func enforceMutuallyExclusiveRrdatasAndRoutingPolicy() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		schema := k8s.GetOpenAPIV3SchemaFromCRD(crd)
		spec := schema.Properties["spec"]
		requireField := func(field string) *apiextensions.JSONSchemaProps {
			return &apiextensions.JSONSchemaProps{
				Required: []string{field},
			}
		}
		if err := crdutil.SetOneOfRuleForObjectOrArray(&spec, []*apiextensions.JSONSchemaProps{
			requireField(rrdatasFieldName),
			requireField(rrdatasRefsFieldName),
			requireField(routingPolicyFieldName),
		}); err != nil {
			return err
		}

		return crdutil.SetSchemaForFieldUnderObjectOrArray("spec", schema, &spec)
	}
	return o
}

func ensureRoutingPoliciesRrDatasRefsFieldsAreMultiKind() ResourceOverride {
	o := ResourceOverride{}
	o.CRDDecorate = func(crd *apiextensions.CustomResourceDefinition) error {
		if err := EnsureReferenceFieldIsMultiKind(crd, []string{routingPolicyFieldName, "wrr"}, rrdatasRefsFieldName, routingPolicyRefKinds); err != nil {
			return fmt.Errorf("error ensuring '%v' field in DNSRecordSet is a multi-kind reference field: %w", rrdatasRefsFieldName, err)
		}
		if err := EnsureReferenceFieldIsMultiKind(crd, []string{routingPolicyFieldName, "geo"}, rrdatasRefsFieldName, routingPolicyRefKinds); err != nil {
			return fmt.Errorf("error ensuring '%v' field in DNSRecordSet is a multi-kind reference field: %w", rrdatasRefsFieldName, err)
		}
		return nil
	}
	return o
}
