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
)

func init() {
	registry.RegisterModel(krm.DataformRepositoryReleaseConfigGVK, NewReleaseConfigModel)
}

func NewReleaseConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &releaseConfigModel{config: *config}, nil
}

var _ directbase.Model = &releaseConfigModel{}

type releaseConfigModel struct {
	config config.ControllerConfig
}

func (m *releaseConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DataformRepositoryReleaseConfig{}
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

	var id *ReleaseConfigIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildReleaseConfigID(repositoryID.project, repositoryID.location, repositoryID.dataform, resourceID)
	} else {
		id, err = asReleaseConfigID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.project != repositoryID.project {
			return nil, fmt.Errorf("DataformRepositoryReleaseConfig %s/%s has repositoryRef changed (project), expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.project, repositoryID.project)
		}
		if id.location != repositoryID.location {
			return nil, fmt.Errorf("DataformRepositoryReleaseConfig %s/%s has repositoryRef changed (location), expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.location, repositoryID.location)
		}
		if id.repository != repositoryID.dataform {
			return nil, fmt.Errorf("DataformRepositoryReleaseConfig %s/%s has repositoryRef changed (repository), expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.repository, repositoryID.dataform)
		}
		if id.releaseConfig != resourceID {
			return nil, fmt.Errorf("DataformRepositoryReleaseConfig %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.releaseConfig, resourceID)
		}
	}

	gcpModel := &model{config: m.config}
	gcpClient, err := gcpModel.client(ctx)
	if err != nil {
		return nil, err
	}

	return &releaseConfigAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *releaseConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type releaseConfigAdapter struct {
	id        *ReleaseConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.DataformRepositoryReleaseConfig
	actual    *dataformpb.ReleaseConfig
}

var _ directbase.Adapter = &releaseConfigAdapter{}

func (a *releaseConfigAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.releaseConfig == "" {
		return false, nil
	}

	req := &dataformpb.GetReleaseConfigRequest{Name: a.id.FullyQualifiedName()}
	actual, err := a.gcpClient.GetReleaseConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DataformRepositoryReleaseConfig %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *releaseConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	mapCtx := &direct.MapContext{}
	resource := DataformRepositoryReleaseConfigSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("converting DataformRepositoryReleaseConfig spec to api: %w", mapCtx.Err())
	}

	req := &dataformpb.CreateReleaseConfigRequest{
		Parent:          a.id.Parent(),
		ReleaseConfig:   resource,
		ReleaseConfigId: a.id.releaseConfig,
	}
	_, err := a.gcpClient.CreateReleaseConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("DataformRepositoryReleaseConfig %s creation failed: %w", a.id.FullyQualifiedName(), err)
	}

	status := &krm.DataformRepositoryReleaseConfigStatus{}
	status.ExternalRef = a.id.AsExternalRef()

	return setStatus(u, status)
}

func (a *releaseConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	updateMask := &fieldmaskpb.FieldMask{}

	if a.desired.Spec.GitCommitish != a.actual.GitCommitish {
		report.AddField("git_commitish", a.actual.GitCommitish, a.desired.Spec.GitCommitish)
		updateMask.Paths = append(updateMask.Paths, "git_commitish")
	}

	if a.desired.Spec.CodeCompilationConfig != nil {
		mapCtx := &direct.MapContext{}
		protoDesired := ReleaseConfigCodeCompilationConfig_ToProto(mapCtx, a.desired.Spec.CodeCompilationConfig)
		if mapCtx.Err() != nil {
			return fmt.Errorf("converting CodeCompilationConfig to api: %w", mapCtx.Err())
		}
		if !reflect.DeepEqual(protoDesired, a.actual.CodeCompilationConfig) {
			report.AddField("code_compilation_config", a.actual.CodeCompilationConfig, protoDesired)
			updateMask.Paths = append(updateMask.Paths, "code_compilation_config")
		}
	} else if a.actual.CodeCompilationConfig != nil {
		report.AddField("code_compilation_config", a.actual.CodeCompilationConfig, nil)
		updateMask.Paths = append(updateMask.Paths, "code_compilation_config")
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
		status := &krm.DataformRepositoryReleaseConfigStatus{}
		status.ExternalRef = a.id.AsExternalRef()
		return setStatus(u, status)
	}

	structuredreporting.ReportDiff(ctx, report)

	mapCtx := &direct.MapContext{}
	resource := DataformRepositoryReleaseConfigSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return fmt.Errorf("converting DataformRepositoryReleaseConfig spec to api: %w", mapCtx.Err())
	}

	resource.Name = a.id.FullyQualifiedName()
	req := &dataformpb.UpdateReleaseConfigRequest{UpdateMask: updateMask, ReleaseConfig: resource}
	_, err := a.gcpClient.UpdateReleaseConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("DataformRepositoryReleaseConfig %s update failed: %w", a.id.FullyQualifiedName(), err)
	}

	status := &krm.DataformRepositoryReleaseConfigStatus{}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *releaseConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.DataformRepositoryReleaseConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DataformRepositoryReleaseConfigSpec_FromProto(mapCtx, a.actual))
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

func (a *releaseConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	req := &dataformpb.DeleteReleaseConfigRequest{Name: a.id.FullyQualifiedName()}
	if err := a.gcpClient.DeleteReleaseConfig(ctx, req); err != nil {
		return false, fmt.Errorf("deleting DataformRepositoryReleaseConfig %s: %w", a.id.FullyQualifiedName(), err)
	}

	return true, nil
}
