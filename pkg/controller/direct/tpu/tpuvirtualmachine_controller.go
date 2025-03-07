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

package tpu

// +tool:controller
// proto.service: google.cloud.tpu.v2.Tpu
// proto.message: google.cloud.tpu.v2.Node
// crd.type: TPUVirtualMachine
// crd.version: v1alpha1

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/googleapis/gax-go/v2"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	pb "cloud.google.com/go/tpu/apiv2/tpupb"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tpu/v1alpha1"
	kcc "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
)

func init() {
	registry.RegisterModel(krm.TPUVirtualMachineGVK, NewTPUVirtualMachineModel)
}

func NewTPUVirtualMachineModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTPUVirtualMachine{config: config}, nil
}

var _ directbase.Model = &modelTPUVirtualMachine{}

type modelTPUVirtualMachine struct {
	config *config.ControllerConfig
}

func (m *modelTPUVirtualMachine) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.TPUVirtualMachine{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	i, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.TPUVirtualMachineIdentity)
	project := &refs.Project{ProjectID: id.ParentID.ProjectID}

	if err := common.NormalizeReferences(ctx, kube, obj, project); err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	tpuClient, err := gcpClient.newClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := TPUVirtualMachineSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &TPUVirtualMachineAdapter{
		id:        id,
		tpuClient: tpuClient,
		desired:   desiredProto,
	}, nil
}

// AdapterForURL returns an adapter for export of an object by URL
// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
func (m *modelTPUVirtualMachine) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //tpu.googleapis.com/projects/PROJECT_NUMBER/locations/LOCATION/nodes/NODE_ID
	if external, ok := strings.CutPrefix(url, "//tpu.googleapis.com/"); ok {
		id := &krm.TPUVirtualMachineIdentity{}
		if err := id.FromExternal(external); err == nil {
			gcpClient, err := newGCPClient(ctx, m.config)
			if err != nil {
				return nil, err
			}
			tpuClient, err := gcpClient.newClient(ctx)
			if err != nil {
				return nil, err
			}

			return &TPUVirtualMachineAdapter{
				id:        id,
				tpuClient: tpuClient,
			}, nil
		}
	}

	return nil, nil
}

type TPUVirtualMachineAdapter struct {
	id        *krm.TPUVirtualMachineIdentity
	tpuClient *TPUV2Client
	desired   *pb.Node
	actual    *pb.Node
}

var _ directbase.Adapter = &TPUVirtualMachineAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TPUVirtualMachineAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TPUVirtualMachine", "name", a.id)

	req := &pb.GetNodeRequest{
		Name: a.id.String(),
	}

	// Because we are using a hand-coded GRPC client, we need to add the headers per https://google.aip.dev/client-libraries/4222
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)

	obj, err := a.tpuClient.GetNode(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting tpu node %q: %w", a.id, err)
	}

	a.actual = obj
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TPUVirtualMachineAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TPUVirtualMachine", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := direct.ProtoClone(a.desired)

	req := &pb.CreateNodeRequest{
		Parent: a.id.ParentID.String(),
		Node:   desired,
		NodeId: a.id.Node,
	}

	{
		// Because we are using a hand-coded GRPC client, we need to add the headers per https://google.aip.dev/client-libraries/4222
		hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}
		ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)

		op, err := a.tpuClient.CreateNode(ctx, req)
		if err != nil {
			return fmt.Errorf("creating tpu vm node %s: %w", a.id, err)
		}
		if !op.GetDone() {
			if err := a.tpuClient.WaitForLRO(ctx, op); err != nil {
				return fmt.Errorf("waiting for creation of tpu vm node %q: %w", a.id, err)
			}
		}
	}

	// Because we are using a hand-coded GRPC client, we need to add the headers per https://google.aip.dev/client-libraries/4222
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(a.id.String()))}
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)

	created, err := a.tpuClient.GetNode(ctx, &pb.GetNodeRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting created tpu vm node %q: %w", a.id, err)
	}

	log.V(2).Info("successfully created tpu vm node", "name", a.id)

	status := &krm.TPUVirtualMachineStatus{}
	status.ObservedState = TPUVirtualMachineObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	var readyCondition *kcc.Condition
	switch created.GetState() {
	case pb.Node_CREATING:
		// We're not actually done yet!
		// Should we set readyCondition?
		createOp.RequestRequeue()
	}

	return createOp.UpdateStatus(ctx, status, readyCondition)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TPUVirtualMachineAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating tpu node", "name", a.id)
	mapCtx := &direct.MapContext{}
	updateMask := fieldmaskpb.FieldMask{}

	desired := direct.ProtoClone(a.desired)

	if desired.Description != a.actual.Description {
		updateMask.Paths = append(updateMask.Paths, "description")
	}

	latest := a.actual

	if len(updateMask.Paths) != 0 {
		{
			desired.Name = a.id.String()
			req := &pb.UpdateNodeRequest{
				UpdateMask: &updateMask,
				Node:       desired,
			}

			// Because we are using a hand-coded GRPC client, we need to add the headers per https://google.aip.dev/client-libraries/4222
			hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(a.id.String()))}
			ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)

			op, err := a.tpuClient.UpdateNode(ctx, req)
			if err != nil {
				return fmt.Errorf("updating tpu vm node %s: %w", a.id, err)
			}
			if !op.GetDone() {
				if err := a.tpuClient.WaitForLRO(ctx, op); err != nil {
					return fmt.Errorf("waiting for update of tpu vm node %q: %w", a.id, err)
				}
			}
		}

		{
			// Because we are using a hand-coded GRPC client, we need to add the headers per https://google.aip.dev/client-libraries/4222
			hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(a.id.String()))}
			ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
			updated, err := a.tpuClient.GetNode(ctx, &pb.GetNodeRequest{Name: a.id.String()})
			if err != nil {
				return fmt.Errorf("getting updated tpu vm node %q: %w", a.id, err)
			}
			log.V(2).Info("successfully updated tpu vm node", "name", a.id)
			latest = updated
		}
	} else {
		log.V(2).Info("no field needs update", "name", a.id)
		latest = a.actual
	}

	status := &krm.TPUVirtualMachineStatus{}
	status.ObservedState = TPUVirtualMachineObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TPUVirtualMachineAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.TPUVirtualMachine{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TPUVirtualMachineSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.ParentID.ProjectID}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj

	u.SetName(a.id.Node)
	u.SetGroupVersionKind(krm.TPUVirtualMachineGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TPUVirtualMachineAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting tpu node", "name", a.id)

	// Because we are using a hand-coded GRPC client, we need to add the headers per https://google.aip.dev/client-libraries/4222
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(a.id.String()))}
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)

	op, err := a.tpuClient.DeleteNode(ctx, &pb.DeleteNodeRequest{Name: a.id.String()})
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent tpu node, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting tpu node %q: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted tpu node", "name", a.id)

	if !op.GetDone() {
		if err := a.tpuClient.WaitForLRO(ctx, op); err != nil {
			return false, fmt.Errorf("waiting for deletion of tpu node %q: %w", a.id, err)
		}
	}
	return true, nil
}
