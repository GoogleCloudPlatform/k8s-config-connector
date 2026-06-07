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

package targetsite

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	gcp "cloud.google.com/go/discoveryengine/apiv1"
	pb "cloud.google.com/go/discoveryengine/apiv1/discoveryenginepb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/discoveryengine/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/discoveryengine"
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

	dataStoreLink, id, err := krm.NewTargetSiteIdentityFromObject(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desired := discoveryengine.DiscoveryEngineDataStoreTargetSiteSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	gcpClient, err := m.client(ctx, dataStoreLink.ProjectID)
	if err != nil {
		return nil, err
	}

	return &targetSiteAdapter{
		gcpClient: gcpClient,
		dataStore: dataStoreLink,
		id:        id,
		desired:   desired,
	}, nil
}

func (m *targetSiteModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	log := klog.FromContext(ctx)
	if strings.HasPrefix(url, "//discoveryengine.googleapis.com/") {
		id, err := krm.ParseTargetSiteExternal(url)
		if err != nil {
			log.V(2).Error(err, "url did not match DiscoveryEngineDataStoreTargetSite format", "url", url)
		} else {
			gcpClient, err := m.client(ctx, id.ProjectID)
			if err != nil {
				return nil, err
			}
			return &targetSiteAdapter{
				gcpClient: gcpClient,
				dataStore: id.DiscoveryEngineDataStoreID,
				id:        id,
			}, nil
		}
	}
	return nil, nil
}

type targetSiteAdapter struct {
	gcpClient *gcp.SiteSearchEngineClient
	dataStore *krm.DiscoveryEngineDataStoreID
	id        *krm.TargetSiteIdentity
	desired   *pb.TargetSite
	actual    *pb.TargetSite
}

var _ directbase.Adapter = &targetSiteAdapter{}

func (a *targetSiteAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	if a.id == nil {
		log.V(2).Info("discoveryengine targetsite has no identity (not yet created)")
		return false, nil
	}
	log.V(2).Info("getting discoveryengine targetsite", "name", a.id.String())

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
	log.V(2).Info("creating discoveryengine targetsite", "datastore", a.dataStore.String())

	desired := proto.Clone(a.desired).(*pb.TargetSite)

	parent := a.dataStore.String() + "/siteSearchEngine"
	req := &pb.CreateTargetSiteRequest{
		Parent:     parent,
		TargetSite: desired,
	}
	op, err := a.gcpClient.CreateTargetSite(ctx, req)
	if err != nil {
		return fmt.Errorf("creating discoveryengine targetsite parent=%s: %w", parent, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("discoveryengine targetsite waiting creation: %w", err)
	}
	log.V(2).Info("successfully created discoveryengine targetsite in gcp", "name", created.GetName())

	id, err := krm.ParseTargetSiteExternal(created.GetName())
	if err != nil {
		return fmt.Errorf("parsing created targetsite name %q: %w", created.GetName(), err)
	}
	a.id = id

	status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = discoveryengine.DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *targetSiteAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating discoveryengine targetsite", "name", a.id.String())

	desired := proto.Clone(a.desired).(*pb.TargetSite)
	desired.Name = a.id.String()

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}

	hasChanges := false
	if !reflect.DeepEqual(a.desired.Type, a.actual.Type) {
		report.AddField("type", a.actual.Type.String(), a.desired.Type.String())
		hasChanges = true
	}
	if !reflect.DeepEqual(a.desired.ExactMatch, a.actual.ExactMatch) {
		report.AddField("exact_match", fmt.Sprintf("%v", a.actual.ExactMatch), fmt.Sprintf("%v", a.desired.ExactMatch))
		hasChanges = true
	}
	if !reflect.DeepEqual(a.desired.ProvidedUriPattern, a.actual.ProvidedUriPattern) {
		report.AddField("provided_uri_pattern", a.actual.ProvidedUriPattern, a.desired.ProvidedUriPattern)
		hasChanges = true
	}

	if !hasChanges {
		log.V(2).Info("no field needs update", "name", a.id.String())
		status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
		mapCtx := &direct.MapContext{}
		status.ObservedState = discoveryengine.DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, a.actual)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		status.ExternalRef = direct.PtrTo(a.id.String())
		return updateOp.UpdateStatus(ctx, status, nil)
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
	log.V(2).Info("successfully updated discoveryengine targetsite", "name", a.id.String())

	status := &krm.DiscoveryEngineDataStoreTargetSiteStatus{}
	mapCtx := &direct.MapContext{}
	status.ObservedState = discoveryengine.DiscoveryEngineDataStoreTargetSiteObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *targetSiteAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)

	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DiscoveryEngineDataStoreTargetSite{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(discoveryengine.DiscoveryEngineDataStoreTargetSiteSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.DataStoreRef = &krm.DiscoveryEngineDataStoreRef{External: a.dataStore.String()}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.TargetSite)
	u.SetGroupVersionKind(krm.DiscoveryEngineDataStoreTargetSiteGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete implements the Adapter interface.
func (a *targetSiteAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting discoveryengine targetsite", "name", a.id.String())

	req := &pb.DeleteTargetSiteRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTargetSite(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting discoveryengine targetsite %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted discoveryengine targetsite", "name", a.id.String())

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting for deletion of discoveryengine targetsite %s: %w", a.id.String(), err)
		}
	}
	return true, nil
}
