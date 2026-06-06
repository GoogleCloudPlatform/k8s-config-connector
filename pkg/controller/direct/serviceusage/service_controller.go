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

package serviceusage

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/serviceusage/apiv1"
	serviceusagepb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/serviceusage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ServiceGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

var _ directbase.Model = &model{}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Service client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.Service{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		id:        id.(*krm.ServiceURIIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.ServiceURIIdentity
	gcpClient *gcp.Client
	desired   *krm.Service
	actual    *serviceusagepb.Service
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Service", "name", a.id)

	req := &serviceusagepb.GetServiceRequest{Name: a.id.String()}
	service, err := a.gcpClient.GetService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Service %q: %w", a.id, err)
	}

	a.actual = service
	if service.State == serviceusagepb.State_ENABLED {
		return true, nil
	}
	return false, nil
}

func (a *Adapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("enabling Service", "name", a.id)

	req := &serviceusagepb.EnableServiceRequest{
		Name: a.id.String(),
	}
	lro, err := a.gcpClient.EnableService(ctx, req)
	if err != nil {
		return fmt.Errorf("enabling Service %s: %w", a.id, err)
	}
	_, err = lro.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for Service %s to be enabled: %w", a.id, err)
	}
	log.V(2).Info("successfully enabled Service", "name", a.id)

	status := &krm.ServiceStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	// Since Service has no spec fields other than projectRef/resourceID,
	// and state is already verified to be ENABLED in Find, there is nothing to update.
	return nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.Service{}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Service)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Service)
	u.SetGroupVersionKind(krm.ServiceGVK)

	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("disabling Service", "name", a.id)

	req := &serviceusagepb.DisableServiceRequest{
		Name: a.id.String(),
	}
	lro, err := a.gcpClient.DisableService(ctx, req)
	if err != nil {
		return false, fmt.Errorf("disabling Service %s: %w", a.id, err)
	}
	_, err = lro.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for Service %s to be disabled: %w", a.id, err)
	}
	log.V(2).Info("successfully disabled Service", "name", a.id)
	return true, nil
}
