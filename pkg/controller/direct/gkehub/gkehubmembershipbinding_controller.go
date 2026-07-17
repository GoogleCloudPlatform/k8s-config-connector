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

	gkehubv1 "google.golang.org/api/gkehub/v1"
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
	registry.RegisterModel(krm.GKEHubMembershipBindingGVK, getGkeHubMembershipBindingModel)
}

func getGkeHubMembershipBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &gkeHubMembershipBindingModel{config: config}, nil
}

type gkeHubMembershipBindingModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &gkeHubMembershipBindingModel{}

type gkeHubMembershipBindingAdapter struct {
	id        *krm.GKEHubMembershipBindingIdentity
	desired   *krm.GKEHubMembershipBinding
	actual    *gkehubv1.MembershipBinding
	hubClient *gkeHubClient
	reader    client.Reader
}

var _ directbase.Adapter = &gkeHubMembershipBindingAdapter{}

// AdapterForObject implements the Model interface.
func (m *gkeHubMembershipBindingModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
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
	obj := &krm.GKEHubMembershipBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	return &gkeHubMembershipBindingAdapter{
		id:        id.(*krm.GKEHubMembershipBindingIdentity),
		desired:   obj,
		hubClient: hubClient,
		reader:    reader,
	}, nil
}

func (m *gkeHubMembershipBindingModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

func (a *gkeHubMembershipBindingAdapter) Find(ctx context.Context) (bool, error) {
	if a.id == nil {
		return false, nil
	}
	actual, err := a.hubClient.membershipBindingClientV1.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting GKEHubMembershipBinding %q: %w", a.id.String(), err)
	}
	a.actual = actual
	return true, nil
}

func (a *gkeHubMembershipBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	op, err := a.hubClient.membershipBindingClientV1.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting GKEHubMembershipBinding %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return false, fmt.Errorf("waiting for GKEHubMembershipBinding deletion %q: %w", a.id.String(), err)
	}
	return true, nil
}

func (a *gkeHubMembershipBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating GKEHubMembershipBinding", "id", a.id.String())

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := GKEHubMembershipBindingSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Scope = a.desired.Spec.ScopeRef.External

	parent := a.id.Parent()
	op, err := a.hubClient.membershipBindingClientV1.Create(parent, desired).MembershipBindingId(a.id.MembershipBinding).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating GKEHubMembershipBinding %q: %w", a.id.String(), err)
	}
	if err := a.waitForOp(ctx, op); err != nil {
		return fmt.Errorf("waiting for GKEHubMembershipBinding creation %q: %w", a.id.String(), err)
	}

	// After creation, we need to get the latest state
	if _, err := a.Find(ctx); err != nil {
		return fmt.Errorf("getting GKEHubMembershipBinding after creation %q: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created GKEHubMembershipBinding", "id", a.id.String())

	// Update status
	status := GKEHubMembershipBindingStatus_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *gkeHubMembershipBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating GKEHubMembershipBinding", "id", a.id.String())

	if err := a.normalizeReferences(ctx); err != nil {
		return err
	}

	mapCtx := &direct.MapContext{}
	desired := GKEHubMembershipBindingSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	if !reflect.DeepEqual(a.actual.Labels, desired.Labels) {
		updateMask := "labels"

		report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
		report.AddField("labels", a.actual.Labels, desired.Labels)
		structuredreporting.ReportDiff(ctx, report)

		op, err := a.hubClient.membershipBindingClientV1.Patch(a.id.String(), desired).UpdateMask(updateMask).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("patching GKEHubMembershipBinding %q: %w", a.id.String(), err)
		}
		if err := a.waitForOp(ctx, op); err != nil {
			return fmt.Errorf("waiting for GKEHubMembershipBinding update %q: %w", a.id.String(), err)
		}

		// After update, we need to get the latest state
		if _, err := a.Find(ctx); err != nil {
			return fmt.Errorf("getting GKEHubMembershipBinding after update %q: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated GKEHubMembershipBinding", "id", a.id.String())
	} else {
		log.V(2).Info("no diff, skipping updating GKEHubMembershipBinding", "id", a.id.String())
	}

	// Update status
	status := GKEHubMembershipBindingStatus_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	externalRef := a.id.String()
	status.ExternalRef = &externalRef

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *gkeHubMembershipBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	mapCtx := &direct.MapContext{}
	spec := GKEHubMembershipBindingSpec_FromAPI(mapCtx, a.actual, a.id)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj := &krm.GKEHubMembershipBinding{
		Spec: *spec,
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: uObj}, nil
}

func (a *gkeHubMembershipBindingAdapter) waitForOp(ctx context.Context, op *gkehubv1.Operation) error {
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

func (a *gkeHubMembershipBindingAdapter) normalizeReferences(ctx context.Context) error {
	if err := a.desired.Spec.MembershipRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
		return err
	}
	if err := a.desired.Spec.ScopeRef.Normalize(ctx, a.reader, a.desired.Namespace); err != nil {
		return err
	}
	return nil
}
