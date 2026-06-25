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
// proto.service: google.cloud.recaptchaenterprise.v1.RecaptchaEnterpriseService
// proto.message: google.cloud.recaptchaenterprise.v1.Key
// crd.type: RecaptchaEnterpriseKey
// crd.version: v1beta1

package recaptchaenterprise

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/recaptchaenterprise/v2/apiv1"
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.RecaptchaEnterpriseKeyGVK, NewKeyModel)
}

func NewKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelKey{config: config}, nil
}

var _ directbase.Model = &modelKey{}

type modelKey struct {
	config *config.ControllerConfig
}

func (m *modelKey) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.RecaptchaEnterpriseKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	keyID, ok := id.(*krm.RecaptchaEnterpriseKeyIdentity)
	if !ok {
		return nil, fmt.Errorf("expected *krm.RecaptchaEnterpriseKeyIdentity, got %T", id)
	}

	gcpClient, err := newReCAPTCHAEnterpriseClient(ctx, m.config)
	if err != nil {
		return nil, err
	}
	return &RecaptchaEnterpriseKeyAdapter{
		gcpClient: gcpClient,
		id:        keyID,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelKey) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type RecaptchaEnterpriseKeyAdapter struct {
	gcpClient *gcp.Client
	id        *krm.RecaptchaEnterpriseKeyIdentity
	desired   *krm.RecaptchaEnterpriseKey
	actual    *pb.Key
	reader    client.Reader
}

var _ directbase.Adapter = &RecaptchaEnterpriseKeyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *RecaptchaEnterpriseKeyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting RecaptchaEnterpriseKey", "name", a.id)

	resourceID := direct.ValueOf(a.desired.Spec.ResourceID)
	if resourceID == "" {
		log.V(2).Info("no resource ID in spec indicates the create intention", "name", a.id)
		return false, nil
	}

	req := &pb.GetKeyRequest{Name: a.id.String()}
	keypb, err := a.gcpClient.GetKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, fmt.Errorf("key %q can't be acquired: %w; unset 'spec.resourceID' if you want to create it", a.id, err)
		}
		return false, fmt.Errorf("getting RecaptchaEnterpriseKey %q: %w", a.id, err)
	}

	a.actual = keypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *RecaptchaEnterpriseKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating RecaptchaEnterpriseKey", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := RecaptchaEnterpriseKeySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateKeyRequest{
		Parent: "projects/" + a.id.Project,
		Key:    resource,
	}
	created, err := a.gcpClient.CreateKey(ctx, req)
	if err != nil {
		return fmt.Errorf("creating RecaptchaEnterpriseKey %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created RecaptchaEnterpriseKey", "name", a.id)

	// Since the ID is server-generated, write it back to spec.resourceID.
	resourceID := lastComponent(created.Name)
	if err := createOp.SetSpecResourceID(ctx, resourceID); err != nil {
		return err
	}

	status := RecaptchaEnterpriseKeyStatus_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *RecaptchaEnterpriseKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating RecaptchaEnterpriseKey", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := RecaptchaEnterpriseKeySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.actual.Name

	paths, err := common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateKeyRequest{
		Key:        desiredPb,
		UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
	}
	updated, err := a.gcpClient.UpdateKey(ctx, req)
	if err != nil {
		return fmt.Errorf("updating RecaptchaEnterpriseKey %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated RecaptchaEnterpriseKey", "name", a.id)

	status := RecaptchaEnterpriseKeyStatus_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *RecaptchaEnterpriseKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.RecaptchaEnterpriseKey{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(RecaptchaEnterpriseKeySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = refs.ProjectRef{
		External: "projects/" + a.id.Project,
	}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Key)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Key)
	u.SetGroupVersionKind(krm.RecaptchaEnterpriseKeyGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *RecaptchaEnterpriseKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting RecaptchaEnterpriseKey", "name", a.id)

	req := &pb.DeleteKeyRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent RecaptchaEnterpriseKey, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting RecaptchaEnterpriseKey %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted RecaptchaEnterpriseKey", "name", a.id)

	return true, nil
}

func lastComponent(s string) string {
	parts := strings.Split(s, "/")
	return parts[len(parts)-1]
}
