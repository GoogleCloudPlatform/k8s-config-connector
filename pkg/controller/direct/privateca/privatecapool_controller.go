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

package privateca

import (
	"context"
	"fmt"
	"strings"

	iampb "cloud.google.com/go/iam/apiv1/iampb"
	api "cloud.google.com/go/security/privateca/apiv1"
	pb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privateca/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
)

func init() {
	registry.RegisterModel(krm.PrivateCACAPoolGVK, newCAPoolModel)
}

func newCAPoolModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	gcpClient, err := newGCPClient(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("building GCP client: %w", err)
	}
	return &caPoolModel{gcpClient: gcpClient}, nil
}

type caPoolModel struct {
	*gcpClient
}

// model implements the Model interface.
var _ directbase.Model = &caPoolModel{}

type caPoolAdapter struct {
	projectID string
	location  string
	caPoolID  string

	desired  *pb.CaPool
	actual   *pb.CaPool
	caClient *api.CertificateAuthorityClient
}

var _ directbase.Adapter = &caPoolAdapter{}

// AdapterForObject implements the Model interface.
func (m *caPoolModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	caClient, err := m.newCertificateAuthorityClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.PrivateCACAPool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	resourceID := direct.ValueOf(obj.Spec.ResourceID)
	if resourceID == "" {
		resourceID = obj.GetName()
	}
	if resourceID == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}

	location := obj.Spec.Location
	if location == "" {
		return nil, fmt.Errorf("cannot resolve location")
	}

	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), obj.Spec.ProjectRef)
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	mapCtx := &direct.MapContext{}
	desired := PrivateCACAPoolSpec_ToProto(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(u.GetLabels())

	return &caPoolAdapter{
		caPoolID:  resourceID,
		location:  location,
		projectID: projectID,
		desired:   desired,
		caClient:  caClient,
	}, nil
}

func (m *caPoolModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// Format is //privateca.googleapis.com/projects/PROJECT_ID/locations/LOCATION/caPools/CA_POOL_ID

	if !strings.HasPrefix(url, "//privateca.googleapis.com/") {
		return nil, nil
	}

	tokens := strings.Split(strings.TrimPrefix(url, "//privateca.googleapis.com/"), "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "caPools" {
		caClient, err := m.newCertificateAuthorityClient(ctx)
		if err != nil {
			return nil, err
		}

		return &caPoolAdapter{
			projectID: tokens[1],
			location:  tokens[3],
			caPoolID:  tokens[5],
			caClient:  caClient,
		}, nil
	}

	return nil, nil
}

// Delete implements the Adapter interface.
func (a *caPoolAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting PrivateCACAPool", "name", a.fullyQualifiedName())

	req := &pb.DeleteCaPoolRequest{Name: a.fullyQualifiedName()}
	op, err := a.caClient.DeleteCaPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent PrivateCACAPool, assuming it was already deleted", "name", a.fullyQualifiedName())
			return true, nil
		}
		return false, fmt.Errorf("deleting PrivateCACAPool %s: %w", a.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully deleted PrivateCACAPool", "name", a.fullyQualifiedName())

	err = op.Wait(ctx)
	if err != nil {
		return false, fmt.Errorf("waiting delete PrivateCACAPool %s: %w", a.fullyQualifiedName(), err)
	}
	return true, nil
}

// Create implements the Adapter interface.
func (a *caPoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating PrivateCACAPool", "id", a.fullyQualifiedName())

	parent := fmt.Sprintf("projects/%s/locations/%s", a.projectID, a.location)

	req := &pb.CreateCaPoolRequest{
		Parent:   parent,
		CaPoolId: a.caPoolID,
		CaPool:   a.desired,
	}
	op, err := a.caClient.CreateCaPool(ctx, req)
	if err != nil {
		return fmt.Errorf("creating PrivateCACAPool %s: %w", a.fullyQualifiedName(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting PrivateCACAPool %s creation: %w", a.fullyQualifiedName(), err)
	}
	log.V(2).Info("successfully created PrivateCACAPool", "name", a.fullyQualifiedName())

	return a.updateStatus(ctx, createOp, created)
}

// Update implements the Adapter interface.
func (a *caPoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating PrivateCACAPool", "name", a.fullyQualifiedName())

	diffs, updateMask, err := comparePrivateCACAPool(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if diffs.HasDiff() {
		diffs.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffs)

		a.desired.Name = a.fullyQualifiedName()
		req := &pb.UpdateCaPoolRequest{
			UpdateMask: updateMask,
			CaPool:     a.desired,
		}
		op, err := a.caClient.UpdateCaPool(ctx, req)
		if err != nil {
			return fmt.Errorf("updating PrivateCACAPool %s: %w", a.fullyQualifiedName(), err)
		}
		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting update PrivateCACAPool %s: %w", a.fullyQualifiedName(), err)
		}
		log.V(2).Info("successfully updated PrivateCACAPool", "name", a.fullyQualifiedName())
		latest = updated
	}

	return a.updateStatus(ctx, updateOp, latest)
}

// Export implements the Adapter interface.
func (a *caPoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}
	u := &unstructured.Unstructured{}

	obj := &krm.PrivateCACAPool{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(PrivateCACAPoolSpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}

	obj.Spec.ProjectRef = &refs.ProjectRef{Name: a.projectID}
	obj.Spec.Location = a.location
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}
	u.Object = uObj
	u.SetName(a.actual.Name)
	u.SetGroupVersionKind(krm.PrivateCACAPoolGVK)
	return u, nil
}

func comparePrivateCACAPool(ctx context.Context, actual, desired *pb.CaPool) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, PrivateCACAPoolSpec_FromProto, PrivateCACAPoolSpec_ToProto)
	if err != nil {
		return nil, nil, err
	}
	maskedActual.Name = desired.Name
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, desired.ProtoReflect(), maskedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}

func (a *caPoolAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *pb.CaPool) error {
	status := &krm.PrivateCACAPoolStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

// Find implements the Adapter interface.
func (a *caPoolAdapter) Find(ctx context.Context) (bool, error) {
	if a.caPoolID == "" {
		return false, nil
	}

	req := &pb.GetCaPoolRequest{
		Name: a.fullyQualifiedName(),
	}
	logMetric, err := a.caClient.GetCaPool(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting logMetric %q: %w", a.fullyQualifiedName(), err)
	}

	a.actual = logMetric

	return true, nil
}

func (a *caPoolAdapter) GetIAMPolicy(ctx context.Context) (*iampb.Policy, error) {
	if a.caPoolID == "" {
		return nil, fmt.Errorf("cannot get iam policy for missing resource")
	}

	req := &iampb.GetIamPolicyRequest{
		Resource: a.fullyQualifiedName(),
	}
	policy, err := a.caClient.GetIamPolicy(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("getting iam policy for %q: %w", a.fullyQualifiedName(), err)
	}

	return policy, nil
}

func (a *caPoolAdapter) SetIAMPolicy(ctx context.Context, policy *iampb.Policy) (*iampb.Policy, error) {
	if a.caPoolID == "" {
		return nil, fmt.Errorf("cannot get iam policy for missing resource")
	}

	req := &iampb.SetIamPolicyRequest{
		Resource: a.fullyQualifiedName(),
		Policy:   policy,
	}
	newPolicy, err := a.caClient.SetIamPolicy(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("setting iam policy for %q: %w", a.fullyQualifiedName(), err)
	}

	return newPolicy, nil
}

func (a *caPoolAdapter) fullyQualifiedName() string {
	return fmt.Sprintf("projects/%s/locations/%s/caPools/%s", a.projectID, a.location, a.caPoolID)
}
