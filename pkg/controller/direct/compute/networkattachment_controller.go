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

// +tool:controller
// proto.service: google.cloud.compute.v1.NetworkAttachments
// proto.message: google.cloud.compute.v1.NetworkAttachment
// crd.type: ComputeNetworkAttachment
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.ComputeNetworkAttachmentGVK, NewNetworkAttachmentModel)
}

func NewNetworkAttachmentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &networkAttachmentModel{config: config}, nil
}

var _ directbase.Model = &networkAttachmentModel{}

type networkAttachmentModel struct {
	config *config.ControllerConfig
}

func (m *networkAttachmentModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeNetworkAttachment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNetworkAttachmentIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	networkAttachmentClient, err := gcpClient.newNetworkAttachmentsClient(ctx)
	if err != nil {
		return nil, err
	}

	return &NetworkAttachmentAdapter{
		gcpClient: networkAttachmentClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *networkAttachmentModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type NetworkAttachmentAdapter struct {
	gcpClient *compute.NetworkAttachmentsClient
	id        *v1alpha1.NetworkAttachmentIdentity
	desired   *krm.ComputeNetworkAttachment
	actual    *computepb.NetworkAttachment
	reader    client.Reader
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
		Project:           a.id.Parent().ProjectID,
		Region:            a.id.Parent().Location,
		NetworkAttachment: a.id.ID(),
	}
	actual, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NetworkAttachment %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NetworkAttachmentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NetworkAttachment", "name", a.id)
	mapCtx := &direct.MapContext{}

	err := a.resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	desired := a.desired.DeepCopy()
	resource := ComputeNetworkAttachmentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = direct.LazyPtr(a.id.ID())

	req := &computepb.InsertNetworkAttachmentRequest{
		Project:                   a.id.Parent().ProjectID,
		Region:                    a.id.Parent().Location,
		NetworkAttachmentResource: resource,
	}
	op, err := a.gcpClient.Insert(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NetworkAttachment %s: %w", a.id, err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("compute NetworkAttachment %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created compute NetworkAttachment in gcp", "name", a.id)

	// Get the created resource
	created, err := a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting NetworkAttachment %s: %w", a.id, err)
	}

	status := &krm.ComputeNetworkAttachmentStatus{}
	status.ObservedState = ComputeNetworkAttachmentObservedState_FromProto(mapCtx, created)
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NetworkAttachmentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkAttachment", "name", a.id)
	mapCtx := &direct.MapContext{}

	err := a.resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	desired := a.desired.DeepCopy()
	resource := ComputeNetworkAttachmentSpec_ToProto(mapCtx, &desired.Spec)
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

	var updated *computepb.NetworkAttachment
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {

		req := &computepb.PatchNetworkAttachmentRequest{
			Project:                   a.id.Parent().ProjectID,
			Region:                    a.id.Parent().Location,
			NetworkAttachment:         a.id.ID(),
			NetworkAttachmentResource: resource,
		}
		op, err := a.gcpClient.Patch(ctx, req)
		if err != nil {
			return fmt.Errorf("updating compute NetworkAttachment %s: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated compute NetworkAttachment", "name", a.id.String())

		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("compute NetworkAttachment %s waiting for update: %w", a.id.String(), err)
		}

		// Get the updated resource
		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting NetworkAttachment %s: %w", a.id, err)
		}
	}

	status := &krm.ComputeNetworkAttachmentStatus{}
	status.ObservedState = ComputeNetworkAttachmentObservedState_FromProto(mapCtx, updated)
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
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComputeNetworkAttachmentGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *NetworkAttachmentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NetworkAttachment", "name", a.id)

	req := &computepb.DeleteNetworkAttachmentRequest{
		Project:           a.id.Parent().ProjectID,
		Region:            a.id.Parent().Location,
		NetworkAttachment: a.id.ID(),
	}
	op, err := a.gcpClient.Delete(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting compute NetworkAttachment %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted compute NetworkAttachment", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of compute NetworkAttachment %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

func (a *NetworkAttachmentAdapter) get(ctx context.Context) (*computepb.NetworkAttachment, error) {
	getReq := &computepb.GetNetworkAttachmentRequest{
		Project:           a.id.Parent().ProjectID,
		Region:            a.id.Parent().Location,
		NetworkAttachment: a.id.ID(),
	}
	resource, err := a.gcpClient.Get(ctx, getReq)
	if err != nil {
		return nil, fmt.Errorf("getting ComputeNetworkAttachment %s: %w", a.id, err)
	}
	return resource, nil
}

func (a *NetworkAttachmentAdapter) resolveDependencies(ctx context.Context, reader client.Reader, obj *krm.ComputeNetworkAttachment) error {
	// resolve subnetwork
	if obj.Spec.SubnetworkRefs != nil {
		subnetworkRefs := obj.Spec.SubnetworkRefs
		for i, subnetworkRef := range subnetworkRefs {
			external, err := subnetworkRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
			if err != nil {
				return err
			}
			subnetworkRefs[i].External = external
		}
	}
	// resolve project
	if obj.Spec.ProducerRejectLists != nil {
		var projects []*refsv1beta1.ProjectRef
		for _, i := range obj.Spec.ProducerRejectLists {
			project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), i)
			if err != nil {
				return err
			}
			i.External = project.ProjectID
			projects = append(projects, i)
		}
		obj.Spec.ProducerRejectLists = projects
	}
	if obj.Spec.ProducerAcceptLists != nil {
		var projects []*refsv1beta1.ProjectRef
		for _, i := range obj.Spec.ProducerAcceptLists {
			project, err := refsv1beta1.ResolveProject(ctx, reader, obj.GetNamespace(), i)
			if err != nil {
				return err
			}
			i.External = project.ProjectID
			projects = append(projects, i)
		}
		obj.Spec.ProducerAcceptLists = projects
	}
	return nil
}
