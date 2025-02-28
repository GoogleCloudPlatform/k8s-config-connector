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
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/monitoring"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	api "google.golang.org/api/networkconnectivity/v1"
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

	gcpClient *api.Service
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &serviceConnectionPolicyAdapter{}

// AdapterForObject implements the Model interface.
func (m *serviceConnectionPolicyModel) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	clientBuilder, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	gcpClient, err := clientBuilder.newNetworkConnectivityClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.NetworkConnectivityServiceConnectionPolicy{}
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

	location := direct.ValueOf(obj.Spec.Location)
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectRef, err := refs.ResolveProject(ctx, kube, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	if err := common.VisitFields(obj, &refNormalizer{ctx: ctx, src: obj, project: *projectRef, kube: kube}); err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkConnectivityServiceConnectionPolicySpec_ToProto(mapCtx, &obj.Spec)
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
	// TODO: What is format?
	// // Format: //monitoring.googleapis.com/projects/PROJECT_NUMBER/dashboards/DASHBOARD_ID
	// if !strings.HasPrefix(url, "//monitoring.googleapis.com/") {
	// 	return nil, nil
	// }

	// tokens := strings.Split(strings.TrimPrefix(url, "//apigee.googleapis.com/"), "/")
	// if len(tokens) == 2 && tokens[0] == "organizations" {
	// 	gcpClient, err := newGCPClient(ctx, m.config)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("building gcp client: %w", err)
	// 	}

	// 	apigeeClient, err := gcpClient.newApigeeClient(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	return &serviceConnectionPolicyAdapter{
	// 		resourceID:   tokens[1],
	// 		apigeeClient: apigeeClient,
	// 	}, nil
	// }

	return nil, nil

}

// Find implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	fqn := a.fullyQualifiedName()
	actual, err := a.gcpClient.Projects.Locations.ServiceConnectionPolicies.Get(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	if err := convertAPIToProto(actual, &a.actual); err != nil {
		return false, err
	}

	return true, nil
}

func (a *serviceConnectionPolicyAdapter) waitForOperation(ctx context.Context, op *api.GoogleLongrunningOperation) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		latest, err := a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation %q: %w", op.Name, err)
		}

		if latest.Done {
			return nil
		}

		time.Sleep(2 * time.Second)
	}
}

// Delete implements the Adapter interface.
func (a *serviceConnectionPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// TODO: Delete via status selfLink?

	fqn := a.fullyQualifiedName()

	op, err := a.gcpClient.Projects.Locations.ServiceConnectionPolicies.Delete(fqn).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("deleting serviceConnectionPolicy %q: %w", fqn, err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
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

	req := &api.ServiceConnectionPolicy{}
	if err := convertProtoToAPI(a.desired, req); err != nil {
		return err
	}

	log.V(0).Info("creating serviceConnectionPolicy", "req", req)
	op, err := a.gcpClient.Projects.Locations.ServiceConnectionPolicies.Create(a.parent(), req).ServiceConnectionPolicyId(a.resourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating serviceConnectionPolicy: %w", err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for create of serviceConnectionPolicy %q: %w", fqn, err)
	}

	created, err := a.gcpClient.Projects.Locations.ServiceConnectionPolicies.Get(fqn).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created serviceConnectionPolicy %q: %w", fqn, err)
	}
	log.V(2).Info("created organization", "serviceConnectionPolicy", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	var createdPB *pb.ServiceConnectionPolicy
	if err := convertAPIToProto(created, &createdPB); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	observedState := NetworkConnectivityServiceConnectionPolicyObservedState_FromProto(mapCtx, createdPB)
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

	// TODO: Where/how do we want to enforce immutability?

	fqn := a.fullyQualifiedName()

	if monitoring.ShouldReconcileBasedOnEtag(ctx, u, a.actual.Etag) {
		req := &api.ServiceConnectionPolicy{}
		if err := convertProtoToAPI(a.desired, req); err != nil {
			return err
		}

		log.V(2).Info("updating serviceConnectionPolicy", "request", req)
		op, err := a.gcpClient.Projects.Locations.ServiceConnectionPolicies.Patch(fqn, req).Context(ctx).Do()
		if err != nil {
			return err
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for update of serviceConnectionPolicy %q: %w", fqn, err)
		}

		// TODO: Other calls

		updated, err := a.gcpClient.Projects.Locations.ServiceConnectionPolicies.Get(fqn).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated serviceConnectionPolicy %q: %w", fqn, err)
		}

		log.V(2).Info("updated serviceConnectionPolicy", "serviceConnectionPolicy", updated)
		if err := convertAPIToProto(updated, &a.actual); err != nil {
			return err
		}
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
		Object: make(map[string]interface{}),
	}
	u.SetName(a.resourceID)
	u.SetGroupVersionKind(krm.NetworkConnectivityServiceConnectionPolicyGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *serviceConnectionPolicyAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/serviceConnectionPolicies/%s", a.projectID, a.location, a.resourceID)
}

func (a *serviceConnectionPolicyAdapter) parent() string {
	return fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location)
}
