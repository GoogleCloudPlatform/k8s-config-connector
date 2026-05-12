// Copyright 2025 Google LLC
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

package binaryauthorization

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/binaryauthorization/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	// TODO(contributor): Update the import with the google cloud client

	// TODO(contributor): Update the import with the google cloud client api protobuf
	api "google.golang.org/api/binaryauthorization/v1"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BinaryAuthorizationPlatformPolicyGVK, NewPlatformPolicyModel)
}

func NewPlatformPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelPlatformPolicy{config: *config}, nil
}

var _ directbase.Model = &modelPlatformPolicy{}

type modelPlatformPolicy struct {
	config config.ControllerConfig
}

func (m *modelPlatformPolicy) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building PlatformPolicy client: %w", err)
	}
	return gcpClient, err
}

func (m *modelPlatformPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BinaryAuthorizationPlatformPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.BinaryAuthorizationPlatformPolicyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}
	if err != nil {
		return nil, err
	}

	// Get binaryauthorization GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &PlatformPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelPlatformPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type PlatformPolicyAdapter struct {
	id        *krm.BinaryAuthorizationPlatformPolicyIdentity
	gcpClient *api.Service
	desired   *krm.BinaryAuthorizationPlatformPolicy
	actual    *api.PlatformPolicy
}

var _ directbase.Adapter = &PlatformPolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *PlatformPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting PlatformPolicy", "name", a.id)

	platformpolicypb, err := a.gcpClient.Projects.Platforms.Policies.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting PlatformPolicy %q: %w", a.id, err)
	}

	a.actual = platformpolicypb
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PlatformPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PlatformPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BinaryAuthorizationPlatformPolicySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	created, err := a.gcpClient.Projects.Platforms.Policies.Create(fmt.Sprintf("projects/%s/platforms/%s", a.id.Project, a.id.Platform), resource).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("PlatformPolicy %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created PlatformPolicy", "name", a.id)

	status := &krm.BinaryAuthorizationPlatformPolicyStatus{}
	status.ObservedState = BinaryAuthorizationPlatformPolicyObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *PlatformPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PlatformPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := BinaryAuthorizationPlatformPolicySpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	paths := make(sets.Set[string])
	// Option 2: manually add all mutable fields.
	// TODO(contributor): If choosing this option, remove the "Option 1" code.
	{
		// manually check fields
		if !reflect.DeepEqual(a.desired.Spec.Description, a.actual.Description) {
			paths.Insert("description")
		}
	}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)

		var err error
		updated, err = a.gcpClient.Projects.Platforms.Policies.ReplacePlatformPolicy(a.id.String(), desiredPb).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("PlatformPolicy %s waiting update: %w", a.id, err)
		}
		log.V(2).Info("successfully updated PlatformPolicy", "name", a.id)
	}

	status := &krm.BinaryAuthorizationPlatformPolicyStatus{}
	status.ObservedState = BinaryAuthorizationPlatformPolicyObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *PlatformPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BinaryAuthorizationPlatformPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BinaryAuthorizationPlatformPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Platform = a.id.Platform
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Policy)
	u.SetGroupVersionKind(krm.BinaryAuthorizationPlatformPolicyGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *PlatformPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PlatformPolicy", "name", a.id)

	_, err := a.gcpClient.Projects.Platforms.Policies.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent PlatformPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting PlatformPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted PlatformPolicy", "name", a.id)

	return true, nil
}
