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

package binaryauthorization

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/binaryauthorization/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func resolveBinaryAuthorizationPolicyRefs(ctx context.Context, reader client.Reader, obj *krm.BinaryAuthorizationPolicy) error {
	if obj.Spec.ProjectRef != nil {
		res, err := refsv1beta1.ResolveProject(ctx, reader, obj.Namespace, obj.Spec.ProjectRef)
		if err != nil {
			return err
		}
		obj.Spec.ProjectRef.External = res.ProjectID
	}

	if obj.Spec.DefaultAdmissionRule != nil {
		if err := resolveAdmissionRuleRefs(ctx, reader, obj, obj.Spec.DefaultAdmissionRule); err != nil {
			return err
		}
	}

	for k, v := range obj.Spec.ClusterAdmissionRules {
		if err := resolveAdmissionRuleRefs(ctx, reader, obj, &v); err != nil {
			return err
		}
		obj.Spec.ClusterAdmissionRules[k] = v
	}

	for k, v := range obj.Spec.KubernetesNamespaceAdmissionRules {
		if err := resolveAdmissionRuleRefs(ctx, reader, obj, &v); err != nil {
			return err
		}
		obj.Spec.KubernetesNamespaceAdmissionRules[k] = v
	}

	for k, v := range obj.Spec.KubernetesServiceAccountAdmissionRules {
		if err := resolveAdmissionRuleRefs(ctx, reader, obj, &v); err != nil {
			return err
		}
		obj.Spec.KubernetesServiceAccountAdmissionRules[k] = v
	}

	for k, v := range obj.Spec.IstioServiceIdentityAdmissionRules {
		if err := resolveAdmissionRuleRefs(ctx, reader, obj, &v); err != nil {
			return err
		}
		obj.Spec.IstioServiceIdentityAdmissionRules[k] = v
	}

	return nil
}

func resolveAdmissionRuleRefs(ctx context.Context, reader client.Reader, src client.Object, rule *krm.AdmissionRule) error {
	for i := range rule.RequireAttestationsBy {
		ref := &rule.RequireAttestationsBy[i]
		if ref.External != "" {
			continue
		}
		if ref.Name == "" {
			return fmt.Errorf("RequireAttestationsBy[%d] has no name or external", i)
		}
		key := types.NamespacedName{
			Namespace: ref.Namespace,
			Name:      ref.Name,
		}
		if key.Namespace == "" {
			key.Namespace = src.GetNamespace()
		}
		u := &unstructured.Unstructured{}
		u.SetGroupVersionKind(schema.GroupVersionKind{
			Group:   "binaryauthorization.cnrm.cloud.google.com",
			Version: "v1beta1",
			Kind:    "BinaryAuthorizationAttestor",
		})
		if err := reader.Get(ctx, key, u); err != nil {
			return err
		}
		// Typically we want the full resource name: projects/*/attestors/*
		// But BinaryAuthorizationAttestor might have its own resolution logic.
		// For now, I'll use a placeholder or check if there's a helper.
		resourceID, _, _ := unstructured.NestedString(u.Object, "spec", "resourceID")
		if resourceID == "" {
			resourceID = u.GetName()
		}
		projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
		if err != nil {
			return err
		}
		ref.External = fmt.Sprintf("projects/%s/attestors/%s", projectID, resourceID)
	}
	return nil
}
