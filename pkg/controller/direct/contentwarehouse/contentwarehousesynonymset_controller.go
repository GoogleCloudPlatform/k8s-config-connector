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

package contentwarehouse

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/contentwarehouse/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	pb "google.golang.org/genproto/googleapis/cloud/contentwarehouse/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ContentWarehouseSynonymSetGVK, newModel)
}

func newModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

var _ directbase.Model = &model{}

func (m *model) client(ctx context.Context) (pb.SynonymSetServiceClient, error) {
	var opts []option.ClientOption

	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}

	opts = append(opts, option.WithEndpoint("contentwarehouse.googleapis.com:443"))

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("dialing contentwarehouse service: %w", err)
	}

	return pb.NewSynonymSetServiceClient(conn), nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	obj := &krm.ContentWarehouseSynonymSet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := ContentWarehouseSynonymSetSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	cid := id.(*krm.ContentWarehouseSynonymSetIdentity)
	desired.Name = cid.String()
	desired.Context = cid.Context

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient: gcpClient,
		id:        cid,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type adapter struct {
	gcpClient pb.SynonymSetServiceClient
	id        *krm.ContentWarehouseSynonymSetIdentity
	desired   *pb.SynonymSet
	actual    *pb.SynonymSet
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting contentwarehouse synonym set", "name", a.id.String())

	req := &pb.GetSynonymSetRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetSynonymSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting contentwarehouse synonym set %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating contentwarehouse synonym set", "name", a.id.String())

	parent := a.id.ParentString()
	req := &pb.CreateSynonymSetRequest{
		Parent:     parent,
		SynonymSet: a.desired,
	}
	actual, err := a.gcpClient.CreateSynonymSet(ctx, req)
	if err != nil {
		return fmt.Errorf("creating contentwarehouse synonym set %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created contentwarehouse synonym set", "name", a.id.String())

	return a.updateStatus(ctx, createOp, actual)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating contentwarehouse synonym set", "name", a.id.String())

	diffs, _, err := compareSynonymSet(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return nil
	}

	log.V(2).Info("fields need update", "name", a.id.String())

	req := &pb.UpdateSynonymSetRequest{
		Name:       a.id.String(),
		SynonymSet: a.desired,
	}

	actual, err := a.gcpClient.UpdateSynonymSet(ctx, req)
	if err != nil {
		return fmt.Errorf("updating contentwarehouse synonym set %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated contentwarehouse synonym set", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, actual)
}

func compareSynonymSet(ctx context.Context, actual, desired *pb.SynonymSet) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := ContentWarehouseSynonymSetSpec_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := ContentWarehouseSynonymSetSpec_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name
	maskedActual.Context = desired.Context

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, actual *pb.SynonymSet) error {
	mapCtx := &direct.MapContext{}
	status := &krm.ContentWarehouseSynonymSetStatus{}
	status.ObservedState = ContentWarehouseSynonymSetObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting contentwarehouse synonym set", "name", a.id.String())

	req := &pb.DeleteSynonymSetRequest{Name: a.id.String()}
	_, err := a.gcpClient.DeleteSynonymSet(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting contentwarehouse synonym set %s: %w", a.id.String(), err)
	}

	return true, nil
}
