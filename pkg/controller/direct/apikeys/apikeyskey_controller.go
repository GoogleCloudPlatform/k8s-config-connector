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

package apikeys

import (
	"context"
	"fmt"
	"reflect"

	api "cloud.google.com/go/apikeys/apiv2"
	pb "cloud.google.com/go/apikeys/apiv2/apikeyspb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/apikeys/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	. "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/mappings" //nolint:revive
)

func init() {
	registry.RegisterModel(krm.APIKeysKeyGVK, newAPIKeysModel)
}

func newAPIKeysModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: *config}, nil
}

type model struct {
	config config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &model{}

var keyMapping = NewMapping(&pb.Key{}, &krm.APIKeysKey{},
	Spec("displayName"),
	Spec("restrictions"),
	// Status("uid"),
	Ignore("createTime"),
	Ignore("updateTime"),
	Ignore("deleteTime"),
	Ignore("etag"),
	Ignore("name"),        // TODO: Should be ResourceID?
	Ignore("annotations"), // TODO: Should not ignore
).
	MapNested(&pb.Restrictions{}, &krm.KeyRestrictions{}, "apiTargets",
		"androidKeyRestrictions", "browserKeyRestrictions", "iosKeyRestrictions", "serverKeyRestrictions").
	MapNested(&pb.AndroidKeyRestrictions{}, &krm.KeyAndroidKeyRestrictions{}, "allowedApplications").
	MapNested(&pb.AndroidApplication{}, &krm.KeyAllowedApplications{}, "packageName", "sha1Fingerprint").
	MapNested(&pb.ApiTarget{}, &krm.KeyApiTargets{}, "methods", "service").
	MapNested(&pb.BrowserKeyRestrictions{}, &krm.KeyBrowserKeyRestrictions{}, "allowedReferrers").
	MapNested(&pb.IosKeyRestrictions{}, &krm.KeyIosKeyRestrictions{}, "allowedBundleIds").
	MapNested(&pb.ServerKeyRestrictions{}, &krm.KeyServerKeyRestrictions{}, "allowedIps").
	MustBuild()

type adapter struct {
	projectID string
	location  string
	keyID     string

	desired *krm.APIKeysKey
	actual  *krm.APIKeysKey

	gcp *api.Client
}

// adapter implements the Adapter interface.
var _ directbase.Adapter = &adapter{}

func (m *model) client(ctx context.Context) (*api.Client, error) {
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}

	gcpClient, err := api.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building apikeys client: %w", err)
	}
	return gcpClient, err
}

// AdapterForObject implements the Model interface.
func (m *model) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	gcp, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: Just fetch this object?
	obj := &krm.APIKeysKey{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	projectID := obj.Spec.ProjectRef.External
	if projectID == "" {
		return nil, fmt.Errorf("unable to determine project")
	}

	// TODO: Use name or request resourceID to be set on create?
	keyID := direct.ValueOf(obj.Spec.ResourceID)
	if keyID == "" {
		return nil, fmt.Errorf("unable to determine resourceID")
	}

	// This is a slightly unusual resource; the location is in the URL,
	// but the location is always "global".
	location := "global"

	return &adapter{
		projectID: projectID,
		location:  location,
		keyID:     keyID,
		desired:   obj,
		gcp:       gcp,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

// Find implements the Adapter interface.
func (a *adapter) Find(ctx context.Context) (bool, error) {
	if a.keyID == "" {
		return false, nil
	}

	req := &pb.GetKeyRequest{
		Name: a.fullyQualifiedName(),
	}
	key, err := a.gcp.GetKey(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			klog.Warningf("key was not found: %v", err)
			return false, nil
		}
		return false, err
	}

	u := &krm.APIKeysKey{}
	if err := keyMapping.Map(key, u); err != nil {
		return false, err
	}
	a.actual = u

	return true, nil
}

// Delete implements the Adapter interface.
func (a *adapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	// TODO: Delete via status selfLink?
	req := &pb.DeleteKeyRequest{
		Name: a.fullyQualifiedName(),
	}
	op, err := a.gcp.DeleteKey(ctx, req)
	if err != nil {
		return false, fmt.Errorf("deleting key: %w", err)
	}

	if _, err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for key deletion to complete: %w", err)
	}

	return true, nil
}

func (a *adapter) buildCreateRequest() (*pb.CreateKeyRequest, error) {
	// You can configure only the `display_name`, `restrictions`, and
	// `annotations` fields.
	desired := &pb.Key{}
	if err := keyMapping.Map(a.desired, desired); err != nil {
		return nil, err
	}

	return &pb.CreateKeyRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location),
		KeyId:  a.keyID,
		Key:    desired,
	}, nil
}

// Create implements the Adapter interface.
func (a *adapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	u := createOp.GetUnstructured()

	log := klog.FromContext(ctx)
	log.V(2).Info("creating object", "u", u)

	req, err := a.buildCreateRequest()
	if err != nil {
		return err
	}

	op, err := a.gcp.CreateKey(ctx, req)
	if err != nil {
		return fmt.Errorf("creating key: %w", err)
	}
	// TODO: Is the resourceID returned if it is dynamically created?  Maybe we should create the UUID?
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for key creation: %w", err)
	}
	log.V(2).Info("created key", "key", created)
	// TODO: Return created object
	return nil
}

// Update implements the Adapter interface.
func (a *adapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	// u := Op.GetUnstructured()

	// TODO: Skip updates if no changes
	// TODO: Where/how do we want to enforce immutability?
	updateMask := &fieldmaskpb.FieldMask{}

	// TODO: I think we can do this with a helper
	if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.Spec.DisplayName) {
		updateMask.Paths = append(updateMask.Paths, "display_name")
	}
	if !reflect.DeepEqual(a.desired.Spec.Restrictions, a.actual.Spec.Restrictions) {
		updateMask.Paths = append(updateMask.Paths, "restrictions")
	}

	// TODO: Annotations
	// if !reflect.DeepEqual(a.desired.Annotations, a.actual.Annotations) {
	// 	updateMask.Paths = append(updateMask.Paths, "annotations")
	// }

	if len(updateMask.Paths) == 0 {
		klog.Warningf("unexpected empty update mask, desired: %v, actual: %v", a.desired, a.actual)
		return nil
	}

	key := &pb.Key{}
	if err := keyMapping.Map(a.desired, key); err != nil {
		return err
	}

	req := &pb.UpdateKeyRequest{
		Key:        key,
		UpdateMask: updateMask,
	}

	req.Key.Name = a.fullyQualifiedName()

	_, err := a.gcp.UpdateKey(ctx, req)
	if err != nil {
		return err
	}
	// TODO: update status in u
	return nil
}

func (a *adapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("apikeyskey %q not found", a.fullyQualifiedName())
	}

	spec := direct.LazyPtr(a.actual.Spec)
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting apikeyskey spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetName(a.keyID)
	u.SetGroupVersionKind(krm.APIKeysKeyGVK)
	u.SetLabels(a.actual.Labels)
	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

func (a *adapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/keys/%s", a.projectID, a.location, a.keyID)
}
