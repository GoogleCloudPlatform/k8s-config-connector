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

package resourcemanager

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/resourcemanager/apiv3"
	tagspb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"

	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.TagsLocationTagBindingGVK, NewTagsLocationTagBindingModel)
}

func NewTagsLocationTagBindingModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelTagsLocationTagBinding{config: *config}, nil
}

var _ directbase.Model = &modelTagsLocationTagBinding{}

type modelTagsLocationTagBinding struct {
	config config.ControllerConfig
}

func (m *modelTagsLocationTagBinding) client(ctx context.Context, location string) (*gcp.TagBindingsClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("https://%s-cloudresourcemanager.googleapis.com", location)
	opts = append(opts, option.WithEndpoint(endpoint))
	gcpClient, err := gcp.NewTagBindingsRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TagBinding client with endpoint: %s, %w", endpoint, err)
	}
	return gcpClient, err
}

func (m *modelTagsLocationTagBinding) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.TagsLocationTagBinding{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewTagsLocationTagBindingIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	// Get tags GCP client
	gcpClient, err := m.client(ctx, id.Location())
	if err != nil {
		return nil, err
	}
	return &TagsLocationTagBindingAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
	}, nil
}

func (m *modelTagsLocationTagBinding) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type TagsLocationTagBindingAdapter struct {
	id        *krm.TagsLocationTagBindingIdentity
	gcpClient *gcp.TagBindingsClient
	desired   *krm.TagsLocationTagBinding
	actual    *tagspb.TagBinding
}

var _ directbase.Adapter = &TagsLocationTagBindingAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *TagsLocationTagBindingAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting TagBinding", "name", a.id)

	req := &tagspb.ListTagBindingsRequest{Parent: a.id.Parent().String()}
	tagsIterator := a.gcpClient.ListTagBindings(ctx, req)
	if tagsIterator == nil {
		fmt.Printf("not found any TagsLocationtagbindings for the resource, %q\n", a.id.Parent().String())
		return false, nil
	}

	// There is no more items when error is iterator.Done.
	var err error
	var tagbindingpb *tagspb.TagBinding
	for tagbindingpb, err = tagsIterator.Next(); err == nil; {
		if tagbindingpb.Name == a.id.String() {
			a.actual = tagbindingpb
			return true, nil
		}
		tagbindingpb, err = tagsIterator.Next()
	}

	return false, fmt.Errorf("getting TagsLocationTagBinding %q: %w", a.id, err)

}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TagsLocationTagBindingAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating New TagsLocationTagBinding", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	resource := TagsLocationTagBindingSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	// TODO(contributor): Complete the gcp "CREATE" or "INSERT" request.
	req := &tagspb.CreateTagBindingRequest{
		TagBinding: resource,
	}
	op, err := a.gcpClient.CreateTagBinding(ctx, req)
	if err != nil {
		return fmt.Errorf("creating TagsLocationTagBinding %s: %w", a.id, err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("TagsLocationTagBinding %s waiting creation: %w", a.id, err)
	}
	log.V(2).Info("successfully created TagsLocationTagBinding", "name", a.id)

	status := &krm.TagsLocationTagBindingStatus{}
	status.ObservedState = TagsLocationTagBindingObservedState_FromProto(mapCtx, created)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	status.ExternalRef = direct.LazyPtr(created.GetName())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *TagsLocationTagBindingAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating TagsLocationTagBinding", "name", a.id)
	// mapCtx := &direct.MapContext{}

	/* 	desiredPb := TagsLocationTagBindingSpec_ToProto(mapCtx, &a.desired.DeepCopy().Spec)
	   	if mapCtx.Err() != nil {
	   		return mapCtx.Err()
	   	} */

	// paths := make(sets.Set[string])
	/* 	// Option 1: This option is good for proto that has `field_mask` for output-only, immutable, required/optional.
	   	// TODO(contributor): If choosing this option, remove the "Option 2" code.
	   	{
	   		var err error
	   		paths, err = common.CompareProtoMessage(desiredPb, a.actual, common.BasicDiff)
	   		if err != nil {
	   			return err
	   		}
	   	} */

	// Option 2: manually add all mutable fields.
	// TODO(contributor): If choosing this option, remove the "Option 1" code.
	/* 	{
		if !reflect.DeepEqual(a.desired.Spec.DisplayName, a.actual.DisplayName) {
			paths = paths.Insert("display_name")
		}
	} */

	status := &krm.TagsLocationTagBindingStatus{}
	/* 	status.ObservedState = TagsLocationTagBindingObservedState_FromProto(mapCtx, updated)
	   	if mapCtx.Err() != nil {
	   		return mapCtx.Err()
	   	} */
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export maps the GCP object to a Config Connector resource `spec`.
func (a *TagsLocationTagBindingAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.TagsLocationTagBinding{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(TagsLocationTagBindingSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	// obj.Spec.ProjectRef = &refs.ProjectRef{External: a.id.Parent().ProjectID}
	// obj.Spec.Location = a.id.Parent().Location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.TagsLocationTagBindingGVK)

	u.Object = uObj
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *TagsLocationTagBindingAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting TagsLocationTagBinding", "name", a.id)

	req := &tagspb.DeleteTagBindingRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteTagBinding(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent TagsLocationTagBinding, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting TagsLocationTagBinding %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted TagsLocationTagBinding", "name", a.id)

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete TagsLocationTagBinding %s: %w", a.id, err)
	}
	return true, nil
}
