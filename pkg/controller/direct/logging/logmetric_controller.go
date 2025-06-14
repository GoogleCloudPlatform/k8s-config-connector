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
	"reflect"
	"strings"

	api "google.golang.org/api/logging/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/logging/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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
	resourceID string
	projectID  string

	desired         *krm.LoggingLogMetric
	actual          *api.LogMetric
	logMetricClient *api.ProjectsMetricsService
}

var _ directbase.Adapter = &logMetricAdapter{}

// AdapterForObject implements the Model interface.
func (m *logMetricModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
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

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if err := LogBucketRef_ConvertToExternal(ctx, reader, obj, &obj.Spec.LoggingLogBucketRef); err != nil {
		return nil, err
	}

	return &logMetricAdapter{
		resourceID:      resourceID,
		projectID:       projectID,
		desired:         obj,
		logMetricClient: projectMetricsService,
	}, nil
}

func (m *logMetricModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format: //logging.googleapis.com/projects/<project>/metrics/<id>
	if !strings.HasPrefix(url, "//logging.googleapis.com/") {
		return nil, nil
	}

	tokens := strings.Split(strings.TrimPrefix(url, "//logging.googleapis.com/"), "/")
	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "metrics" {
		gcpClient, err := newGCPClient(ctx, m.config)
		if err != nil {
			return nil, err
		}

		projectMetricsService, err := gcpClient.newProjectMetricsService(ctx)
		if err != nil {
			return nil, err
		}

		return &logMetricAdapter{
			projectID:       tokens[1],
			resourceID:      tokens[3],
			logMetricClient: projectMetricsService,
		}, nil
	}

	return nil, nil
}

func (a *logMetricAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

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
	if a.resourceID == "" {
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

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating object", "u", u)

	projectID := a.projectID
	if projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}
	filter := a.desired.Spec.Filter
	if filter == "" {
		return fmt.Errorf("filter is empty")
	}
	if a.desired.Spec.LoggingLogBucketRef != nil {
		bucket, err := LogBucketRef_Parse(ctx, a.desired.Spec.LoggingLogBucketRef.External)
		if err != nil {
			return err
		}

		// validate that the bucket is in the same project
		if bucket.ProjectID() != a.projectID {
			return fmt.Errorf("LoggingLogBucket %q is not in the same project %q", bucket.FQN(), a.projectID)
		}
	}

	logMetric := convertKCCtoAPI(&a.desired.Spec)
	logMetric.Name = a.resourceID

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

	status := &krm.LoggingLogMetricStatus{}
	if err := logMetricStatusToKRM(created, status); err != nil {
		return err
	}

	return setStatus(u, status)
}

func logMetricStatusToKRM(in *api.LogMetric, out *krm.LoggingLogMetricStatus) error {
	out.CreateTime = direct.LazyPtr(in.CreateTime)
	out.UpdateTime = direct.LazyPtr(in.UpdateTime)

	out.MetricDescriptor = convertAPItoKRM_MetricDescriptorStatus(in.MetricDescriptor)

	return nil
}

func (a *logMetricAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)

	latest := a.actual

	if a.hasChanges(ctx, u) {
		update := new(api.LogMetric)
		*update = *a.actual

		if direct.ValueOf(a.desired.Spec.Description) != a.actual.Description {
			update.Description = direct.ValueOf(a.desired.Spec.Description)
		}
		if direct.ValueOf(a.desired.Spec.Disabled) != a.actual.Disabled {
			update.Disabled = direct.ValueOf(a.desired.Spec.Disabled)
		}
		if a.desired.Spec.Filter != a.actual.Filter {
			// todo acpana: revisit UX, err out if filter of desired is empty
			if a.desired.Spec.Filter != "" {
				update.Filter = a.desired.Spec.Filter
			} else {
				// filter is a REQUIRED field
				update.Filter = a.actual.Filter
			}
		}

		if !compareMetricDescriptors(a.desired.Spec.MetricDescriptor, a.actual.MetricDescriptor) {
			if err := validateImmutableFieldsUpdated(a.desired.Spec.MetricDescriptor, a.actual.MetricDescriptor); err != nil {
				return fmt.Errorf("logMetric update failed: %w", err)
			}
			update.MetricDescriptor = convertKCCtoAPIForMetricDescriptor(a.desired.Spec.MetricDescriptor)
		}

		if !reflect.DeepEqual(a.desired.Spec.LabelExtractors, a.actual.LabelExtractors) {
			update.LabelExtractors = a.desired.Spec.LabelExtractors
		}

		if !compareBucketOptions(a.desired.Spec.BucketOptions, a.actual.BucketOptions) {
			update.BucketOptions = convertKCCtoAPIForBucketOptions(a.desired.Spec.BucketOptions)
		}

		if direct.ValueOf(a.desired.Spec.ValueExtractor) != a.actual.ValueExtractor {
			update.ValueExtractor = direct.ValueOf(a.desired.Spec.ValueExtractor)
		}
		if a.desired.Spec.LoggingLogBucketRef != nil && a.desired.Spec.LoggingLogBucketRef.External != a.actual.BucketName {
			update.BucketName = a.desired.Spec.LoggingLogBucketRef.External
		}

		diffs, err := ListFieldDiffs(a.actual, update)
		if err != nil {
			// Don't return an error as we're only logging
			log.Error(err, "computing changed field paths (for logging)")
		}
		log.Info("updating logMetric", "diffs", diffs)

		// DANGER: this is an upsert; it will create the LogMetric if it doesn't exists
		// but this behavior is consistent with the DCL backed behavior we provide for this resource.
		// todo acpana: look for / switch to a better method and/or use etags etc
		updated, err := a.logMetricClient.Update(a.fullyQualifiedName(), update).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("logMetric update failed: %w", err)
		}
		latest = updated
	}

	status := &krm.LoggingLogMetricStatus{}
	if err := logMetricStatusToKRM(latest, status); err != nil {
		return err
	}

	// actualUpdate may not contain the description for the metric descriptor.
	if latest.Description != "" {
		if status.MetricDescriptor != nil {
			status.MetricDescriptor.Description = &latest.Description
		}
	}

	return setStatus(u, status)
}

func (a *logMetricAdapter) hasChanges(ctx context.Context, u *unstructured.Unstructured) bool {
	log := klog.FromContext(ctx)

	if u.GetGeneration() != getObservedGeneration(u) {
		log.V(2).Info("generation does not match", "generation", u.GetGeneration(), "observedGeneration", getObservedGeneration(u))
		return true
	}

	gcpUpdateTime := a.actual.UpdateTime
	if gcpUpdateTime == "" {
		log.V(2).Info("updateTime is not set in GCP")
		return true
	}

	obj := &krm.LoggingLogMetric{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		log.Error(err, "error converting from unstructured")
		return true
	}

	// if observed generation matches and status is initially empty after creation, no changes required yet.
	if obj.Status.UpdateTime == nil {
		log.V(2).Info("status.updateTime is not set")
		return false
	}
	if gcpUpdateTime != direct.ValueOf(obj.Status.UpdateTime) {
		log.V(2).Info("status.updateTime does not match gcp updateTime", "status.updateTime", obj.Status.UpdateTime, "gcpUpdateTime", gcpUpdateTime)
		return true
	}

	if obj.Status.Conditions != nil {
		// if there was a previously failing update let's make sure we give
		// the update a chance to heal or keep marking it as failed

		ready := false
		for _, condition := range obj.Status.Conditions {
			if condition.Type == v1alpha1.ReadyConditionType {
				if condition.Status == corev1.ConditionTrue {
					ready = true
				}
			}
		}

		if !ready {
			log.V(2).Info("status.conditions indicates object is not ready yet")
			return true
		}
	}

	log.V(2).Info("status.updateTime matches gcp updateTime", "status.updateTime", obj.Status.UpdateTime, "gcpUpdateTime", gcpUpdateTime)
	return false
}

func (a *logMetricAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("logMetric %q not found", a.fullyQualifiedName())
	}

	un, err := convertAPItoKRM_LoggingLogMetric(a.projectID, a.actual)
	if err != nil {
		return nil, fmt.Errorf("error converting logMetric to unstructured %w", err)
	}

	// TODO(acpana): revisit if we want to include mutable but unreadable fields in our export
	if a.desired != nil {
		if a.desired.Spec.MetricDescriptor != nil && a.desired.Spec.MetricDescriptor.LaunchStage != nil {
			if err := unstructured.SetNestedField(un.Object,
				*a.desired.Spec.MetricDescriptor.LaunchStage,
				"spec", "metricDescriptor", "launchStage",
			); err != nil {
				return nil, fmt.Errorf("could not set metricDescriptor.launchStage mutable but unreadable field %w", err)
			}
		}
		if a.desired.Spec.MetricDescriptor != nil && a.desired.Spec.MetricDescriptor.Metadata != nil {
			if a.desired.Spec.MetricDescriptor.Metadata.IngestDelay != nil {
				if err := unstructured.SetNestedField(un.Object,
					*a.desired.Spec.MetricDescriptor.Metadata.IngestDelay,
					"spec", "metricDescriptor", "metadata", "ingestDelay",
				); err != nil {
					return nil, fmt.Errorf("could not set metricDescriptor.metadata.ingestDelay mutable but unreadable field %w", err)
				}
			}
			if a.desired.Spec.MetricDescriptor.Metadata.SamplePeriod != nil {
				if err := unstructured.SetNestedField(un.Object,
					*a.desired.Spec.MetricDescriptor.Metadata.SamplePeriod,
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
	return MakeFQN(a.projectID, a.resourceID)
}

// MakeFQN constructions a fully qualified name for a LogMetric resources
// to be used in API calls. The format expected is: "projects/[PROJECT_ID]/metrics/[METRIC_ID]".
// Func assumes values are well formed and validated.
func MakeFQN(projectID, metricID string) string {
	return fmt.Sprintf("projects/%s/metrics/%s", projectID, metricID)
}
