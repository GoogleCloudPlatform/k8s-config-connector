// Copyright 2025 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"

	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ refsv1beta1.Ref = &ComputeTargetHTTPSProxyRef{}

type ComputeTargetHTTPSProxyRef struct {
	// Allowed value: string of the format `projects/{{project}}/global/targetHttpsProxies/{{value}}` or `projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{value}}`, where {{value}} is the `name` field of a `ComputeTargetHTTPSProxy` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

func (r *ComputeTargetHTTPSProxyRef) GetGVK() schema.GroupVersionKind {
	return ComputeTargetHTTPSProxyGVK
}

func (r *ComputeTargetHTTPSProxyRef) GetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Name:      r.Name,
		Namespace: r.Namespace,
	}
}

func (r *ComputeTargetHTTPSProxyRef) GetExternal() string {
	return r.External
}

func (r *ComputeTargetHTTPSProxyRef) SetExternal(ref string) {
	r.External = ref
}

func (r *ComputeTargetHTTPSProxyRef) ValidateExternal(ref string) error {
	id := &ComputeTargetHTTPSProxyIdentity{}
	if err := id.FromExternal(ref); err != nil {
		return err
	}
	return nil
}

func (r *ComputeTargetHTTPSProxyRef) Normalize(ctx context.Context, reader client.Reader, defaultNamespace string) error {
	return refsv1beta1.Normalize(ctx, reader, r, defaultNamespace)
}

// A reference to a ComputeURLMap resource.
type ComputeURLMapRef struct {
	// Allowed value: The `selfLink` field of a `ComputeURLMap` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

var ComputeURLMapGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeURLMap",
}

func (r *ComputeURLMapRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on ComputeURLMap reference")
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeURLMapGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced ComputeURLMap %s: %w", key, err)
	}

	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return "", fmt.Errorf("cannot get selfLink for referenced ComputeURLMap %v (status.selfLink is empty)", key)
	}
	return selfLink, nil
}

// A reference to a ComputeSSLCertificate resource.
type ComputeSSLCertificateRef struct {
	// Allowed value: string of the format `projects/{{project}}/global/sslCertificates/{{value}}` or `projects/{{project}}/regions/{{region}}/sslCertificates/{{value}}`, where {{value}} is the `name` field of a `ComputeSSLCertificate` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

var ComputeSSLCertificateGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeSSLCertificate",
}

func (r *ComputeSSLCertificateRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on ComputeSSLCertificate reference")
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeSSLCertificateGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced ComputeSSLCertificate %s: %w", key, err)
	}

	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return "", fmt.Errorf("cannot get selfLink for referenced ComputeSSLCertificate %v (status.selfLink is empty)", key)
	}
	return selfLink, nil
}

// A reference to a ComputeSSLPolicy resource.
type ComputeSSLPolicyRef struct {
	// Allowed value: The `selfLink` field of a `ComputeSSLPolicy` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

var ComputeSSLPolicyGVK = schema.GroupVersionKind{
	Group:   "compute.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "ComputeSSLPolicy",
}

func (r *ComputeSSLPolicyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on ComputeSSLPolicy reference")
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(ComputeSSLPolicyGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced ComputeSSLPolicy %s: %w", key, err)
	}

	selfLink, _, err := unstructured.NestedString(u.Object, "status", "selfLink")
	if err != nil || selfLink == "" {
		return "", fmt.Errorf("cannot get selfLink for referenced ComputeSSLPolicy %v (status.selfLink is empty)", key)
	}
	return selfLink, nil
}

// A reference to a CertificateManagerCertificate resource.
type CertificateManagerCertificateRef struct {
	// Allowed value: The `externalRef` field of a `CertificateManagerCertificate` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

var CertificateManagerCertificateGVK = schema.GroupVersionKind{
	Group:   "certificatemanager.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "CertificateManagerCertificate",
}

func (r *CertificateManagerCertificateRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on CertificateManagerCertificate reference")
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(CertificateManagerCertificateGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced CertificateManagerCertificate %s: %w", key, err)
	}

	// CertificateManagerCertificate is an unmigrated Terraform-based CRD and does not expose status.externalRef.
	// We must manually construct the external reference.
	resourceID, err := refsv1beta1.GetResourceID(u)
	if err != nil {
		return "", err
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return "", err
	}

	location, err := refsv1beta1.GetLocation(u)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("//certificatemanager.googleapis.com/projects/%s/locations/%s/certificates/%s", projectID, location, resourceID), nil
}

// A reference to a CertificateManagerCertificateMap resource.
type CertificateManagerCertificateMapRef struct {
	// Allowed value: string of the format `//certificatemanager.googleapis.com/projects/{{project}}/locations/global/certificateMaps/{{value}}`, where {{value}} is the `name` field of a `CertificateManagerCertificateMap` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

var CertificateManagerCertificateMapGVK = schema.GroupVersionKind{
	Group:   "certificatemanager.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "CertificateManagerCertificateMap",
}

func (r *CertificateManagerCertificateMapRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on CertificateManagerCertificateMap reference")
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(CertificateManagerCertificateMapGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced CertificateManagerCertificateMap %s: %w", key, err)
	}

	// CertificateManagerCertificateMap is an unmigrated Terraform-based CRD and does not expose status.externalRef.
	// We must manually construct the external reference.
	resourceID, err := refsv1beta1.GetResourceID(u)
	if err != nil {
		return "", err
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("//certificatemanager.googleapis.com/projects/%s/locations/global/certificateMaps/%s", projectID, resourceID), nil
}

// A reference to a NetworkSecurityServerTLSPolicy resource.
type NetworkSecurityServerTLSPolicyRef struct {
	// Allowed value: string of the format `projects/{{project}}/locations/{{location}}/serverTlsPolicies/{{value}}`, where {{value}} is the `name` field of a `NetworkSecurityServerTLSPolicy` resource.
	External string `json:"external,omitempty"`

	// Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name string `json:"name,omitempty"`

	// Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	Namespace string `json:"namespace,omitempty"`
}

var NetworkSecurityServerTLSPolicyGVK = schema.GroupVersionKind{
	Group:   "networksecurity.cnrm.cloud.google.com",
	Version: "v1beta1",
	Kind:    "NetworkSecurityServerTLSPolicy",
}

func (r *NetworkSecurityServerTLSPolicyRef) NormalizedExternal(ctx context.Context, reader client.Reader, otherNamespace string) (string, error) {
	if r.External != "" && r.Name != "" {
		return "", fmt.Errorf("cannot specify both name and external on NetworkSecurityServerTLSPolicy reference")
	}
	if r.External != "" {
		return r.External, nil
	}

	if r.Namespace == "" {
		r.Namespace = otherNamespace
	}
	key := types.NamespacedName{Name: r.Name, Namespace: r.Namespace}
	u := &unstructured.Unstructured{}
	u.SetGroupVersionKind(NetworkSecurityServerTLSPolicyGVK)
	if err := reader.Get(ctx, key, u); err != nil {
		if apierrors.IsNotFound(err) {
			return "", k8s.NewReferenceNotFoundError(u.GroupVersionKind(), key)
		}
		return "", fmt.Errorf("reading referenced NetworkSecurityServerTLSPolicy %s: %w", key, err)
	}

	// NetworkSecurityServerTLSPolicy is an unmigrated Terraform-based CRD and does not expose status.externalRef.
	// We must manually construct the external reference.
	resourceID, err := refsv1beta1.GetResourceID(u)
	if err != nil {
		return "", err
	}

	projectID, err := refsv1beta1.ResolveProjectID(ctx, reader, u)
	if err != nil {
		return "", err
	}

	location, err := refsv1beta1.GetLocation(u)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("projects/%s/locations/%s/serverTlsPolicies/%s", projectID, location, resourceID), nil
}
