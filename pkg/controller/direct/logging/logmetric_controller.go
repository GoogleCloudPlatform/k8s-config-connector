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
	"time"

	api "google.golang.org/api/logging/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resources/logging/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

const ctrlName = "logmetric-controller"

// AddLogMetricController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddLogMetricController(mgr manager.Manager, config *controller.Config, opts directbase.Deps) error {
	gvk := krm.LoggingLogMetricGVK

	// todo(acpana): plumb context throughout direct
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &logMetricModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m, opts)
}

type logMetricModel struct {
	*gcpClient
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
	projectMetricsService, err := m.newProjectMetricsService(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.LoggingLogMetric{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	resourceID := ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	projectRef, err := ResolveProject(ctx, reader, obj, &obj.Spec.ProjectRef)
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

func (a *logMetricAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	logMetric, err := a.logMetricClient.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting logMetric %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = logMetric

	return true, nil
}

// Delete implements the Adapter interface.
func (a *logMetricAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deleted
	if a.resourceID == "" {
		return false, nil
	}

	_, err := a.logMetricClient.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting log metric %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

func (a *logMetricAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating object", "u", u)

	projectID := a.projectID
	if projectID == "" {
		return fmt.Errorf("project is empty")
	}
	name := a.desired.GetName()
	if name == "" {
		return fmt.Errorf("name is empty")
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

	logMetric := convertKCCtoAPI(a.desired)

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
	out.CreateTime = nil
	if in.CreateTime != "" {
		parsed, err := time.Parse(time.RFC3339, in.CreateTime)
		if err != nil {
			return fmt.Errorf("cannot parse createTime %q: %w", in.CreateTime, err)
		}
		mt := metav1.NewTime(parsed.UTC())
		out.CreateTime = &mt
	}

	out.UpdateTime = nil
	if in.UpdateTime != "" {
		parsed, err := time.Parse(time.RFC3339, in.UpdateTime)
		if err != nil {
			return fmt.Errorf("cannot parse updateTime %q: %w", in.UpdateTime, err)
		}
		mt := metav1.NewTime(parsed.UTC())
		out.UpdateTime = &mt
	}

	out.MetricDescriptor = convertAPItoKRM_MetricDescriptorStatus(in.MetricDescriptor)

	return nil
}

func (a *logMetricAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	update := new(api.LogMetric)
	*update = *a.actual

	if ValueOf(a.desired.Spec.Description) != a.actual.Description {
		update.Description = ValueOf(a.desired.Spec.Description)
	}
	if ValueOf(a.desired.Spec.Disabled) != a.actual.Disabled {
		update.Disabled = ValueOf(a.desired.Spec.Disabled)
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
		update.MetricDescriptor = convertKCCtoAPIForMetricDescriptor(a.desired.Spec.MetricDescriptor)
	}

	if !reflect.DeepEqual(a.desired.Spec.LabelExtractors, a.actual.LabelExtractors) {
		update.LabelExtractors = a.desired.Spec.LabelExtractors
	}

	if !compareBucketOptions(a.desired.Spec.BucketOptions, a.actual.BucketOptions) {
		update.BucketOptions = convertKCCtoAPIForBucketOptions(a.desired.Spec.BucketOptions)
	}

	if ValueOf(a.desired.Spec.ValueExtractor) != a.actual.ValueExtractor {
		update.ValueExtractor = ValueOf(a.desired.Spec.ValueExtractor)
	}
	if a.desired.Spec.LoggingLogBucketRef != nil && a.desired.Spec.LoggingLogBucketRef.External != a.actual.BucketName {
		update.BucketName = a.desired.Spec.LoggingLogBucketRef.External
	}

	// DANGER: this is an upsert; it will create the LogMetric if it doesn't exists
	// but this behavior is consistent with the DCL backed behavior we provide for this resource.
	// todo acpana: look for / switch to a better method and/or use etags etc
	actualUpdate, err := a.logMetricClient.Update(a.fullyQualifiedName(), update).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("logMetric update failed: %w", err)
	}

	status := &krm.LoggingLogMetricStatus{}
	if err := logMetricStatusToKRM(actualUpdate, status); err != nil {
		return err
	}

	// actualUpdate may not contain the description for the metric descriptor.
	if update.Description != "" {
		if status.MetricDescriptor != nil {
			status.MetricDescriptor.Description = &update.Description
		}
	}

	return setStatus(u, status)
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
