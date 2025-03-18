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

// +tool:controller
// proto.service: google.cloud.aiplatform.v1beta1.NotebookService
// proto.message: google.cloud.aiplatform.v1beta1.NotebookRuntimeTemplate
// crd.type: ColabRuntimeTemplate
// crd.version: v1alpha1

package colab

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/colab/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.ColabRuntimeTemplateGVK, NewRuntimeTemplateModel)
}

func NewRuntimeTemplateModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &runtimeTemplateModel{config: *config}, nil
}

var _ directbase.Model = &runtimeTemplateModel{}

type runtimeTemplateModel struct {
	config config.ControllerConfig
}

func (m *runtimeTemplateModel) client(ctx context.Context, projectID string) (*gcp.NotebookClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewNotebookRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building aiplatform colabruntimetemplate client: %w", err)
	}

	return gcpClient, err
}

func (m *runtimeTemplateModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ColabRuntimeTemplate{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewNotebookRuntimeTemplateIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := ColabRuntimeTemplateSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &runtimeTemplateAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *runtimeTemplateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		id, err := krm.ParseNotebookRuntimeTemplateExternal(url)
		if err != nil {
			log.V(2).Error(err, "url did not match ColabRuntimeTemplate format", "url", url)
		} else {
			gcpClient, err := m.client(ctx, id.Parent().ProjectID)
			if err != nil {
				return nil, err
			}
			return &runtimeTemplateAdapter{
				gcpClient: gcpClient,
				id:        id,
			}, nil
		}
	}
	return nil, nil
}

type runtimeTemplateAdapter struct {
	gcpClient *gcp.NotebookClient
	id        *krm.NotebookRuntimeTemplateIdentity
	desired   *pb.NotebookRuntimeTemplate
	actual    *pb.NotebookRuntimeTemplate
}

var _ directbase.Adapter = &runtimeTemplateAdapter{}

func (a *runtimeTemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting aiplatform colabruntimetemplate", "name", a.id)

	req := &pb.GetNotebookRuntimeTemplateRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetNotebookRuntimeTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting aiplatform colabruntimetemplate %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *runtimeTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating aiplatform colabruntimetemplate", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	req := &pb.CreateNotebookRuntimeTemplateRequest{
		Parent:                    a.id.Parent().String(),
		NotebookRuntimeTemplate:   desired,
		NotebookRuntimeTemplateId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateNotebookRuntimeTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating aiplatform colabruntimetemplate %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("aiplatform colabruntimetemplate %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created aiplatform colabruntimetemplate in gcp", "name", a.id)

	status := &krm.ColabRuntimeTemplateStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ColabRuntimeTemplateObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *runtimeTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating aiplatform colabruntimetemplate", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	// TODO(user): Update the field if applicable.
	updateMask := &fieldmaskpb.FieldMask{}
	if !reflect.DeepEqual(a.desired.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(a.desired.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if !reflect.DeepEqual(a.desired.Etag, a.actual.Etag) {
		updateMask.Paths = append(updateMask.Paths, "etag")
	}
	if !reflect.DeepEqual(a.desired.Labels, a.actual.Labels) {
		updateMask.Paths = append(updateMask.Paths, "labels")
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}
	req := &pb.UpdateNotebookRuntimeTemplateRequest{
		UpdateMask:              updateMask,
		NotebookRuntimeTemplate: desired,
	}
	updated, err := a.gcpClient.UpdateNotebookRuntimeTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("updating aiplatform colabruntimetemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated aiplatform colabruntimetemplate", "name", a.id)

	status := &krm.ColabRuntimeTemplateStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = ColabRuntimeTemplateObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *runtimeTemplateAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.ColabRuntimeTemplate{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ColabRuntimeTemplateSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.ColabRuntimeTemplateGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *runtimeTemplateAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting aiplatform colabruntimetemplate", "name", a.id)

	req := &pb.DeleteNotebookRuntimeTemplateRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteNotebookRuntimeTemplate(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting aiplatform colabruntimetemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted aiplatform colabruntimetemplate", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of aiplatform colabruntimetemplate %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
