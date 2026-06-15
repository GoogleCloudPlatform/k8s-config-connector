// Copyright 2024 Google LLC
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
	"sort"
	"strings"

	gcp "cloud.google.com/go/logging/apiv2"
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"google.golang.org/api/option"
	apipb "google.golang.org/genproto/googleapis/api"
	labelpb "google.golang.org/genproto/googleapis/api/label"
	metricpb "google.golang.org/genproto/googleapis/api/metric"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

const ctrlName = "logmetric-controller"

func init() {
	registry.RegisterModel(krm.LoggingLogMetricGVK, NewLogMetricModel)
}

func NewLogMetricModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &logMetricModel{config: config}, nil
}

type logMetricModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &logMetricModel{}

type logMetricAdapter struct {
	id *krm.LoggingLogMetricIdentity

	desired         *pb.LogMetric
	actual          *pb.LogMetric
	logMetricClient *gcp.MetricsClient
}

var _ directbase.Adapter = &logMetricAdapter{}

func (m *logMetricModel) client(ctx context.Context) (*gcp.MetricsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewMetricsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Logging Metrics client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *logMetricModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.LoggingLogMetric{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := LoggingLogMetricSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &logMetricAdapter{
		id:              id.(*krm.LoggingLogMetricIdentity),
		desired:         desired,
		logMetricClient: gcpClient,
	}, nil
}

func (m *logMetricModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//logging.googleapis.com/") {
		return nil, nil
	}

	url = strings.TrimPrefix(url, "//logging.googleapis.com/")

	id := &krm.LoggingLogMetricIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &logMetricAdapter{
		id:              id,
		logMetricClient: gcpClient,
	}, nil
}

func (a *logMetricAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting LoggingLogMetric", "name", a.id)

	req := &pb.GetLogMetricRequest{MetricName: a.id.String()}
	actual, err := a.logMetricClient.GetLogMetric(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting LoggingLogMetric %q: %w", a.id, err)
	}

	a.actual = actual

	return true, nil
}

// Delete implements the Adapter interface.
func (a *logMetricAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Already deleted
	if a.id.Metric == "" {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting LoggingLogMetric", "name", a.id)

	req := &pb.DeleteLogMetricRequest{MetricName: a.id.String()}
	err := a.logMetricClient.DeleteLogMetric(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting log metric %s: %w", a.id, err)
	}

	return true, nil
}

func (a *logMetricAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	projectID := a.id.Project
	if projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.id.Metric == "" {
		return fmt.Errorf("resourceID is empty")
	}
	if a.desired.Filter == "" {
		return fmt.Errorf("filter is empty")
	}
	if a.desired.BucketName != "" {
		bucket, err := LogBucketRef_Parse(ctx, a.desired.BucketName)
		if err != nil {
			return err
		}

		// validate that the bucket is in the same project
		if bucket.ProjectID() != a.id.Project {
			return fmt.Errorf("LoggingLogBucket %q is not in the same project %q", bucket.FQN(), a.id.Project)
		}
	}

	logMetric := a.desired
	logMetric.Name = a.id.Metric

	req := &pb.CreateLogMetricRequest{
		Parent: a.id.ParentString(),
		Metric: logMetric,
	}
	log.V(2).Info("creating logMetric", "request", req, "name", logMetric.Name)
	created, err := a.logMetricClient.CreateLogMetric(ctx, req)
	if err != nil {
		return fmt.Errorf("logMetric %s creation failed: %w", logMetric.Name, err)
	}

	log.V(2).Info("created logMetric", "logMetric", created)

	resourceID := created.Name
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *logMetricAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating logMetric", "name", a.id.String())

	u := updateOp.GetUnstructured()
	diffResults, _, err := compareLogMetric(ctx, a.actual, a.desired, u)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffResults.HasDiff() {
		diffResults.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffResults)

		a.desired.Name = a.id.Metric

		req := &pb.UpdateLogMetricRequest{
			MetricName: a.fullyQualifiedName(),
			Metric:     a.desired,
		}

		updated, err := a.logMetricClient.UpdateLogMetric(ctx, req)
		if err != nil {
			return fmt.Errorf("logMetric update failed: %w", err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *logMetricAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.LogMetric) error {
	mapCtx := &direct.MapContext{}
	status := LoggingLogMetricStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// latest.Description may be set but may not map directly to status.MetricDescriptor.Description if needed,
	// let's match the old behavior just in case:
	if latest.Description != "" {
		if status.MetricDescriptor != nil {
			status.MetricDescriptor.Description = &latest.Description
		}
	}

	return op.UpdateStatus(ctx, status, nil)
}

func (a *logMetricAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("logMetric %q not found", a.fullyQualifiedName())
	}

	u := &unstructured.Unstructured{}

	obj := &krm.LoggingLogMetric{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(LoggingLogMetricSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	if a.desired != nil {
		if a.desired.MetricDescriptor != nil {
			if obj.Spec.MetricDescriptor == nil {
				obj.Spec.MetricDescriptor = &krm.LogmetricMetricDescriptor{}
			}
			if a.desired.MetricDescriptor.LaunchStage != apipb.LaunchStage_LAUNCH_STAGE_UNSPECIFIED {
				obj.Spec.MetricDescriptor.LaunchStage = direct.LazyPtr(a.desired.MetricDescriptor.LaunchStage.String())
			}
			if a.desired.MetricDescriptor.Metadata != nil {
				if obj.Spec.MetricDescriptor.Metadata == nil {
					obj.Spec.MetricDescriptor.Metadata = &krm.LogmetricMetadata{}
				}
				if a.desired.MetricDescriptor.Metadata.IngestDelay != nil {
					obj.Spec.MetricDescriptor.Metadata.IngestDelay = direct.LazyPtr(a.desired.MetricDescriptor.Metadata.IngestDelay.AsDuration().String())
				}
				if a.desired.MetricDescriptor.Metadata.SamplePeriod != nil {
					obj.Spec.MetricDescriptor.Metadata.SamplePeriod = direct.LazyPtr(a.desired.MetricDescriptor.Metadata.SamplePeriod.AsDuration().String())
				}
			}
		}
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.LoggingLogMetricGVK)
	return u, nil
}

func (a *logMetricAdapter) fullyQualifiedName() string {
	return a.id.String()
}

func compareLogMetric(ctx context.Context, actual, desired *pb.LogMetric, u *unstructured.Unstructured) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, LoggingLogMetricSpec_FromProto, LoggingLogMetricSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.Clone(desired).(*pb.LogMetric)

	populateDefaults := func(obj *pb.LogMetric) {
		if obj.MetricDescriptor == nil {
			obj.MetricDescriptor = &metricpb.MetricDescriptor{}
		}
		if obj.MetricDescriptor.MetricKind == metricpb.MetricDescriptor_METRIC_KIND_UNSPECIFIED {
			obj.MetricDescriptor.MetricKind = metricpb.MetricDescriptor_DELTA
		}
		if obj.MetricDescriptor.ValueType == metricpb.MetricDescriptor_VALUE_TYPE_UNSPECIFIED {
			obj.MetricDescriptor.ValueType = metricpb.MetricDescriptor_INT64
		}
		if obj.MetricDescriptor.Unit == "" {
			obj.MetricDescriptor.Unit = "1"
		}
		sort.Slice(obj.MetricDescriptor.Labels, func(i, j int) bool {
			return obj.MetricDescriptor.Labels[i].Key < obj.MetricDescriptor.Labels[j].Key
		})
		for _, label := range obj.MetricDescriptor.Labels {
			if label.ValueType == labelpb.LabelDescriptor_ValueType(0) {
				label.ValueType = labelpb.LabelDescriptor_ValueType(1)
			}
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	gcpUpdateTime := actual.UpdateTime
	updateTimeMatches := false
	if gcpUpdateTime != nil && u != nil {
		obj := &krm.LoggingLogMetric{}
		if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err == nil {
			ready := false
			for _, condition := range obj.Status.Conditions {
				if condition.Type == "Ready" {
					if condition.Status == "True" {
						ready = true
					}
				}
			}
			mapCtx := &direct.MapContext{}
			gcpUpdateTimeStr := direct.StringTimestamp_FromProto(mapCtx, gcpUpdateTime)
			if mapCtx.Err() == nil && gcpUpdateTimeStr != nil && ready && obj.Status.UpdateTime != nil && *gcpUpdateTimeStr == *obj.Status.UpdateTime && u.GetGeneration() == getObservedGeneration(u) {
				updateTimeMatches = true
			}
		}
	}

	if updateTimeMatches {
		if maskedActual.MetricDescriptor != nil {
			// If the updateTime matches and the object is ready, the GCP side matches KRM status.
			// We can safely assume the unreadable fields (LaunchStage and Metadata) match, so we align them to avoid false diffs.
			maskedActual.MetricDescriptor.LaunchStage = clonedDesired.MetricDescriptor.LaunchStage
			maskedActual.MetricDescriptor.Metadata = clonedDesired.MetricDescriptor.Metadata
		}
	} else {
		// If updateTime doesn't match, we cannot guarantee that the mutable-but-unreadable fields are up to date.
		// Therefore, we force a diff in the metricDescriptor to trigger an update request.
		if maskedActual.MetricDescriptor != nil {
			if clonedDesired.MetricDescriptor.LaunchStage == apipb.LaunchStage_LAUNCH_STAGE_UNSPECIFIED {
				maskedActual.MetricDescriptor.LaunchStage = apipb.LaunchStage_EARLY_ACCESS
			} else {
				maskedActual.MetricDescriptor.LaunchStage = apipb.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
			}
		}
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
