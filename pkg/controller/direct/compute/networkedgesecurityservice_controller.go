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
	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sort"
)

func init() {
	registry.RegisterModel(krm.ComputeNetworkEdgeSecurityServiceGVK, NewNetworkEdgeSecurityServiceModel)
}

func NewNetworkEdgeSecurityServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNetworkEdgeSecurityService{config: config}, nil
}

var _ directbase.Model = &modelNetworkEdgeSecurityService{}

type modelNetworkEdgeSecurityService struct {
	config *config.ControllerConfig
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

	// resolve securityPolicyRef
	securityPolicyRef := obj.Spec.SecurityPolicyRef
	if securityPolicyRef != nil {
		external, err := securityPolicyRef.NormalizedExternal(ctx, reader, obj.Namespace)
		if err != nil {
			return nil, err
		}
		obj.Spec.SecurityPolicyRef.External = external
	}

	// Get compute GCP client
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}

	networkEdgeSecurityServicesClient, err := gcpClient.newNetworkEdgeSecurityServicesClient(ctx)
	if err != nil {
		return nil, err
	}

	return &NetworkEdgeSecurityServiceAdapter{
		gcpClient: networkEdgeSecurityServicesClient,
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
		Project:                    a.id.Parent().ProjectID,
		Region:                     a.id.Parent().Location,
		NetworkEdgeSecurityService: a.id.ID(),
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeNetworkEdgeSecurityService %q: %w", a.id, err)
	}

	a.actual = actual
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
	resource.Name = direct.LazyPtr(a.id.ID())

	req := &computepb.InsertNetworkEdgeSecurityServiceRequest{
		Project:                            a.id.Parent().ProjectID,
		Region:                             a.id.Parent().Location,
		NetworkEdgeSecurityServiceResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeNetworkEdgeSecurityService %s waiting creation: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created ComputeNetworkEdgeSecurityService", "name", a.id)

	// Get the created resource
	created := &computepb.NetworkEdgeSecurityService{}
	getReq := &computepb.GetNetworkEdgeSecurityServiceRequest{
		Project:                    a.id.Parent().ProjectID,
		Region:                     a.id.Parent().Location,
		NetworkEdgeSecurityService: a.id.ID(),
	}
	created, err = a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting compute ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}
	status := &krm.ComputeNetworkEdgeSecurityServiceStatus{}
	status.ObservedState = ComputeNetworkEdgeSecurityServiceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
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
	resource.Name = direct.LazyPtr(a.id.ID())
	// An up-to-date fingerprint must be provided in order to patch
	resource.Fingerprint = a.actual.Fingerprint

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.ComputeNetworkEdgeSecurityServiceStatus{}
		status.ObservedState = ComputeNetworkEdgeSecurityServiceObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	// updateMask is a comma-separated list of fully qualified names of fields.
	var stringSlice []string
	for path := range paths {
		stringSlice = append(stringSlice, path)
	}

	sort.Strings(stringSlice)
	//updateMask := strings.Join(stringSlice, ",")

	req := &computepb.PatchNetworkEdgeSecurityServiceRequest{
		Project:                            a.id.Parent().ProjectID,
		Region:                             a.id.Parent().Location,
		NetworkEdgeSecurityService:         a.id.ID(),
		NetworkEdgeSecurityServiceResource: resource,
		UpdateMask:                         direct.LazyPtr("description,security_policy"),
	}
	op, err := a.gcpClient.Patch(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute ComputeNetworkEdgeSecurityService %s waiting for update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated ComputeNetworkEdgeSecurityService", "name", a.id)

	// Get the updated resource
	updated := &computepb.NetworkEdgeSecurityService{}
	getReq := &computepb.GetNetworkEdgeSecurityServiceRequest{
		Project:                    a.id.Parent().ProjectID,
		Region:                     a.id.Parent().Location,
		NetworkEdgeSecurityService: a.id.ID(),
	}
	updated, err = a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	// Fetch the updated resource to get the latest state including new fingerprint.
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("fetching updated ComputeNetworkEdgeSecurityService %s: %w", a.id, err)
	}

	status := &krm.ComputeNetworkEdgeSecurityServiceStatus{}
	status.ObservedState = ComputeNetworkEdgeSecurityServiceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
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

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location // Region is the location for this resource.
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.GetName())
	u.SetGroupVersionKind(krm.ComputeNetworkEdgeSecurityServiceGVK)
	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *NetworkEdgeSecurityServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeNetworkEdgeSecurityService", "name", a.id)

	req := &computepb.DeleteNetworkEdgeSecurityServiceRequest{
		Project:                    a.id.Parent().ProjectID,
		Region:                     a.id.Parent().Location,
		NetworkEdgeSecurityService: a.id.ID(),
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting ComputeNetworkEdgeSecurityService %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted ComputeNetworkEdgeSecurityService", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of ComputeNetworkEdgeSecurityService %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
