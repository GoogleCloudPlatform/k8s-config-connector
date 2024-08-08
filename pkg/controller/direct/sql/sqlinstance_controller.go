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
	refs    *SQLInstanceInternalRefs
	actual  *api.DatabaseInstance

	sqlOperationsClient *api.OperationsService
	sqlInstancesClient  *api.InstancesService
	sqlUsersClient      *api.UsersService
}

var _ directbase.Adapter = &sqlInstanceAdapter{}

func (m *sqlInstanceModel) AdapterForObject(ctx context.Context, kube client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.SQLInstance{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
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

	refs, err := NormalizeSQLInstance(ctx, kube, obj)
	if err != nil {
		return nil, err
	}

	return &sqlInstanceAdapter{
		projectID:           projectID,
		resourceID:          resourceID,
		desired:             obj.DeepCopy(),
		refs:                refs,
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
	log.V(2).Info("found cloudsql instance", "actual", a.actual)

	return true, nil
}

func (a *sqlInstanceAdapter) Create(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating instance", "desired", a.desired)

	if a.projectID == "" {
		return fmt.Errorf("project is empty")
	}
	if a.resourceID == "" {
		return fmt.Errorf("resourceID is empty")
	}

	desiredGCP, err := SQLInstanceKRMToGCP(a.desired, a.refs)
	if err != nil {
		return err
	}

	op, err := a.sqlInstancesClient.Insert(a.projectID, desiredGCP).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("create SQLInstance %s failed: %w", a.desired.Name, err)
	}

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
			return fmt.Errorf("wait SQLInstance %s creation failed: %w", a.desired.Name, err)
		}
		op, err = a.sqlOperationsClient.Get(a.projectID, op.Name).Do()
		if err != nil {
			return fmt.Errorf("get SQLInstance %s create operation %s failed: %w", a.desired.Name, op.Name, err)
		}
	}

	created, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("get SQLInstance %s failed: %w", a.desired.Name, err)
	}

	users, err := a.sqlUsersClient.List(a.projectID, a.resourceID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("list SQLInstance %s users failed: %w", a.desired.Name, err)
	}

	if users != nil {
		for _, user := range users.Items {
			if user.Name == "root" && strings.HasPrefix(created.DatabaseVersion, "MYSQL") {
				// Delete "root" user to match Terraform behavior, to improve default security.
				// Ref: https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance
				op, err := a.sqlUsersClient.Delete(a.projectID, a.resourceID).Context(ctx).Name(user.Name).Do()
				if err != nil {
					return fmt.Errorf("delete SQLInstance %s root user failed: %w", a.desired.Name, err)
				}
				for {
					log.V(2).Info("polling", "op", op)

					if op.Status == "DONE" {
						break
					}
					if err := gax.Sleep(ctx, pollingBackoff.Pause()); err != nil {
						return fmt.Errorf("wait SQLInstance %s delete user failed: %w", a.desired.Name, err)
					}
					op, err = a.sqlOperationsClient.Get(a.projectID, op.Name).Do()
					if err != nil {
						return fmt.Errorf("get SQLInstance %s delete root user operation %s failed: %w", a.desired.Name, op.Name, err)
					}
				}
			}
		}
	}

	log.V(2).Info("instance created", "op", op, "instance", created)

	status := &krm.SQLInstanceStatus{}
	if err := Convert_SQLInstance_API_v1_To_KRM_status(created, status); err != nil {
		return fmt.Errorf("update SQLInstance status failed: %w", err)
	}
	return setStatus(u, status)
}

func (a *sqlInstanceAdapter) Update(ctx context.Context, u *unstructured.Unstructured) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating instance", "desired", a.desired)

	// First, handle database version updates
	if a.desired.Spec.DatabaseVersion != nil && *a.desired.Spec.DatabaseVersion != a.actual.DatabaseVersion {
		newVersionDb := &api.DatabaseInstance{
			DatabaseVersion: *a.desired.Spec.DatabaseVersion,
		}
		op, err := a.sqlInstancesClient.Patch(a.projectID, *a.desired.Spec.ResourceID, newVersionDb).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("patch SQLInstance %s version failed: %w", *a.desired.Spec.ResourceID, err)
		}

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
				return fmt.Errorf("wait SQLInstance %s version patch failed: %w", *a.desired.Spec.ResourceID, err)
			}
			op, err = a.sqlOperationsClient.Get(a.projectID, op.Name).Do()
			if err != nil {
				return fmt.Errorf("get SQLInstance %s version patch operation %s failed: %w", *a.desired.Spec.ResourceID, op.Name, err)
			}
		}

		updated, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("get SQLInstance %s failed: %w", *a.desired.Spec.ResourceID, err)
		}

		log.V(2).Info("instance version updated", "op", op, "instance", updated)

		a.actual = updated
	}

	// Next, update rest of the fields
	merged, diffDetected, err := MergeDesiredSQLInstanceWithActual(a.desired, a.refs, a.actual)
	if err != nil {
		return fmt.Errorf("diff SQL instances failed: %w", err)
	}

	if diffDetected {
		op, err := a.sqlInstancesClient.Update(a.projectID, merged.Name, merged).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("update SQLInstance %s failed: %w", merged.Name, err)
		}

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
				return fmt.Errorf("wait SQLInstance %s update failed: %w", merged.Name, err)
			}
			op, err = a.sqlOperationsClient.Get(a.projectID, op.Name).Do()
			if err != nil {
				return fmt.Errorf("get SQLInstance %s update operation %s failed: %w", merged.Name, op.Name, err)
			}
		}

		updated, err := a.sqlInstancesClient.Get(a.projectID, a.resourceID).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("get SQLInstance %s failed: %w", merged.Name, err)
		}

		log.V(2).Info("instance updated", "op", op, "instance", updated)

		status := &krm.SQLInstanceStatus{}
		if err := Convert_SQLInstance_API_v1_To_KRM_status(updated, status); err != nil {
			return fmt.Errorf("update SQLInstance status failed: %w", err)
		}
		return setStatus(u, status)
	}

	return nil
}

// Delete implements the Adapter interface.
func (a *sqlInstanceAdapter) Delete(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting instance", "actual", a.actual)

	if a.resourceID == "" {
		return false, nil
	}
	op, err := a.sqlInstancesClient.Delete(a.projectID, a.resourceID).Context(ctx).Do()
	if err != nil {
		return false, fmt.Errorf("deleting SQLInstance %s: %w", a.resourceID, err)
	}

	log.V(2).Info("deleted instance", "op", op)

	return true, nil
}

func (a *sqlInstanceAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func setStatus(u *unstructured.Unstructured, typedStatus any) error {
	status, err := runtime.DefaultUnstructuredConverter.ToUnstructured(typedStatus)
	if err != nil {
		return fmt.Errorf("error converting status to unstructured: %w", err)
	}

	old, _, _ := unstructured.NestedMap(u.Object, "status")
	if old != nil {
		status["conditions"] = old["conditions"]
		status["observedGeneration"] = old["observedGeneration"]
	}

	u.Object["status"] = status

	return nil
}
