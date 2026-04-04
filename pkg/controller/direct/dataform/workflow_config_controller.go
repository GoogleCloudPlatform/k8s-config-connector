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

package dataform

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataform/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/dataform/apiv1beta1"
	dataformpb "cloud.google.com/go/dataform/apiv1beta1/dataformpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.DataformRepositoryWorkflowConfigGVK, NewWorkflowConfigModel)
}

func NewWorkflowConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &workflowConfigModel{config: *config}, nil
}

var _ directbase.Model = &workflowConfigModel{}

type workflowConfigModel struct {
	config config.ControllerConfig
}

func (m *workflowConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataformRepositoryWorkflowConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	repositoryRef, err := obj.Spec.RepositoryRef.FullyQualifiedName(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	repositoryID, err := asID(repositoryRef)
	if err != nil {
		return nil, err
	}

	var id *WorkflowConfigIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildWorkflowConfigID(repositoryID.project, repositoryID.location, repositoryID.dataform, resourceID)
	} else {
		id, err = asWorkflowConfigID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.project != repositoryID.project {
			return nil, fmt.Errorf("DataformRepositoryWorkflowConfig %s/%s has repositoryRef changed (project), expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.project, repositoryID.project)
		}
		if id.location != repositoryID.location {
			return nil, fmt.Errorf("DataformRepositoryWorkflowConfig %s/%s has repositoryRef changed (location), expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.location, repositoryID.location)
		}
		if id.repository != repositoryID.dataform {
			return nil, fmt.Errorf("DataformRepositoryWorkflowConfig %s/%s has repositoryRef changed (repository), expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.repository, repositoryID.dataform)
		}
		if id.workflowConfig != resourceID {
			return nil, fmt.Errorf("DataformRepositoryWorkflowConfig %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.workflowConfig, resourceID)
		}
	}

	if err := resolveWorkflowConfigRefs(ctx, reader, obj); err != nil {
		return nil, err
	}

	gcpModel := &model{config: m.config}
	gcpClient, err := gcpModel.client(ctx)
	if err != nil {
		return nil, err
	}

	return &workflowConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func resolveWorkflowConfigRefs(ctx context.Context, reader client.Reader, obj *krm.DataformRepositoryWorkflowConfig) error {
	releaseConfigRef, err := obj.Spec.ReleaseConfigRef.FullyQualifiedName(ctx, reader, obj.GetNamespace())
	if err != nil {
		return err
	}
	obj.Spec.ReleaseConfigRef.External = releaseConfigRef

	if obj.Spec.InvocationConfig != nil && obj.Spec.InvocationConfig.ServiceAccountRef != nil {
		if err := obj.Spec.InvocationConfig.ServiceAccountRef.Resolve(ctx, reader, obj); err != nil {
			return err
		}
	}
	return nil
}

func (m *workflowConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type workflowConfigAdapter struct {
	id        *WorkflowConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.DataformRepositoryWorkflowConfig
	actual    *dataformpb.WorkflowConfig
}

var _ directbase.Adapter = &workflowConfigAdapter{}

func (a *workflowConfigAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.workflowConfig == "" {
		return false, nil
	}

	req := &dataformpb.GetWorkflowConfigRequest{Name: a.id.FullyQualifiedName()}
	actual, err := a.gcpClient.GetWorkflowConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataformRepositoryWorkflowConfig %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *workflowConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	mapCtx := &direct.MapContext{}
	resource := DataformRepositoryWorkflowConfigSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("converting DataformRepositoryWorkflowConfig spec to api: %w", mapCtx.Err())
	}

	req := &dataformpb.CreateWorkflowConfigRequest{
		Parent:           a.id.Parent(),
		WorkflowConfig:   resource,
		WorkflowConfigId: a.id.workflowConfig,
	}
	_, err := a.gcpClient.CreateWorkflowConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("DataformRepositoryWorkflowConfig %s creation failed: %w", a.id.FullyQualifiedName(), err)
	}

	status := &krm.DataformRepositoryWorkflowConfigStatus{}
	status.ExternalRef = a.id.AsExternalRef()

	return setStatus(u, status)
}

func (a *workflowConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	updateMask := &fieldmaskpb.FieldMask{}

	if a.desired.Spec.ReleaseConfigRef.External != a.actual.ReleaseConfig {
		report.AddField("release_config", a.actual.ReleaseConfig, a.desired.Spec.ReleaseConfigRef.External)
		updateMask.Paths = append(updateMask.Paths, "release_config")
	}

	if a.desired.Spec.InvocationConfig != nil {
		mapCtx := &direct.MapContext{}
		protoDesired := WorkflowConfigInvocationConfig_ToProto(mapCtx, a.desired.Spec.InvocationConfig)
		if mapCtx.Err() != nil {
			return fmt.Errorf("converting InvocationConfig to api: %w", mapCtx.Err())
		}
		if !reflect.DeepEqual(protoDesired, a.actual.InvocationConfig) {
			report.AddField("invocation_config", a.actual.InvocationConfig, protoDesired)
			updateMask.Paths = append(updateMask.Paths, "invocation_config")
		}
	} else if a.actual.InvocationConfig != nil {
		report.AddField("invocation_config", a.actual.InvocationConfig, nil)
		updateMask.Paths = append(updateMask.Paths, "invocation_config")
	}

	if direct.ValueOf(a.desired.Spec.CronSchedule) != a.actual.CronSchedule {
		report.AddField("cron_schedule", a.actual.CronSchedule, direct.ValueOf(a.desired.Spec.CronSchedule))
		updateMask.Paths = append(updateMask.Paths, "cron_schedule")
	}

	if direct.ValueOf(a.desired.Spec.TimeZone) != a.actual.TimeZone {
		report.AddField("time_zone", a.actual.TimeZone, direct.ValueOf(a.desired.Spec.TimeZone))
		updateMask.Paths = append(updateMask.Paths, "time_zone")
	}

	if direct.ValueOf(a.desired.Spec.Disabled) != a.actual.Disabled {
		report.AddField("disabled", a.actual.Disabled, direct.ValueOf(a.desired.Spec.Disabled))
		updateMask.Paths = append(updateMask.Paths, "disabled")
	}

	if len(updateMask.Paths) == 0 {
		status := &krm.DataformRepositoryWorkflowConfigStatus{}
		status.ExternalRef = a.id.AsExternalRef()
		return setStatus(u, status)
	}

	structuredreporting.ReportDiff(ctx, report)

	mapCtx := &direct.MapContext{}
	resource := DataformRepositoryWorkflowConfigSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("converting DataformRepositoryWorkflowConfig spec to api: %w", mapCtx.Err())
	}

	resource.Name = a.id.FullyQualifiedName()
	req := &dataformpb.UpdateWorkflowConfigRequest{UpdateMask: updateMask, WorkflowConfig: resource}
	_, err := a.gcpClient.UpdateWorkflowConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("DataformRepositoryWorkflowConfig %s update failed: %w", a.id.FullyQualifiedName(), err)
	}

	status := &krm.DataformRepositoryWorkflowConfigStatus{}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *workflowConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataformRepositoryWorkflowConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataformRepositoryWorkflowConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.RepositoryRef.External = a.id.Parent()
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *workflowConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &dataformpb.DeleteWorkflowConfigRequest{Name: a.id.FullyQualifiedName()}
	if err := a.gcpClient.DeleteWorkflowConfig(ctx, req); err != nil {
		return false, fmt.Errorf("deleting DataformRepositoryWorkflowConfig %s: %w", a.id.FullyQualifiedName(), err)
	}

	return true, nil
}
