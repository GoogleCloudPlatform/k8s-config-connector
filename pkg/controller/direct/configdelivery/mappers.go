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

package configdelivery

import (
	pb "cloud.google.com/go/configdelivery/apiv1/configdeliverypb"
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/configdelivery/v1alpha1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func FleetPackage_CloudBuildRepository_FromProto(mapCtx *direct.MapContext, in *pb.FleetPackage_CloudBuildRepository) *krm.FleetPackage_CloudBuildRepository {
	if in == nil {
		return nil
	}
	out := &krm.FleetPackage_CloudBuildRepository{}
	out.VariantsPattern = direct.LazyPtr(in.GetVariantsPattern())
	if in.GetName() != "" {
		out.RepositoryRef = &cloudbuildv1beta1.RepositoryRef{External: in.GetName()}
	}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Tag = direct.LazyPtr(in.GetTag())
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &refsv1beta1.IAMServiceAccountRef{External: in.GetServiceAccount()}
	}
	return out
}

func FleetPackage_CloudBuildRepository_ToProto(mapCtx *direct.MapContext, in *krm.FleetPackage_CloudBuildRepository) *pb.FleetPackage_CloudBuildRepository {
	if in == nil {
		return nil
	}
	out := &pb.FleetPackage_CloudBuildRepository{}
	if oneof := FleetPackage_CloudBuildRepository_VariantsPattern_ToProto(mapCtx, in.VariantsPattern); oneof != nil {
		out.Variants = oneof
	}
	if in.RepositoryRef != nil {
		out.Name = in.RepositoryRef.External
	}
	out.Path = direct.ValueOf(in.Path)
	out.Tag = direct.ValueOf(in.Tag)
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	return out
}

func FleetPackage_ResourceBundleTag_FromProto(mapCtx *direct.MapContext, in *pb.FleetPackage_ResourceBundleTag) *krm.FleetPackage_ResourceBundleTag {
	if in == nil {
		return nil
	}
	out := &krm.FleetPackage_ResourceBundleTag{}
	if in.GetName() != "" {
		out.ResourceBundleRef = &krm.ConfigDeliveryResourceBundleRef{External: in.GetName()}
	}
	out.Tag = direct.LazyPtr(in.GetTag())
	return out
}

func FleetPackage_ResourceBundleTag_ToProto(mapCtx *direct.MapContext, in *krm.FleetPackage_ResourceBundleTag) *pb.FleetPackage_ResourceBundleTag {
	if in == nil {
		return nil
	}
	out := &pb.FleetPackage_ResourceBundleTag{}
	if in.ResourceBundleRef != nil {
		out.Name = in.ResourceBundleRef.External
	}
	out.Tag = direct.ValueOf(in.Tag)
	return out
}
