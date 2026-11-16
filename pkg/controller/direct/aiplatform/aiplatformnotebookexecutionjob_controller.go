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

// +generated:controller
// krm.group: aiplatform.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: AIPlatformNotebookExecutionJob

package aiplatform

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/aiplatform/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1"
	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AIPlatformNotebookExecutionJobGVK, NewNotebookExecutionJobModel)
}

func NewNotebookExecutionJobModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelNotebookExecutionJob{config: *config}, nil
}

var _ directbase.Model = &modelNotebookExecutionJob{}

type modelNotebookExecutionJob struct {
	config config.ControllerConfig
}

func (m *modelNotebookExecutionJob) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AIPlatformNotebookExecutionJob{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.AIPlatformNotebookExecutionJobIdentity)

	// Get aiplatform GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &NotebookExecutionJobAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelNotebookExecutionJob) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AIPlatformNotebookExecutionJobIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}
	return &NotebookExecutionJobAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelNotebookExecutionJob) client(ctx context.Context, location string) (*gcp.NotebookClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-aiplatform.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewNotebookClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building NotebookExecutionJob client: %w", err)
	}
	return gcpClient, err
}

type NotebookExecutionJobAdapter struct {
	id        *krm.AIPlatformNotebookExecutionJobIdentity
	gcpClient *gcp.NotebookClient
	desired   *krm.AIPlatformNotebookExecutionJob
	actual    *aiplatformpb.NotebookExecutionJob
}

var _ directbase.Adapter = &NotebookExecutionJobAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *NotebookExecutionJobAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting NotebookExecutionJob", "name", a.id)

	req := &aiplatformpb.GetNotebookExecutionJobRequest{
		Name: a.id.String(),
	}
	notebookexecutionjobpb, err := a.gcpClient.GetNotebookExecutionJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting NotebookExecutionJob %q: %w", a.id, err)
	}

	a.actual = notebookexecutionjobpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NotebookExecutionJobAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating NotebookExecutionJob", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AIPlatformNotebookExecutionJobSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &aiplatformpb.CreateNotebookExecutionJobRequest{
		Parent:                 a.id.ParentString(),
		NotebookExecutionJobId: a.id.NotebookExecutionJob,
		NotebookExecutionJob:   resource,
	}
	op, err := a.gcpClient.CreateNotebookExecutionJob(ctx, req)
	if err != nil {
		return fmt.Errorf("creating NotebookExecutionJob %s: %w", a.id, err)
	}
	_, err = op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("NotebookExecutionJob %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created NotebookExecutionJob", "name", a.id)

	getReq := &aiplatformpb.GetNotebookExecutionJobRequest{
		Name: a.id.String(),
	}
	created, err := a.gcpClient.GetNotebookExecutionJob(ctx, getReq)
	if err != nil {
		return fmt.Errorf("getting created NotebookExecutionJob %s: %w", a.id, err)
	}

	status := &krm.AIPlatformNotebookExecutionJobStatus{}
	status.ObservedState = AIPlatformNotebookExecutionJobObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *NotebookExecutionJobAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NotebookExecutionJob", "name", a.id)

	normalize := func(ctx context.Context, p *aiplatformpb.NotebookExecutionJob) error {
		if p.GetDirectNotebookSource() != nil && len(p.GetDirectNotebookSource().GetContent()) == 0 && a.desired.Spec.DirectNotebookSource != nil {
			p.GetDirectNotebookSource().Content = a.desired.Spec.DirectNotebookSource.Content
		}
		return nil
	}

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AIPlatformNotebookExecutionJobSpec_FromProto,
		AIPlatformNotebookExecutionJobSpec_ToProto,
		normalize,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating AIPlatformNotebookExecutionJob is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.AIPlatformNotebookExecutionJobStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AIPlatformNotebookExecutionJobObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *NotebookExecutionJobAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AIPlatformNotebookExecutionJob{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AIPlatformNotebookExecutionJobSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refsv1beta1.ProjectRef{External: a.id.Project},
		Location:   a.id.Location,
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.NotebookExecutionJob)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.NotebookExecutionJob)
	u.SetGroupVersionKind(krm.AIPlatformNotebookExecutionJobGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *NotebookExecutionJobAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting NotebookExecutionJob", "name", a.id)

	req := &aiplatformpb.DeleteNotebookExecutionJobRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteNotebookExecutionJob(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent NotebookExecutionJob, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting NotebookExecutionJob %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted NotebookExecutionJob", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete NotebookExecutionJob %s: %w", a.id, err)
	}
	return true, nil
}
