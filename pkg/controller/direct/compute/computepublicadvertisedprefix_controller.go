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

package compute

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComputePublicAdvertisedPrefixGVK, NewComputePublicAdvertisedPrefixModel)
}

func NewComputePublicAdvertisedPrefixModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computePublicAdvertisedPrefixModel{config: config}, nil
}

var _ directbase.Model = &computePublicAdvertisedPrefixModel{}

type computePublicAdvertisedPrefixModel struct {
	config *config.ControllerConfig
}

func (m *computePublicAdvertisedPrefixModel) client(ctx context.Context) (*gcp.PublicAdvertisedPrefixesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewPublicAdvertisedPrefixesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PublicAdvertisedPrefixes client: %w", err)
	}
	return gcpClient, err
}

func (m *computePublicAdvertisedPrefixModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputePublicAdvertisedPrefix{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to ComputePublicAdvertisedPrefix: %w", err)
	}

	idValue, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idValue.(*krm.PublicAdvertisedPrefixIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ComputePublicAdvertisedPrefixAdapter{
		id:        id,
		projectId: id.ParentID.ProjectID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *computePublicAdvertisedPrefixModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support discovery
	return nil, nil
}

type ComputePublicAdvertisedPrefixAdapter struct {
	id        *krm.PublicAdvertisedPrefixIdentity
	projectId string
	gcpClient *gcp.PublicAdvertisedPrefixesClient
	desired   *krm.ComputePublicAdvertisedPrefix
	actual    *computepb.PublicAdvertisedPrefix
}

var _ directbase.Adapter = &ComputePublicAdvertisedPrefixAdapter{}

// Find retrieves the GCP resource.
// Return true, nil if the resource is found.
// Return false, nil if the resource is not found.
// Return false, err if an error occurred.
func (a *ComputePublicAdvertisedPrefixAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputePublicAdvertisedPrefix", "name", a.id.ResourceID)

	req := &computepb.GetPublicAdvertisedPrefixeRequest{
		Project:                a.projectId,
		PublicAdvertisedPrefix: a.id.ResourceID,
	}
	publicAdvertisedPrefix, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputePublicAdvertisedPrefix %q: %w", a.id.ResourceID, err)
	}

	a.actual = publicAdvertisedPrefix
	return true, nil
}

// Create creates the resource in GCP based on `a.desired`.
func (a *ComputePublicAdvertisedPrefixAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputePublicAdvertisedPrefix", "name", a.id.ResourceID)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputePublicAdvertisedPrefixSpecToAPI(mapCtx, &desired.Spec)
	resource.Name = direct.LazyPtr(a.id.ResourceID)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &computepb.InsertPublicAdvertisedPrefixeRequest{
		Project:                        a.projectId,
		PublicAdvertisedPrefixResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputePublicAdvertisedPrefix %q: %w", a.id.ResourceID, err)
	}
	if err := op.Wait(ctx); err != nil {
		return fmt.Errorf("waiting for creation of ComputePublicAdvertisedPrefix %q: %w", a.id.ResourceID, err)
	}
	log.V(2).Info("successfully created ComputePublicAdvertisedPrefix", "name", a.id.ResourceID)

	u := createOp.GetUnstructured()
	status := &krm.ComputePublicAdvertisedPrefixStatus{}
	if err := computePublicAdvertisedPrefixStatusToKRM(resource, status); err != nil {
		return fmt.Errorf("converting status to KRM: %w", err)
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return setPublicAdvertisedPrefixStatus(u, status)
}

// Update updates the resource in GCP based on `a.desired`.
func (a *ComputePublicAdvertisedPrefixAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputePublicAdvertisedPrefix", "name", a.id.ResourceID)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := ComputePublicAdvertisedPrefixSpecToAPI(mapCtx, &desired.Spec)
	resource.Name = direct.LazyPtr(a.id.ResourceID)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if !reflect.DeepEqual(resource.Description, a.actual.Description) {
		req := &computepb.PatchPublicAdvertisedPrefixeRequest{
			Project:                        a.projectId,
			PublicAdvertisedPrefix:         a.id.ResourceID,
			PublicAdvertisedPrefixResource: resource,
		}

		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ComputePublicAdvertisedPrefix %q: %w", a.id.ResourceID, err)
		}
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for update of ComputePublicAdvertisedPrefix %q: %w", a.id.ResourceID, err)
		}
		log.V(2).Info("successfully updated ComputePublicAdvertisedPrefix", "name", a.id.ResourceID)
	}

	// Fetch actual again if needed, this ensures the object reflects reality.
	// But since we just patched, `resource` is relatively accurate for the `setPublicAdvertisedPrefixStatus`.

	u := updateOp.GetUnstructured()
	status := &krm.ComputePublicAdvertisedPrefixStatus{}
	if err := computePublicAdvertisedPrefixStatusToKRM(resource, status); err != nil {
		return fmt.Errorf("converting status to KRM: %w", err)
	}
	return setPublicAdvertisedPrefixStatus(u, status)
}

// Export mimics the Find operation but creates an exact declarative representation
// of the GCP resource.
func (a *ComputePublicAdvertisedPrefixAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	var obj krm.ComputePublicAdvertisedPrefix
	obj.TypeMeta.Kind = krm.ComputePublicAdvertisedPrefixGVK.Kind
	obj.TypeMeta.APIVersion = krm.ComputePublicAdvertisedPrefixGVK.GroupVersion().String()
	obj.ObjectMeta.Name = a.id.ResourceID

	mapCtx := &direct.MapContext{}
	ComputePublicAdvertisedPrefixSpecFromAPI(mapCtx, a.actual, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: m}, nil
}

// Delete gracefully deletes the resource from GCP.
func (a *ComputePublicAdvertisedPrefixAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputePublicAdvertisedPrefix", "name", a.id.ResourceID)

	req := &computepb.DeletePublicAdvertisedPrefixeRequest{
		Project:                a.projectId,
		PublicAdvertisedPrefix: a.id.ResourceID,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for ComputePublicAdvertisedPrefix, not found", "name", a.id.ResourceID)
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputePublicAdvertisedPrefix %q: %w", a.id.ResourceID, err)
	}
	log.V(2).Info("successfully deleted ComputePublicAdvertisedPrefix", "name", a.id.ResourceID)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete ComputePublicAdvertisedPrefix %q: %w", a.id.ResourceID, err)
	}
	return true, nil
}

func computePublicAdvertisedPrefixStatusToKRM(in *computepb.PublicAdvertisedPrefix, out *krm.ComputePublicAdvertisedPrefixStatus) error {
	out.ObservedState = &krm.ComputePublicAdvertisedPrefixObservedState{}
	if in.Fingerprint != nil {
		out.ObservedState.Fingerprint = direct.LazyPtr(*in.Fingerprint)
	}
	if in.SharedSecret != nil {
		out.ObservedState.SharedSecret = in.SharedSecret
	}
	if in.Status != nil {
		out.ObservedState.Status = in.Status
	}
	if in.SelfLink != nil {
		out.ObservedState.SelfLink = in.SelfLink
	}
	return nil
}

func ComputePublicAdvertisedPrefixSpecToAPI(ctx *direct.MapContext, in *krm.ComputePublicAdvertisedPrefixSpec) *computepb.PublicAdvertisedPrefix {
	if in == nil {
		return nil
	}
	out := &computepb.PublicAdvertisedPrefix{}
	out.Description = in.Description
	out.DnsVerificationIp = in.DNSVerificationIP
	out.IpCidrRange = in.IPCidrRange
	out.PdpScope = in.PdpScope
	return out
}

func ComputePublicAdvertisedPrefixSpecFromAPI(ctx *direct.MapContext, in *computepb.PublicAdvertisedPrefix, out *krm.ComputePublicAdvertisedPrefixSpec) {
	if in == nil {
		return
	}
	out.Description = in.Description
	out.DNSVerificationIP = in.DnsVerificationIp
	out.IPCidrRange = in.IpCidrRange
	out.PdpScope = in.PdpScope
}

func setPublicAdvertisedPrefixStatus(u *unstructured.Unstructured, status interface{}) error {
	statusMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(status)
	if err != nil {
		return fmt.Errorf("converting status to unstructured: %w", err)
	}
	return unstructured.SetNestedField(u.Object, statusMap, "status")
}
