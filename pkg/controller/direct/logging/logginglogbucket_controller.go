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
	registry.RegisterModel(krm.LoggingLogBucketGVK, NewLoggingLogBucketModel)
}

func NewLoggingLogBucketModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelLoggingLogBucket{config: *config}, nil
}

var _ directbase.Model = &modelLoggingLogBucket{}

type modelLoggingLogBucket struct {
	config config.ControllerConfig
}

func (m *modelLoggingLogBucket) client(ctx context.Context) (*gcp.ConfigClient, error) {
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

func (m *modelLoggingLogBucket) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.LoggingLogBucket{}
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

	return &LoggingLogBucketAdapter{
		id:        id.(*krm.LogBucketIdentity),
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelLoggingLogBucket) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.LogBucketIdentity{}
	if err := id.FromExternal(url); err != nil {
		// Not recognized
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &LoggingLogBucketAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type LoggingLogBucketAdapter struct {
	id        *krm.LogBucketIdentity
	gcpClient *gcp.ConfigClient
	desired   *krm.LoggingLogBucket
	actual    *loggingpb.LogBucket
}

var _ directbase.Adapter = &LoggingLogBucketAdapter{}

func (a *LoggingLogBucketAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLogBucket", "name", a.id)

	req := &loggingpb.GetBucketRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetBucket(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLogBucket %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *LoggingLogBucketAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("creating LoggingLogBucket", "id", fqn)

	mapCtx := &direct.MapContext{}
	desired := LoggingLogBucketSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.ParentString()
	resourceID := a.id.Bucket

	req := &loggingpb.CreateBucketRequest{
		Parent:   parent,
		Bucket:   desired,
		BucketId: resourceID,
	}
	created, err := a.gcpClient.CreateBucket(ctx, req)
	if err != nil {
		return fmt.Errorf("creating LogBucket %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created LogBucket", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *LoggingLogBucketAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating LoggingLogBucket", "name", a.id.String())

	mapCtx := &direct.MapContext{}
	desired := LoggingLogBucketSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if a.desired.Spec.Description == nil {
		desired.Description = a.actual.GetDescription()
	}
	if a.desired.Spec.EnableAnalytics == nil {
		desired.AnalyticsEnabled = a.actual.GetAnalyticsEnabled()
	}
	if a.desired.Spec.Locked == nil {
		desired.Locked = a.actual.GetLocked()
	}
	if a.desired.Spec.RetentionDays == nil || *a.desired.Spec.RetentionDays == 0 {
		desired.RetentionDays = a.actual.GetRetentionDays()
	}

	diffs, updateMask, err := compareBucket(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		// Cloud Logging rejects UpdateBucket requests that update analytics_enabled
		// at the same time as other bucket fields ("Buckets cannot be upgraded at the
		// same time as making updates to other bucket fields"). Split into two calls
		// when analytics_enabled is updated alongside other fields.
		hasAnalyticsUpdate := false
		var otherPaths []string
		for _, p := range updateMask.GetPaths() {
			if p == "analytics_enabled" {
				hasAnalyticsUpdate = true
			} else {
				otherPaths = append(otherPaths, p)
			}
		}

		if hasAnalyticsUpdate && len(otherPaths) > 0 {
			analyticsReq := &loggingpb.UpdateBucketRequest{
				Name:       a.id.String(),
				Bucket:     desired,
				UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"analytics_enabled"}},
			}
			updated, err := a.gcpClient.UpdateBucket(ctx, analyticsReq)
			if err != nil {
				return fmt.Errorf("updating LoggingLogBucket %s (analytics_enabled): %w", a.id.String(), err)
			}
			latest = updated

			otherReq := &loggingpb.UpdateBucketRequest{
				Name:       a.id.String(),
				Bucket:     desired,
				UpdateMask: &fieldmaskpb.FieldMask{Paths: otherPaths},
			}
			updated, err = a.gcpClient.UpdateBucket(ctx, otherReq)
			if err != nil {
				return fmt.Errorf("updating LoggingLogBucket %s: %w", a.id.String(), err)
			}
			latest = updated
		} else {
			req := &loggingpb.UpdateBucketRequest{
				Name:       a.id.String(),
				Bucket:     desired,
				UpdateMask: updateMask,
			}

			updated, err := a.gcpClient.UpdateBucket(ctx, req)
			if err != nil {
				return fmt.Errorf("updating LoggingLogBucket %s: %w", a.id.String(), err)
			}
			latest = updated
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *LoggingLogBucketAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *loggingpb.LogBucket) error {
	mapCtx := &direct.MapContext{}
	status := LoggingLogBucketStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *LoggingLogBucketAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLogBucket{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLogBucketSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.Location = a.id.Location

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingLogBucketGVK)
	return u, nil
}

func (a *LoggingLogBucketAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LogBucket", "name", a.id)

	if a.id.Bucket == "_Default" || a.id.Bucket == "_Required" {
		log.V(2).Info("skipping delete for default/required LogBucket", "name", a.id.String())
		return true, nil
	}

	req := &loggingpb.DeleteBucketRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteBucket(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent LogBucket, assuming it was already deleted", "name", a.id.String())
			return true, nil
		}
		return false, fmt.Errorf("deleting LogBucket %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted LogBucket", "name", a.id)
	return true, nil
}

func compareBucket(ctx context.Context, actual, desired *loggingpb.LogBucket) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *loggingpb.LogBucket
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := LoggingLogBucketSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = LoggingLogBucketSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
