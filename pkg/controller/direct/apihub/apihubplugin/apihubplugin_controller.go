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

package apihubplugin

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/apihub/apiv1"
	pb "cloud.google.com/go/apihub/apiv1/apihubpb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apihub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/apihub"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

const (
	ctrlName = "apihub-plugin-controller"
)

func init() {
	registry.RegisterModel(krm.APIHubPluginGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.ApiHubPluginClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewApiHubPluginRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building APIHubPlugin client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.APIHubPlugin{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.APIHubPluginIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.APIHubPluginIdentity
	gcpClient *gcp.ApiHubPluginClient
	desired   *krm.APIHubPlugin
	actual    *pb.Plugin
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting APIHubPlugin", "name", a.id.String())

	req := &pb.GetPluginRequest{Name: a.id.String()}
	pbObj, err := a.gcpClient.GetPlugin(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting APIHubPlugin %q: %w", a.id.String(), err)
	}

	a.actual = pbObj
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating APIHubPlugin")
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := apihub.APIHubPluginSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	// Parent format: projects/{project}/locations/{location}
	parentString := a.id.ParentString()

	req := &pb.CreatePluginRequest{
		Parent:   parentString,
		PluginId: a.id.Plugin,
		Plugin:   resource,
	}

	created, err := a.gcpClient.CreatePlugin(ctx, req)
	if err != nil {
		return fmt.Errorf("creating APIHubPlugin %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created APIHubPlugin", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating APIHubPlugin", "name", a.id.String())

	diffs, err := comparePlugin(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		return fmt.Errorf("APIHubPlugin is immutable and cannot be updated. Detected changes: %v", diffs.Fields)
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func comparePlugin(ctx context.Context, actual *pb.Plugin, desiredKRM *krm.APIHubPlugin) (*structuredreporting.Diff, error) {
	mapCtx := &direct.MapContext{}
	desired := apihub.APIHubPluginSpec_ToProto(mapCtx, &desiredKRM.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	maskedActual, err := mappers.OnlySpecFields(actual, apihub.APIHubPluginSpec_FromProto, apihub.APIHubPluginSpec_ToProto)
	if err != nil {
		return nil, err
	}
	diffs, _, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, err
	}
	return diffs, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.APIHubPlugin{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(apihub.APIHubPluginSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.Location = &a.id.Location
	obj.Spec.ProjectRef = &refs.ProjectRef{External: fmt.Sprintf("projects/%s", a.id.Project)}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting APIHubPlugin", "name", a.id.String())

	req := &pb.DeletePluginRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeletePlugin(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting APIHubPlugin %s: %w", a.id.String(), err)
	}

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting for APIHubPlugin %s deletion: %w", a.id.String(), err)
	}

	return true, nil
}

func (a *Adapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Plugin) error {
	mapCtx := &direct.MapContext{}
	status := &krm.APIHubPluginStatus{}
	status.ObservedState = apihub.APIHubPluginObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}
