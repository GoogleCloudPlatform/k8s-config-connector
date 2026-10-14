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

package oslogin

import (
	"context"
	"fmt"
	"strings"

	gcp "cloud.google.com/go/oslogin/apiv1"
	osloginpb "cloud.google.com/go/oslogin/apiv1/osloginpb"
	commonpb "cloud.google.com/go/oslogin/common/commonpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	apiscommon "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/oslogin/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.OSLoginSSHPublicKeyGVK, NewOSLoginSSHPublicKeyModel)
}

func NewOSLoginSSHPublicKeyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelOSLoginSSHPublicKey{config: *config}, nil
}

var _ directbase.Model = &modelOSLoginSSHPublicKey{}

type modelOSLoginSSHPublicKey struct {
	config config.ControllerConfig
}

func (m *modelOSLoginSSHPublicKey) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building OSLogin REST client: %w", err)
	}
	return gcpClient, nil
}

func (m *modelOSLoginSSHPublicKey) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.OSLoginSSHPublicKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.OSLoginSSHPublicKeyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := OSLoginSSHPublicKeySpec_v1alpha1_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &osloginSSHPublicKeyAdapter{
		gcpClient:      gcpClient,
		id:             id,
		desired:        desired,
		desiredProject: obj.Spec.Project,
	}, nil
}

func (m *modelOSLoginSSHPublicKey) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	u := &krm.OSLoginSSHPublicKeyIdentity{}
	if err := u.FromExternal(url); err != nil {
		return nil, nil
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &osloginSSHPublicKeyAdapter{
		gcpClient: gcpClient,
		id:        u,
	}, nil
}

type osloginSSHPublicKeyAdapter struct {
	gcpClient      *gcp.Client
	id             *krm.OSLoginSSHPublicKeyIdentity
	desired        *commonpb.SshPublicKey
	desiredProject *string
	actual         *commonpb.SshPublicKey
}

var _ directbase.Adapter = &osloginSSHPublicKeyAdapter{}

func (a *osloginSSHPublicKeyAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.Fingerprint == "" {
		return false, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("finding OSLoginSSHPublicKey", "name", a.id.String())

	req := &osloginpb.GetSshPublicKeyRequest{
		Name: a.id.String(),
	}
	resp, err := a.gcpClient.GetSshPublicKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting OSLoginSSHPublicKey %q: %w", a.id.String(), err)
	}

	a.actual = resp
	return true, nil
}

func (a *osloginSSHPublicKeyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating OSLoginSSHPublicKey", "user", a.id.User)

	req := &osloginpb.ImportSshPublicKeyRequest{
		Parent:       a.id.ParentString(),
		SshPublicKey: a.desired,
		ProjectId:    apiscommon.ValueOf(a.desiredProject),
	}
	resp, err := a.gcpClient.ImportSshPublicKey(ctx, req)
	if err != nil {
		return fmt.Errorf("importing OSLoginSSHPublicKey: %w", err)
	}

	var createdKey *commonpb.SshPublicKey
	if resp.GetLoginProfile() != nil {
		for _, k := range resp.GetLoginProfile().GetSshPublicKeys() {
			if strings.TrimSpace(k.GetKey()) == strings.TrimSpace(a.desired.GetKey()) {
				createdKey = k
				break
			}
		}
	}
	if createdKey == nil {
		return fmt.Errorf("imported SSH public key not found in response login profile")
	}

	a.id.Fingerprint = createdKey.GetFingerprint()
	log.V(2).Info("successfully imported OSLoginSSHPublicKey", "name", a.id.String())

	return a.updateStatus(ctx, createOp, createdKey)
}

func (a *osloginSSHPublicKeyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	if a.id.Fingerprint == "" {
		return fmt.Errorf("cannot update OSLoginSSHPublicKey with empty fingerprint")
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("updating OSLoginSSHPublicKey", "name", a.id.String())

	diffs, updateMask, err := compareOSLoginSSHPublicKey(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		updatePb := &commonpb.SshPublicKey{}
		for _, path := range updateMask.Paths {
			switch path {
			case "expiration_time_usec", "expirationTimeUsec":
				updatePb.ExpirationTimeUsec = a.desired.ExpirationTimeUsec
			default:
				return fmt.Errorf("field %q is immutable and cannot be updated", path)
			}
		}

		req := &osloginpb.UpdateSshPublicKeyRequest{
			Name:         a.id.String(),
			SshPublicKey: updatePb,
			UpdateMask:   updateMask,
		}

		updated, err := a.gcpClient.UpdateSshPublicKey(ctx, req)
		if err != nil {
			return fmt.Errorf("updating OSLoginSSHPublicKey %s: %w", a.id.String(), err)
		}
		latest = updated
		log.V(2).Info("successfully updated OSLoginSSHPublicKey", "name", a.id.String())
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *osloginSSHPublicKeyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	if a.id.Fingerprint == "" {
		return true, nil
	}

	log := klog.FromContext(ctx)
	log.V(2).Info("deleting OSLoginSSHPublicKey", "name", a.id.String())

	req := &osloginpb.DeleteSshPublicKeyRequest{
		Name: a.id.String(),
	}
	err := a.gcpClient.DeleteSshPublicKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting OSLoginSSHPublicKey %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted OSLoginSSHPublicKey", "name", a.id.String())
	return true, nil
}

func (a *osloginSSHPublicKeyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("OSLoginSSHPublicKey %q not found", a.id.String())
	}

	mapCtx := &direct.MapContext{}
	spec := OSLoginSSHPublicKeySpec_v1alpha1_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	spec.User = a.id.User
	spec.ResourceID = &a.id.Fingerprint

	obj := &krm.OSLoginSSHPublicKey{
		TypeMeta: metav1.TypeMeta{
			Kind:       krm.OSLoginSSHPublicKeyGVK.Kind,
			APIVersion: krm.OSLoginSSHPublicKeyGVK.GroupVersion().String(),
		},
		Spec: *spec,
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting to unstructured: %w", err)
	}

	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *osloginSSHPublicKeyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *commonpb.SshPublicKey) error {
	mapCtx := &direct.MapContext{}
	status := OSLoginSSHPublicKeyStatus_v1alpha1_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareOSLoginSSHPublicKey(ctx context.Context, actual, desired *commonpb.SshPublicKey) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, OSLoginSSHPublicKeySpec_v1alpha1_FromProto, OSLoginSSHPublicKeySpec_v1alpha1_ToProto)
	if err != nil {
		return nil, nil, err
	}
	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
