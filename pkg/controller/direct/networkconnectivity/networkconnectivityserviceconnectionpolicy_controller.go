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

// +tool:controller
// proto.service: google.cloud.networkconnectivity.v1.CrossNetworkAutomation
// proto.message: google.cloud.networkconnectivity.v1.ServiceConnectionPolicy
// crd.type: NetworkConnectivityServiceConnectionPolicy
// crd.version: v1alpha1

package networkconnectivity

import (
	"context"
	"fmt"
	"reflect"

	gcpapi "cloud.google.com/go/networkconnectivity/apiv1"
	pb "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/monitoring"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.NetworkConnectivityServiceConnectionPolicyGVK, newServiceConnectionPolicyModel)
}

func newServiceConnectionPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &serviceConnectionPolicyModel{config: config}, nil
}

type serviceConnectionPolicyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &serviceConnectionPolicyModel{}

type serviceConnectionPolicyAdapter struct {
	projectID  string
	location   string
	resourceID string

	desired *pb.ServiceConnectionPolicy
	actual  *pb.ServiceConnectionPolicy

	gcpClient *gcpapi.CrossNetworkAutomationClient
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &serviceConnectionPolicyAdapter{}

// AdapterForObject implements the Model interface.
func (m *serviceConnectionPolicyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	clientBuilder, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	gcpClient, err := clientBuilder.newCrossNetworkAutomationClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("building crossnetworkautomation client: %w", err)
	}

	desired := &krm.NetworkConnectivityServiceConnectionPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, desired); err != nil {
		return nil, fmt.Errorf("error converting from unstructured: %w", err)
	}

	projectRef, err := refs.ResolveProject(ctx, kube, u.GetNamespace(), &desired.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID

	location := direct.ValueOf(desired.Spec.Location)
	resourceID := direct.ValueOf(desired.Spec.ResourceID)
	if resourceID == "" {
		resourceID = u.GetName()
	}

	if err := common.VisitFields(desired, &refNormalizer{ctx: ctx, src: desired, project: *projectRef, kube: kube}); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkConnectivityServiceConnectionPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &serviceConnectionPolicyAdapter{
		projectID:  projectID,
		resourceID: resourceID,
		location:   location,
		desired:    desiredProto,
		gcpClient:  gcpClient,
	}, nil
}

func (m *serviceConnectionPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *serviceConnectionPolicyAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", a.projectID, a.location, a.resourceID)
}

func (a *serviceConnectionPolicyAdapter) parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location)
}

// Find implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	fqn := a.fullyQualifiedName()
	req := &pb.GetServiceConnectionPolicyRequest{
		Name: fqn,
	}
	actual, err := a.gcpClient.GetServiceConnectionPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	a.actual = actual
	return true, nil
}

// Delete implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	exists, err := a.Find(ctx)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	fqn := a.fullyQualifiedName()
	req := &pb.DeleteServiceConnectionPolicyRequest{
		Name: fqn,
	}
	op, err := a.gcpClient.DeleteServiceConnectionPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting serviceConnectionPolicy %q: %w", fqn, err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for delete of serviceConnectionPolicy %q: %w", fqn, err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	fqn := a.fullyQualifiedName()

	req := &pb.CreateServiceConnectionPolicyRequest{
		Parent:                    a.parent(),
		ServiceConnectionPolicyId: a.resourceID,
		ServiceConnectionPolicy:   a.desired,
	}

	log.V(0).Info("creating serviceConnectionPolicy", "req", req)
	op, err := a.gcpClient.CreateServiceConnectionPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating serviceConnectionPolicy: %w", err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for create of serviceConnectionPolicy %q: %w", fqn, err)
	}
	log.V(2).Info("created serviceConnectionPolicy", "serviceConnectionPolicy", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	mapCtx := &direct.MapContext{}
	observedState := NetworkConnectivityServiceConnectionPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

// Update implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating object", "u", u)

	fqn := a.fullyQualifiedName()

	if monitoring.ShouldReconcileBasedOnEtag(ctx, u, a.actual.GetEtag()) {
		paths := []string{}
		if !reflect.DeepEqual(a.desired.PscConfig, a.actual.PscConfig) {
			paths = append(paths, "psc_config")
		}
		if a.desired.Description != a.actual.Description {
			paths = append(paths, "description")
		}

		req := &pb.UpdateServiceConnectionPolicyRequest{
			ServiceConnectionPolicy: a.desired,
			UpdateMask:              &fieldmaskpb.FieldMask{Paths: paths},
		}
		a.desired.Name = fqn

		log.V(2).Info("updating serviceConnectionPolicy", "request", req)
		op, err := a.gcpClient.UpdateServiceConnectionPolicy(ctx, req)
		if err != nil {
			return err
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of serviceConnectionPolicy %q: %w", fqn, err)
		}

		log.V(2).Info("updated serviceConnectionPolicy", "serviceConnectionPolicy", updated)
		a.actual = updated
	}

	mapCtx := &direct.MapContext{}
	observedState := NetworkConnectivityServiceConnectionPolicyObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setObservedState(u, observedState)
}

func (a *serviceConnectionPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("serviceConnectionPolicy %q not found", a.fullyQualifiedName())
	}

	mc := &direct.MapContext{}
	spec := NetworkConnectivityServiceConnectionPolicySpec_FromProto(mc, a.actual)
	if err := mc.Err(); err != nil {
		return nil, fmt.Errorf("error converting serviceConnectionPolicy from API %w", err)
	}

	spec.ProjectRef.External = a.projectID

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting serviceConnectionPolicy spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: specObj,
	}
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.NetworkConnectivityServiceConnectionPolicyGVK)
	return u, nil
}
