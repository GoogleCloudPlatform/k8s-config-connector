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
// proto.service: google.cloud.discoveryengine.v1beta.ServingConfigService
// proto.message: google.cloud.discoveryengine.v1beta.ServingConfig
// crd.type: DiscoveryEngineServingConfig
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1beta"
	pb "cloud.google.com/go/discoveryengine/apiv1beta/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineServingConfigGVK, NewServingConfigModel)
}

func NewServingConfigModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &servingConfigModel{config: *config}, nil
}

var _ directbase.Model = &servingConfigModel{}

type servingConfigModel struct {
	config config.ControllerConfig
}

func (m *servingConfigModel) client(ctx context.Context, projectID string) (*gcp.ServingConfigClient, error) {
	var opts []option.ClientOption

	config := m.config

	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewServingConfigRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine serving config client: %w", err)
	}

	return gcpClient, err
}

func (m *servingConfigModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineServingConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.DiscoveryEngineServingConfigIdentity)

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineServingConfigSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	return &servingConfigAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *servingConfigModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		trimmed := strings.TrimPrefix(url, "//discoveryengine.googleapis.com/")
		id := &krm.DiscoveryEngineServingConfigIdentity{}
		if err := id.FromExternal(trimmed); err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineServingConfig format", "url", url)
			return nil, nil
		}
		gcpClient, err := m.client(ctx, id.Project)
		if err != nil {
			return nil, err
		}
		return &servingConfigAdapter{
			gcpClient: gcpClient,
			id:        id,
		}, nil
	}
	return nil, nil
}

type servingConfigAdapter struct {
	gcpClient *gcp.ServingConfigClient
	id        *krm.DiscoveryEngineServingConfigIdentity
	desired   *pb.ServingConfig
	actual    *pb.ServingConfig
}

var _ directbase.Adapter = &servingConfigAdapter{}

func (a *servingConfigAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	fqn := a.id.String()
	log.V(2).Info("getting discoveryengine serving config", "name", fqn)

	req := &pb.GetServingConfigRequest{Name: fqn}
	actual, err := a.gcpClient.GetServingConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine serving config %q from gcp: %w", fqn, err)
	}

	a.actual = actual
	return true, nil
}

func (a *servingConfigAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating (updating pre-existing) discoveryengine serving config", "name", a.id)

	desired := proto.CloneOf(a.desired)
	desired.Name = a.id.String()

	req := &pb.UpdateServingConfigRequest{
		ServingConfig: desired,
	}
	updated, err := a.gcpClient.UpdateServingConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating discoveryengine serving config %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully reconciled discoveryengine serving config in gcp", "name", a.id)

	return a.updateStatus(ctx, createOp, updated)
}

func (a *servingConfigAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine serving config", "name", a.id)

	diffs, updateMask, err := compareServingConfig(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.CloneOf(a.desired)
		desired.Name = a.id.String()

		req := &pb.UpdateServingConfigRequest{
			ServingConfig: desired,
			UpdateMask:    updateMask,
		}
		updated, err := a.gcpClient.UpdateServingConfig(ctx, req)
		if err != nil {
			return fmt.Errorf("updating discoveryengine serving config %s: %w", a.id.String(), err)
		}
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareServingConfig(ctx context.Context, actual, desired *pb.ServingConfig) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DiscoveryEngineServingConfigSpec_v1alpha1_FromProto, DiscoveryEngineServingConfigSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore name if needed

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.ServingConfig) {
		// populate any defaults if necessary
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *servingConfigAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.ServingConfig) error {
	mapCtx := &direct.MapContext{}
	status := &krm.DiscoveryEngineServingConfigStatus{}
	status.ObservedState = DiscoveryEngineServingConfigObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *servingConfigAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineServingConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineServingConfigSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	obj.Spec.Collection = a.id.Collection

	if a.id.Engine != "" {
		obj.Spec.EngineRef = &krm.DiscoveryEngineEngineRef{
			External: fmt.Sprintf("projects/%s/locations/%s/collections/%s/engines/%s", a.id.Project, a.id.Location, a.id.Collection, a.id.Engine),
		}
	}
	if a.id.DataStore != "" {
		obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{
			External: fmt.Sprintf("projects/%s/locations/%s/collections/%s/dataStores/%s", a.id.Project, a.id.Location, a.id.Collection, a.id.DataStore),
		}
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ServingConfig)
	u.SetGroupVersionKind(krm.DiscoveryEngineServingConfigGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *servingConfigAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting (noop) discoveryengine serving config", "name", a.id)
	return true, nil
}
