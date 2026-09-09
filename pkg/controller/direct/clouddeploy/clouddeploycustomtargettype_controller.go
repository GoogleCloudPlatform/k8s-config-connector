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

package clouddeploy

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/deploy/apiv1"
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CustomTargetTypeGVK, NewCustomTargetTypeModel)
}

func NewCustomTargetTypeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCustomTargetType{config: *config}, nil
}

type modelCustomTargetType struct {
	config config.ControllerConfig
}

func (m *modelCustomTargetType) client(ctx context.Context) (*gcp.CloudDeployClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudDeployRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CustomTargetType client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelCustomTargetType) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudDeployCustomTargetType{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := krm.NewCustomTargetTypeIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := CustomTargetTypeSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.String()

	return &CustomTargetTypeAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelCustomTargetType) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type CustomTargetTypeAdapter struct {
	id        *krm.CustomTargetTypeIdentity
	gcpClient *gcp.CloudDeployClient
	desired   *pb.CustomTargetType
	actual    *pb.CustomTargetType
}

func (a *CustomTargetTypeAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding CustomTargetType", "name", a.id.String())

	req := &pb.GetCustomTargetTypeRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetCustomTargetType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CustomTargetType %s: %w", a.id.String(), err)
	}
	a.actual = actual
	return true, nil
}

func (a *CustomTargetTypeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CustomTargetType", "name", a.id.String())

	req := &pb.CreateCustomTargetTypeRequest{
		Parent:             a.id.Parent().String(),
		CustomTargetType:   a.desired,
		CustomTargetTypeId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateCustomTargetType(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CustomTargetType %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("CustomTargetType %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created CustomTargetType", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *CustomTargetTypeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CustomTargetType", "name", a.id.String())

	diffs, updateMask, err := compareCustomTargetType(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	var updated *pb.CustomTargetType
	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)
		req := &pb.UpdateCustomTargetTypeRequest{
			UpdateMask:       updateMask,
			CustomTargetType: a.desired,
		}
		op, err := a.gcpClient.UpdateCustomTargetType(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CustomTargetType %s: %w", a.id.String(), err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("CustomTargetType %s waiting update: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated CustomTargetType", "name", a.id.String())
	} else {
		log.V(2).Info("no field needs update", "name", a.id.String())
		updated = a.actual
	}

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *CustomTargetTypeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CustomTargetType", "name", a.id.String())

	req := &pb.DeleteCustomTargetTypeRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteCustomTargetType(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting CustomTargetType %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete CustomTargetType %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted CustomTargetType", "name", a.id.String())
	return true, nil
}

func (a *CustomTargetTypeAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CustomTargetType) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CustomTargetTypeStatus{}
	status.ObservedState = CustomTargetTypeObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *CustomTargetTypeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudDeployCustomTargetType{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CustomTargetTypeSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("error converting to unstructured: %w", err)
	}

	u.Object = uObj
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.CustomTargetTypeGVK)

	return u, nil
}

func compareCustomTargetType(ctx context.Context, actual, desired *pb.CustomTargetType) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CustomTargetTypeSpec_v1alpha1_FromProto, CustomTargetTypeSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.CustomTargetType)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
