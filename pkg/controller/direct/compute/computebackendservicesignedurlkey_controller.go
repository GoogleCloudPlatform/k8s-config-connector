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

// +tool:controller
// proto.service: google.cloud.compute.v1.BackendServices
// proto.message: google.cloud.compute.v1.SignedUrlKey
// crd.type: ComputeBackendServiceSignedURLKey
// crd.version: v1alpha1

package compute

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1alpha1"
	secret "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1/secret"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.ComputeBackendServiceSignedURLKeyGVK, NewComputeBackendServiceSignedURLKeyModel)
}

func NewComputeBackendServiceSignedURLKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &computeBackendServiceSignedURLKeyModel{config: config}, nil
}

var _ directbase.Model = &computeBackendServiceSignedURLKeyModel{}

type computeBackendServiceSignedURLKeyModel struct {
	config *config.ControllerConfig
}

func (m *computeBackendServiceSignedURLKeyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.ComputeBackendServiceSignedURLKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	backendServicesClient, err := gcpClient.newBackendServicesClient(ctx)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	mapCtx := &direct.MapContext{}
	desired := obj.DeepCopy()
	resource := ComputeBackendServiceSignedURLKeySpec_v1alpha1_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	keyId := id.(*krm.ComputeBackendServiceSignedURLKeyIdentity)
	resource.KeyName = &keyId.SignedUrlKey

	return &ComputeBackendServiceSignedURLKeyAdapter{
		gcpClient: backendServicesClient,
		id:        keyId,
		desired:   resource,
		keyValue:  &obj.Spec.KeyValue,
		namespace: obj.Namespace,
		reader:    reader,
	}, nil
}

func (m *computeBackendServiceSignedURLKeyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type ComputeBackendServiceSignedURLKeyAdapter struct {
	gcpClient *compute.BackendServicesClient
	id        *krm.ComputeBackendServiceSignedURLKeyIdentity
	desired   *pb.SignedUrlKey
	actual    *pb.SignedUrlKey
	keyValue  *secret.Legacy
	namespace string
	reader    client.Reader
}

var _ directbase.Adapter = &ComputeBackendServiceSignedURLKeyAdapter{}

func (a *ComputeBackendServiceSignedURLKeyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting ComputeBackendServiceSignedURLKey", "name", a.id)

	req := &pb.GetBackendServiceRequest{
		Project:        a.id.Project,
		BackendService: a.id.BackendService,
	}
	backendService, err := a.gcpClient.Get(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BackendService for ComputeBackendServiceSignedURLKey %q: %w", a.id, err)
	}

	if backendService.CdnPolicy != nil {
		for _, keyName := range backendService.CdnPolicy.SignedUrlKeyNames {
			if keyName == a.id.SignedUrlKey {
				keyNameCopy := keyName
				a.actual = &pb.SignedUrlKey{
					KeyName: &keyNameCopy,
				}
				return true, nil
			}
		}
	}

	return false, nil
}

func (a *ComputeBackendServiceSignedURLKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating ComputeBackendServiceSignedURLKey", "name", a.id)

	secretVal, err := a.keyValue.ReadSecretValue(ctx, "spec.keyValue", a.namespace, a.reader)
	if err != nil {
		return fmt.Errorf("resolving keyValue secret: %w", err)
	}
	if secretVal == nil {
		return fmt.Errorf("keyValue secret is empty")
	}

	keyName := a.id.SignedUrlKey
	req := &pb.AddSignedUrlKeyBackendServiceRequest{
		Project:        a.id.Project,
		BackendService: a.id.BackendService,
		SignedUrlKeyResource: &pb.SignedUrlKey{
			KeyName:  &keyName,
			KeyValue: secretVal,
		},
	}
	op, err := a.gcpClient.AddSignedUrlKey(ctx, req)
	if err != nil {
		return fmt.Errorf("adding ComputeBackendServiceSignedURLKey %s: %w", a.id, err)
	}
	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return fmt.Errorf("waiting for ComputeBackendServiceSignedURLKey %s create: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully created ComputeBackendServiceSignedURLKey", "name", a.id)

	created := &pb.SignedUrlKey{
		KeyName: &keyName,
	}

	return a.updateStatus(ctx, createOp, created)
}

func (a *ComputeBackendServiceSignedURLKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating ComputeBackendServiceSignedURLKey", "name", a.id)

	diffs, _, err := compareComputeBackendServiceSignedURLKey(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)
		return fmt.Errorf("ComputeBackendServiceSignedURLKey is immutable and cannot be updated")
	}

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *ComputeBackendServiceSignedURLKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting ComputeBackendServiceSignedURLKey", "name", a.id)

	req := &pb.DeleteSignedUrlKeyBackendServiceRequest{
		Project:        a.id.Project,
		BackendService: a.id.BackendService,
		KeyName:        a.id.SignedUrlKey,
	}
	op, err := a.gcpClient.DeleteSignedUrlKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting ComputeBackendServiceSignedURLKey %s: %w", a.id, err)
	}
	if !op.Done() {
		if err := op.Wait(ctx); err != nil {
			return false, fmt.Errorf("waiting for ComputeBackendServiceSignedURLKey %s delete: %w", a.id, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeBackendServiceSignedURLKey", "name", a.id)
	return true, nil
}

func (a *ComputeBackendServiceSignedURLKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.ComputeBackendServiceSignedURLKey{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(ComputeBackendServiceSignedURLKeySpec_v1alpha1_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.actual.GetKeyName())
	u.SetGroupVersionKind(krm.ComputeBackendServiceSignedURLKeyGVK)
	return u, nil
}

func (a *ComputeBackendServiceSignedURLKeyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SignedUrlKey) error {
	status := &krm.ComputeBackendServiceSignedURLKeyStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func compareComputeBackendServiceSignedURLKey(ctx context.Context, actual, desired *pb.SignedUrlKey) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, ComputeBackendServiceSignedURLKeySpec_v1alpha1_FromProto, ComputeBackendServiceSignedURLKeySpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.KeyName = desired.KeyName

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.SignedUrlKey) {
		// KeyValue is write-only in GCP and not returned in Get.
		// To avoid a false positive diff, we populate it from desired if it is nil.
		if obj.KeyValue == nil && clonedDesired.KeyValue != nil {
			obj.KeyValue = clonedDesired.KeyValue
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
