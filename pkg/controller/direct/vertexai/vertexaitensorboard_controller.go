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

package vertexai

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/aiplatform/apiv1beta1"
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.VertexAITensorboardGVK, NewTensorboardModel)
}

func NewTensorboardModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTensorboard{config: *config}, nil
}

var _ directbase.Model = &modelTensorboard{}

type modelTensorboard struct {
	config config.ControllerConfig
}

func (m *modelTensorboard) client(ctx context.Context, location string) (*gcp.TensorboardClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-aiplatform.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewTensorboardRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("error building Tensorboard client: %w", err)
	}
	return gcpClient, err
}

func (m *modelTensorboard) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.VertexAITensorboard{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idRaw, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idRaw.(*krm.VertexAITensorboardIdentity)

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	// Resolve KMS Key Ref if any
	if obj.Spec.EncryptionSpec != nil && obj.Spec.EncryptionSpec.KMSKeyRef != nil {
		resolvedKMSKey, err := refs.ResolveKMSCryptoKeyRef(ctx, reader, obj, obj.Spec.EncryptionSpec.KMSKeyRef)
		if err != nil {
			return nil, fmt.Errorf("resolving encryptionSpec.kmsKeyRef: %w", err)
		}
		obj.Spec.EncryptionSpec.KMSKeyRef = resolvedKMSKey
	}

	// Get vertexai GCP client
	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := VertexAITensorboardSpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &TensorboardAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *modelTensorboard) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.VertexAITensorboardIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx, id.Location)
	if err != nil {
		return nil, err
	}

	return &TensorboardAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type TensorboardAdapter struct {
	id        *krm.VertexAITensorboardIdentity
	gcpClient *gcp.TensorboardClient
	desired   *pb.Tensorboard
	actual    *pb.Tensorboard
	reader    client.Reader
}

var _ directbase.Adapter = &TensorboardAdapter{}

func (a *TensorboardAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting VertexAITensorboard", "name", a.id)

	req := &pb.GetTensorboardRequest{Name: a.id.String()}
	resp, err := a.gcpClient.GetTensorboard(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting VertexAITensorboard %q: %w", a.id, err)
	}

	a.actual = resp
	return true, nil
}

func (a *TensorboardAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating VertexAITensorboard", "name", a.id)

	req := &pb.CreateTensorboardRequest{
		Parent:      a.id.ParentString(),
		Tensorboard: a.desired,
	}
	op, err := a.gcpClient.CreateTensorboard(ctx, req)
	if err != nil {
		return fmt.Errorf("creating VertexAITensorboard %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting VertexAITensorboard %s creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created VertexAITensorboard", "name", a.id)

	// Fetch fully-populated resource after creation
	latest, err := a.gcpClient.GetTensorboard(ctx, &pb.GetTensorboardRequest{Name: created.Name})
	if err != nil {
		return fmt.Errorf("getting VertexAITensorboard %s after creation: %w", a.id, err)
	}

	return a.updateStatus(ctx, createOp, latest)
}

func (a *TensorboardAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating VertexAITensorboard", "name", a.id)

	diffs, updateMask, err := a.compare(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		desired := proto.Clone(a.desired).(*pb.Tensorboard)
		desired.Name = a.id.String()

		req := &pb.UpdateTensorboardRequest{
			Tensorboard: desired,
			UpdateMask:  updateMask,
		}
		op, err := a.gcpClient.UpdateTensorboard(ctx, req)
		if err != nil {
			return fmt.Errorf("updating VertexAITensorboard %s: %w", a.id, err)
		}
		_, err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting VertexAITensorboard %s update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated VertexAITensorboard", "name", a.id)

		// Fetch fully-populated resource after update
		latest, err = a.gcpClient.GetTensorboard(ctx, &pb.GetTensorboardRequest{Name: a.id.String()})
		if err != nil {
			return fmt.Errorf("getting VertexAITensorboard %s after update: %w", a.id, err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *TensorboardAdapter) compare(ctx context.Context, actual, desired *pb.Tensorboard) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, VertexAITensorboardSpec_v1alpha1_FromProto, VertexAITensorboardSpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *TensorboardAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.Tensorboard) error {
	mapCtx := &direct.MapContext{}
	status := &krm.VertexAITensorboardStatus{}
	status.ObservedState = VertexAITensorboardObservedState_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return op.UpdateStatus(ctx, status, nil)
}

func (a *TensorboardAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.VertexAITensorboard{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(VertexAITensorboardSpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Region = a.id.Location
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Tensorboard)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.Tensorboard)
	u.SetGroupVersionKind(krm.VertexAITensorboardGVK)
	return u, nil
}

func (a *TensorboardAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting VertexAITensorboard", "name", a.id)

	req := &pb.DeleteTensorboardRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTensorboard(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent VertexAITensorboard, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting VertexAITensorboard %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted VertexAITensorboard", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete VertexAITensorboard %s: %w", a.id, err)
	}
	return true, nil
}
