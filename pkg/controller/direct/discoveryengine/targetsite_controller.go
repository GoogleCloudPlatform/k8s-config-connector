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

// +tool:controller
// proto.service: google.cloud.discoveryengine.v1.SiteSearchEngineService
// proto.message: google.cloud.discoveryengine.v1.TargetSite
// crd.type: DiscoveryEngineDataStoreTargetSite
// crd.version: v1alpha1

package discoveryengine

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.DiscoveryEngineDataStoreTargetSiteGVK, NewTargetSiteModel)
}

func NewTargetSiteModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &targetSiteModel{config: *config}, nil
}

var _ directbase.Model = &targetSiteModel{}

type targetSiteModel struct {
	config config.ControllerConfig
}

func (m *targetSiteModel) client(ctx context.Context, projectID string) (*gcp.SiteSearchEngineClient, error) {
	var opts []option.ClientOption

	config := m.config

	// Workaround for an unusual behaviour (bug?):
	//  the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = projectID
	}

	opts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := gcp.NewSiteSearchEngineRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building discoveryengine sitesearchengine client: %w", err)
	}

	return gcpClient, err
}

func (m *targetSiteModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DiscoveryEngineDataStoreTargetSite{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	dataStoreID, id, err := krm.NewTargetSiteIdentityFromObject(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := DiscoveryEngineDataStoreTargetSiteSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, dataStoreID.ProjectID)
	if err != nil {
		return nil, err
	}

	return &targetSiteAdapter{
		gcpClient:   gcpClient,
		dataStoreID: dataStoreID,
		id:          id,
		desired:     desired,
	}, nil
}

func (m *targetSiteModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Not implemented
	return nil, nil
}

type targetSiteAdapter struct {
	gcpClient   *gcp.SiteSearchEngineClient
	dataStoreID *krm.DiscoveryEngineDataStoreID
	id          *krm.TargetSiteIdentity
	desired     *pb.TargetSite
	actual      *pb.TargetSite
}

var _ directbase.Adapter = &targetSiteAdapter{}

func (a *targetSiteAdapter) Find(ctx context.Context) (bool, error) {
	if a.id == nil {
		return false, nil
	}
	log := klog.FromContext(ctx)
	log.V(2).Info("getting discoveryengine targetsite", "name", a.id)

	req := &pb.GetTargetSiteRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetTargetSite(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting discoveryengine targetsite %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *targetSiteAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating discoveryengine targetsite", "dataStoreID", a.dataStoreID)

	desired := direct.ProtoClone(a.desired)

	req := &pb.CreateTargetSiteRequest{
		Parent:     a.dataStoreID.String() + "/siteSearchEngine",
		TargetSite: desired,
	}
	op, err := a.gcpClient.CreateTargetSite(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine targetsite: %w", err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("discoveryengine targetsite waiting creation: %w", err)
	}
	log.V(2).Info("successfully created discoveryengine targetsite in gcp")

	// Since TargetSite resource ID is server-generated, we need to parse it from the created object
	id, err := krm.ParseTargetSiteExternal(created.Name)
	if err != nil {
		return fmt.Errorf("parsing created targetsite name %q: %w", created.Name, err)
	}
	a.id = id

	status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *targetSiteAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine targetsite", "name", a.id)

	desired := direct.ProtoClone(a.desired)
	desired.Name = a.id.String()

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	// Most fields in TargetSite are immutable or input-only.
	// We check for drifts and report them, but only update if there are mutable fields.
	// Based on the API, TargetSite doesn't seem to have many mutable fields.
	// If we find that no fields can be updated, we just return nil to avoid UpdateFailed loop.

	needsUpdate := false
	if !reflect.DeepEqual(a.desired.Type, a.actual.Type) {
		report.AddField("type", a.actual.Type, a.desired.Type)
		needsUpdate = true
	}

	// provided_uri_pattern and exact_match are InputOnly, so we check for drift but don't include in updateMask if they are immutable.
	if !reflect.DeepEqual(a.desired.ProvidedUriPattern, a.actual.ProvidedUriPattern) {
		report.AddField("provided_uri_pattern", a.actual.ProvidedUriPattern, a.desired.ProvidedUriPattern)
		// Assuming immutable
	}
	if !reflect.DeepEqual(a.desired.ExactMatch, a.actual.ExactMatch) {
		report.AddField("exact_match", a.actual.ExactMatch, a.desired.ExactMatch)
		// Assuming immutable
	}

	if !needsUpdate {
		log.V(2).Info("no field needs update for discoveryengine targetsite", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, report)

	req := &pb.UpdateTargetSiteRequest{
		TargetSite: desired,
	}
	op, err := a.gcpClient.UpdateTargetSite(ctx, req)
	if err != nil {
		return fmt.Errorf("updating discoveryengine targetsite %s: %w", a.id.String(), err)
	}
	updated, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("discoveryengine targetsite %s waiting update: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated discoveryengine targetsite", "name", a.id)

	status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *targetSiteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineDataStoreTargetSite{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(DiscoveryEngineDataStoreTargetSiteSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{
		External: a.dataStoreID.String(),
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.TargetSite)
	u.SetGroupVersionKind(krm.DiscoveryEngineDataStoreTargetSiteGVK)

	return u, nil
}

// Delete implements the Adapter interface.
func (a *targetSiteAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine targetsite", "name", a.id)

	req := &pb.DeleteTargetSiteRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTargetSite(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting discoveryengine targetsite %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine targetsite", "name", a.id)

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of discoveryengine targetsite %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
