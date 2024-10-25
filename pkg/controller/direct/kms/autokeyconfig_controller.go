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

package kms

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/kms/apiv1"

	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ctrlName      = "kms-controller"
	serviceDomain = "//cloudkms.googleapis.com"
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

	var folder *refs.Folder
	var err error
	folder, err = refs.ResolveFolderFromAnnotation(ctx, reader, obj)
	if err != nil {
		folder, err = refs.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
	}
	if err != nil || folder.FolderID == "" {
		return nil, fmt.Errorf("unable to resolve folder for autokeyConfig name: %s", obj.GetName())
	}
	var id *KMSAutokeyConfigIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(folder.FolderID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.FolderID != folder.FolderID {
			return nil, fmt.Errorf("KMSAutokeyConfig %s/%s has spec.folderRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.FolderID, folder.FolderID)
		}
	}

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
	// TODO: Support URLs
	return nil, nil
}

type Adapter struct {
	id        *KMSAutokeyConfigIdentity
	gcpClient *gcp.AutokeyAdminClient
	desired   *krm.KMSAutokeyConfig
	actual    *kmspb.AutokeyConfig
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting KMSAutokeyConfig", "name", a.id.FullyQualifiedName())
	req := &kmspb.GetAutokeyConfigRequest{Name: a.id.FullyQualifiedName()}
	autokeyconfigpb, err := a.gcpClient.GetAutokeyConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting KMSAutokeyConfig %q: %w", a.id.FullyQualifiedName(), err)
	}

	a.actual = autokeyconfigpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating AutokeyConfig", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := KMSAutokeyConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	updated, err := a.updateAutokeyConfig(ctx, resource)
	if err != nil {
		return err
	}
	status := &krm.KMSAutokeyConfigStatus{}
	status.ObservedState = KMSAutokeyConfigStatusObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	return setStatus(u, status)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating AutokeyConfig", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := KMSAutokeyConfigSpec_ToProto(mapCtx, &desired.Spec)
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
	return setStatus(u, status)
}

func (a *Adapter) updateAutokeyConfig(ctx context.Context, resource *kmspb.AutokeyConfig) (*kmspb.AutokeyConfig, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	// To populate a.actual calling a.Find()
	isExist, err := a.Find(ctx)
	if !isExist {
		return nil, fmt.Errorf("updateAutokeyConfig failed as AutokeyConfig does not exist, name: %s", a.id.FullyQualifiedName())
	}
	if err != nil {
		return nil, err
	}
	updateMask := &fieldmaskpb.FieldMask{}
	if resource.KeyProject != "" && !reflect.DeepEqual(resource.KeyProject, a.actual.KeyProject) {
		updateMask.Paths = append(updateMask.Paths, "key_project")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.FullyQualifiedName())
		return nil, nil
	}
	req := &kmspb.UpdateAutokeyConfigRequest{
		UpdateMask:    updateMask,
		AutokeyConfig: resource,
	}
	updated, err := a.gcpClient.UpdateAutokeyConfig(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("updating AutokeyConfig %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated AutokeyConfig", "name", a.id.FullyQualifiedName())
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
	obj.Spec.FolderRef = &refs.FolderRef{External: a.id.Parent.String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting AutokeyConfig", "name", a.id.FullyQualifiedName())

	log.V(2).Info("no-op, cannot deleted AutokeyConfig", "name", a.id.FullyQualifiedName())

	return true, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
		status["externalRef"] = old["externalRef"]
	}

	u.Object["status"] = status

	return nil
}
