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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/privateca/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
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

	desired  *krm.PrivateCACAPool
	actual   *pb.CaPool
	caClient *api.CertificateAuthorityClient
}

var _ directbase.Adapter = &caPoolAdapter{}

// AdapterForObject implements the Model interface.
func (m *caPoolModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	caClient, err := m.newCertificateAuthorityClient(ctx)
	if err != nil {
		return nil, err
	}

	obj := &krm.PrivateCACAPool{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
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

	projectRef, err := refs.ResolveProject(ctx, reader, obj.GetNamespace(), refs.AsProjectRef(&obj.Spec.ProjectRef))
	if err != nil {
		return nil, err
	}
	projectID := projectRef.ProjectID
	if projectID == "" {
		return nil, fmt.Errorf("cannot resolve project")
	}

	return &caPoolAdapter{
		caPoolID:  resourceID,
		location:  location,
		projectID: projectID,
		desired:   obj,
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
	return false, fmt.Errorf("not implemented")
}

// Create implements the Adapter interface.
func (a *caPoolAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return fmt.Errorf("not implemented")
}

// Update implements the Adapter interface.
func (a *caPoolAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return fmt.Errorf("not implemented")
}

// Export implements the Adapter interface.
func (a *caPoolAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, fmt.Errorf("not implemented")
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
