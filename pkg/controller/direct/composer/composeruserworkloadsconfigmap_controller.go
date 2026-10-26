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
// krm.group: composer.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ComposerUserWorkloadsConfigMap

package composer

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/composer/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/orchestration/airflow/service/apiv1"
	servicepb "cloud.google.com/go/orchestration/airflow/service/apiv1/servicepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ComposerUserWorkloadsConfigMapGVK, NewUserWorkloadsConfigMapModel)
}

func NewUserWorkloadsConfigMapModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelUserWorkloadsConfigMap{config: *config}, nil
}

var _ directbase.Model = &modelUserWorkloadsConfigMap{}

type modelUserWorkloadsConfigMap struct {
	config config.ControllerConfig
}

func (m *modelUserWorkloadsConfigMap) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComposerUserWorkloadsConfigMap{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComposerUserWorkloadsConfigMapIdentity)

	// Get composer GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &UserWorkloadsConfigMapAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelUserWorkloadsConfigMap) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//composer.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComposerUserWorkloadsConfigMapIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &UserWorkloadsConfigMapAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelUserWorkloadsConfigMap) client(ctx context.Context) (*gcp.EnvironmentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewEnvironmentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building UserWorkloadsConfigMap client: %w", err)
	}
	return gcpClient, err
}

type UserWorkloadsConfigMapAdapter struct {
	id        *krm.ComposerUserWorkloadsConfigMapIdentity
	gcpClient *gcp.EnvironmentsClient
	desired   *krm.ComposerUserWorkloadsConfigMap
	actual    *servicepb.UserWorkloadsConfigMap
}

var _ directbase.Adapter = &UserWorkloadsConfigMapAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *UserWorkloadsConfigMapAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting UserWorkloadsConfigMap", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &servicepb.GetUserWorkloadsConfigMapRequest{
		Name: a.id.String(),
	}
	userworkloadsconfigmappb, err := a.gcpClient.GetUserWorkloadsConfigMap(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting UserWorkloadsConfigMap %q: %w", a.id, err)
	}

	a.actual = userworkloadsconfigmappb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *UserWorkloadsConfigMapAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating UserWorkloadsConfigMap", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComposerUserWorkloadsConfigMapSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &servicepb.CreateUserWorkloadsConfigMapRequest{
		Parent:                 a.id.ParentString(),
		UserWorkloadsConfigMap: resource,
	}
	created, err := a.gcpClient.CreateUserWorkloadsConfigMap(ctx, req)
	if err != nil {
		return fmt.Errorf("creating UserWorkloadsConfigMap %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created UserWorkloadsConfigMap", "name", a.id)

	status := &krm.ComposerUserWorkloadsConfigMapStatus{}
	parsedID := &krm.ComposerUserWorkloadsConfigMapIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.UserWorkloadsConfigMap = parsedID.UserWorkloadsConfigMap
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *UserWorkloadsConfigMapAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating UserWorkloadsConfigMap", "name", a.id)

	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComposerUserWorkloadsConfigMapSpec_v1alpha1_FromProto,
		ComposerUserWorkloadsConfigMapSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &servicepb.UpdateUserWorkloadsConfigMapRequest{
			UserWorkloadsConfigMap: resource,
		}
		_, err = a.gcpClient.UpdateUserWorkloadsConfigMap(ctx, req)
		if err != nil {
			return fmt.Errorf("updating UserWorkloadsConfigMap %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated UserWorkloadsConfigMap", "name", a.id)
	}

	status := &krm.ComposerUserWorkloadsConfigMapStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *UserWorkloadsConfigMapAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComposerUserWorkloadsConfigMap{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComposerUserWorkloadsConfigMapSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComposerUserWorkloadsConfigMapGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *UserWorkloadsConfigMapAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting UserWorkloadsConfigMap", "name", a.id)

	req := &servicepb.DeleteUserWorkloadsConfigMapRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteUserWorkloadsConfigMap(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent UserWorkloadsConfigMap, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting UserWorkloadsConfigMap %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted UserWorkloadsConfigMap", "name", a.id)
	return true, nil
}
