// Copyright 2025 Google LLC
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

package compute

import (
	"context"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResolveComputeForwardingRuleRefs(ctx context.Context, kube client.Reader, obj *krm.ComputeForwardingRule) error {
	// TODO: resolve BackendServiceRef
	// TODO: resolve IpAddress.AddressRef
	// TODO: resolve NetworkRef
	// TODO: resolve SubnetworkRef
	// TODO: resolve Target.ServiceAttachmentRef
	// TODO: resolve Target.TargetGRPCProxyRef
	// TODO: resolve Target.TargetHTTPProxyRef
	// TODO: resolve Target.TargetHTTPSProxyRef
	// TODO: resolve Target.TargetSSLProxyRef
	// TODO: resolve Target.TargetTCPProxyRef
	// TODO: resolve Target.TargetVPNGatewayRef
	return nil
}
