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

package kmsautokeyconfig

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(user): Update the import with the google cloud client
	gcp "cloud.google.com/go/kms/apiv1"

	// TODO(user): Update the import with the google cloud client api protobuf
	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName = "kms-autokeyconfig-controller"
)

func init() {
	registry.RegisterModel(krm.KMSAutokeyConfigGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.AutokeyAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAutokeyAdminRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building AutokeyConfig client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.KMSAutokeyConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAutokeyConfigIdentity(ctx, reader, obj)
	if err != nil {
		return nil, fmt.Errorf("unable to resolve folder for autokeyConfig name: %s, err: %w", obj.GetName(), err)
	}
	var keyProject *refs.Project
	if obj.Spec.KeyProjectRef != nil {
		var err error
		keyProject, err = refs.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.KeyProjectRef)
		if err != nil {
			return nil, fmt.Errorf("unable to resolve key project for autokeyConfig naem: %s, err: %w", obj.GetName(), err)
		}
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:                id,
		desiredKeyProject: keyProject,
		gcpClient:         gcpClient,
		desired:           obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id                *krm.KMSAutokeyConfigIdentity
	desiredKeyProject *refs.Project
	gcpClient         *gcp.AutokeyAdminClient
	desired           *krm.KMSAutokeyConfig
	actual            *kmspb.AutokeyConfig
}

var _ directbase.Adapter = &Adapter{}

// Find return true if AutokeyConfig exist and user has permission to read it.
// Else it will return false and error.
func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting KMSAutokeyConfig", "name", a.id)

	req := &kmspb.GetAutokeyConfigRequest{Name: a.id.String()}
	autokeyconfigpb, err := a.gcpClient.GetAutokeyConfig(ctx, req)
	if err != nil {
		return false, fmt.Errorf("getting KMSAutokeyConfig %q: %w", a.id, err)
	}

	a.actual = autokeyconfigpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("Create operation not supported for AutokeyConfig resource.")
	return fmt.Errorf("Create operation not supported for AutokeyConfig resource")
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating AutokeyConfig", "name", a.id)
	mapCtx := &direct.MapContext{}

	resource := KMSAutokeyConfig_FromFields(mapCtx, a.id, a.desiredKeyProject)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updated, err := a.updateAutokeyConfig(ctx, resource)
	if err != nil {
		return err
	}

	status := &krm.KMSAutokeyConfigStatus{}
	status.ObservedState = KMSAutokeyConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) updateAutokeyConfig(ctx context.Context, resource *kmspb.AutokeyConfig) (*kmspb.AutokeyConfig, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	// To populate a.actual calling a.Find()
	isExist, err := a.Find(ctx)
	if !isExist {
		return nil, fmt.Errorf("updateAutokeyConfig failed as AutokeyConfig does not exist, name: %s", a.id)
	}
	if err != nil {
		return nil, err
	}
	updateMask := &fieldmaskpb.FieldMask{}
	if resource.KeyProject != "" && !reflect.DeepEqual(resource.KeyProject, a.actual.KeyProject) {
		updateMask.Paths = append(updateMask.Paths, "key_project")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil, nil
	}
	req := &kmspb.UpdateAutokeyConfigRequest{
		UpdateMask:    updateMask,
		AutokeyConfig: resource,
	}
	updated, err := a.gcpClient.UpdateAutokeyConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("updating AutokeyConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated AutokeyConfig", "name", a.id)
	return updated, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.KMSAutokeyConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(KMSAutokeyConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	parent := a.id.Parent()
	obj.Spec.FolderRef = &refs.FolderRef{External: parent.FolderID}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
// Note: Delete operation is not supported for GCP AutokeyConfig resource.
// However in KCC, the user has full flexibility to delete the KCC AutokeyConfig resource.
// To make this KCC operation effective, as part of KCC AutokeyConfig deletion we will update the AutokeyConfig resource in GCP with empty key_project which will prevent further use of AutokeyConfig.
// Because of the above decision we will update the observedstate for AutokeyConfig with state = UNINITIALIZED
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting AutokeyConfig", "name", a.id)
	mapCtx := &direct.MapContext{}
	// make a copy of the a.actual i.e. from krm.AutokeyConfig to kmspb.AutokeyConfig
	tempKrmAutokeyResource := AutokeyConfig_FromProto(mapCtx, a.actual)
	resource := AutokeyConfig_ToProto(mapCtx, tempKrmAutokeyResource)
	updated, err := a.updateAutokeyConfig(ctx, resource)
	if err != nil {
		return false, fmt.Errorf("updating AutokeyConfig %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted AutokeyConfig in KCC by resetting the key_project", "name", a.id)
	status := &krm.KMSAutokeyConfigStatus{}
	// The state in ObservedState is expected to be UNINITIALIZED as we have set the key_project to empty
	status.ObservedState = KMSAutokeyConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return false, mapCtx.Err()
	}
	// TODO: uncomment once we found a valid solution
	//deleteOp.UpdateStatus(ctx, status, nil)
	return true, nil
}
