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

package forwardingrule

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveComputeAddress(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeAddressRef) (*refs.ComputeAddressRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
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

	computeAddress, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeAddress",
	})
	if err != nil {
		return nil, err
	}

	// targetField: address
	// See compute servicemappings for details
	// Because `spec.address` field is optional, we can't guarantee it always
	// exists in a successfully reconciled ComputeAddress CR, so we should use
	// the `status.observedState.address` instead.
	address, _, err := unstructured.NestedString(computeAddress.Object, "status", "observedState", "address")
	if err != nil || address == "" {
		return nil, fmt.Errorf("cannot get address for referenced %s %v (status.observedState.address is empty)", computeAddress.GetKind(), computeAddress.GetNamespace())
	}
	return &refs.ComputeAddressRef{
		External: address}, nil
}

func ResolveComputeServiceAttachment(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeServiceAttachmentRef) (*refs.ComputeServiceAttachmentRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
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

	computeServiceAttachment, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeServiceAttachment",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	selfLink, _, err := unstructured.NestedString(computeServiceAttachment.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return nil, fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", computeServiceAttachment.GetKind(), computeServiceAttachment.GetNamespace())
	}
	return &refs.ComputeServiceAttachmentRef{
		External: selfLink}, nil
}

func ResolveComputeTargetGrpcProxy(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeTargetGrpcProxyRef) (*refs.ComputeTargetGrpcProxyRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
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

	computeTargetGrpcProxy, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetGRPCProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	selfLink, _, err := unstructured.NestedString(computeTargetGrpcProxy.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return nil, fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", computeTargetGrpcProxy.GetKind(), computeTargetGrpcProxy.GetNamespace())
	}
	return &refs.ComputeTargetGrpcProxyRef{
		External: selfLink}, nil
}

func ResolveComputeTargetHTTPProxy(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeTargetHTTPProxyRef) (*refs.ComputeTargetHTTPProxyRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
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

	computeTargetHTTPProxy, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetHTTPProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	selfLink, _, err := unstructured.NestedString(computeTargetHTTPProxy.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return nil, fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", computeTargetHTTPProxy.GetKind(), computeTargetHTTPProxy.GetNamespace())
	}
	return &refs.ComputeTargetHTTPProxyRef{
		External: selfLink}, nil
}

func ResolveComputeTargetHTTPSProxy(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeTargetHTTPSProxyRef) (*refs.ComputeTargetHTTPSProxyRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
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

	computeTargetHTTPSProxy, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetHTTPSProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	selfLink, _, err := unstructured.NestedString(computeTargetHTTPSProxy.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return nil, fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", computeTargetHTTPSProxy.GetKind(), computeTargetHTTPSProxy.GetNamespace())
	}
	return &refs.ComputeTargetHTTPSProxyRef{
		External: selfLink}, nil
}

func ResolveComputeTargetSSLProxy(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeTargetSSLProxyRef) (*refs.ComputeTargetSSLProxyRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
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

	computeTargetSSLProxy, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetSSLProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	selfLink, _, err := unstructured.NestedString(computeTargetSSLProxy.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return nil, fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", computeTargetSSLProxy.GetKind(), computeTargetSSLProxy.GetNamespace())
	}
	return &refs.ComputeTargetSSLProxyRef{
		External: selfLink}, nil
}

func ResolveComputeTargetTCPProxy(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeTargetTCPProxyRef) (*refs.ComputeTargetTCPProxyRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
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

	computeTargetTCPProxy, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetTCPProxy",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	selfLink, _, err := unstructured.NestedString(computeTargetTCPProxy.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return nil, fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", computeTargetTCPProxy.GetKind(), computeTargetTCPProxy.GetNamespace())
	}
	return &refs.ComputeTargetTCPProxyRef{
		External: selfLink}, nil
}

func ResolveComputeTargetVPNGateway(ctx context.Context, reader client.Reader, src client.Object, ref *refs.ComputeTargetVPNGatewayRef) (*refs.ComputeTargetVPNGatewayRef, error) {
	if ref == nil {
		return nil, nil
	}

	if ref.External != "" {
		if ref.Name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on reference")
		}
		return ref, nil
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

	computeTargetVPNGateway, err := resolveResourceName(ctx, reader, key, schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetVPNGateway",
	})
	if err != nil {
		return nil, err
	}

	// targetField: self_link
	// See compute servicemappings for details
	selfLink, _, err := unstructured.NestedString(computeTargetVPNGateway.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return nil, fmt.Errorf("cannot get selfLink for referenced %s %v (status.selfLink is empty)", computeTargetVPNGateway.GetKind(), computeTargetVPNGateway.GetNamespace())
	}
	return &refs.ComputeTargetVPNGatewayRef{
		External: selfLink}, nil
}

func resolveResourceName(ctx context.Context, reader client.Reader, key client.ObjectKey, gvk schema.GroupVersionKind) (*unstructured.Unstructured, error) {
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
