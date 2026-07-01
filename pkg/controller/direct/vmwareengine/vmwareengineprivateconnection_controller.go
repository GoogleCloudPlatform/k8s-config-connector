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

// +tool:controller
// proto.service: google.cloud.vmwareengine.v1.VmwareEngine
// proto.message: google.cloud.vmwareengine.v1.PrivateConnection
// crd.type: VMwareEnginePrivateConnection
// crd.version: v1alpha1

package vmwareengine

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/vmwareengine/apiv1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.VMwareEnginePrivateConnectionGVK, NewPrivateConnectionModel)
}

func NewPrivateConnectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &privateConnectionModel{config: *config}, nil
}

var _ directbase.Model = &privateConnectionModel{}

type privateConnectionModel struct {
	config config.ControllerConfig
}

func (m *privateConnectionModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VMwareEnginePrivateConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewPrivateConnectionIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get VMwareEngine GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newClient(ctx)
	if err != nil {
		return nil, err
	}
	return &privateConnectionAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *privateConnectionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type privateConnectionAdapter struct {
	gcpClient *gcp.Client
	id        *krm.PrivateConnectionIdentity
	desired   *krm.VMwareEnginePrivateConnection
	actual    *pb.PrivateConnection
	reader    client.Reader
}

var _ directbase.Adapter = &privateConnectionAdapter{}

func (a *privateConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting vmwareengine private connection", "name", a.id)

	req := &pb.GetPrivateConnectionRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetPrivateConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting vmwareengine private connection %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *privateConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating vmwareengine private connection", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEnginePrivateConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreatePrivateConnectionRequest{
		Parent:              a.id.Parent().String(),
		PrivateConnectionId: a.id.ID(),
		PrivateConnection:   resource,
	}
	op, err := a.gcpClient.CreatePrivateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating vmwareengine private connection %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vmwareengine private connection %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created vmwareengine private connection in gcp", "name", a.id)

	status := &krm.VMwareEnginePrivateConnectionStatus{}
	status.ObservedState = VMwareEnginePrivateConnectionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *privateConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating vmwareengine private connection", "name", a.id)

	if err := a.normalizeReferenceFields(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := VMwareEnginePrivateConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	paths := []string{}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		report.AddField("description", a.actual.Description, resource.Description)
		paths = append(paths, "description")
	}
	if desired.Spec.RoutingMode != nil && !reflect.DeepEqual(resource.RoutingMode, a.actual.RoutingMode) {
		report.AddField("routing_mode", a.actual.RoutingMode, resource.RoutingMode)
		paths = append(paths, "routing_mode")
	}

	var updated *pb.PrivateConnection
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdatePrivateConnectionRequest{
			PrivateConnection: resource,
			UpdateMask:        &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdatePrivateConnection(ctx, req)
		if err != nil {
			return fmt.Errorf("updating vmwareengine private connection %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("vmwareengine private connection %s waiting for update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated vmwareengine private connection", "name", a.id)
	}

	status := &krm.VMwareEnginePrivateConnectionStatus{}
	status.ObservedState = VMwareEnginePrivateConnectionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *privateConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VMwareEnginePrivateConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VMwareEnginePrivateConnectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VMwareEnginePrivateConnectionGVK)
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *privateConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting vmwareengine private connection", "name", a.id)

	req := &pb.DeletePrivateConnectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeletePrivateConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent PrivateConnection, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting vmwareengine private connection %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted vmwareengine private connection", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete PrivateConnection %s: %w", a.id, err)
	}
	return true, nil
}

func (a *privateConnectionAdapter) normalizeReferenceFields(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.VMwareEngineNetworkRef != nil {
		if _, err := obj.Spec.VMwareEngineNetworkRef.NormalizedExternal(ctx, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if obj.Spec.ServiceNetworkRef != nil {
		if err := obj.Spec.ServiceNetworkRef.Normalize(ctx, a.reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	return nil
}
