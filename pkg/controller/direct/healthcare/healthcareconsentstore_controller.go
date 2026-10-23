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

package healthcare

import (
	"context"
	"fmt"
	"strings"

	api "google.golang.org/api/healthcare/v1"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/healthcare/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.HealthcareConsentStoreGVK, NewConsentStoreModel)
}

func NewConsentStoreModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &consentStoreModel{config: *config}, nil
}

var _ directbase.Model = &consentStoreModel{}

type consentStoreModel struct {
	config config.ControllerConfig
}

func (m *consentStoreModel) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building healthcare client: %w", err)
	}
	return gcpClient, nil
}

func (m *consentStoreModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.HealthcareConsentStore{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.HealthcareConsentStoreIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &ConsentStoreAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *consentStoreModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	id := &krm.HealthcareConsentStoreIdentity{}
	if err := id.FromExternal(url); err != nil {
		return nil, nil
	}
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ConsentStoreAdapter{
		id:        id,
		gcpClient: gcpClient,
	}, nil
}

type ConsentStoreAdapter struct {
	id        *krm.HealthcareConsentStoreIdentity
	gcpClient *api.Service
	desired   *krm.HealthcareConsentStore
	actual    *api.ConsentStore
}

var _ directbase.Adapter = &ConsentStoreAdapter{}

func (a *ConsentStoreAdapter) buildDesired(ctx context.Context, u *unstructured.Unstructured) (*api.ConsentStore, error) {
	mapCtx := &direct.MapContext{}
	desired := HealthcareConsentStoreSpec_ToAPI(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desired.Labels = label.GCPLabels(u)
	return desired, nil
}

func (a *ConsentStoreAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting HealthcareConsentStore", "name", a.id.String())

	actual, err := a.gcpClient.Projects.Locations.Datasets.ConsentStores.Get(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting HealthcareConsentStore %q from GCP: %w", a.id.String(), err)
	}
	a.actual = actual
	return true, nil
}

func (a *ConsentStoreAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating HealthcareConsentStore", "name", a.id.String())

	u := createOp.GetUnstructured()
	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	call := a.gcpClient.Projects.Locations.Datasets.ConsentStores.Create(a.id.ParentString(), desired)
	if a.id.ConsentStore != "" {
		call.ConsentStoreId(a.id.ConsentStore)
	}
	created, err := call.Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating HealthcareConsentStore %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully created HealthcareConsentStore", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func (a *ConsentStoreAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating HealthcareConsentStore", "name", a.id.String())

	u := updateOp.GetUnstructured()
	desired, err := a.buildDesired(ctx, u)
	if err != nil {
		return err
	}

	diffResults, updatePaths, err := compareConsentStore(ctx, a.actual, desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if !diffResults.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
	} else {
		log.V(2).Info("fields need update", "name", a.id.String())

		diffResults.Object = u
		structuredreporting.ReportDiff(ctx, diffResults)

		patch := &api.ConsentStore{}
		for _, path := range updatePaths {
			switch path {
			case "defaultConsentTtl":
				patch.DefaultConsentTtl = desired.DefaultConsentTtl
				if patch.DefaultConsentTtl == "" {
					patch.ForceSendFields = append(patch.ForceSendFields, "DefaultConsentTtl")
				}
			case "enableConsentCreateOnUpdate":
				patch.EnableConsentCreateOnUpdate = desired.EnableConsentCreateOnUpdate
			case "labels":
				patch.Labels = desired.Labels
				if len(patch.Labels) == 0 {
					patch.NullFields = append(patch.NullFields, "Labels")
				}
			}
		}

		updateMask := strings.Join(updatePaths, ",")
		updated, err := a.gcpClient.Projects.Locations.Datasets.ConsentStores.Patch(a.id.String(), patch).
			UpdateMask(updateMask).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating HealthcareConsentStore %q: %w", a.id.String(), err)
		}
		log.V(2).Info("successfully updated HealthcareConsentStore", "name", a.id.String())
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareConsentStore(ctx context.Context, actual, desired *api.ConsentStore) (*structuredreporting.Diff, []string, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, HealthcareConsentStoreSpec_FromAPI, HealthcareConsentStoreSpec_ToAPI)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Labels = actual.Labels

	clonedDesired := common.DeepCopy(desired)

	return diffs.GoogleAPI.Diff(ctx, maskedActual, clonedDesired)
}

func (a *ConsentStoreAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.ConsentStore) error {
	mapCtx := &direct.MapContext{}
	status := HealthcareConsentStoreStatus_FromAPI(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ConsentStoreAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.HealthcareConsentStore{}
	mapCtx := &direct.MapContext{}
	spec := HealthcareConsentStoreSpec_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	spec.Dataset = a.id.ParentString()
	spec.ResourceID = direct.LazyPtr(a.id.ConsentStore)
	obj.Spec = *spec
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ConsentStore)
	u.SetGroupVersionKind(krm.HealthcareConsentStoreGVK)
	if len(a.actual.Labels) > 0 {
		u.SetLabels(a.actual.Labels)
	}

	return u, nil
}

func (a *ConsentStoreAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting HealthcareConsentStore", "name", a.id.String())

	_, err := a.gcpClient.Projects.Locations.Datasets.ConsentStores.Delete(a.id.String()).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return true, nil
		}
		return false, fmt.Errorf("deleting HealthcareConsentStore %q: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted HealthcareConsentStore", "name", a.id.String())

	return true, nil
}
