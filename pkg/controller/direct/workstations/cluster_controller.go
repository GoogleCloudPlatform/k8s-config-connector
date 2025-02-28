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

package workstations

import (
	"context"
	"fmt"
	"reflect"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"

	gcp "cloud.google.com/go/workstations/apiv1"
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "workstations-controller"
)

func init() {
	registry.RegisterModel(krm.WorkstationClusterGVK, NewWorkstationClusterModel)
	fuzztesting.RegisterKRMFuzzer(workstationclusterFuzzer())
}

func workstationclusterFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.WorkstationCluster{},
		WorkstationClusterSpec_FromProto, WorkstationClusterSpec_ToProto,
		WorkstationClusterObservedState_FromProto, WorkstationClusterObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")

	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".reconciling")
	f.UnimplementedFields.Insert(".degraded")
	f.UnimplementedFields.Insert(".conditions")
	f.UnimplementedFields.Insert(".private_cluster_config.cluster_hostname")
	f.UnimplementedFields.Insert(".private_cluster_config.service_attachment_uri")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".private_cluster_config")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".subnetwork")
	f.SpecFields.Insert(".network")

	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".control_plane_ip")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".uid")

	return f
}

func NewWorkstationClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkstationCluster{config: *config}, nil
}

var _ directbase.Model = &modelWorkstationCluster{}

type modelWorkstationCluster struct {
	config config.ControllerConfig
}

func (m *modelWorkstationCluster) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building WorkstationCluster client: %w", err)
	}
	return gcpClient, err
}

func (m *modelWorkstationCluster) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.WorkstationCluster{}
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

	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	// Get location
	location := obj.Spec.Location

	var id *WorkstationClusterIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(projectID, location, resourceID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.Project != projectID {
			return nil, fmt.Errorf("WorkstationCluster %s/%s has spec.projectRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Project, projectID)
		}
		if id.Parent.Location != location {
			return nil, fmt.Errorf("WorkstationCluster %s/%s has spec.location changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Location, location)
		}
		if id.WorkstationCluster != resourceID {
			return nil, fmt.Errorf("WorkstationCluster  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.WorkstationCluster, resourceID)
		}
	}

	if err := NormalizeWorkstationCluster(ctx, reader, obj); err != nil {
		return nil, err
	}

	// Get workstations GCP client
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

func (m *modelWorkstationCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *WorkstationClusterIdentity
	gcpClient *gcp.Client
	desired   *krm.WorkstationCluster
	actual    *pb.WorkstationCluster
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting WorkstationCluster", "name", a.id.FullyQualifiedName())

	req := &pb.GetWorkstationClusterRequest{Name: a.id.FullyQualifiedName()}
	workstationclusterpb, err := a.gcpClient.GetWorkstationCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WorkstationCluster %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = workstationclusterpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating WorkstationCluster", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkstationClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateWorkstationClusterRequest{
		Parent:               a.id.Parent.String(),
		WorkstationClusterId: a.id.WorkstationCluster,
		WorkstationCluster:   resource,
	}
	op, err := a.gcpClient.CreateWorkstationCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WorkstationCluster %s: %w", a.id.FullyQualifiedName(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkstationCluster %s waiting creation: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created WorkstationCluster", "name", a.id.FullyQualifiedName())

	status := &krm.WorkstationClusterStatus{}
	status.ObservedState = WorkstationClusterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating WorkstationCluster", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkstationClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.FullyQualifiedName()

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(resource.Annotations, a.actual.Annotations) {
		updateMask.Paths = append(updateMask.Paths, "annotations")
	}
	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}

	req := &pb.UpdateWorkstationClusterRequest{
		UpdateMask:         updateMask,
		WorkstationCluster: resource,
	}
	op, err := a.gcpClient.UpdateWorkstationCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating WorkstationCluster %s: %w", a.id.FullyQualifiedName(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkstationCluster %s waiting update: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated WorkstationCluster", "name", a.id.FullyQualifiedName())

	status := &krm.WorkstationClusterStatus{}
	status.ObservedState = WorkstationClusterObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.WorkstationCluster{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(WorkstationClusterSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = refs.ProjectRef{Name: a.id.Parent.Project}
	obj.Spec.Location = a.id.Parent.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting WorkstationCluster", "name", a.id.FullyQualifiedName())

	req := &pb.DeleteWorkstationClusterRequest{Name: a.id.FullyQualifiedName()}
	op, err := a.gcpClient.DeleteWorkstationCluster(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting WorkstationCluster %s: %w", a.id.FullyQualifiedName(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		// todo (b/368419476): Workstation service does not provide a valid response on success.
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting delete WorkstationCluster %s: %w", a.id.FullyQualifiedName(), err)
		}
	}

	log.V(2).Info("successfully deleted WorkstationCluster", "name", a.id.FullyQualifiedName())
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
