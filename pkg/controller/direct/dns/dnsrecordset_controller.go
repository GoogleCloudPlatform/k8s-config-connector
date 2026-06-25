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

package dns

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/dns/v1"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DNSRecordSetGVK, NewRecordSetModel)
}

func NewRecordSetModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &recordSetModel{config: config}, nil
}

var _ directbase.Model = &recordSetModel{}

type recordSetModel struct {
	config *config.ControllerConfig
}

func (m *recordSetModel) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DNS client: %w", err)
	}
	return gcpClient, err
}

func (m *recordSetModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DNSRecordSet{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.DNSRecordSetIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format here, so we follow the pattern in the skill.
	mapCtx := &direct.MapContext{}
	desired := DNSRecordSetSpec_ToAPI(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get DNS GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &DNSRecordSetAdapter{
		id:         id,
		recordType: obj.Spec.Type,
		gcpClient:  gcpClient,
		desired:    desired,
	}, nil
}

func (m *recordSetModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DNSRecordSetAdapter struct {
	id         *krm.DNSRecordSetIdentity
	recordType string
	gcpClient  *api.Service
	desired    *api.ResourceRecordSet
	actual     *api.ResourceRecordSet
}

var _ directbase.Adapter = &DNSRecordSetAdapter{}

func (a *DNSRecordSetAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DNSRecordSet", "name", a.id, "type", a.recordType)

	resp, err := a.gcpClient.ResourceRecordSets.List(a.id.Project, a.id.ManagedZone).Name(a.id.Name).Type(a.recordType).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("listing DNSRecordSet %q (type %s): %w", a.id, a.recordType, err)
	}

	if len(resp.Rrsets) == 0 {
		return false, nil
	}

	a.actual = resp.Rrsets[0]
	return true, nil
}

func (a *DNSRecordSetAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DNSRecordSet", "name", a.id, "type", a.recordType)

	// We use the Changes API because it's what the legacy TF provider uses, and keeping
	// the HTTP traffic aligned prevents overwhelming golden http.log diffs during cross-comparison.
	change := &api.Change{
		Additions: []*api.ResourceRecordSet{a.desired},
	}
	_, err := a.gcpClient.Changes.Create(a.id.Project, a.id.ManagedZone, change).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating DNSRecordSet %s (type %s) via changes: %w", a.id, a.recordType, err)
	}

	log.V(2).Info("successfully created DNSRecordSet via changes", "name", a.id, "type", a.recordType)

	// Fetch the created resource to populate the status properly
	found, err := a.Find(ctx)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("created DNSRecordSet but could not find it afterwards")
	}

	return a.updateStatus(ctx, createOp, a.actual)
}

func (a *DNSRecordSetAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DNSRecordSet", "name", a.id, "type", a.recordType)

	diffResults, err := compareRecordSet(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if !diffResults.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id, "type", a.recordType)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "type", a.recordType)

		diffResults.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffResults)

		// We use the Changes API because it's what the legacy TF provider uses, and keeping
		// the HTTP traffic aligned prevents overwhelming golden http.log diffs during cross-comparison.
		change := &api.Change{
			Deletions: []*api.ResourceRecordSet{a.actual},
			Additions: []*api.ResourceRecordSet{a.desired},
		}
		_, err := a.gcpClient.Changes.Create(a.id.Project, a.id.ManagedZone, change).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating DNSRecordSet %s (type %s) via changes: %w", a.id, a.recordType, err)
		}
		log.V(2).Info("successfully updated DNSRecordSet via changes", "name", a.id, "type", a.recordType)

		// Commented out standard PATCH for future reference / option to switch to PATCH later:
		/*
			patched, err := a.gcpClient.ResourceRecordSets.Patch(a.id.Project, a.id.ManagedZone, a.id.Name, a.recordType, a.desired).Context(ctx).Do()
			if err != nil {
				return fmt.Errorf("updating DNSRecordSet %s (type %s): %w", a.id, a.recordType, err)
			}
			log.V(2).Info("successfully patched DNSRecordSet", "name", a.id, "type", a.recordType)
			latest = patched
		*/

		// Fetch the updated resource to populate the status properly
		found, err := a.Find(ctx)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("updated DNSRecordSet but could not find it afterwards")
		}
		latest = a.actual
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareRecordSet(ctx context.Context, actual, desired *api.ResourceRecordSet) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DNSRecordSetSpec_FromAPI, DNSRecordSetSpec_ToAPI)
	if err != nil {
		return nil, err
	}

	desired = common.DeepCopy(desired)

	populateDefaults := func(obj *api.ResourceRecordSet) {
		obj.Kind = "dns#resourceRecordSet"
		if obj.SignatureRrdatas == nil {
			obj.SignatureRrdatas = []string{}
		}
	}
	populateDefaults(maskedActual)
	populateDefaults(desired)

	diffs, _, err := diffs.GoogleAPI.Diff(ctx, maskedActual, desired)
	return diffs, err
}

func (a *DNSRecordSetAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.ResourceRecordSet) error {
	mapCtx := &direct.MapContext{}
	status := DNSRecordSetStatus_FromAPI(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *DNSRecordSetAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DNSRecordSet", "name", a.id, "type", a.recordType)

	// Since we need to match the existing data exactly to delete, we should first read actual
	// if it is not already loaded (e.g., if Find was not run, or to ensure we have the correct actual).
	found, err := a.Find(ctx)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}

	// We use the Changes API because it's what the legacy TF provider uses, and keeping
	// the HTTP traffic aligned prevents overwhelming golden http.log diffs during cross-comparison.
	change := &api.Change{
		Deletions: []*api.ResourceRecordSet{a.actual},
	}
	_, err = a.gcpClient.Changes.Create(a.id.Project, a.id.ManagedZone, change).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting DNSRecordSet %s (type %s) via changes: %w", a.id, a.recordType, err)
	}

	log.V(2).Info("successfully deleted DNSRecordSet via changes", "name", a.id, "type", a.recordType)
	return true, nil
}

func (a *DNSRecordSetAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DNSRecordSet{}
	mapCtx := &direct.MapContext{}
	spec := DNSRecordSetSpec_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	spec.ManagedZoneRef = krm.DNSManagedZoneRef{
		External: a.id.ManagedZone,
	}
	obj.Spec = *spec
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Name)
	u.SetGroupVersionKind(krm.DNSRecordSetGVK)

	return u, nil
}
