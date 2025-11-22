// Copyright 2025 Google LLC
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

package configdelivery

import (
	"context"
	"fmt"

	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdelivery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/configdelivery/apiv1"
	configdeliverypb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krmv1alpha1.ConfigDeliveryResourceBundleGVK, NewResourceBundleModel)
}

func NewResourceBundleModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelResourceBundle{config: *config}, nil
}

var _ directbase.Model = &modelResourceBundle{}

type modelResourceBundle struct {
	config config.ControllerConfig
}

func (m *modelResourceBundle) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building ResourceBundle client: %w", err)
	}
	return gcpClient, err
}

func (m *modelResourceBundle) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krmv1alpha1.ConfigDeliveryResourceBundle{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}

	copied := obj.DeepCopy()
	desired := ConfigDeliveryResourceBundleSpec_ToProto(mapCtx, &copied.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ResourceBundleAdapter{
		id:        id.(*krmv1alpha1.ResourceBundleIdentity),
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelResourceBundle) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ResourceBundleAdapter struct {
	id        *krmv1alpha1.ResourceBundleIdentity
	gcpClient *gcp.Client
	desired   *configdeliverypb.ResourceBundle
	actual    *configdeliverypb.ResourceBundle
}

var _ directbase.Adapter = &ResourceBundleAdapter{}

func (a *ResourceBundleAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ResourceBundle", "name", a.id)

	req := &configdeliverypb.GetResourceBundleRequest{Name: a.id.String()}
	resourceBundlepb, err := a.gcpClient.GetResourceBundle(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ResourceBundle %q: %w", a.id, err)
	}

	a.actual = resourceBundlepb
	return true, nil
}

func (a *ResourceBundleAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ResourceBundle", "name", a.id)
	req := &configdeliverypb.CreateResourceBundleRequest{
		Parent:           a.id.Parent().String(),
		ResourceBundle:   a.desired,
		ResourceBundleId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateResourceBundle(ctx, req)
	if err != nil {
		return fmt.Errorf("creating ResourceBundle %s: %w", a.id, err)
	}

	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for create of ResourceBundle %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created ResourceBundle", "name", a.id)

	mapCtx := &direct.MapContext{}
	status := &krmv1alpha1.ConfigDeliveryResourceBundleStatus{}
	status.ObservedState = ConfigDeliveryResourceBundleObservedState_v1alpha1_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ResourceBundleAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ResourceBundle", "name", a.id)

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(a.desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}
	a.desired.Name = a.id.String()

	req := &configdeliverypb.UpdateResourceBundleRequest{
		UpdateMask:     updateMask,
		ResourceBundle: a.desired,
	}
	op, err := a.gcpClient.UpdateResourceBundle(ctx, req)
	if err != nil {
		return fmt.Errorf("updating ResourceBundle %s: %w", a.id, err)
	}

	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for update of ResourceBundle %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated ResourceBundle", "name", a.id)

	mapCtx := &direct.MapContext{}

	status := &krmv1alpha1.ConfigDeliveryResourceBundleStatus{}
	status.ObservedState = ConfigDeliveryResourceBundleObservedState_v1alpha1_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *ResourceBundleAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krmv1alpha1.ConfigDeliveryResourceBundle{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ConfigDeliveryResourceBundleSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	externalRef := a.actual.GetName()
	var id *krmv1alpha1.ResourceBundleIdentity
	if err := id.FromExternal(externalRef); err != nil {
		return nil, fmt.Errorf("parsing external ref %q: %w", externalRef, err)
	}
	obj.Spec.Parent = id.Parent().String()
	obj.Spec.ResourceID = direct.LazyPtr(a.id.ID())
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krmv1alpha1.ConfigDeliveryResourceBundleGVK)

	u.Object = uObj
	return u, nil
}

func (a *ResourceBundleAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ResourceBundle", "name", a.id)

	req := &configdeliverypb.DeleteResourceBundleRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteResourceBundle(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent ResourceBundle, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting ResourceBundle %s: %w", a.id, err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for delete of ResourceBundle %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted ResourceBundle", "name", a.id)

	return true, nil
}
