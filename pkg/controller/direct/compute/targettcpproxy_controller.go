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

package compute

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ComputeTargetTCPProxyGVK, NewTargetTCPProxyModel)
}

func NewTargetTCPProxyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &targetTCPProxyModel{config: config}, nil
}

type targetTCPProxyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &targetTCPProxyModel{}

type targetTCPProxyAdapter struct {
	id                             *krm.TargetTCPProxyIdentity
	targetTcpProxiesClient         *gcp.TargetTcpProxiesClient
	regionalTargetTcpProxiesClient *gcp.RegionTargetTcpProxiesClient
	desired                        *krm.ComputeTargetTCPProxy
	actual                         *computepb.TargetTcpProxy
	reader                         client.Reader
}

var _ directbase.Adapter = &targetTCPProxyAdapter{}

func (m *targetTCPProxyModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeTargetTCPProxy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTargetTCPProxyIdentity(ctx, reader, obj, u)
	if err != nil {
		return nil, err
	}

	targetTCPProxyAdapter := &targetTCPProxyAdapter{
		id:      id,
		desired: obj,
		reader:  reader,
	}

	// Get location
	parent := id.Parent()
	location := parent.Location

	// Handle API/TF default values
	if obj.Spec.ProxyBind != nil && *obj.Spec.ProxyBind == false {
		obj.Spec.ProxyBind = nil
	}
	if obj.Spec.ProxyHeader == nil {
		obj.Spec.ProxyHeader = direct.PtrTo("NONE")
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	// Get GCP client
	if location == "global" {
		targetTcpProxiesClient, err := gcpClient.newTargetTcpProxiesClient(ctx)
		if err != nil {
			return nil, err
		}
		targetTCPProxyAdapter.targetTcpProxiesClient = targetTcpProxiesClient
	} else {
		regionalTargetTcpProxiesClient, err := gcpClient.newRegionalTargetTcpProxiesClient(ctx)
		if err != nil {
			return nil, err
		}
		targetTCPProxyAdapter.regionalTargetTcpProxiesClient = regionalTargetTcpProxiesClient
	}
	return targetTCPProxyAdapter, nil
}

func (m *targetTCPProxyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *targetTCPProxyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeTargetTCPProxy", "name", a.id)

	targetTCPProxy, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeTargetTCPProxy %q: %w", a.id, err)
	}
	a.actual = targetTCPProxy
	return true, nil
}

func (a *targetTCPProxyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	var err error

	err = resolveTargetTcpProxyRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeTargetTCPProxy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	targetTCPProxy := ComputeTargetTCPProxySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.Parent()

	tokens := strings.Split(a.id.String(), "/")
	targetTCPProxy.Name = direct.LazyPtr(tokens[len(tokens)-1])

	op := &gcp.Operation{}
	if parent.Location == "global" {
		req := &computepb.InsertTargetTcpProxyRequest{
			Project:                parent.ProjectID,
			TargetTcpProxyResource: targetTCPProxy,
		}
		op, err = a.targetTcpProxiesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertRegionTargetTcpProxyRequest{
			Project:                parent.ProjectID,
			Region:                 parent.Location,
			TargetTcpProxyResource: targetTCPProxy,
		}
		op, err = a.regionalTargetTcpProxiesClient.Insert(ctx, req)
	}

	if err != nil {
		return fmt.Errorf("creating ComputeTargetTCPProxy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeTargetTCPProxy %s create failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputeTargetTCPProxy", "name", a.id)

	// Get the created resource
	created := &computepb.TargetTcpProxy{}
	created, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeTargetTCPProxy %s: %w", a.id, err)
	}

	status := &krm.ComputeTargetTCPProxyStatus{}
	status = ComputeTargetTCPProxyStatus_FromProto(mapCtx, created)

	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *targetTCPProxyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	var err error

	err = resolveTargetTcpProxyRefs(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeTargetTCPProxy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	targetTCPProxy := ComputeTargetTCPProxySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.Parent()
	tokens := strings.Split(a.id.String(), "/")
	updated := &computepb.TargetTcpProxy{}

	// Assign API output-only values
	// todo: https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/4455
	targetTCPProxy.CreationTimestamp = a.actual.CreationTimestamp
	targetTCPProxy.Id = a.actual.Id
	targetTCPProxy.SelfLink = a.actual.SelfLink
	targetTCPProxy.Kind = a.actual.Kind
	// Convert region `europe-west4` to proto region format `https://www.googleapis.com/compute/v1/projects/projectId/regions/europe-west4`
	// Prevent diff when comparing with proto message
	targetTCPProxy.Region = direct.LazyPtr(fmt.Sprintf("https://www.googleapis.com/compute/v1/%s", parent))
	targetTCPProxy.Name = direct.LazyPtr(a.id.ID())
	paths, err := common.CompareProtoMessage(targetTCPProxy, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())

		// Even though there is no update, we still want to update KRM status
		updated = a.actual
	} else {
		// Changes on resource spec are detected
		if parent.Location != "global" {
			// Regional ComputeTargetTCPProxy API does not support Update
			return fmt.Errorf("update operation not supported for regional ComputeTargetTCPProxy")
		}

		op := &gcp.Operation{}
		if !reflect.DeepEqual(targetTCPProxy.ProxyHeader, a.actual.ProxyHeader) {
			setProxyHeaderReq := &computepb.SetProxyHeaderTargetTcpProxyRequest{
				Project: parent.ProjectID,
				TargetTcpProxiesSetProxyHeaderRequestResource: &computepb.TargetTcpProxiesSetProxyHeaderRequest{ProxyHeader: targetTCPProxy.ProxyHeader},
				TargetTcpProxy: tokens[len(tokens)-1],
			}
			op, err = a.targetTcpProxiesClient.SetProxyHeader(ctx, setProxyHeaderReq)
			if err != nil {
				return fmt.Errorf("updating ComputeTargetTCPProxy proxy header %s: %w", a.id, err)
			}
			if !op.Done() {
				err = op.Wait(ctx)
				if err != nil {
					return fmt.Errorf("waiting ComputeTargetTCPProxy proxy header %s update failed: %w", a.id, err)
				}
			}
			log.V(2).Info("successfully updated ComputeTargetTCPProxy proxy header", "name", a.id)
		}

		if !reflect.DeepEqual(targetTCPProxy.Service, a.actual.Service) {
			setBackendServiceReq := &computepb.SetBackendServiceTargetTcpProxyRequest{
				Project: parent.ProjectID,
				TargetTcpProxiesSetBackendServiceRequestResource: &computepb.TargetTcpProxiesSetBackendServiceRequest{Service: targetTCPProxy.Service},
				TargetTcpProxy: tokens[len(tokens)-1],
			}
			op, err = a.targetTcpProxiesClient.SetBackendService(ctx, setBackendServiceReq)
			if err != nil {
				return fmt.Errorf("updating ComputeTargetTCPProxy backend service %s: %w", a.id, err)
			}
			if !op.Done() {
				err = op.Wait(ctx)
				if err != nil {
					return fmt.Errorf("waiting ComputeTargetTCPProxy backend service %s update failed: %w", a.id, err)
				}
			}
			log.V(2).Info("successfully updated ComputeTargetTCPProxy backend service", "name", a.id)
		}

		// Get the updated resource
		updated, err = a.get(ctx)
		if err != nil {
			return fmt.Errorf("getting ComputeTargetTCPProxy %s: %w", a.id, err)
		}
	}

	status := &krm.ComputeTargetTCPProxyStatus{}
	status = ComputeTargetTCPProxyStatus_FromProto(mapCtx, updated)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *targetTCPProxyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("targetTcpProxy %s not found", a.id)
	}

	mc := &direct.MapContext{}
	spec := ComputeTargetTCPProxySpec_FromProto(mc, a.actual)
	// Convert proto region format `https://www.googleapis.com/compute/v1/projects/projectId/regions/europe-west4` to `europe-west4`
	if spec.Location != nil {
		region := strings.Split(*spec.Location, "/")
		spec.Location = direct.LazyPtr(region[len(region)-1])
	}
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting targetTcpProxy spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetGroupVersionKind(krm.ComputeTargetTCPProxyGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *targetTCPProxyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeTargetTcpProxy", "name", a.id)

	parent := a.id.Parent()

	var err error
	op := &gcp.Operation{}
	tokens := strings.Split(a.id.String(), "/")
	if parent.Location == "global" {
		delReq := &computepb.DeleteTargetTcpProxyRequest{
			Project:        parent.ProjectID,
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		op, err = a.targetTcpProxiesClient.Delete(ctx, delReq)
	} else {
		delReq := &computepb.DeleteRegionTargetTcpProxyRequest{
			Project:        parent.ProjectID,
			Region:         parent.Location,
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		op, err = a.regionalTargetTcpProxiesClient.Delete(ctx, delReq)
	}

	if err != nil {
		return false, fmt.Errorf("deleting ComputeTargetTcpProxy %s: %w", a.id, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeTargetTcpProxy %s delete failed: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeTargetTcpProxy", "name", a.id)
	return true, nil
}

func (a *targetTCPProxyAdapter) get(ctx context.Context) (*computepb.TargetTcpProxy, error) {
	parent := a.id.Parent()
	location := parent.Location

	tokens := strings.Split(a.id.String(), "/")
	if location == "global" {
		getReq := &computepb.GetTargetTcpProxyRequest{
			Project:        parent.ProjectID,
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		return a.targetTcpProxiesClient.Get(ctx, getReq)
	} else {
		getReq := &computepb.GetRegionTargetTcpProxyRequest{
			Project:        parent.ProjectID,
			Region:         location,
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		return a.regionalTargetTcpProxiesClient.Get(ctx, getReq)
	}
}
