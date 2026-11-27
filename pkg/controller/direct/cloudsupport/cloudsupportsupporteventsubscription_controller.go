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
// krm.group: cloudsupport.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: CloudSupportSupportEventSubscription

package cloudsupport

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudsupport/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/support/apiv2"
	supportpb "cloud.google.com/go/support/apiv2/supportpb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CloudSupportSupportEventSubscriptionGVK, NewSupportEventSubscriptionModel)
}

func NewSupportEventSubscriptionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelSupportEventSubscription{config: *config}, nil
}

var _ directbase.Model = &modelSupportEventSubscription{}

type modelSupportEventSubscription struct {
	config config.ControllerConfig
}

func (m *modelSupportEventSubscription) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudSupportSupportEventSubscription{}
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
	id := i.(*krm.CloudSupportSupportEventSubscriptionIdentity)

	// Get cloudsupport GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SupportEventSubscriptionAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelSupportEventSubscription) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//cloudsupport.googleapis.com/") {
		return nil, nil
	}
	id := &krm.CloudSupportSupportEventSubscriptionIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &SupportEventSubscriptionAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelSupportEventSubscription) client(ctx context.Context) (*gcp.SupportEventSubscriptionClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewSupportEventSubscriptionRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building SupportEventSubscription client: %w", err)
	}
	return gcpClient, err
}

type SupportEventSubscriptionAdapter struct {
	id        *krm.CloudSupportSupportEventSubscriptionIdentity
	gcpClient *gcp.SupportEventSubscriptionClient
	desired   *krm.CloudSupportSupportEventSubscription
	actual    *supportpb.SupportEventSubscription
}

var _ directbase.Adapter = &SupportEventSubscriptionAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *SupportEventSubscriptionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting SupportEventSubscription", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &supportpb.GetSupportEventSubscriptionRequest{
		Name: a.id.String(),
	}
	supporteventsubscriptionpb, err := a.gcpClient.GetSupportEventSubscription(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting SupportEventSubscription %q: %w", a.id, err)
	}

	a.actual = supporteventsubscriptionpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SupportEventSubscriptionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating SupportEventSubscription", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := CloudSupportSupportEventSubscriptionSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &supportpb.CreateSupportEventSubscriptionRequest{
		Parent:                   a.id.ParentString(),
		SupportEventSubscription: resource,
	}
	created, err := a.gcpClient.CreateSupportEventSubscription(ctx, req)
	if err != nil {
		return fmt.Errorf("creating SupportEventSubscription %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created SupportEventSubscription", "name", created.GetName())

	if err := a.id.FromExternal(created.GetName()); err != nil {
		return fmt.Errorf("parsing created SupportEventSubscription name %q: %w", created.GetName(), err)
	}

	status := &krm.CloudSupportSupportEventSubscriptionStatus{}
	status.ObservedState = CloudSupportSupportEventSubscriptionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *SupportEventSubscriptionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating SupportEventSubscription", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		CloudSupportSupportEventSubscriptionSpec_FromProto,
		CloudSupportSupportEventSubscriptionSpec_ToProto,
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
		req := &supportpb.UpdateSupportEventSubscriptionRequest{
			UpdateMask:               updateMask,
			SupportEventSubscription: resource,
		}
		updated, err = a.gcpClient.UpdateSupportEventSubscription(ctx, req)
		if err != nil {
			return fmt.Errorf("updating SupportEventSubscription %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated SupportEventSubscription", "name", a.id)
	}

	status := &krm.CloudSupportSupportEventSubscriptionStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = CloudSupportSupportEventSubscriptionObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *SupportEventSubscriptionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CloudSupportSupportEventSubscription{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(CloudSupportSupportEventSubscriptionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.OrganizationRef = &refsv1beta1.OrganizationRef{External: a.id.ParentString()}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.SupportEventSubscription)
	u.SetGroupVersionKind(krm.CloudSupportSupportEventSubscriptionGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *SupportEventSubscriptionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting SupportEventSubscription", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	req := &supportpb.DeleteSupportEventSubscriptionRequest{
		Name: a.id.String(),
	}
	_, err := a.gcpClient.DeleteSupportEventSubscription(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent SupportEventSubscription, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting SupportEventSubscription %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted SupportEventSubscription", "name", a.id)
	return true, nil
}
