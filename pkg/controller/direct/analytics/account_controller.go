// Copyright 2025 Google LLC
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

package analytics

import (
	"context"
	"fmt"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"

	gcp "cloud.google.com/go/analytics/admin/apiv1alpha"
	analyticspb "cloud.google.com/go/analytics/admin/apiv1alpha/adminpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.AnalyticsAccountGVK, NewAccountModel)
}

func NewAccountModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelAccount{config: *config}, nil
}

var _ directbase.Model = &modelAccount{}

type modelAccount struct {
	config config.ControllerConfig
}

func (m *modelAccount) client(ctx context.Context) (*gcp.AnalyticsAdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewAnalyticsAdminClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building Account client: %w", err)
	}
	return gcpClient, err
}

func (m *modelAccount) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.AnalyticsAccount{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAccountIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get analytics GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}
	return &AccountAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelAccount) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type AccountAdapter struct {
	id        *krm.AccountIdentity
	gcpClient *gcp.AnalyticsAdminClient
	desired   *krm.AnalyticsAccount
	actual    *analyticspb.Account
}

var _ directbase.Adapter = &AccountAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *AccountAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting Account", "name", a.id)

	// Step I: Try to find the account.
	// Option 1: Get Account via account ID.
	if a.id.ID() != "" {
		req := &analyticspb.GetAccountRequest{Name: a.id.String()}
		accountpb, err := a.gcpClient.GetAccount(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				log.V(2).Info("not found via GetAccount()", "name", a.id)
				return false, nil
			}
			return false, fmt.Errorf("getting Account %q: %w", a.id, err)
		}
		a.actual = accountpb
		return true, nil
	}

	// Option 2: Get Account via the combination of displayName and regionCode.
	//
	// When a.id.ID() is nil, it's possible that the account was just provisioned in the last reconciliation, and the
	// controller hasn't populated `status.externalRef` yet. To try to get the resource ID, we'll do a LIST first and
	// look for the account that matches the desired display name and region code.
	accountIterator := a.gcpClient.ListAccounts(ctx, &analyticspb.ListAccountsRequest{})
	if accountIterator != nil {
		// The iterator doesn't return an error when there is no more items. Could it be a bug?
		prevName := ""
		for accountpb, err := accountIterator.Next(); err == nil; {
			currName := accountpb.Name
			if currName != "" && prevName == currName {
				// Iteration should be done now.
				break
			}
			prevName = currName
			if accountpb.DisplayName == *a.desired.Spec.DisplayName && accountpb.RegionCode == *a.desired.Spec.RegionCode {
				a.actual = accountpb
				resourceID, err := krm.ParseAccountExternal(accountpb.Name)
				if err != nil {
					return false, fmt.Errorf("parsing Account %q: %w", accountpb.Name, err)
				}
				a.id.SetID(resourceID)
				return true, nil
			}
			continue
		}
	}

	// Step II: Try to identify what's missing.
	idInfo := fmt.Sprintf("displayName %q and regionCode %q", *a.desired.Spec.DisplayName, *a.desired.Spec.RegionCode)
	log.V(2).Info("not found via ListAccount()", "identification info", idInfo)

	if a.desired.Status.ObservedState != nil {
		accountTicketID := a.desired.Status.ObservedState.AccountTicketID
		if accountTicketID != nil { // Reconciliation after creation ticket is provisioned.
			namespacedName := types.NamespacedName{
				Namespace: a.desired.Namespace,
				Name:      a.desired.Name,
			}
			if *accountTicketID != "" {
				return false, k8s.NewManualStepNotCompletedError(a.desired.GroupVersionKind(), namespacedName,
					fmt.Sprintf("creation of a Google Analytics account is not completed. Go to ToS page "+
						"https://analytics.google.com/analytics/web/?provisioningSignup=false#/termsofservice/%s and "+
						"accept the Terms of Service", *accountTicketID))
			} else {
				return false, k8s.NewManualStepNotCompletedError(a.desired.GroupVersionKind(), namespacedName,
					fmt.Sprintf("provisioning of a Google Analytics account ticket returned an empty ticket ID"))
			}
		}
	}
	return false, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AccountAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating Account", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := AnalyticsAccountSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &analyticspb.ProvisionAccountTicketRequest{
		Account:     resource,
		RedirectUri: *a.desired.Spec.RedirectURI,
	}
	resp, err := a.gcpClient.ProvisionAccountTicket(ctx, req)
	if err != nil {
		return fmt.Errorf("provisioning Account ticket %s: %w", a.id, err)
	}
	log.V(2).Info("successfully provisioned Account", "name", a.id)

	status := &krm.AnalyticsAccountStatus{}
	status.ObservedState = AnalyticsAccountObservedState_FromAccountTicketID(mapCtx, resp.AccountTicketId)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	// status.ExternalRef cannot be set at creation time.
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *AccountAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating Account", "name", a.id)
	mapCtx := &direct.MapContext{}

	desiredPb := AnalyticsAccountSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desiredPb.Name = a.id.String()

	paths := make(sets.Set[string])
	// Option 1: This option is good for proto that has `field_mask` for output-only, immutable, required/optional.
	// TODO(contributor): If choosing this option, remove the "Option 2" code.
	{
		var err error
		paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
		if err != nil {
			return err
		}
	}

	//// Option 2: manually add all mutable fields.
	//// TODO(contributor): If choosing this option, remove the "Option 1" code.
	//{
	//	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
	//		paths = paths.Insert("display_name")
	//	}
	//}

	updated := a.actual
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id, "paths", paths)
		updateMask := &fieldmaskpb.FieldMask{
			Paths: sets.List(paths),
		}

		// TODO(contributor): Complete the gcp "UPDATE" or "PATCH" request.
		req := &analyticspb.UpdateAccountRequest{
			UpdateMask: updateMask,
			Account:    desiredPb,
		}
		var err error
		updated, err = a.gcpClient.UpdateAccount(ctx, req)
		if err != nil {
			return fmt.Errorf("updating Account %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated Account", "name", a.id)
	}

	status := &krm.AnalyticsAccountStatus{}
	status.ExternalRef = direct.LazyPtr(a.id.String())
	status.ObservedState = AnalyticsAccountObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *AccountAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
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
	if a.actual != nil {
		resourceID, err := krm.ParseAccountExternal(a.actual.Name)
		if err != nil {
			return nil, fmt.Errorf("parsing Account %q: %w", a.actual.Name, err)
		}
		obj.Spec.ResourceID = direct.LazyPtr(resourceID)
	}
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	if a.actual != nil {
		u.SetName("non-existent-account")
	} else {
		// The resource ID for an Account is pure number so it needs prefix to construct a KRM object name.
		u.SetName(fmt.Sprintf("account-%s", *obj.Spec.ResourceID))
	}
	u.SetGroupVersionKind(krm.AnalyticsAccountGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *AccountAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting Account", "name", a.id)

	// The service-generated ID may not exist yet.
	if a.id.ID() == "" {
		log.V(2).Info("skipping delete for Account without an ID", "name", a.id)
		return true, nil
	}

	req := &analyticspb.DeleteAccountRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteAccount(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent Account, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting Account %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted Account", "name", a.id)
	return true, nil
}
