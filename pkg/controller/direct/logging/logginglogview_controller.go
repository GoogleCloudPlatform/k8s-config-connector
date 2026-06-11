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
	registry.RegisterModel(krm.LoggingLogViewGVK, NewLoggingLogViewModel)
}

func NewLoggingLogViewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLoggingLogView{config: *config}, nil
}

var _ directbase.Model = &modelLoggingLogView{}

type modelLoggingLogView struct {
	config config.ControllerConfig
}

func (m *modelLoggingLogView) client(ctx context.Context) (*gcp.ConfigClient, error) {
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

func (m *modelLoggingLogView) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LoggingLogView{}
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
	desiredPb := LoggingLogViewSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredPb.Name = id.String()

	return &LoggingLogViewAdapter{
		id:        id.(*krm.LoggingLogViewIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
	}, nil
}

func (m *modelLoggingLogView) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type LoggingLogViewAdapter struct {
	id        *krm.LoggingLogViewIdentity
	gcpClient *gcp.ConfigClient
	desired   *loggingpb.LogView
	actual    *loggingpb.LogView
}

var _ directbase.Adapter = &LoggingLogViewAdapter{}

func (a *LoggingLogViewAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLogView", "name", a.id)

	req := &loggingpb.GetViewRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetView(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLogView %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *LoggingLogViewAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating LoggingLogView", "id", fqn)

	parent := a.id.ParentString()
	resourceID := a.id.View

	req := &loggingpb.CreateViewRequest{
		Parent: parent,
		View:   a.desired,
		ViewId: resourceID,
	}
	created, err := a.gcpClient.CreateView(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LogView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created LogView", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *LoggingLogViewAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LoggingLogView", "name", a.id.String())

	diffs, updateMask, err := compareView(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &loggingpb.UpdateViewRequest{
			Name:       a.id.String(),
			View:       a.desired,
			UpdateMask: updateMask,
		}

		updated, err := a.gcpClient.UpdateView(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LoggingLogView %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *LoggingLogViewAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *loggingpb.LogView) error {
	mapCtx := &direct.MapContext{}
	status := LoggingLogViewStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLogViewAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLogView{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLogViewSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingLogViewGVK)
	return u, nil
}

func (a *LoggingLogViewAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LogView", "name", a.id)

	req := &loggingpb.DeleteViewRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteView(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LogView, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LogView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LogView", "name", a.id)
	return true, nil
}

func compareView(ctx context.Context, actual, desired *loggingpb.LogView) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *loggingpb.LogView
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := LoggingLogViewSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = LoggingLogViewSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
