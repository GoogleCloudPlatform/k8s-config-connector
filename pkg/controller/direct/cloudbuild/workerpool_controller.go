/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudbuild

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	gcp "cloud.google.com/go/cloudbuild/apiv1/v2"
	cloudbuildpb "cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.GroupVersionKind, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

const ctrlName = "cloudbuild-controller"

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building cloudbuild client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.CloudBuildWorkerPool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	// Get GCP Project
	projectRef, err := refs.ResolveProject(ctx, reader, obj, obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	// Get location
	location := obj.Spec.Location

	var id *CloudBuildWorkerPoolIdentity

	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.project != projectID {
			return nil, fmt.Errorf("CloudBuildWorkerPool %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.project, projectID)
		}
		if id.location != location {
			return nil, fmt.Errorf("CloudBuildWorkerPool %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.location, location)
		}
		// TODO: need to support more cases
		if id.workerpool != resourceID {
			return nil, fmt.Errorf("CloudBuildWorkerPool  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.workerpool, resourceID)
		}
	}

	// Get computeNetwork
	if obj.Spec.PrivatePoolConfig.NetworkConfig != nil {
		networkRef, err := refs.ResolveComputeNetwork(ctx, reader, obj, &obj.Spec.PrivatePoolConfig.NetworkConfig.PeeredNetworkRef)
		if err != nil {
			return nil, err

		}
		obj.Spec.PrivatePoolConfig.NetworkConfig.PeeredNetworkRef.External = networkRef.String()
	}

	// Get CloudBuild GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //cloudbuild.googleapis.com/projects/<project>/lcoations/<location>/workerPools/<id>
	if !strings.HasPrefix(url, "//cloudbuild.googleapis.com/") {
		return nil, nil
	}

	tokens := strings.Split(strings.TrimPrefix(url, "//cloudbuild.googleapis.com/"), "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "workerPools" {
		// Get CloudBuild GCP client
		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}

		return &Adapter{
			id: &CloudBuildWorkerPoolIdentity{
				project:    tokens[1],
				location:   tokens[3],
				workerpool: tokens[5],
			},
			gcpClient: gcpClient,
		}, nil
	}

	return nil, nil
}

type Adapter struct {
	id        *CloudBuildWorkerPoolIdentity
	gcpClient *gcp.Client
	desired   *krm.CloudBuildWorkerPool
	actual    *cloudbuildpb.WorkerPool
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	req := &cloudbuildpb.GetWorkerPoolRequest{Name: a.id.FullyQualifiedName()}
	workerpoolpb, err := a.gcpClient.GetWorkerPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting cloudbuildworkerpool %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = workerpoolpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating object", "u", u)

	desired := a.desired.DeepCopy()

	mapCtx := &direct.MapContext{}
	wp := CloudBuildWorkerPoolSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	wp.Name = a.id.FullyQualifiedName()
	req := &cloudbuildpb.CreateWorkerPoolRequest{
		Parent:       a.id.Parent(),
		WorkerPoolId: a.id.workerpool,
		WorkerPool:   wp,
	}
	op, err := a.gcpClient.CreateWorkerPool(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudbuildworkerpool %s creating failed: %w", wp.Name, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("cloudbuildworkerpool %s waiting creation failed: %w", wp.Name, err)
	}

	status := &krm.CloudBuildWorkerPoolStatus{}
	status.ObservedState = CloudBuildWorkerPoolObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	updateMask := &fieldmaskpb.FieldMask{}

	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}

	typedConfig, ok := a.actual.Config.(*cloudbuildpb.WorkerPool_PrivatePoolV1Config)
	if !ok {
		return fmt.Errorf("unable to convert cloudbuildworkerpool %s config to workerpool PrivatePoolV1Config", a.actual.Name)
	}
	actualConfig := typedConfig.PrivatePoolV1Config
	desiredConfig := a.desired.Spec.PrivatePoolConfig

	if desiredConfig.NetworkConfig != nil {
		switch actualConfig.NetworkConfig.EgressOption {
		case cloudbuildpb.PrivatePoolV1Config_NetworkConfig_EGRESS_OPTION_UNSPECIFIED:
			if !reflect.DeepEqual(direct.ValueOf(desiredConfig.NetworkConfig.EgressOption), "UNSPECIFIED") {
				updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.network_config.egress_option")
			}
		case cloudbuildpb.PrivatePoolV1Config_NetworkConfig_NO_PUBLIC_EGRESS:
			if !reflect.DeepEqual(direct.ValueOf(desiredConfig.NetworkConfig.EgressOption), "NO_PUBLIC_EGRESS") {
				updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.network_config.egress_option")
			}
		case cloudbuildpb.PrivatePoolV1Config_NetworkConfig_PUBLIC_EGRESS:
			if !reflect.DeepEqual(direct.ValueOf(desiredConfig.NetworkConfig.EgressOption), "PUBLIC_EGRESS") {
				updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.network_config.egress_option")
			}
		}
		expectedIPRange := direct.ValueOf(desiredConfig.NetworkConfig.PeeredNetworkIPRange)
		if expectedIPRange != "" && !reflect.DeepEqual(expectedIPRange, actualConfig.NetworkConfig.PeeredNetworkIpRange) {
			updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.network_config.peered_network_ip_range")
		}

		// TODO: better handle the network complexity
		// 1. peered_network is an immutable field. whether/when shall we validate
		// 2. the gcp workerpool stores the network with "project_number", different from the spec which uses the "project_id".
		//    * projects/<project_number>/global/networks/<network_id>
		//    * projects/<project_id>/global/networks/<network_id>
		desiredNetwork := strings.Split(desiredConfig.NetworkConfig.PeeredNetworkRef.External, "/")
		actualNetwork := strings.Split(actualConfig.NetworkConfig.PeeredNetwork, "/")
		if len(desiredNetwork) == 5 && len(actualNetwork) == 5 && !reflect.DeepEqual(desiredNetwork[4], actualNetwork[4]) {
			return fmt.Errorf("peered_network is immutable field")
		}
	}
	if !reflect.DeepEqual(desiredConfig.WorkerConfig.DiskSizeGb, actualConfig.WorkerConfig.DiskSizeGb) {
		updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.worker_config.disk_size_gb")
	}
	if !reflect.DeepEqual(desiredConfig.WorkerConfig.MachineType, actualConfig.WorkerConfig.MachineType) {
		updateMask.Paths = append(updateMask.Paths, "private_pool_v1_config.worker_config.machine_type")
	}
	if len(updateMask.Paths) == 0 {
		klog.Warningf("unexpected empty update mask, desired: %v, actual: %v", a.desired, a.actual)
		return nil
	}

	desired := a.desired.DeepCopy()
	mapCtx := &direct.MapContext{}
	wp := CloudBuildWorkerPoolSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	wp.Name = a.id.FullyQualifiedName()
	wp.Etag = a.actual.Etag
	req := &cloudbuildpb.UpdateWorkerPoolRequest{
		WorkerPool: wp,
		UpdateMask: updateMask,
	}
	op, err := a.gcpClient.UpdateWorkerPool(ctx, req)
	if err != nil {
		return fmt.Errorf("cloudbuildworkerpool %s updating failed: %w", wp.Name, err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("cloudbuildworkerpool %s waiting update failed: %w", wp.Name, err)
	}
	status := &krm.CloudBuildWorkerPoolStatus{}
	status.ObservedState = CloudBuildWorkerPoolObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return fmt.Errorf("update workerpool status %w", mapCtx.Err())
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudBuildWorkerPool{}
	obj.SetGroupVersionKind(krm.GroupVersionKind)
	obj.SetName(a.actual.Name)

	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudBuildWorkerPoolSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: *a.id.AsExternalRef()}
	obj.Spec.Location = a.id.location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
// TODO: Delete can rely on status.externalRef and do not need spec.projectRef.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &cloudbuildpb.DeleteWorkerPoolRequest{Name: a.id.FullyQualifiedName(), AllowMissing: true}
	op, err := a.gcpClient.DeleteWorkerPool(ctx, req)
	if err != nil {
		// likely a server bug. worker_pool can be successfully deleted.
		if !strings.Contains(err.Error(), "(line 12:3): missing \"value\" field") {
			return false, fmt.Errorf("deleting cloudbuildworkerpool %s: %w", a.id.FullyQualifiedName(), err)
		}
	}
	err = op.Wait(ctx)
	if err != nil {
		// likely a server bug. worker_pool can be successfully deleted.
		if !strings.Contains(err.Error(), "(line 12:3): missing \"value\" field") {
			return false, fmt.Errorf("waiting delete cloudbuildworkerpool %s: %w", a.id.FullyQualifiedName(), err)
		}
	}
	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
