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

package networkservices

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/edgecacheservice/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"

	"google.golang.org/api/option"
	transport_grpc "google.golang.org/api/transport/grpc"
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

func (m *modelEdgeCacheService) client(ctx context.Context) (pb.EdgeCacheServicesClient, error) {
	clientOpts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}

	// Use standard transport to dial.
	// Prepend default endpoint so it can be overridden by clientOpts (e.g. for MockGCP)
	endpoint := "networkservices.googleapis.com:443"
	conn, err := transport_grpc.Dial(ctx, append([]option.ClientOption{option.WithEndpoint(endpoint)}, clientOpts...)...)
	if err != nil {
		return nil, fmt.Errorf("dialing edgecacheservice: %w", err)
	}

	return pb.NewEdgeCacheServicesClient(conn), nil
}

func (m *modelEdgeCacheService) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetworkServicesEdgeCacheService{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EdgeCacheServiceAdapter{
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelEdgeCacheService) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type EdgeCacheServiceAdapter struct {
	gcpClient pb.EdgeCacheServicesClient
	desired   *krm.NetworkServicesEdgeCacheService
	actual    *pb.EdgeCacheService
}

var _ directbase.Adapter = &EdgeCacheServiceAdapter{}

func (a *EdgeCacheServiceAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)

	project := a.desired.Spec.ProjectRef.External
	location := a.desired.Spec.Location
	resourceID := a.desired.Spec.ResourceID
	if resourceID == nil {
		resourceID = &a.desired.Name
	}
	name := fmt.Sprintf("projects/%s/locations/%s/edgeCacheServices/%s", project, location, *resourceID)

	log.V(2).Info("getting EdgeCacheService", "name", name)

	actual, err := a.gcpClient.GetEdgeCacheService(ctx, &pb.GetEdgeCacheServiceRequest{Name: name})
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting EdgeCacheService %q: %w", name, err)
	}

	a.actual = actual
	return true, nil
}

func (a *EdgeCacheServiceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)

	project := a.desired.Spec.ProjectRef.External
	location := a.desired.Spec.Location
	resourceID := a.desired.Spec.ResourceID
	if resourceID == nil {
		resourceID = &a.desired.Name
	}
	parentName := fmt.Sprintf("projects/%s/locations/%s", project, location)
	name := fmt.Sprintf("%s/edgeCacheServices/%s", parentName, *resourceID)

	log.V(2).Info("creating EdgeCacheService", "name", name)

	mapCtx := &direct.MapContext{}
	desired := EdgeCacheServiceSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Ensure labels are set
	if desired.Labels == nil {
		desired.Labels = make(map[string]string)
	}
	desired.Labels["managed-by-cnrm"] = "true"

	op, err := a.gcpClient.CreateEdgeCacheService(ctx, &pb.CreateEdgeCacheServiceRequest{
		Parent:             parentName,
		EdgeCacheServiceId: *resourceID,
		EdgeCacheService:   desired,
	})
	if err != nil {
		return fmt.Errorf("creating EdgeCacheService %s: %w", name, err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for EdgeCacheService %s creation: %w", name, err)
	}
	log.V(2).Info("successfully created EdgeCacheService", "name", name)

	// Fetch the created object to get the status
	created, err := a.gcpClient.GetEdgeCacheService(ctx, &pb.GetEdgeCacheServiceRequest{Name: name})
	if err != nil {
		return fmt.Errorf("getting created EdgeCacheService %s: %w", name, err)
	}

	status := &krm.NetworkServicesEdgeCacheServiceStatus{}
	status.ObservedState = EdgeCacheServiceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(name)
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *EdgeCacheServiceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)

	project := a.desired.Spec.ProjectRef.External
	location := a.desired.Spec.Location
	resourceID := a.desired.Spec.ResourceID
	if resourceID == nil {
		resourceID = &a.desired.Name
	}
	name := fmt.Sprintf("projects/%s/locations/%s/edgeCacheServices/%s", project, location, *resourceID)

	log.V(2).Info("updating EdgeCacheService", "name", name)

	mapCtx := &direct.MapContext{}
	desired := EdgeCacheServiceSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if desired.Labels == nil {
		desired.Labels = make(map[string]string)
	}
	desired.Labels["managed-by-cnrm"] = "true"
	desired.Name = name

	// Compute update mask
	updateMask := &fieldmaskpb.FieldMask{Paths: []string{
		"description", "labels", "routing", "disable_quic", "disable_http2",
		"require_tls", "edge_ssl_certificates", "edge_security_policy", "log_config",
	}}

	op, err := a.gcpClient.UpdateEdgeCacheService(ctx, &pb.UpdateEdgeCacheServiceRequest{
		EdgeCacheService: desired,
		UpdateMask:       updateMask,
	})
	if err != nil {
		return fmt.Errorf("updating EdgeCacheService %s: %w", name, err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for EdgeCacheService %s update: %w", name, err)
	}
	log.V(2).Info("successfully updated EdgeCacheService", "name", name)

	updated, err := a.gcpClient.GetEdgeCacheService(ctx, &pb.GetEdgeCacheServiceRequest{Name: name})
	if err != nil {
		return fmt.Errorf("getting updated EdgeCacheService %s: %w", name, err)
	}

	status := &krm.NetworkServicesEdgeCacheServiceStatus{}
	status.ObservedState = EdgeCacheServiceObservedState_FromProto(mapCtx, updated)
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
	obj.Spec = direct.ValueOf(EdgeCacheServiceSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	parts := strings.Split(a.actual.Name, "/")
	if len(parts) == 6 {
		obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
			ProjectRef: &refs.ProjectRef{External: parts[1]},
			Location:   parts[3],
		}
		obj.Spec.ResourceID = &parts[5]
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(direct.ValueOf(obj.Spec.ResourceID))
	u.SetGroupVersionKind(krm.NetworkServicesEdgeCacheServiceGVK)
	u.Object = uObj

	return u, nil
}

func (a *EdgeCacheServiceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)

	project := a.desired.Spec.ProjectRef.External
	location := a.desired.Spec.Location
	resourceID := a.desired.Spec.ResourceID
	if resourceID == nil {
		resourceID = &a.desired.Name
	}
	name := fmt.Sprintf("projects/%s/locations/%s/edgeCacheServices/%s", project, location, *resourceID)

	log.V(2).Info("deleting EdgeCacheService", "name", name)

	op, err := a.gcpClient.DeleteEdgeCacheService(ctx, &pb.DeleteEdgeCacheServiceRequest{Name: name})
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent EdgeCacheService", "name", name)
			return true, nil
		}
		return false, fmt.Errorf("deleting EdgeCacheService %s: %w", name, err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for EdgeCacheService %s deletion: %w", name, err)
	}

	log.V(2).Info("successfully deleted EdgeCacheService", "name", name)
	return true, nil
}

func (a *EdgeCacheServiceAdapter) waitForOperation(ctx context.Context, op *longrunningpb.Operation) error {
	if op.Done {
		if op.GetError() != nil {
			return fmt.Errorf("operation failed: %v", op.GetError())
		}
		return nil
	}
	return nil
}
