// Copyright 2024 Google LLC
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

// +tool:controller
// proto.service: google.cloud.edgecontainer.v1.Service
// proto.message: google.cloud.edgecontainer.v1.VpnConnection
// crd.type: EdgeContainerVpnConnection
// crd.version: v1alpha1

package edgecontainer

import (
	"context"
	"fmt"

	edgecontainer "cloud.google.com/go/edgecontainer/apiv1"
	edgecontainerpb "cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/edgecontainer/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.EdgeContainerVpnConnectionGVK, NewVPNConnectionModel)
}

func NewVPNConnectionModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &vpnConnectionModel{config: *config}, nil
}

var _ directbase.Model = &vpnConnectionModel{}

type vpnConnectionModel struct {
	config config.ControllerConfig
}

func (m *vpnConnectionModel) client(ctx context.Context, projectID string) (*edgecontainer.Client, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := edgecontainer.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building edgecontainer vpnconnection client: %w", err)
	}

	return gcpClient, err
}

func (m *vpnConnectionModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.EdgeContainerVpnConnection{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewVpnConnectionIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &vpnConnectionAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *vpnConnectionModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type vpnConnectionAdapter struct {
	gcpClient *edgecontainer.Client
	id        *krm.VpnConnectionIdentity
	desired   *krm.EdgeContainerVpnConnection
	actual    *edgecontainerpb.VpnConnection
	reader    client.Reader
}

var _ directbase.Adapter = &vpnConnectionAdapter{}

func (a *vpnConnectionAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("getting vpn connection", "name", a.id)

	req := &edgecontainerpb.GetVpnConnectionRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetVpnConnection(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting vpn connection %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *vpnConnectionAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.Info("creating vpn connection", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := EdgeContainerVpnConnectionSpec_ToProto(mapCtx, &desired.Spec)
	if err := normalizeReferences(ctx, a.reader, a.desired); err != nil {
		return err
	}

	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	req := &edgecontainerpb.CreateVpnConnectionRequest{
		Parent:          a.id.Parent().String(),
		VpnConnectionId: a.id.ID(),
		VpnConnection:   resource,
	}
	op, err := a.gcpClient.CreateVpnConnection(ctx, req)
	if err != nil {
		return fmt.Errorf("creating vpn connection %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("vpn connection %s waiting creation: %w", a.id.String(), err)
	}
	log.Info("successfully created vpn connection in gcp", "name", a.id)

	status := &krm.EdgeContainerVpnConnectionStatus{}
	status.ObservedState = EdgeContainerVpnConnectionObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// EdgeContainerVpnConnection does not support update.
func (a *vpnConnectionAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// EdgeContainerVPNConnection does not support update.
	log := klog.FromContext(ctx)
	log.Info("updating vpn connection", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	desiredpb := EdgeContainerVpnConnectionSpec_ToProto(mapCtx, &desired.Spec)
	paths, err := common.CompareProtoMessage(desiredpb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	if len(paths) != 0 {
		log.V(2).Info("This resource does not support update", "name", a.id.String())
		return nil
	}

	status := &krm.EdgeContainerVpnConnectionStatus{}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *vpnConnectionAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.Info("deleting vpn connection", "name", a.id)

	req := &edgecontainerpb.DeleteVpnConnectionRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteVpnConnection(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting vpn connection %s: %w", a.id.String(), err)
	}
	log.Info("successfully deleted vpn connection", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of vpn connection %s: %w", a.id.String(), err)
		}
	}

	return true, nil
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *vpnConnectionAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EdgeContainerVpnConnection{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EdgeContainerVpnConnectionSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &v1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.EdgeContainerVpnConnectionGVK)

	u.Object = uObj
	return u, nil
}

func normalizeReferences(ctx context.Context, reader client.Reader, desired *krm.EdgeContainerVpnConnection) error {
	if desired.Spec.EdgeContainerClusterRef != nil {
		if _, err := desired.Spec.EdgeContainerClusterRef.NormalizedExternal(ctx, reader, desired.GetNamespace()); err != nil {
			return err
		}
	}
	return nil
}
