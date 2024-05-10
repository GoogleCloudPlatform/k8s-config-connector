// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package spanner

import (
	"context"
	"fmt"
	"reflect"
	"regexp"

	instanceapi "cloud.google.com/go/spanner/admin/instance/apiv1"
	"cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

var spannerInstanceIDRegexp = regexp.MustCompile("projects/(.+)/instances/(.+)")
var spannerInstanceConfigIDRegexp = regexp.MustCompile("projects/(.+)/instanceConfigs/(.+)")

// AddKeyReconciler creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddSpannerInstanceReconciler(mgr manager.Manager, config *controller.Config, opts directbase.Deps) error {
	gvk := krm.SpannerInstanceGVK

	return directbase.Add(mgr, gvk, &model{config: *config}, opts)
}

// model implements the Model interface.
var _ directbase.Model = &model{}

type model struct {
	config controller.Config
}

func (m *model) client(ctx context.Context) (*instanceapi.InstanceAdminClient, error) {

	var opts []option.ClientOption
	if m.config.UserAgent != "" {
		opts = append(opts, option.WithUserAgent(m.config.UserAgent))
	}
	if m.config.HTTPClient != nil {
		opts = append(opts, option.WithHTTPClient(m.config.HTTPClient))
	}
	if m.config.UserProjectOverride && m.config.BillingProject != "" {
		opts = append(opts, option.WithQuotaProject(m.config.BillingProject))
	}

	gcpClient, err := instanceapi.NewInstanceAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SpannerInstance client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.SpannerInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// TODO(yuwenma): following current behavior. But do we have better option?
	instanceID := directbase.ValueOf(obj.Spec.ResourceID)
	if instanceID == "" {
		instanceID = obj.GetName()
	}

	// TODO(yuwenma): following current behavior. But do we have better option?
	projectID, ok := u.GetAnnotations()[k8s.ProjectIDAnnotation]
	if !ok {
		projectID = u.GetNamespace()
	}
	return &adapter{
		projectID:  projectID,
		InstanceID: instanceID,
		desired:    obj,
		gcp:        gcp,
	}, nil
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &adapter{}

type adapter struct {
	projectID  string
	InstanceID string

	desired *krm.SpannerInstance
	actual  *krm.SpannerInstance

	gcp *instanceapi.InstanceAdminClient
}

// Find implements the Adapter interface.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	if a.InstanceID == "" {
		return false, nil
	}

	req := &instancepb.GetInstanceRequest{
		Name: a.fullyQualifiedName(),
	}
	instance, err := a.gcp.GetInstance(ctx, req)
	if err != nil {
		if directbase.IsNotFound(err) {
			klog.Warningf("SpannerInstance was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.SpannerInstance{}
	if err := Convert_v1_SpannerInstance_API_To_v1beta1_SpannerInstance_KRM(instance, u, a); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink
	req := &instancepb.DeleteInstanceRequest{
		Name: a.fullyQualifiedName(),
	}
	if err := a.gcp.DeleteInstance(ctx, req); err != nil {
		if directbase.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting key: %w", err)
	}
	return true, nil
}

// Create implements the Adapter interface.
func (a *adapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	desired := &instancepb.Instance{}

	if err := Convert_v1beta1_SpannerInstance_KRM_To_v1_SpannerInstance_API(a.desired, desired, a); err != nil {
		return err
	}
	req := &instancepb.CreateInstanceRequest{
		Parent:     fmt.Sprintf("projects/%s", a.projectID),
		InstanceId: a.InstanceID,
		Instance:   desired,
	}

	op, err := a.gcp.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating spannerInstance: %w", err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for spannerInstance creation: %w", err)
	}
	log.V(2).Info("created spannerInstance", "spannerInstance", created)
	// TODO: Return created object
	return nil
}

// Update implements the Adapter interface.
func (a *adapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	// TODO: Skip updates if no changes
	// TODO: Where/how do we want to enforce immutability?
	updateMask := &fieldmaskpb.FieldMask{}

	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.Spec.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(a.desired.Spec.NumNodes, a.actual.Spec.NumNodes) {
		updateMask.Paths = append(updateMask.Paths, "node_count")
	}
	if !reflect.DeepEqual(a.desired.Spec.ProcessingUnits, a.actual.Spec.ProcessingUnits) {
		updateMask.Paths = append(updateMask.Paths, "processing_units")
	}
	if !reflect.DeepEqual(a.desired.Spec.Config, a.actual.Spec.Config) {
		updateMask.Paths = append(updateMask.Paths, "config")
	}

	// TODO: Annotations
	// if !reflect.DeepEqual(a.desired.Annotations, a.actual.Annotations) {
	// 	updateMask.Paths = append(updateMask.Paths, "annotations")
	// }

	if len(updateMask.Paths) == 0 {
		klog.Warningf("unexpected empty update mask, desired: %v, actual: %v", a.desired, a.actual)
		return nil
	}

	instance := &instancepb.Instance{}
	if err := Convert_v1beta1_SpannerInstance_KRM_To_v1_SpannerInstance_API(a.desired, instance, a); err != nil {
		return err
	}

	req := &instancepb.UpdateInstanceRequest{
		Instance:  instance,
		FieldMask: updateMask,
	}

	req.Instance.Name = a.fullyQualifiedName()

	_, err := a.gcp.UpdateInstance(ctx, req)
	if err != nil {

		return err
	}
	// TODO: update status in u
	return nil
}

func (a *adapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/instances/%s", a.projectID, a.InstanceID)
}

func Convert_v1beta1_SpannerInstance_KRM_To_v1_SpannerInstance_API(in *krm.SpannerInstance, out *instancepb.Instance, i directbase.Adapter) error {
	// TODO(yuwenma): auto conversion
	out.DisplayName = in.Spec.DisplayName
	if in.Spec.NumNodes != nil {
		out.NodeCount = int32(*in.Spec.NumNodes)
	}
	if in.Spec.ProcessingUnits != nil {
		out.ProcessingUnits = int32(*in.Spec.ProcessingUnits)
	}

	// custom update
	a, ok := i.(*adapter)
	if !ok {
		return fmt.Errorf("unable to cast %s to adapter", i)
	}
	out.Name = a.fullyQualifiedName()
	out.Config = fmt.Sprintf("projects/%s/instanceConfigs/%s", a.projectID, in.Spec.Config)
	return nil
}

func Convert_v1_SpannerInstance_API_To_v1beta1_SpannerInstance_KRM(in *instancepb.Instance, out *krm.SpannerInstance, a directbase.Adapter) error {
	// TODO(yuwenma): auto conversion
	out.Name = in.Name
	out.Spec.DisplayName = in.DisplayName
	out.Spec.NumNodes = new(int)
	*out.Spec.NumNodes = int(in.NodeCount)
	out.Spec.ProcessingUnits = new(int)
	*out.Spec.ProcessingUnits = int(in.ProcessingUnits)

	// custom update
	segments := spannerInstanceConfigIDRegexp.FindStringSubmatch(in.Config)
	out.Spec.Config = segments[2]
	return nil
}
