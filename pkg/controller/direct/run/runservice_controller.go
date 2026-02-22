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

package run

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"google.golang.org/protobuf/proto"

	gcp "cloud.google.com/go/run/apiv2"

	pb "cloud.google.com/go/run/apiv2/runpb"
	runpb "cloud.google.com/go/run/apiv2/runpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.RunServiceGVK, NewServiceModel)
}

func NewServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelService{config: *config}, nil
}

var _ directbase.Model = &modelService{}

type modelService struct {
	config config.ControllerConfig
}

func (m *modelService) client(ctx context.Context) (*gcp.ServicesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewServicesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Service client: %w", err)
	}
	return gcpClient, err
}

func (m *modelService) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.RunService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewServiceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}
	if err := ResolveRunServiceRefs(ctx, reader, obj); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()
	desired := RunServiceSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	// Get run GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ServiceAdapter{
		id:                 id,
		gcpClient:          gcpClient,
		desired:            desired,
		lastModifiedCookie: obj.Status.LastModifiedCookie,
	}, nil
}

func (m *modelService) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if s, ok := strings.CutPrefix(url, "//run.googleapis.com/"); ok {
		// Direct controller for RunService only handles v2.
		s = strings.TrimPrefix(s, "v2/")

		var id krm.RunServiceIdentity
		if err := id.FromExternal(s); err != nil {
			log.V(2).Error(err, "url did not match RunService format", "url", url)
			return nil, nil
		}

		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &ServiceAdapter{
			gcpClient: gcpClient,
			id:        &id,
		}, nil
	}
	return nil, nil
}

type ServiceAdapter struct {
	id                 *krm.RunServiceIdentity
	gcpClient          *gcp.ServicesClient
	desired            *runpb.Service
	actual             *runpb.Service
	lastModifiedCookie *string
}

var _ directbase.Adapter = &ServiceAdapter{}

func (a *ServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Service", "name", a.id)

	req := &runpb.GetServiceRequest{Name: a.id.String()}
	found, err := a.gcpClient.GetService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Service %q: %w", a.id, err)
	}

	a.actual = found
	return true, nil
}

func (a *ServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Service", "name", a.id)
	req := &runpb.CreateServiceRequest{
		Parent:    fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		Service:   a.desired,
		ServiceId: a.id.Service,
	}
	op, err := a.gcpClient.CreateService(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Service %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for creation of service %q: %w", a.id, err)
	}
	log.V(2).Info("successfully created Service", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krm.RunServiceStatus{}
	status.ObservedState = RunServiceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	newCookie, err := common.NewLegacyCookie(a.desired, created)
	if err != nil {
		return fmt.Errorf("composing cookie: %w", err)
	}
	log.V(2).Info("Service cookie added", "name", a.id, "new-cookie", newCookie.String())
	status.LastModifiedCookie = direct.LazyPtr(newCookie.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Service", "name", a.id)

	currentCookie, err := common.NewLegacyCookie(a.desired, a.actual)
	if err != nil {
		return err
	}
	if currentCookie.Equal(a.lastModifiedCookie) {
		log.V(2).Info("resource is up to date", "name", a.id)
		return a.updateStatus(ctx, a.actual, updateOp)
	}

	updateService := proto.Clone(a.desired).(*runpb.Service)
	updateService.Name = a.actual.Name
	req := &runpb.UpdateServiceRequest{
		Service: updateService,
	}
	op, err := a.gcpClient.UpdateService(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Service %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Service %s waiting update: %w", a.id, err)
	}
	log.Info("successfully updated Service", "name", a.id)
	return a.updateStatus(ctx, updated, updateOp)
}

func (a *ServiceAdapter) updateStatus(ctx context.Context, updated *pb.Service, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	status := &krm.RunServiceStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = RunServiceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	updatedCookie, err := common.NewLegacyCookie(a.desired, updated)
	if err != nil {
		return err
	}
	status.LastModifiedCookie = direct.LazyPtr(updatedCookie.String())
	status.ExternalRef = direct.LazyPtr(a.id.String())

	if !updatedCookie.Equal(a.lastModifiedCookie) {
		log.Info("Service cookie updated", "name", a.id, "old-cookie", direct.ValueOf(a.lastModifiedCookie),
			"new-cookie", updatedCookie.String())
	}

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.RunService{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(RunServiceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Service)
	u.SetGroupVersionKind(krm.RunServiceGVK)

	return u, nil
}

func (a *ServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Service", "name", a.id)

	name := a.id.String()
	req := &runpb.DeleteServiceRequest{Name: name}
	op, err := a.gcpClient.DeleteService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Service, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Service %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Service", "name", a.id)

	if _, err = op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting delete Service %s: %w", a.id, err)
	}
	return true, nil
}
