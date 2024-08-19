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

package compute

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func resolveResource(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef, gvk schema.GroupVersionKind) (*unstructured.Unstructured, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return nil, nil
	}

	if ref.Name == "" {
		return nil, fmt.Errorf("must specify either name or external on reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = src.GetNamespace()
	}

	resource := &unstructured.Unstructured{}
	resource.SetGroupVersionKind(gvk)
	if err := reader.Get(ctx, key, resource); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, k8s.NewReferenceNotFoundError(resource.GroupVersionKind(), key)
		}
		return nil, fmt.Errorf("error reading referenced %v %v: %w", gvk.Kind, key, err)
	}

	return resource, nil
}

func ResolveComputeNetwork(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeNetwork, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeNetwork",
	})
	if err != nil {
		return nil, err
	}
	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeNetwork.GetSelfLink()}, nil
}

func ResolveComputeSubnetwork(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeSubnetwork, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeSubnetwork",
	})
	if err != nil {
		return nil, err
	}
	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeSubnetwork.GetSelfLink(),
	}, nil
}

func ResolveComputeAddress(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeAddress, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeAddress",
	})
	if err != nil {
		return nil, err
	}

	// targetField: address
	// See compute servicemappings for details
	address, _, _ := unstructured.NestedString(computeAddress.Object, "spec", "address")
	if address == "" {
		return nil, fmt.Errorf("cannot get address for referenced %s %v (spec.address is empty)", computeAddress.GetKind(), computeAddress.GetNamespace())
	}
	return &v1alpha1.ResourceRef{
		External: address}, nil
}

func ResolveComputeBackendService(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeBackendService, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeBackendService",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeBackendService.GetSelfLink()}, nil
}

func ResolveComputeServiceAttachment(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeServiceAttachment, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeServiceAttachment",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeServiceAttachment.GetSelfLink()}, nil
}

func ResolveComputeTargetGrpcProxy(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeTargetGrpcProxy, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetGrpcProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeTargetGrpcProxy.GetSelfLink()}, nil
}

func ResolveComputeTargetHTTPProxy(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeTargetHTTPProxy, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetHTTPProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeTargetHTTPProxy.GetSelfLink()}, nil
}

func ResolveComputeTargetHTTPSProxy(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeTargetHTTPSProxy, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetHTTPSProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeTargetHTTPSProxy.GetSelfLink()}, nil
}

func ResolveComputeTargetSSLProxy(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeTargetSSLProxy, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetSSLProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeTargetSSLProxy.GetSelfLink()}, nil
}

func ResolveComputeTargetTCPProxy(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeTargetTCPProxy, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetTCPProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeTargetTCPProxy.GetSelfLink()}, nil
}

func ResolveComputeTargetVPNGateway(ctx context.Context, reader client.Reader, src client.Object, ref *v1alpha1.ResourceRef) (*v1alpha1.ResourceRef, error) {
	computeTargetVPNGateway, err := resolveResource(ctx, reader, src, ref, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetVPNGateway",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	return &v1alpha1.ResourceRef{
		External: computeTargetVPNGateway.GetSelfLink()}, nil
}
