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

package firebase

import (
	"context"
	"fmt"
	"time"

	api "google.golang.org/api/firebase/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firebase/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.FirebaseProjectGVK, NewModel, registry.CannotBeDeleted())
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.FirebaseProject{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idBase, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := idBase.(*krm.FirebaseProjectIdentity)

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &projectAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.FirebaseProjectIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &projectAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type projectAdapter struct {
	id        *krm.FirebaseProjectIdentity
	gcpClient *api.Service
	desired   *krm.FirebaseProject
	actual    *api.FirebaseProject
}

var _ directbase.Adapter = &projectAdapter{}

func (a *projectAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FirebaseProject", "name", a.id)

	project, err := a.gcpClient.Projects.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirebaseProject %s: %w", a.id, err)
	}

	a.actual = project
	return true, nil
}

func (a *projectAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirebaseProject", "name", a.id)

	req := &api.AddFirebaseRequest{}
	lro, err := a.gcpClient.Projects.AddFirebase(a.id.String(), req).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating FirebaseProject %s: %w", a.id, err)
	}

	if err := a.waitForOp(ctx, lro); err != nil {
		return fmt.Errorf("waiting for creation of FirebaseProject %s: %w", a.id, err)
	}

	latest, err := a.gcpClient.Projects.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("getting created FirebaseProject %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created FirebaseProject", "name", a.id)

	return a.updateStatus(ctx, op, latest)
}

func (a *projectAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirebaseProject", "name", a.id)

	// FirebaseProject in KRM Spec has no mutable fields (only projectRef and resourceID).
	// GCP FirebaseProject resources are immutable through KCC's spec.
	return a.updateStatus(ctx, op, a.actual)
}

func (a *projectAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.FirebaseProject{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FirebaseProjectSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = refs.ProjectRef{External: a.id.Project}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Project)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Project)
	u.SetGroupVersionKind(krm.FirebaseProjectGVK)

	u.Object = uObj
	return u, nil
}

func (a *projectAdapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirebaseProject (no-op as Firebase cannot be deleted)", "name", a.id)

	// Firebase Projects cannot be deleted from Google Cloud.
	// Returning true allows removal from K8s without error.
	return true, nil
}

func (a *projectAdapter) waitForOp(ctx context.Context, op *api.Operation) error {
	if op.Done {
		if op.Error != nil {
			return fmt.Errorf("operation failed: %v", op.Error.Message)
		}
		return nil
	}
	return common.WaitForDoneOrTimeout(ctx, 2*time.Second, func() (bool, error) {
		current, err := a.gcpClient.Operations.Get(op.Name).Context(ctx).Do()
		if err != nil {
			return false, fmt.Errorf("getting operation status of %q: %w", op.Name, err)
		}
		if current.Done {
			if current.Error != nil {
				return true, fmt.Errorf("operation %q completed with error: %s", op.Name, current.Error.Message)
			}
			return true, nil
		}
		return false, nil
	})
}

func (a *projectAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.FirebaseProject) error {
	mapCtx := &direct.MapContext{}
	status := FirebaseProjectStatus_FromAPI(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
