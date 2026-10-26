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

// +generated:controller
// krm.group: gkemulticloud.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: GKEMultiCloudAzureNodePool

package gkemulticloud

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkemulticloud/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/gkemulticloud/apiv1"
	gkemulticloudpb "cloud.google.com/go/gkemulticloud/apiv1/gkemulticloudpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.GKEMultiCloudAzureNodePoolGVK, NewAzureNodePoolModel)
}

func NewAzureNodePoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAzureNodePool{config: *config}, nil
}

var _ directbase.Model = &modelAzureNodePool{}

type modelAzureNodePool struct {
	config config.ControllerConfig
}

func (m *modelAzureNodePool) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.GKEMultiCloudAzureNodePool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.GKEMultiCloudAzureNodePoolIdentity)

	// Get gkemulticloud GCP client
	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &AzureNodePoolAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAzureNodePool) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}
func (m *modelAzureNodePool) client(ctx context.Context, location string) (*gcp.AzureClustersClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-gkemulticloud.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewAzureClustersClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AzureNodePool client: %w", err)
	}
	return gcpClient, err
}

type AzureNodePoolAdapter struct {
	id        *krm.GKEMultiCloudAzureNodePoolIdentity
	gcpClient *gcp.AzureClustersClient
	desired   *krm.GKEMultiCloudAzureNodePool
	actual    *gkemulticloudpb.AzureNodePool
}

var _ directbase.Adapter = &AzureNodePoolAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AzureNodePoolAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AzureNodePool", "name", a.id)

	req := &gkemulticloudpb.GetAzureNodePoolRequest{
		Name: a.id.String(),
	}
	azurenodepoolpb, err := a.gcpClient.GetAzureNodePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AzureNodePool %q: %w", a.id, err)
	}

	a.actual = azurenodepoolpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AzureNodePoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AzureNodePool", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEMultiCloudAzureNodePoolSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &gkemulticloudpb.CreateAzureNodePoolRequest{
		Parent:          a.id.ParentString(),
		AzureNodePoolId: a.id.AzureNodePool,
		AzureNodePool:   resource,
	}
	op, err := a.gcpClient.CreateAzureNodePool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AzureNodePool %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AzureNodePool %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created AzureNodePool", "name", a.id)

	status := &krm.GKEMultiCloudAzureNodePoolStatus{}
	status.ObservedState = GKEMultiCloudAzureNodePoolObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AzureNodePoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AzureNodePool", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		GKEMultiCloudAzureNodePoolSpec_FromProto,
		GKEMultiCloudAzureNodePoolSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &gkemulticloudpb.UpdateAzureNodePoolRequest{
			UpdateMask:    updateMask,
			AzureNodePool: resource,
		}
		op, err := a.gcpClient.UpdateAzureNodePool(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AzureNodePool %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("AzureNodePool %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated AzureNodePool", "name", a.id)
	}

	status := &krm.GKEMultiCloudAzureNodePoolStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = GKEMultiCloudAzureNodePoolObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AzureNodePoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEMultiCloudAzureNodePool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEMultiCloudAzureNodePoolSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.GKEMultiCloudAzureNodePoolGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AzureNodePoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AzureNodePool", "name", a.id)

	req := &gkemulticloudpb.DeleteAzureNodePoolRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteAzureNodePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent AzureNodePool, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting AzureNodePool %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AzureNodePool", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete AzureNodePool %s: %w", a.id, err)
	}
	return true, nil
}
