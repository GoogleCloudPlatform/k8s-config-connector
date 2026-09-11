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

package cloudnumberregistry

import (
	"context"
	"fmt"
	"time"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudnumberregistry/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/cloudnumberregistry/pb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CloudNumberRegistryRegistryBookGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (pb.CloudNumberRegistryClient, longrunningpb.OperationsClient, error) {
	var opts []option.ClientOption

	config := m.config
	opts, err := config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}

	opts = append(opts, option.WithEndpoint("cloudnumberregistry.googleapis.com:443"))
	opts = append(opts, option.WithScopes("https://www.googleapis.com/auth/cloud-platform"))

	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("dialing cloudnumberregistry service: %w", err)
	}

	return pb.NewCloudNumberRegistryClient(conn), longrunningpb.NewOperationsClient(conn), nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudNumberRegistryRegistryBook{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, err
	}
	mapCtx := &direct.MapContext{}
	desired := CloudNumberRegistryRegistryBookSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, operationsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &adapter{
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		id:               id.(*krm.CloudNumberRegistryRegistryBookIdentity),
		desired:          desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type adapter struct {
	gcpClient        pb.CloudNumberRegistryClient
	operationsClient longrunningpb.OperationsClient
	id               *krm.CloudNumberRegistryRegistryBookIdentity
	desired          *pb.RegistryBook
	actual           *pb.RegistryBook
}

var _ directbase.Adapter = &adapter{}

func (a *adapter) Find(ctx context.Context) (bool, error) {
	if a.id.RegistryBook == "" {
		return false, nil
	}

	req := &pb.GetRegistryBookRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetRegistryBook(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting RegistryBook %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating RegistryBook", "name", a.id)

	req := &pb.CreateRegistryBookRequest{
		Parent:         a.id.ParentString(),
		RegistryBook:   a.desired,
		RegistryBookId: a.id.RegistryBook,
	}
	op, err := a.gcpClient.CreateRegistryBook(ctx, req)
	if err != nil {
		return fmt.Errorf("creating RegistryBook %s: %w", a.id.String(), err)
	}

	createdOp, err := a.waitForLRO(ctx, op)
	if err != nil {
		return fmt.Errorf("waiting for RegistryBook %s creation: %w", a.id.String(), err)
	}

	log.Info("successfully created RegistryBook in gcp", "name", a.id)

	var created pb.RegistryBook
	if err := createdOp.GetResponse().UnmarshalTo(&created); err != nil {
		return fmt.Errorf("unmarshalling RegistryBook from LRO response: %w", err)
	}

	return a.updateStatus(ctx, createOp, &created)
}

func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating RegistryBook", "name", a.id)

	desired := proto.Clone(a.desired).(*pb.RegistryBook)
	desired.Name = a.id.String()

	diffs, updateMask, err := a.compare(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.Info("no diff detected, skipping update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateRegistryBookRequest{
		RegistryBook: desired,
		UpdateMask:   updateMask,
	}
	op, err := a.gcpClient.UpdateRegistryBook(ctx, req)
	if err != nil {
		return fmt.Errorf("updating RegistryBook %s: %w", a.id.String(), err)
	}

	updatedOp, err := a.waitForLRO(ctx, op)
	if err != nil {
		return fmt.Errorf("waiting for RegistryBook %s update: %w", a.id.String(), err)
	}

	log.Info("successfully updated RegistryBook", "name", a.id)

	var updated pb.RegistryBook
	if err := updatedOp.GetResponse().UnmarshalTo(&updated); err != nil {
		return fmt.Errorf("unmarshalling RegistryBook from LRO response: %w", err)
	}

	return a.updateStatus(ctx, updateOp, &updated)
}

func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting RegistryBook", "name", a.id)

	req := &pb.DeleteRegistryBookRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteRegistryBook(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting RegistryBook %s: %w", a.id.String(), err)
	}

	_, err = a.waitForLRO(ctx, op)
	if err != nil {
		return false, fmt.Errorf("waiting for RegistryBook %s deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudNumberRegistryRegistryBook{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudNumberRegistryRegistryBookSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.RegistryBook)
	u.SetGroupVersionKind(krm.CloudNumberRegistryRegistryBookGVK)

	u.Object = uObj
	return u, nil
}

func (a *adapter) waitForLRO(ctx context.Context, op *longrunningpb.Operation) (*longrunningpb.Operation, error) {
	pollInterval := 1 * time.Second
	return common.WaitForOperation(ctx, pollInterval, func(op *longrunningpb.Operation) (bool, error) {
		return op.Done, nil
	}, func() (*longrunningpb.Operation, error) {
		return a.operationsClient.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.Name})
	})
}

func (a *adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.RegistryBook) error {
	status := &krm.CloudNumberRegistryRegistryBookStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudNumberRegistryRegistryBookObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *adapter) compare(ctx context.Context, actual, desired *pb.RegistryBook) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CloudNumberRegistryRegistryBookSpec_FromProto, CloudNumberRegistryRegistryBookSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
