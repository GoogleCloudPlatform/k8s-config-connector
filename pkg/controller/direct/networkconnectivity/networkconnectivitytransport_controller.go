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

package networkconnectivity

import (
	"context"
	"fmt"
	"strings"
	"time"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	api "google.golang.org/api/networkconnectivity/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.NetworkConnectivityTransportGVK, NewTransportModel)
}

func NewTransportModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &transportModel{config: *config}, nil
}

var _ directbase.Model = &transportModel{}

type transportModel struct {
	config config.ControllerConfig
}

func (m *transportModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkConnectivityTransport{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idIdentity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idIdentity.(*krm.NetworkConnectivityTransportIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := NetworkConnectivityTransportSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newNetworkConnectivityClient(ctx)
	if err != nil {
		return nil, err
	}

	return &transportAdapter{
		gcpClient:  client,
		id:         id,
		desiredKRM: &obj.Spec,
		desired:    desired,
	}, nil
}

func (m *transportModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type transportAdapter struct {
	gcpClient  *api.Service
	id         *krm.NetworkConnectivityTransportIdentity
	desiredKRM *krm.NetworkConnectivityTransportSpec
	desired    *pb.Transport
	actual     *pb.Transport
}

var _ directbase.Adapter = &transportAdapter{}

func (a *transportAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networkconnectivity transport", "name", a.id)
	fqn := a.id.String()
	actual, err := a.gcpClient.Projects.Locations.Transports.Get(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networkconnectivity transport %q from gcp: %w", fqn, err)
	}

	if err := convertAPIToProto(actual, &a.actual); err != nil {
		return false, err
	}

	return true, nil
}

func (a *transportAdapter) waitForOperation(ctx context.Context, op *api.GoogleLongrunningOperation) error {
	_, err := common.WaitForOperation(ctx, 2*time.Second, func(current *api.GoogleLongrunningOperation) (bool, error) {
		if !current.Done {
			return false, nil
		}
		if current.Error != nil {
			return true, fmt.Errorf("operation failed with code %d: %s", current.Error.Code, current.Error.Message)
		}
		return true, nil
	}, func() (*api.GoogleLongrunningOperation, error) {
		return a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
	})
	return err
}

func (a *transportAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()
	log := klog.FromContext(ctx)
	log.V(2).Info("creating networkconnectivity transport", "name", a.id)

	req := &api.Transport{}
	if err := convertProtoToAPI(a.desired, req); err != nil {
		return err
	}

	fqn := a.id.String()
	parent := a.id.ParentString()
	op, err := a.gcpClient.Projects.Locations.Transports.Create(parent, req).TransportId(a.id.Transport).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating networkconnectivity transport %s: %w", fqn, err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for create of transport %q: %w", fqn, err)
	}

	log.V(2).Info("successfully created networkconnectivity transport in gcp", "name", a.id)

	created, err := a.gcpClient.Projects.Locations.Transports.Get(fqn).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created transport %q: %w", fqn, err)
	}

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	var createdPB *pb.Transport
	if err := convertAPIToProto(created, &createdPB); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, createdPB)
}

func (a *transportAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networkconnectivity transport", "name", a.id)

	diffs, updateMask, err := compareTransport(ctx, a.desiredKRM, a.actual)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)
		req := &api.Transport{}
		if err := convertProtoToAPI(a.desired, req); err != nil {
			return err
		}

		fqn := a.id.String()
		paths := strings.Join(updateMask.GetPaths(), ",")
		op, err := a.gcpClient.Projects.Locations.Transports.Patch(fqn, req).UpdateMask(paths).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating networkconnectivity transport %s: %w", fqn, err)
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for update of transport %q: %w", fqn, err)
		}

		log.V(2).Info("successfully updated networkconnectivity transport", "name", fqn)

		updated, err := a.gcpClient.Projects.Locations.Transports.Get(fqn).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated transport %q: %w", fqn, err)
		}
		if err := convertAPIToProto(updated, &a.actual); err != nil {
			return err
		}
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func compareTransport(ctx context.Context, desiredKRM *krm.NetworkConnectivityTransportSpec, actual *pb.Transport) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	return common.CompareBrownfieldSpec(
		ctx,
		desiredKRM,
		actual,
		NetworkConnectivityTransportSpec_FromProto,
		NetworkConnectivityTransportSpec_ToProto,
		nil,
	)
}

func (a *transportAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networkconnectivity transport", "name", a.id)

	fqn := a.id.String()
	op, err := a.gcpClient.Projects.Locations.Transports.Delete(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting networkconnectivity transport %s: %w", fqn, err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for delete of transport %q: %w", fqn, err)
	}

	return true, nil
}

func (a *transportAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("transport %q not found", a.id.String())
	}

	mc := &direct.MapContext{}
	spec := NetworkConnectivityTransportSpec_FromProto(mc, a.actual)
	if err := mc.Err(); err != nil {
		return nil, fmt.Errorf("error converting transport from proto: %w", err)
	}

	spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	spec.Location = &a.id.Location
	spec.ResourceID = &a.id.Transport

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting transport spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.id.Transport)
	u.SetGroupVersionKind(krm.NetworkConnectivityTransportGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *transportAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Transport) error {
	mapCtx := &direct.MapContext{}
	observedState := NetworkConnectivityTransportObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status := &krm.NetworkConnectivityTransportStatus{
		ExternalRef:   &externalRef,
		ObservedState: observedState,
	}

	return op.UpdateStatus(ctx, status, nil)
}
