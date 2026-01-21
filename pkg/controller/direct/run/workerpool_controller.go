// Copyright 2025 Google LLC
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

package run

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/run/apiv2"
	pb "cloud.google.com/go/run/apiv2/runpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.RunWorkerPoolGVK, NewWorkerPoolModel)
}

func NewWorkerPoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkerPool{config: *config}, nil
}

var _ directbase.Model = &modelWorkerPool{}

type modelWorkerPool struct {
	config config.ControllerConfig
}

func (m *modelWorkerPool) client(ctx context.Context) (*gcp.WorkerPoolsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewWorkerPoolsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building WorkerPool client: %w", err)
	}
	return gcpClient, err
}

func (m *modelWorkerPool) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.RunWorkerPool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewWorkerPoolIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()
	desired := RunWorkerPoolSpec_v1alpha1_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get run GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WorkerPoolAdapter{
		id:                 id,
		gcpClient:          gcpClient,
		desired:            desired,
		lastModifiedCookie: obj.Status.LastModifiedCookie,
	}, nil
}

func (m *modelWorkerPool) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if s, ok := strings.CutPrefix(url, "//run.googleapis.com/"); ok {
		s = strings.TrimPrefix(s, "v2/")

		var id krm.WorkerPoolIdentity
		if err := id.FromExternal(s); err != nil {
			log.V(2).Error(err, "url did not match RunWorkerPool format", "url", url)
			return nil, nil
		}

		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &WorkerPoolAdapter{
			gcpClient: gcpClient,
			id:        &id,
		}, nil
	}
	return nil, nil
}

type WorkerPoolAdapter struct {
	id                 *krm.WorkerPoolIdentity
	gcpClient          *gcp.WorkerPoolsClient
	desired            *pb.WorkerPool
	actual             *pb.WorkerPool
	lastModifiedCookie *string
}

var _ directbase.Adapter = &WorkerPoolAdapter{}

// Find retrieves the GCP resource.
func (a *WorkerPoolAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting WorkerPool", "name", a.id)

	req := &pb.GetWorkerPoolRequest{Name: a.id.String()}
	found, err := a.gcpClient.GetWorkerPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WorkerPool %q: %w", a.id, err)
	}

	a.actual = found
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WorkerPoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating WorkerPool", "name", a.id)
	req := &pb.CreateWorkerPoolRequest{
		Parent:       a.id.Parent().String(),
		WorkerPool:   a.desired,
		WorkerPoolId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateWorkerPool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WorkerPool %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of WorkerPool %q: %w", a.id, err)
	}
	log.V(2).Info("successfully created WorkerPool", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.RunWorkerPoolStatus{}
	status.ObservedState = RunWorkerPoolObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	newCookie, err := common.NewCookie(a.desired, created)
	if err != nil {
		return fmt.Errorf("composing cookie: %w", err)
	}
	log.V(2).Info("WorkerPool cookie added", "name", a.id, "new-cookie", newCookie.String())
	status.LastModifiedCookie = direct.LazyPtr(newCookie.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WorkerPoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating WorkerPool", "name", a.id)

	currentCookie, err := common.NewCookie(a.desired, a.actual)
	if err != nil {
		return err
	}
	if currentCookie.Equal(a.lastModifiedCookie) {
		log.V(2).Info("resource is up to date", "name", a.id)
		return a.updateStatus(ctx, a.actual, updateOp)
	}

	// We need to set the name for the update request, but we don't want to modify a.desired
	// as that would change the computed hash (specHash).
	updateWorkerPool := proto.Clone(a.desired).(*pb.WorkerPool)
	updateWorkerPool.Name = a.actual.Name
	req := &pb.UpdateWorkerPoolRequest{
		WorkerPool: updateWorkerPool,
	}
	op, err := a.gcpClient.UpdateWorkerPool(ctx, req)
	if err != nil {
		return fmt.Errorf("updating WorkerPool %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkerPool %s waiting update: %w", a.id, err)
	}
	log.Info("successfully updated WorkerPool", "name", a.id)
	return a.updateStatus(ctx, updated, updateOp)
}

func (a *WorkerPoolAdapter) updateStatus(ctx context.Context, updated *pb.WorkerPool, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	status := &krm.RunWorkerPoolStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = RunWorkerPoolObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())

	updatedCookie, err := common.NewCookie(a.desired, updated)
	if err != nil {
		return err
	}
	status.LastModifiedCookie = direct.LazyPtr(updatedCookie.String())

	if !updatedCookie.Equal(a.lastModifiedCookie) {
		log.Info("WorkerPool cookie updated", "name", a.id, "old-cookie", direct.ValueOf(a.lastModifiedCookie),
			"new-cookie", updatedCookie.String())
	}

	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *WorkerPoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.RunWorkerPool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(RunWorkerPoolSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef.External = a.id.Parent().ProjectID
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.RunWorkerPoolGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *WorkerPoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting WorkerPool", "name", a.id)

	name := a.id.String()
	req := &pb.DeleteWorkerPoolRequest{Name: name}
	op, err := a.gcpClient.DeleteWorkerPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent WorkerPool, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting WorkerPool %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted WorkerPool", "name", a.id)

	if _, err = op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting delete WorkerPool %s: %w", a.id, err)
	}
	return true, nil
}
