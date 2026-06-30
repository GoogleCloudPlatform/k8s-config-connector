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
	"strings"
	"time"

	"google.golang.org/api/option"
	gcp "google.golang.org/api/serviceusage/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/serviceusage/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.ServiceGVK, NewServiceModel)
}

func NewServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &serviceModel{config: *config}, nil
}

var _ directbase.Model = &serviceModel{}

type serviceModel struct {
	config config.ControllerConfig
}

func (m *serviceModel) client(ctx context.Context) (*gcp.Service, error) {
	var opts []option.ClientOption

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building serviceusage client: %w", err)
	}

	return gcpClient, nil
}

func (m *serviceModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	obj := &krm.Service{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := ServiceSpec_ToAPI(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &serviceAdapter{
		gcpClient:   gcpClient,
		id:          id.(*krm.ServiceURIIdentity),
		desired:     desired,
		desiredKube: obj,
	}, nil
}

func (m *serviceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if s, ok := strings.CutPrefix(url, "//serviceusage.googleapis.com/"); ok {
		s = strings.TrimPrefix(s, "v1/")

		var id krm.ServiceURIIdentity
		err := id.FromExternal(s)
		if err != nil {
			log.V(2).Error(err, "url did not match Service format", "url", url)
			return nil, nil
		}

		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &serviceAdapter{
			gcpClient: gcpClient,
			id:        &id,
		}, nil
	}
	return nil, nil
}

type serviceAdapter struct {
	gcpClient   *gcp.Service
	id          *krm.ServiceURIIdentity
	desired     *gcp.GoogleApiServiceusageV1Service
	desiredKube *krm.Service
	actual      *gcp.GoogleApiServiceusageV1Service
}

var _ directbase.Adapter = &serviceAdapter{}

func (a *serviceAdapter) Find(ctx context.Context) (bool, error) {
	name := a.id.String()
	service, err := a.gcpClient.Services.Get(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting service %q: %w", name, err)
	}

	a.actual = service
	if service.State == "ENABLED" {
		return true, nil
	}
	return false, nil
}

func (a *serviceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("enabling service", "id", fqn)

	op, err := a.gcpClient.Services.Enable(fqn, &gcp.EnableServiceRequest{}).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("enabling service %q: %w", fqn, err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for service %q enabling: %w", fqn, err)
	}

	return nil
}

func (a *serviceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// Service only has ProjectRef and ResourceID, which are immutable.
	// So Update is a no-op as the service is already enabled.
	return nil
}

func (a *serviceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	u := deleteOp.GetUnstructured()
	disableOnDestroy, err := GetDisableOnDestroy(u)
	if err != nil {
		return false, err
	}

	if !disableOnDestroy {
		return true, nil
	}

	fqn := a.id.String()
	req := &gcp.DisableServiceRequest{}
	disableDependentServices, err := GetDisableDependentServices(u)
	if err != nil {
		return false, err
	}
	req.DisableDependentServices = disableDependentServices

	op, err := a.gcpClient.Services.Disable(fqn, req).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("disabling service %q: %w", fqn, err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for service %q disabling: %w", fqn, err)
	}

	return true, nil
}

func (a *serviceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("service %q not found", a.id.Service)
	}

	obj := &krm.Service{}
	obj.SetName(a.id.Service)
	obj.SetGroupVersionKind(krm.ServiceGVK)
	if a.desiredKube != nil {
		obj.SetLabels(a.desiredKube.Labels)
	}

	obj.Spec = krm.ServiceSpec{
		ProjectRef: &refs.ProjectRef{
			External: a.id.Project,
		},
		ResourceID: &a.id.Service,
	}

	unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("error converting service to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{Object: unstructuredMap}
	return u, nil
}

func (a *serviceAdapter) waitForOp(ctx context.Context, op *gcp.Operation) error {
	return common.WaitForDoneOrTimeout(ctx, 2*time.Second, func() (bool, error) {
		current, err := a.gcpClient.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return false, fmt.Errorf("getting operation status of %q: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return true, fmt.Errorf("operation %q completed with error: %s", op.Name, current.Error.Message)
			} else {
				return true, nil
			}
		}
		return false, nil
	})
}

func getBoolAnnotation(o client.Object, key string, defaultValue bool) (bool, error) {
	v, found := o.GetAnnotations()[key]
	if !found {
		return defaultValue, nil
	}
	v = strings.TrimSpace(v)
	v = strings.ToLower(v)
	switch v {
	case "false":
		return false, nil
	case "true":
		return true, nil
	default:
		return false, fmt.Errorf("unhandled value for %s annotation %q", key, v)
	}
}

// GetDisableOnDestroy returns the value of the cnrm.cloud.google.com/disable-on-destroy annotation
func GetDisableOnDestroy(o client.Object) (bool, error) {
	return getBoolAnnotation(o, "cnrm.cloud.google.com/disable-on-destroy", true)
}

// GetDisableDependentServices returns the value of the cnrm.cloud.google.com/disable-dependent-services annotation
func GetDisableDependentServices(o client.Object) (bool, error) {
	return getBoolAnnotation(o, "cnrm.cloud.google.com/disable-dependent-services", false)
}
