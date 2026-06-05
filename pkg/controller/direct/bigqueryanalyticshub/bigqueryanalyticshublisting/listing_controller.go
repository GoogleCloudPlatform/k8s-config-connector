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

package bigqueryanalyticshublisting

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	gcp "cloud.google.com/go/bigquery/analyticshub/apiv1"
	bigqueryanalyticshubpb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

const (
	listingCtrlName = "bigqueryanalyticshub-listing-controller"
)

func init() {
	registry.RegisterModel(krm.BigQueryAnalyticsHubListingGVK, NewListingModel)
}

func NewListingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelListing{config: *config}, nil
}

var _ directbase.Model = &modelListing{}

type modelListing struct {
	config config.ControllerConfig
}

func (m *modelListing) client(ctx context.Context, project string) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions(config.WithDefaultQuotaProject(project))
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Listing client: %w", err)
	}
	return gcpClient, err
}

func (m *modelListing) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigQueryAnalyticsHubListing{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}

	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	desired := &bigqueryanalyticshubpb.Listing{}
	mapCtx := &direct.MapContext{}
	desired = BigQueryAnalyticsHubListingSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Get bigqueryanalyticshub GCP client
	gcpClient, err := m.client(ctx, id.(*krm.BigQueryAnalyticsHubListingIdentity).Project)
	if err != nil {
		return nil, err
	}
	return &ListingAdapter{
		id:        id.(*krm.BigQueryAnalyticsHubListingIdentity),
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *modelListing) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ListingAdapter struct {
	id        *krm.BigQueryAnalyticsHubListingIdentity
	gcpClient *gcp.Client
	desired   *bigqueryanalyticshubpb.Listing
	actual    *bigqueryanalyticshubpb.Listing
}

var _ directbase.Adapter = &ListingAdapter{}

func (a *ListingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Listing", "name", a.id.String())

	req := &bigqueryanalyticshubpb.GetListingRequest{Name: a.id.String()}
	listingpb, err := a.gcpClient.GetListing(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Listing %q: %w", a.id.String(), err)
	}

	a.actual = listingpb
	return true, nil
}

func (a *ListingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Listing", "name", a.id.String())

	parent := fmt.Sprintf("projects/%s/locations/%s/dataExchanges/%s", a.id.Project, a.id.Location, a.id.DataExchange)
	req := &bigqueryanalyticshubpb.CreateListingRequest{
		Parent:    parent,
		Listing:   a.desired,
		ListingId: a.id.Listing,
	}
	created, err := a.gcpClient.CreateListing(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Listing %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created Listing", "name", a.id.String())

	return a.updateStatus(ctx, createOp, created)
}

func compareListing(ctx context.Context, actual, desired *bigqueryanalyticshubpb.Listing) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	var maskedActual *bigqueryanalyticshubpb.Listing
	{
		mapCtx := &direct.MapContext{}
		spec := BigQueryAnalyticsHubListingSpec_FromProto(mapCtx, actual)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
		maskedActual = BigQueryAnalyticsHubListingSpec_ToProto(mapCtx, spec)
		if mapCtx.Err() != nil {
			return nil, nil, mapCtx.Err()
		}
	}
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *ListingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Listing", "name", a.id.String())

	diffs, updateMask, err := compareListing(ctx, a.actual, a.desired)
	if err != nil {
		return fmt.Errorf("comparing Listing %s: %w", a.id.String(), err)
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id.String())
		return nil
	}

	// Update the Name field in desired before sending
	desiredCopy := &bigqueryanalyticshubpb.Listing{}
	if a.desired != nil {
		desiredCopy = a.desired
	}
	desiredCopy.Name = a.id.String()

	req := &bigqueryanalyticshubpb.UpdateListingRequest{
		UpdateMask: updateMask,
		Listing:    desiredCopy,
	}
	updated, err := a.gcpClient.UpdateListing(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Listing %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully updated Listing", "name", a.id.String())

	return a.updateStatus(ctx, updateOp, updated)
}

func (a *ListingAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *bigqueryanalyticshubpb.Listing) error {
	mapCtx := &direct.MapContext{}
	status := &krm.BigQueryAnalyticsHubListingStatus{}
	status.ObservedState = BigQueryAnalyticsHubListingObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := a.id.String()
	status.ExternalRef = &externalRef
	return op.UpdateStatus(ctx, status, nil)
}

func (a *ListingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.BigQueryAnalyticsHubListing{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(BigQueryAnalyticsHubListingSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Project}
	obj.Spec.Location = a.id.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *ListingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Listing", "name", a.id.String())

	req := &bigqueryanalyticshubpb.DeleteListingRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteListing(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Listing %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted Listing", "name", a.id.String())

	return true, nil
}
