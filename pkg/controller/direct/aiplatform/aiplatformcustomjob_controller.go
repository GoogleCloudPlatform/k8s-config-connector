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

package aiplatform

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformCustomJobGVK, NewAIPlatformCustomJobModel)
}

func NewAIPlatformCustomJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCustomJob{config: config}, nil
}

type modelCustomJob struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &modelCustomJob{}

func (m *modelCustomJob) client(ctx context.Context, location string) (*gcp.JobClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewJobClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CustomJob client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelCustomJob) AdapterForObject(ctx context.Context, reader *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := reader.Object
	obj := &krm.AIPlatformCustomJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader.Reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.AIPlatformCustomJobIdentity)

	// Always call common.NormalizeReferences to resolve any resource references
	if err := common.NormalizeReferences(ctx, reader.Reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// IAMServiceAccountRef does not implement the refs.Ref interface, so resolve explicitly
	if obj.Spec.JobSpec != nil && obj.Spec.JobSpec.ServiceAccountRef != nil {
		if err := obj.Spec.JobSpec.ServiceAccountRef.Resolve(ctx, reader.Reader, obj); err != nil {
			return nil, fmt.Errorf("resolving serviceAccountRef: %w", err)
		}
	}

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
	id        *krm.AIPlatformCustomJobIdentity
	gcpClient *gcp.JobClient
	desired   *krm.AIPlatformCustomJob
	actual    *pb.CustomJob
}

var _ directbase.Adapter = &CustomJobAdapter{}

func (a *CustomJobAdapter) Find(ctx context.Context) (bool, error) {
	// Rule 2: Find() Pre-Check Guard for service-generated ID
	if a.id.CustomJob == "" {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("getting CustomJob", "name", a.id)

	req := &pb.GetCustomJobRequest{Name: a.id.String()}
	customJob, err := a.gcpClient.GetCustomJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = customJob
	return true, nil
}

func (a *CustomJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CustomJob", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := AIPlatformCustomJobSpec_ToProto(mapCtx, &desired.Spec)
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

	return a.updateStatus(ctx, createOp, created)
}

func (a *CustomJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CustomJob (checking for mutations)", "name", a.id)

	if a.actual == nil {
		return fmt.Errorf("actual state not found for CustomJob %s", a.id)
	}

	_, diffs, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformCustomJobSpec_FromProto,
		AIPlatformCustomJobSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("AIPlatformCustomJob is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())
	}

	log.V(2).Info("no mutations detected on immutable CustomJob", "name", a.id)
	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *CustomJobAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CustomJob) error {
	mapCtx := &direct.MapContext{}
	status := &krm.AIPlatformCustomJobStatus{}
	status.ObservedState = AIPlatformCustomJobObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if err := a.id.FromExternal(latest.Name); err != nil {
		return fmt.Errorf("parsing CustomJob name %q: %w", latest.Name, err)
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *CustomJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("CustomJob %s not found", a.id)
	}

	mapCtx := &direct.MapContext{}
	obj := &krm.AIPlatformCustomJob{}
	obj.Spec = direct.ValueOf(AIPlatformCustomJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = &a.id.Location
	obj.Spec.ResourceID = &a.id.CustomJob

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.CustomJob)
	u.SetGroupVersionKind(krm.AIPlatformCustomJobGVK)

	return u, nil
}

func (a *CustomJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CustomJob", "name", a.id)

	if a.id.CustomJob == "" {
		return true, nil
	}

	// 1. Get the latest CustomJob state before deletion
	getReq := &pb.GetCustomJobRequest{
		Name: a.id.String(),
	}
	actual, err := a.gcpClient.GetCustomJob(ctx, getReq)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("getting CustomJob before deletion %q: %w", a.id.String(), err)
	}

	// 2. If the state is not terminal, initiate cancellation and wait
	isTerminal := func(state pb.JobState) bool {
		return state == pb.JobState_JOB_STATE_SUCCEEDED ||
			state == pb.JobState_JOB_STATE_FAILED ||
			state == pb.JobState_JOB_STATE_CANCELLED
	}

	if !isTerminal(actual.State) {
		if actual.State != pb.JobState_JOB_STATE_CANCELLING {
			log.V(2).Info("cancelling AIPlatformCustomJob before deletion", "name", a.id.String(), "currentState", actual.State)
			cancelReq := &pb.CancelCustomJobRequest{
				Name: a.id.String(),
			}
			if err := a.gcpClient.CancelCustomJob(ctx, cancelReq); err != nil {
				log.Error(err, "error cancelling AIPlatformCustomJob", "name", a.id.String())
				return false, fmt.Errorf("cancelling AIPlatformCustomJob %q: %w", a.id.String(), err)
			}
		}
		return false, fmt.Errorf("AIPlatformCustomJob %q is in state %s; waiting for terminal state before deletion", a.id.String(), actual.State)
	}

	// 3. Delete the CustomJob once in a terminal state
	req := &pb.DeleteCustomJobRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteCustomJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CustomJob, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting CustomJob %s: %w", a.id, err)
	}
	log.V(2).Info("waiting for CustomJob deletion operation", "name", a.id)
	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for CustomJob %s deletion: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CustomJob", "name", a.id)
	return true, nil
}
