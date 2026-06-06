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
	"time"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/diffs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/label"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/mappers"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	api "google.golang.org/api/dns/v1"
	"google.golang.org/api/option"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
)

func init() {
	registry.RegisterModel(krm.DNSManagedZoneGVK, NewModel)
}

func NewModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &model{config: config}, nil
}

var _ directbase.Model = &model{}

type model struct {
	config *config.ControllerConfig
}

func (m *model) client(ctx context.Context) (*api.Service, error) {
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

func (m *model) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.DNSManagedZone{}
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
	id, ok := idVal.(*krm.DNSManagedZoneIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", idVal)
	}

	// Convert the KRM spec to API format here, so we follow the pattern in the skill.
	mapCtx := &direct.MapContext{}
	desired := DNSManagedZoneSpec_ToAPI(mapCtx, &obj.Spec)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	desired.Name = id.ManagedZone
	desired.Labels = label.GCPLabels(obj)

	if err := normalizeURLs(desired); err != nil {
		return nil, fmt.Errorf("normalizing network URLs: %w", err)
	}

	// Get DNS GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, err
	}

	return &DNSManagedZoneAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   desired,
	}, nil
}

func (m *model) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

type DNSManagedZoneAdapter struct {
	id        *krm.DNSManagedZoneIdentity
	gcpClient *api.Service
	desired   *api.ManagedZone
	actual    *api.ManagedZone
}

var _ directbase.Adapter = &DNSManagedZoneAdapter{}

func (a *DNSManagedZoneAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting DNSManagedZone", "name", a.id)

	resource, err := a.gcpClient.ManagedZones.Get(a.id.Project, a.id.ManagedZone).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting DNSManagedZone %q: %w", a.id, err)
	}

	a.actual = resource
	return true, nil
}

func (a *DNSManagedZoneAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating DNSManagedZone", "name", a.id)

	created, err := a.gcpClient.ManagedZones.Create(a.id.Project, a.desired).Context(ctx).Do()
	if err != nil {
		return fmt.Errorf("creating DNSManagedZone %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created DNSManagedZone", "name", a.id)

	return a.updateStatus(ctx, createOp, created)
}

func (a *DNSManagedZoneAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating DNSManagedZone", "name", a.id)

	diffResults, err := compareZone(ctx, a.actual, a.desired)
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

		op, err := a.gcpClient.ManagedZones.Update(a.id.Project, a.id.ManagedZone, a.desired).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("updating DNSManagedZone %s: %w", a.id, err)
		}
		log.V(2).Info("successfully triggered update of DNSManagedZone, waiting for operation", "name", a.id, "op", op.Id)
		if err := a.waitForDNSOp(ctx, op); err != nil {
			return fmt.Errorf("waiting for DNSManagedZone update operation %s: %w", op.Id, err)
		}
		log.V(2).Info("DNSManagedZone update operation completed", "name", a.id, "op", op.Id)

		latest, err = a.gcpClient.ManagedZones.Get(a.id.Project, a.id.ManagedZone).Context(ctx).Do()
		if err != nil {
			return fmt.Errorf("getting DNSManagedZone after update %s: %w", a.id, err)
		}
	}

	return a.updateStatus(ctx, updateOp, latest)
}

func compareZone(ctx context.Context, actual, desired *api.ManagedZone) (*structuredreporting.Diff, error) {
	maskedActual, err := mappers.OnlySpecFields(actual, DNSManagedZoneSpec_FromAPI, DNSManagedZoneSpec_ToAPI)
	if err != nil {
		return nil, err
	}

	desired = common.DeepCopy(desired)

	setDefaults := func(zone *api.ManagedZone) error {
		if err := normalizeURLs(zone); err != nil {
			return err
		}
		zone.Kind = "dns#managedZone"
		if zone.DnssecConfig != nil {
			zone.DnssecConfig.Kind = "dns#managedZoneDnsSecConfig"
			for i := range zone.DnssecConfig.DefaultKeySpecs {
				zone.DnssecConfig.DefaultKeySpecs[i].Kind = "dns#dnsKeySpec"
			}
		}
		if zone.PeeringConfig != nil && zone.PeeringConfig.TargetNetwork != nil {
			zone.PeeringConfig.TargetNetwork.Kind = "dns#managedZonePeeringConfigTargetNetwork"
		}
		if zone.PrivateVisibilityConfig != nil {
			zone.PrivateVisibilityConfig.Kind = "dns#managedZonePrivateVisibilityConfig"
			for i := range zone.PrivateVisibilityConfig.GkeClusters {
				zone.PrivateVisibilityConfig.GkeClusters[i].Kind = "dns#managedZonePrivateVisibilityConfigGKECluster"
			}
			for i := range zone.PrivateVisibilityConfig.Networks {
				zone.PrivateVisibilityConfig.Networks[i].Kind = "dns#managedZonePrivateVisibilityConfigNetwork"
			}
		}
		if zone.ServiceDirectoryConfig != nil {
			zone.ServiceDirectoryConfig.Kind = "dns#managedZoneServiceDirectoryConfig"
			if zone.ServiceDirectoryConfig.Namespace != nil {
				zone.ServiceDirectoryConfig.Namespace.Kind = "dns#managedZoneServiceDirectoryConfigNamespace"
			}
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

func (a *DNSManagedZoneAdapter) updateStatus(ctx context.Context, op directbase.Operation, latest *api.ManagedZone) error {
	mapCtx := &direct.MapContext{}
	status := DNSManagedZoneStatus_FromAPI(mapCtx, latest)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *DNSManagedZoneAdapter) waitForDNSOp(ctx context.Context, op *api.Operation) error {
	if op.Status == "done" {
		return nil
	}
	pollInterval := 2 * time.Second
	return common.WaitForDoneOrTimeout(ctx, pollInterval, func() (bool, error) {
		currentOp, err := a.gcpClient.ManagedZoneOperations.Get(a.id.Project, a.id.ManagedZone, op.Id).Context(ctx).Do()
		if err != nil {
			return false, err
		}
		if currentOp.Status == "done" {
			return true, nil
		}
		return false, nil
	})
}

func (a *DNSManagedZoneAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("Find() not called")
	}

	obj := &krm.DNSManagedZone{}
	mapCtx := &direct.MapContext{}
	spec := DNSManagedZoneSpec_FromAPI(mapCtx, a.actual)
	if mapCtx.Err() != nil {
		return nil, mapCtx.Err()
	}
	spec.ResourceID = direct.LazyPtr(a.id.ManagedZone)
	obj.Spec = *spec
	uObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return nil, err
	}

	u := &unstructured.Unstructured{Object: uObj}
	u.SetName(a.id.ManagedZone)
	u.SetGroupVersionKind(krm.DNSManagedZoneGVK)

	return u, nil
}

func (a *DNSManagedZoneAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting DNSManagedZone", "name", a.id)

	err := a.gcpClient.ManagedZones.Delete(a.id.Project, a.id.ManagedZone).Context(ctx).Do()
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("skipping delete for non-existent DNSManagedZone, assuming it was already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting DNSManagedZone %s: %w", a.id, err)
	}
	log.V(2).Info("successfully deleted DNSManagedZone", "name", a.id)

	return true, nil
}

func normalizeURLs(zone *api.ManagedZone) error {
	var errs []error
	prefix := "https://www.googleapis.com/compute/v1/"
	if zone.PeeringConfig != nil && zone.PeeringConfig.TargetNetwork != nil {
		netUrl := zone.PeeringConfig.TargetNetwork.NetworkUrl
		if netUrl != "" {
			id := &computev1beta1.NetworkIdentity{}
			if err := id.FromExternal(netUrl); err != nil {
				errs = append(errs, fmt.Errorf("invalid peering config target network URL %q: %w", netUrl, err))
			} else {
				zone.PeeringConfig.TargetNetwork.NetworkUrl = prefix + id.String()
			}
		}
	}
	if zone.PrivateVisibilityConfig != nil {
		for i := range zone.PrivateVisibilityConfig.Networks {
			netUrl := zone.PrivateVisibilityConfig.Networks[i].NetworkUrl
			if netUrl != "" {
				id := &computev1beta1.NetworkIdentity{}
				if err := id.FromExternal(netUrl); err != nil {
					errs = append(errs, fmt.Errorf("invalid private visibility config network URL %q: %w", netUrl, err))
				} else {
					zone.PrivateVisibilityConfig.Networks[i].NetworkUrl = prefix + id.String()
				}
			}
		}
	}
	return errors.Join(errs...)
}
