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

package securesourcemanager

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"k8s.io/klog/v2"

	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/securesourcemanager/apis/v1alpha1"

	ssm "cloud.google.com/go/securesourcemanager/apiv1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings"
)

// Add creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddInstanceController(mgr manager.Manager, config *controller.Config) error {
	gvk := krm.SecureSourceManagerInstanceGVK

	return directbase.Add(mgr, gvk, &instanceModel{config: *config})
}

type instanceModel struct {
	config controller.Config
}

var instanceMapping = NewMapping(&pb.Instance{}, &krm.SecureSourceManagerInstance{},
	SpecResourceRef("kmsKey", &refMapKMSKey{}),

	Status("state"),
	Status("stateNote"),
	Status("hostConfig"),

	Ignore("createTime"), // Not interesting
	Ignore("updateTime"), // Not interesting

	Ignore("projectRef"), // Handled in code
	Ignore("resourceID"), // Handled in code
	Ignore("location"),   // Handled in code

	TODO("labels"),
).
	MapNested(&pb.Instance_HostConfig{}, &krm.SecureSourceManagerInstance_HostConfig{}, "html", "api", "gitHttp", "gitSsh").
	MustBuild()

type instanceAdapter struct {
	projectID  string
	location   string
	instanceID string

	desired *krm.SecureSourceManagerInstance
	actual  *krm.SecureSourceManagerInstance

	gcp *ssm.Client
}

func (m *instanceModel) client(ctx context.Context) (*ssm.Client, error) {
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

	// if m.config.Endpoint != "" {
	// 	opts = append(opts, option.WithEndpoint(m.config.Endpoint))
	// }

	gcpClient, err := ssm.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building instance client: %w", err)
	}
	return gcpClient, err
}

func (m *instanceModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.SecureSourceManagerInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("unable to determine location")
	}

	instanceID := obj.Spec.ResourceID
	if instanceID == "" {
		instanceID = obj.GetName()
	}
	if instanceID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	return &instanceAdapter{
		projectID:  projectID,
		location:   location,
		instanceID: instanceID,
		desired:    obj,
		gcp:        gcp,
	}, nil
}

func (a *instanceAdapter) Find(ctx context.Context) (bool, error) {
	if a.instanceID == "" {
		return false, nil
	}

	req := &pb.GetInstanceRequest{
		Name: a.fullyQualifiedName(),
	}
	gcpObject, err := a.gcp.GetInstance(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			klog.Warningf("instance was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.SecureSourceManagerInstance{}
	if err := instanceMapping.Map(gcpObject, u, nil); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

func (a *instanceAdapter) Delete(ctx context.Context) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteInstanceRequest{
		Name: a.fullyQualifiedName(),
	}
	op, err := a.gcp.DeleteInstance(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting instance: %w", err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for instance deletion: %w", err)
	}

	return true, nil
}

func (a *instanceAdapter) Create(ctx context.Context, obj *unstructured.Unstructured) error {
	desired := &pb.Instance{}
	if err := instanceMapping.MapSpec(a.desired, desired); err != nil {
		return err
	}

	desired.Name = a.fullyQualifiedName()

	req := &pb.CreateInstanceRequest{
		Parent:     fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location),
		InstanceId: a.instanceID,
		Instance:   desired,
	}

	op, err := a.gcp.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating instance: %w", err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for instance creation: %w", err)
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("created instance", "instance", created)
	// TODO: Return created object
	return nil
}

func (a *instanceAdapter) Update(ctx context.Context) (*unstructured.Unstructured, error) {
	// desired := &pb.Instance{}
	// if err := instanceMapping.Map(a.desired, desired); err != nil {
	// 	return nil, err
	// }

	// desired.Name = a.fullyQualifiedName()

	// req := &pb.UpdateInstanceRequest{
	// 	Instance: desired,
	// }

	// updated, err := a.gcp.UpdateInstance(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }
	// log := klog.FromContext(ctx)
	// log.V(2).Info("updated instance", "instance", updated)
	// // TODO: Return updated object
	// return nil, nil
	return nil, fmt.Errorf("update is not supported")
}

func (a *instanceAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", a.projectID, a.location, a.instanceID)
}
