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

package commerceproducer

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/commerceproducer/apiv1beta"
	pb "cloud.google.com/go/commerceproducer/apiv1beta/commerceproducerpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/commerceproducer/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CommerceProducerPrivateOfferGVK, NewCommerceProducerPrivateOfferModel)
}

func NewCommerceProducerPrivateOfferModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelCommerceProducerPrivateOffer{config: *config}, nil
}

var _ directbase.Model = &modelCommerceProducerPrivateOffer{}

type modelCommerceProducerPrivateOffer struct {
	config config.ControllerConfig
}

func (m *modelCommerceProducerPrivateOffer) client(ctx context.Context) (*gcp.CommerceTransactionClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewCommerceTransactionRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building CommerceProducerPrivateOffer client: %w", err)
	}
	return gcpClient, err
}

func (m *modelCommerceProducerPrivateOffer) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CommerceProducerPrivateOffer{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Normalize resource references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.CommerceProducerPrivateOfferIdentity)

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	desiredProto := CommerceProducerPrivateOfferSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	desiredProto.Name = id.String()

	return &CommerceProducerPrivateOfferAdapter{
		id:           id,
		gcpClient:    gcpClient,
		desiredProto: desiredProto,
	}, nil
}

func (m *modelCommerceProducerPrivateOffer) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type CommerceProducerPrivateOfferAdapter struct {
	id           *krm.CommerceProducerPrivateOfferIdentity
	gcpClient    *gcp.CommerceTransactionClient
	desiredProto *pb.PrivateOffer
	actual       *pb.PrivateOffer
}

var _ directbase.Adapter = &CommerceProducerPrivateOfferAdapter{}

func (a *CommerceProducerPrivateOfferAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting CommerceProducerPrivateOffer", "name", a.id)

	req := &pb.GetPrivateOfferRequest{
		Name: a.id.String(),
		View: pb.PrivateOfferView_PRIVATE_OFFER_VIEW_FULL,
	}
	privateOfferpb, err := a.gcpClient.GetPrivateOffer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting CommerceProducerPrivateOffer %q: %w", a.id, err)
	}

	a.actual = privateOfferpb
	return true, nil
}

func (a *CommerceProducerPrivateOfferAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating CommerceProducerPrivateOffer", "name", a.id)

	req := &pb.CreatePrivateOfferRequest{
		Parent:       a.id.ParentString(),
		PrivateOffer: a.desiredProto,
	}
	created, err := a.gcpClient.CreatePrivateOffer(ctx, req)
	if err != nil {
		return fmt.Errorf("creating CommerceProducerPrivateOffer %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created CommerceProducerPrivateOffer", "name", a.id)

	// Fetch full view after create to ensure all status fields are populated
	getReq := &pb.GetPrivateOfferRequest{
		Name: created.GetName(),
		View: pb.PrivateOfferView_PRIVATE_OFFER_VIEW_FULL,
	}
	fullCreated, err := a.gcpClient.GetPrivateOffer(ctx, getReq)
	if err != nil {
		// Fallback to returned object if get fails
		fullCreated = created
	}

	return a.updateStatus(ctx, createOp, fullCreated)
}

func (a *CommerceProducerPrivateOfferAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating CommerceProducerPrivateOffer", "name", a.id)

	diffs, updateMask, err := compareCommerceProducerPrivateOffer(ctx, a.actual, a.desiredProto)
	if err != nil {
		return fmt.Errorf("comparing CommerceProducerPrivateOffer %s: %w", a.id, err)
	}

	latest := a.actual
	if !diffs.HasDiff() {
		log.V(2).Info("no changes detected for CommerceProducerPrivateOffer", "name", a.id)
	} else {
		// Report exact diffs
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		req := &pb.UpdatePrivateOfferRequest{
			UpdateMask:   updateMask,
			PrivateOffer: a.desiredProto,
		}
		updated, err := a.gcpClient.UpdatePrivateOffer(ctx, req)
		if err != nil {
			return fmt.Errorf("updating CommerceProducerPrivateOffer %s: %w", a.id, err)
		}
		latest = updated
		log.V(2).Info("successfully updated CommerceProducerPrivateOffer", "name", a.id)

		// Fetch full view after update
		getReq := &pb.GetPrivateOfferRequest{
			Name: latest.GetName(),
			View: pb.PrivateOfferView_PRIVATE_OFFER_VIEW_FULL,
		}
		fullLatest, err := a.gcpClient.GetPrivateOffer(ctx, getReq)
		if err == nil {
			latest = fullLatest
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func (a *CommerceProducerPrivateOfferAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	desired := &krm.CommerceProducerPrivateOffer{}
	mapCtx := &direct.MapContext{}
	desired.Spec = direct.ValueOf(CommerceProducerPrivateOfferSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Spec.ProjectRef = &refsv1beta1.ProjectRef{External: a.id.Project}
	desired.Spec.Location = direct.LazyPtr(a.id.Location)

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(desired)
	if err != nil {
		return nil, err
	}

	u.Object = uObj
	u.SetName(a.id.PrivateOffer)
	u.SetGroupVersionKind(krm.CommerceProducerPrivateOfferGVK)

	return u, nil
}

func (a *CommerceProducerPrivateOfferAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting CommerceProducerPrivateOffer", "name", a.id)

	req := &pb.DeletePrivateOfferRequest{
		Name:  a.id.String(),
		Force: true,
	}
	err := a.gcpClient.DeletePrivateOffer(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent CommerceProducerPrivateOffer, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting CommerceProducerPrivateOffer %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted CommerceProducerPrivateOffer", "name", a.id)

	return true, nil
}

func compareCommerceProducerPrivateOffer(ctx context.Context, actual, desired *pb.PrivateOffer) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CommerceProducerPrivateOfferSpec_FromProto, CommerceProducerPrivateOfferSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name

	clonedDesired := proto.CloneOf(desired)

	populateDefaults := func(obj *pb.PrivateOffer) {
		// Add any server-side or GCP defaults if known, or leave empty
	}
	populateDefaults(maskedActual)
	populateDefaults(clonedDesired)

	diffs, updateMask, err := common.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *CommerceProducerPrivateOfferAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.PrivateOffer) error {
	mapCtx := &direct.MapContext{}
	status := &krm.CommerceProducerPrivateOfferStatus{}
	status.ObservedState = CommerceProducerPrivateOfferObservedState_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	externalRef := latest.GetName()
	if externalRef == "" {
		externalRef = a.id.String()
	}
	status.ExternalRef = direct.LazyPtr(externalRef)
	return op.UpdateStatus(ctx, status, nil)
}
