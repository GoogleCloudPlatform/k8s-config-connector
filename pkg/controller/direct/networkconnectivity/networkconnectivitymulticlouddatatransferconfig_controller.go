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

// +tool:controller
// proto.service: google.cloud.networkconnectivity.v1.HubService
// proto.message: google.cloud.networkconnectivity.v1.MulticloudDataTransferConfig
// crd.type: NetworkConnectivityMulticloudDataTransferConfig
// crd.version: v1alpha1

package networkconnectivity

import (
	"context"
	"fmt"
	"strings"
	"time"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/networkconnectivity/v1"
	api "google.golang.org/api/networkconnectivity/v1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkconnectivity/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.NetworkConnectivityMulticloudDataTransferConfigGVK, NewMulticloudDataTransferConfigModel)
}

func NewMulticloudDataTransferConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &multicloudDataTransferConfigModel{config: *config}, nil
}

var _ directbase.Model = &multicloudDataTransferConfigModel{}

type multicloudDataTransferConfigModel struct {
	config config.ControllerConfig
}

func (m *multicloudDataTransferConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.NetworkConnectivityMulticloudDataTransferConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idIdentity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idIdentity.(*krm.NetworkConnectivityMulticloudDataTransferConfigIdentity)

	// normalize reference fields
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Get networkconnectivity GCP client
	gcpClient, err := newGCPClient(ctx, &m.config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newNetworkConnectivityClient(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkConnectivityMulticloudDataTransferConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &multicloudDataTransferConfigAdapter{
		gcpClient: client,
		id:        id,
		desired:   desiredProto,
	}, nil
}

func (m *multicloudDataTransferConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs/Export
	return nil, nil
}

type multicloudDataTransferConfigAdapter struct {
	gcpClient *api.Service
	id        *krm.NetworkConnectivityMulticloudDataTransferConfigIdentity
	desired   *pb.MulticloudDataTransferConfig
	actual    *pb.MulticloudDataTransferConfig
}

var _ directbase.Adapter = &multicloudDataTransferConfigAdapter{}

func (a *multicloudDataTransferConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networkconnectivity multicloudDataTransferConfig", "name", a.id)
	fqn := a.id.String()
	actual, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Get(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networkconnectivity multicloudDataTransferConfig %q from gcp: %w", a.id.String(), err)
	}

	if err := convertAPIToProto(actual, &a.actual); err != nil {
		return false, err
	}

	return true, nil
}

func (a *multicloudDataTransferConfigAdapter) waitForOperation(ctx context.Context, op *api.GoogleLongrunningOperation) error {
	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		latest, err := a.gcpClient.Projects.Locations.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation %q: %w", op.Name, err)
		}

		if latest.Done {
			return nil
		}

		time.Sleep(2 * time.Second)
	}
}

func (a *multicloudDataTransferConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating networkconnectivity multicloudDataTransferConfig", "name", a.id)

	req := &api.MulticloudDataTransferConfig{}
	if err := convertProtoToAPI(a.desired, req); err != nil {
		return err
	}

	fqn := a.id.String()
	parent := a.id.ParentString()

	op, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Create(parent, req).MulticloudDataTransferConfigId(a.id.MulticloudDataTransferConfig).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating multicloudDataTransferConfig %q: %w", fqn, err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for create of multicloudDataTransferConfig %q: %w", fqn, err)
	}

	created, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Get(fqn).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created multicloudDataTransferConfig %q: %w", fqn, err)
	}

	if err := convertAPIToProto(created, &a.actual); err != nil {
		return err
	}

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	return a.updateStatus(ctx, createOp, a.actual)
}

func (a *multicloudDataTransferConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating networkconnectivity multicloudDataTransferConfig", "name", a.id.String())

	diffs, updateMask, err := compareConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &api.MulticloudDataTransferConfig{}
		if err := convertProtoToAPI(a.desired, req); err != nil {
			return err
		}

		fqn := a.id.String()
		updateMaskStr := strings.Join(updateMask.Paths, ",")
		op, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Patch(fqn, req).UpdateMask(updateMaskStr).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating multicloudDataTransferConfig %q: %w", fqn, err)
		}

		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for update of multicloudDataTransferConfig %q: %w", fqn, err)
		}

		updated, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Get(fqn).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated multicloudDataTransferConfig %q: %w", fqn, err)
		}

		if err := convertAPIToProto(updated, &latest); err != nil {
			return err
		}
		a.actual = latest
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *multicloudDataTransferConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.MulticloudDataTransferConfig) error {
	mapCtx := &direct.MapContext{}
	status := &krm.NetworkConnectivityMulticloudDataTransferConfigStatus{}
	status.ObservedState = NetworkConnectivityMulticloudDataTransferConfigObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return op.UpdateStatus(ctx, status, nil)
}

func (a *multicloudDataTransferConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting multicloudDataTransferConfig", "name", a.id)

	fqn := a.id.String()

	op, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Delete(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting multicloudDataTransferConfig %q: %w", fqn, err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for delete of multicloudDataTransferConfig %q: %w", fqn, err)
	}

	return true, nil
}

func compareConfig(ctx context.Context, actual, desired *pb.MulticloudDataTransferConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *pb.MulticloudDataTransferConfig
	{
		// A "trick" to only compare spec fields - round trip via the spec
		mapCtx := &direct.MapContext{}
		spec := NetworkConnectivityMulticloudDataTransferConfigSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = NetworkConnectivityMulticloudDataTransferConfigSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *multicloudDataTransferConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkConnectivityMulticloudDataTransferConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkConnectivityMulticloudDataTransferConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.NetworkConnectivityMulticloudDataTransferConfigGVK)
	return u, nil
}
