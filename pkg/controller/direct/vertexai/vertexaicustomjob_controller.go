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

package vertexai

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"

	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.VertexAICustomJobGVK, NewCustomJobModel)
}

func NewCustomJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCustomJob{config: *config}, nil
}

var _ directbase.Model = &modelCustomJob{}

type modelCustomJob struct {
	config config.ControllerConfig
}

func (m *modelCustomJob) client(ctx context.Context, location string) (*gcp.JobClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	aiplatformurl := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(aiplatformurl))
	gcpClient, err := gcp.NewJobRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error building CustomJob client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCustomJob) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAICustomJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.VertexAICustomJobIdentity)

	// Always call common.NormalizeReferences to resolve any resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if obj.Spec.JobSpec != nil && obj.Spec.JobSpec.ServiceAccountRef != nil {
		if err := obj.Spec.JobSpec.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
			return nil, fmt.Errorf("resolving serviceAccountRef: %w", err)
		}
	}

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &CustomJobAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelCustomJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type CustomJobAdapter struct {
	id        *krm.VertexAICustomJobIdentity
	gcpClient *gcp.JobClient
	desired   *krm.VertexAICustomJob
	actual    *pb.CustomJob
}

var _ directbase.Adapter = &CustomJobAdapter{}

func (a *CustomJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CustomJob", "name", a.id)

	req := &pb.GetCustomJobRequest{Name: a.id.String()}
	customJob, err := a.gcpClient.GetCustomJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) || direct.IsBadRequest(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CustomJob %q: %w", a.id, err)
	}

	a.actual = customJob
	return true, nil
}

func (a *CustomJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CustomJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := VertexAICustomJobSpec_ToProto(mapCtx, &desired.Spec)

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateCustomJobRequest{
		Parent:    a.id.ParentString(),
		CustomJob: resource,
	}
	created, err := a.gcpClient.CreateCustomJob(ctx, req)
	if err != nil {
		return fmt.Errorf("CustomJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CustomJob", "name", a.id)

	status := &krm.VertexAICustomJobStatus{}
	status.ObservedState = VertexAICustomJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *CustomJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CustomJob (checking for mutations)", "name", a.id)

	if a.actual == nil {
		return fmt.Errorf("actual state not found for CustomJob %s", a.id)
	}

	desired := a.desired.DeepCopy()
	mapCtx := &direct.MapContext{}
	desiredProto := VertexAICustomJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, err := a.compare(ctx, a.actual, desiredProto)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("VertexAICustomJob is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *CustomJobAdapter) compare(ctx context.Context, actual, desired *pb.CustomJob) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VertexAICustomJobSpec_FromProto, VertexAICustomJobSpec_ToProto)
	if err != nil {
		return nil, err
	}
	diffs, _, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func (a *CustomJobAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CustomJob) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAICustomJobStatus{}
	status.ObservedState = VertexAICustomJobObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(latest.Name)
	return op.UpdateStatus(ctx, status, nil)
}

func (a *CustomJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAICustomJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAICustomJobSpec_FromProto(mapCtx, a.actual))
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
	u.SetGroupVersionKind(krm.VertexAICustomJobGVK)

	u.Object = uObj
	return u, nil
}

func (a *CustomJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CustomJob", "name", a.id)

	req := &pb.DeleteCustomJobRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCustomJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) || direct.IsBadRequest(err) {
			log.V(2).Info("skipping delete for non-existent CustomJob, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CustomJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully initiated deletion of CustomJob", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete CustomJob %s: %w", a.id, err)
	}
	return true, nil
}
