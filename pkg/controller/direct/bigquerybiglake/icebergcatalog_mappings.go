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

package bigquerybiglake

import (
	biglakepb "cloud.google.com/go/biglake/apiv1/biglakepb"
	krmbigquerybiglakev1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerybiglake/v1alpha1"
	krmsecretmanagerv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secretmanager/v1beta1"
	krmservicedirectoryv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/servicedirectory/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func IcebergCatalog_FederatedCatalogOptions_v1alpha1_FromProto(mapCtx *direct.MapContext, in *biglakepb.IcebergCatalog_FederatedCatalogOptions) *krmbigquerybiglakev1alpha1.IcebergCatalog_FederatedCatalogOptions {
	if in == nil {
		return nil
	}
	out := &krmbigquerybiglakev1alpha1.IcebergCatalog_FederatedCatalogOptions{}
	out.UnityCatalogInfo = IcebergCatalog_FederatedCatalogOptions_UnityCatalogInfo_v1alpha1_FromProto(mapCtx, in.GetUnityCatalogInfo())
	out.GlueCatalogInfo = IcebergCatalog_FederatedCatalogOptions_GlueCatalogInfo_v1alpha1_FromProto(mapCtx, in.GetGlueCatalogInfo())
	out.SnowflakeCatalogInfo = IcebergCatalog_FederatedCatalogOptions_SnowflakeCatalogInfo_v1alpha1_FromProto(mapCtx, in.GetSnowflakeCatalogInfo())
	if in.GetSecretName() != "" {
		out.SecretRef = &krmsecretmanagerv1beta1.SecretManagerSecretRef{External: in.GetSecretName()}
	}
	if in.GetServiceDirectoryName() != "" {
		out.ServiceDirectoryRef = &krmservicedirectoryv1beta1.ServiceDirectoryServiceRef{External: in.GetServiceDirectoryName()}
	}
	out.RefreshOptions = IcebergCatalog_FederatedCatalogOptions_RefreshOptions_v1alpha1_FromProto(mapCtx, in.GetRefreshOptions())
	return out
}

func IcebergCatalog_FederatedCatalogOptions_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmbigquerybiglakev1alpha1.IcebergCatalog_FederatedCatalogOptions) *biglakepb.IcebergCatalog_FederatedCatalogOptions {
	if in == nil {
		return nil
	}
	out := &biglakepb.IcebergCatalog_FederatedCatalogOptions{}
	if oneof := IcebergCatalog_FederatedCatalogOptions_UnityCatalogInfo_v1alpha1_ToProto(mapCtx, in.UnityCatalogInfo); oneof != nil {
		out.RemoteCatalogInfo = &biglakepb.IcebergCatalog_FederatedCatalogOptions_UnityCatalogInfo_{UnityCatalogInfo: oneof}
	}
	if oneof := IcebergCatalog_FederatedCatalogOptions_GlueCatalogInfo_v1alpha1_ToProto(mapCtx, in.GlueCatalogInfo); oneof != nil {
		out.RemoteCatalogInfo = &biglakepb.IcebergCatalog_FederatedCatalogOptions_GlueCatalogInfo_{GlueCatalogInfo: oneof}
	}
	if oneof := IcebergCatalog_FederatedCatalogOptions_SnowflakeCatalogInfo_v1alpha1_ToProto(mapCtx, in.SnowflakeCatalogInfo); oneof != nil {
		out.RemoteCatalogInfo = &biglakepb.IcebergCatalog_FederatedCatalogOptions_SnowflakeCatalogInfo_{SnowflakeCatalogInfo: oneof}
	}
	if in.SecretRef != nil && in.SecretRef.External != "" {
		out.SecretName = &in.SecretRef.External
	}
	if in.ServiceDirectoryRef != nil && in.ServiceDirectoryRef.External != "" {
		out.ServiceDirectoryName = &in.ServiceDirectoryRef.External
	}
	out.RefreshOptions = IcebergCatalog_FederatedCatalogOptions_RefreshOptions_v1alpha1_ToProto(mapCtx, in.RefreshOptions)
	return out
}
