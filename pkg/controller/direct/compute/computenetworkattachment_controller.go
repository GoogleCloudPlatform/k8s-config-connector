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
// proto.service: google.cloud.compute.v1.NetworkAttachments
// proto.message: google.cloud.compute.v1.NetworkAttachment
// crd.type: ComputeNetworkAttachment
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/diff"
)

func init() {
	registry.RegisterModel(krm.ComputeNetworkAttachmentGVK, NewNetworkAttachmentModel)
}

func NewNetworkAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNetworkAttachment{config: *config}, nil
}

var _ directbase.Model = &modelNetworkAttachment{}

type modelNetworkAttachment struct {
	config config.ControllerConfig
}

func (m *modelNetworkAttachment) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeNetworkAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewComputeNetworkAttachmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get compute GCP client
	gcpClient, err := newGCPClient(ctx, &m.config, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	computeClient, err := gcpClient.newComputeV1Client(ctx)
	if err != nil {
		return nil, err
	}
	return &NetworkAttachmentAdapter{
		id:        id,
		gcpClient: computeClient,
		desired:   obj,
		reader:    reader,
		resourceOverrides: resourceoverrides.Get(
			v1beta1.ComputeNetworkAttachmentGVK.Kind,
			direct.OverridesKey,
			directbase.OverridesFieldManagerName,
			directbase.OverridesDeletionPolicy,
			directbase.OverridesDefaultLabels,
			directbase.OverridesProjectIDLabel,
		),
	}, nil
}

func (m *modelNetworkAttachment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NetworkAttachmentAdapter struct {
	id        *krm.ComputeNetworkAttachmentIdentity
	gcpClient *computeClient
	desired   *krm.ComputeNetworkAttachment
	actual    *computepb.NetworkAttachment
	reader    client.Reader
	resourceOverrides	resourceoverrides.ResourceOverrides
}

var _ directbase.Adapter = &NetworkAttachmentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *NetworkAttachmentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NetworkAttachment", "name", a.id)

	req := &computepb.GetNetworkAttachmentRequest{
		Project:           a.id.Project,
		Region:            a.id.Region,
		NetworkAttachment: a.id.Name,
	}
	networkattachmentpb, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkAttachment %q: %w", a.id, err)
	}

	a.actual = networkattachmentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NetworkAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkAttachment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputeNetworkAttachmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &computepb.InsertNetworkAttachmentRequest{
		Project:                a.id.Project,
		Region:                 a.id.Region,
		NetworkAttachmentResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkAttachment %s: %w", a.id, err)
	}

	// use default krm UpdateStatusHandler
	status := &krm.ComputeNetworkAttachmentStatus{}
	var pending *directbase.PendingStatus
	if op.Status != nil && op.GetStatus() != "DONE" {
		status.ObservedGeneration = &a.desired.Generation
		pending = &directbase.PendingStatus{
			OperationType: directbase.CreateOperation,
			Status: v1alpha1.Status{
				Conditions: []v1alpha1.Condition{{
					Type:    v1alpha1.ConditionType(krm.UpToDate),
					Status:  v1alpha1.ConditionFalse,
					Reason:  "Updating",
					Message: fmt.Sprintf("Operation %q is in progress", op.Name),
				}},
			},
		}
	}

	if err := createOp.UpdateStatus(ctx, status, pending); err != nil {
		return err
	}

	log.V(2).Info("successfully created NetworkAttachment", "name", a.id)
	return nil
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NetworkAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkAttachment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComputeNetworkAttachmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = &a.id.Name // we need to set the name so that GCP API can identify the resource

	paths, err := a.resourceOverrides.GetUpdateDiff(ctx, a.reader, a.desired, a.actual, desired.Spec.Subnetworks)
	if err != nil {
		return err
	}
	if len(paths) != 0 {
		return fmt.Errorf("updating NetworkAttachment is not supported, fields: %v", paths)
	}

	// still need to update status (in the event of acquiring an existing resource)
	status := &krm.ComputeNetworkAttachmentStatus{}
	status.ObservedState = ComputeNetworkAttachmentObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *NetworkAttachmentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeNetworkAttachment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeNetworkAttachmentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Region
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ComputeNetworkAttachmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *NetworkAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkAttachment", "name", a.id)

	req := &computepb.DeleteNetworkAttachmentRequest{
		Project:           a.id.Project,
		Region:            a.id.Region,
		NetworkAttachment: a.id.Name,
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent NetworkAttachment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting NetworkAttachment %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted NetworkAttachment", "name", a.id)
	return false, direct.CheckOperation(ctx, op, a.config.UserAgent)
}

// This function convert ComputeNetworkAttachment object to proto message for creating it.
// Note: this function updates the MapContext so it must not be used to convert the status.
func ComputeNetworkAttachmentSpec_ToProto(mapCtx *direct.MapContext, spec *krm.ComputeNetworkAttachmentSpec) (res *computepb.NetworkAttachment) {
	res = &computepb.NetworkAttachment{}

	res.Name = spec.Name

	if spec.Description != nil {
		res.Description = spec.Description
	}
	if spec.ConnectionPreference != nil {
		res.ConnectionPreference = spec.ConnectionPreference
	}
	if spec.Subnetworks != nil {
		res.Subnetworks = direct.StringSlicePtrVal(spec.Subnetworks)
	}
	if spec.ProducerAcceptLists != nil {
		res.ProducerAcceptLists = direct.StringSlicePtrVal(spec.ProducerAcceptLists)
	}
	if spec.ProducerRejectLists != nil {
		res.ProducerRejectLists = direct.StringSlicePtrVal(spec.ProducerRejectLists)
	}

	return
}

func ComputeNetworkAttachmentObservedState_FromProto(mapCtx *direct.MapContext, p *computepb.NetworkAttachment) (obs *krm.ComputeNetworkAttachmentStatus) {
	if p == nil {
		return nil
	}

	obs = &krm.ComputeNetworkAttachmentStatus{}

	if p.Fingerprint != nil {
		obs.Fingerprint = direct.PtrTo(p.GetFingerprint())
	}
	if p.Kind != nil {
		obs.Kind = direct.PtrTo(p.GetKind())
	}
	if p.Network != nil {
		obs.Network = direct.PtrTo(p.GetNetwork())
	}
	if p.SelfLink != nil {
		obs.SelfLink = direct.PtrTo(p.GetSelfLink())
	}

	if p.ConnectionEndpoints != nil {
		obs.ConnectionEndpoints = []*krm.NetworkAttachmentConnectionEndpoints{}
		for _, v := range p.GetConnectionEndpoints() {
			tmp := ComputeNetworkAttachmentConnectionEndpoints_FromProto(mapCtx, v)
			obs.ConnectionEndpoints = append(obs.ConnectionEndpoints, tmp)
		}
	}
	return
}

func ComputeNetworkAttachmentConnectionEndpoints_FromProto(mapCtx *direct.MapContext, p *computepb.NetworkAttachmentConnectedEndpoint) (res *krm.NetworkAttachmentConnectionEndpoints) {
	if p == nil {
		return
	}

	res = &krm.NetworkAttachmentConnectionEndpoints{}
	res.IpAddress = p.IpAddress
	res.IPv6Address = p.Ipv6Address
	res.ProjectIDOrNum = p.ProjectIdOrNum
	res.SecondaryIpCidrRange = p.SecondaryIpCidrRange
	res.Status = p.Status
	res.Subnetwork = p.Subnetwork

	return
}

```
</out>


