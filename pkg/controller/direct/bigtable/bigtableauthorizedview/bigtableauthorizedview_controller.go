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

package bigtableauthorizedview

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/bigtable"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.BigtableAuthorizedViewGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config config.ControllerConfig
}

func (m *model) client(ctx context.Context, projectID, instanceID string) (*gcp.AdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building bigtable admin client options: %w", err)
	}
	gcpClient, err := gcp.NewAdminClient(ctx, projectID, instanceID, opts...)
	if err != nil {
		return nil, fmt.Errorf("building bigtable admin client: %w", err)
	}
	return gcpClient, nil
}

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableAuthorizedView{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAuthorizedViewIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.ProjectID(), id.InstanceID())
	if err != nil {
		return nil, err
	}
	return &Adapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type Adapter struct {
	id        *krm.AuthorizedViewIdentity
	gcpClient *gcp.AdminClient
	desired   *krm.BigtableAuthorizedView
	actual    *gcp.AuthorizedViewInfo
}

var _ directbase.Adapter = &Adapter{}

func (a *Adapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableAuthorizedView", "name", a.id)

	info, err := a.gcpClient.AuthorizedViewInfo(ctx, a.id.TableID(), a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableAuthorizedView %q: %w", a.id, err)
	}

	a.actual = info
	return true, nil
}

func (a *Adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableAuthorizedView", "name", a.id)

	var subsetView *gcp.SubsetViewConf
	if a.desired.Spec.SubsetView != nil {
		subsetView = &gcp.SubsetViewConf{
			RowPrefixes: a.desired.Spec.SubsetView.RowPrefixes,
		}
	}

	var deletionProtection gcp.DeletionProtection
	if a.desired.Spec.DeletionProtection != nil {
		if *a.desired.Spec.DeletionProtection {
			deletionProtection = gcp.Protected
		} else {
			deletionProtection = gcp.Unprotected
		}
	}

	conf := gcp.AuthorizedViewConf{
		TableID:            a.id.TableID(),
		AuthorizedViewID:   a.id.ID(),
		AuthorizedView:     subsetView,
		DeletionProtection: deletionProtection,
	}

	err := a.gcpClient.CreateAuthorizedView(ctx, &conf)
	if err != nil {
		return fmt.Errorf("creating BigtableAuthorizedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created BigtableAuthorizedView", "name", a.id)

	status := &krm.BigtableAuthorizedViewStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = &krm.BigtableAuthorizedViewObservedState{}

	if err := createOp.UpdateStatus(ctx, status, nil); err != nil {
		return err
	}

	// Write resourceID into spec.
	if err := unstructured.SetNestedField(createOp.GetUnstructured().Object, a.id.ID(), "spec", "resourceID"); err != nil {
		return fmt.Errorf("error setting spec.resourceID: %w", err)
	}
	return nil
}

func (a *Adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableAuthorizedView", "name", a.id)

	var subsetView *gcp.SubsetViewConf
	if a.desired.Spec.SubsetView != nil {
		subsetView = &gcp.SubsetViewConf{
			RowPrefixes: a.desired.Spec.SubsetView.RowPrefixes,
		}
	}

	var deletionProtection gcp.DeletionProtection
	if a.desired.Spec.DeletionProtection != nil {
		if *a.desired.Spec.DeletionProtection {
			deletionProtection = gcp.Protected
		} else {
			deletionProtection = gcp.Unprotected
		}
	}

	conf := gcp.AuthorizedViewConf{
		TableID:            a.id.TableID(),
		AuthorizedViewID:   a.id.ID(),
		AuthorizedView:     subsetView,
		DeletionProtection: deletionProtection,
	}

	var actualRowPrefixes [][]byte
	if a.actual.AuthorizedView != nil {
		if sv, ok := a.actual.AuthorizedView.(*gcp.SubsetViewInfo); ok {
			actualRowPrefixes = sv.RowPrefixes
		}
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	needsUpdate := false

	if subsetView != nil && !reflect.DeepEqual(subsetView.RowPrefixes, actualRowPrefixes) {
		report.AddField("subset_view.row_prefixes", actualRowPrefixes, subsetView.RowPrefixes)
		needsUpdate = true
	}

	if deletionProtection != a.actual.DeletionProtection {
		report.AddField("deletion_protection", a.actual.DeletionProtection, deletionProtection)
		needsUpdate = true
	}

	if !needsUpdate {
		log.V(2).Info("no field needs update", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	updateConf := gcp.UpdateAuthorizedViewConf{
		AuthorizedViewConf: conf,
	}

	err := a.gcpClient.UpdateAuthorizedView(ctx, updateConf)
	if err != nil {
		return fmt.Errorf("updating BigtableAuthorizedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully updated BigtableAuthorizedView", "name", a.id)

	status := &krm.BigtableAuthorizedViewStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = &krm.BigtableAuthorizedViewObservedState{}

	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *Adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableAuthorizedView", "name", a.id)

	err := a.gcpClient.DeleteAuthorizedView(ctx, a.id.TableID(), a.id.ID())
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting BigtableAuthorizedView %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted BigtableAuthorizedView", "name", a.id)
	return true, nil
}

func (a *Adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigtableAuthorizedView{}
	obj.Spec.SubsetView = &krm.AuthorizedView_SubsetView{}
	if a.actual.AuthorizedView != nil {
		if sv, ok := a.actual.AuthorizedView.(*gcp.SubsetViewInfo); ok {
			obj.Spec.SubsetView.RowPrefixes = sv.RowPrefixes
		}
	}
	if a.actual.DeletionProtection == gcp.Protected {
		obj.Spec.DeletionProtection = direct.LazyPtr(true)
	} else if a.actual.DeletionProtection == gcp.Unprotected {
		obj.Spec.DeletionProtection = direct.LazyPtr(false)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("error converting to unstructured: %w", err)
	}

	u.Object = uObj
	return u, nil
}
