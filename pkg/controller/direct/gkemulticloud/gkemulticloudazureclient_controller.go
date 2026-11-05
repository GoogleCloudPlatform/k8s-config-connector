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
// krm.kind: GKEMultiCloudAzureClient

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
	registry.RegisterModel(krm.GKEMultiCloudAzureClientGVK, NewAzureClientModel)
}

func NewAzureClientModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAzureClient{config: *config}, nil
}

var _ directbase.Model = &modelAzureClient{}

type modelAzureClient struct {
	config config.ControllerConfig
}

func (m *modelAzureClient) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.GKEMultiCloudAzureClient{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.GKEMultiCloudAzureClientIdentity)

	// Get gkemulticloud GCP client
	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &AzureClientAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAzureClient) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}
func (m *modelAzureClient) client(ctx context.Context, location string) (*gcp.AzureClustersClient, error) {
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
		return nil, fmt.Errorf("building AzureClient client: %w", err)
	}
	return gcpClient, err
}

type AzureClientAdapter struct {
	id        *krm.GKEMultiCloudAzureClientIdentity
	gcpClient *gcp.AzureClustersClient
	desired   *krm.GKEMultiCloudAzureClient
	actual    *gkemulticloudpb.AzureClient
}

var _ directbase.Adapter = &AzureClientAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AzureClientAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AzureClient", "name", a.id)

	req := &gkemulticloudpb.GetAzureClientRequest{
		Name: a.id.String(),
	}
	azureclientpb, err := a.gcpClient.GetAzureClient(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AzureClient %q: %w", a.id, err)
	}

	a.actual = azureclientpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AzureClientAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AzureClient", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEMultiCloudAzureClientSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &gkemulticloudpb.CreateAzureClientRequest{
		Parent:        a.id.ParentString(),
		AzureClientId: a.id.AzureClient,
		AzureClient:   resource,
	}
	op, err := a.gcpClient.CreateAzureClient(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AzureClient %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AzureClient %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created AzureClient", "name", a.id)

	status := &krm.GKEMultiCloudAzureClientStatus{}
	status.ObservedState = GKEMultiCloudAzureClientObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AzureClientAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AzureClient", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		GKEMultiCloudAzureClientSpec_FromProto,
		GKEMultiCloudAzureClientSpec_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating GKEMultiCloudAzureClient is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.GKEMultiCloudAzureClientStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = GKEMultiCloudAzureClientObservedState_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AzureClientAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEMultiCloudAzureClient{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEMultiCloudAzureClientSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.GKEMultiCloudAzureClientGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AzureClientAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AzureClient", "name", a.id)

	req := &gkemulticloudpb.DeleteAzureClientRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteAzureClient(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent AzureClient, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting AzureClient %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AzureClient", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete AzureClient %s: %w", a.id, err)
	}
	return true, nil
}
