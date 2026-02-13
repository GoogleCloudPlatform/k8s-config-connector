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
	"strings"
	"time"

	api "google.golang.org/api/gkehub/v1beta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.GKEHubMembershipBindingGVK, getGKEHubMembershipBindingModel)
}

func getGKEHubMembershipBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gkeHubMembershipBindingModel{config: config}, nil
}

type gkeHubMembershipBindingModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &gkeHubMembershipBindingModel{}

type gkeHubMembershipBindingAdapter struct {
	id      *GKEHubMembershipBindingIdentity
	desired *krm.GKEHubMembershipBinding
	actual  *api.MembershipBinding

	hubClient *gkeHubClient
}

var _ directbase.Adapter = &gkeHubMembershipBindingAdapter{}

// AdapterForObject implements the Model interface.
func (m *gkeHubMembershipBindingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	gcpClient, err := newGCPClient(m.config)
	if err != nil {
		return nil, err
	}
	hubClient, err := gcpClient.newGkeHubClient(ctx)
	if err != nil {
		return nil, err
	}
	obj := &krm.GKEHubMembershipBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID, err := direct.ResolveProjectID(ctx, reader, obj.GetNamespace(), &obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}

	if err := obj.Spec.MembershipRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	membershipID := obj.Spec.MembershipRef.External

	if err := obj.Spec.ScopeRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, err
	}
	scopeID := obj.Spec.ScopeRef.External

	return &gkeHubMembershipBindingAdapter{
		id: &GKEHubMembershipBindingIdentity{
			Project:      projectID,
			Location:     obj.Spec.Location,
			MembershipID: membershipID,
			ID:           resourceID,
		},
		desired:   obj,
		hubClient: hubClient,
	}, nil
}

func (m *gkeHubMembershipBindingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *gkeHubMembershipBindingAdapter) Find(ctx context.Context) (bool, error) {
	if a.id.ID == "" {
		return false, nil
	}
	binding, err := a.hubClient.membershipBindingClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting membership binding %q: %w", a.id.String(), err)
	}
	a.actual = binding
	return true, nil
}

func (a *gkeHubMembershipBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	op, err := a.hubClient.membershipBindingClient.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting membership binding %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for membership binding deletion %q: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *gkeHubMembershipBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating membership binding", "id", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.krmToApi(mapCtx)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent := a.id.Parent()
	op, err := a.hubClient.membershipBindingClient.Create(parent, desired).MembershipBindingId(a.id.ID).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating membership binding %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for membership binding creation %q: %w", a.id.String(), err)
	}

	actual, err := a.hubClient.membershipBindingClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting membership binding %q after creation: %w", a.id.String(), err)
	}
	a.actual = actual

	return a.setID(createOp.GetUnstructured())
}

func (a *gkeHubMembershipBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating membership binding", "id", a.id)

	mapCtx := &direct.MapContext{}
	desired := a.krmToApi(mapCtx)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := []string{}
	if !reflect.DeepEqual(a.desired.Spec.Labels, a.actual.Labels) {
		updateMask = append(updateMask, "labels")
	}
	if a.desired.Spec.ScopeRef.External != a.actual.Scope {
		updateMask = append(updateMask, "scope")
	}

	if len(updateMask) == 0 {
		return nil
	}

	op, err := a.hubClient.membershipBindingClient.Patch(a.id.String(), desired).UpdateMask(strings.Join(updateMask, ",")).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating membership binding %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for membership binding update %q: %w", a.id.String(), err)
	}

	actual, err := a.hubClient.membershipBindingClient.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting membership binding %q after update: %w", a.id.String(), err)
	}
	a.actual = actual

	return nil
}

func (a *gkeHubMembershipBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	mapCtx := &direct.MapContext{}
	obj := &krm.GKEHubMembershipBinding{}
	obj.Spec = a.apiToKrm(mapCtx)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *gkeHubMembershipBindingAdapter) waitForOp(ctx context.Context, op *api.Operation) error {
	retryPeriod := baseDelay
	timeoutAt := time.Now().Add(timeoutDuration)
	for {
		current, err := a.hubClient.operationClient.Get(op.Name).Context(ctx).Do()
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
		retryPeriod = retryPeriod * 2
	}
}

func (a *gkeHubMembershipBindingAdapter) setID(u *unstructured.Unstructured) error {
	if err := unstructured.SetNestedField(u.Object, a.id.String(), "status", "externalRef"); err != nil {
		return err
	}
	return nil
}

func (a *gkeHubMembershipBindingAdapter) krmToApi(mapCtx *direct.MapContext) *api.MembershipBinding {
	out := &api.MembershipBinding{}
	out.Labels = a.desired.Spec.Labels
	out.Scope = a.desired.Spec.ScopeRef.External
	return out
}

func (a *gkeHubMembershipBindingAdapter) apiToKrm(mapCtx *direct.MapContext) krm.GKEHubMembershipBindingSpec {
	out := krm.GKEHubMembershipBindingSpec{}
	out.ProjectRef = refs.ProjectRef{External: a.id.Project}
	out.Location = a.id.Location
	out.ResourceID = direct.LazyPtr(a.id.ID)
	out.MembershipRef = krm.GKEHubMembershipRef{External: a.id.MembershipID}
	out.ScopeRef = krm.GKEHubScopeRef{External: a.actual.Scope}
	out.Labels = a.actual.Labels
	return out
}

type GKEHubMembershipBindingIdentity struct {
	Project      string
	Location     string
	MembershipID string
	ID           string
}

func (i *GKEHubMembershipBindingIdentity) String() string {
	return fmt.Sprintf("%s/bindings/%s", i.MembershipID, i.ID)
}

func (i *GKEHubMembershipBindingIdentity) Parent() string {
	return i.MembershipID
}
