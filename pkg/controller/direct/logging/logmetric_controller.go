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

	api "google.golang.org/api/logging/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/diffs"
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

	desired         *api.LogMetric
	actual          *api.LogMetric
	logMetricClient *api.ProjectsMetricsService
}

var _ directbase.Adapter = &logMetricAdapter{}

// AdapterForObject implements the Model interface.
func (m *logMetricModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	projectMetricsService, err := gcpClient.newProjectMetricsService(ctx)
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
	desired := LoggingLogMetricSpec_ToAPI(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &logMetricAdapter{
		id:              id.(*krm.LoggingLogMetricIdentity),
		desired:         desired,
		logMetricClient: projectMetricsService,
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

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}

	projectMetricsService, err := gcpClient.newProjectMetricsService(ctx)
	if err != nil {
		return nil, err
	}

	return &logMetricAdapter{
		id:              id,
		logMetricClient: projectMetricsService,
	}, nil
}

func (a *logMetricAdapter) Find(ctx context.Context) (bool, error) {
	logMetric, err := a.logMetricClient.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting logMetric %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = logMetric

	return true, nil
}

// Delete implements the Adapter interface.
func (a *logMetricAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// Already deleted
	if a.id.Metric == "" {
		return false, nil
	}

	_, err := a.logMetricClient.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting log metric %s: %w", a.fullyQualifiedName(), err)
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

	createRequest := a.logMetricClient.Create("projects/"+projectID, logMetric)
	log.V(2).Info("creating logMetric", "request", &createRequest, "name", logMetric.Name)
	created, err := createRequest.Context(ctx).Do()
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

func logMetricStatusToKRM(id *krm.LoggingLogMetricIdentity, in *api.LogMetric, out *krm.LoggingLogMetricStatus) error {
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)

	out.MetricDescriptor = convertAPItoKRM_MetricDescriptorStatus(in.MetricDescriptor)

	return nil
}

func (a *logMetricAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating logMetric", "name", a.id.String())

	u := updateOp.GetUnstructured()
	diffResults, err := compareLogMetric(ctx, a.actual, a.desired, u)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffResults.HasDiff() {
		diffResults.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffResults)

		update := common.DeepCopy(a.desired)
		update.Name = a.id.Metric

		updated, err := a.logMetricClient.Update(a.fullyQualifiedName(), update).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("logMetric update failed: %w", err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *logMetricAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.LogMetric) error {
	status := &krm.LoggingLogMetricStatus{}
	if err := logMetricStatusToKRM(a.id, latest, status); err != nil {
		return err
	}

	// actualUpdate may not contain the description for the metric descriptor.
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

	un, err := convertAPItoKRM_LoggingLogMetric(a.id.Project, a.actual)
	if err != nil {
		return nil, fmt.Errorf("error converting logMetric to unstructured %w", err)
	}

	// TODO(acpana): revisit if we want to include mutable but unreadable fields in our export
	if a.desired != nil {
		if a.desired.MetricDescriptor != nil && a.desired.MetricDescriptor.LaunchStage != "" {
			if err := unstructured.SetNestedField(un.Object,
				a.desired.MetricDescriptor.LaunchStage,
				"spec", "metricDescriptor", "launchStage",
			); err != nil {
				return nil, fmt.Errorf("could not set metricDescriptor.launchStage mutable but unreadable field %w", err)
			}
		}
		if a.desired.MetricDescriptor != nil && a.desired.MetricDescriptor.Metadata != nil {
			if a.desired.MetricDescriptor.Metadata.IngestDelay != "" {
				if err := unstructured.SetNestedField(un.Object,
					a.desired.MetricDescriptor.Metadata.IngestDelay,
					"spec", "metricDescriptor", "metadata", "ingestDelay",
				); err != nil {
					return nil, fmt.Errorf("could not set metricDescriptor.metadata.ingestDelay mutable but unreadable field %w", err)
				}
			}
			if a.desired.MetricDescriptor.Metadata.SamplePeriod != "" {
				if err := unstructured.SetNestedField(un.Object,
					a.desired.MetricDescriptor.Metadata.SamplePeriod,
					"spec", "metricDescriptor", "metadata", "samplePeriod",
				); err != nil {
					return nil, fmt.Errorf("could not set metricDescriptor.metadata.samplePeriod mutable but unreadable field %w", err)
				}
			}
		}
	}

	return un, nil
}

func (a *logMetricAdapter) fullyQualifiedName() string {
	return a.id.String()
}

func compareLogMetric(ctx context.Context, actual, desired *api.LogMetric, u *unstructured.Unstructured) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, LoggingLogMetricSpec_FromAPI, LoggingLogMetricSpec_ToAPI)
	if err != nil {
		return nil, err
	}

	clonedDesired := common.DeepCopy(desired)

	populateDefaults := func(obj *api.LogMetric) {
		if obj.MetricDescriptor != nil {
			if obj.MetricDescriptor.MetricKind == "" {
				obj.MetricDescriptor.MetricKind = "DELTA"
			}
			if obj.MetricDescriptor.ValueType == "" {
				obj.MetricDescriptor.ValueType = "INT64"
			}
			if obj.MetricDescriptor.Unit == "" {
				obj.MetricDescriptor.Unit = "1"
			}
			sort.Slice(obj.MetricDescriptor.Labels, func(i, j int) bool {
				return obj.MetricDescriptor.Labels[i].Key < obj.MetricDescriptor.Labels[j].Key
			})
			for _, label := range obj.MetricDescriptor.Labels {
				if label.ValueType == "" {
					label.ValueType = "STRING"
				}
			}
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	// If the user did not specify MetricDescriptor, we mask it out from actual to avoid comparing server-generated defaults.
	if desired.MetricDescriptor == nil {
		maskedActual.MetricDescriptor = nil
	}

	gcpUpdateTime := actual.UpdateTime
	updateTimeMatches := false
	if gcpUpdateTime != "" && u != nil {
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
			if ready && obj.Status.UpdateTime != nil && gcpUpdateTime == *obj.Status.UpdateTime && u.GetGeneration() == getObservedGeneration(u) {
				updateTimeMatches = true
			}
		}
	}

	if updateTimeMatches {
		if maskedActual.MetricDescriptor != nil && clonedDesired.MetricDescriptor != nil {
			// If the updateTime matches and the object is ready, the GCP side matches KRM status.
			// We can safely assume the unreadable fields (LaunchStage and Metadata) match, so we align them to avoid false diffs.
			maskedActual.MetricDescriptor.LaunchStage = clonedDesired.MetricDescriptor.LaunchStage
			maskedActual.MetricDescriptor.Metadata = clonedDesired.MetricDescriptor.Metadata
		}
	} else {
		if maskedActual.MetricDescriptor != nil && clonedDesired.MetricDescriptor != nil {
			// If updateTime doesn't match (meaning a new update/reconciliation cycle), we cannot
			// guarantee that mutable-but-unreadable fields are up to date.
			// Therefore, we force a diff in the metricDescriptor to trigger an update request.
			maskedActual.MetricDescriptor.LaunchStage = "FORCE_DIFF_UNREADABLE_FIELD"
		}
	}

	diffResults, _, err := diffs.GoogleAPI.Diff(ctx, maskedActual, clonedDesired)
	if err != nil {
		return nil, err
	}
	return diffResults, nil
}

func LoggingLogMetricSpec_FromAPI(mapCtx *direct.MapContext, in *api.LogMetric) *krm.LoggingLogMetricSpec {
	if in == nil {
		return nil
	}
	out := &krm.LoggingLogMetricSpec{}
	out.Description = direct.LazyPtr(in.Description)
	out.Disabled = direct.LazyPtr(in.Disabled)
	out.Filter = in.Filter
	out.LabelExtractors = in.LabelExtractors
	out.MetricDescriptor = convertAPItoKRM_MetricDescriptor(in.MetricDescriptor)
	out.BucketOptions = convertAPItoKRM_BucketOptions(in.BucketOptions)
	out.ValueExtractor = direct.LazyPtr(in.ValueExtractor)
	if in.BucketName != "" {
		out.LoggingLogBucketRef = &krm.LoggingLogBucketRef{
			External: in.BucketName,
		}
	}
	return out
}

func LoggingLogMetricSpec_ToAPI(mapCtx *direct.MapContext, in *krm.LoggingLogMetricSpec) *api.LogMetric {
	return convertKCCtoAPI(in)
}
