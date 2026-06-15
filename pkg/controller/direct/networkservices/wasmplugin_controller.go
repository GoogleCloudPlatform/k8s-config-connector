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

package networkservices

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	gcp "cloud.google.com/go/networkservices/apiv1"
	networkservicespb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.NetworkServicesWasmPluginGVK, NewWasmPluginModel)
}

func NewWasmPluginModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelWasmPlugin{config: *config}, nil
}

var _ directbase.Model = &modelWasmPlugin{}

type modelWasmPlugin struct {
	config config.ControllerConfig
}

func (m *modelWasmPlugin) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building WasmPlugin client: %w", err)
	}
	return gcpClient, err
}

func (m *modelWasmPlugin) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	desired := &krm.NetworkServicesWasmPlugin{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &desired); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", desired, err)
	}

	id, err := krm.NewWasmPluginIdentity(ctx, reader, desired)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := NetworkServicesWasmPluginSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desiredProto.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &WasmPluginAdapter{
		id:           id,
		gcpClient:    gcpClient,
		desired:      desired,
		reader:       reader,
		desiredProto: desiredProto,
	}, nil
}

func (m *modelWasmPlugin) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type WasmPluginAdapter struct {
	id           *krm.WasmPluginIdentity
	gcpClient    *gcp.Client
	desired      *krm.NetworkServicesWasmPlugin
	reader       client.Reader
	actual       *networkservicespb.WasmPlugin
	desiredProto *networkservicespb.WasmPlugin
}

var _ directbase.Adapter = &WasmPluginAdapter{}

func (a *WasmPluginAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting WasmPlugin", "name", a.id)

	req := &networkservicespb.GetWasmPluginRequest{Name: a.id.String()}
	wasmpluginpb, err := a.gcpClient.GetWasmPlugin(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting WasmPlugin %q: %w", a.id, err)
	}

	// GCP GetWasmPlugin does not return versions inline. We must list WasmPluginVersion child resources to populate actual.Versions.
	listReq := &networkservicespb.ListWasmPluginVersionsRequest{Parent: a.id.String()}
	versionIterator := a.gcpClient.ListWasmPluginVersions(ctx, listReq)
	if versionIterator != nil {
		for versionpb, err := versionIterator.Next(); err == nil; {
			tokens := strings.Split(versionpb.Name, "/versions/")
			if len(tokens) == 2 {
				versionID := tokens[1]
				if wasmpluginpb.Versions == nil {
					wasmpluginpb.Versions = make(map[string]*networkservicespb.WasmPlugin_VersionDetails)
				}
				vd := &networkservicespb.WasmPlugin_VersionDetails{
					Description: versionpb.Description,
					Labels:      versionpb.Labels,
					ImageUri:    versionpb.ImageUri,
				}
				switch source := versionpb.PluginConfigSource.(type) {
				case *networkservicespb.WasmPluginVersion_PluginConfigData:
					vd.PluginConfigSource = &networkservicespb.WasmPlugin_VersionDetails_PluginConfigData{
						PluginConfigData: source.PluginConfigData,
					}
				case *networkservicespb.WasmPluginVersion_PluginConfigUri:
					vd.PluginConfigSource = &networkservicespb.WasmPlugin_VersionDetails_PluginConfigUri{
						PluginConfigUri: source.PluginConfigUri,
					}
				}
				wasmpluginpb.Versions[versionID] = vd
			}
			versionpb, err = versionIterator.Next()
		}
	}

	a.actual = wasmpluginpb
	return true, nil
}

func (a *WasmPluginAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating WasmPlugin", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := proto.Clone(a.desiredProto).(*networkservicespb.WasmPlugin)
	resource.Name = a.id.String()

	req := &networkservicespb.CreateWasmPluginRequest{
		Parent:       a.id.Parent().String(),
		WasmPluginId: a.id.ID(),
		WasmPlugin:   resource,
	}
	op, err := a.gcpClient.CreateWasmPlugin(ctx, req)
	if err != nil {
		return fmt.Errorf("creating WasmPlugin %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("WasmPlugin %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created WasmPlugin", "name", a.id)

	status := &krm.NetworkServicesWasmPluginStatus{}
	status.ObservedState = NetworkServicesWasmPluginObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *WasmPluginAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating WasmPlugin", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := proto.Clone(a.desiredProto).(*networkservicespb.WasmPlugin)
	resource.Name = a.id.String()

	// Preserve system labels (goog- or go-)
	if a.actual.Labels != nil {
		if resource.Labels == nil {
			resource.Labels = make(map[string]string)
		}
		for k, v := range a.actual.Labels {
			if strings.HasPrefix(k, "goog-") || strings.HasPrefix(k, "go-") {
				resource.Labels[k] = v
			}
		}
	}

	diff, err := common.CompareProtoMessage(a.actual, resource, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing WasmPlugin %s: %w", a.id, err)
	}

	updated := a.actual
	if diff.Len() == 0 {
		log.V(2).Info("no changes detected for WasmPlugin", "name", a.id)
	} else {
		sortedPaths := diff.UnsortedList()
		sort.Strings(sortedPaths)
		updateMask := &fieldmaskpb.FieldMask{Paths: sortedPaths}

		req := &networkservicespb.UpdateWasmPluginRequest{
			UpdateMask: updateMask,
			WasmPlugin: resource,
		}
		op, err := a.gcpClient.UpdateWasmPlugin(ctx, req)
		if err != nil {
			return fmt.Errorf("updating WasmPlugin %s: %w", a.id, err)
		}
		updated, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("WasmPlugin %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated WasmPlugin", "name", a.id)
	}

	status := &krm.NetworkServicesWasmPluginStatus{}
	status.ObservedState = NetworkServicesWasmPluginObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *WasmPluginAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.NetworkServicesWasmPlugin{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(NetworkServicesWasmPluginSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refs.ProjectRef{External: a.id.Parent().ProjectID},
		Location:   a.id.Parent().Location,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkServicesWasmPluginGVK)
	if a.actual.Labels != nil {
		u.SetLabels(a.actual.Labels)
	}

	u.Object = uObj
	return u, nil
}

func (a *WasmPluginAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting WasmPlugin", "name", a.id)

	req := &networkservicespb.DeleteWasmPluginRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteWasmPlugin(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent WasmPlugin, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting WasmPlugin %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted WasmPlugin", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete WasmPlugin %s: %w", a.id, err)
	}
	return true, nil
}
