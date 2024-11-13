// Copyright 2024 Google LLC
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

package containerattached

import (
	"context"
	// "crypto/tls"
	"errors"
	"fmt"
	"log"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/containerattached/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/gkemulticloud/apiv1"
	cloudresourcemanager "cloud.google.com/go/resourcemanager/apiv3"

	containerattachedpb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	"github.com/googleapis/gax-go/v2/apierror"
	"google.golang.org/api/option"
	// "google.golang.org/grpc/credentials"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ContainerAttachedClusterGVK, NewContainerAttachedClusterModel)
}

func NewContainerAttachedClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelContainerAttachedCluster{config: *config}, nil
}

var _ directbase.Model = &modelContainerAttachedCluster{}

type modelContainerAttachedCluster struct {
	config config.ControllerConfig
}

func loggingUnaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			err := invoker(ctx, method, req, reply, cc, opts...)
			log.Printf(":HF: Invoked method: %v, %+v", method, err)
			md, ok := metadata.FromOutgoingContext(ctx)
			log.Printf(":HF: md: %+v, ok: %t", md, ok)
			if ok {
					log.Println("Metadata:")
					for k, v := range md {
							log.Printf("Key: %v, Value: %v", k, v)
					}
			}
			reqb, merr := protojson.Marshal(req.(protoreflect.ProtoMessage))
			if merr == nil {
					log.Printf(":HF: Request: %s", reqb)
			}
			log.Printf(":HF: merr: %+v", merr)
			return err
	}
}

func (m *modelContainerAttachedCluster) client(ctx context.Context, endpoint string) (*gcp.AttachedClustersClient, error) {
	var opts []option.ClientOption
	// Not working ("WithHTTPClient is incompatible with gRPC dial options"). GRPCClientOptions() gives the same error, and
	// the implementation looks functionally identical.
	// opts, err := m.config.GRPCClientOptions()
	// if err != nil {
	// 	return nil, err
	// }

	// Added an interceptor to add more logging, but couldn't find anything useful.
	// opts = append(opts, option.WithGRPCDialOption(grpc.WithUnaryInterceptor(loggingUnaryInterceptor())))
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewAttachedClustersClient(ctx, opts...)
	// gcpClient, err := gcp.NewAttachedClustersClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("building AttachedCluster client: %w", err)
	}
	return gcpClient, err
}

func (m *modelContainerAttachedCluster) projectsClient(ctx context.Context) (*cloudresourcemanager.ProjectsClient, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	crmClient, err := cloudresourcemanager.NewProjectsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building cloudresourcemanager client: %w", err)
	}
	return crmClient, err
}

func endpoint(location string) string {
	return fmt.Sprintf("%s-gkemulticloud.googleapis.com:443", location)
}

func (m *modelContainerAttachedCluster) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	log.V(0).Info(":HF: AdapterForObject")
	obj := &krm.ContainerAttachedCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewContainerAttachedClusterRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get project number
	project, err := refs.ResolveProject(ctx, reader, u, &obj.Spec.Fleet.ProjectRef)
	if err != nil {
					return nil, err
	}
	projectNumber := project.ProjectNumber
	obj.Spec.Fleet.ProjectRef.External = fmt.Sprintf("projects/%v", projectNumber)
	// projectsClient, err := m.projectsClient(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// // err = krm.ResolveFleetProjectRef(ctx, reader, &obj.Spec.Fleet.ProjectRef, obj)
	// err = obj.Spec.Fleet.ProjectRef.ResolveExternal(ctx, projectsClient)
	// if err != nil {
	// 	// return nil, err
	// 	log.V(0).Info(fmt.Sprintf("HF: ResolveFleetProjectRef error: %v", err))
	// }
	// log.V(0).Info(fmt.Sprintf("HF: ResolveFleetProjectRef: %v", obj.Spec.Fleet.ProjectRef))

	// Get containerattached GCP client
	endpoint := endpoint(id.Location)
	gcpClient, err := m.client(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	return &ContainerAttachedClusterAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelContainerAttachedCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	log.V(0).Info(":HF: AdapterForURL")
	// TODO: Support URLs
	return nil, nil
}

type ContainerAttachedClusterAdapter struct {
	id        *krm.ContainerAttachedClusterRef
	gcpClient *gcp.AttachedClustersClient
	desired   *krm.ContainerAttachedCluster
	actual    *containerattachedpb.AttachedCluster
}

var _ directbase.Adapter = &ContainerAttachedClusterAdapter{}

func (a *ContainerAttachedClusterAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	//log.V(2).Info("getting ContainerAttachedCluster", "name", a.id.External)
	log.V(0).Info(":HF: getting ContainerAttachedCluster", "name", a.id.External, "id", a.id)

	req := &containerattachedpb.GetAttachedClusterRequest{Name: a.id.External}
	attachedclusterpb, err := a.gcpClient.GetAttachedCluster(ctx, req)
	if err != nil {
		log.V(0).Info(":HF: getting ContainerAttachedCluster: error", "err", err)

		var ae *apierror.APIError
		if errors.As(err, &ae) {
				log.V(0).Info(ae.Reason())
				log.V(0).Info(":HF:", "help", ae.Details().Help.GetLinks())
		}
		if s, ok := status.FromError(err); ok {
      log.V(0).Info(s.Message())
      for _, d := range s.Proto().Details {
         log.V(0).Info(":HF:", "d", d)
      }
   }

		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ContainerAttachedCluster %q: %w", a.id.External, err)
	}

	log.V(0).Info(":HF: getting ContainerAttachedCluster: success", "actual", attachedclusterpb)
	a.actual = attachedclusterpb
	return true, nil
}

func (a *ContainerAttachedClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	// log.V(2).Info("creating AttachedCluster", "name", a.id.Name)
	log.V(0).Info(":HF: creating AttachedCluster", "name", a.id.Name, "id", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ContainerAttachedClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(user): Complete the gcp "CREATE" or "INSERT" request with required fields.
	parent, err := a.id.Parent()
	if err != nil {
		return err
	}
	req := &containerattachedpb.CreateAttachedClusterRequest{
		Parent:          parent,
		AttachedCluster: resource,
		AttachedClusterId: a.id.Name,
	}
	op, err := a.gcpClient.CreateAttachedCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AttachedCluster %s: %w", a.id.External, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AttachedCluster %s waiting creation: %w", a.id.External, err)
	}
	log.V(2).Info("successfully created AttachedCluster", "name", a.id.External)

	status := &krm.ContainerAttachedClusterStatus{}
	status.ObservedState = ContainerAttachedClusterStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ContainerAttachedClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	// log.V(2).Info("updating AttachedCluster", "name", a.id.External)
	log.V(0).Info(":HF: updating AttachedCluster", "name", a.id.Name, "id", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ContainerAttachedClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// :TODO: Should ToProto be filling out the name?
	resource.Name = a.id.External

	// Mask of fields to update. At least one path must be supplied in
	// this field.
	//
	//   - `annotations`.
	//   - `authorization.admin_users`.
	//   - `binary_authorization.evaluation_mode`.
	//   - `description`.
	//   - `logging_config.component_config.enable_components`.
	//   - `monitoring_config.managed_prometheus_config.enabled`.
	//   - `platform_version`.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.Spec.Annotations, a.actual.Annotations) {
		updateMask.Paths = append(updateMask.Paths, "annotations")
	}
	if !reflect.DeepEqual(a.desired.Spec.Authorization.AdminUsers, a.actual.Authorization.AdminUsers) {
		updateMask.Paths = append(updateMask.Paths, "authorization.admin_users")
	}
	if !reflect.DeepEqual(a.desired.Spec.BinaryAuthorization.EvaluationMode, a.actual.BinaryAuthorization.EvaluationMode) {
		updateMask.Paths = append(updateMask.Paths, "binary_authorization.evaluation_mode")
	}
	if !reflect.DeepEqual(a.desired.Spec.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(a.desired.Spec.LoggingConfig.ComponentConfig.EnableComponents, a.actual.LoggingConfig.ComponentConfig.EnableComponents) {
		updateMask.Paths = append(updateMask.Paths, "logging_config.component_config.enable_components")
	}
	if !reflect.DeepEqual(a.desired.Spec.MonitoringConfig.ManagedPrometheusConfig.Enabled, a.actual.MonitoringConfig.ManagedPrometheusConfig.Enabled) {
		updateMask.Paths = append(updateMask.Paths, "monitoring_config.managed_prometheus_config.enabled")
	}
	if !reflect.DeepEqual(a.desired.Spec.PlatformVersion, a.actual.PlatformVersion) {
		updateMask.Paths = append(updateMask.Paths, "platform_version")
	}
	log.V(0).Info(":HF: updating AttachedCluster: resource", "resource", resource, "mask", updateMask, "actual", a.actual)

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		return nil
	}
	req := &containerattachedpb.UpdateAttachedClusterRequest{
		UpdateMask:      updateMask,
		AttachedCluster: resource,
	}
	op, err := a.gcpClient.UpdateAttachedCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AttachedCluster %s: %w", a.id.External, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AttachedCluster %s waiting update: %w", a.id.External, err)
	}
	log.V(2).Info("successfully updated AttachedCluster", "name", a.id.External)

	status := &krm.ContainerAttachedClusterStatus{}
	status.ObservedState = ContainerAttachedClusterStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ContainerAttachedClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ContainerAttachedCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ContainerAttachedClusterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent, err := a.id.Parent()
	if err != nil {
		return nil, err
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: parent}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Name)
	u.SetGroupVersionKind(krm.ContainerAttachedClusterGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *ContainerAttachedClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AttachedCluster", "name", a.id.External)

	req := &containerattachedpb.DeleteAttachedClusterRequest{Name: a.id.External}
	op, err := a.gcpClient.DeleteAttachedCluster(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting AttachedCluster %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully deleted AttachedCluster", "name", a.id.External)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete AttachedCluster %s: %w", a.id.External, err)
	}
	return true, nil
}
