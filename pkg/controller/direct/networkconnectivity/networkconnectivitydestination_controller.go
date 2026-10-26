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
// krm.group: networkconnectivity.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: NetworkConnectivityDestination

package networkconnectivity

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/networkconnectivity/apiv1"
	networkconnectivitypb "cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkConnectivityDestinationGVK, NewDestinationModel)
}

func NewDestinationModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelDestination{config: *config}, nil
}

var _ directbase.Model = &modelDestination{}

type modelDestination struct {
	config config.ControllerConfig
}

func (m *modelDestination) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkConnectivityDestination{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.NetworkConnectivityDestinationIdentity)

	// Get networkconnectivity GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DestinationAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelDestination) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//networkconnectivity.googleapis.com/") {
		return nil, nil
	}
	id := &krm.NetworkConnectivityDestinationIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &DestinationAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelDestination) client(ctx context.Context) (*gcp.DataTransferClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewDataTransferClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Destination client: %w", err)
	}
	return gcpClient, err
}

type DestinationAdapter struct {
	id        *krm.NetworkConnectivityDestinationIdentity
	gcpClient *gcp.DataTransferClient
	desired   *krm.NetworkConnectivityDestination
	actual    *networkconnectivitypb.Destination
}

var _ directbase.Adapter = &DestinationAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *DestinationAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Destination", "name", a.id)

	req := &networkconnectivitypb.GetDestinationRequest{
		Name: a.id.String(),
	}
	destinationpb, err := a.gcpClient.GetDestination(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Destination %q: %w", a.id, err)
	}

	a.actual = destinationpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DestinationAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Destination", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkConnectivityDestinationSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &networkconnectivitypb.CreateDestinationRequest{
		Parent:        a.id.ParentString(),
		DestinationId: a.id.Destination,
		Destination:   resource,
	}
	op, err := a.gcpClient.CreateDestination(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Destination %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("Destination %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created Destination", "name", a.id)

	status := &krm.NetworkConnectivityDestinationStatus{}
	status.ObservedState = NetworkConnectivityDestinationObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *DestinationAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Destination", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpec(
		ctx,
		&a.desired.Spec,
		a.actual,
		NetworkConnectivityDestinationSpec_FromProto,
		NetworkConnectivityDestinationSpec_ToProto,
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
		req := &networkconnectivitypb.UpdateDestinationRequest{
			UpdateMask:  updateMask,
			Destination: resource,
		}
		op, err := a.gcpClient.UpdateDestination(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Destination %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("Destination %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Destination", "name", a.id)
	}

	status := &krm.NetworkConnectivityDestinationStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = NetworkConnectivityDestinationObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *DestinationAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkConnectivityDestination{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkConnectivityDestinationSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NetworkConnectivityDestinationGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *DestinationAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Destination", "name", a.id)

	req := &networkconnectivitypb.DeleteDestinationRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteDestination(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Destination, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Destination %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Destination", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete Destination %s: %w", a.id, err)
	}
	return true, nil
}
