package iam

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/iam/v1beta1"
	api "google.golang.org/api/iam/v1"
)

func ServiceAccountSpec_FromProto(ctx *MapContext, in *api.ServiceAccount) *krm.IAMServiceAccountSpec {
	if in == nil {
		return nil
	}
	out := &krm.IAMServiceAccountSpec{}
	out.DisplayName = LazyPtr(in.DisplayName)
	out.Description = LazyPtr(in.Description)
	out.Disabled = LazyPtr(in.Disabled)
	return out
}

func ServiceAccountStatus_FromProto(ctx *MapContext, in *api.ServiceAccount) *krm.IAMServiceAccountStatus {
	if in == nil {
		return nil
	}
	out := &krm.IAMServiceAccountStatus{}
	out.Email = LazyPtr(in.Email)
	out.UniqueId = LazyPtr(in.UniqueId)
	out.Member = LazyPtr("serviceAccount:" + in.Email)
	out.Name = LazyPtr(in.Name)
	return out
}

func ServiceAccountSpec_ToProto(ctx *MapContext, in *krm.IAMServiceAccountSpec) *api.ServiceAccount {
	if in == nil {
		return nil
	}
	out := &api.ServiceAccount{}
	out.DisplayName = ValueOf(in.DisplayName)
	out.Description = ValueOf(in.Description)
	out.Disabled = ValueOf(in.Disabled)
	return out
}
