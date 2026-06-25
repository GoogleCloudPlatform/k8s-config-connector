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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dns/v1beta1"
	apirefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	api "google.golang.org/api/dns/v1"
)

func DNSRecordSetSpec_FromAPI(mapCtx *direct.MapContext, in *api.ResourceRecordSet) *krm.DNSRecordSetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DNSRecordSetSpec{}
	out.Name = in.Name
	out.RoutingPolicy = RecordsetRoutingPolicy_FromAPI(mapCtx, in.RoutingPolicy)
	out.Rrdatas = in.Rrdatas
	out.Ttl = direct.LazyPtr(in.Ttl)
	out.Type = in.Type
	return out
}

func DNSRecordSetSpec_ToAPI(mapCtx *direct.MapContext, in *krm.DNSRecordSetSpec) *api.ResourceRecordSet {
	if in == nil {
		return nil
	}
	out := &api.ResourceRecordSet{}
	out.Name = in.Name
	out.RoutingPolicy = RecordsetRoutingPolicy_ToAPI(mapCtx, in.RoutingPolicy)
	out.Rrdatas = in.Rrdatas
	for _, ref := range in.RrdatasRefs {
		if ref.External != nil && *ref.External != "" {
			out.Rrdatas = append(out.Rrdatas, *ref.External)
		}
	}
	out.Ttl = direct.ValueOf(in.Ttl)
	out.Type = in.Type
	return out
}

func DNSRecordSetStatus_FromAPI(mapCtx *direct.MapContext, in *api.ResourceRecordSet) *krm.DNSRecordSetStatus {
	if in == nil {
		return nil
	}
	out := &krm.DNSRecordSetStatus{}
	return out
}

func DNSRecordSetStatus_ToAPI(mapCtx *direct.MapContext, in *krm.DNSRecordSetStatus) *api.ResourceRecordSet {
	if in == nil {
		return nil
	}
	out := &api.ResourceRecordSet{}
	return out
}

func RecordsetRoutingPolicy_FromAPI(mapCtx *direct.MapContext, in *api.RRSetRoutingPolicy) *krm.RecordsetRoutingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.RecordsetRoutingPolicy{}
	if in.Geo != nil {
		out.EnableGeoFencing = direct.LazyPtr(in.Geo.EnableFencing)
		for _, item := range in.Geo.Items {
			out.Geo = append(out.Geo, *RecordsetGeo_FromAPI(mapCtx, item))
		}
	}
	out.PrimaryBackup = RecordsetPrimaryBackup_FromAPI(mapCtx, in.PrimaryBackup)
	if in.Wrr != nil {
		for _, item := range in.Wrr.Items {
			out.Wrr = append(out.Wrr, *RecordsetWrr_FromAPI(mapCtx, item))
		}
	}
	return out
}

func RecordsetRoutingPolicy_ToAPI(mapCtx *direct.MapContext, in *krm.RecordsetRoutingPolicy) *api.RRSetRoutingPolicy {
	if in == nil {
		return nil
	}
	out := &api.RRSetRoutingPolicy{}
	if len(in.Geo) > 0 || in.EnableGeoFencing != nil {
		geo := &api.RRSetRoutingPolicyGeoPolicy{}
		geo.EnableFencing = direct.ValueOf(in.EnableGeoFencing)
		for i := range in.Geo {
			if apiItem := RecordsetGeo_ToAPI(mapCtx, &in.Geo[i]); apiItem != nil {
				geo.Items = append(geo.Items, apiItem)
			}
		}
		out.Geo = geo
	}
	out.PrimaryBackup = RecordsetPrimaryBackup_ToAPI(mapCtx, in.PrimaryBackup)
	if len(in.Wrr) > 0 {
		wrr := &api.RRSetRoutingPolicyWrrPolicy{}
		for i := range in.Wrr {
			if apiItem := RecordsetWrr_ToAPI(mapCtx, &in.Wrr[i]); apiItem != nil {
				wrr.Items = append(wrr.Items, apiItem)
			}
		}
		out.Wrr = wrr
	}
	return out
}

func RecordsetGeo_FromAPI(mapCtx *direct.MapContext, in *api.RRSetRoutingPolicyGeoPolicyGeoPolicyItem) *krm.RecordsetGeo {
	if in == nil {
		return nil
	}
	out := &krm.RecordsetGeo{}
	out.Location = in.Location
	out.HealthCheckedTargets = RecordsetHealthCheckedTargets_FromAPI(mapCtx, in.HealthCheckedTargets)
	for _, r := range in.Rrdatas {
		out.RrdatasRefs = append(out.RrdatasRefs, krm.RecordsetRrdatasRefs{
			External: direct.LazyPtr(r),
		})
	}
	return out
}

func RecordsetGeo_ToAPI(mapCtx *direct.MapContext, in *krm.RecordsetGeo) *api.RRSetRoutingPolicyGeoPolicyGeoPolicyItem {
	if in == nil {
		return nil
	}
	out := &api.RRSetRoutingPolicyGeoPolicyGeoPolicyItem{}
	out.Location = in.Location
	out.HealthCheckedTargets = RecordsetHealthCheckedTargets_ToAPI(mapCtx, in.HealthCheckedTargets)
	for _, ref := range in.RrdatasRefs {
		if ref.External != nil && *ref.External != "" {
			out.Rrdatas = append(out.Rrdatas, *ref.External)
		}
	}
	return out
}

func RecordsetBackupGeo_FromAPI(mapCtx *direct.MapContext, in *api.RRSetRoutingPolicyGeoPolicyGeoPolicyItem) *krm.RecordsetBackupGeo {
	if in == nil {
		return nil
	}
	out := &krm.RecordsetBackupGeo{}
	out.Location = in.Location
	out.HealthCheckedTargets = RecordsetHealthCheckedTargets_FromAPI(mapCtx, in.HealthCheckedTargets)
	for _, r := range in.Rrdatas {
		out.RrdatasRefs = append(out.RrdatasRefs, krm.RecordsetRrdatasRefs{
			External: direct.LazyPtr(r),
		})
	}
	return out
}

func RecordsetBackupGeo_ToAPI(mapCtx *direct.MapContext, in *krm.RecordsetBackupGeo) *api.RRSetRoutingPolicyGeoPolicyGeoPolicyItem {
	if in == nil {
		return nil
	}
	out := &api.RRSetRoutingPolicyGeoPolicyGeoPolicyItem{}
	out.Location = in.Location
	out.HealthCheckedTargets = RecordsetHealthCheckedTargets_ToAPI(mapCtx, in.HealthCheckedTargets)
	for _, ref := range in.RrdatasRefs {
		if ref.External != nil && *ref.External != "" {
			out.Rrdatas = append(out.Rrdatas, *ref.External)
		}
	}
	return out
}

func RecordsetWrr_FromAPI(mapCtx *direct.MapContext, in *api.RRSetRoutingPolicyWrrPolicyWrrPolicyItem) *krm.RecordsetWrr {
	if in == nil {
		return nil
	}
	out := &krm.RecordsetWrr{}
	out.Weight = in.Weight
	out.HealthCheckedTargets = RecordsetHealthCheckedTargets_FromAPI(mapCtx, in.HealthCheckedTargets)
	for _, r := range in.Rrdatas {
		out.RrdatasRefs = append(out.RrdatasRefs, krm.RecordsetRrdatasRefs{
			External: direct.LazyPtr(r),
		})
	}
	return out
}

func RecordsetWrr_ToAPI(mapCtx *direct.MapContext, in *krm.RecordsetWrr) *api.RRSetRoutingPolicyWrrPolicyWrrPolicyItem {
	if in == nil {
		return nil
	}
	out := &api.RRSetRoutingPolicyWrrPolicyWrrPolicyItem{}
	out.Weight = in.Weight
	out.HealthCheckedTargets = RecordsetHealthCheckedTargets_ToAPI(mapCtx, in.HealthCheckedTargets)
	for _, ref := range in.RrdatasRefs {
		if ref.External != nil && *ref.External != "" {
			out.Rrdatas = append(out.Rrdatas, *ref.External)
		}
	}
	return out
}

func RecordsetPrimaryBackup_FromAPI(mapCtx *direct.MapContext, in *api.RRSetRoutingPolicyPrimaryBackupPolicy) *krm.RecordsetPrimaryBackup {
	if in == nil {
		return nil
	}
	out := &krm.RecordsetPrimaryBackup{}
	if in.BackupGeoTargets != nil {
		out.EnableGeoFencingForBackups = direct.LazyPtr(in.BackupGeoTargets.EnableFencing)
		for _, item := range in.BackupGeoTargets.Items {
			out.BackupGeo = append(out.BackupGeo, *RecordsetBackupGeo_FromAPI(mapCtx, item))
		}
	}
	if in.PrimaryTargets != nil {
		out.Primary.InternalLoadBalancers = RecordsetHealthCheckedTargets_FromAPI(mapCtx, in.PrimaryTargets).InternalLoadBalancers
	}
	out.TrickleRatio = direct.LazyPtr(in.TrickleTraffic)
	return out
}

func RecordsetPrimaryBackup_ToAPI(mapCtx *direct.MapContext, in *krm.RecordsetPrimaryBackup) *api.RRSetRoutingPolicyPrimaryBackupPolicy {
	if in == nil {
		return nil
	}
	out := &api.RRSetRoutingPolicyPrimaryBackupPolicy{}
	if len(in.BackupGeo) > 0 || in.EnableGeoFencingForBackups != nil {
		backupGeo := &api.RRSetRoutingPolicyGeoPolicy{}
		backupGeo.EnableFencing = direct.ValueOf(in.EnableGeoFencingForBackups)
		for i := range in.BackupGeo {
			if apiItem := RecordsetBackupGeo_ToAPI(mapCtx, &in.BackupGeo[i]); apiItem != nil {
				backupGeo.Items = append(backupGeo.Items, apiItem)
			}
		}
		out.BackupGeoTargets = backupGeo
	}
	if len(in.Primary.InternalLoadBalancers) > 0 {
		hct := &krm.RecordsetHealthCheckedTargets{
			InternalLoadBalancers: in.Primary.InternalLoadBalancers,
		}
		out.PrimaryTargets = RecordsetHealthCheckedTargets_ToAPI(mapCtx, hct)
	}
	out.TrickleTraffic = direct.ValueOf(in.TrickleRatio)
	return out
}

func RecordsetHealthCheckedTargets_FromAPI(mapCtx *direct.MapContext, in *api.RRSetRoutingPolicyHealthCheckTargets) *krm.RecordsetHealthCheckedTargets {
	if in == nil {
		return nil
	}
	out := &krm.RecordsetHealthCheckedTargets{}
	for _, item := range in.InternalLoadBalancers {
		out.InternalLoadBalancers = append(out.InternalLoadBalancers, *RecordsetInternalLoadBalancers_FromAPI(mapCtx, item))
	}
	return out
}

func RecordsetHealthCheckedTargets_ToAPI(mapCtx *direct.MapContext, in *krm.RecordsetHealthCheckedTargets) *api.RRSetRoutingPolicyHealthCheckTargets {
	if in == nil {
		return nil
	}
	out := &api.RRSetRoutingPolicyHealthCheckTargets{}
	for i := range in.InternalLoadBalancers {
		if apiItem := RecordsetInternalLoadBalancers_ToAPI(mapCtx, &in.InternalLoadBalancers[i]); apiItem != nil {
			out.InternalLoadBalancers = append(out.InternalLoadBalancers, apiItem)
		}
	}
	return out
}

func RecordsetInternalLoadBalancers_FromAPI(mapCtx *direct.MapContext, in *api.RRSetRoutingPolicyLoadBalancerTarget) *krm.RecordsetInternalLoadBalancers {
	if in == nil {
		return nil
	}
	out := &krm.RecordsetInternalLoadBalancers{}
	out.IpAddressRef = computev1beta1.ComputeAddressRef{External: in.IpAddress}
	out.IpProtocol = in.IpProtocol
	out.LoadBalancerType = in.LoadBalancerType
	out.NetworkRef = computev1beta1.ComputeNetworkRef{External: in.NetworkUrl}
	out.Port = in.Port
	out.ProjectRef = apirefs.ProjectRef{External: in.Project}
	if in.Region != "" {
		out.RegionRef = &krm.RegionRef{External: in.Region}
	}
	return out
}

func RecordsetInternalLoadBalancers_ToAPI(mapCtx *direct.MapContext, in *krm.RecordsetInternalLoadBalancers) *api.RRSetRoutingPolicyLoadBalancerTarget {
	if in == nil {
		return nil
	}
	out := &api.RRSetRoutingPolicyLoadBalancerTarget{}
	out.IpAddress = in.IpAddressRef.External
	out.IpProtocol = in.IpProtocol
	out.LoadBalancerType = in.LoadBalancerType
	out.NetworkUrl = in.NetworkRef.External
	out.Port = in.Port
	out.Project = in.ProjectRef.External
	if in.RegionRef != nil {
		out.Region = in.RegionRef.External
	}
	return out
}
