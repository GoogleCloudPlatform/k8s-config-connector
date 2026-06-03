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

package wasmplugin

import (
	"context"
	"fmt"
	"sort"
	"strings"

	parent "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/networkservices/apiv1"
	networkservicespb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	proto "google.golang.org/protobuf/proto"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
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
	obj := &krm.NetworkServicesWasmPlugin{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewWasmPluginIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get networkservices GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &WasmPluginAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelWasmPlugin) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type WasmPluginAdapter struct {
	id        *krm.WasmPluginIdentity
	gcpClient *gcp.Client
	desired   *krm.NetworkServicesWasmPlugin
	actual    *networkservicespb.WasmPlugin
}

var _ directbase.Adapter = &WasmPluginAdapter{}

// Find retrieves the GCP resource.
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

	// List versions for the WasmPlugin to populate actual.Versions
	var versions map[string]*networkservicespb.WasmPlugin_VersionDetails
	reqVersions := &networkservicespb.ListWasmPluginVersionsRequest{
		Parent: a.id.String(),
	}
	it := a.gcpClient.ListWasmPluginVersions(ctx, reqVersions)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return false, fmt.Errorf("listing WasmPluginVersions for %q: %w", a.id, err)
		}
		if versions == nil {
			versions = make(map[string]*networkservicespb.WasmPlugin_VersionDetails)
		}
		tokens := strings.Split(resp.GetName(), "/versions/")
		versionID := tokens[len(tokens)-1]

		details := &networkservicespb.WasmPlugin_VersionDetails{
			Description: resp.GetDescription(),
			Labels:      resp.GetLabels(),
			ImageUri:    resp.GetImageUri(),
			CreateTime:  resp.GetCreateTime(),
			UpdateTime:  resp.GetUpdateTime(),
		}
		if data := resp.GetPluginConfigData(); data != nil {
			details.PluginConfigSource = &networkservicespb.WasmPlugin_VersionDetails_PluginConfigData{PluginConfigData: data}
		} else if uri := resp.GetPluginConfigUri(); uri != "" {
			details.PluginConfigSource = &networkservicespb.WasmPlugin_VersionDetails_PluginConfigUri{PluginConfigUri: uri}
		}
		versions[versionID] = details
	}
	wasmpluginpb.Versions = versions

	a.actual = wasmpluginpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WasmPluginAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating WasmPlugin", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkServicesWasmPluginSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

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

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *WasmPluginAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating WasmPlugin", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := NetworkServicesWasmPluginSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()
	resource.Labels = make(map[string]string)
	for k, v := range a.desired.GetObjectMeta().GetLabels() {
		resource.Labels[k] = v
	}
	resource.Labels["managed-by-cnrm"] = "true"

	// Compare with actual state to only update changed fields
	// Clear server-generated fields from actual before comparing to avoid false diffs
	actualClone := proto.Clone(a.actual).(*networkservicespb.WasmPlugin)
	for _, v := range actualClone.GetVersions() {
		v.CreateTime = nil
		v.UpdateTime = nil
		v.ImageDigest = ""
		v.PluginConfigDigest = ""
	}

	diff, err := common.CompareProtoMessage(actualClone, resource, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing WasmPlugin %s: %w", a.id, err)
	}

	// Double-check map of versions to avoid false positive diffs due to map of interface values
	actualSpec := NetworkServicesWasmPluginSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	if versionsEqual(desired.Spec.Versions, actualSpec.Versions) {
		diff.Delete("versions")
	}

	updated := a.actual
	if diff.Len() == 0 {
		log.V(2).Info("no changes detected for WasmPlugin", "name", a.id)
	} else {
		topLevelFieldPaths := sets.New[string]()
		for path := range diff {
			tokens := strings.Split(path, ".")
			topLevelFieldPaths.Insert(tokens[0])
		}
		sortedPaths := sets.List(topLevelFieldPaths)
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
			return fmt.Errorf("waiting update WasmPlugin %s: %w", a.id, err)
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

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *WasmPluginAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.NetworkServicesWasmPlugin{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkServicesWasmPluginSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectAndLocationRef = &parent.ProjectAndLocationRef{
		ProjectRef: &refs.ProjectRef{External: a.id.Parent().ProjectID},
		Location:   a.id.Parent().Location,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.NetworkServicesWasmPluginGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *WasmPluginAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting WasmPlugin", "name", a.id)

	req := &networkservicespb.DeleteWasmPluginRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteWasmPlugin(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
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

func versionsEqual(a, b map[string]krm.WasmPlugin_VersionDetails) bool {
	if len(a) != len(b) {
		return false
	}
	for k, va := range a {
		vb, ok := b[k]
		if !ok {
			return false
		}
		if direct.ValueOf(va.Description) != direct.ValueOf(vb.Description) {
			return false
		}
		if direct.ValueOf(va.ImageURI) != direct.ValueOf(vb.ImageURI) {
			return false
		}
		if direct.ValueOf(va.PluginConfigData) != direct.ValueOf(vb.PluginConfigData) {
			return false
		}
		if direct.ValueOf(va.PluginConfigURI) != direct.ValueOf(vb.PluginConfigURI) {
			return false
		}
		if len(va.Labels) != len(vb.Labels) {
			return false
		}
		for lk, lva := range va.Labels {
			lvb, lok := vb.Labels[lk]
			if !lok || lva != lvb {
				return false
			}
		}
	}
	return true
}
