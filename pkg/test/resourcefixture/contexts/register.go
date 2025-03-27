// Copyright 2022 Google LLC
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

package contexts

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/lifecyclehandler"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/nasa9084/go-openapi"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclconversion "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/kcclite"
	dcllivestate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/livestate"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	testreconciler "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/controller/reconciler"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
)

type ResourceContext struct {
	ResourceGVK  schema.GroupVersionKind
	ResourceKind string

	SkipNoChange bool
	SkipUpdate   bool
	SkipDelete   bool

	// Hack: Optionally wait before getting the object in GCP. This is to work around some issues with troublesome
	// services in GCP that claim to be done with creating / updating the resource before it is actually available.
	PostModifyDelay time.Duration

	// If true, skip drift detection test.
	SkipDriftDetection bool

	// fields related to DCL-based resources
	DCLSchema *openapi.Schema
}

var (
	resourceContextMap = map[string]ResourceContext{}
	emptyGVK           = schema.GroupVersionKind{}
)

func GetResourceContext(fixture resourcefixture.ResourceFixture, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) (ResourceContext, error) {
	rc, ok := resourceContextMap[fixture.Name]
	if !ok {
		rc = ResourceContext{
			ResourceGVK:  fixture.GVK,
			ResourceKind: fixture.GVK.Kind,
		}
	}

	if rc.ResourceGVK == emptyGVK {
		rc.ResourceGVK = fixture.GVK
	}

	// If CRD has DCL controller label, fetch DCL schema.
	crd, err := crdloader.GetCRDForGVK(rc.ResourceGVK)
	if err != nil {
		return ResourceContext{}, err
	}
	if crd.GetLabels()[crdgeneration.Dcl2CRDLabel] == "true" {
		s, err := dclschemaloader.GetDCLSchemaForGVK(rc.ResourceGVK, serviceMetadataLoader, dclSchemaLoader)
		if err != nil {
			panic(fmt.Sprintf("error getting the DCL schema for GVK %v: %v", rc.ResourceGVK, err))
		}
		rc.DCLSchema = s
	}

	return rc, nil
}

func (rc ResourceContext) SupportsLabels(smLoader *servicemappingloader.ServiceMappingLoader, u *unstructured.Unstructured) bool {
	rt, err := testreconciler.ReconcilerTypeForObject(u)
	if err != nil {
		panic(fmt.Errorf("error getting reconciler type: %w", err))
	}

	if rt == testreconciler.ReconcilerTypeTerraform {
		// For tf based resources, resolve the label info from ResourceConfig
		resourceConfig := rc.getResourceConfig(smLoader)
		return resourceConfig.MetadataMapping.Labels != ""
	} else if rt == testreconciler.ReconcilerTypeDCL {
		_, _, found, err := dclextension.GetLabelsFieldSchema(rc.DCLSchema)
		if err != nil {
			panic(fmt.Errorf("error getting the DCL schema for labels field: %w", err))
		}
		return found
	} else {
		// Labels are not supported for Direct resources yet.
		return false
	}
}

func (rc ResourceContext) getResourceConfig(smLoader *servicemappingloader.ServiceMappingLoader) v1alpha1.ResourceConfig {
	for _, sm := range smLoader.GetServiceMappings() {
		for _, resourceConfig := range sm.Spec.Resources {
			if resourceConfig.Kind == rc.ResourceKind {
				return resourceConfig
			}
		}
	}
	panic(fmt.Errorf("no resource config found for kind: %v", rc.ResourceKind))
}

func (rc ResourceContext) IsAutoGenerated(smLoader *servicemappingloader.ServiceMappingLoader, u *unstructured.Unstructured) bool {
	rt, err := testreconciler.ReconcilerTypeForObject(u)
	if err != nil {
		panic(fmt.Errorf("error getting reconciler type: %w", err))
	}

	if rt == testreconciler.ReconcilerTypeTerraform {
		resourceConfig := rc.getResourceConfig(smLoader)
		return resourceConfig.AutoGenerated
	} else {
		return false
	}
}

func (rc ResourceContext) Create(ctx context.Context, _ *testing.T, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader, config *mmdcl.Config, dclConverter *dclconversion.Converter) (*unstructured.Unstructured, error) {
	rt, err := testreconciler.ReconcilerTypeForObject(u)
	if err != nil {
		panic(fmt.Errorf("error getting reconciler type: %w", err))
	}

	if rt == testreconciler.ReconcilerTypeDirect {
		return directCreate(ctx, u, c)
	}
	if rt == testreconciler.ReconcilerTypeDCL {
		return dclCreate(ctx, u, config, c, dclConverter, smLoader)
	}
	return terraformCreate(ctx, u, provider, c, smLoader)
}

func (rc ResourceContext) Get(ctx context.Context, _ *testing.T, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader, cfg *mmdcl.Config, dclConverter *dclconversion.Converter, httpClient *http.Client) (*unstructured.Unstructured, error) {
	rt, err := testreconciler.ReconcilerTypeForObject(u)
	if err != nil {
		panic(fmt.Errorf("error getting reconciler type: %w", err))
	}

	if rt == testreconciler.ReconcilerTypeDirect {
		result, err := directExport(ctx, u, c)
		if result == nil && err == nil {
			return nil, fmt.Errorf("%v uses direct controller and Export() is not implemented yet", u.GetKind())
		}
		return result, err
	}
	if rt == testreconciler.ReconcilerTypeDCL {
		return dclGet(ctx, u, cfg, c, dclConverter, smLoader)
	}
	return terraformGet(ctx, u, provider, c, smLoader)
}

func (rc ResourceContext) Delete(ctx context.Context, _ *testing.T, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader, cfg *mmdcl.Config, dclConverter *dclconversion.Converter, httpClient *http.Client) error {
	rt, err := testreconciler.ReconcilerTypeForObject(u)
	if err != nil {
		panic(fmt.Errorf("error getting reconciler type: %w", err))
	}

	if rt == testreconciler.ReconcilerTypeDirect {
		return directDelete(ctx, u, c)
	}
	if rt == testreconciler.ReconcilerTypeDCL {
		return dclDelete(ctx, u, cfg, c, dclConverter, smLoader)
	}
	return terraformDelete(ctx, u, provider, c, smLoader)
}

func terraformDelete(ctx context.Context, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) error {
	resource, liveState, err := getTerraformResourceAndLiveState(ctx, u, provider, c, smLoader)
	if err != nil {
		return err
	}
	if liveState.Empty() {
		return fmt.Errorf("resource '%v' of type '%v' cannot be deleted as it does not exist", u.GetName(), u.GroupVersionKind())
	}
	_, diagnostics := resource.TFResource.Apply(ctx, liveState, &terraform.InstanceDiff{Destroy: true}, provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return fmt.Errorf("error deleting resource: %w", err)
	}
	return err
}

func terraformCreate(ctx context.Context, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	resource, liveState, err := getTerraformResourceAndLiveState(ctx, u, provider, c, smLoader)
	if err != nil {
		return nil, err
	}
	if !liveState.Empty() {
		return nil, fmt.Errorf("resource '%v' of type '%v' cannot be created as it already exists", u.GetName(), u.GroupVersionKind())
	}
	config, _, err := krmtotf.KRMResourceToTFResourceConfig(resource, c, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error expanding resource configuration: %w", err)
	}
	diff, err := resource.TFResource.Diff(ctx, liveState, config, provider.Meta())
	if err != nil {
		return nil, fmt.Errorf("error calculating diff: %w", err)
	}
	newState, diagnostics := resource.TFResource.Apply(ctx, liveState, diff, provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return nil, fmt.Errorf("error applying resource change: %w", err)
	}
	return resourceToKRM(resource, newState)
}

func terraformGet(ctx context.Context, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	resource, liveState, err := getTerraformResourceAndLiveState(ctx, u, provider, c, smLoader)
	if err != nil {
		return nil, err
	}
	if liveState.Empty() {
		return nil, fmt.Errorf("resource '%v' of type '%v' is not found", u.GetName(), u.GroupVersionKind())
	}
	return resourceToKRM(resource, liveState)
}

func dclCreate(ctx context.Context, u *unstructured.Unstructured, config *mmdcl.Config, kubeClient client.Client, converter *dclconversion.Converter, serviceMappingLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	resource, err := newDCLResource(u, converter)
	if err != nil {
		return nil, err
	}
	liveLite, err := dcllivestate.FetchLiveState(ctx, resource, config, converter, serviceMappingLoader, kubeClient)
	if err != nil {
		return nil, err
	}
	if liveLite != nil {
		return nil, fmt.Errorf("resource '%v' of type '%v' cannot be created as it already exists", u.GetName(), u.GroupVersionKind())
	}
	lite, err := kcclite.ToKCCLite(resource, converter.MetadataLoader, converter.SchemaLoader, serviceMappingLoader, kubeClient)
	if err != nil {
		return nil, fmt.Errorf("error converting KCC full to KCC lite: %w", err)
	}
	dclObj, err := converter.KRMObjectToDCLObject(lite)
	if err != nil {
		return nil, fmt.Errorf("error converting KCC lite to DCL resource: %w", err)
	}
	createdDCLObj, err := dclunstruct.Apply(ctx, config, dclObj, dclcontroller.LifecycleParams...)
	if err != nil {
		return nil, fmt.Errorf("error applying the desired resource: %w", err)
	}
	// get the new state in KCC lite format
	newStateLite, err := converter.DCLObjectToKRMObject(createdDCLObj)
	if err != nil {
		return nil, fmt.Errorf("error converting DCL resource to KCC lite: %w", err)
	}
	return dclStateToKRM(resource, newStateLite, converter.MetadataLoader)
}

func dclGet(ctx context.Context, u *unstructured.Unstructured, config *mmdcl.Config, kubeClient client.Client, converter *dclconversion.Converter, serviceMappingLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	resource, err := newDCLResource(u, converter)
	if err != nil {
		return nil, err
	}
	liveLite, err := dcllivestate.FetchLiveState(ctx, resource, config, converter, serviceMappingLoader, kubeClient)
	if err != nil {
		return nil, err
	}
	if liveLite == nil {
		return nil, fmt.Errorf("resource '%v' of type '%v' is not found", u.GetName(), u.GroupVersionKind())
	}
	return dclStateToKRM(resource, liveLite, converter.MetadataLoader)
}

func dclDelete(ctx context.Context, u *unstructured.Unstructured, config *mmdcl.Config, kubeClient client.Client, converter *dclconversion.Converter, serviceMappingLoader *servicemappingloader.ServiceMappingLoader) error {
	resource, err := newDCLResource(u, converter)
	if err != nil {
		return err
	}
	lite, err := kcclite.ToKCCLiteBestEffort(resource, converter.MetadataLoader, converter.SchemaLoader, serviceMappingLoader, kubeClient)
	if err != nil {
		return fmt.Errorf("error converting KCC full to KCC lite: %w", err)
	}
	dclObj, err := converter.KRMObjectToDCLObject(lite)
	if err != nil {
		return fmt.Errorf("error converting KCC lite to DCL resource: %w", err)
	}
	if err := dclunstruct.Delete(ctx, config, dclObj); err != nil {
		return fmt.Errorf("error deleting the resource %v: %w", resource.GetNamespacedName(), err)
	}
	return nil
}

func newDCLResource(u *unstructured.Unstructured, converter *dclconversion.Converter) (*dcl.Resource, error) {
	s, err := dclschemaloader.GetDCLSchemaForGVK(u.GroupVersionKind(), converter.MetadataLoader, converter.SchemaLoader)
	if err != nil {
		return nil, err
	}
	resource, err := dcl.NewResource(u, s)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func dclStateToKRM(resource *dcl.Resource, liveState *unstructured.Unstructured, smLoader dclmetadata.ServiceMetadataLoader) (*unstructured.Unstructured, error) {
	spec, status, err := kcclite.ResolveSpecAndStatus(liveState, resource, smLoader)
	if err != nil {
		return nil, err
	}
	resource.Spec = spec
	resource.Status = status
	resource.Labels = liveState.GetLabels()
	return resource.MarshalAsUnstructured()
}

func resourceToKRM(resource *krmtotf.Resource, state *terraform.InstanceState) (*unstructured.Unstructured, error) {
	resource.Spec, resource.Status = krmtotf.ResolveSpecAndStatusWithResourceID(resource, state)
	resource.Labels = krmtotf.GetLabelsFromState(resource, state)
	// Apply post-actuation transformation.
	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource, state, nil); err != nil {
		return nil, fmt.Errorf("error applying post-actuation transformation to resource '%v': %w", resource.GetNamespacedName(), err)
	}
	return resource.MarshalAsUnstructured()
}

func getTerraformResourceAndLiveState(ctx context.Context, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*krmtotf.Resource,
	*terraform.InstanceState, error) {
	resource, err := newTerraformResource(u, provider, smLoader)
	if err != nil {
		return nil, nil, err
	}
	// Apply pre-actuation transformation.
	if err := resourceoverrides.Handler.PreActuationTransform(&resource.Resource); err != nil {
		return nil, nil, fmt.Errorf("error applying pre-actuation transformation to resource '%s': %w", u.GetName(), err)
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, provider, c, smLoader)
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching live state: %w", err)
	}
	return resource, liveState, nil
}

func newTerraformResource(u *unstructured.Unstructured, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader) (*krmtotf.Resource, error) {
	sm, err := smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return nil, err
	}
	resource, err := krmtotf.NewResource(u, sm, provider)
	if err != nil {
		return nil, fmt.Errorf("could not parse resource %s: %w", u.GetName(), err)
	}
	return resource, nil
}

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "is not found")
}

func getAdapter(ctx context.Context, u *unstructured.Unstructured, c client.Client) (directbase.Adapter, error) {
	gvk := u.GroupVersionKind()
	model, err := registry.GetModel(gvk.GroupKind())
	if err != nil {
		return nil, err
	}
	return model.AdapterForObject(ctx, c, u)
}

func directExport(ctx context.Context, u *unstructured.Unstructured, c client.Client) (*unstructured.Unstructured, error) {
	a, err := getAdapter(ctx, u, c)
	if err != nil {
		return nil, err
	}

	found, err := a.Find(ctx)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("GVK %s '%v' is not found", u.GroupVersionKind(), u.GetName())
	}

	unst, err := a.Export(ctx)
	if err != nil {
		return nil, err
	}
	return unst, nil
}

func directCreate(ctx context.Context, u *unstructured.Unstructured, c client.Client) (*unstructured.Unstructured, error) {
	a, err := getAdapter(ctx, u, c)
	if err != nil {
		return nil, err
	}

	found, err := a.Find(ctx)
	if err != nil {
		return nil, err
	}
	if found {
		return nil, fmt.Errorf("GVK %s '%v' already exist", u.GroupVersionKind(), u.GetName())
	}

	op := directbase.NewCreateOperation(lifecyclehandler.LifecycleHandler{}, c, u)
	err = a.Create(ctx, op)
	if err != nil {
		return nil, err
	}
	return directExport(ctx, u, c)
}

func directDelete(ctx context.Context, u *unstructured.Unstructured, c client.Client) error {
	maxRetries := 3
	baseDelay := 100 * time.Millisecond // starting delay
	var err error

	a, err := getAdapter(ctx, u, c)
	if err != nil {
		return err
	}
	// the very first call to find can be true but subsequent calls may return false
	// with the underlying Delete call already succeeding! So if we do not find this resource
	// in subsequent calls but found it originally, we can consider it deleted.
	firstFind := false

	for i := 0; i < maxRetries; i++ {
		found, err := a.Find(ctx)
		if err == nil && found {
			firstFind = true
			op := directbase.NewDeleteOperation(c, u)
			_, err = a.Delete(ctx, op)
			if err == nil {
				return nil // success
			}
		} else if err == nil && !found {
			if firstFind {
				return nil // success
			}
			return fmt.Errorf("GVK %s '%v' is not found", u.GroupVersionKind(), u.GetName())
		}

		// Exponential backoff
		time.Sleep(baseDelay * (1 << i)) // delays: 100ms * 2^0, 100ms * 2^1, 100ms * 2^2
	}

	return fmt.Errorf("failed to delete after %d retries: %w", maxRetries, err)
}
