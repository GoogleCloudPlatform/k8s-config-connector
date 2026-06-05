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

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "cloud.google.com/go/networksecurity/apiv1/networksecuritypb"
	option "google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	grpc "google.golang.org/grpc"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
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

func (m *model) client(ctx context.Context) (pb.InterceptClient, *grpc.ClientConn, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, option.WithEndpoint("networksecurity.googleapis.com:443"))
	conn, err := gtransport.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("dialing networksecurity gRPC: %w", err)
	}
	return pb.NewInterceptClient(conn), conn, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkSecurityInterceptEndpointGroup{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNetworkSecurityInterceptEndpointGroupIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	copied := obj.DeepCopy()
	desired := NetworkSecurityInterceptEndpointGroupSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, conn, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		conn:      conn,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.NetworkSecurityInterceptEndpointGroupIdentity
	conn      *grpc.ClientConn
	gcpClient pb.InterceptClient
	desired   *pb.InterceptEndpointGroup
	actual    *pb.InterceptEndpointGroup
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting InterceptEndpointGroup", "name", a.id)

	req := &pb.GetInterceptEndpointGroupRequest{Name: a.id.String()}
	found, err := a.gcpClient.GetInterceptEndpointGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting InterceptEndpointGroup %q: %w", a.id, err)
	}

	a.actual = found
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating InterceptEndpointGroup", "name", a.id)

	req := &pb.CreateInterceptEndpointGroupRequest{
		Parent:                   fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		InterceptEndpointGroupId: a.id.Interceptendpointgroup,
		InterceptEndpointGroup:   a.desired,
	}
	op, err := a.gcpClient.CreateInterceptEndpointGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("creating InterceptEndpointGroup %s: %w", a.id, err)
	}
	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("waiting for creation of InterceptEndpointGroup %q: %w", a.id, err)
	}

	// Fetch created state
	latest, err := a.gcpClient.GetInterceptEndpointGroup(ctx, &pb.GetInterceptEndpointGroupRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching created InterceptEndpointGroup %q: %w", a.id, err)
	}

	log.V(2).Info("successfully created InterceptEndpointGroup", "name", a.id)

	status := &krm.NetworkSecurityInterceptEndpointGroupStatus{}
	if err := a.updateStatus(ctx, status, latest); err != nil {
		return err
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating InterceptEndpointGroup", "name", a.id)

	updateReq := proto.Clone(a.desired).(*pb.InterceptEndpointGroup)
	updateReq.Name = a.actual.Name

	paths, err := common.CompareProtoMessage(updateReq, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		status := &krm.NetworkSecurityInterceptEndpointGroupStatus{}
		if err := a.updateStatus(ctx, status, a.actual); err != nil {
			return err
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &pb.UpdateInterceptEndpointGroupRequest{
		InterceptEndpointGroup: updateReq,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		},
	}
	op, err := a.gcpClient.UpdateInterceptEndpointGroup(ctx, req)
	if err != nil {
		return fmt.Errorf("updating InterceptEndpointGroup %s: %w", a.id, err)
	}
	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("InterceptEndpointGroup %s waiting update: %w", a.id, err)
	}

	// Fetch updated state
	latest, err := a.gcpClient.GetInterceptEndpointGroup(ctx, &pb.GetInterceptEndpointGroupRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching updated InterceptEndpointGroup %q: %w", a.id, err)
	}

	log.Info("successfully updated InterceptEndpointGroup", "name", a.id)

	status := &krm.NetworkSecurityInterceptEndpointGroupStatus{}
	if err := a.updateStatus(ctx, status, latest); err != nil {
		return err
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting InterceptEndpointGroup", "name", a.id)

	req := &pb.DeleteInterceptEndpointGroupRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteInterceptEndpointGroup(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting InterceptEndpointGroup %s: %w", a.id, err)
	}
	err = a.waitForOperation(ctx, op)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("waiting for deletion of InterceptEndpointGroup %q: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted InterceptEndpointGroup", "name", a.id)
	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkSecurityInterceptEndpointGroup{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *NetworkSecurityInterceptEndpointGroupSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Re-inject project/location back to spec because mapper doesn't know about it
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) updateStatus(ctx context.Context, status *krm.NetworkSecurityInterceptEndpointGroupStatus, updated *pb.InterceptEndpointGroup) error {
	mapCtx := &direct.MapContext{}
	status.ObservedState = NetworkSecurityInterceptEndpointGroupObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return nil
}

func (a *Adapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
	operationsClient := longrunningpb.NewOperationsClient(a.conn)
	return common.WaitForDoneOrTimeout(ctx, 2*time.Second, func() (bool, error) {
		latest, err := operationsClient.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.Name})
		if err != nil {
			return false, fmt.Errorf("getting operation %q: %w", op.Name, err)
		}
		return latest.Done, nil
	})
}
