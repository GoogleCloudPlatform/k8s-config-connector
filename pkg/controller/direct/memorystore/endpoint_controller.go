// Copyright 2026 Google LLC
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

package memorystore

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/memorystore/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcpcompute "cloud.google.com/go/compute/apiv1"
	gcp "cloud.google.com/go/memorystore/apiv1"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	pb "cloud.google.com/go/memorystore/apiv1/memorystorepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.MemorystoreInstanceEndpointGVK, NewEndpointModel)
}

func NewEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEndpoint{config: *config}, nil
}

var _ directbase.Model = &modelEndpoint{}

type modelEndpoint struct {
	config config.ControllerConfig
}

func (m *modelEndpoint) client(ctx context.Context) (*gcp.Client, *gcpcompute.ForwardingRulesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("building InstanceEndpoint client: %w", err)
	}
	cmpClient, err := gcpcompute.NewForwardingRulesRESTClient(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("building ForwardingRules client: %w", err)
	}
	return gcpClient, cmpClient, err
}

func (m *modelEndpoint) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.MemorystoreInstanceEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewEndpointIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := resolveEndpointReferences(ctx, reader, obj); err != nil {
		return nil, err
	}

	// Get memorystore and compute GCP client
	gcpClient, cmpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EndpointAdapter{
		id:        id,
		gcpClient: gcpClient,
		cmpClient: cmpClient,
		desired:   obj,
	}, nil
}

func resolveEndpointReferences(ctx context.Context, reader client.Reader, obj *krm.MemorystoreInstanceEndpoint) error {
	for _, endpoint := range obj.Spec.Endpoints {
		for _, connection := range endpoint.Connections {
			if connection.PscConnection != nil {
				autoConnection := connection.PscConnection
				if err := autoConnection.ForwardingRuleRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *modelEndpoint) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type EndpointAdapter struct {
	id        *krm.EndpointIdentity
	gcpClient *gcp.Client
	cmpClient *gcpcompute.ForwardingRulesClient
	desired   *krm.MemorystoreInstanceEndpoint
	actual    *pb.Instance
}

var _ directbase.Adapter = &EndpointAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *EndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Endpoint", "name", a.id)

	req := &pb.GetInstanceRequest{Name: a.id.String()}
	instancepb, err := a.gcpClient.GetInstance(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Instance %q: %w", a.id, err)
	}

	a.actual = instancepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating user created instance endoints", "name", a.id)

	actual, err := a.updateConnections(ctx, a.desired.Spec.Endpoints, createOp.GetUnstructured())
	if err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	status := &krm.MemorystoreInstanceEndpointStatus{}
	status.ObservedState = MemorystoreInstanceEndpointObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	a.actual = actual
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating user created instance endoints", "name", a.id)

	actual, err := a.updateConnections(ctx, a.desired.Spec.Endpoints, updateOp.GetUnstructured())
	if err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	status := &krm.MemorystoreInstanceEndpointStatus{}
	status.ObservedState = MemorystoreInstanceEndpointObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	a.actual = actual
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.MemorystoreInstanceEndpoint{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(MemorystoreInstanceEndpointSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.InstanceRef = &refs.MemorystoreInstanceRef{External: a.id.Parent()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.MemorystoreInstanceEndpointGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting user created instance endoints", "name", a.id)

	if _, err := a.updateConnections(ctx, nil, deleteOp.GetUnstructured()); err != nil {
		return false, err
	}
	return true, nil
}

func (a *EndpointAdapter) updateConnections(ctx context.Context, userCreated []krm.Endpoint, obj *unstructured.Unstructured) (*pb.Instance, error) {
	mapCtx := &direct.MapContext{}
	oldFRs := make(map[string]struct{})
	for _, endpoint := range a.actual.Endpoints {
		for _, conn := range endpoint.Connections {
			if pscConnection := conn.GetPscConnection(); pscConnection != nil {
				oldFRs[removePrefixFromExternal(mapCtx, "forwardingrule", pscConnection.ForwardingRule)] = struct{}{}
			}
		}
	}
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	newFRs := make(map[string]struct{})
	for _, endpoint := range userCreated {
		for _, conn := range endpoint.Connections {
			if pscConnection := conn.PscConnection; pscConnection != nil {
				newFRs[removePrefixFromExternal(mapCtx, "forwardingrule", pscConnection.ForwardingRuleRef.External)] = struct{}{}
			}
		}
	}
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	if reflect.DeepEqual(oldFRs, newFRs) {
		return a.actual, nil
	}

	var endpoints []*pb.Instance_InstanceEndpoint
	if len(a.actual.Endpoints) > 0 && a.actual.Endpoints[0] != nil {
		if a.actual.Endpoints[0].Connections[0].GetPscAutoConnection() != nil {
			endpoints = append(endpoints, a.actual.Endpoints[0])
		}
	}
	for _, uc := range userCreated {
		endpoints = append(endpoints, Endpoint_ToProto(mapCtx, &uc))
	}
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	for _, endpoint := range endpoints {
		for _, conn := range endpoint.GetConnections() {
			if pscConnection := conn.GetPscConnection(); pscConnection != nil {
				tokens := strings.Split(removePrefixFromExternal(mapCtx, "forwardingrule", pscConnection.ForwardingRule), "/")
				req := &computepb.GetForwardingRuleRequest{
					Project:        tokens[1],
					Region:         tokens[3],
					ForwardingRule: tokens[5],
				}
				rsp, err := a.cmpClient.Get(ctx, req)
				if err != nil {
					return nil, err
				}
				pscConnection.IpAddress = rsp.GetIPAddress()
				pscConnection.PscConnectionId = fmt.Sprintf("%d", rsp.GetPscConnectionId())
				pscConnection.Network = removePrefixFromExternal(mapCtx, "network", rsp.GetNetwork())
				pscConnection.ServiceAttachment = removePrefixFromExternal(mapCtx, "target", rsp.GetTarget())
			}
		}
	}

	report := &structuredreporting.Diff{Object: obj}
	report.AddField("endpoints", a.actual.Endpoints, endpoints)

	updateReq := &pb.UpdateInstanceRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"endpoints"},
		},
		Instance: &pb.Instance{
			Name:      a.actual.Name,
			Endpoints: endpoints,
			// Decouple the lifecycle of the forwarding rules and their endpoints
			// The same strategy is used in the custom delete handler for the Terraform provider.
			// see https://github.com/GoogleCloudPlatform/magic-modules/blob/4304a90191f754b36d5f392da331be609c8b471a/mmv1/templates/terraform/custom_delete/memorystore_instance_desired_user_created_endpoints.go.tmpl#L36
			AsyncInstanceEndpointsDeletionEnabled: direct.LazyPtr(true),
		},
	}
	op, err := a.gcpClient.UpdateInstance(ctx, updateReq)
	if err != nil {
		return nil, fmt.Errorf("updating instance %s: %w", a.id, err)
	}
	actual, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("instance %s waiting update: %w", a.id, err)
	}
	return actual, nil
}

func removePrefixFromExternal(mapCtx *direct.MapContext, fieldName, fieldValue string) string {
	if strings.HasPrefix(fieldValue, "projects/") {
		return fieldValue
	}
	tokens := strings.Split(fieldValue, "/projects/")
	if len(tokens) != 2 {
		mapCtx.Errorf("invalid %s: %s", fieldName, fieldValue)
		return fieldValue
	}
	return "projects/" + tokens[1]
}
