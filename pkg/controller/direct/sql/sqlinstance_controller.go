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
	"strings"
	"time"

	api "google.golang.org/api/sqladmin/v1beta4"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/googleapis/gax-go/v2"
)

const ctrlName = "sqlinstance-controller"

func init() {
	registry.RegisterModel(krm.SQLInstanceGVK, newSQLInstanceModel)
}

func newSQLInstanceModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &sqlInstanceModel{config: config}, nil
}

type sqlInstanceModel struct {
	config *config.ControllerConfig
}

var _ directbase.Model = &sqlInstanceModel{}

type sqlInstanceAdapter struct {
	projectID  string
	resourceID string

	desired *krm.SQLInstance
	actual  *api.DatabaseInstance

	sqlOperationsClient *api.OperationsService
	sqlInstancesClient  *api.InstancesService
	sqlUsersClient      *api.UsersService
}

var _ directbase.Adapter = &sqlInstanceAdapter{}

func (m *sqlInstanceModel) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SQLInstance{}
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

	gcpClient, err := newGCPClient(ctx, m.config)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}

	if err := ResolveSQLInstanceRefs(ctx, kube, obj); err != nil {
		return nil, err
	}

	return &sqlInstanceAdapter{
		projectID:           projectID,
		resourceID:          resourceID,
		desired:             obj.DeepCopy(),
		sqlOperationsClient: gcpClient.sqlOperationsClient(),
		sqlInstancesClient:  gcpClient.sqlInstancesClient(),
		sqlUsersClient:      gcpClient.sqlUsersClient(),
	}, nil
}

func (m *sqlInstanceModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *sqlInstanceAdapter) Find(ctx context.Context) (bool, error) {
	if a.resourceID == "" {
		return false, nil
	}

	instance, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
	if err != nil {
		return false, nil
	}

	a.actual = instance

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("found SQLInstance", "actual", a.actual)

	return true, nil
}

func (a *sqlInstanceAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating SQLInstance", "desired", a.desired)

	if a.projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	if a.desired.Spec.CloneSource != nil {
		return a.cloneInstance(ctx, u, log)
	} else {
		return a.insertInstance(ctx, u, log)
	}
}

func (a *sqlInstanceAdapter) cloneInstance(ctx context.Context, u *unstructured.Unstructured, log klog.Logger) error {
	desiredGCP, err := SQLInstanceCloneKRMToGCP(a.desired)
	if err != nil {
		return err
	}

	sourceInstance := a.desired.Spec.CloneSource.SQLInstanceRef.External
	op, err := a.sqlInstancesClient.Clone(a.projectID, sourceInstance, desiredGCP).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("cloning SQLInstance %s failed: %w", a.desired.Name, err)
	}
	if err := a.pollForLROCompletion(ctx, op, "clone"); err != nil {
		return err
	}

	created, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting SQLInstance %s failed: %w", a.desired.Name, err)
	}

	log.V(2).Info("instance cloned", "op", op, "instance", created)

	status, err := SQLInstanceStatusGCPToKRM(created)
	if err != nil {
		return fmt.Errorf("updating SQLInstance status failed: %w", err)
	}
	return setStatus(u, status)
}

func (a *sqlInstanceAdapter) insertInstance(ctx context.Context, u *unstructured.Unstructured, log klog.Logger) error {
	desiredGCP, err := SQLInstanceKRMToGCP(a.desired, a.actual)
	if err != nil {
		return err
	}

	op, err := a.sqlInstancesClient.Insert(a.projectID, desiredGCP).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating SQLInstance %s failed: %w", a.desired.Name, err)
	}
	if err := a.pollForLROCompletion(ctx, op, "create"); err != nil {
		return err
	}

	created, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting SQLInstance %s failed: %w", a.desired.Name, err)
	}

	users, err := a.sqlUsersClient.List(a.projectID, a.resourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("listing SQLInstance %s users failed: %w", a.desired.Name, err)
	}

	if users != nil {
		for _, user := range users.Items {
			if user.Name == "root" && strings.HasPrefix(created.DatabaseVersion, "MYSQL") {
				// Delete "root" user to match Terraform behavior, to improve default security.
				// Ref: https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance
				op, err := a.sqlUsersClient.Delete(a.projectID, a.resourceID).Context(ctx).Name(user.Name).Host(user.Host).Do()
				if err != nil {
					return fmt.Errorf("deleting SQLInstance %s root user failed: %w", a.desired.Name, err)
				}
				if err := a.pollForLROCompletion(ctx, op, "delete root user"); err != nil {
					return err
				}
			}
		}
	}

	log.V(2).Info("instance created", "op", op, "instance", created)

	status, err := SQLInstanceStatusGCPToKRM(created)
	if err != nil {
		return fmt.Errorf("updating SQLInstance status failed: %w", err)
	}
	return setStatus(u, status)
}

func (a *sqlInstanceAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	u := updateOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("updating SQLInstance", "desired", a.desired)

	// First, handle database version updates
	if a.desired.Spec.DatabaseVersion != nil && *a.desired.Spec.DatabaseVersion != a.actual.DatabaseVersion {
		newVersionDb := &api.DatabaseInstance{
			DatabaseVersion: direct.ValueOf(a.desired.Spec.DatabaseVersion),
		}
		op, err := a.sqlInstancesClient.Patch(a.projectID, a.resourceID, newVersionDb).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("patching SQLInstance %s version failed: %w", a.resourceID, err)
		}
		if err := a.pollForLROCompletion(ctx, op, "version patch"); err != nil {
			return err
		}

		updated, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting SQLInstance %s failed: %w", a.resourceID, err)
		}

		log.V(2).Info("instance version updated", "op", op, "instance", updated)

		a.actual = updated
	}

	// Next, handle database edition updates
	if a.desired.Spec.Settings.Edition != nil && *a.desired.Spec.Settings.Edition != a.actual.Settings.Edition {
		newEditionDb := &api.DatabaseInstance{
			Settings: &api.Settings{
				Edition: direct.ValueOf(a.desired.Spec.Settings.Edition),
				// ENTERPRISE_PLUS edition has limitations on the allowable set of tiers that can be used. Therefore, when
				// modifying the edition, we should also allow modifications to the tier at the same time, so that the
				// user can update from an invalid tier to a valid tier (when going from ENTERPRISE -> ENTERPRISE_PLUS).
				Tier: a.desired.Spec.Settings.Tier,
			},
		}
		op, err := a.sqlInstancesClient.Patch(a.projectID, a.resourceID, newEditionDb).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("patching SQLInstance %s edition failed: %w", a.resourceID, err)
		}
		if err := a.pollForLROCompletion(ctx, op, "edition patch"); err != nil {
			return err
		}

		updated, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting SQLInstance %s failed: %w", a.resourceID, err)
		}

		log.V(2).Info("instance edition updated", "op", op, "instance", updated)

		a.actual = updated
	}

	// Finally, update rest of the fields
	desiredGCP, err := SQLInstanceKRMToGCP(a.desired, a.actual)
	if err != nil {
		return err
	}

	if !InstancesMatch(desiredGCP, a.actual) {
		updateOp.RecordUpdatingEvent()

		op, err := a.sqlInstancesClient.Update(a.projectID, desiredGCP.Name, desiredGCP).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating SQLInstance %s failed: %w", desiredGCP.Name, err)
		}
		if err := a.pollForLROCompletion(ctx, op, "update"); err != nil {
			return err
		}

		updated, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting SQLInstance %s failed: %w", desiredGCP.Name, err)
		}

		log.V(2).Info("instance updated", "op", op, "instance", updated)

		status, err := SQLInstanceStatusGCPToKRM(updated)
		if err != nil {
			return fmt.Errorf("updating SQLInstance status failed: %w", err)
		}
		return setStatus(u, status)
	}

	return nil
}

// Delete implements the Adapter interface.
func (a *sqlInstanceAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting SQLInstance", "actual", a.actual)

	op, err := a.sqlInstancesClient.Delete(a.projectID, a.resourceID).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("deleting SQLInstance %s failed: %w", a.resourceID, err)
	}
	if err := a.pollForLROCompletion(ctx, op, "delete"); err != nil {
		return false, err
	}

	log.V(2).Info("deleted SQLInstance", "op", op)

	return true, nil
}

func (a *sqlInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("SQLInstance %q not found", a.resourceID)
	}

	sqlInstance, err := SQLInstanceGCPToKRM(a.actual)
	if err != nil {
		return nil, fmt.Errorf("converting SQLInstance from API failed: %w", err)
	}

	sqlInstanceObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(sqlInstance)
	if err != nil {
		return nil, fmt.Errorf("converting SQLInstance spec to unstructured failed: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: sqlInstanceObj,
	}
	u.SetName(a.resourceID)
	u.SetGroupVersionKind(krm.SQLInstanceGVK)

	return u, nil
}

func (a *sqlInstanceAdapter) pollForLROCompletion(ctx context.Context, op *api.Operation, verb string) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	var err error

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
			return fmt.Errorf("waiting for SQLInstance %s %s failed: %w", a.resourceID, verb, err)
		}
		op, err = a.sqlOperationsClient.Get(a.projectID, op.Name).Do()
		if err != nil {
			return fmt.Errorf("getting SQLInstance %s %s operation %s failed: %w", a.resourceID, verb, op.Name, err)
		}
	}

	return nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("converting status to unstructured failed: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
	}

	u.Object["status"] = status

	return nil
}
