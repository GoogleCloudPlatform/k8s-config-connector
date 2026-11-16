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
// krm.group: configdeployment.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ConfigDeployment

package configdeployment

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdeployment/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/export"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/config/apiv1"
	configpb "cloud.google.com/go/config/apiv1/configpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ConfigDeploymentGVK, NewDeploymentModel)
}

func NewDeploymentModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDeployment{config: *config}, nil
}

var _ directbase.Model = &modelDeployment{}

type modelDeployment struct {
	config config.ControllerConfig
}

func (m *modelDeployment) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConfigDeployment{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ConfigDeploymentIdentity)

	// Get configdeployment GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDeployment) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//config.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ConfigDeploymentIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DeploymentAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDeployment) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Deployment client: %w", err)
	}
	return gcpClient, err
}

type DeploymentAdapter struct {
	id        *krm.ConfigDeploymentIdentity
	gcpClient *gcp.Client
	desired   *krm.ConfigDeployment
	actual    *configpb.Deployment
}

var _ directbase.Adapter = &DeploymentAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DeploymentAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Deployment", "name", a.id)

	req := &configpb.GetDeploymentRequest{
		Name: a.id.String(),
	}
	deploymentpb, err := a.gcpClient.GetDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Deployment %q: %w", a.id, err)
	}

	a.actual = deploymentpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeploymentAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Deployment", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ConfigDeploymentSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = label.GCPLabels(a.desired)

	req := &configpb.CreateDeploymentRequest{
		Parent:       a.id.ParentString(),
		DeploymentId: a.id.Deployment,
		Deployment:   resource,
	}
	op, err := a.gcpClient.CreateDeployment(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Deployment %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Deployment %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Deployment", "name", a.id)

	if fetched, err := a.gcpClient.GetDeployment(ctx, &configpb.GetDeploymentRequest{Name: a.id.String()}); err == nil {
		created = fetched
	}

	status := &krm.ConfigDeploymentStatus{}
	status.ObservedState = ConfigDeploymentObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DeploymentAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Deployment", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpecAndLabels(
		ctx,
		updateOp.GetUnstructured(),
		&a.desired.Spec,
		a.actual,
		"labels",
		ConfigDeploymentSpec_FromProto,
		ConfigDeploymentSpec_ToProto,
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
		req := &configpb.UpdateDeploymentRequest{
			UpdateMask: updateMask,
			Deployment: resource,
		}
		op, err := a.gcpClient.UpdateDeployment(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Deployment %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Deployment %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Deployment", "name", a.id)

		if fetched, err := a.gcpClient.GetDeployment(ctx, &configpb.GetDeploymentRequest{Name: a.id.String()}); err == nil {
			updated = fetched
		}
	}

	status := &krm.ConfigDeploymentStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ConfigDeploymentObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DeploymentAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConfigDeployment{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConfigDeploymentSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ConfigDeploymentGVK)
	export.SetLabels(u, a.actual.Labels)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DeploymentAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Deployment", "name", a.id)

	req := &configpb.DeleteDeploymentRequest{
		Name:  a.id.String(),
		Force: true,
	}
	op, err := a.gcpClient.DeleteDeployment(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Deployment, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Deployment %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Deployment", "name", a.id)
	_, err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Deployment %s: %w", a.id, err)
	}
	return true, nil
}
