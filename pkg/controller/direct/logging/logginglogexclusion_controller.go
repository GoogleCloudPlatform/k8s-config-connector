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

package logging

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/logging/apiv2"
	loggingpb "cloud.google.com/go/logging/apiv2/loggingpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.LoggingLogExclusionGVK, NewLoggingLogExclusionModel)
}

func NewLoggingLogExclusionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLoggingLogExclusion{config: *config}, nil
}

var _ directbase.Model = &modelLoggingLogExclusion{}

type modelLoggingLogExclusion struct {
	config config.ControllerConfig
}

func (m *modelLoggingLogExclusion) client(ctx context.Context) (*gcp.ConfigClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Logging Config client: %w", err)
	}
	return gcpClient, err
}

func (m *modelLoggingLogExclusion) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LoggingLogExclusion{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := LoggingLogExclusionSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredPb.Name = id.String()

	return &LoggingLogExclusionAdapter{
		id:        id.(*krm.LoggingLogExclusionIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelLoggingLogExclusion) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LoggingLogExclusionAdapter struct {
	id        *krm.LoggingLogExclusionIdentity
	gcpClient *gcp.ConfigClient
	desired   *loggingpb.LogExclusion
	actual    *loggingpb.LogExclusion
}

var _ directbase.Adapter = &LoggingLogExclusionAdapter{}

func (a *LoggingLogExclusionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLogExclusion", "name", a.id)

	req := &loggingpb.GetExclusionRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetExclusion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLogExclusion %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *LoggingLogExclusionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating LoggingLogExclusion", "id", fqn)

	parent := a.id.ParentString()

	// Since we got the name from identity, let's make sure the ID field is set on the pb.LogExclusion name
	// Although name typically matches "projects/[PROJECT_ID]/exclusions/[EXCLUSION_ID]",
	// let's follow the standard pattern.
	a.desired.Name = a.id.ID()

	req := &loggingpb.CreateExclusionRequest{
		Parent:    parent,
		Exclusion: a.desired,
	}
	created, err := a.gcpClient.CreateExclusion(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LogExclusion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created LogExclusion", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *LoggingLogExclusionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LoggingLogExclusion", "name", a.id.String())

	diffs, updateMask, err := compareExclusion(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &loggingpb.UpdateExclusionRequest{
			Name:       a.id.String(),
			Exclusion:  a.desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateExclusion(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LoggingLogExclusion %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *LoggingLogExclusionAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *loggingpb.LogExclusion) error {
	mapCtx := &direct.MapContext{}
	status := LoggingLogExclusionStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLogExclusionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLogExclusion{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLogExclusionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingLogExclusionGVK)
	return u, nil
}

func (a *LoggingLogExclusionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LoggingLogExclusion", "name", a.id)

	req := &loggingpb.DeleteExclusionRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteExclusion(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LoggingLogExclusion, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LoggingLogExclusion %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LoggingLogExclusion", "name", a.id)
	return true, nil
}

func compareExclusion(ctx context.Context, actual, desired *loggingpb.LogExclusion) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, LoggingLogExclusionSpec_FromProto, LoggingLogExclusionSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
