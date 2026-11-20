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

// +generated:controller
// krm.group: ftp.cnrm.cloud.google.com
// krm.version: v1alpha1
// krm.kind: FTPUser

package ftp

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/ftp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/ftp/apiv1"
	ftppb "cloud.google.com/go/ftp/apiv1/ftppb"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.FTPUserGVK, NewUserModel)
}

func NewUserModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelUser{config: *config}, nil
}

var _ directbase.Model = &modelUser{}

type modelUser struct {
	config config.ControllerConfig
}

func (m *modelUser) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.FTPUser{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}
	if err := common.NormalizeGreenfieldReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}
	i, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := i.(*krm.FTPUserIdentity)

	// Get ftp GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &UserAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelUser) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	if !strings.HasPrefix(url, "//ftp.googleapis.com/") {
		return nil, nil
	}
	id := &krm.FTPUserIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &UserAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}
func (m *modelUser) client(ctx context.Context) (*gcp.CloudFtpClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCloudFtpRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building User client: %w", err)
	}
	return gcpClient, err
}

type UserAdapter struct {
	id        *krm.FTPUserIdentity
	gcpClient *gcp.CloudFtpClient
	desired   *krm.FTPUser
	actual    *ftppb.User
}

var _ directbase.Adapter = &UserAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *UserAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting User", "name", a.id)

	req := &ftppb.GetUserRequest{
		Name: a.id.String(),
		View: ftppb.UserView_USER_VIEW_FULL,
	}
	userpb, err := a.gcpClient.GetUser(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting User %q: %w", a.id, err)
	}

	a.actual = userpb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *UserAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating User", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := FTPUserSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Labels = label.GCPLabels(a.desired)

	req := &ftppb.CreateUserRequest{
		Parent: a.id.ParentString(),
		UserId: a.id.User,
		User:   resource,
	}
	op, err := a.gcpClient.CreateUser(ctx, req)
	if err != nil {
		return fmt.Errorf("creating User %s: %w", a.id, err)
	}
	if _, err := op.Wait(ctx); err != nil {
		return fmt.Errorf("User %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created User", "name", a.id)

	fullUser, err := a.gcpClient.GetUser(ctx, &ftppb.GetUserRequest{
		Name: a.id.String(),
		View: ftppb.UserView_USER_VIEW_FULL,
	})
	if err != nil {
		return fmt.Errorf("getting User %s after creation: %w", a.id, err)
	}

	status := &krm.FTPUserStatus{}
	status.ObservedState = FTPUserObservedState_FromProto(mapCtx, fullUser)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func normalizeUser(ctx context.Context, u *ftppb.User) error {
	for _, cred := range u.UserCredentials {
		// Cloud FTP strips any trailing comment from ssh_public_key_body.
		// Normalizing to "<algorithm> <base64-key>" prevents infinite diffs during reconciliation.
		if cred.SshPublicKeyBody != "" {
			parts := strings.Fields(cred.SshPublicKeyBody)
			if len(parts) >= 2 {
				cred.SshPublicKeyBody = parts[0] + " " + parts[1]
			}
		}
	}
	return nil
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *UserAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating User", "name", a.id)

	resource, report, updateMask, err := common.CompareGreenfieldSpecAndLabels(
		ctx,
		updateOp.GetUnstructured(),
		&a.desired.Spec,
		a.actual,
		"labels",
		FTPUserSpec_FromProto,
		FTPUserSpec_ToProto,
		normalizeUser,
	)
	if err != nil {
		return err
	}
	updated := a.actual
	if !report.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		report.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, report)
		resource.Name = a.id.String()
		req := &ftppb.UpdateUserRequest{
			UpdateMask: updateMask,
			User:       resource,
		}
		op, err := a.gcpClient.UpdateUser(ctx, req)
		if err != nil {
			return fmt.Errorf("updating User %s: %w", a.id, err)
		}
		if _, err := op.Wait(ctx); err != nil {
			return fmt.Errorf("User %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated User", "name", a.id)

		fullUser, err := a.gcpClient.GetUser(ctx, &ftppb.GetUserRequest{
			Name: a.id.String(),
			View: ftppb.UserView_USER_VIEW_FULL,
		})
		if err != nil {
			return fmt.Errorf("getting User %s after update: %w", a.id, err)
		}
		updated = fullUser
	}

	status := &krm.FTPUserStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = FTPUserObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *UserAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.FTPUser{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FTPUserSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.String())
	u.SetGroupVersionKind(krm.FTPUserGVK)

	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *UserAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting User", "name", a.id)

	req := &ftppb.DeleteUserRequest{
		Name: a.id.String(),
	}
	op, err := a.gcpClient.DeleteUser(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent User, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting User %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted User", "name", a.id)
	err = op.Wait(ctx)
	if err != nil {
		if direct.IsNotFound(err) || strings.Contains(err.Error(), "missing \"value\" field") {
			return true, nil
		}
		return false, fmt.Errorf("waiting delete User %s: %w", a.id, err)
	}
	return true, nil
}
