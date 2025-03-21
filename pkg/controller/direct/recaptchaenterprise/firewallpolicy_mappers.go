package recaptchaenterprise

import (
	pb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/recaptchaenterprise/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ReCAPTCHAEnterpriseFirewallPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FirewallPolicy) *krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState {
	if in == nil {
		return nil
	}
	return &krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState{}
}

func ReCAPTCHAEnterpriseFirewallPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReCAPTCHAEnterpriseFirewallPolicyObservedState) *pb.FirewallPolicy {
	if in == nil {
		return nil
	}
	return &pb.FirewallPolicy{}
}
