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
// krm.group: configdelivery.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ConfigDeliveryRelease

package configdelivery

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdelivery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/configdelivery/apiv1"
	configdeliverypb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ConfigDeliveryReleaseGVK, NewReleaseModel)
}

func NewReleaseModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelRelease{config: *config}, nil
}

var _ directbase.Model = &modelRelease{}

type modelRelease struct {
	config config.ControllerConfig
}

func (m *modelRelease) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ConfigDeliveryRelease{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ConfigDeliveryReleaseIdentity)

	// Get configdelivery GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReleaseAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelRelease) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//configdelivery.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ConfigDeliveryReleaseIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReleaseAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelRelease) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Release client: %w", err)
	}
	return gcpClient, err
}

type ReleaseAdapter struct {
	id        *krm.ConfigDeliveryReleaseIdentity
	gcpClient *gcp.Client
	desired   *krm.ConfigDeliveryRelease
	actual    *configdeliverypb.Release
}

var _ directbase.Adapter = &ReleaseAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ReleaseAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Release", "name", a.id)

	req := &configdeliverypb.GetReleaseRequest{
		Name: a.id.String(),
	}
	releasepb, err := a.gcpClient.GetRelease(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Release %q: %w", a.id, err)
	}

	a.actual = releasepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReleaseAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Release", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ConfigDeliveryReleaseSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &configdeliverypb.CreateReleaseRequest{
		Parent:    a.id.ParentString(),
		ReleaseId: a.id.Release,
		Release:   resource,
	}
	op, err := a.gcpClient.CreateRelease(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Release %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Release %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Release", "name", a.id)

	status := &krm.ConfigDeliveryReleaseStatus{}
	status.ObservedState = ConfigDeliveryReleaseObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReleaseAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Release", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ConfigDeliveryReleaseSpec_FromProto,
		ConfigDeliveryReleaseSpec_ToProto,
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
		req := &configdeliverypb.UpdateReleaseRequest{
			UpdateMask: updateMask,
			Release:    resource,
		}
		op, err := a.gcpClient.UpdateRelease(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Release %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Release %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Release", "name", a.id)
	}

	status := &krm.ConfigDeliveryReleaseStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ConfigDeliveryReleaseObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReleaseAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ConfigDeliveryRelease{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConfigDeliveryReleaseSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ConfigDeliveryReleaseGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReleaseAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Release", "name", a.id)

	req := &configdeliverypb.DeleteReleaseRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteRelease(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Release, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Release %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Release", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Release %s: %w", a.id, err)
	}
	return true, nil
}
