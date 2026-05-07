package gkehub

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	gkehubv1 "google.golang.org/api/gkehub/v1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1alpha1"
)

func GKEHubScopeRBACRoleBindingSpec_ToAPI(mapCtx *direct.MapContext, in *krm.GKEHubScopeRBACRoleBindingSpec) (*gkehubv1.RBACRoleBinding) {
	if in == nil {
		return nil
	}
	out := &gkehubv1.RBACRoleBinding{}
	if in.Role != nil {
		out.Role = &gkehubv1.Role{
			CustomRole:     direct.ValueOf(in.Role.CustomRole),
			PredefinedRole: direct.ValueOf(in.Role.PredefinedRole),
		}
	}
	out.User = direct.ValueOf(in.User)
	out.Group = direct.ValueOf(in.Group)
	out.Labels = in.Labels
	return out
}

func GKEHubScopeRBACRoleBindingStatus_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.RBACRoleBinding) (*krm.GKEHubScopeRBACRoleBindingStatus) {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubScopeRBACRoleBindingStatus{}
	out.ObservedState = &krm.GKEHubScopeRBACRoleBindingObservedState{
		CreateTime: direct.LazyPtr(in.CreateTime),
		UpdateTime: direct.LazyPtr(in.UpdateTime),
		DeleteTime: direct.LazyPtr(in.DeleteTime),
		Uid:        direct.LazyPtr(in.Uid),
	}
	if in.State != nil {
		out.ObservedState.State = &krm.RBACRoleBindingStateStatus{
			Code: direct.LazyPtr(in.State.Code),
		}
	}
	return out
}

func GKEHubScopeRBACRoleBindingSpec_FromAPI(mapCtx *direct.MapContext, in *gkehubv1.RBACRoleBinding, id *krm.GKEHubScopeRBACRoleBindingIdentity) *krm.GKEHubScopeRBACRoleBindingSpec {
	if in == nil {
		return nil
	}
	out := &krm.GKEHubScopeRBACRoleBindingSpec{}
	out.ScopeRef = &krm.GKEHubScopeRef{
		External: "projects/" + id.ProjectID + "/locations/" + id.Location + "/scopes/" + id.ScopeID,
	}
	out.RBACRoleBindingID = direct.LazyPtr(id.RBACRoleBindingID)
	if in.Role != nil {
		out.Role = &krm.RBACRoleBindingRole{
			CustomRole:     direct.LazyPtr(in.Role.CustomRole),
			PredefinedRole: direct.LazyPtr(in.Role.PredefinedRole),
		}
	}
	out.User = direct.LazyPtr(in.User)
	out.Group = direct.LazyPtr(in.Group)
	out.Labels = in.Labels
	return out
}

func GKEHubScopeRBACRoleBindingStatus_ToAPI(mapCtx *direct.MapContext, in *krm.GKEHubScopeRBACRoleBindingStatus) *gkehubv1.RBACRoleBinding {
	if in == nil || in.ObservedState == nil {
		return nil
	}
	out := &gkehubv1.RBACRoleBinding{}
	out.CreateTime = direct.ValueOf(in.ObservedState.CreateTime)
	out.UpdateTime = direct.ValueOf(in.ObservedState.UpdateTime)
	out.DeleteTime = direct.ValueOf(in.ObservedState.DeleteTime)
	out.Uid = direct.ValueOf(in.ObservedState.Uid)
	if in.ObservedState.State != nil {
		out.State = &gkehubv1.RBACRoleBindingLifecycleState{
			Code: direct.ValueOf(in.ObservedState.State.Code),
		}
	}
	return out
}
