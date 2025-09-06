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
// proto.service: google.cloud.networksecurity.v1beta1.NetworkSecurity
// proto.message: google.cloud.networksecurity.v1beta1.AuthorizationPolicy
// crd.type: NetworkSecurityAuthorizationPolicy
// crd.version: v1beta1

package networksecurity

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networksecurity/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"

	networksecurity "cloud.google.com/go/networksecurity/apiv1beta1"
	"cloud.google.com/go/networksecurity/apiv1beta1/networksecuritypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func init() {
	registry.RegisterModel(krm.NetworkSecurityAuthorizationPolicyGVK, NewAuthorizationPolicyModel)
}

// cleanupAuthorizationPolicy removes empty objects from the policy's lists.
// This is to avoid perpetual diffs when the API returns empty objects in lists.
func cleanupAuthorizationPolicy(policy *networksecuritypb.AuthorizationPolicy) {
	if policy == nil {
		return
	}

	var cleanedRules []*networksecuritypb.AuthorizationPolicy_Rule
	for _, rule := range policy.Rules {
		if rule == nil || proto.Equal(rule, &networksecuritypb.AuthorizationPolicy_Rule{}) {
			continue
		}

		var cleanedSources []*networksecuritypb.AuthorizationPolicy_Rule_Source
		for _, source := range rule.Sources {
			if source != nil && !proto.Equal(source, &networksecuritypb.AuthorizationPolicy_Rule_Source{}) {
				cleanedSources = append(cleanedSources, source)
			}
		}
		rule.Sources = cleanedSources

		var cleanedDests []*networksecuritypb.AuthorizationPolicy_Rule_Destination
		for _, dest := range rule.Destinations {
			if dest != nil && !proto.Equal(dest, &networksecuritypb.AuthorizationPolicy_Rule_Destination{}) {
				cleanedDests = append(cleanedDests, dest)
			}
		}
		rule.Destinations = cleanedDests

		// A rule is empty if it has no sources and no destinations.
		// The proto has only sources and destinations fields.
		if len(rule.Sources) > 0 || len(rule.Destinations) > 0 {
			cleanedRules = append(cleanedRules, rule)
		}
	}
	policy.Rules = cleanedRules
}

func NewAuthorizationPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &authorizationPolicyModel{config: *config}, nil
}

var _ directbase.Model = &authorizationPolicyModel{}

type authorizationPolicyModel struct {
	config config.ControllerConfig
}

func (m *authorizationPolicyModel) client(ctx context.Context) (*networksecurity.Client, error) {
	var opts []option.ClientOption
	config := m.config

	restClientOpts, err := config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	opts = append(opts, restClientOpts...)

	gcpClient, err := networksecurity.NewRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building networksecurity authorizationpolicy client: %w", err)
	}

	return gcpClient, err
}

func (m *authorizationPolicyModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.NetworkSecurityAuthorizationPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	id, err := krm.NewAuthorizationPolicyIdentity(ctx, reader, obj)
	if err != nil {
		return nil, err
	}

	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &authorizationPolicyAdapter{
		gcpClient: gcpClient,
		id:        id,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *authorizationPolicyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type authorizationPolicyAdapter struct {
	gcpClient *networksecurity.Client
	id        *krm.AuthorizationPolicyIdentity
	desired   *krm.NetworkSecurityAuthorizationPolicy
	actual    *networksecuritypb.AuthorizationPolicy
	reader    client.Reader
}

var _ directbase.Adapter = &authorizationPolicyAdapter{}

// Find retrieves the GCP resource.
// Return true means the object is found. This triggers Adapter `Update` call.
// Return false means the object is not found. This triggers Adapter `Create` call.
// Return a non-nil error requeues the requests.
func (a *authorizationPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting networksecurity authorizationpolicy", "name", a.id)

	req := &networksecuritypb.GetAuthorizationPolicyRequest{Name: a.id.String()}
	actual, err := a.gcpClient.GetAuthorizationPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting networksecurity authorizationpolicy %q from gcp: %w", a.id.String(), err)
	}

	cleanupAuthorizationPolicy(actual)
	a.actual = actual
	return true, nil
}

// Create creates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *authorizationPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating networksecurity authorizationpolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := NetworkSecurityAuthorizationPolicySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Labels = label.NewGCPLabelsFromK8sLabels(a.desired.Labels)
	req := &networksecuritypb.CreateAuthorizationPolicyRequest{
		Parent:                a.id.Parent().String(),
		AuthorizationPolicyId: a.id.ID(),
		AuthorizationPolicy:   desired,
	}
	op, err := a.gcpClient.CreateAuthorizationPolicy(ctx, req)
	if err != nil {
		return fmt.Errorf("creating networksecurity authorizationpolicy %s: %w", a.id.String(), err)
	}
	created, err := op.Wait(ctx)
	if err != nil {
		return fmt.Errorf("waiting for create of networksecurity authorizationpolicy %s: %w", a.id.String(), err)
	}

	log.V(2).Info("successfully created networksecurity authorizationpolicy in gcp", "name", a.id)

	status := &krm.NetworkSecurityAuthorizationPolicyStatus{}

	status.CreateTime = direct.StringTimestamp_FromProto(mapCtx, created.GetCreateTime())
	status.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, created.GetUpdateTime())
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return createOp.UpdateStatus(ctx, status, nil)
}

// Update updates the resource in GCP based on `spec` and update the Config Connector object `status` based on the GCP response.
func (a *authorizationPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating NetworkSecurityAuthorizationPolicy", "name", a.id)
	mapCtx := &direct.MapContext{}

	desired := NetworkSecurityAuthorizationPolicySpec_ToProto(mapCtx, &a.desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	desired.Name = a.id.String()
	desired.Labels = label.NewGCPLabelsFromK8sLabels(a.desired.Labels)

	diff, err := common.CompareProtoMessage(desired, a.actual, common.BasicDiff)
	if err != nil {
		return err
	}
	log.V(2).Info("diff", "diff", diff)

	status := &krm.NetworkSecurityAuthorizationPolicyStatus{}

	if len(diff) == 0 {
		log.V(2).Info("no update needed")
		status.CreateTime = direct.StringTimestamp_FromProto(mapCtx, a.actual.GetCreateTime())
		status.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, a.actual.GetUpdateTime())
	} else {
		req := &networksecuritypb.UpdateAuthorizationPolicyRequest{
			AuthorizationPolicy: desired,
			UpdateMask:          &fieldmaskpb.FieldMask{Paths: sets.List(diff)},
		}
		op, err := a.gcpClient.UpdateAuthorizationPolicy(ctx, req)
		if err != nil {
			if direct.IsNotFound(err) {
				return nil
			}
			return fmt.Errorf("updating networksecurity authorizationpolicy %s: %w", a.id.String(), err)
		}

		updated, err := op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting for update of networksecurity authorizationpolicy %s: %w", a.id.String(), err)
		}

		log.V(2).Info("successfully updated networksecurity authorizationpolicy in gcp", "name", a.id)

		status.CreateTime = direct.StringTimestamp_FromProto(mapCtx, updated.GetCreateTime())
		status.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, updated.GetUpdateTime())
		if mapCtx.Err() != nil {
			return mapCtx.Err()
		}
	}
	status.ExternalRef = direct.PtrTo(a.id.String())
	return updateOp.UpdateStatus(ctx, status, nil)
}

// Export implements the Adapter interface.
func (a *authorizationPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	log := klog.FromContext(ctx)
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.NetworkSecurityAuthorizationPolicy{}
	mapCtx := &direct.MapContext{}
	obj.Spec = direct.ValueOf(NetworkSecurityAuthorizationPolicySpec_FromProto(mapCtx, a.actual))
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	obj.Labels = a.actual.Labels

	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ID())
	u.SetGroupVersionKind(krm.NetworkSecurityAuthorizationPolicyGVK)

	log.Info("exported object", "obj", u, "gvk", u.GroupVersionKind())
	return u, nil
}

// Delete the resource from GCP service when the corresponding Config Connector resource is deleted.
func (a *authorizationPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting networksecurity authorizationpolicy", "name", a.id)

	req := &networksecuritypb.DeleteAuthorizationPolicyRequest{Name: a.id.String()}
	op, err := a.gcpClient.DeleteAuthorizationPolicy(ctx, req)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("deleting networksecurity authorizationpolicy %s: %w", a.id.String(), err)
	}

	if err := op.Wait(ctx); err != nil {
		return false, fmt.Errorf("waiting for delete of networksecurity authorizationpolicy %s: %w", a.id.String(), err)
	}
	return true, nil
}
