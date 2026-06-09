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

package dns

import (
	"context"
	"errors"
	"fmt"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/dns/v1"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DNSPolicyGVK, NewPolicyModel)
}

func NewPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &policyModel{config: config}, nil
}

var _ directbase.Model = &policyModel{}

type policyModel struct {
	config *config.ControllerConfig
}

func (m *policyModel) client(ctx context.Context) (*api.Service, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := api.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building DNS client: %w", err)
	}
	return gcpClient, err
}

func (m *policyModel) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DNSPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Always call common.NormalizeReferences to resolve references
	if err := common.NormalizeReferences(ctx, reader, obj, nil); err != nil {
		return nil, fmt.Errorf("normalizing references: %w", err)
	}

	idVal, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := idVal.(*krm.DNSPolicyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format here, so we follow the pattern in the skill.
	mapCtx := &direct.MapContext{}
	desired := DNSPolicySpec_ToAPI(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.Policy

	if err := normalizePolicyURLs(desired); err != nil {
		return nil, fmt.Errorf("normalizing network URLs: %w", err)
	}

	// Get DNS GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &DNSPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *policyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DNSPolicyAdapter struct {
	id        *krm.DNSPolicyIdentity
	gcpClient *api.Service
	desired   *api.Policy
	actual    *api.Policy
}

var _ directbase.Adapter = &DNSPolicyAdapter{}

func (a *DNSPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DNSPolicy", "name", a.id)

	resource, err := a.gcpClient.Policies.Get(a.id.Project, a.id.Policy).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DNSPolicy %q: %w", a.id, err)
	}

	a.actual = resource
	return true, nil
}

func (a *DNSPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DNSPolicy", "name", a.id)

	created, err := a.gcpClient.Policies.Create(a.id.Project, a.desired).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating DNSPolicy %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created DNSPolicy", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *DNSPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DNSPolicy", "name", a.id)

	diffResults, err := comparePolicy(ctx, a.actual, a.desired)
	if err != nil {
		return err
	}

	latest := a.actual
	if !diffResults.HasDiff() {
		log.V(2).Info("no field needs update", "name", a.id)
	} else {
		log.V(2).Info("fields need update", "name", a.id)

		diffResults.Object = updateOp.GetUnstructured()
		structuredreporting.ReportDiff(ctx, diffResults)

		desired := common.DeepCopy(a.desired)
		// Workaround: Id is required in update calls
		desired.Id = a.actual.Id

		resp, err := a.gcpClient.Policies.Update(a.id.Project, a.id.Policy, desired).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating DNSPolicy %s: %w", a.id, err)
		}
		log.V(2).Info("successfully updated DNSPolicy", "name", a.id)
		latest = resp.Policy
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func comparePolicy(ctx context.Context, actual, desired *api.Policy) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DNSPolicySpec_FromAPI, DNSPolicySpec_ToAPI)
	if err != nil {
		return nil, err
	}

	desired = common.DeepCopy(desired)

	setDefaults := func(p *api.Policy) error {
		if err := normalizePolicyURLs(p); err != nil {
			return err
		}
		p.Kind = "dns#policy"
		if p.AlternativeNameServerConfig != nil {
			p.AlternativeNameServerConfig.Kind = "dns#policyAlternativeNameServerConfig"
			for i := range p.AlternativeNameServerConfig.TargetNameServers {
				p.AlternativeNameServerConfig.TargetNameServers[i].Kind = "dns#policyTargetNameServer"
			}
		}
		for i := range p.Networks {
			p.Networks[i].Kind = "dns#policyNetwork"
		}
		return nil
	}

	if err := setDefaults(maskedActual); err != nil {
		return nil, fmt.Errorf("normalizing actual network URLs: %w", err)
	}
	if err := setDefaults(desired); err != nil {
		return nil, fmt.Errorf("normalizing desired network URLs: %w", err)
	}

	diffs, _, err := diffs.GoogleAPI.Diff(ctx, maskedActual, desired)
	return diffs, err
}

func (a *DNSPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.Policy) error {
	mapCtx := &direct.MapContext{}
	status := DNSPolicyStatus_FromAPI(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *DNSPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DNSPolicy{}
	mapCtx := &direct.MapContext{}
	spec := DNSPolicySpec_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	spec.ResourceID = direct.LazyPtr(a.id.Policy)
	obj.Spec = *spec
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.Policy)
	u.SetGroupVersionKind(krm.DNSPolicyGVK)

	return u, nil
}

func (a *DNSPolicyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DNSPolicy", "name", a.id)

	err := a.gcpClient.Policies.Delete(a.id.Project, a.id.Policy).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DNSPolicy, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DNSPolicy %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DNSPolicy", "name", a.id)

	return true, nil
}

func normalizePolicyURLs(policy *api.Policy) error {
	var errs []error
	prefix := "https://www.googleapis.com/compute/v1/"
	for i := range policy.Networks {
		netUrl := policy.Networks[i].NetworkUrl
		if netUrl != "" {
			id := &computev1beta1.NetworkIdentity{}
			if err := id.FromExternal(netUrl); err != nil {
				errs = append(errs, fmt.Errorf("invalid network URL %q: %w", netUrl, err))
			} else {
				policy.Networks[i].NetworkUrl = prefix + id.String()
			}
		}
	}
	return errors.Join(errs...)
}
