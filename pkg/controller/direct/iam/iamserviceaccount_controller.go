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

package iam

import (
	"context"
	"fmt"
	"strings"

	gcp "google.golang.org/api/iam/v1"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	adminpb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/iam/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func init() {
	registry.RegisterModel(krm.IAMServiceAccountGVK, NewIAMServiceAccountModel)
}

func NewIAMServiceAccountModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &iamServiceAccountModel{config: *config}, nil
}

var _ directbase.Model = &iamServiceAccountModel{}

type iamServiceAccountModel struct {
	config config.ControllerConfig
}

func (m *iamServiceAccountModel) client(ctx context.Context) (*gcp.Service, error) {
	var opts []option.ClientOption

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building iam v1 client: %w", err)
	}

	return gcpClient, nil
}

func (m *iamServiceAccountModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader
	obj := &krm.IAMServiceAccount{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, kube, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	id, err := obj.GetIdentity(ctx, kube)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &iamServiceAccountAdapter{
		gcpClient:   gcpClient,
		id:          id.(*krm.IAMServiceAccountIdentity),
		desiredKube: obj,
	}, nil
}

func (m *iamServiceAccountModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if s, ok := strings.CutPrefix(url, "//iam.googleapis.com/"); ok {
		s = strings.TrimPrefix(s, "v1/")

		var id krm.IAMServiceAccountIdentity
		err := id.FromExternal(s)
		if err != nil {
			log.V(2).Error(err, "url did not match IAMServiceAccount format", "url", url)
			return nil, nil
		}

		gcpClient, err := m.client(ctx)
		if err != nil {
			return nil, err
		}
		return &iamServiceAccountAdapter{
			gcpClient: gcpClient,
			id:        &id,
		}, nil
	}
	return nil, nil
}

type iamServiceAccountAdapter struct {
	gcpClient   *gcp.Service
	id          *krm.IAMServiceAccountIdentity
	desiredKube *krm.IAMServiceAccount
	actual      *adminpb.ServiceAccount
}

var _ directbase.Adapter = &iamServiceAccountAdapter{}

func (a *iamServiceAccountAdapter) gcpName() string {
	return fmt.Sprintf("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", a.id.Project, a.id.Account, a.id.Project)
}

func (a *iamServiceAccountAdapter) Find(ctx context.Context) (bool, error) {
	name := a.gcpName()
	sa, err := a.gcpClient.Projects.ServiceAccounts.Get(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting service account %q: %w", name, err)
	}

	a.actual = ServiceAccount_FromREST(sa)
	return true, nil
}

func (a *iamServiceAccountAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	mapCtx := &direct.MapContext{}
	desired := IAMServiceAccountSpec_ToProto(mapCtx, &a.desiredKube.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	sa := ServiceAccount_ToREST(desired)
	parent := "projects/" + a.id.Project
	req := &gcp.CreateServiceAccountRequest{
		AccountId:      a.id.Account,
		ServiceAccount: sa,
	}

	created, err := a.gcpClient.Projects.ServiceAccounts.Create(parent, req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating service account %q in %s: %w", a.id.Account, parent, err)
	}

	a.actual = ServiceAccount_FromREST(created)
	return a.updateStatus(ctx, createOp, a.actual)
}

func (a *iamServiceAccountAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	mapCtx := &direct.MapContext{}
	desired := IAMServiceAccountSpec_ToProto(mapCtx, &a.desiredKube.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	diffs, _, err := compareServiceAccount(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		return nil
	}

	name := a.gcpName()

	// If disabled field changed
	if a.actual.Disabled != desired.Disabled {
		if desired.Disabled {
			_, err = a.gcpClient.Projects.ServiceAccounts.Disable(name, &gcp.DisableServiceAccountRequest{}).Context(ctx).Do()
			if err != nil {
				return fmt.Errorf("disabling service account %q: %w", name, err)
			}
		} else {
			_, err = a.gcpClient.Projects.ServiceAccounts.Enable(name, &gcp.EnableServiceAccountRequest{}).Context(ctx).Do()
			if err != nil {
				return fmt.Errorf("enabling service account %q: %w", name, err)
			}
		}
	}

	// If DisplayName or Description changed
	var updateMask []string
	if a.actual.Description != desired.Description {
		updateMask = append(updateMask, "description")
	}
	if a.actual.DisplayName != desired.DisplayName {
		updateMask = append(updateMask, "display_name")
	}

	if len(updateMask) > 0 {
		req := &gcp.PatchServiceAccountRequest{
			UpdateMask: strings.Join(updateMask, ","),
			ServiceAccount: &gcp.ServiceAccount{
				DisplayName: desired.DisplayName,
				Description: desired.Description,
				Etag:        string(a.actual.Etag),
			},
		}
		_, err = a.gcpClient.Projects.ServiceAccounts.Patch(name, req).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("patching service account %q: %w", name, err)
		}
	}

	// Fetch latest state to update status accurately
	latest, err := a.gcpClient.Projects.ServiceAccounts.Get(name).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting updated service account %q: %w", name, err)
	}

	a.actual = ServiceAccount_FromREST(latest)
	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *iamServiceAccountAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	name := a.gcpName()
	_, err := a.gcpClient.Projects.ServiceAccounts.Delete(name).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting service account %q: %w", name, err)
	}

	return true, nil
}

func (a *iamServiceAccountAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("service account %q not found", a.id.Account)
	}

	obj := &krm.IAMServiceAccount{}
	obj.SetName(a.id.Account)
	obj.SetGroupVersionKind(krm.IAMServiceAccountGVK)
	if a.desiredKube != nil {
		obj.SetLabels(a.desiredKube.Labels)
	}

	mapCtx := &direct.MapContext{}
	spec := IAMServiceAccountSpec_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	spec.ResourceID = &a.id.Account
	obj.Spec = *spec

	unstructuredMap, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("error converting IAMServiceAccount to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{Object: unstructuredMap}
	return u, nil
}

func (a *iamServiceAccountAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *adminpb.ServiceAccount) error {
	mapCtx := &direct.MapContext{}
	status := IAMServiceAccountStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareServiceAccount(ctx context.Context, actual, desired *adminpb.ServiceAccount) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual := &adminpb.ServiceAccount{
		DisplayName: actual.DisplayName,
		Description: actual.Description,
		Disabled:    actual.Disabled,
	}

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func ServiceAccount_FromREST(in *gcp.ServiceAccount) *adminpb.ServiceAccount {
	if in == nil {
		return nil
	}
	return &adminpb.ServiceAccount{
		Name:           in.Name,
		ProjectId:      in.ProjectId,
		UniqueId:       in.UniqueId,
		Email:          in.Email,
		DisplayName:    in.DisplayName,
		Etag:           []byte(in.Etag),
		Description:    in.Description,
		Oauth2ClientId: in.Oauth2ClientId,
		Disabled:       in.Disabled,
	}
}

func ServiceAccount_ToREST(in *adminpb.ServiceAccount) *gcp.ServiceAccount {
	if in == nil {
		return nil
	}
	return &gcp.ServiceAccount{
		Name:           in.Name,
		ProjectId:      in.ProjectId,
		UniqueId:       in.UniqueId,
		Email:          in.Email,
		DisplayName:    in.DisplayName,
		Etag:           string(in.Etag),
		Description:    in.Description,
		Oauth2ClientId: in.Oauth2ClientId,
		Disabled:       in.Disabled,
	}
}
