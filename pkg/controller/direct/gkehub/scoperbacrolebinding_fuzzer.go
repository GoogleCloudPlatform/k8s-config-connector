package gkehub

import (
	gkehubv1 "google.golang.org/api/gkehub/v1"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer_NoProto(fuzzGKEHubScopeRBACRoleBinding())
}

func fuzzGKEHubScopeRBACRoleBinding() fuzztesting.KRMFuzzer_NoProto {
	f := fuzztesting.NewKRMTypedFuzzer_NoProto(&gkehubv1.RBACRoleBinding{},
		func(ctx *direct.MapContext, in *gkehubv1.RBACRoleBinding) *krm.GKEHubScopeRBACRoleBindingSpec {
			return GKEHubScopeRBACRoleBindingSpec_FromAPI(ctx, in, &krm.GKEHubScopeRBACRoleBindingIdentity{})
		},
		GKEHubScopeRBACRoleBindingSpec_ToAPI,
		func(ctx *direct.MapContext, in *gkehubv1.RBACRoleBinding) *krm.GKEHubScopeRBACRoleBindingStatus {
			return GKEHubScopeRBACRoleBindingStatus_FromAPI(ctx, in)
		},
		GKEHubScopeRBACRoleBindingStatus_ToAPI,
	)

	f.SpecField(".Role")
	f.SpecField(".User")
	f.SpecField(".Group")
	f.SpecField(".Labels")

	f.StatusField(".CreateTime")
	f.StatusField(".UpdateTime")
	f.StatusField(".DeleteTime")
	f.StatusField(".Uid")
	f.StatusField(".State")

	f.Unimplemented_NotYetTriaged(".Name")
	f.Unimplemented_NotYetTriaged(".ForceSendFields")
	f.Unimplemented_NotYetTriaged(".NullFields")
	f.Unimplemented_NotYetTriaged(".ServerResponse")

	f.Unimplemented_NotYetTriaged(".Role.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".Role.NullFields")
	f.Unimplemented_NotYetTriaged(".State.ForceSendFields")
	f.Unimplemented_NotYetTriaged(".State.NullFields")

	return f
}
