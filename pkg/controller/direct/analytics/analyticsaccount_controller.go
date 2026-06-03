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
// proto.service: google.analytics.admin.v1alpha.AnalyticsAdminService
// proto.message: google.analytics.admin.v1alpha.Account
// crd.type: AnalyticsAccount
// crd.version: v1alpha1

package analytics

import (
	"context"
	"errors"
	"fmt"

	gcp "cloud.google.com/go/analytics/admin/apiv1alpha"
	pb "cloud.google.com/go/analytics/admin/apiv1alpha/adminpb"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.AnalyticsAccountGVK, NewAccountModel)
}

func NewAccountModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &accountModel{config: *config}, nil
}

var _ directbase.Model = &accountModel{}

type accountModel struct {
	config config.ControllerConfig
}

func (m *accountModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.AnalyticsAccount{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := identity.(*krm.AccountIdentity)

	config := m.config
	// Get GCP client
	gcpClient, err := newGCPClient(ctx, &config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newAnalyticsAdminClient(ctx)
	if err != nil {
		return nil, err
	}
	return &accountAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *accountModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type accountAdapter struct {
	gcpClient *gcp.AnalyticsAdminClient
	id        *krm.AccountIdentity
	desired   *krm.AnalyticsAccount
	actual    *pb.Account
	reader    client.Reader
}

var _ directbase.Adapter = &accountAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *accountAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting analytics account", "name", a.id)

	if a.id.ID() == "" {
		id, acc, err := a.discoverAccountID(ctx)
		if err != nil {
			return false, err
		}
		if id != "" {
			if err := a.id.FromExternal(id); err != nil {
				return false, err
			}
			a.actual = acc
			return true, nil
		}

		if a.desired.Status.ObservedState != nil && a.desired.Status.ObservedState.AccountTicketID != nil {
			log.V(2).Info("account ticket already provisioned, awaiting user acceptance and resourceID setting", "ticket", *a.desired.Status.ObservedState.AccountTicketID)
			return true, nil
		}

		log.V(2).Info("no resource ID in get indicates the create intention", "name", a.id)
		return false, nil
	}

	req := &pb.GetAccountRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetAccount(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting analytics account %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

func (a *accountAdapter) discoverAccountID(ctx context.Context) (string, *pb.Account, error) {
	log := klog.FromContext(ctx)
	req := &pb.ListAccountsRequest{}
	it := a.gcpClient.ListAccounts(ctx, req)
	for {
		acc, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return "", nil, fmt.Errorf("listing analytics accounts: %w", err)
		}
		if acc.DisplayName == direct.ValueOf(a.desired.Spec.DisplayName) {
			log.V(2).Info("automatically discovered analytics account ID via display name", "name", acc.Name)
			return acc.Name, acc, nil
		}
	}
	return "", nil, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *accountAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating analytics account", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AnalyticsAccountSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.ProvisionAccountTicketRequest{
		Account:     resource,
		RedirectUri: direct.ValueOf(desired.Spec.RedirectURI),
	}
	res, err := a.gcpClient.ProvisionAccountTicket(ctx, req)
	if err != nil {
		return fmt.Errorf("provisioning analytics account ticket %s: %w", a.id, err)
	}
	log.V(2).Info("successfully provisioned analytics account ticket", "name", a.id)

	status := &krm.AnalyticsAccountStatus{}
	status.ObservedState = AnalyticsAccountObservedState_FromAccountTicketID(mapCtx, res.AccountTicketId)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// Try to discover the newly created account immediately (will succeed in mockgcp synchronously)
	id, acc, err := a.discoverAccountID(ctx)
	if err != nil {
		log.Error(err, "error discovering account ID after provisioning ticket")
	}
	if id != "" {
		status.ExternalRef = direct.LazyPtr(id)
		status.ObservedState = AnalyticsAccountObservedState_FromProto(mapCtx, acc)
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
	}

	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *accountAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating analytics account", "name", a.id)
	mapCtx := &direct.MapContext{}

	if a.id.ID() == "" {
		// Awaiting TOS acceptance.
		status := &krm.AnalyticsAccountStatus{}
		if a.desired.Status.ObservedState != nil {
			status.ObservedState = a.desired.Status.ObservedState.DeepCopy()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	desired := a.desired.DeepCopy()
	resource := AnalyticsAccountSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String()

	paths, report, err := common.CompareProtoMessageStructuredDiff(resource, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}

	var updated *pb.Account
	if paths.Len() == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		structuredreporting.ReportDiff(ctx, report)
		req := &pb.UpdateAccountRequest{
			Account:    resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: sets.List(paths)},
		}
		var err error
		updated, err = a.gcpClient.UpdateAccount(ctx, req)
		if err != nil {
			return fmt.Errorf("updating analytics account %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated analytics account", "name", a.id)
	}

	status := &krm.AnalyticsAccountStatus{}
	status.ObservedState = AnalyticsAccountObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.Name)

	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *accountAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AnalyticsAccount{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AnalyticsAccountSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.AnalyticsAccountGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *accountAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting analytics account", "name", a.id)

	req := &pb.DeleteAccountRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAccount(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent analytics account, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting analytics account %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted analytics account", "name", a.id)

	return true, nil
}
