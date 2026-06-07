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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.LoggingLogSinkGVK, NewLoggingLogSinkModel)
}

func NewLoggingLogSinkModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLoggingLogSink{config: *config}, nil
}

var _ directbase.Model = &modelLoggingLogSink{}

type modelLoggingLogSink struct {
	config config.ControllerConfig
}

func (m *modelLoggingLogSink) client(ctx context.Context) (*gcp.ConfigClient, error) {
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

func (m *modelLoggingLogSink) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LoggingLogSink{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	if err := m.resolveReferences(ctx, reader, obj); err != nil {
		return nil, fmt.Errorf("resolving destination references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	sinkID := id.(*krm.LoggingLogSinkIdentity).ID()

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredPb := LoggingLogSinkSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredPb.Name = sinkID

	return &LoggingLogSinkAdapter{
		id:        id.(*krm.LoggingLogSinkIdentity),
		gcpClient: gcpClient,
		desired:   desiredPb,
		spec:      &obj.Spec,
	}, nil
}

func (m *modelLoggingLogSink) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (m *modelLoggingLogSink) resolveReferences(ctx context.Context, reader client.Reader, obj *krm.LoggingLogSink) error {
	dest := &obj.Spec.Destination
	if dest.PubSubTopicRef != nil {
		ext, err := dest.PubSubTopicRef.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return err
		}
		dest.PubSubTopicRef.External = ext
	}
	if dest.BigQueryDatasetRef != nil {
		if err := dest.BigQueryDatasetRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if dest.LoggingLogBucketRef != nil {
		if err := dest.LoggingLogBucketRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	if dest.StorageBucketRef != nil {
		if err := dest.StorageBucketRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
			return err
		}
	}
	return nil
}

type LoggingLogSinkAdapter struct {
	id        *krm.LoggingLogSinkIdentity
	gcpClient *gcp.ConfigClient
	desired   *loggingpb.LogSink
	actual    *loggingpb.LogSink
	spec      *krm.LoggingLogSinkSpec
}

var _ directbase.Adapter = &LoggingLogSinkAdapter{}

func (a *LoggingLogSinkAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLogSink", "name", a.id)

	req := &loggingpb.GetSinkRequest{SinkName: a.id.String()}
	actual, err := a.gcpClient.GetSink(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLogSink %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *LoggingLogSinkAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating LoggingLogSink", "id", fqn)

	parent := a.parentString()

	req := &loggingpb.CreateSinkRequest{
		Parent: parent,
		Sink:   a.desired,
	}
	if a.spec.UniqueWriterIdentity != nil {
		req.UniqueWriterIdentity = *a.spec.UniqueWriterIdentity
	}

	created, err := a.gcpClient.CreateSink(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LogSink %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created LogSink", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *LoggingLogSinkAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LoggingLogSink", "name", a.id.String())

	diffs, updateMask, err := compareSink(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &loggingpb.UpdateSinkRequest{
			SinkName:   a.id.String(),
			Sink:       a.desired,
			UpdateMask: updateMask,
		}
		if a.spec.UniqueWriterIdentity != nil {
			req.UniqueWriterIdentity = *a.spec.UniqueWriterIdentity
		}

		updated, err := a.gcpClient.UpdateSink(ctx, req)
		if err != nil {
			return fmt.Errorf("updating LoggingLogSink %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *LoggingLogSinkAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *loggingpb.LogSink) error {
	mapCtx := &direct.MapContext{}
	status := LoggingLogSinkStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLogSinkAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLogSink{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLogSinkSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingLogSinkGVK)
	return u, nil
}

func (a *LoggingLogSinkAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LogSink", "name", a.id)

	req := &loggingpb.DeleteSinkRequest{SinkName: a.id.String()}
	err := a.gcpClient.DeleteSink(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LogSink, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LogSink %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LogSink", "name", a.id)
	return true, nil
}

func (a *LoggingLogSinkAdapter) parentString() string {
	if a.id.Project != "" {
		return fmt.Sprintf("projects/%s", a.id.Project)
	}
	if a.id.Folder != "" {
		return fmt.Sprintf("folders/%s", a.id.Folder)
	}
	if a.id.Organization != "" {
		return fmt.Sprintf("organizations/%s", a.id.Organization)
	}
	if a.id.BillingAccount != "" {
		return fmt.Sprintf("billingAccounts/%s", a.id.BillingAccount)
	}
	return ""
}

func compareSink(ctx context.Context, actual, desired *loggingpb.LogSink) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *loggingpb.LogSink
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := LoggingLogSinkSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = LoggingLogSinkSpec_ToProto(mapCtx, spec)
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
