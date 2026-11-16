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
// krm.group: appengine.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: AppEngineAuthorizedCertificate

package appengine

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/appengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/appengine/apiv1"
	appenginepb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.AppEngineAuthorizedCertificateGVK, NewAuthorizedCertificateModel)
}

func NewAuthorizedCertificateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAuthorizedCertificate{config: *config}, nil
}

var _ directbase.Model = &modelAuthorizedCertificate{}

type modelAuthorizedCertificate struct {
	config config.ControllerConfig
}

func (m *modelAuthorizedCertificate) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AppEngineAuthorizedCertificate{}
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
	id := i.(*krm.AppEngineAuthorizedCertificateIdentity)

	// Get appengine GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AuthorizedCertificateAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAuthorizedCertificate) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//appengine.googleapis.com/") {
		return nil, nil
	}
	id := &krm.AppEngineAuthorizedCertificateIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AuthorizedCertificateAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelAuthorizedCertificate) client(ctx context.Context) (*gcp.AuthorizedCertificatesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAuthorizedCertificatesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AuthorizedCertificate client: %w", err)
	}
	return gcpClient, err
}

type AuthorizedCertificateAdapter struct {
	id        *krm.AppEngineAuthorizedCertificateIdentity
	gcpClient *gcp.AuthorizedCertificatesClient
	desired   *krm.AppEngineAuthorizedCertificate
	actual    *appenginepb.AuthorizedCertificate
}

var _ directbase.Adapter = &AuthorizedCertificateAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AuthorizedCertificateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting AuthorizedCertificate", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &appenginepb.GetAuthorizedCertificateRequest{
		Name: a.id.String(),
	}
	authorizedcertificatepb, err := a.gcpClient.GetAuthorizedCertificate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting AuthorizedCertificate %q: %w", a.id, err)
	}

	a.actual = authorizedcertificatepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AuthorizedCertificateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating AuthorizedCertificate", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AppEngineAuthorizedCertificateSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &appenginepb.CreateAuthorizedCertificateRequest{
		Parent:      a.id.ParentString(),
		Certificate: resource,
	}
	created, err := a.gcpClient.CreateAuthorizedCertificate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AuthorizedCertificate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created AuthorizedCertificate", "name", a.id)

	status := &krm.AppEngineAuthorizedCertificateStatus{}
	status.ObservedState = AppEngineAuthorizedCertificateObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.AppEngineAuthorizedCertificateIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.AuthorizedCertificate = parsedID.AuthorizedCertificate
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AuthorizedCertificateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating AuthorizedCertificate", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		AppEngineAuthorizedCertificateSpec_v1alpha1_FromProto,
		AppEngineAuthorizedCertificateSpec_v1alpha1_ToProto,
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
		req := &appenginepb.UpdateAuthorizedCertificateRequest{
			Name:        a.id.String(),
			UpdateMask:  updateMask,
			Certificate: resource,
		}
		updated, err = a.gcpClient.UpdateAuthorizedCertificate(ctx, req)
		if err != nil {
			return fmt.Errorf("updating AuthorizedCertificate %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated AuthorizedCertificate", "name", a.id)
	}

	status := &krm.AppEngineAuthorizedCertificateStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = AppEngineAuthorizedCertificateObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AuthorizedCertificateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AppEngineAuthorizedCertificate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AppEngineAuthorizedCertificateSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.AppEngineAuthorizedCertificateGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AuthorizedCertificateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting AuthorizedCertificate", "name", a.id)

	req := &appenginepb.DeleteAuthorizedCertificateRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteAuthorizedCertificate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent AuthorizedCertificate, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting AuthorizedCertificate %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AuthorizedCertificate", "name", a.id)
	return true, nil
}
