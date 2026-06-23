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
// See the License for the_specific language governing permissions and
// limitations under the License.

package compute

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ComputeDiskGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeDisk",
	}
	ComputeAddressGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeAddress",
	}
	ComputeNetworkGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNetwork",
	}
	ComputeSubnetworkGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeSubnetwork",
	}
	ComputeImageGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeImage",
	}
	ComputeInstanceTemplateGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeInstanceTemplate",
	}
	ComputeResourcePolicyGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeResourcePolicy",
	}
	KMSCryptoKeyGVK = schema.GroupVersionKind{
		Group:   "kms.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "KMSCryptoKey",
	}
	IAMServiceAccountGVK = schema.GroupVersionKind{
		Group:   "iam.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "IAMServiceAccount",
	}
)

func resolveComputeInstanceRefs(ctx context.Context, reader client.Reader, obj *krm.ComputeInstance) error {
	defaultNamespace := obj.GetNamespace()

	if obj.Spec.InstanceTemplateRef != nil {
		if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, obj.Spec.InstanceTemplateRef, ComputeInstanceTemplateGVK); err != nil {
			return err
		}
	}

	for i := range obj.Spec.ResourcePolicies {
		if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, &obj.Spec.ResourcePolicies[i], ComputeResourcePolicyGVK); err != nil {
			return err
		}
	}

	if obj.Spec.BootDisk != nil {
		if obj.Spec.BootDisk.KmsKeyRef != nil {
			if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, obj.Spec.BootDisk.KmsKeyRef, KMSCryptoKeyGVK); err != nil {
				return err
			}
		}
		if obj.Spec.BootDisk.SourceDiskRef != nil {
			if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, obj.Spec.BootDisk.SourceDiskRef, ComputeDiskGVK); err != nil {
				return err
			}
		}
		if obj.Spec.BootDisk.InitializeParams != nil && obj.Spec.BootDisk.InitializeParams.SourceImageRef != nil {
			if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, obj.Spec.BootDisk.InitializeParams.SourceImageRef, ComputeImageGVK); err != nil {
				return err
			}
		}
	}

	for i := range obj.Spec.AttachedDisk {
		disk := &obj.Spec.AttachedDisk[i]
		if disk.KmsKeyRef != nil {
			if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, disk.KmsKeyRef, KMSCryptoKeyGVK); err != nil {
				return err
			}
		}
		if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, &disk.SourceDiskRef, ComputeDiskGVK); err != nil {
			return err
		}
	}

	for i := range obj.Spec.NetworkInterface {
		ni := &obj.Spec.NetworkInterface[i]
		if ni.NetworkRef != nil {
			if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, ni.NetworkRef, ComputeNetworkGVK); err != nil {
				return err
			}
		}
		if ni.SubnetworkRef != nil {
			if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, ni.SubnetworkRef, ComputeSubnetworkGVK); err != nil {
				return err
			}
		}
		if ni.NetworkIpRef != nil {
			if err := resolveResourceRef(ctx, reader, defaultNamespace, ni.NetworkIpRef, ComputeAddressGVK); err != nil {
				return err
			}
		}
		for j := range ni.AccessConfig {
			ac := &ni.AccessConfig[j]
			if ac.NatIpRef != nil {
				if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, ac.NatIpRef, ComputeAddressGVK); err != nil {
					return err
				}
			}
		}
	}

	if obj.Spec.ServiceAccount != nil && obj.Spec.ServiceAccount.ServiceAccountRef != nil {
		if err := resolveInstanceResourceRef(ctx, reader, defaultNamespace, obj.Spec.ServiceAccount.ServiceAccountRef, IAMServiceAccountGVK); err != nil {
			return err
		}
	}

	return nil
}

func resolveInstanceResourceRef(ctx context.Context, reader client.Reader, defaultNamespace string, ref *krm.InstanceResourceRef, gvk schema.GroupVersionKind) error {
	if ref == nil {
		return nil
	}
	if ref.External != "" {
		ref.External = apirefs.TrimComputeURIPrefix(ref.External)
		return nil
	}
	if ref.Name == "" {
		return nil
	}
	namespace := ref.Namespace
	if namespace == "" {
		namespace = defaultNamespace
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	key := types.NamespacedName{
		Namespace: namespace,
		Name:      ref.Name,
	}
	if err := reader.Get(ctx, key, u); err != nil {
		return fmt.Errorf("getting referenced %s %s: %w", gvk.Kind, key, err)
	}

	externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
	if externalRef != "" {
		ref.External = apirefs.TrimComputeURIPrefix(externalRef)
		return nil
	}
	selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
	if selfLink != "" {
		ref.External = apirefs.TrimComputeURIPrefix(selfLink)
		return nil
	}
	email, _, _ := unstructured.NestedString(u.Object, "status", "email")
	if email != "" {
		ref.External = email
		return nil
	}

	return fmt.Errorf("referenced %s %s does not have status.externalRef, status.selfLink, or status.email populated", gvk.Kind, key)
}

func resolveResourceRef(ctx context.Context, reader client.Reader, defaultNamespace string, ref *k8sv1alpha1.ResourceRef, gvk schema.GroupVersionKind) error {
	if ref == nil {
		return nil
	}
	if ref.External != "" {
		ref.External = apirefs.TrimComputeURIPrefix(ref.External)
		return nil
	}
	if ref.Name == "" {
		return nil
	}
	namespace := ref.Namespace
	if namespace == "" {
		namespace = defaultNamespace
	}

	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(gvk)
	key := types.NamespacedName{
		Namespace: namespace,
		Name:      ref.Name,
	}
	if err := reader.Get(ctx, key, u); err != nil {
		return fmt.Errorf("getting referenced %s %s: %w", gvk.Kind, key, err)
	}

	externalRef, _, _ := unstructured.NestedString(u.Object, "status", "externalRef")
	if externalRef != "" {
		ref.External = apirefs.TrimComputeURIPrefix(externalRef)
		return nil
	}
	selfLink, _, _ := unstructured.NestedString(u.Object, "status", "selfLink")
	if selfLink != "" {
		ref.External = apirefs.TrimComputeURIPrefix(selfLink)
		return nil
	}
	email, _, _ := unstructured.NestedString(u.Object, "status", "email")
	if email != "" {
		ref.External = email
		return nil
	}

	return fmt.Errorf("referenced %s %s does not have status.externalRef, status.selfLink, or status.email populated", gvk.Kind, key)
}
