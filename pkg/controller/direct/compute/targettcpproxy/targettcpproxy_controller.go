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

package targettcpproxy

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	"google.golang.org/api/option"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	kccpredicate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/predicate"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ctrlName = "firewallpolicyrule-controller"

func init() {
	rg := &TargetTCPProxyReconcileGate{}
	registry.RegisterModelWithReconcileGate(krm.ComputeTargetTCPProxyGVK, NewTargetTCPProxyModel, rg)
}

func NewTargetTCPProxyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &targetTCPProxyModel{config: config}, nil
}

type TargetTCPProxyReconcileGate struct {
	optIn kccpredicate.OptInToDirectReconciliation
}

var _ kccpredicate.ReconcileGate = &TargetTCPProxyReconcileGate{}

func (r *TargetTCPProxyReconcileGate) ShouldReconcile(o *unstructured.Unstructured) bool {
	if r.optIn.ShouldReconcile(o) {
		return true
	}
	obj := &krm.ComputeTargetTCPProxy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(o.Object, &obj); err != nil {
		return false
	}
	// Run the direct reconciler only when spec.location is specified and not global
	return obj.Spec.Location != nil && obj.Spec.Location != direct.PtrTo("global")
}

type targetTCPProxyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &targetTCPProxyModel{}

type targetTCPProxyAdapter struct {
	id                             *krm.ComputeTargetTCPProxyRef
	targetTcpProxiesClient         *gcp.TargetTcpProxiesClient
	regionalTargetTcpProxiesClient *gcp.RegionTargetTcpProxiesClient
	desired                        *krm.ComputeTargetTCPProxy
	actual                         *computepb.TargetTcpProxy
	reader                         client.Reader
}

var _ directbase.Adapter = &targetTCPProxyAdapter{}

func (m *targetTCPProxyModel) client(ctx context.Context) (*gcp.TargetTcpProxiesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewTargetTcpProxiesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TargetTcpProxy client: %w", err)
	}
	return gcpClient, err
}

func (m *targetTCPProxyModel) regionalClient(ctx context.Context) (*gcp.RegionTargetTcpProxiesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRegionTargetTcpProxiesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TargetTcpProxy client: %w", err)
	}
	return gcpClient, err
}

func (m *targetTCPProxyModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeTargetTCPProxy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	computeTargetTCPProxyRef, err := krm.NewComputeTargetTCPProxyRef(ctx, reader, obj, u)
	if err != nil {
		return nil, err
	}

	targetTCPProxyAdapter := &targetTCPProxyAdapter{
		id:      computeTargetTCPProxyRef,
		desired: obj,
		reader:  reader,
	}

	// Get location
	parent, err := computeTargetTCPProxyRef.Parent()
	if err != nil {
		return nil, fmt.Errorf("get ComputeTargetTCPProxyAdapter parent %s: %w", computeTargetTCPProxyRef.External, err)
	}
	location := parent.Location

	// Handle API/TF default values
	if obj.Spec.ProxyBind != nil && *obj.Spec.ProxyBind == false {
		obj.Spec.ProxyBind = nil
	}
	if obj.Spec.ProxyHeader == nil {
		obj.Spec.ProxyHeader = direct.PtrTo("NONE")
	}

	// Get GCP client
	if location == "global" {
		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, fmt.Errorf("building gcp client: %w", err)
		}
		targetTCPProxyAdapter.targetTcpProxiesClient = gcpClient
	} else {
		gcpClient, err := m.regionalClient(ctx)
		if err != nil {
			return nil, fmt.Errorf("building gcp client: %w", err)
		}
		targetTCPProxyAdapter.regionalTargetTcpProxiesClient = gcpClient
	}
	return targetTCPProxyAdapter, nil
}

func (m *targetTCPProxyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *targetTCPProxyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeTargetTCPProxy", "name", a.id.External)

	targetTCPProxy, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeTargetTCPProxy %q: %w", a.id.External, err)
	}
	a.actual = targetTCPProxy
	return true, nil
}

func (a *targetTCPProxyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	// Resolve dependencies
	err := resolveBackendService(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeTargetTCPProxy", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	targetTCPProxy := ComputeTargetTCPProxySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.id.Parent()
	if err != nil {
		return fmt.Errorf("get ComputeTargetTCPProxy parent %s: %w", a.id.External, err)
	}
	location := parent.Location

	tokens := strings.Split(a.id.External, "/")
	targetTCPProxy.Name = direct.LazyPtr(tokens[len(tokens)-1])

	op := &gcp.Operation{}
	if location == "global" {
		req := &computepb.InsertTargetTcpProxyRequest{
			Project:                parent.ProjectID,
			TargetTcpProxyResource: targetTCPProxy,
		}
		op, err = a.targetTcpProxiesClient.Insert(ctx, req)
	} else {
		req := &computepb.InsertRegionTargetTcpProxyRequest{
			Project:                parent.ProjectID,
			Region:                 location,
			TargetTcpProxyResource: targetTCPProxy,
		}
		op, err = a.regionalTargetTcpProxiesClient.Insert(ctx, req)
	}

	if err != nil {
		return fmt.Errorf("creating ComputeTargetTCPProxy %s: %w", a.id.External, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeTargetTCPProxy %s create failed: %w", a.id.External, err)
		}
	}
	log.V(2).Info("successfully created ComputeTargetTCPProxy", "name", a.id.External)

	// Get the created resource
	created := &computepb.TargetTcpProxy{}
	created, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeTargetTCPProxy %s: %w", a.id.External, err)
	}

	status := &krm.ComputeTargetTCPProxyStatus{}
	status = ComputeTargetTCPProxyStatus_FromProto(mapCtx, created)

	parent, err = a.id.Parent()
	if err != nil {
		return err
	}

	externalRef := parent.String() + "/targetTcpProxies/" + direct.ValueOf(created.Name)
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *targetTCPProxyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// Resolve dependencies
	err := resolveBackendService(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating ComputeTargetTCPProxy", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	targetTCPProxy := ComputeTargetTCPProxySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	op := &gcp.Operation{}
	updated := &computepb.TargetTcpProxy{}

	parent, err := a.id.Parent()
	if err != nil {
		return fmt.Errorf("get ComputeTargetTCPProxy parent %s: %w", a.id.External, err)
	}
	location := parent.Location

	// Regional API does not support Update
	if location != "global" {
		return fmt.Errorf("update operation not supported for resource %v %v",
			a.desired.GroupVersionKind(), k8s.GetNamespacedName(a.desired))
	}

	tokens := strings.Split(a.id.External, "/")
	targetTCPProxy.Name = direct.LazyPtr(tokens[len(tokens)-1])

	if !reflect.DeepEqual(targetTCPProxy.ProxyHeader, a.actual.ProxyHeader) {
		setProxyHeaderReq := &computepb.SetProxyHeaderTargetTcpProxyRequest{
			Project: parent.ProjectID,
			TargetTcpProxiesSetProxyHeaderRequestResource: &computepb.TargetTcpProxiesSetProxyHeaderRequest{ProxyHeader: targetTCPProxy.ProxyHeader},
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		op, err = a.targetTcpProxiesClient.SetProxyHeader(ctx, setProxyHeaderReq)
		if err != nil {
			return fmt.Errorf("updating ComputeTargetTCPProxy proxy header %s: %w", a.id.External, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeTargetTCPProxy proxy header %s update failed: %w", a.id.External, err)
			}
		}
		log.V(2).Info("successfully updated ComputeTargetTCPProxy proxy header", "name", a.id.External)
	}

	if !reflect.DeepEqual(targetTCPProxy.Service, a.actual.Service) {
		setBackendServiceReq := &computepb.SetBackendServiceTargetTcpProxyRequest{
			Project: parent.ProjectID,
			TargetTcpProxiesSetBackendServiceRequestResource: &computepb.TargetTcpProxiesSetBackendServiceRequest{Service: targetTCPProxy.Service},
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		op, err = a.targetTcpProxiesClient.SetBackendService(ctx, setBackendServiceReq)
		if err != nil {
			return fmt.Errorf("updating ComputeTargetTCPProxy backend service %s: %w", a.id.External, err)
		}
		if !op.Done() {
			err = op.Wait(ctx)
			if err != nil {
				return fmt.Errorf("waiting ComputeTargetTCPProxy backend service %s update failed: %w", a.id.External, err)
			}
		}
		log.V(2).Info("successfully updated ComputeTargetTCPProxy backend service", "name", a.id.External)

	}

	// Get the updated resource
	updated, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeTargetTCPProxy %s: %w", a.id.External, err)
	}

	status := &krm.ComputeTargetTCPProxyStatus{}
	status = ComputeTargetTCPProxyStatus_FromProto(mapCtx, updated)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *targetTCPProxyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("targetTcpProxy %s not found", a.id.External)
	}

	mc := &direct.MapContext{}
	spec := ComputeTargetTCPProxySpec_FromProto(mc, a.actual)
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
	log.V(2).Info("deleting ComputeTargetTcpProxy", "name", a.id.External)

	parent, err := a.id.Parent()
	if err != nil {
		return false, fmt.Errorf("get ComputeTargetTcpProxy parent %s: %w", a.id.External, err)
	}
	location := parent.Location

	op := &gcp.Operation{}
	tokens := strings.Split(a.id.External, "/")
	if location == "global" {
		delReq := &computepb.DeleteTargetTcpProxyRequest{
			Project:        parent.ProjectID,
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		op, err = a.targetTcpProxiesClient.Delete(ctx, delReq)
	} else {
		delReq := &computepb.DeleteRegionTargetTcpProxyRequest{
			Project:        parent.ProjectID,
			Region:         location,
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		op, err = a.regionalTargetTcpProxiesClient.Delete(ctx, delReq)
	}

	if err != nil {
		return false, fmt.Errorf("deleting ComputeTargetTcpProxy %s: %w", a.id.External, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeTargetTcpProxy %s delete failed: %w", a.id.External, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeTargetTcpProxy", "name", a.id.External)
	return true, nil
}

func (a *targetTCPProxyAdapter) get(ctx context.Context) (*computepb.TargetTcpProxy, error) {
	parent, err := a.id.Parent()
	if err != nil {
		return nil, fmt.Errorf("get ComputeTargetTcpProxy parent %s: %w", a.id.External, err)
	}
	location := parent.Location

	tokens := strings.Split(a.id.External, "/")
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

// This function get the normalized external values and convert it to the API required format
func resolveBackendService(ctx context.Context, reader client.Reader, obj *krm.ComputeTargetTCPProxy) error {
	// API required format: selfLink
	computeBasePath := "https://www.googleapis.com/compute/v1/"
	ref := obj.Spec.BackendServiceRef
	if ref != nil {
		// Get normalized external
		_, err := ref.NormalizedExternal(ctx, reader, obj.GetNamespace())
		if err != nil {
			return fmt.Errorf("failed to get BackendServiceRef: %w", err)
		}
		// Convert normalized external to API required format
		v := ref.External
		_, err = krm.ParseBackendServiceExternal(v)
		// Value follows KCC external format, likely it's created by direct controller
		if err == nil {
			// add the compute prefix in front
			obj.Spec.BackendServiceRef.External = computeBasePath + v
			return nil
		}
		// For backward compatibility, we also accept values that does not match the KCC external format.(likely it's created by legacy controller)
		// Return the value as is and let the API handle it
		obj.Spec.BackendServiceRef.External = v
		return nil
	}
	return nil
}
