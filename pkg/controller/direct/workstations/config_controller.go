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

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
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
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.WorkstationConfigGVK, NewWorkstationConfigModel)
	fuzztesting.RegisterKRMFuzzer(workstationConfigFuzzer())
}

func workstationConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.WorkstationConfig{},
		WorkstationConfigSpec_FromProto, WorkstationConfigSpec_ToProto,
		WorkstationConfigObservedState_FromProto, WorkstationConfigObservedState_ToProto,
	)

	f.UnimplementedFields.Insert(".name")
	f.UnimplementedFields.Insert(".reconciling")
	f.UnimplementedFields.Insert(".conditions")

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".annotations")
	f.SpecFields.Insert(".labels")
	f.SpecFields.Insert(".idle_timeout")
	f.SpecFields.Insert(".running_timeout")
	f.SpecFields.Insert(".host")
	f.SpecFields.Insert(".persistent_directories")
	f.SpecFields.Insert(".container")
	f.SpecFields.Insert(".encryption_key")
	f.SpecFields.Insert(".readiness_checks")
	f.SpecFields.Insert(".replica_zones")

	f.StatusFields.Insert(".uid")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")
	f.StatusFields.Insert(".delete_time")
	f.StatusFields.Insert(".etag")
	f.StatusFields.Insert(".host.gce_instance.pooled_instances")
	f.StatusFields.Insert(".degraded")

	return f
}

func NewWorkstationConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWorkstationConfig{config: *config}, nil
}

var _ directbase.Model = &modelWorkstationConfig{}

type modelWorkstationConfig struct {
	config config.ControllerConfig
}

func (m *modelWorkstationConfig) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building WorkstationConfig client: %w", err)
	}
	return gcpClient, err
}

func (m *modelWorkstationConfig) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.WorkstationConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewWorkstationConfigIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get workstations GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WorkstationConfigAdapter{
		id:        id,
		k8sClient: reader,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWorkstationConfig) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type WorkstationConfigAdapter struct {
	id        *krm.WorkstationConfigIdentity
	k8sClient client.Reader
	gcpClient *gcp.Client
	desired   *krm.WorkstationConfig
	actual    *pb.WorkstationConfig
}

var _ directbase.Adapter = &WorkstationConfigAdapter{}

func (a *WorkstationConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting WorkstationConfig", "name", a.id.String())

	req := &pb.GetWorkstationConfigRequest{Name: a.id.String()}
	workstationconfigpb, err := a.gcpClient.GetWorkstationConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WorkstationConfig %q: %w", a.id.String(), err)
	}

	a.actual = workstationconfigpb
	return true, nil
}

func (a *WorkstationConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating WorkstationConfig", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	// Resolve references
	if err := ResolveWorkstationConfigRefs(ctx, a.k8sClient, desired); err != nil {
		return err
	}
	// Convert to proto
	resource := WorkstationConfigSpec_ToProto(mapCtx, &desired.Spec)
	ApplyWorkstationConfigGCPDefaults(mapCtx, &desired.Spec, resource, nil)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set name manually, because it is not filled-in by WorkstationConfigSpec_ToProto.
	resource.Name = a.id.String()

	req := &pb.CreateWorkstationConfigRequest{
		Parent:              a.id.Parent().String(),
		WorkstationConfigId: a.id.ID(),
		WorkstationConfig:   resource,
	}
	op, err := a.gcpClient.CreateWorkstationConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WorkstationConfig %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting WorkstationConfig %s creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created WorkstationConfig", "name", a.id.String())

	status := &krm.WorkstationConfigStatus{}
	status.ObservedState = WorkstationConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *WorkstationConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating WorkstationConfig", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	// Resolve references
	if err := ResolveWorkstationConfigRefs(ctx, a.k8sClient, desired); err != nil {
		return err
	}
	// Convert to proto
	resource := WorkstationConfigSpec_ToProto(mapCtx, &desired.Spec)
	ApplyWorkstationConfigGCPDefaults(mapCtx, &desired.Spec, resource, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set name and etag manually, because they are not filled-in by WorkstationConfigSpec_ToProto.
	resource.Name = a.id.String()
	resource.Etag = a.actual.Etag

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.WorkstationConfigStatus{}
		status.ObservedState = WorkstationConfigObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateWorkstationConfigRequest{
		WorkstationConfig: resource,
		UpdateMask:        &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}

	op, err := a.gcpClient.UpdateWorkstationConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating WorkstationConfig %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting WorkstationConfig %s update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated WorkstationConfig", "name", a.id.String())

	status := &krm.WorkstationConfigStatus{}
	status.ObservedState = WorkstationConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *WorkstationConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.WorkstationConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(WorkstationConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Parent = &krm.WorkstationClusterRef{External: a.id.Parent().String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.WorkstationConfigGVK)

	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *WorkstationConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting WorkstationConfig", "name", a.id.String())

	req := &pb.DeleteWorkstationConfigRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteWorkstationConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent WorkstationConfig, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting WorkstationConfig %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted WorkstationConfig", "name", a.id.String())

	_, err = op.Wait(ctx)
	if err != nil {
		// todo (b/368419476): Workstation service does not provide a valid response on success.
		if err.Error() != "unsupported result type <nil>: <nil>" {
			return false, fmt.Errorf("waiting delete WorkstationConfig %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}
