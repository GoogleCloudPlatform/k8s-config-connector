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
// krm.group: chronicle.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: ChronicleReferenceList

package chronicle

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/chronicle/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/chronicle/apiv1"
	chroniclepb "cloud.google.com/go/chronicle/apiv1/chroniclepb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.ChronicleReferenceListGVK, NewReferenceListModel)
}

func NewReferenceListModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelReferenceList{config: *config}, nil
}

var _ directbase.Model = &modelReferenceList{}

type modelReferenceList struct {
	config config.ControllerConfig
}

func (m *modelReferenceList) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ChronicleReferenceList{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.ChronicleReferenceListIdentity)

	// Get chronicle GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReferenceListAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelReferenceList) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//chronicle.googleapis.com/") {
		return nil, nil
	}
	id := &krm.ChronicleReferenceListIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ReferenceListAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelReferenceList) client(ctx context.Context) (*gcp.ReferenceListClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewReferenceListRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ReferenceList client: %w", err)
	}
	return gcpClient, err
}

type ReferenceListAdapter struct {
	id        *krm.ChronicleReferenceListIdentity
	gcpClient *gcp.ReferenceListClient
	desired   *krm.ChronicleReferenceList
	actual    *chroniclepb.ReferenceList
}

var _ directbase.Adapter = &ReferenceListAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *ReferenceListAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ReferenceList", "name", a.id)

	req := &chroniclepb.GetReferenceListRequest{
		Name: a.id.String(),
	}
	referencelistpb, err := a.gcpClient.GetReferenceList(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ReferenceList %q: %w", a.id, err)
	}

	a.actual = referencelistpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReferenceListAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ReferenceList", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := ChronicleReferenceListSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &chroniclepb.CreateReferenceListRequest{
		Parent:          a.id.ParentString(),
		ReferenceListId: a.id.ReferenceList,
		ReferenceList:   resource,
	}
	created, err := a.gcpClient.CreateReferenceList(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ReferenceList %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created ReferenceList", "name", a.id)

	status := &krm.ChronicleReferenceListStatus{}
	status.ObservedState = ChronicleReferenceListObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *ReferenceListAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ReferenceList", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		ChronicleReferenceListSpec_FromProto,
		ChronicleReferenceListSpec_ToProto,
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
		req := &chroniclepb.UpdateReferenceListRequest{
			UpdateMask:    updateMask,
			ReferenceList: resource,
		}
		updated, err = a.gcpClient.UpdateReferenceList(ctx, req)
		if err != nil {
			return fmt.Errorf("updating ReferenceList %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated ReferenceList", "name", a.id)
	}

	status := &krm.ChronicleReferenceListStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ChronicleReferenceListObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *ReferenceListAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ChronicleReferenceList{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ChronicleReferenceListSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.ChronicleReferenceListGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *ReferenceListAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deletion not supported for ReferenceList, removing finalizer without deleting GCP resource", "name", a.id)
	return true, nil
}
