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

package interceptendpointgroup

import (
	"context"
	"fmt"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityInterceptEndpointGroupGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (pb.InterceptClient, longrunningpb.OperationsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))
	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("dialing networksecurity service: %w", err)
	}
	return pb.NewInterceptClient(conn), longrunningpb.NewOperationsClient(conn), nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityInterceptEndpointGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	// Get networksecurity GCP client
	gcpClient, operationsClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := NetworkSecurityInterceptEndpointGroupSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:               id.(*krm.NetworkSecurityInterceptEndpointGroupIdentity),
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		desired:          desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id               *krm.NetworkSecurityInterceptEndpointGroupIdentity
	gcpClient        pb.InterceptClient
	operationsClient longrunningpb.OperationsClient
	desired          *pb.InterceptEndpointGroup
	actual           *pb.InterceptEndpointGroup
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting InterceptEndpointGroup", "name", a.id)

	req := &pb.GetInterceptEndpointGroupRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetInterceptEndpointGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting InterceptEndpointGroup %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating InterceptEndpointGroup", "name", a.id)

	req := &pb.CreateInterceptEndpointGroupRequest{
		Parent:                   a.id.ParentString(),
		InterceptEndpointGroupId: a.id.Interceptendpointgroup,
		InterceptEndpointGroup:   a.desired,
	}
	op, err := a.gcpClient.CreateInterceptEndpointGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating InterceptEndpointGroup %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("InterceptEndpointGroup %s waiting creation: %w", a.id, err)
	}

	// Fetch fully-populated resource after creation
	latest, err := a.gcpClient.GetInterceptEndpointGroup(ctx, &pb.GetInterceptEndpointGroupRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting InterceptEndpointGroup after creation: %w", err)
	}

	log.V(2).Info("successfully created InterceptEndpointGroup", "name", a.id)

	return a.updateStatus(ctx, createOp, latest)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InterceptEndpointGroup", "name", a.id)

	diffs, updateMask, err := compareInterceptEndpointGroup(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	// Report diffs to the user
	diffs.Object = updateOp.GetUnstructured()
	structuredreporting.ReportDiff(ctx, diffs)

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	log.V(2).Info("fields need update", "name", a.id, "updateMask", updateMask)

	req := &pb.UpdateInterceptEndpointGroupRequest{
		UpdateMask:             updateMask,
		InterceptEndpointGroup: a.desired,
	}
	req.InterceptEndpointGroup.Name = a.id.String()

	op, err := a.gcpClient.UpdateInterceptEndpointGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("updating InterceptEndpointGroup %s: %w", a.id, err)
	}
	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("InterceptEndpointGroup %s waiting update: %w", a.id, err)
	}

	// Fetch fully-populated resource after update
	latest, err := a.gcpClient.GetInterceptEndpointGroup(ctx, &pb.GetInterceptEndpointGroupRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting InterceptEndpointGroup after update: %w", err)
	}

	log.V(2).Info("successfully updated InterceptEndpointGroup", "name", a.id)

	return a.updateStatus(ctx, updateOp, latest)
}

func compareInterceptEndpointGroup(ctx context.Context, actual, desired *pb.InterceptEndpointGroup) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	mapCtx := &direct.MapContext{}
	spec := NetworkSecurityInterceptEndpointGroupSpec_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual := NetworkSecurityInterceptEndpointGroupSpec_ToProto(mapCtx, spec)
	if mapCtx.Err() != nil {
		return nil, nil, mapCtx.Err()
	}
	maskedActual.Name = desired.Name

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityInterceptEndpointGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkSecurityInterceptEndpointGroupSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Interceptendpointgroup)
	u.SetGroupVersionKind(krm.NetworkSecurityInterceptEndpointGroupGVK)

	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting InterceptEndpointGroup", "name", a.id)

	req := &pb.DeleteInterceptEndpointGroupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInterceptEndpointGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent InterceptEndpointGroup, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting InterceptEndpointGroup %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return false, fmt.Errorf("waiting delete InterceptEndpointGroup %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted InterceptEndpointGroup", "name", a.id)
	return true, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.InterceptEndpointGroup) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkSecurityInterceptEndpointGroupStatus{}
	status.ObservedState = NetworkSecurityInterceptEndpointGroupObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
	for {
		req := &longrunningpb.GetOperationRequest{
			Name: op.GetName(),
		}
		current, err := a.operationsClient.GetOperation(ctx, req)
		if err != nil {
			return fmt.Errorf("getting operation %q: %w", op.GetName(), err)
		}
		if current.GetDone() {
			if current.GetError() != nil {
				return fmt.Errorf("operation failed: %s", current.GetError().GetMessage())
			}
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(1 * time.Second):
		}
	}
}
