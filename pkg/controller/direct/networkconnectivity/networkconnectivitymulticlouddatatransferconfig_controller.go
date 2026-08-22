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
// See the License for the3 specific language governing permissions and
// limitations under the License.

package networkconnectivity

import (
	"context"
	"fmt"
	"strings"
	"time"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

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
	"google.golang.org/protobuf/proto"
)

func init() {
	registry.RegisterModel(krm.NetworkConnectivityMulticloudDataTransferConfigGVK, newMulticloudDataTransferConfigModel)
}

func newMulticloudDataTransferConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &multicloudDataTransferConfigModel{config: config}, nil
}

type multicloudDataTransferConfigModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &multicloudDataTransferConfigModel{}

type multicloudDataTransferConfigAdapter struct {
	projectID  string
	location   string
	resourceID string

	id      *krm.NetworkConnectivityMulticloudDataTransferConfigIdentity
	desired *pb.MulticloudDataTransferConfig
	actual  *pb.MulticloudDataTransferConfig

	gcpClient *api.Service
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &multicloudDataTransferConfigAdapter{}

// AdapterForObject implements the Model interface.
func (m *multicloudDataTransferConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	clientBuilder, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	gcpClient, err := clientBuilder.newNetworkConnectivityClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.NetworkConnectivityMulticloudDataTransferConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idIdentity, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}
	id := idIdentity.(*krm.NetworkConnectivityMulticloudDataTransferConfigIdentity)

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkConnectivityMulticloudDataTransferConfigSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &multicloudDataTransferConfigAdapter{
		projectID:  id.Project,
		location:   id.Location,
		resourceID: id.MulticloudDataTransferConfig,
		id:         id,
		desired:    desiredProto,
		gcpClient:  gcpClient,
	}, nil
}

func (m *multicloudDataTransferConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find retrieves the GCP resource.
func (a *multicloudDataTransferConfigAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	fqn := a.id.String()
	actual, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Get(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, err
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

// Delete implements the Adapter interface.
func (a *multicloudDataTransferConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	exists, err := a.Find(ctx)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}

	fqn := a.id.String()

	op, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Delete(fqn).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting MulticloudDataTransferConfig %q: %w", fqn, err)
	}

	if err := a.waitForOperation(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for delete of MulticloudDataTransferConfig %q: %w", fqn, err)
	}

	return true, nil
}

// Create implements the Adapter interface.
func (a *multicloudDataTransferConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	fqn := a.id.String()

	req := &api.MulticloudDataTransferConfig{}
	if err := convertProtoToAPI(a.desired, req); err != nil {
		return err
	}

	log.V(0).Info("creating MulticloudDataTransferConfig", "req", req)
	op, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Create(a.id.ParentString(), req).MulticloudDataTransferConfigId(a.resourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating MulticloudDataTransferConfig: %w", err)
	}
	if err := a.waitForOperation(ctx, op); err != nil {
		return fmt.Errorf("waiting for create of MulticloudDataTransferConfig %q: %w", fqn, err)
	}

	created, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Get(fqn).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created MulticloudDataTransferConfig %q: %w", fqn, err)
	}
	log.V(2).Info("created MulticloudDataTransferConfig", "MulticloudDataTransferConfig", created)

	resourceID := lastComponent(created.Name)
	if err := unstructured.SetNestedField(u.Object, resourceID, "spec", "resourceID"); err != nil {
		return fmt.Errorf("setting spec.resourceID: %w", err)
	}

	var createdPB *pb.MulticloudDataTransferConfig
	if err := convertAPIToProto(created, &createdPB); err != nil {
		return err
	}

	return a.updateStatus(ctx, createOp, createdPB)
}

// Update implements the Adapter interface.
func (a *multicloudDataTransferConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating object", "u", u)

	fqn := a.id.String()

	mapCtx := &direct.MapContext{}
	maskedActualSpec := NetworkConnectivityMulticloudDataTransferConfigSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual := NetworkConnectivityMulticloudDataTransferConfigSpec_ToProto(mapCtx, maskedActualSpec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	maskedActual.Name = a.id.String()

	clonedDesired := proto.Clone(a.desired).(*pb.MulticloudDataTransferConfig)
	clonedDesired.Name = a.id.String()

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		structuredreporting.ReportDiff(ctx, diffs)

		req := &api.MulticloudDataTransferConfig{}
		if err := convertProtoToAPI(clonedDesired, req); err != nil {
			return err
		}

		log.V(2).Info("updating MulticloudDataTransferConfig", "request", req)
		op, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Patch(fqn, req).UpdateMask(strings.Join(updateMask.GetPaths(), ",")).Context(ctx).Do()
		if err != nil {
			return err
		}
		if err := a.waitForOperation(ctx, op); err != nil {
			return fmt.Errorf("waiting for update of MulticloudDataTransferConfig %q: %w", fqn, err)
		}

		updated, err := a.gcpClient.Projects.Locations.MulticloudDataTransferConfigs.Get(fqn).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting updated MulticloudDataTransferConfig %q: %w", fqn, err)
		}

		log.V(2).Info("updated MulticloudDataTransferConfig", "MulticloudDataTransferConfig", updated)
		if err := convertAPIToProto(updated, &a.actual); err != nil {
			return err
		}
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *multicloudDataTransferConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("MulticloudDataTransferConfig %q not found", a.id.String())
	}

	mc := &direct.MapContext{}
	spec := NetworkConnectivityMulticloudDataTransferConfigSpec_FromProto(mc, a.actual)
	if err := mc.Err(); err != nil {
		return nil, fmt.Errorf("error converting MulticloudDataTransferConfig from API %w", err)
	}

	spec.ProjectRef = &refs.ProjectRef{External: a.projectID}
	spec.Location = &a.location

	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting MulticloudDataTransferConfig spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.resourceID)
	u.SetGroupVersionKind(krm.NetworkConnectivityMulticloudDataTransferConfigGVK)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
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
