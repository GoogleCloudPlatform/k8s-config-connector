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
// krm.group: resourcemanager.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ResourceManagerTagHold

package resourcemanager

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/resourcemanager/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/resourcemanager/apiv3"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ResourceManagerTagHoldGVK, NewTagHoldModel)
}

func NewTagHoldModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTagHold{config: *config}, nil
}

var _ directbase.Model = &modelTagHold{}

type modelTagHold struct {
	config config.ControllerConfig
}

func (m *modelTagHold) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ResourceManagerTagHold{}
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
	id := i.(*krm.ResourceManagerTagHoldIdentity)

	// Get resourcemanager GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TagHoldAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTagHold) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//cloudresourcemanager.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ResourceManagerTagHoldIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TagHoldAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelTagHold) client(ctx context.Context) (*gcp.TagHoldsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewTagHoldsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TagHold client: %w", err)
	}
	return gcpClient, err
}

type TagHoldAdapter struct {
	id        *krm.ResourceManagerTagHoldIdentity
	gcpClient *gcp.TagHoldsClient
	desired   *krm.ResourceManagerTagHold
	actual    *resourcemanagerpb.TagHold
}

var _ directbase.Adapter = &TagHoldAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TagHoldAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TagHold", "name", a.id)
	if !a.id.HasIdentitySpecified() {
		return false, nil
	}

	req := &resourcemanagerpb.ListTagHoldsRequest{
		Parent: a.id.ParentString(),
	}
	it := a.gcpClient.ListTagHolds(ctx, req)
	for {
		item, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			if direct.IsNotFound(err) {
				return false, nil
			}
			return false, fmt.Errorf("listing TagHold in %q: %w", a.id.ParentString(), err)
		}
		if item.Name == a.id.String() {
			a.actual = item
			return true, nil
		}
	}
	return false, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TagHoldAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TagHold", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ResourceManagerTagHoldSpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &resourcemanagerpb.CreateTagHoldRequest{
		Parent:  a.id.ParentString(),
		TagHold: resource,
	}

	// Newly created TagValues in GCP Resource Manager require ~60-75 seconds for IAM permissions
	// (resourcemanager.tagHolds.create) to replicate to the TagHold backend.
	time.Sleep(75 * time.Second)

	var op *gcp.CreateTagHoldOperation
	err := wait.PollUntilContextTimeout(ctx, 5*time.Second, 1*time.Minute, true, func(ctx context.Context) (bool, error) {
		var createErr error
		op, createErr = a.gcpClient.CreateTagHold(ctx, req)
		if createErr != nil {
			if direct.IsPermissionDenied(createErr) || strings.Contains(createErr.Error(), "(or it may not exist)") {
				log.V(2).Info("tagValue not yet propagated to TagHold service, retrying CreateTagHold", "error", createErr)
				return false, nil
			}
			return false, createErr
		}
		return true, nil
	})
	if err != nil {
		return fmt.Errorf("creating TagHold %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("TagHold %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created TagHold", "name", a.id)

	status := &krm.ResourceManagerTagHoldStatus{}
	status.ObservedState = ResourceManagerTagHoldObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	parsedID := &krm.ResourceManagerTagHoldIdentity{}
	if err := parsedID.FromExternal(created.Name); err != nil {
		return err
	}
	a.id.TagHold = parsedID.TagHold
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TagHoldAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating TagHold", "name", a.id)

	_, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ResourceManagerTagHoldSpec_v1alpha1_FromProto,
		ResourceManagerTagHoldSpec_v1alpha1_ToProto,
		nil,
	)
	if err != nil {
		return err
	}
	if report.HasDiff() {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		return fmt.Errorf("updating ResourceManagerTagHold is not supported, fields modified: %v", updateMask.Paths)
	}

	status := &krm.ResourceManagerTagHoldStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ResourceManagerTagHoldObservedState_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TagHoldAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ResourceManagerTagHold{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ResourceManagerTagHoldSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ResourceManagerTagHoldGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TagHoldAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TagHold", "name", a.id)

	if !a.id.HasIdentitySpecified() {
		return true, nil
	}

	req := &resourcemanagerpb.DeleteTagHoldRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteTagHold(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent TagHold, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting TagHold %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted TagHold", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete TagHold %s: %w", a.id, err)
	}
	return true, nil
}
