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

package leaser

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/lease/leasable"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"

	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	leaseHolderKey     = "cnrm-lease-holder-id"
	leaseExpirationKey = "cnrm-lease-expiration"
)

var (
	zeroUnixTime = time.Unix(0, 0)
	leaseKeys    = []string{leaseHolderKey, leaseExpirationKey}
)

// Leaser locks for a period of time by GCP resources by applying 'owner' and 'expiration' labels to the resource. After they are applied,
// Leaser will only allow a caller to Obtain the resource if one of the following is true:
// 1. The owner id is the same as what is saved in the labels
// 2. The expiration time is before the current time
//
// No protections are made for race conditions: the last writer will win.
type Leaser struct {
	tfProvider *tfschema.Provider
	kubeClient client.Client
	smLoader   *servicemappingloader.ServiceMappingLoader
}

func NewLeaser(tfProvider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader, kubeClient client.Client) *Leaser {
	return &Leaser{
		tfProvider: tfProvider,
		kubeClient: kubeClient,
		smLoader:   smLoader,
	}
}

// Obtains a lease, for 'ownerID', on the given unstructured for 'duration' time. If the owner already
// has a lease on the unstructured the lease is renewed to 'duration' time if the 'renewalMin' is greater than the remaining
// time on the lease.
//
// To 'always' renew the lease, pass in a 'renewalMin' that is equal to the duration.
func (l *Leaser) Obtain(ctx context.Context, u *unstructured.Unstructured, ownerID string, duration time.Duration, renewalMin time.Duration) error {
	if err := l.validateUnstructuredSupportsLocking(u); err != nil {
		return err
	}
	if renewalMin > duration {
		return fmt.Errorf("invalid argument, renewalMin '%v' is greater than duration '%v'", renewalMin, duration)
	}
	resource, liveState, err := l.getResourceAndLiveState(ctx, u)
	if err != nil {
		return err
	}
	ok, err := l.softObtain(resource.Labels, ownerID, duration, renewalMin)
	if err != nil {
		return err
	}
	if !ok {
		// the lease was not obtained because it is still held by this owner
		return nil
	}
	config, _, err := krmtotf.KRMResourceToTFResourceConfig(resource, l.kubeClient, l.smLoader)
	if err != nil {
		return fmt.Errorf("error expanding resource configuration: %w", err)
	}
	diff, err := resource.TFResource.Diff(ctx, liveState, config, l.tfProvider.Meta())
	if err != nil {
		return fmt.Errorf("error calculating diff: %w", err)
	}
	if diff.Empty() {
		return nil
	}
	_, diagnostics := resource.TFResource.Apply(ctx, liveState, diff, l.tfProvider.Meta())

	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return fmt.Errorf("error applying resource change: %w", err)
	}
	return nil
}

// Soft obtain obtains a lease for 'ownerID' on the given resource for 'duration' time. See the comment on Obtain(...) for more.
//
// It does not write the results to GCP so the caller must apply the changes to GCP if persistence is desired
func (l *Leaser) SoftObtain(resource *k8s.Resource, liveLabels map[string]string, ownerID string, duration time.Duration, renewalMin time.Duration) error {
	if _, err := l.softObtain(liveLabels, ownerID, duration, renewalMin); err != nil {
		return err
	}
	if resource.Labels == nil {
		resource.Labels = liveLabels
	} else {
		if val, ok := liveLabels[leaseHolderKey]; ok {
			resource.Labels[leaseHolderKey] = val
		}
		if val, ok := liveLabels[leaseExpirationKey]; ok {
			resource.Labels[leaseExpirationKey] = val
		}
	}
	return nil
}

// checks to see if the lease is obtainable, if not an error is returned. If it is obtainable then the lease is renewed if necessary. If it is unnecessary
// to renew the lease as the time period is still within the renewalMin window then 'false' is returned for the 'ok' parameter
func (l *Leaser) softObtain(labels map[string]string, ownerID string, duration time.Duration, renewalMin time.Duration) (ok bool, err error) {
	leaseHolder, expirationTime := getLeaseHolderAndExpirationTime(labels)
	if !canObtainLease(ownerID, leaseHolder, expirationTime) {
		return false, fmt.Errorf("resource is under lease by '%v' for an additional %v second(s)", leaseHolder, expirationTime.Sub(time.Now()))
	}
	if !shouldRenewOrObtainLease(renewalMin, expirationTime) {
		return false, nil
	}
	setLeaseHolder(labels, ownerID, duration)
	return true, nil
}

func (l *Leaser) Release(ctx context.Context, u *unstructured.Unstructured, ownerID string) error {
	if err := l.validateUnstructuredSupportsLocking(u); err != nil {
		return err
	}
	resource, liveState, err := l.getResourceAndLiveState(ctx, u)
	if err != nil {
		return err
	}
	leaseHolder, expirationTime := getLeaseHolderAndExpirationTime(resource.Labels)
	if leaseHolder == "" {
		return fmt.Errorf("resource is not under management by '%v' or any other owner", ownerID)
	}
	now := time.Now()
	if leaseHolder != ownerID {
		return fmt.Errorf("resource is under lease by '%v' for an additional %v second(s)", leaseHolder, expirationTime.Sub(now))
	}
	if expirationTime.Before(now) {
		return fmt.Errorf("unable to release lease: expired %v second(s) ago", now.Sub(expirationTime))
	}
	delete(resource.Labels, leaseHolderKey)
	config, _, err := krmtotf.KRMResourceToTFResourceConfig(resource, l.kubeClient, l.smLoader)
	if err != nil {
		return fmt.Errorf("error expanding resource configuration: %w", err)
	}
	diff, err := resource.TFResource.Diff(ctx, liveState, config, l.tfProvider.Meta())
	if err != nil {
		return fmt.Errorf("error calculating diff: %w", err)
	}
	_, diagnostics := resource.TFResource.Apply(ctx, liveState, diff, l.tfProvider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return fmt.Errorf("error applying resource change: %w", err)
	}
	return nil
}

func (l *Leaser) GetOwnerAndExpirationTime(ctx context.Context, u *unstructured.Unstructured) (string, time.Time, error) {
	if err := l.validateUnstructuredSupportsLocking(u); err != nil {
		return "", zeroUnixTime, err
	}
	resource, _, err := l.getResourceAndLiveState(ctx, u)
	if err != nil {
		return "", zeroUnixTime, err
	}
	leaseHolder, expirationTime := getLeaseHolderAndExpirationTime(resource.Labels)
	return leaseHolder, expirationTime, nil
}

func (l *Leaser) getResourceAndLiveState(ctx context.Context, u *unstructured.Unstructured) (*krmtotf.Resource,
	*terraform.InstanceState, error) {
	sm, err := l.smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return nil, nil, fmt.Errorf("error getting service mapping for gvk '%v': %w", u.GroupVersionKind(), err)
	}
	resource, err := krmtotf.NewResource(u, sm, l.tfProvider)
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing resource %s: %w", u.GetName(), err)
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, l.tfProvider, l.kubeClient, l.smLoader)
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching live state: %w", err)
	}
	if liveState.Empty() {
		return nil, nil, fmt.Errorf("resource '%v' of type '%v': not found", u.GetName(), u.GroupVersionKind())
	}
	resource.Labels = krmtotf.GetLabelsFromState(resource, liveState)
	resource.Annotations = krmtotf.GetAnnotationsFromState(resource, liveState)
	resource.Spec, resource.Status = krmtotf.ResolveSpecAndStatusWithResourceID(resource, liveState)

	return resource, liveState, nil
}

// Get the keys that are set on a resource to obtain a lease
func GetLabelKeys() []string {
	return leaseKeys
}

func (l *Leaser) UnstructuredSupportsLeasing(u *unstructured.Unstructured) (ok bool, err error) {
	sm, err := l.smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return false, fmt.Errorf("error getting service mapping: %w", err)
	}
	rc, err := servicemappingloader.GetResourceConfig(sm, u)
	if err != nil {
		return false, fmt.Errorf("error getting resource config: %w", err)
	}
	return leasable.ResourceConfigSupportsLeasing(rc, l.tfProvider.ResourcesMap)
}

func (l *Leaser) validateUnstructuredSupportsLocking(u *unstructured.Unstructured) error {
	ok, err := l.UnstructuredSupportsLeasing(u)
	if err != nil {
		return fmt.Errorf("error determining if gvk '%v' supports locking: %w", u.GroupVersionKind(), err)
	}
	if !ok {
		return fmt.Errorf("gvk '%v' does not support locking", u.GroupVersionKind())
	}
	return nil
}

func shouldRenewOrObtainLease(minDuration time.Duration, expirationTime time.Time) bool {
	return time.Until(expirationTime) < minDuration
}

func canObtainLease(ownerID string, curLeaseHolder string, expirationTime time.Time) bool {
	if curLeaseHolder == "" || curLeaseHolder == ownerID {
		return true
	}
	return time.Now().After(expirationTime)
}

func setLeaseHolder(labels map[string]string, ownerID string, duration time.Duration) {
	labels[leaseHolderKey] = ownerID
	labels[leaseExpirationKey] = strconv.FormatInt(time.Now().Add(duration).Unix(), 10)
}

func getLeaseHolderAndExpirationTime(labels map[string]string) (string, time.Time) {
	leaseHolder, ok := labels[leaseHolderKey]
	if !ok {
		return "", zeroUnixTime
	}
	leaseExpirationString, ok := labels[leaseExpirationKey]
	if !ok {
		return leaseHolder, zeroUnixTime
	}
	unixTime, err := strconv.ParseInt(leaseExpirationString, 10, 64)
	if err != nil {
		return leaseHolder, zeroUnixTime
	}
	return leaseHolder, time.Unix(unixTime, 0)
}
