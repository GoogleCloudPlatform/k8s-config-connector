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

package aistreamscluster

import (
	"context"
	"fmt"
	"reflect"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aistreams/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"google.golang.org/api/option"
	"google.golang.org/api/transport/grpc"
	lro "google.golang.org/genproto/googleapis/longrunning"
	pb "google.golang.org/genproto/googleapis/partner/aistreams/v1alpha1"
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

func (m *model) client(ctx context.Context) (pb.AIStreamsClient, lro.OperationsClient, error) {
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, nil, err
	}
	opts = append(opts, option.WithEndpoint("aistreams.googleapis.com:443"))
	conn, err := grpc.Dial(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("building aistreams client: %w", err)
	}
	return pb.NewAIStreamsClient(conn), lro.NewOperationsClient(conn), nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader

	obj := &krm.AIStreamsCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.ParseAIStreamsClusterIdentity(direct.ValueOf(obj.Status.ExternalRef))
	if err != nil {
		idInterf, err := obj.GetIdentity(ctx, reader)
		if err != nil {
			return nil, err
		}
		id = idInterf.(*krm.AIStreamsClusterIdentity)
	}

	mapCtx := &direct.MapContext{}

	return &Adapter{
		id:     id,
		obj:    obj,
		mapCtx: mapCtx,
		model:  m,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id     *krm.AIStreamsClusterIdentity
	obj    *krm.AIStreamsCluster
	actual *pb.Cluster
	mapCtx *direct.MapContext
	model  *model
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(a.id.String())
	log.V(2).Info("getting AIStreamsCluster", "name", a.id.String())

	gcpClient, _, err := a.model.client(ctx)
	if err != nil {
		return false, err
	}

	req := &pb.GetClusterRequest{
		Name: a.id.String(),
	}
	cluster, err := gcpClient.GetCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AIStreamsCluster %q: %w", a.id.String(), err)
	}

	a.actual = cluster
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(a.id.String())
	log.V(2).Info("creating AIStreamsCluster", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	gcpClient, lroClient, err := a.model.client(ctx)
	if err != nil {
		return err
	}

	desired := AIStreamsClusterSpec_ToProto(mapCtx, &a.obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	req := &pb.CreateClusterRequest{
		Parent:    parent,
		ClusterId: a.id.ID(),
		Cluster:   desired,
	}

	op, err := gcpClient.CreateCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AIStreamsCluster %s: %w", a.id.String(), err)
	}

	if err := a.waitLRO(ctx, lroClient, op); err != nil {
		return fmt.Errorf("AIStreamsCluster %s waiting creation: %w", a.id.String(), err)
	}

	created, err := gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: a.id.String()})
	if err != nil {
		return fmt.Errorf("getting created AIStreamsCluster: %w", err)
	}

	status := &krm.AIStreamsClusterStatus{}
	status.ObservedState = AIStreamsClusterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(a.id.String())
	log.V(2).Info("updating AIStreamsCluster", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	gcpClient, lroClient, err := a.model.client(ctx)
	if err != nil {
		return err
	}

	desired := AIStreamsClusterSpec_ToProto(mapCtx, &a.obj.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(desired.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	if len(updateMask.Paths) > 0 {
		req := &pb.UpdateClusterRequest{
			Cluster:    desired,
			UpdateMask: updateMask,
		}

		op, err := gcpClient.UpdateCluster(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AIStreamsCluster %s: %w", a.id.String(), err)
		}

		if err := a.waitLRO(ctx, lroClient, op); err != nil {
			return fmt.Errorf("AIStreamsCluster %s waiting update: %w", a.id.String(), err)
		}

		a.actual, err = gcpClient.GetCluster(ctx, &pb.GetClusterRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("getting updated AIStreamsCluster: %w", err)
		}
	}

	status := &krm.AIStreamsClusterStatus{}
	status.ObservedState = AIStreamsClusterObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(a.id.String())
	log.V(2).Info("deleting AIStreamsCluster", "name", a.id.String())

	gcpClient, lroClient, err := a.model.client(ctx)
	if err != nil {
		return false, err
	}

	req := &pb.DeleteClusterRequest{
		Name: a.id.String(),
	}

	op, err := gcpClient.DeleteCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting AIStreamsCluster %s: %w", a.id.String(), err)
	}

	err = a.waitLRO(ctx, lroClient, op)
	if err != nil {
		return false, err
	}
	return false, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called or no object found")
	}
	u := &unstructured.Unstructured{}
	obj := &krm.AIStreamsCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIStreamsClusterSpec_FromProto(mapCtx, a.actual))
	if err := mapCtx.Err(); err != nil {
		return nil, err
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: "projects/" + a.id.Project}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.AIStreamsClusterGVK)
	return u, nil
}

func (a *Adapter) waitLRO(ctx context.Context, lroClient lro.OperationsClient, op *lro.Operation) error {
	if op.Done {
		if op.GetError() != nil {
			return fmt.Errorf("operation failed: %v", op.GetError().Message)
		}
		return nil
	}

	err := common.WaitForDoneOrTimeout(ctx, 5*time.Second, func() (bool, error) {
		req := &lro.GetOperationRequest{
			Name: op.Name,
		}
		pollOp, err := lroClient.GetOperation(ctx, req)
		if err != nil {
			return false, err
		}
		if pollOp.Done {
			if pollOp.GetError() != nil {
				return true, fmt.Errorf("operation failed: %v", pollOp.GetError().Message)
			}
			return true, nil
		}
		return false, nil
	})
	return err
}
