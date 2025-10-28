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

// +tool:controller
// proto.service: google.cloud.essentialcontacts.v1.EssentialContactsService
// proto.message: google.cloud.essentialcontacts.v1.Contact
// crd.type: EssentialContactsContact
// crd.version: v1alpha1

package essentialcontacts

import (
	"context"
	"fmt"
	"reflect"

	gcp "cloud.google.com/go/essentialcontacts/apiv1"
	pb "cloud.google.com/go/essentialcontacts/apiv1/essentialcontactspb"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/essentialcontacts/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(krm.EssentialContactsContactGVK, NewContactModel)
}

func NewContactModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &contactModel{config: *config}, nil
}

var _ directbase.Model = &contactModel{}

type contactModel struct {
	config config.ControllerConfig
}

func (m *contactModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.EssentialContactsContact{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewContactIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	config := m.config
	// the service requires that a quota project be set
	if !config.UserProjectOverride || config.BillingProject == "" {
		config.UserProjectOverride = true
		config.BillingProject = id.Parent().ProjectID
	}
	// Get GCP client
	gcpClient, err := newGCPClient(ctx, &config)
	if err != nil {
		return nil, err
	}
	client, err := gcpClient.newEssentialContactsClient(ctx)
	if err != nil {
		return nil, err
	}
	return &contactAdapter{
		gcpClient: client,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *contactModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type contactAdapter struct {
	gcpClient *gcp.Client
	id        *krm.ContactIdentity
	desired   *krm.EssentialContactsContact
	actual    *pb.Contact
	reader    client.Reader
}

var _ directbase.Adapter = &contactAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *contactAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting essential contact", "name", a.id)

	// This resource has a server-generated ID. This means user should not know
	// the ID before the resource is created, and 'metadata.name' won't be used
	// as the default resource ID. So empty value for 'spec.resourceID' should
	// also be valid:
	// 1. When 'spec.resourceID' is not set or set to an empty value, the
	//    intention is to create the resource.
	// 2. When 'spec.resourceID' is set, the intention is to acquire an existing
	//    resource.
	//    2.1. When 'spec.resourceID' is set but the corresponding GCP resource
	//         is not found, then it is a real error.
	if a.id.ID() == "" {
		log.V(2).Info("no resource ID in get indicates the create intention", "name", a.id)
		return false, nil
	}

	req := &pb.GetContactRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetContact(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting essential contact %q: %w", a.id, err)
	}

	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *contactAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating essential contact", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := EssentialContactsContactSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	req := &pb.CreateContactRequest{
		Parent:  a.id.Parent().String(),
		Contact: resource,
	}
	created, err := a.gcpClient.CreateContact(ctx, req)
	if err != nil {
		return fmt.Errorf("creating essential contact %s: %w", a.id, err)
	}
	log.V(2).Info("successfully created essential contact", "name", a.id)

	status := &krm.EssentialContactsContactStatus{}
	status.ObservedState = EssentialContactsContactObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(created.Name)
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *contactAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating essential contact", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := EssentialContactsContactSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String() // Set name for update request

	paths := []string{}
	if !reflect.DeepEqual(resource.NotificationCategorySubscriptions, a.actual.NotificationCategorySubscriptions) {
		paths = append(paths, "notification_category_subscriptions")
	}
	if !reflect.DeepEqual(resource.LanguageTag, a.actual.LanguageTag) {
		paths = append(paths, "language_tag")
	}

	var updated *pb.Contact
	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		updated = a.actual
	} else {
		req := &pb.UpdateContactRequest{
			Contact:    resource,
			UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		}
		var err error
		updated, err = a.gcpClient.UpdateContact(ctx, req)
		if err != nil {
			return fmt.Errorf("updating essential contact %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated essential contact", "name", a.id)
	}

	status := &krm.EssentialContactsContactStatus{}
	status.ObservedState = EssentialContactsContactObservedState_FromProto(mapCtx, updated)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.LazyPtr(updated.Name)

	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *contactAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.EssentialContactsContact{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(EssentialContactsContactSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set parent ref based on the actual name format
	parentIdentity, resourceID, err := krm.ParseContactExternal(a.id.String())
	if err != nil {
		return nil, fmt.Errorf("parsing parent from actual name %q: %w", a.id, err)
	}
	if parentIdentity.ProjectID != "" {
		obj.Spec.ProjectRef = &refv1beta1.ProjectRef{
			External: parentIdentity.String(),
		}
	} else if parentIdentity.FolderID != "" {
		obj.Spec.FolderRef = &refv1beta1.FolderRef{
			External: parentIdentity.String(),
		}
	} else if parentIdentity.OrganizationID != "" {
		obj.Spec.OrganizationRef = &refv1beta1.OrganizationRef{
			External: parentIdentity.String(),
		}
	} else {
		return nil, fmt.Errorf("unknown parent type in name %q", a.actual.Name)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(resourceID)
	u.SetGroupVersionKind(krm.EssentialContactsContactGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *contactAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting essential contact", "name", a.id)

	req := &pb.DeleteContactRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteContact(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent essential contact, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting essential contact %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted essential contact", "name", a.id)

	return true, nil
}
