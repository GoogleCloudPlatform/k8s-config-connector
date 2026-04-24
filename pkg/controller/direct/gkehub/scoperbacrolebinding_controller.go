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

package gkehub

import (
	"context"
	"fmt"
	"reflect"
	"time"

	gkehubapi "google.golang.org/api/gkehub/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.GKEHubScopeRBACRoleBindingGVK, getGkeHubScopeRBACRoleBindingModel)
}

func getGkeHubScopeRBACRoleBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gkeHubScopeRBACRoleBindingModel{config: config}, nil
}

type gkeHubScopeRBACRoleBindingModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &gkeHubScopeRBACRoleBindingModel{}

type gkeHubScopeRBACRoleBindingAdapter struct {
	id        *krm.GKEHubScopeRBACRoleBindingIdentity
	desired   *krm.GKEHubScopeRBACRoleBinding
	actual    *gkehubapi.RBACRoleBinding
	hubClient *gkeHubClient
	reader    client.Reader
}

var _ directbase.Adapter = &gkeHubScopeRBACRoleBindingAdapter{}

// AdapterForObject implements the Model interface.
func (m *gkeHubScopeRBACRoleBindingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.Object
	reader := op.Reader
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}
	hubClient, err := gcpClient.newGkeHubClient(ctx)
	if err != nil {
		return nil, err
	}
	obj := &krm.GKEHubScopeRBACRoleBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	return &gkeHubScopeRBACRoleBindingAdapter{
		id:        id.(*krm.GKEHubScopeRBACRoleBindingIdentity),
		desired:   obj,
		hubClient: hubClient,
		reader:    reader,
	}, nil
}

func (m *gkeHubScopeRBACRoleBindingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Find(ctx context.Context) (bool, error) {
	if a.id == nil {
		return false, nil
	}
	actual, err := a.hubClient.rbacrolebindingClientV1.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GKEHubScopeRBACRoleBinding %q: %w", a.id.String(), err)
	}
	a.actual = actual
	return true, nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	op, err := a.hubClient.rbacrolebindingClientV1.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting GKEHubScopeRBACRoleBinding %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for GKEHubScopeRBACRoleBinding deletion %q: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GKEHubScopeRBACRoleBinding", "id", a.id.String())

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := GKEHubScopeRBACRoleBindingSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.Parent().String()
	op, err := a.hubClient.rbacrolebindingClientV1.Create(parent, desired).RbacrolebindingId(a.id.ID()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating GKEHubScopeRBACRoleBinding %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for GKEHubScopeRBACRoleBinding creation %q: %w", a.id.String(), err)
	}

	// After creation, we need to get the latest state
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("getting GKEHubScopeRBACRoleBinding after creation %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created GKEHubScopeRBACRoleBinding", "id", a.id.String())

	// Update status
	status := GKEHubScopeRBACRoleBindingStatus_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GKEHubScopeRBACRoleBinding", "id", a.id.String())

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := GKEHubScopeRBACRoleBindingSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Manual diffing because we don't have a proto Message for RBACRoleBinding
	diff := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	if !reflect.DeepEqual(a.actual.Role, desired.Role) {
		diff.AddField("role", a.actual.Role, desired.Role)
	}
	if a.actual.User != desired.User {
		diff.AddField("user", a.actual.User, desired.User)
	}
	if a.actual.Group != desired.Group {
		diff.AddField("group", a.actual.Group, desired.Group)
	}

	if !diff.HasDiff() {
		log.V(2).Info("no field needs update", "id", a.id.String())
		// Ensure ExternalRef is set even if no update is needed
		if a.desired.Status.ExternalRef == nil {
			status := GKEHubScopeRBACRoleBindingStatus_FromProto(mapCtx, a.actual)
			status.ExternalRef = direct.LazyPtr(a.id.String())
			return updateOp.UpdateStatus(ctx, status, nil)
		}
		return nil
	}

	// Report the detected differences
	structuredreporting.ReportDiff(ctx, diff)

	// Check immutability
	if a.actual.User != desired.User {
		return fmt.Errorf("user is immutable")
	}
	if a.actual.Group != desired.Group {
		return fmt.Errorf("group is immutable")
	}

	// Update only if role changed
	if !reflect.DeepEqual(a.actual.Role, desired.Role) {
		updateOp.RecordUpdatingEvent()
		op, err := a.hubClient.rbacrolebindingClientV1.Patch(a.id.String(), desired).UpdateMask("role").Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("patching GKEHubScopeRBACRoleBinding %q: %w", a.id.String(), err)
		}
		if err := a.waitForOp(ctx, op); err != nil {
			return fmt.Errorf("waiting for GKEHubScopeRBACRoleBinding update %q: %w", a.id.String(), err)
		}

		// After update, we need to get the latest state
		if _, err := a.Find(ctx); err != nil {
			return fmt.Errorf("getting GKEHubScopeRBACRoleBinding after update %q: %w", a.id.String(), err)
		}
	}

	log.V(2).Info("successfully updated GKEHubScopeRBACRoleBinding", "id", a.id.String())

	// Update status
	status := GKEHubScopeRBACRoleBindingStatus_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *gkeHubScopeRBACRoleBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	spec := &krm.GKEHubScopeRBACRoleBindingSpec{
		ProjectRef: a.desired.Spec.ProjectRef,
		Location:   a.desired.Spec.Location,
		ScopeRef:   a.desired.Spec.ScopeRef,
		Role: &krm.GKEHubScopeRBACRoleBindingRole{
			PredefinedRole: direct.LazyPtr(a.actual.Role.PredefinedRole),
			CustomRole:     direct.LazyPtr(a.actual.Role.CustomRole),
		},
		User:       direct.LazyPtr(a.actual.User),
		Group:      direct.LazyPtr(a.actual.Group),
		ResourceID: a.desired.Spec.ResourceID,
	}

	obj := &krm.GKEHubScopeRBACRoleBinding{
		Spec: *spec,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) normalizeReferences(ctx context.Context) error {
	if a.desired.Spec.ProjectRef != nil {
		if err := a.desired.Spec.ProjectRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return err
		}
	}
	if a.desired.Spec.ScopeRef != nil {
		if err := a.desired.Spec.ScopeRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
			return err
		}
	}
	return nil
}

func (a *gkeHubScopeRBACRoleBindingAdapter) waitForOp(ctx context.Context, op *gkehubapi.Operation) error {
	retryPeriod := 5 * time.Second
	timeoutDuration := 20 * time.Minute
	timeoutAt := time.Now().Add(timeoutDuration)
	for {
		current, err := a.hubClient.operationClientV1.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting operation status of %q failed: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return fmt.Errorf("operation %q completed with error: %v", op.Name, current.Error)
			} else {
				return nil
			}
		}
		if time.Now().After(timeoutAt) {
			return fmt.Errorf("operation timed out waiting for LRO after %s", timeoutDuration.String())
		}
		time.Sleep(retryPeriod)
		if retryPeriod < 30*time.Second {
			retryPeriod = retryPeriod * 2
		}
	}
}
