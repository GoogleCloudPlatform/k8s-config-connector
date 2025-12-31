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
// proto.service: google.cloud.asset.v1.AssetService
// proto.message: google.cloud.asset.v1.Feed
// crd.type: AssetFeed
// crd.version: v1beta1

package asset

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/asset/apiv1"
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.AssetFeedGVK, NewFeedModel)
}

func NewFeedModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &feedModel{config: *config}, nil
}

var _ directbase.Model = &feedModel{}

type feedModel struct {
	config config.ControllerConfig
}

func (m *feedModel) client(ctx context.Context, projectID string) (*gcp.Client, error) {
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

	gcpClient, err := gcp.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building asset feed client: %w", err)
	}

	return gcpClient, err
}

func (m *feedModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.AssetFeed{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewFeedIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx, id.Parent().ProjectID)
	if err != nil {
		return nil, err
	}

	return &feedAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *feedModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type feedAdapter struct {
	gcpClient *gcp.Client
	id        *krm.FeedIdentity
	desired   *krm.AssetFeed
	actual    *pb.Feed
	reader    client.Reader
}

var _ directbase.Adapter = &feedAdapter{}

func (a *feedAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting asset feed", "name", a.id)

	name := a.id.String()
	if a.actual != nil {
		name = a.actual.Name
	}
	req := &pb.GetFeedRequest{Name: name}
	actual, err := a.gcpClient.GetFeed(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting asset feed %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *feedAdapter) normalizeReferences(ctx context.Context) error {
	obj := a.desired
	if obj.Spec.FeedOutputConfig.PubsubDestination != nil && obj.Spec.FeedOutputConfig.PubsubDestination.TopicRef != nil {
		ref := obj.Spec.FeedOutputConfig.PubsubDestination.TopicRef
		_, err := ref.NormalizedExternal(ctx, a.reader, obj.GetNamespace())
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *feedAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating asset feed", "name", a.id)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := AssetFeedSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	resource.Name = a.id.String() // Name is required for CreateFeed

	req := &pb.CreateFeedRequest{
		Parent: a.id.Parent().String(),
		FeedId: a.id.ID(),
		Feed:   resource,
	}

	// Attempt to create the feed
	actual, err := a.gcpClient.CreateFeed(ctx, req)

	// Regardless of CreateFeed result (success or AlreadyExists),
	// we need to ensure a.actual is populated by calling Find.
	found := false
	var findErr error

	if err == nil {
		log.V(2).Info("successfully created asset feed in gcp, attempting to fetch it", "name", a.id)
		a.actual = actual
		found = true
	} else if errors.IsAlreadyExists(err) {
		// Resource already exists. Log and attempt to Find to populate a.actual.
		log.V(2).Info("asset feed already exists during creation attempt, attempting to fetch it", "name", a.id.String(), "warning", err)
		found, findErr = a.Find(ctx)
		if findErr != nil {
			return fmt.Errorf("fetching existing asset feed %q after CreateFeed returned AlreadyExists failed: %w", a.id.String(), findErr)
		}
		if !found {
			// This is unexpected if CreateFeed returned AlreadyExists
			return fmt.Errorf("asset feed %q not found even though CreateFeed returned AlreadyExists", a.id.String())
		}
	} else {
		// Other error during creation
		return fmt.Errorf("creating asset feed %s: %w", a.id.String(), err)
	}

	// Update status
	status := &krm.AssetFeedStatus{}
	status.ExternalRef = direct.LazyPtr(a.actual.Name)
	// Pass the fetched 'a.actual' to UpdateStatus for potential status mapping
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *feedAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating asset feed", "name", a.id, "actual", a.actual.Name)
	mapCtx := &direct.MapContext{}

	if err := a.normalizeReferences(ctx); err != nil {
		return fmt.Errorf("normalizing references: %w", err)
	}

	desired := a.desired.DeepCopy()
	resource := AssetFeedSpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	resource.Name = a.id.String() // Name must be set for update

	paths, err := common.CompareProtoMessage(resource, a.actual, common.BasicDiff)
	if err != nil {
		return fmt.Errorf("comparing proto messages: %w", err)
	}

	if len(paths) == 0 {
		log.V(2).Info("no field needs update", "name", a.id)
		// Update status even if no fields changed in spec
		status := &krm.AssetFeedStatus{}

		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
		return updateOp.UpdateStatus(ctx, status, nil)
	}

	report := &structuredreporting.Diff{Object: updateOp.GetUnstructured()}
	for path := range paths {
		report.AddField(path, nil, nil)
	}
	structuredreporting.ReportDiff(ctx, report)

	log.V(2).Info("updating asset feed fields", "paths", sets.List(paths))

	updateMask := &fieldmaskpb.FieldMask{Paths: sets.List(paths)}
	req := &pb.UpdateFeedRequest{
		Feed:       resource,
		UpdateMask: updateMask,
	}
	_, err = a.gcpClient.UpdateFeed(ctx, req)
	if err != nil {
		return fmt.Errorf("updating asset feed %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully updated asset feed", "name", a.id)

	status := &krm.AssetFeedStatus{}

	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *feedAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.AssetFeed{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(AssetFeedSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	parentRef, id, err := krm.ParseFeedExternal(a.actual.Name)
	if err != nil {
		return nil, fmt.Errorf("parsing parent from name %q: %w", a.actual.Name, err)
	}
	if parentRef.ProjectID != "" {
		obj.Spec.Parent.ProjectRef = &refs.ProjectRef{External: parentRef.String()}
	} else if parentRef.FolderID != "" {
		obj.Spec.Parent.FolderRef = &refs.FolderRef{External: parentRef.String()}
	} else if parentRef.OrganizationID != "" {
		obj.Spec.Parent.OrganizationRef = &refs.OrganizationRef{External: parentRef.String()}
	} else {
		return nil, fmt.Errorf("unknown parent type in name %q", a.actual.Name)
	}

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	// Initialize u with the converted object map
	u = &unstructured.Unstructured{Object: uObj} // Assign to existing 'u'
	// Set GVK and Name, which will modify u.Object directly
	u.SetGroupVersionKind(krm.AssetFeedGVK)
	u.SetName(id)

	return u, nil
}

func (a *feedAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info(" deleting asset feed", "name", a.id)

	req := &pb.DeleteFeedRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteFeed(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info(" skipping delete for non-existent asset feed, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting asset feed %s: %w", a.id.String(), err)
	}
	log.V(2).Info(" successfully deleted asset feed", "name", a.id)

	return true, nil
}
