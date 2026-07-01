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

package aistreams

import (
	"context"
	"fmt"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aistreams/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	common "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"google.golang.org/api/option"
	"google.golang.org/api/transport"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	pb "google.golang.org/genproto/googleapis/partner/aistreams/v1alpha1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIStreamsClusterGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

type Adapter struct {
	id               *krm.AIStreamsClusterIdentity
	gcpClient        pb.AIStreamsClient
	operationsClient longrunningpb.OperationsClient
	conn             *grpc.ClientConn
	desired          *pb.Cluster
	actual           *pb.Cluster
	model            *model
}

var _ directbase.Adapter = &Adapter{}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIStreamsCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	clusterID := id.(*krm.AIStreamsClusterIdentity)

	var opts []option.ClientOption
	opts, err = m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	// We use DialGRPC directly because there is no official high-level Go client library
	// generated for the aistreams service. DialGRPC requires an explicit endpoint since
	// it does not have a default endpoint built-in.
	opts = append(opts, option.WithEndpoint("aistreams.googleapis.com:443"))
	conn, err := transport.DialGRPC(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("dialing aistreams gRPC endpoint: %w", err)
	}

	gcpClient := pb.NewAIStreamsClient(conn)
	operationsClient := longrunningpb.NewOperationsClient(conn)

	mapCtx := &direct.MapContext{}
	desired := AIStreamsClusterSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		conn.Close()
		return nil, mapCtx.Err()
	}

	return &Adapter{
		id:               clusterID,
		gcpClient:        gcpClient,
		operationsClient: operationsClient,
		conn:             conn,
		desired:          desired,
		model:            m,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("finding AIStreamsCluster", "id", a.id)

	req := &pb.GetClusterRequest{
		Name: a.id.String(),
	}
	cluster, err := a.gcpClient.GetCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AIStreamsCluster %s: %w", a.id, err)
	}

	a.actual = cluster
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AIStreamsCluster", "id", a.id)

	req := &pb.CreateClusterRequest{
		Parent:    fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		ClusterId: a.id.Cluster,
		Cluster:   a.desired,
	}
	op, err := a.gcpClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AIStreamsCluster %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("waiting for AIStreamsCluster %s creation: %w", a.id, err)
	}

	refetched, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching AIStreamsCluster %s after creation: %w", a.id, err)
	}
	a.actual = refetched

	status := &krm.AIStreamsClusterStatus{}
	if err := a.updateStatus(ctx, status, refetched); err != nil {
		return err
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AIStreamsCluster", "id", a.id)

	updateReq := proto.CloneOf(a.desired)
	updateReq.Name = a.actual.Name

	paths, err := common.CompareProtoMessage(updateReq, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "id", a.id)
		status := &krm.AIStreamsClusterStatus{}
		if err := a.updateStatus(ctx, status, a.actual); err != nil {
			return err
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &pb.UpdateClusterRequest{
		Cluster: updateReq,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: paths.UnsortedList(),
		},
	}

	op, err := a.gcpClient.UpdateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AIStreamsCluster %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		return fmt.Errorf("waiting for AIStreamsCluster %s update: %w", a.id, err)
	}

	refetched, err := a.gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("fetching AIStreamsCluster %s after update: %w", a.id, err)
	}
	a.actual = refetched

	status := &krm.AIStreamsClusterStatus{}
	if err := a.updateStatus(ctx, status, refetched); err != nil {
		return err
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AIStreamsCluster", "id", a.id)

	req := &pb.DeleteClusterRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting AIStreamsCluster %s: %w", a.id, err)
	}

	err = a.waitForOperation(ctx, op)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("waiting for deletion of AIStreamsCluster %q: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted AIStreamsCluster", "id", a.id)
	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIStreamsCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *AIStreamsClusterSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = direct.LazyPtr(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) updateStatus(ctx context.Context, status *krm.AIStreamsClusterStatus, cluster *pb.Cluster) error {
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIStreamsClusterObservedState_FromProto(mapCtx, cluster)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return nil
}

func (a *Adapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
	if op.Done {
		if op.GetError() != nil {
			return fmt.Errorf("operation failed: %s", op.GetError().GetMessage())
		}
		return nil
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(5 * time.Second):
			current, err := a.operationsClient.GetOperation(ctx, &longrunningpb.GetOperationRequest{Name: op.Name})
			if err != nil {
				return fmt.Errorf("getting operation %q: %w", op.Name, err)
			}
			if current.Done {
				if current.GetError() != nil {
					return fmt.Errorf("operation failed: %s", current.GetError().GetMessage())
				}
				return nil
			}
		}
	}
}
