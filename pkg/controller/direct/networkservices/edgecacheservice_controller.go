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

package networkservices

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/networkservices/apiv1"
	networkservicespb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetworkServicesEdgeCacheServiceGVK, NewEdgeCacheServiceModel)
}

func NewEdgeCacheServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEdgeCacheService{config: *config}, nil
}

var _ directbase.Model = &modelEdgeCacheService{}

type modelEdgeCacheService struct {
	config config.ControllerConfig
}

func (m *modelEdgeCacheService) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building EdgeCacheService client: %w", err)
	}
	return gcpClient, err
}

func (m *modelEdgeCacheService) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetworkServicesEdgeCacheService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEdgeCacheServiceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EdgeCacheServiceAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEdgeCacheService) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EdgeCacheServiceAdapter struct {
	id        *krm.EdgeCacheServiceIdentity
	gcpClient *gcp.Client
	desired   *krm.NetworkServicesEdgeCacheService
	actual    *networkservicespb.EdgeCacheService
}

var _ directbase.Adapter = &EdgeCacheServiceAdapter{}

func (a *EdgeCacheServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EdgeCacheService", "name", a.id)

	req := &networkservicespb.GetEdgeCacheServiceRequest{Name: a.id.String()}
	pb, err := a.gcpClient.GetEdgeCacheService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EdgeCacheService %q: %w", a.id, err)
	}

	a.actual = pb
	return true, nil
}

func (a *EdgeCacheServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EdgeCacheService", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkServicesEdgeCacheServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	req := &networkservicespb.CreateEdgeCacheServiceRequest{
		Parent:             a.id.Parent().String(),
		EdgeCacheServiceId: a.id.ID(),
		EdgeCacheService:   resource,
	}
	op, err := a.gcpClient.CreateEdgeCacheService(ctx, req)
	if err != nil {
		return fmt.Errorf("creating EdgeCacheService %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("EdgeCacheService %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created EdgeCacheService", "name", a.id)

	status := &krm.NetworkServicesEdgeCacheServiceStatus{}
	status.ObservedState = NetworkServicesEdgeCacheServiceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *EdgeCacheServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EdgeCacheService", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkServicesEdgeCacheServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	paths := []string{
		"description",
		"routing",
		"labels",
		"disable_quic",
		"require_tls",
		"edge_ssl_certificates",
		"ssl_policy",
		"log_config",
		"disable_http2",
		"edge_security_policy",
	}

	req := &networkservicespb.UpdateEdgeCacheServiceRequest{
		EdgeCacheService: resource,
		UpdateMask:       &fieldmaskpb.FieldMask{Paths: paths},
	}

	op, err := a.gcpClient.UpdateEdgeCacheService(ctx, req)
	if err != nil {
		return fmt.Errorf("updating EdgeCacheService %s: %w", a.id, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("EdgeCacheService %s waiting update: %w", a.id, err)
	}

	status := &krm.NetworkServicesEdgeCacheServiceStatus{}
	status.ObservedState = NetworkServicesEdgeCacheServiceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *EdgeCacheServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkServicesEdgeCacheService{}
	mapCtx := &direct.MapContext{}
	obj.Spec = *NetworkServicesEdgeCacheServiceSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkServicesEdgeCacheServiceGVK)

	u.Object = uObj
	return u, nil
}

func (a *EdgeCacheServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EdgeCacheService", "name", a.id)

	req := &networkservicespb.DeleteEdgeCacheServiceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEdgeCacheService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting EdgeCacheService %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted EdgeCacheService", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete EdgeCacheService %s: %w", a.id, err)
	}
	return true, nil
}
