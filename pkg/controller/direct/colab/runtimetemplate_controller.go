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
	"strings"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/colab/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
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

func (m *runtimeTemplateModel) client(ctx context.Context, projectID, location string) (*gcp.NotebookClient, error) {
	var opts []option.ClientOption
	config := m.config
	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewNotebookRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building colabruntimetemplate client: %w", err)
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

	gcpClient, err := m.client(ctx, id.Parent().ProjectID, id.Parent().Location)
	if err != nil {
		return nil, err
	}

	return &runtimeTemplateAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *runtimeTemplateModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//aiplatform.googleapis.com/") {
		id, err := krm.ParseNotebookRuntimeTemplateExternal(url)
		if err != nil {
			log.V(2).Error(err, "url did not match ColabRuntimeTemplate format", "url", url)
		} else {
			gcpClient, err := m.client(ctx, id.Parent().ProjectID, id.Parent().Location)
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
	desired   *krm.ColabRuntimeTemplate
	actual    *pb.NotebookRuntimeTemplate

	reader client.Reader
}

var _ directbase.Adapter = &runtimeTemplateAdapter{}

func (a *runtimeTemplateAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting colabruntimetemplate", "name", a.id)

	req := &pb.GetNotebookRuntimeTemplateRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetNotebookRuntimeTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting colabruntimetemplate %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *runtimeTemplateAdapter) normalizeReferences(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.NetworkSpec != nil {
		if obj.Spec.NetworkSpec.NetworkRef != nil {
			if err := obj.Spec.NetworkSpec.NetworkRef.Normalize(ctx, a.reader, obj); err != nil {
				return err
			}
		}
		if obj.Spec.NetworkSpec.SubnetworkRef != nil {
			subnetworkRef, err := refs.ResolveComputeSubnetwork(ctx, a.reader, obj, obj.Spec.NetworkSpec.SubnetworkRef)
			if err != nil {
				return err
			}
			obj.Spec.NetworkSpec.SubnetworkRef = subnetworkRef
		}
	}
	if obj.Spec.ServiceAccountRef != nil {
		if err := obj.Spec.ServiceAccountRef.Resolve(ctx, a.reader, obj); err != nil {
			return err
		}
	}
	if obj.Spec.EncryptionSpec != nil && obj.Spec.EncryptionSpec.KMSKeyRef != nil {
		ref := obj.Spec.EncryptionSpec.KMSKeyRef
		_, err := ref.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *runtimeTemplateAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating colabruntimetemplate", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("resolving references: %w", err)
	}

	desiredPb := ColabRuntimeTemplateSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateNotebookRuntimeTemplateRequest{
		Parent:                    a.id.Parent().String(),
		NotebookRuntimeTemplate:   desiredPb,
		NotebookRuntimeTemplateId: a.id.ID(),
	}
	op, err := a.gcpClient.CreateNotebookRuntimeTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("creating colabruntimetemplate %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("colabruntimetemplate %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created colabruntimetemplate in gcp", "name", a.id)

	status := &krm.ColabRuntimeTemplateStatus{}
	status.ObservedState = ColabRuntimeTemplateObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *runtimeTemplateAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating colabruntimetemplate", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("resolving references: %w", err)
	}

	desiredPb := ColabRuntimeTemplateSpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	// Set default fields.
	if desiredPb.NotebookRuntimeType == pb.NotebookRuntimeType_NOTEBOOK_RUNTIME_TYPE_UNSPECIFIED {
		desiredPb.NotebookRuntimeType = a.actual.NotebookRuntimeType
	}
	if desiredPb.IdleShutdownConfig == nil {
		desiredPb.IdleShutdownConfig = a.actual.IdleShutdownConfig
	}
	if desiredPb.DataPersistentDiskSpec == nil {
		desiredPb.DataPersistentDiskSpec = a.actual.DataPersistentDiskSpec
	}

	paths := make(sets.Set[string])
	var err error
	paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	// Remove output-only fields.
	// The returned "name" is in the format of "projects/{{projectNumber}}/locations/{{location}}/notebookRuntimeTemplates/{{notebookruntimetemplateID}}",
	// so there is always this unexpected diff.
	paths.Delete("name", "etag")
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// Still update the status to cover the use case of acquisition.
		if a.desired.Status.ExternalRef == nil {
			observedState := ColabRuntimeTemplateObservedState_FromProto(mapCtx, a.actual)
			if mapCtx.Err() != nil {
				return mapCtx.Err()
			}

			a.desired.Status.ExternalRef = direct.PtrTo(a.id.String())
			a.desired.Status.ObservedState = observedState

			return updateOp.UpdateStatus(ctx, a.desired.Status, nil)
		}
	}
	updateMask := &fieldmaskpb.FieldMask{
		Paths: sets.List(paths),
	}
	req := &pb.UpdateNotebookRuntimeTemplateRequest{
		UpdateMask:              updateMask,
		NotebookRuntimeTemplate: desiredPb,
	}
	// Currently, the only allowed update mask path is
	// "encryption_spec.kms_key_name".
	updated, err := a.gcpClient.UpdateNotebookRuntimeTemplate(ctx, req)
	if err != nil {
		return fmt.Errorf("updating colabruntimetemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated colabruntimetemplate", "name", a.id)

	status := &krm.ColabRuntimeTemplateStatus{}
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
	log.V(2).Info("deleting colabruntimetemplate", "name", a.id)

	req := &pb.DeleteNotebookRuntimeTemplateRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteNotebookRuntimeTemplate(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting colabruntimetemplate %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted colabruntimetemplate", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of colabruntimetemplate %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
