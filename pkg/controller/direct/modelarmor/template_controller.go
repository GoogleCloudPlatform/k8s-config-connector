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

package modelarmor

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/modelarmor/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/modelarmor/apiv1"
	pb "cloud.google.com/go/modelarmor/apiv1/modelarmorpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.ModelArmorTemplateGVK, NewTemplateModel)
}

func NewTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTemplate{config: *config}, nil
}

var _ directbase.Model = &modelTemplate{}

type modelTemplate struct {
	config config.ControllerConfig
}

func (m *modelTemplate) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Template client: %w", err)
	}
	return gcpClient, err
}

func (m *modelTemplate) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ModelArmorTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTemplateIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get modelarmor GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &TemplateAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelTemplate) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type TemplateAdapter struct {
	id        *krm.TemplateIdentity
	gcpClient *gcp.Client
	desired   *krm.ModelArmorTemplate
	actual    *pb.Template
	reader    client.Reader
}

var _ directbase.Adapter = &TemplateAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Template", "name", a.id)

	req := &pb.GetTemplateRequest{Name: a.id.String()}
	templatepb, err := a.gcpClient.GetTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Template %q: %w", a.id, err)
	}

	a.actual = templatepb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Template", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.desired.DeepCopy()
	resource := ModelArmorTemplateSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateTemplateRequest{
		Parent:     a.id.Parent().String(),
		TemplateId: a.id.ID(),
		Template:   resource,
	}
	created, err := a.gcpClient.CreateTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Template %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created Template", "name", a.id)

	status := &krm.ModelArmorTemplateStatus{}
	status.ObservedState = ModelArmorTemplateObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Template", "name", a.id)

	mapCtx := &direct.MapContext{}
	desiredPb := ModelArmorTemplateSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Set the name field to ensure the GCP API can identity the resource during UpdateTemplate().
	desiredPb.Name = a.id.String()

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.ModelArmorTemplateStatus{}
		status.ObservedState = ModelArmorTemplateObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateTemplateRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: sets.List(paths)},
		Template: desiredPb,
	}
	updated, err := a.gcpClient.UpdateTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Template %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated Template", "name", a.id.String())

	status := &krm.ModelArmorTemplateStatus{}
	status.ExternalRef = direct.LazyPtr(updated.Name)
	status.ObservedState = ModelArmorTemplateObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ModelArmorTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ModelArmorTemplateSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.ModelArmorTemplateGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Template", "name", a.id)

	req := &pb.DeleteTemplateRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteTemplate(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Template %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Template", "name", a.id)

	return true, nil
}
