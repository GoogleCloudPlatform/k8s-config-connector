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
	"strings"

	api "cloud.google.com/go/securesourcemanager/apiv1"
	pb "cloud.google.com/go/securesourcemanager/apiv1/securesourcemanagerpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/refresolver"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securesourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.SecureSourceManagerInstanceGVK, newSecureSourceManagerInstanceModel)
}

func newSecureSourceManagerInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &SecureSourceManagerInstanceModel{config: config}, nil
}

type SecureSourceManagerInstanceModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &SecureSourceManagerInstanceModel{}

type instanceAdapter struct {
	projectID  string
	location   string
	resourceID string

	desired *pb.Instance
	actual  *pb.Instance

	client *api.Client
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &instanceAdapter{}

// AdapterForObject implements the Model interface.
func (m *SecureSourceManagerInstanceModel) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	ssmClient, err := gcpClient.newClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.SecureSourceManagerInstance{}
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

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectRef, err := refs.ResolveProject(ctx, kube, obj, &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if err := refresolver.ResolveReferences(ctx, kube, obj, projectRef); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := SecureSourceManagerInstanceSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &instanceAdapter{
		projectID:  projectID,
		location:   location,
		resourceID: resourceID,
		desired:    desiredProto,
		client:     ssmClient,
	}, nil
}

func (m *SecureSourceManagerInstanceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// The url format should match the Cloud-Asset-Inventory format: https://cloud.google.com/asset-inventory/docs/resource-name-format
	// Format: //securesourcemanager.googleapis.com/projects/PROJECT_NUMBER/dashboards/DASHBOARD_ID
	if !strings.HasPrefix(url, "//securesourcemanager.googleapis.com/") {
		return nil, nil
	}

	tokens := strings.Split(strings.TrimPrefix(url, "//securesourcemanager.googleapis.com/"), "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "instances" {
		gcpClient, err := newGCPClient(ctx, m.config)
		if err != nil {
			return nil, fmt.Errorf("building gcp client: %w", err)
		}

		client, err := gcpClient.newClient(ctx)
		if err != nil {
			return nil, err
		}

		return &instanceAdapter{
			projectID:  tokens[1],
			location:   tokens[3],
			resourceID: tokens[5],
			client:     client,
		}, nil
	}

	return nil, nil
}

// Find implements the Adapter interface.
func (a *instanceAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	req := &pb.GetInstanceRequest{
		Name: a.fullyQualifiedName(),
	}
	instance, err := a.client.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = instance

	return true, nil
}

// Delete implements the Adapter interface.
func (a *instanceAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deleted
	if a.resourceID == "" {
		return false, nil
	}

	// TODO: Delete via status selfLink?
	req := &pb.DeleteInstanceRequest{
		Name: a.fullyQualifiedName(),
	}

	op, err := a.client.DeleteInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting SecureSourceManagerInstance %s: %w", a.fullyQualifiedName(), err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for SecureSourceManagerInstance delete %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

func (a *instanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("instance %q not found", a.fullyQualifiedName())
	}

	mc := &direct.MapContext{}
	spec := SecureSourceManagerInstanceSpec_FromProto(mc, a.actual)
	if err := mc.Err(); err != nil {
		return nil, fmt.Errorf("building SecureSourceManagerInstance from API: %w", err)
	}

	spec.ProjectRef.External = a.projectID

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("converting SecureSourceManagerInstance to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.resourceID)
	u.SetGroupVersionKind(krm.SecureSourceManagerInstanceGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Create implements the Adapter interface.
func (a *instanceAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	parent := "projects/" + a.projectID + "/locations/" + a.location

	instance := proto.Clone(a.desired).(*pb.Instance)

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

	// For new resources, we do not write spec.resourceID

	mapCtx := &direct.MapContext{}
	observedState := SecureSourceManagerInstanceObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return direct.SetObservedState(u, observedState)
}

// Update implements the Adapter interface.
func (a *instanceAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	// return fmt.Errorf("update of SecureSourceManagerInstance not supported")
	klog.Infof("update of SecureSourceManagerInstance not supported")
	return nil
}

func (a *instanceAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/instances/%s", a.projectID, a.location, a.resourceID)
}
