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
// proto.service: google.cloud.asset.v1.AssetService
// proto.message: google.cloud.asset.v1.SavedQuery
// crd.type: CloudAssetSavedQuery
// crd.version: v1alpha1

package cloudasset

import (
	"context"
	"fmt"

	gcp "cloud.google.com/go/asset/apiv1"
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudasset/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.CloudAssetSavedQueryGVK, NewSavedQueryModel)
}

func NewSavedQueryModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &savedQueryModel{config: *config}, nil
}

var _ directbase.Model = &savedQueryModel{}

type savedQueryModel struct {
	config config.ControllerConfig
}

func (m *savedQueryModel) client(ctx context.Context, projectID string) (*gcp.Client, error) {
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
		return nil, fmt.Errorf("building cloudasset saved query client: %w", err)
	}

	return gcpClient, err
}

func (m *savedQueryModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.CloudAssetSavedQuery{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	res, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id := res.(*krm.CloudAssetSavedQueryIdentity)

	gcpClient, err := m.client(ctx, id.Project)
	if err != nil {
		return nil, err
	}

	mapCtx := &direct.MapContext{}
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	desired := CloudAssetSavedQuerySpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	return &savedQueryAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   desired,
		reader:    reader,
	}, nil
}

func (m *savedQueryModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type savedQueryAdapter struct {
	gcpClient *gcp.Client
	id        *krm.CloudAssetSavedQueryIdentity
	desired   *pb.SavedQuery
	actual    *pb.SavedQuery
	reader    client.Reader
}

var _ directbase.Adapter = &savedQueryAdapter{}

func (a *savedQueryAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting cloudasset saved query", "name", a.id)

	name := a.id.String()
	if a.actual != nil {
		name = a.actual.Name
	}
	req := &pb.GetSavedQueryRequest{Name: name}
	actual, err := a.gcpClient.GetSavedQuery(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting cloudasset saved query %q from gcp: %w", a.id.String(), err)
	}

	a.actual = actual
	return true, nil
}

func (a *savedQueryAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating cloudasset saved query", "name", a.id)

	// resource.Name is output only for SavedQuery

	req := &pb.CreateSavedQueryRequest{
		Parent:       "projects/" + a.id.Project,
		SavedQueryId: a.id.SavedQuery,
		SavedQuery:   a.desired,
	}

	// Attempt to create the saved query
	actual, err := a.gcpClient.CreateSavedQuery(ctx, req)

	// Regardless of CreateSavedQuery result (success or AlreadyExists),
	// we need to ensure a.actual is populated by calling Find.
	found := false
	var findErr error

	if err == nil {
		log.V(2).Info("successfully created cloudasset saved query in gcp, attempting to fetch it", "name", a.id)
		a.actual = actual
		found = true
	} else if errors.IsAlreadyExists(err) {
		// Resource already exists. Log and attempt to Find to populate a.actual.
		log.V(2).Info("cloudasset saved query already exists during creation attempt, attempting to fetch it", "name", a.id.String(), "warning", err)
		found, findErr = a.Find(ctx)
		if findErr != nil {
			return fmt.Errorf("fetching existing cloudasset saved query %q after CreateSavedQuery returned AlreadyExists failed: %w", a.id.String(), findErr)
		}
		if !found {
			// This is unexpected if CreateSavedQuery returned AlreadyExists
			return fmt.Errorf("cloudasset saved query %q not found even though CreateSavedQuery returned AlreadyExists", a.id.String())
		}
	} else {
		// Other error during creation
		return fmt.Errorf("creating cloudasset saved query %s: %w", a.id.String(), err)
	}

	return a.updateStatus(ctx, createOp, a.actual)
}

func (a *savedQueryAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating cloudasset saved query", "name", a.id, "actual", a.actual.Name)

	a.desired.Name = a.actual.Name // Name must be set for update based on actual state

	diffs, updateMask, err := compareCloudAssetSavedQuery(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("no fields need update", "name", a.id)
		return a.updateStatus(ctx, updateOp, a.actual)
	}

	structuredreporting.ReportDiff(ctx, diffs)

	req := &pb.UpdateSavedQueryRequest{
		SavedQuery: a.desired,
		UpdateMask: updateMask,
	}
	actual, err := a.gcpClient.UpdateSavedQuery(ctx, req)
	if err != nil {
		return fmt.Errorf("updating cloudasset saved query %s: %w", a.id.String(), err)
	}
	a.actual = actual // Update internal actual state with the response
	log.V(2).Info("successfully updated cloudasset saved query", "name", a.id)

	return a.updateStatus(ctx, updateOp, a.actual)
}

func (a *savedQueryAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() must be called before Export()")
	}
	mapCtx := &direct.MapContext{}

	obj := &krm.CloudAssetSavedQuery{}
	obj.Spec = direct.ValueOf(CloudAssetSavedQuerySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	// Set parent reference from the 'name' field
	parsed, match, err := krm.CloudAssetSavedQueryIdentityFormatRelative.Parse(a.actual.Name)
	if err != nil {
		return nil, fmt.Errorf("parsing parent from name %q: %w", a.actual.Name, err)
	}
	if !match {
		parsed, match, err = krm.CloudAssetSavedQueryIdentityFormat.Parse(a.actual.Name)
		if err != nil || !match {
			return nil, fmt.Errorf("unknown parent format in name %q", a.actual.Name)
		}
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{External: parsed.Project}

	// Set observed state
	status := CloudAssetSavedQueryStatus_FromProto(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Status = *status

	// Convert to unstructured
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, fmt.Errorf("converting CloudAssetSavedQuery to unstructured: %w", err)
	}
	u := &unstructured.Unstructured{Object: uObj}

	// Set standard Kubernetes metadata
	u.SetGroupVersionKind(krm.CloudAssetSavedQueryGVK)
	u.SetName(parsed.SavedQuery) // Use the parsed ID from the name field
	u.SetNamespace(a.id.Project)

	// Set external reference annotation if available
	if status.ExternalRef != nil {
		annotations := u.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}
		annotations["cnrm.cloud.google.com/external-ref"] = *status.ExternalRef
		u.SetAnnotations(annotations)
	}

	return u, nil
}

func (a *savedQueryAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting cloudasset saved query", "name", a.id)

	req := &pb.DeleteSavedQueryRequest{Name: a.id.String()}
	err := a.gcpClient.DeleteSavedQuery(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			// Return success if not found (assume it was already deleted).
			log.V(2).Info("skipping delete for non-existent cloudasset saved query, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting cloudasset saved query %s: %w", a.id.String(), err)
	}
	log.V(2).Info("successfully deleted cloudasset saved query", "name", a.id)

	return true, nil
}

func (a *savedQueryAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.SavedQuery) error {
	mapCtx := &direct.MapContext{}
	status := CloudAssetSavedQueryStatus_FromProto(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func compareCloudAssetSavedQuery(ctx context.Context, actual, desired *pb.SavedQuery) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, CloudAssetSavedQuerySpec_FromProto, CloudAssetSavedQuerySpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name // Restore any non-spec identifier fields if needed

	clonedDesired := proto.CloneOf(desired)

	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

// Helper function to get owner reference. Needed for parent resolution.
func ownerFromObject(obj metav1.Object) *types.NamespacedName {
	return &types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
}
