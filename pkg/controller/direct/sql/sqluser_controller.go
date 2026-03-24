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

package sql

import (
	"context"
	"fmt"
	"time"

	api "google.golang.org/api/sqladmin/v1beta4"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/googleapis/gax-go/v2"
)

const sqlUserCtrlName = "sqluser-controller"

func init() {
	registry.RegisterModel(krm.SQLUserGVK, newSQLUserModel)
}

func newSQLUserModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sqlUserModel{config: config}, nil
}

type sqlUserModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &sqlUserModel{}

type sqlUserAdapter struct {
	projectID  string
	instanceID string
	resourceID string
	host       string

	desired *krm.SQLUser
	actual  *api.User

	sqlOperationsClient *api.OperationsService
	sqlUsersClient      *api.UsersService
}

var _ directbase.Adapter = &sqlUserAdapter{}

func (m *sqlUserModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	kube := op.Reader

	obj := &krm.SQLUser{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("converting to %T failed: %w", obj, err)
	}

	resourceID, err := refs.GetResourceID(u)
	if err != nil {
		return nil, err
	}
	obj.Spec.ResourceID = &resourceID

	projectID, ok := u.GetAnnotations()[k8s.ProjectIDAnnotation]
	if !ok {
		projectID = u.GetNamespace()
	}

	// Resolve instanceRef to get the instance name.
	instanceID, err := ResolveSQLUserInstanceRef(ctx, kube, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	desiredCopy := obj.DeepCopy()
	return &sqlUserAdapter{
		projectID:           projectID,
		instanceID:          instanceID,
		resourceID:          resourceID,
		host:                direct.ValueOf(obj.Spec.Host),
		desired:             desiredCopy,
		sqlOperationsClient: gcpClient.sqlOperationsClient(),
		sqlUsersClient:      gcpClient.sqlUsersClient(),
	}, nil
}

func (m *sqlUserModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *sqlUserAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(sqlUserCtrlName)

	if a.resourceID == "" {
		return false, nil
	}

	// The Users API doesn't have a Get method, so we use List and filter.
	users, err := a.sqlUsersClient.List(a.projectID, a.instanceID).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("listing SQLUsers for instance %s failed: %w", a.instanceID, err)
	}

	for _, user := range users.Items {
		if user.Name == a.resourceID {
			// For MySQL, match on host too. For Postgres, host is empty.
			if a.host == "" || user.Host == a.host {
				log.V(2).Info("found SQLUser", "name", user.Name, "host", user.Host)
				a.actual = user
				return true, nil
			}
		}
	}

	return false, nil
}

func (a *sqlUserAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(sqlUserCtrlName)
	log.V(2).Info("creating SQLUser", "name", a.resourceID)

	desired, err := SQLUserKRMToGCP(a.desired)
	if err != nil {
		return err
	}
	desired.Instance = a.instanceID
	desired.Project = a.projectID

	op, err := a.sqlUsersClient.Insert(a.projectID, a.instanceID, desired).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating SQLUser %s failed: %w", a.resourceID, err)
	}

	if err := a.pollForLROCompletion(ctx, op, "create"); err != nil {
		return err
	}

	// Re-read the created user to get server-populated fields.
	created, err := a.findUser(ctx)
	if err != nil {
		return fmt.Errorf("reading created SQLUser %s failed: %w", a.resourceID, err)
	}

	status := &krm.SQLUserStatusFields{}
	if details := SQLUserStatusGCPToKRM(created); details != nil {
		status.SqlServerUserDetails = []krm.SQLUserSqlServerUserDetailsStatus{*details}
	}

	return setStatus(createOp.GetUnstructured(), status)
}

func (a *sqlUserAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(sqlUserCtrlName)
	log.V(2).Info("updating SQLUser", "name", a.resourceID)

	desired, err := SQLUserKRMToGCP(a.desired)
	if err != nil {
		return err
	}
	desired.Instance = a.instanceID
	desired.Project = a.projectID

	op, err := a.sqlUsersClient.Update(a.projectID, a.instanceID, desired).
		Host(a.host).
		Name(a.resourceID).
		Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating SQLUser %s failed: %w", a.resourceID, err)
	}

	if err := a.pollForLROCompletion(ctx, op, "update"); err != nil {
		return err
	}

	// Re-read the updated user.
	updated, err := a.findUser(ctx)
	if err != nil {
		return fmt.Errorf("reading updated SQLUser %s failed: %w", a.resourceID, err)
	}

	status := &krm.SQLUserStatusFields{}
	if details := SQLUserStatusGCPToKRM(updated); details != nil {
		status.SqlServerUserDetails = []krm.SQLUserSqlServerUserDetailsStatus{*details}
	}

	return setStatus(updateOp.GetUnstructured(), status)
}

func (a *sqlUserAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(sqlUserCtrlName)
	log.V(2).Info("deleting SQLUser", "name", a.resourceID)

	op, err := a.sqlUsersClient.Delete(a.projectID, a.instanceID).
		Host(a.host).
		Name(a.resourceID).
		Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent SQLUser", "name", a.resourceID)
			return true, nil
		}
		return false, fmt.Errorf("deleting SQLUser %s failed: %w", a.resourceID, err)
	}

	if err := a.pollForLROCompletion(ctx, op, "delete"); err != nil {
		return false, err
	}

	return true, nil
}

func (a *sqlUserAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("SQLUser %q not found", a.resourceID)
	}

	spec := SQLUserGCPToKRM(a.actual)

	sqlUser := &krm.SQLUser{
		Spec: *spec,
	}
	sqlUser.SetGroupVersionKind(krm.SQLUserGVK)
	sqlUser.SetName(a.resourceID)

	sqlUserObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(sqlUser)
	if err != nil {
		return nil, fmt.Errorf("converting SQLUser spec to unstructured failed: %w", err)
	}

	return &unstructured.Unstructured{Object: sqlUserObj}, nil
}

func (a *sqlUserAdapter) findUser(ctx context.Context) (*api.User, error) {
	users, err := a.sqlUsersClient.List(a.projectID, a.instanceID).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	for _, user := range users.Items {
		if user.Name == a.resourceID {
			if a.host == "" || user.Host == a.host {
				return user, nil
			}
		}
	}
	return nil, fmt.Errorf("user %s not found in instance %s", a.resourceID, a.instanceID)
}

func (a *sqlUserAdapter) pollForLROCompletion(ctx context.Context, op *api.Operation, verb string) error {
	log := klog.FromContext(ctx).WithName(sqlUserCtrlName)

	pollingBackoff := gax.Backoff{
		Initial:    time.Second,
		Max:        time.Minute,
		Multiplier: 2,
	}
	for {
		log.V(2).Info("polling", "op", op)

		if op.Status == "DONE" {
			break
		}
		if err := gax.Sleep(ctx, pollingBackoff.Pause()); err != nil {
			return fmt.Errorf("waiting for SQLUser %s %s failed: %w", a.resourceID, verb, err)
		}
		var err error
		op, err = a.sqlOperationsClient.Get(a.projectID, op.Name).Do()
		if err != nil {
			return fmt.Errorf("getting SQLUser %s %s operation %s failed: %w", a.resourceID, verb, op.Name, err)
		}
	}

	return nil
}
