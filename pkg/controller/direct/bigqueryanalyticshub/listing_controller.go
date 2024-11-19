// Copyright 2024 Google LLC
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

package bigqueryanalyticshub

import (
	"context"
	"fmt"
	"reflect"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	gcp "cloud.google.com/go/bigquery/analyticshub/apiv1"
	bigqueryanalyticshubpb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func (m *modelListing) client(ctx context.Context) (*gcp.Client, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Listing client: %w", err)
	}
	return gcpClient, err
}

func (m *modelListing) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.BigQueryAnalyticsHubListing{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewBigQueryAnalyticsHubListingRef(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	if err := resolveOptionalReferences(ctx, reader, obj); err != nil {
		return nil, err
	}
	// Get bigqueryanalyticshub GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &ListingAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func resolveOptionalReferences(ctx context.Context, reader client.Reader, obj *krm.BigQueryAnalyticsHubListing) error {
	if ref := obj.Spec.DataExchangeRef; ref != nil {
		_, err := refs.ResolveDataExchangeRef(ctx, reader, obj, ref)
		if err != nil {
			return err
		}
	}

	if obj.Spec.Source != nil && obj.Spec.Source.BigQueryDatasetSource != nil {
		if ref := obj.Spec.Source.BigQueryDatasetSource.Dataset; ref != nil {
			if _, err := refs.ResolveBigQueryDataset(ctx, reader, obj, ref); err != nil {
				return err
			}

			for _, selectedResource := range obj.Spec.Source.BigQueryDatasetSource.SelectedResources {
				if ref := selectedResource.TableRef; ref != nil {
					if _, err := refs.ResolveBigQueryTable(ctx, reader, obj, ref); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (m *modelListing) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type ListingAdapter struct {
	id        *krm.BigQueryAnalyticsHubListingRef
	gcpClient *gcp.Client
	desired   *krm.BigQueryAnalyticsHubListing
	actual    *bigqueryanalyticshubpb.Listing
}

var _ directbase.Adapter = &ListingAdapter{}

func (a *ListingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(listingCtrlName)
	log.V(2).Info("getting Listing", "name", a.id.External)

	req := &bigqueryanalyticshubpb.GetListingRequest{Name: a.id.External}
	listingpb, err := a.gcpClient.GetListing(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting Listing %q: %w", a.id.External, err)
	}

	a.actual = listingpb
	return true, nil
}

func (a *ListingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx).WithName(listingCtrlName)
	log.V(2).Info("creating Listing", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigQueryAnalyticsHubListingSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.id.Parent()
	if err != nil {
		return err
	}

	req := &bigqueryanalyticshubpb.CreateListingRequest{
		Parent:    parent.String(),
		Listing:   resource,
		ListingId: a.desired.GetName(),
	}
	created, err := a.gcpClient.CreateListing(ctx, req)
	if err != nil {
		return fmt.Errorf("creating Listing %s: %w", a.id.External, err)
	}

	log.V(2).Info("successfully created Listing", "name", a.id.External)

	status := &krm.BigQueryAnalyticsHubListingStatus{}
	status.ObservedState = BigQueryAnalyticsHubListingObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = &a.id.External
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *ListingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx).WithName(listingCtrlName)
	log.V(2).Info("updating Listing", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := BigQueryAnalyticsHubListingSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	updateMask := &fieldmaskpb.FieldMask{}
	if a.desired.Spec.DisplayName != nil && !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if a.desired.Spec.Description != nil && !reflect.DeepEqual(a.desired.Spec.Description, a.actual.Description) {
		updateMask.Paths = append(updateMask.Paths, "description")
	}
	if a.desired.Spec.PrimaryContact != nil && !reflect.DeepEqual(a.desired.Spec.PrimaryContact, a.actual.PrimaryContact) {
		updateMask.Paths = append(updateMask.Paths, "primary_contact")
	}
	if a.desired.Spec.Documentation != nil && !reflect.DeepEqual(a.desired.Spec.Documentation, a.actual.Documentation) {
		updateMask.Paths = append(updateMask.Paths, "documentation")
	}
	if a.desired.Spec.DiscoveryType != nil && !reflect.DeepEqual(a.desired.Spec.DiscoveryType, a.actual.DiscoveryType.String()) {
		updateMask.Paths = append(updateMask.Paths, "discovery_type")
	}
	if a.desired.Spec.RequestAccess != nil && reflect.DeepEqual(a.desired.Spec.RequestAccess, a.actual.RequestAccess) {
		updateMask.Paths = append(updateMask.Paths, "request_access")
	}

	// NOT YET
	// if a.desired.Spec.Icon != nil && !reflect.DeepEqual(a.desired.Spec.Icon, a.actual.Icon) {
	// 	updateMask.Paths = append(updateMask.Paths, "icon")
	// }
	if a.desired.Spec.DataProvider != nil {
		mapCtx := &direct.MapContext{}
		toProto := DataProvider_ToProto(mapCtx, a.desired.Spec.DataProvider)
		if mapCtx.Err() != nil {
			return fmt.Errorf("converting data provider: %w", mapCtx.Err())
		}

		if !reflect.DeepEqual(toProto, a.actual.DataProvider) {
			updateMask.Paths = append(updateMask.Paths, "data_provider")
		}
	}

	if a.desired.Spec.Publisher != nil {
		mapCtx := &direct.MapContext{}
		toProto := Publisher_ToProto(mapCtx, a.desired.Spec.Publisher)
		if mapCtx.Err() != nil {
			return fmt.Errorf("converting publisher: %w", mapCtx.Err())
		}

		if !reflect.DeepEqual(toProto, a.actual.Publisher) {
			updateMask.Paths = append(updateMask.Paths, "publisher")
		}
	}

	if a.desired.Spec.Categories != nil {
		mapCtx := &direct.MapContext{}
		toProto := Categories_ToProto(mapCtx, a.desired.Spec.Categories)
		if mapCtx.Err() != nil {
			return fmt.Errorf("converting categories: %w", mapCtx.Err())
		}
		if !reflect.DeepEqual(toProto, a.actual.Categories) {
			updateMask.Paths = append(updateMask.Paths, "categories")
		}
	}

	if len(updateMask.Paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id.External)
		return nil
	}

	req := &bigqueryanalyticshubpb.UpdateListingRequest{
		UpdateMask: updateMask,
		Listing:    resource,
	}
	updated, err := a.gcpClient.UpdateListing(ctx, req)
	if err != nil {
		return fmt.Errorf("updating Listing %s: %w", a.id.External, err)
	}

	log.V(2).Info("successfully updated Listing", "name", a.id.External)

	status := &krm.BigQueryAnalyticsHubListingStatus{}
	status.ObservedState = BigQueryAnalyticsHubListingObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	return updateOp.UpdateStatus(ctx, status, nil)
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

	parent, err := a.id.Parent()
	if err != nil {
		return nil, err
	}
	obj.Spec.ProjectRef = &refs.ProjectRef{External: parent.String()}
	obj.Spec.Location = parent.Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	return u, nil
}

// Delete implements the Adapter interface.
func (a *ListingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(listingCtrlName)
	log.V(2).Info("deleting Listing", "name", a.id.External)

	req := &bigqueryanalyticshubpb.DeleteListingRequest{Name: a.id.External}
	err := a.gcpClient.DeleteListing(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting Listing %s: %w", a.id.External, err)
	}
	log.V(2).Info("successfully deleted Listing", "name", a.id.External)

	return true, nil
}
