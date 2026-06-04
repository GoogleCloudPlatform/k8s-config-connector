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

package certificatemanagertrustconfig

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/certificatemanager/apiv1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/certificatemanager/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/certificatemanager"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.CertificateManagerTrustConfigGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TrustConfig client: %w", err)
	}
	return gcpClient, err
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CertificateManagerTrustConfig{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	typedID, ok := id.(*krm.CertificateManagerTrustConfigIdentity)
	if !ok {
		return nil, fmt.Errorf("expected CertificateManagerTrustConfigIdentity, got %T", id)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        typedID,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.CertificateManagerTrustConfigIdentity
	gcpClient *gcp.Client
	desired   *krm.CertificateManagerTrustConfig
	actual    *pb.TrustConfig
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TrustConfig", "name", a.id.String())

	req := &pb.GetTrustConfigRequest{Name: a.id.String()}
	trustconfigpb, err := a.gcpClient.GetTrustConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting TrustConfig %q: %w", a.id.String(), err)
	}

	a.actual = trustconfigpb
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating TrustConfig")
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := certificatemanager.CertificateManagerTrustConfigSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := fmt.Sprintf("projects/%s/locations/%s", a.id.Project, a.id.Location)
	resource.Name = a.id.String()

	req := &pb.CreateTrustConfigRequest{
		Parent:        parent,
		TrustConfigId: a.id.TrustConfig,
		TrustConfig:   resource,
	}

	op, err := a.gcpClient.CreateTrustConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TrustConfig %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("TrustConfig %s waiting creation: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created TrustConfig", "name", a.id.String())

	status := &krm.CertificateManagerTrustConfigStatus{}
	status.ObservedState = certificatemanager.CertificateManagerTrustConfigObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := created.Name
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating TrustConfig", "name", a.id.String())
	mapCtx := &direct.MapContext{}

	resource := certificatemanager.CertificateManagerTrustConfigSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	// Preserve fields that are read-only or not in spec but needed for update
	resource.Etag = a.actual.Etag

	updateMask := &fieldmaskpb.FieldMask{}

	// Compare fields to build update mask
	if !reflect.DeepEqual(a.desired.Spec.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}
	if !reflect.DeepEqual(resource.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(resource.TrustStores, a.actual.TrustStores) {
		updateMask.Paths = append(updateMask.Paths, "trust_stores")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.CertificateManagerTrustConfigStatus{}
		status.ObservedState = certificatemanager.CertificateManagerTrustConfigObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		externalRef := a.actual.Name
		status.ExternalRef = &externalRef
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	req := &pb.UpdateTrustConfigRequest{
		TrustConfig: resource,
		UpdateMask:  updateMask,
	}

	op, err := a.gcpClient.UpdateTrustConfig(ctx, req)
	if err != nil {
		return fmt.Errorf("updating TrustConfig %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("TrustConfig %s waiting update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated TrustConfig", "name", a.id.String())

	status := &krm.CertificateManagerTrustConfigStatus{}
	status.ObservedState = certificatemanager.CertificateManagerTrustConfigObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := updated.Name
	status.ExternalRef = &externalRef
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.CertificateManagerTrustConfig{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(certificatemanager.CertificateManagerTrustConfigSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TrustConfig", "name", a.id.String())

	req := &pb.DeleteTrustConfigRequest{
		Name: a.id.String(),
	}

	op, err := a.gcpClient.DeleteTrustConfig(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting TrustConfig %s: %w", a.id.String(), err)
	}
	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("TrustConfig %s waiting deletion: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted TrustConfig", "name", a.id.String())

	return true, nil
}
