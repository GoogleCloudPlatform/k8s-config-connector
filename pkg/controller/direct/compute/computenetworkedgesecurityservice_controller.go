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

// +tool:controller
// proto.service: google.cloud.compute.v1.NetworkEdgeSecurityServices
// proto.message: google.cloud.compute.v1.NetworkEdgeSecurityService
// crd.type: ComputeNetworkEdgeSecurityService
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"
	"reflect"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeNetworkEdgeSecurityServiceGVK, NewNetworkEdgeSecurityServiceModel)
}

func NewNetworkEdgeSecurityServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNetworkEdgeSecurityService{config: *config}, nil
}

var _ directbase.Model = &modelNetworkEdgeSecurityService{}

type modelNetworkEdgeSecurityService struct {
	config config.ControllerConfig
}

func (m *modelNetworkEdgeSecurityService) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeNetworkEdgeSecurityService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNetworkEdgeSecurityServiceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// normalize reference fields
	if obj.Spec.SecurityPolicyRef != nil {
		if _, err := obj.Spec.SecurityPolicyRef.NormalizedExternal(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, err
		}
	}

	// Get compute GCP client
	gcpClient, err := newComputeRESTClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}

	computeClient := gcpClient.NetworkEdgeSecurityServices()

	return &NetworkEdgeSecurityServiceAdapter{
		gcpClient: computeClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelNetworkEdgeSecurityService) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NetworkEdgeSecurityServiceAdapter struct {
	gcpClient *compute.NetworkEdgeSecurityServicesClient
	id        *krm.NetworkEdgeSecurityServiceIdentity
	desired   *krm.ComputeNetworkEdgeSecurityService
	actual    *computepb.NetworkEdgeSecurityService
	reader    client.Reader
}

var _ directbase.Adapter = &NetworkEdgeSecurityServiceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *NetworkEdgeSecurityServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeNetworkEdgeSecurityService", "name", a.id)

	req := &computepb.GetNetworkEdgeSecurityServiceRequest{
		Project:                    a.id.ProjectID,
		Region:                     a.id.Location,
		NetworkEdgeSecurityService: a.id.Name,
	}
	networkedgesecurityservicepb, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeNetworkEdgeSecurityService %q: %w", a.id, err)
	}

	a.actual = networkedgesecurityservicepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NetworkEdgeSecurityServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeNetworkEdgeSecurityService", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputeNetworkEdgeSecurityServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = &a.id.Name

	req := &computepb.InsertNetworkEdgeSecurityServiceRequest{
		Project:                            a.id.ProjectID,
		Region:                             a.id.Location,
		NetworkEdgeSecurityServiceResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	// Wait for the regional operation to complete.
	if err = waitForRegionOp(ctx, op, a.id.ProjectID, a.id.Location, "create", a.gcpClient.RestBaseClient()); err != nil {
		return err
	}
	log.V(2).Info("successfully created ComputeNetworkEdgeSecurityService", "name", a.id)

	// Fetch the newly created resource to get the output-only fields.
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("fetching created ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	status := &krm.ComputeNetworkEdgeSecurityServiceStatus{}
	status.ObservedState = ComputeNetworkEdgeSecurityServiceObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	// Compute API does not return the resource body on create, only the operation.
	// We must fetch the created resource to get the status fields.
	// status.SelfLink = a.actual.SelfLink // Not present in ObservedState, but could be added if needed.
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NetworkEdgeSecurityServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeNetworkEdgeSecurityService", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputeNetworkEdgeSecurityServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	needsUpdate := false
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		needsUpdate = true
	}
	if desired.Spec.SecurityPolicyRef != nil {
		// Need to normalize the reference again as it might not be fully resolved in the desired object.
		spRefExt, err := desired.Spec.SecurityPolicyRef.NormalizedExternal(ctx, a.reader, a.desired.GetNamespace())
		if err != nil {
			return err
		}
		if !reflect.DeepEqual(&spRefExt, a.actual.SecurityPolicy) {
			needsUpdate = true
		}
	} else if a.actual.SecurityPolicy != nil { // Check if the field should be unset
		needsUpdate = true
	}

	if !needsUpdate {
		log.V(2).Info("no field needs update", "name", a.id)
		// Update status even if no fields changed.
		status := &krm.ComputeNetworkEdgeSecurityServiceStatus{}
		status.ObservedState = ComputeNetworkEdgeSecurityServiceObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.LazyPtr(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)

	}

	// Add fingerprint for optimistic locking
	resource.Fingerprint = a.actual.Fingerprint

	req := &computepb.PatchNetworkEdgeSecurityServiceRequest{
		Project:                            a.id.ProjectID,
		Region:                             a.id.Location,
		NetworkEdgeSecurityService:         a.id.Name,
		NetworkEdgeSecurityServiceResource: resource,
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	// Wait for the regional operation to complete.
	if err = waitForRegionOp(ctx, op, a.id.ProjectID, a.id.Location, "update", a.gcpClient.RestBaseClient()); err != nil {
		return err
	}
	log.V(2).Info("successfully updated ComputeNetworkEdgeSecurityService", "name", a.id)

	// Fetch the updated resource to get the latest state including new fingerprint.
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("fetching updated ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	status := &krm.ComputeNetworkEdgeSecurityServiceStatus{}
	status.ObservedState = ComputeNetworkEdgeSecurityServiceObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	// status.SelfLink = a.actual.SelfLink // Not present in ObservedState, but could be added if needed.
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *NetworkEdgeSecurityServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeNetworkEdgeSecurityService{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeNetworkEdgeSecurityServiceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Handle references
	if a.actual.GetSecurityPolicy() != "" {
		obj.Spec.SecurityPolicyRef = &refs.ComputeSecurityPolicyRef{
			External: a.actual.GetSecurityPolicy(),
		}
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.ProjectID}
	obj.Spec.Location = a.id.Location // Region is the location for this resource.
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeNetworkEdgeSecurityServiceGVK)
	// Retain KCC annotations/labels if desired.
	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *NetworkEdgeSecurityServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeNetworkEdgeSecurityService", "name", a.id)

	req := &computepb.DeleteNetworkEdgeSecurityServiceRequest{
		Project:                    a.id.ProjectID,
		Region:                     a.id.Location,
		NetworkEdgeSecurityService: a.id.Name,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent ComputeNetworkEdgeSecurityService, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}
	log.V(2).Info("successfully initiated deletion of ComputeNetworkEdgeSecurityService", "name", a.id)

	// Wait for the regional operation to complete.
	if err = waitForRegionOp(ctx, op, a.id.ProjectID, a.id.Location, "delete", a.gcpClient.RestBaseClient()); err != nil {
		return false, fmt.Errorf("waiting delete ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	return true, nil
}

// Helper function to wait for regional Compute operations
func waitForRegionOp(ctx context.Context, op *compute.Operation, project, region, operationDesc string, client *compute.BaseClient) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("Waiting for regional operation", "operation", op.Proto().GetName(), "region", region, "description", operationDesc)
	return op.Wait(ctx)
	// Note: The above `op.Wait()` uses the default polling interval/timeout.
	// For fine-grained control or use with a shared poller, you might need a custom implementation:
	/*
		_, err := direct.WaitForComputeRegionOperation(ctx, client, project, region, proto.String(op.Proto().GetName()))
		if err != nil {
			return fmt.Errorf("waiting for region operation %q to complete: %w", *op.Proto().Name, err)
		}
		log.V(2).Info("Regional operation completed", "operation", op.Proto().GetName(), "region", region, "description", operationDesc)
		return nil
	*/
}
