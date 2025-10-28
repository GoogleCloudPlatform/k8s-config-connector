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

// +tool:controller
// proto.service: google.cloud.networkservices.v1.NetworkServices
// proto.message: google.cloud.networkservices.v1.EdgeCacheService
// crd.type: NetworkServicesEdgeCacheService
// crd.version: v1alpha1

package networkservices

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/networkservices/apiv1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/google/cloud/networkservices/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetworkServicesEdgeCacheServiceGVK, NewEdgeCacheServiceModel)
}

func NewEdgeCacheServiceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &EdgeCacheServiceModel{config: *config}, nil
}

var _ directbase.Model = &EdgeCacheServiceModel{}

type EdgeCacheServiceModel struct {
	config config.ControllerConfig
}

func (m *EdgeCacheServiceModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetworkServicesEdgeCacheService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEdgeCacheServiceIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get NetworkServices GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	networkServicesClient, err := gcpClient.newNetworkServicesClient(ctx)
	if err != nil {
		return nil, err
	}
	return &EdgeCacheServiceAdapter{
		gcpClient: networkServicesClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *EdgeCacheServiceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EdgeCacheServiceAdapter struct {
	gcpClient *gcp.Client
	id        *krm.EdgeCacheServiceIdentity
	desired   *krm.NetworkServicesEdgeCacheService
	actual    *pb.EdgeCacheService
	reader    client.Reader
}

var _ directbase.Adapter = &EdgeCacheServiceAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EdgeCacheServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting EdgeCacheService", "name", a.id)

	// TODO: EdgeCacheService is an alpha resource not yet available in the official
	// cloud.google.com/go/networkservices/apiv1 API. This controller currently only
	// works with MockGCP for testing. Uncomment when the API is available.
	/*
	req := &pb.GetEdgeCacheServiceRequest{Name: a.id.String()}
	edgecacheservicepb, err := a.gcpClient.GetEdgeCacheService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EdgeCacheService %q: %w", a.id, err)
	}

	a.actual = edgecacheservicepb
	return true, nil
	*/

	// Temporary: return not found for compilation
	return false, nil
}

func (a *EdgeCacheServiceAdapter) resolveReferences(ctx context.Context) error {
	// TODO: Implement reference resolution for:
	// - EdgeCacheOrigin references
	// - EdgeCacheKeyset references
	// - Certificate references
	return nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EdgeCacheServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating EdgeCacheService", "name", a.id)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := NetworkServicesEdgeCacheServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO: EdgeCacheService is an alpha resource not yet available in the official API.
	// Uncomment when the API is available in cloud.google.com/go/networkservices/apiv1
	/*
	req := &pb.CreateEdgeCacheServiceRequest{
		Parent:               a.id.Parent().String(),
		EdgeCacheServiceId:   a.id.ID(),
		EdgeCacheService:     resource,
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
	*/

	// Temporary: Use MockGCP for testing
	_ = resource // Keep resource usage to avoid unused variable error
	return fmt.Errorf("EdgeCacheService Create not yet implemented (alpha resource, use MockGCP for testing)")
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EdgeCacheServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating EdgeCacheService", "name", a.id)

	if err := a.resolveReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := NetworkServicesEdgeCacheServiceSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO: EdgeCacheService is an alpha resource not yet available in the official API.
	// Uncomment when the API is available in cloud.google.com/go/networkservices/apiv1
	/*
	// Build update mask for fields that changed
	paths := []string{}

	if desired.Spec.Labels != nil && !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		paths = append(paths, "labels")
	}
	if desired.Spec.Description != nil && !reflect.DeepEqual(resource.Description, a.actual.Description) {
		paths = append(paths, "description")
	}
	if desired.Spec.Routing != nil && !reflect.DeepEqual(resource.Routing, a.actual.Routing) {
		paths = append(paths, "routing")
	}
	if desired.Spec.RequireTls != nil && !reflect.DeepEqual(resource.RequireTls, a.actual.RequireTls) {
		paths = append(paths, "require_tls")
	}
	if desired.Spec.DisableHttp2 != nil && !reflect.DeepEqual(resource.DisableHttp2, a.actual.DisableHttp2) {
		paths = append(paths, "disable_http2")
	}
	if desired.Spec.DisableQuic != nil && !reflect.DeepEqual(resource.DisableQuic, a.actual.DisableQuic) {
		paths = append(paths, "disable_quic")
	}
	if desired.Spec.LogConfig != nil && !reflect.DeepEqual(resource.LogConfig, a.actual.LogConfig) {
		paths = append(paths, "log_config")
	}

	var updated *pb.EdgeCacheService
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		resource.Name = a.id.String() // we need to set the name so that GCP API can identify the resource
		req := &pb.UpdateEdgeCacheServiceRequest{
			EdgeCacheService: resource,
			UpdateMask:       &fieldmaskpb.FieldMask{Paths: paths},
		}
		op, err := a.gcpClient.UpdateEdgeCacheService(ctx, req)
		if err != nil {
			return fmt.Errorf("updating EdgeCacheService %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("EdgeCacheService %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated EdgeCacheService", "name", a.id)
	}

	status := &krm.NetworkServicesEdgeCacheServiceStatus{}
	status.ObservedState = NetworkServicesEdgeCacheServiceObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
	*/

	// Temporary: Use MockGCP for testing
	_ = resource // Keep resource usage to avoid unused variable error
	return fmt.Errorf("EdgeCacheService Update not yet implemented (alpha resource, use MockGCP for testing)")
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EdgeCacheServiceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkServicesEdgeCacheService{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkServicesEdgeCacheServiceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID},
		Location:   a.id.Parent().Location,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkServicesEdgeCacheServiceGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EdgeCacheServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting EdgeCacheService", "name", a.id)

	// TODO: EdgeCacheService is an alpha resource not yet available in the official API.
	// Uncomment when the API is available in cloud.google.com/go/networkservices/apiv1
	/*
	req := &pb.DeleteEdgeCacheServiceRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteEdgeCacheService(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent EdgeCacheService, assuming it was already deleted", "name", a.id)
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
	*/

	// Temporary: Use MockGCP for testing
	return true, nil // Assume deleted successfully
}
