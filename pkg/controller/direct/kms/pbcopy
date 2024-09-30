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
	//"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
//	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	folderref "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/kms/folderref"
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
	ctrlName = "kms-controller"
	// TODO(user): Confirm service domain
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

	// Get ResourceID
	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	//TODO add folder ref to kmsautokeyconfig
	folderRef, err := folderref.ResolveFolder(ctx, reader, obj, obj.Spec.FolderRef)
	if err != nil {
		return nil, err
	}
	folderID := folderRef.FolderID
	if folderID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}
	/*projectRef := &refs.ProjectRef{
		Name:      obj.Spec.ProjectRef.Name,
		Namespace: obj.Spec.ProjectRef.Namespace,
		External:  obj.Spec.ProjectRef.External,
	}
	project, err := refs.ResolveProject(ctx, reader, obj, projectRef)
	if err != nil {
		return nil, err
	}*/
	var id *KMSAutokeyConfigIdentity
	externalRef := direct.ValueOf(obj.Status.ExternalRef)
	if externalRef == "" {
		id = BuildID(folderID)
	} else {
		id, err = asID(externalRef)
		if err != nil {
			return nil, err
		}

		if id.Parent.Folder != folderID {
			return nil, fmt.Errorf("KMSAutokeyConfig %s/%s has spec.folderRef changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.Parent.Folder, folderID)
		}
		/*if id.AutokeyConfig != resourceID {
			return nil, fmt.Errorf("KMSAutokeyConfig  %s/%s has metadata.name or spec.resourceID changed, expect %s, got %s",
				u.GetNamespace(), u.GetName(), id.AutokeyConfig, resourceID)
		}*/
	}

	// TODO(kcc): GetGCPClient as interface method.
	// Get kms GCP client
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

	// TODO(user): write the gcp "GET" operation.
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
	/*
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating AutokeyConfig", "name", a.id.FullyQualifiedName())
	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := KMSAutokeyConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(user): Complete the gcp "CREATE" or "INSERT" request with required fields.
	req := &kmspb.CreateAutokeyConfigRequest{
		Parent:        a.id.Parent.String(),
		AutokeyConfig: resource,
	}
	op, err := a.gcpClient.CreateAutokeyConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating AutokeyConfig %s: %w", a.id.FullyQualifiedName(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("AutokeyConfig %s waiting creation: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created AutokeyConfig", "name", a.id.FullyQualifiedName())
	status := &krm.KMSAutokeyConfigStatus{}
	status.ObservedState = KMSAutokeyConfigStatusObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = a.id.AsExternalRef()
	log.V(2).Info("no-op, AutokeyConfig is already created")
	*/
	return setStatus(nil, nil)
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

	// TODO(user): Update the field if applicable.
	updateMask := &fieldmaskpb.FieldMask{}
	/*if projectID != "" && !reflect.DeepEqual(projectID, a.actual.KeyProject) {
		updateMask.Paths = append(updateMask.Paths, "KeyProject")
	}*/

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.FullyQualifiedName())
		return nil
	}
	// TODO(user): Complete the gcp "UPDATE" or "PATCH" request with required fields.
	req := &kmspb.UpdateAutokeyConfigRequest{
		//Name:          a.id.FullyQualifiedName(),
		UpdateMask:    updateMask,
		AutokeyConfig: resource,
	}
	updated, err := a.gcpClient.UpdateAutokeyConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating AutokeyConfig %s: %w", a.id.FullyQualifiedName(), err)
	}
	log.V(2).Info("successfully updated AutokeyConfig", "name", a.id.FullyQualifiedName())

	status := &krm.KMSAutokeyConfigStatus{}
	status.ObservedState = KMSAutokeyConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return setStatus(u, status)
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
	// TODO(user): Update other resource reference
	obj.Spec.FolderRef = &folderref.FolderRef{Name: a.id.Parent.Folder}
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

	/*req := &kmspb.DeleteAutokeyConfigRequest{Name: a.id.FullyQualifiedName()}
	op, err := a.gcpClient.DeleteAutokeyConfig(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting AutokeyConfig %s: %w", a.id.FullyQualifiedName(), err)
	}
	*/
	log.V(2).Info("no-op, cannot deleted AutokeyConfig", "name", a.id.FullyQualifiedName())

	/*err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete AutokeyConfig %s: %w", a.id.FullyQualifiedName(), err)
	}*/
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
