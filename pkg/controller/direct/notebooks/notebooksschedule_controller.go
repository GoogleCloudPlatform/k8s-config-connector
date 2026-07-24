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

package notebooks

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/notebooks/apiv1"
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	dataprocv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NotebooksScheduleGVK, NewScheduleModel)
}

func NewScheduleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSchedule{config: *config}, nil
}

var _ directbase.Model = &modelSchedule{}

type modelSchedule struct {
	config config.ControllerConfig
}

func (m *modelSchedule) client(ctx context.Context) (*gcp.NotebookClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewNotebookClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Schedule client: %w", err)
	}
	return gcpClient, err
}

func (m *modelSchedule) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NotebooksSchedule{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Manually resolve IAMServiceAccountRef
	if obj.Spec.ExecutionTemplate != nil && obj.Spec.ExecutionTemplate.ServiceAccountRef != nil {
		if err := obj.Spec.ExecutionTemplate.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
			return nil, fmt.Errorf("resolving serviceAccountRef: %w", err)
		}
	}

	// Manually normalize ClusterRef
	if obj.Spec.ExecutionTemplate != nil && obj.Spec.ExecutionTemplate.DataprocParameters != nil && obj.Spec.ExecutionTemplate.DataprocParameters.ClusterRef != nil {
		if err := obj.Spec.ExecutionTemplate.DataprocParameters.ClusterRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return nil, fmt.Errorf("normalizing clusterRef: %w", err)
		}

		id := &dataprocv1beta1.DataprocClusterIdentity{}
		if err := id.FromExternal(obj.Spec.ExecutionTemplate.DataprocParameters.ClusterRef.External); err == nil {
			obj.Spec.ExecutionTemplate.DataprocParameters.ClusterRef.External = dataprocv1beta1.DataprocClusterIdentityFormatRelative.ToString(*id)
		}
	}

	id, err := krm.NewScheduleIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get notebooks GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := NotebooksScheduleSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &ScheduleAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelSchedule) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ScheduleAdapter struct {
	id        *krm.NotebooksScheduleIdentity
	gcpClient *gcp.NotebookClient
	desired   *pb.Schedule
	actual    *pb.Schedule
}

var _ directbase.Adapter = &ScheduleAdapter{}

func (a *ScheduleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Schedule", "name", a.id)

	req := &pb.GetScheduleRequest{Name: a.id.String()}
	schedulepb, err := a.gcpClient.GetSchedule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Schedule %q: %w", a.id, err)
	}

	a.actual = schedulepb
	return true, nil
}

func (a *ScheduleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Schedule", "name", a.id)

	req := &pb.CreateScheduleRequest{
		Parent:     a.id.ParentString(),
		ScheduleId: a.id.ID(),
		Schedule:   a.desired,
	}
	op, err := a.gcpClient.CreateSchedule(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Schedule %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for Schedule %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Schedule", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *ScheduleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Schedule", "name", a.id)

	diffs, _, err := compareSchedule(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no diff detected for Schedule", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	return fmt.Errorf("NotebooksSchedule resource is immutable and cannot be updated")
}

func (a *ScheduleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NotebooksSchedule{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NotebooksScheduleSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.NotebooksScheduleGVK)

	u.Object = uObj
	return u, nil
}

func (a *ScheduleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Schedule", "name", a.id)

	req := &pb.DeleteScheduleRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteSchedule(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent Schedule, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Schedule %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Schedule", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete Schedule %s: %w", a.id, err)
	}
	return true, nil
}

func compareSchedule(ctx context.Context, actual, desired *pb.Schedule) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, NotebooksScheduleSpec_v1alpha1_FromProto, NotebooksScheduleSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.Schedule) {
		// Even if empty, it's a good pattern to define and populate GCP/server defaults here
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *ScheduleAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Schedule) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NotebooksScheduleStatus{}
	status.ObservedState = NotebooksScheduleObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}
