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

	gcp "cloud.google.com/go/vmwareengine/apiv1"
	pb "cloud.google.com/go/vmwareengine/apiv1/vmwareenginepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vmwareengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
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

	// Normalize references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
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

	mapCtx := &direct.MapContext{}
	desired := VMwareEnginePrivateConnectionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &privateConnectionAdapter{
		gcpClient: client,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *privateConnectionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type privateConnectionAdapter struct {
	gcpClient *gcp.Client
	id        *krm.PrivateConnectionIdentity
	desired   *pb.PrivateConnection
	actual    *pb.PrivateConnection
}

var _ directbase.Adapter = &privateConnectionAdapter{}

func (a *privateConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting vmwareengine privateconnection", "name", a.id)

	req := &pb.GetPrivateConnectionRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetPrivateConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting vmwareengine privateconnection %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *privateConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating vmwareengine privateconnection", "name", a.id)

	req := &pb.CreatePrivateConnectionRequest{
		Parent:              a.id.Parent().String(),
		PrivateConnectionId: a.id.ID(),
		PrivateConnection:   a.desired,
	}
	op, err := a.gcpClient.CreatePrivateConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating vmwareengine privateconnection %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vmwareengine privateconnection %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created vmwareengine privateconnection in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *privateConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating vmwareengine privateconnection", "name", a.id)

	diffs, updateMask, err := comparePrivateConnection(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		a.desired.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdatePrivateConnectionRequest{
			PrivateConnection: a.desired,
			UpdateMask:        updateMask,
		}
		op, err := a.gcpClient.UpdatePrivateConnection(ctx, req)
		if err != nil {
			return fmt.Errorf("updating vmwareengine privateconnection %s: %w", a.id.String(), err)
		}
		latest, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("vmwareengine privateconnection %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated vmwareengine privateconnection", "name", a.id)
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func comparePrivateConnection(ctx context.Context, actual, desired *pb.PrivateConnection) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VMwareEnginePrivateConnectionSpec_FromProto, VMwareEnginePrivateConnectionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *privateConnectionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.PrivateConnection) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VMwareEnginePrivateConnectionStatus{}
	status.ObservedState = VMwareEnginePrivateConnectionObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
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
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.VMwareEnginePrivateConnectionGVK)
	u.Object = uObj
	return u, nil
}

func (a *privateConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting vmwareengine privateconnection", "name", a.id)

	req := &pb.DeletePrivateConnectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeletePrivateConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting vmwareengine privateconnection %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted vmwareengine privateconnection", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete vmwareengine privateconnection %s: %w", a.id, err)
	}
	return true, nil
}
