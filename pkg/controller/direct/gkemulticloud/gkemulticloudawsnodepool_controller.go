// Copyright 2025 Google LLC
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
// krm.kind: GKEMultiCloudAWSNodePool

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
	registry.RegisterModel(krm.GKEMultiCloudAWSNodePoolGVK, NewAwsNodePoolModel)
}

func NewAwsNodePoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAwsNodePool{config: *config}, nil
}

var _ directbase.Model = &modelAwsNodePool{}

type modelAwsNodePool struct {
	config config.ControllerConfig
}

func (m *modelAwsNodePool) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.GKEMultiCloudAWSNodePool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.GKEMultiCloudAWSNodePoolIdentity)

	// Get gkemulticloud GCP client
	gcpClient, err := m.client(ctx, direct.ValueOf(obj.Spec.Location))
	if err != nil {
		return nil, err
	}
	return &AwsNodePoolAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAwsNodePool) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}
func (m *modelAwsNodePool) client(ctx context.Context, location string) (*gcp.AwsClustersClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	if location != "" && location != "global" {
		opts = append(opts, option.WithEndpoint(fmt.Sprintf("%s-gkemulticloud.googleapis.com:443", location)))
	}
	gcpClient, err := gcp.NewAwsClustersClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AwsNodePool client: %w", err)
	}
	return gcpClient, err
}

type AwsNodePoolAdapter struct {
	id        *krm.GKEMultiCloudAWSNodePoolIdentity
	gcpClient *gcp.AwsClustersClient
	desired   *krm.GKEMultiCloudAWSNodePool
	actual    *gkemulticloudpb.AwsNodePool
}

var _ directbase.Adapter = &AwsNodePoolAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AwsNodePoolAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AwsNodePool", "name", a.id)

	req := &gkemulticloudpb.GetAwsNodePoolRequest{
		Name: a.id.String(),
	}
	awsnodepoolpb, err := a.gcpClient.GetAwsNodePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AwsNodePool %q: %w", a.id, err)
	}

	a.actual = awsnodepoolpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AwsNodePoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AwsNodePool", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := GKEMultiCloudAWSNodePoolSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &gkemulticloudpb.CreateAwsNodePoolRequest{
		Parent:        a.id.ParentString(),
		AwsNodePoolId: a.id.AwsNodePool,
		AwsNodePool:   resource,
	}
	op, err := a.gcpClient.CreateAwsNodePool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AwsNodePool %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AwsNodePool %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created AwsNodePool", "name", a.id)

	status := &krm.GKEMultiCloudAWSNodePoolStatus{}
	status.ObservedState = GKEMultiCloudAWSNodePoolObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AwsNodePoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AwsNodePool", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		GKEMultiCloudAWSNodePoolSpec_FromProto,
		GKEMultiCloudAWSNodePoolSpec_ToProto,
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
		req := &gkemulticloudpb.UpdateAwsNodePoolRequest{
			UpdateMask:  updateMask,
			AwsNodePool: resource,
		}
		op, err := a.gcpClient.UpdateAwsNodePool(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AwsNodePool %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("AwsNodePool %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated AwsNodePool", "name", a.id)
	}

	status := &krm.GKEMultiCloudAWSNodePoolStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = GKEMultiCloudAWSNodePoolObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AwsNodePoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.GKEMultiCloudAWSNodePool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(GKEMultiCloudAWSNodePoolSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.GKEMultiCloudAWSNodePoolGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AwsNodePoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AwsNodePool", "name", a.id)

	req := &gkemulticloudpb.DeleteAwsNodePoolRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteAwsNodePool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent AwsNodePool, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting AwsNodePool %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AwsNodePool", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete AwsNodePool %s: %w", a.id, err)
	}
	return true, nil
}
