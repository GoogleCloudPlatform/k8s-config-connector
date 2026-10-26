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
// krm.kind: ComposerUserWorkloadsSecret

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
	registry.RegisterModel(krm.ComposerUserWorkloadsSecretGVK, NewUserWorkloadsSecretModel)
}

func NewUserWorkloadsSecretModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelUserWorkloadsSecret{config: *config}, nil
}

var _ directbase.Model = &modelUserWorkloadsSecret{}

type modelUserWorkloadsSecret struct {
	config config.ControllerConfig
}

func (m *modelUserWorkloadsSecret) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComposerUserWorkloadsSecret{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ComposerUserWorkloadsSecretIdentity)

	// Get composer GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &UserWorkloadsSecretAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelUserWorkloadsSecret) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//composer.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ComposerUserWorkloadsSecretIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &UserWorkloadsSecretAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelUserWorkloadsSecret) client(ctx context.Context) (*gcp.EnvironmentsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewEnvironmentsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building UserWorkloadsSecret client: %w", err)
	}
	return gcpClient, err
}

type UserWorkloadsSecretAdapter struct {
	id        *krm.ComposerUserWorkloadsSecretIdentity
	gcpClient *gcp.EnvironmentsClient
	desired   *krm.ComposerUserWorkloadsSecret
	actual    *servicepb.UserWorkloadsSecret
}

var _ directbase.Adapter = &UserWorkloadsSecretAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *UserWorkloadsSecretAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting UserWorkloadsSecret", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &servicepb.GetUserWorkloadsSecretRequest{
		Name: a.id.String(),
	}
	userworkloadssecretpb, err := a.gcpClient.GetUserWorkloadsSecret(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting UserWorkloadsSecret %q: %w", a.id, err)
	}

	a.actual = userworkloadssecretpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *UserWorkloadsSecretAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating UserWorkloadsSecret", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ComposerUserWorkloadsSecretSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &servicepb.CreateUserWorkloadsSecretRequest{
		Parent:              a.id.ParentString(),
		UserWorkloadsSecret: resource,
	}
	created, err := a.gcpClient.CreateUserWorkloadsSecret(ctx, req)
	if err != nil {
		return fmt.Errorf("creating UserWorkloadsSecret %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created UserWorkloadsSecret", "name", a.id)

	status := &krm.ComposerUserWorkloadsSecretStatus{}
	parsedID := &krm.ComposerUserWorkloadsSecretIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.UserWorkloadsSecret = parsedID.UserWorkloadsSecret
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *UserWorkloadsSecretAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating UserWorkloadsSecret", "name", a.id)

	resource, report, _, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ComposerUserWorkloadsSecretSpec_v1alpha1_FromProto,
		ComposerUserWorkloadsSecretSpec_v1alpha1_ToProto,
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
		req := &servicepb.UpdateUserWorkloadsSecretRequest{
			UserWorkloadsSecret: resource,
		}
		_, err = a.gcpClient.UpdateUserWorkloadsSecret(ctx, req)
		if err != nil {
			return fmt.Errorf("updating UserWorkloadsSecret %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated UserWorkloadsSecret", "name", a.id)
	}

	status := &krm.ComposerUserWorkloadsSecretStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *UserWorkloadsSecretAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComposerUserWorkloadsSecret{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComposerUserWorkloadsSecretSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ComposerUserWorkloadsSecretGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *UserWorkloadsSecretAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting UserWorkloadsSecret", "name", a.id)

	req := &servicepb.DeleteUserWorkloadsSecretRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteUserWorkloadsSecret(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent UserWorkloadsSecret, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting UserWorkloadsSecret %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted UserWorkloadsSecret", "name", a.id)
	return true, nil
}
