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

package firebasehosting

import (
	"context"
	"fmt"
	"strings"

	api "google.golang.org/api/firebasehosting/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/firebasehosting/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.FirebaseHostingSiteGVK, NewModel)
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
	obj := &krm.FirebaseHostingSite{}
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
	id := idBase.(*krm.FirebaseHostingSiteIdentity)

	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, err
	}

	return &SiteAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.FirebaseHostingSiteIdentity{}
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
	return &SiteAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type SiteAdapter struct {
	id        *krm.FirebaseHostingSiteIdentity
	gcpClient *api.Service
	desired   *krm.FirebaseHostingSite
	actual    *api.Site
}

var _ directbase.Adapter = &SiteAdapter{}

func (a *SiteAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting FirebaseHostingSite", "name", a.id)

	site, err := a.gcpClient.Projects.Sites.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting FirebaseHostingSite %s: %w", a.id, err)
	}

	a.actual = site
	return true, nil
}

func (a *SiteAdapter) Create(ctx context.Context, op *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating FirebaseHostingSite", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := FirebaseHostingSiteSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	created, err := a.gcpClient.Projects.Sites.Create(a.id.ParentString(), desired).SiteId(a.id.Site).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating FirebaseHostingSite %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created FirebaseHostingSite", "name", a.id)

	return a.updateStatus(ctx, op, created)
}

func (a *SiteAdapter) Update(ctx context.Context, op *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating FirebaseHostingSite", "name", a.id)

	mapCtx := &direct.MapContext{}
	desired := FirebaseHostingSiteSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	report := &structuredreporting.Diff{Object: op.GetUnstructured()}
	var paths []string
	if desired.AppId != a.actual.AppId {
		report.AddField("app_id", a.actual.AppId, desired.AppId)
		paths = append(paths, "appId")
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		return a.updateStatus(ctx, op, a.actual)
	}

	structuredreporting.ReportDiff(ctx, report)

	patch := &api.Site{
		AppId: desired.AppId,
	}
	if patch.AppId == "" {
		patch.ForceSendFields = []string{"AppId"}
	}

	updateMask := strings.Join(paths, ",")
	updated, err := a.gcpClient.Projects.Sites.Patch(a.id.String(), patch).UpdateMask(updateMask).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("updating FirebaseHostingSite %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated FirebaseHostingSite", "name", a.id)

	return a.updateStatus(ctx, op, updated)
}

func (a *SiteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.FirebaseHostingSite{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(FirebaseHostingSiteSpec_FromAPI(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = refs.ProjectRef{External: a.id.Project}
	obj.Spec.ResourceID = direct.LazyPtr(a.id.Site)
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.Site)
	u.SetGroupVersionKind(krm.FirebaseHostingSiteGVK)

	u.Object = uObj
	return u, nil
}

func (a *SiteAdapter) Delete(ctx context.Context, op *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting FirebaseHostingSite", "name", a.id)

	_, err := a.gcpClient.Projects.Sites.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting FirebaseHostingSite %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted FirebaseHostingSite", "name", a.id)

	return true, nil
}

func (a *SiteAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.Site) error {
	mapCtx := &direct.MapContext{}
	status := FirebaseHostingSiteStatus_FromAPI(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}
