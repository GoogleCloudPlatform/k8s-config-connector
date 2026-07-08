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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/workstations/apiv1"
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
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
	f.UnimplementedFields.Insert(".domain_config")

	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".reconciling")
	f.UnimplementedFields.Insert(".degraded")
	f.UnimplementedFields.Insert(".conditions")
	f.UnimplementedFields.Insert(".private_cluster_config.cluster_hostname")
	f.UnimplementedFields.Insert(".private_cluster_config.service_attachment_uri")
	f.UnimplementedFields.Insert(".gateway_config")
	f.UnimplementedFields.Insert(".workstation_authorization_url")
	f.UnimplementedFields.Insert(".tags")
	f.UnimplementedFields.Insert(".workstation_launch_url")

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

func (m *modelWorkstationCluster) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.WorkstationCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	workstationClusterID := id.(*krm.WorkstationClusterIdentity)

	if err := NormalizeWorkstationCluster(ctx, reader, obj); err != nil {
		return nil, err
	}

	// Get workstations GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        workstationClusterID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWorkstationCluster) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *krm.WorkstationClusterIdentity
	gcpClient *gcp.Client
	desired   *krm.WorkstationCluster
	actual    *pb.WorkstationCluster
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting WorkstationCluster", "name", a.id.String())

	req := &pb.GetWorkstationClusterRequest{Name: a.id.String()}
	workstationclusterpb, err := a.gcpClient.GetWorkstationCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WorkstationCluster %q: %w", a.id.String(), err)
	}

	a.actual = workstationclusterpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating WorkstationCluster", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkstationClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateWorkstationClusterRequest{
		Parent:               fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location),
		WorkstationClusterId: a.id.WorkstationCluster,
		WorkstationCluster:   resource,
	}
	op, err := a.gcpClient.CreateWorkstationCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WorkstationCluster %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkstationCluster %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created WorkstationCluster", "name", a.id.String())

	status := &krm.WorkstationClusterStatus{}
	status.ObservedState = WorkstationClusterObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating WorkstationCluster", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := WorkstationClusterSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(resource.Annotations, a.actual.Annotations) {
		report.AddField("annotations", a.actual.Annotations, resource.Annotations)
		updateMask.Paths = append(updateMask.Paths, "annotations")
	}
	if !reflect.DeepEqual(resource.Labels, a.actual.Labels) {
		report.AddField("labels", a.actual.Labels, resource.Labels)
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	if len(updateMask.Paths) == 0 {
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateWorkstationClusterRequest{
		UpdateMask:         updateMask,
		WorkstationCluster: resource,
	}
	op, err := a.gcpClient.UpdateWorkstationCluster(ctx, req)
	if err != nil {
		return fmt.Errorf("updating WorkstationCluster %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WorkstationCluster %s waiting update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated WorkstationCluster", "name", a.id.String())

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
	obj.Spec.ProjectRef = refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting WorkstationCluster", "name", a.id.String())

	req := &pb.DeleteWorkstationClusterRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteWorkstationCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent WorkstationCluster, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting WorkstationCluster %s: %w", a.id.String(), err)
	}

	_, err = op.Wait(ctx)
	if err != nil {
		// todo (b/368419476): Workstation service does not provide a valid response on success.
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting delete WorkstationCluster %s: %w", a.id.String(), err)
		}
	}

	log.V(2).Info("successfully deleted WorkstationCluster", "name", a.id.String())
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
