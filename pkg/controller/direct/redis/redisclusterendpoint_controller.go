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

package redis

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcpcompute "cloud.google.com/go/compute/apiv1"
	gcp "cloud.google.com/go/redis/cluster/apiv1"

	computepb "cloud.google.com/go/compute/apiv1/computepb"
	pb "cloud.google.com/go/redis/cluster/apiv1/clusterpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.RedisClusterEndpointGVK, NewEndpointModel)
}

func NewEndpointModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelEndpoint{config: *config}, nil
}

var _ directbase.Model = &modelEndpoint{}

type modelEndpoint struct {
	config config.ControllerConfig
}

func (m *modelEndpoint) client(ctx context.Context) (*gcp.CloudRedisClusterClient, *gcpcompute.ForwardingRulesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, nil, err
	}
	gcpClient, err := gcp.NewCloudRedisClusterRESTClient(ctx, opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("building Redis Cluster client: %w", err)
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
	obj := &krm.RedisClusterEndpoint{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if obj.Spec.ClusterRef == nil {
		return nil, fmt.Errorf("spec.clusterRef is required")
	}
	err := obj.Spec.ClusterRef.Normalize(ctx, reader, obj.GetNamespace())
	if err != nil {
		return nil, err
	}

	gcpClient, cmpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &EndpointAdapter{
		id:        obj.Spec.ClusterRef.GetExternal(),
		gcpClient: gcpClient,
		cmpClient: cmpClient,
		reader:    reader,
		desired:   obj,
	}, nil
}

func resolveEndpointReferences(ctx context.Context, reader client.Reader, obj *krm.RedisClusterEndpoint) error {
	for _, endpoint := range obj.Spec.ClusterEndpoints {
		for _, connection := range endpoint.Connections {
			if connection.PSCConnection != nil {
				pscConnection := connection.PSCConnection
				if err := pscConnection.ForwardingRuleRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (m *modelEndpoint) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// no support for URLs - RedisClusterEndpoint is only a proxy object to RedisCluster.
	return nil, nil
}

type EndpointAdapter struct {
	id        string
	gcpClient *gcp.CloudRedisClusterClient
	cmpClient *gcpcompute.ForwardingRulesClient
	reader    client.Reader
	desired   *krm.RedisClusterEndpoint
	actual    *pb.Cluster
}

var _ directbase.Adapter = &EndpointAdapter{}

// Find retrieves the GCP resource.
func (a *EndpointAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Endpoint", "name", a.id)

	req := &pb.GetClusterRequest{Name: a.id}
	clusterpb, err := a.gcpClient.GetCluster(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting RedisCluster %q: %w", a.id, err)
	}

	a.actual = clusterpb
	for _, endpoint := range a.actual.ClusterEndpoints {
		for _, conn := range endpoint.Connections {
			if pscConnection := conn.GetPscConnection(); pscConnection != nil {
				return true, nil
			}
		}
	}
	return false, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EndpointAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating user created cluster endpoints", "name", a.id)

	if err := resolveEndpointReferences(ctx, a.reader, a.desired); err != nil {
		return err
	}

	actual, err := a.updateConnections(ctx, a.desired.Spec.ClusterEndpoints, createOp.GetUnstructured())
	if err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	status := &krm.RedisClusterEndpointStatus{}
	status.ObservedState = RedisClusterEndpointObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *EndpointAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating user created cluster endpoints", "name", a.id)

	if err := resolveEndpointReferences(ctx, a.reader, a.desired); err != nil {
		return err
	}

	actual, err := a.updateConnections(ctx, a.desired.Spec.ClusterEndpoints, updateOp.GetUnstructured())
	if err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	status := &krm.RedisClusterEndpointStatus{}
	status.ObservedState = RedisClusterEndpointObservedState_FromProto(mapCtx, actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *EndpointAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	// no support for export - RedisClusterEndpoint is only a proxy object to RedisCluster.
	return nil, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *EndpointAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting user created cluster endpoints", "name", a.id)

	if _, err := a.updateConnections(ctx, nil, deleteOp.GetUnstructured()); err != nil {
		return false, err
	}
	return true, nil
}

func (a *EndpointAdapter) updateConnections(ctx context.Context, userCreated []krm.ClusterEndpoint_ClusterEndpoint, obj *unstructured.Unstructured) (*pb.Cluster, error) {
	mapCtx := &direct.MapContext{}
	oldFRs := make(map[string]struct{})
	for _, endpoint := range a.actual.ClusterEndpoints {
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
			if pscConnection := conn.PSCConnection; pscConnection != nil {
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

	var endpoints []*pb.ClusterEndpoint
	for _, uc := range userCreated {
		endpoints = append(endpoints, ClusterEndpoint_ClusterEndpoint_ToProto(mapCtx, &uc))
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
				pscConnection.ProjectId = tokens[1]
				pscConnection.Address = rsp.GetIPAddress()
				pscConnection.PscConnectionId = fmt.Sprintf("%d", rsp.GetPscConnectionId())
				pscConnection.Network = removePrefixFromExternal(mapCtx, "network", rsp.GetNetwork())
				pscConnection.ServiceAttachment = removePrefixFromExternal(mapCtx, "target", rsp.GetTarget())
			}
		}
	}

	report := &structuredreporting.Diff{Object: obj}
	report.AddField("clusterEndpoints", a.actual.ClusterEndpoints, endpoints)

	updateReq := &pb.UpdateClusterRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"cluster_endpoints"},
		},
		Cluster: &pb.Cluster{
			Name:             a.actual.Name,
			ClusterEndpoints: endpoints,
		},
	}
	op, err := a.gcpClient.UpdateCluster(ctx, updateReq)
	if err != nil {
		return nil, fmt.Errorf("updating redisCluster %s: %w", a.id, err)
	}
	actual, err := op.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("redisCluster %s waiting update: %w", a.id, err)
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
