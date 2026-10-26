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
	"sigs.k8s.io/controller-runtime/pkg/client"

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
	registry.RegisterModel(krm.CloudNumberRegistryIpamAdminScopeGVK, NewIpamAdminScopeModel)
}

func NewIpamAdminScopeModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &ipamAdminScopeModel{config: *config}, nil
}

var _ directbase.Model = &ipamAdminScopeModel{}

type ipamAdminScopeModel struct {
	config config.ControllerConfig
}

func (m *ipamAdminScopeModel) client(ctx context.Context) (pb.CloudNumberRegistryClient, longrunningpb.OperationsClient, error) {
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

func (m *ipamAdminScopeModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudNumberRegistryIpamAdminScope{}
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
	desired := CloudNumberRegistryIpamAdminScopeSpec_ToProto(mapCtx, &obj.Spec)
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}

	gcpClient, operationsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ipamAdminScopeAdapter{
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		id:               id.(*krm.CloudNumberRegistryIpamAdminScopeIdentity),
		desired:          desired,
		reader:           reader,
		namespace:        obj.Namespace,
	}, nil
}

func (m *ipamAdminScopeModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ipamAdminScopeAdapter struct {
	gcpClient        pb.CloudNumberRegistryClient
	operationsClient longrunningpb.OperationsClient
	id               *krm.CloudNumberRegistryIpamAdminScopeIdentity
	desired          *pb.IpamAdminScope
	actual           *pb.IpamAdminScope

	reader    client.Reader
	namespace string
}

var _ directbase.Adapter = &ipamAdminScopeAdapter{}

func (a *ipamAdminScopeAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.IpamAdminScope == "" {
		return false, nil
	}

	req := &pb.GetIpamAdminScopeRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetIpamAdminScope(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting IpamAdminScope %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *ipamAdminScopeAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating IpamAdminScope", "name", a.id)

	req := &pb.CreateIpamAdminScopeRequest{
		Parent:           a.id.ParentString(),
		IpamAdminScope:   a.desired,
		IpamAdminScopeId: a.id.IpamAdminScope,
	}
	op, err := a.gcpClient.CreateIpamAdminScope(ctx, req)
	if err != nil {
		return fmt.Errorf("creating IpamAdminScope %s: %w", a.id.String(), err)
	}

	createdOp, err := a.waitForLRO(ctx, op)
	if err != nil {
		return fmt.Errorf("waiting for IpamAdminScope %s creation: %w", a.id.String(), err)
	}

	log.Info("successfully created IpamAdminScope in gcp", "name", a.id)

	var created pb.IpamAdminScope
	if err := createdOp.GetResponse().UnmarshalTo(&created); err != nil {
		return fmt.Errorf("unmarshalling IpamAdminScope from LRO response: %w", err)
	}

	return a.updateStatus(ctx, createOp, &created)
}

func (a *ipamAdminScopeAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("updating IpamAdminScope", "name", a.id)

	desired := proto.Clone(a.desired).(*pb.IpamAdminScope)
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

	req := &pb.UpdateIpamAdminScopeRequest{
		IpamAdminScope: desired,
		UpdateMask:     updateMask,
	}
	op, err := a.gcpClient.UpdateIpamAdminScope(ctx, req)
	if err != nil {
		return fmt.Errorf("updating IpamAdminScope %s: %w", a.id.String(), err)
	}

	updatedOp, err := a.waitForLRO(ctx, op)
	if err != nil {
		return fmt.Errorf("waiting for IpamAdminScope %s update: %w", a.id.String(), err)
	}

	log.Info("successfully updated IpamAdminScope", "name", a.id)

	var updated pb.IpamAdminScope
	if err := updatedOp.GetResponse().UnmarshalTo(&updated); err != nil {
		return fmt.Errorf("unmarshalling IpamAdminScope from LRO response: %w", err)
	}

	return a.updateStatus(ctx, updateOp, &updated)
}

func (a *ipamAdminScopeAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting IpamAdminScope", "name", a.id)

	req := &pb.DeleteIpamAdminScopeRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteIpamAdminScope(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting IpamAdminScope %s: %w", a.id.String(), err)
	}

	_, err = a.waitForLRO(ctx, op)
	if err != nil {
		return false, fmt.Errorf("waiting for IpamAdminScope %s deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *ipamAdminScopeAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudNumberRegistryIpamAdminScope{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudNumberRegistryIpamAdminScopeSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.IpamAdminScope)
	u.SetGroupVersionKind(krm.CloudNumberRegistryIpamAdminScopeGVK)

	u.Object = uObj
	return u, nil
}

func (a *ipamAdminScopeAdapter) waitForLRO(ctx context.Context, op *longrunningpb.Operation) (*longrunningpb.Operation, error) {
	pollInterval := 1 * time.Second
	return common.WaitForOperation(ctx, pollInterval, func(op *longrunningpb.Operation) (bool, error) {
		if op.Done {
			if op.GetError() != nil {
				return true, fmt.Errorf("operation failed: %s", op.GetError().GetMessage())
			}
			return true, nil
		}
		return false, nil
	}, func() (*longrunningpb.Operation, error) {
		return a.operationsClient.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.Name})
	})
}

func (a *ipamAdminScopeAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.IpamAdminScope) error {
	status := &krm.CloudNumberRegistryIpamAdminScopeStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudNumberRegistryIpamAdminScopeObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ipamAdminScopeAdapter) compare(ctx context.Context, actual, desired *pb.IpamAdminScope) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	clonedActual := proto.Clone(actual).(*pb.IpamAdminScope)
	clonedDesired := proto.Clone(desired).(*pb.IpamAdminScope)

	maskedActual, err := mappers.OnlySpecFields(clonedActual, CloudNumberRegistryIpamAdminScopeSpec_FromProto, CloudNumberRegistryIpamAdminScopeSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = clonedDesired.Name

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
