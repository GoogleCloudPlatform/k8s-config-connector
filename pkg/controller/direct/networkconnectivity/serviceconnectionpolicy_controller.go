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

package networkconnectivity

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	api "google.golang.org/api/networkconnectivity/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
)

// AddServiceConnectionPolicyController creates a new controller and adds it to the Manager.
// The Manager will set fields on the Controller and start it when the Manager is started.
func AddServiceConnectionPolicyController(mgr manager.Manager, config *controller.Config, opts directbase.Deps) error {
	gvk := krm.ServiceConnectionPolicyGVK

	// TODO: Share gcp client (any value in doing so)?
	ctx := context.TODO()
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return err
	}
	m := &serviceConnectionPolicyModel{gcpClient: gcpClient}
	return directbase.Add(mgr, gvk, m, opts)
}

type serviceConnectionPolicyModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &serviceConnectionPolicyModel{}

type serviceConnectionPolicyAdapter struct {
	projectID  string
	location   string
	resourceID string

	desiredProto *api.ServiceConnectionPolicy
	actual       *api.ServiceConnectionPolicy

	*gcpClient
	client *api.Service
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &serviceConnectionPolicyAdapter{}

// AdapterForObject implements the Model interface.
func (m *serviceConnectionPolicyModel) AdapterForObject(ctx context.Context, u *unstructured.Unstructured) (directbase.Adapter, error) {
	networkConnectivityClient, err := m.newNetworkConnectivityClient(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.NetworkConnectivityServiceConnectionPolicy{}
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

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	{
		tokens := strings.Split(projectID, "/")
		if len(tokens) == 1 {
			projectID = tokens[0]
		} else if len(tokens) == 2 && tokens[0] == "projects" {
			projectID = tokens[1]
		} else {
			return nil, fmt.Errorf("cannot resolve project from name %q", projectID)
		}
	}

	mapCtx := &MapContext{
		//	kube: kube,
	}
	desiredProto := ServiceConnectionPolicySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &serviceConnectionPolicyAdapter{
		projectID:  projectID,
		location:   location,
		resourceID: resourceID,
		// desired:          obj,
		desiredProto: desiredProto,
		gcpClient:    m.gcpClient,
		client:       networkConnectivityClient,
	}, nil
}

// Find implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	name := a.fullyQualifiedName()
	found, err := a.client.Projects.Locations.ServiceConnectionPolicies.Get(name).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = found

	return true, nil
}

// Delete implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Delete(ctx context.Context) (bool, error) {
	// Already deleted
	if a.resourceID == "" {
		return false, nil
	}

	// TODO: Delete via status selfLink?

	op, err := a.client.Projects.Locations.ServiceConnectionPolicies.Delete(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		if IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting serviceConnectionPolicy %s: %w", a.fullyQualifiedName(), err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for serviceConnectionPolicy delete %s: %w", a.fullyQualifiedName(), err)
	}

	return true, nil
}

func (a *serviceConnectionPolicyAdapter) waitForOp(ctx context.Context, op *api.GoogleLongrunningOperation) error {
	// TODO: Only wait a short time, then update status?
	for {
		current, err := a.client.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation status of %q: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error)
			} else {
				return nil
			}
		}
		time.Sleep(2 * time.Second)
	}
}

// Create implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("creating object", "u", u)

	// TODO: Should be ref
	// parent := a.desired.Spec.Parent

	parent := "projects/" + a.projectID + "/locations/" + a.location

	// TODO: Deep copy?
	policy := a.desiredProto

	log.V(0).Info("creating ServiceConnectionPolicy", "policy", policy)
	op, err := a.client.Projects.Locations.ServiceConnectionPolicies.Create(parent, policy).ServiceConnectionPolicyId(a.resourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating serviceConnectionPolicy: %w", err)
	}

	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for serviceConnectionPolicy create %s: %w", a.fullyQualifiedName(), err)
	}

	created, err := a.client.Projects.Locations.ServiceConnectionPolicies.Get(a.fullyQualifiedName()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created serviceConnectionPolicy: %w", err)
	}

	log.V(0).Info("created serviceConnectionPolicy", "serviceConnectionPolicy", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	mapCtx := &MapContext{
		// kube: kube,
	}
	observedState := ServiceConnectionPolicyState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

// Update implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(0).Info("updating object", "u", u)

	// TODO: Where/how do we want to enforce immutability?

	var latest *api.ServiceConnectionPolicy

	updateMask := &fieldmaskpb.FieldMask{}

	if (a.desiredProto.Description) != (a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(a.desiredProto.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if (a.desiredProto.Network) != (a.actual.Network) {
		updateMask.Paths = append(updateMask.Paths, "network")
	}
	if !reflect.DeepEqual(a.desiredProto.PscConfig, a.actual.PscConfig) {
		updateMask.Paths = append(updateMask.Paths, "psc_config")
	}
	if (a.desiredProto.ServiceClass) != (a.actual.ServiceClass) {
		updateMask.Paths = append(updateMask.Paths, "service_class")
	}

	if len(updateMask.Paths) != 0 {

		// TODO: Deep copy?
		policy := a.desiredProto

		op, err := a.client.Projects.Locations.ServiceConnectionPolicies.Patch(a.fullyQualifiedName(), policy).UpdateMask(strings.Join(updateMask.Paths, ",")).Context(ctx).Do()
		if err != nil {
			return err
		}

		if err := a.waitForOp(ctx, op); err != nil {
			return fmt.Errorf("waiting for serviceConnectionPolicy update %s: %w", a.fullyQualifiedName(), err)
		}

		updated, err := a.client.Projects.Locations.ServiceConnectionPolicies.Get(a.fullyQualifiedName()).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated serviceConnectionPolicy: %w", err)
		}
		log.V(0).Info("updated serviceConnectionPolicy", "serviceConnectionPolicy", updated)

		latest = updated
	} else {
		latest = a.actual
	}

	mapCtx := &MapContext{
		// kube: kube,
	}
	observedState := ServiceConnectionPolicyState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

func (a *serviceConnectionPolicyAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", a.projectID, a.location, a.resourceID)
}
