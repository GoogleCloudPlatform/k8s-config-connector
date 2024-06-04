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

	api "cloud.google.com/go/securesourcemanager/apiv1"
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"google.golang.org/protobuf/proto"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/references"
)

func init() {
	directbase.ControllerBuilder.RegisterModel(krm.SecureSourceManagerInstanceGVK, newInstanceModel)
}

func newInstanceModel(config *controller.Config) directbase.Model {
	return &instanceModel{config: config}
}

type instanceModel struct {
	// *gcpClient
	config *controller.Config
}

// model implements the Model interface.
var _ directbase.Model = &instanceModel{}

type clusterAdapter struct {
	projectID  string
	location   string
	resourceID string

	desired *pb.Instance
	actual  *pb.Instance

	client *api.Client
}

var _ directbase.Adapter = &clusterAdapter{}

// AdapterForObject implements the Model interface.
func (m *instanceModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	klog.FromContext(ctx).V(0).Info("creating adapter", "u", u)
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.SecureSourceManagerInstance{}
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

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectRef, err := references.ResolveProject(ctx, reader, obj, &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	mapCtx := &MapContext{}
	desired := SecureSourceManagerInstanceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &clusterAdapter{
		resourceID: resourceID,
		projectID:  projectID,
		location:   location,
		desired:    desired,
		client:     client,
	}, nil
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &clusterAdapter{}

// Find implements the Adapter interface.
func (a *clusterAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}
	req := &pb.GetInstanceRequest{
		Name: a.fullyQualifiedName(),
	}
	instance, err := a.client.GetInstance(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = instance

	return true, nil
}

// Delete implements the Adapter interface.
func (a *clusterAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deleted
	if a.resourceID == "" {
		return false, nil
	}

	req := &pb.DeleteInstanceRequest{
		Name: a.fullyQualifiedName(),
	}
	op, err := a.client.DeleteInstance(ctx, req)
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting instance %s: %w", a.fullyQualifiedName(), err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for instance deletion: %w", err)
	}
	return true, nil
}

// Create implements the Adapter interface.
func (a *clusterAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	parent := "projects/" + a.projectID + "/locations/" + a.location

	instance := proto.Clone(a.desired).(*pb.Instance)
	// instance.Name = a.fullyQualifiedName()

	req := &pb.CreateInstanceRequest{
		Parent:     parent,
		InstanceId: a.resourceID,
		Instance:   instance,
	}
	log.V(0).Info("creating instance", "request", req)
	op, err := a.client.CreateInstance(ctx, req)
	if err != nil {
		return fmt.Errorf("creating instance %s: %w", a.fullyQualifiedName(), err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for instance creation: %w", err)
	}

	log.V(0).Info("created instance", "instance", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	mapCtx := &MapContext{}
	observedState := SecureSourceManagerInstanceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

// Update implements the Adapter interface.
func (a *clusterAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	// return fmt.Errorf("update of SecureSourceManagerInstance not supported")
	klog.Infof("update of SecureSourceManagerInstance not supported")
	return nil
}

func (a *clusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("unimplemented")
}

func (a *clusterAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", a.projectID, a.location, a.resourceID)
}
