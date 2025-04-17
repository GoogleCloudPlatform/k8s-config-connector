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

// +generated:mapper
// krm.group: eventarc.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.eventarc.v1

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	connectorv1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connector/v1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EventarcChannelObservedState_PubsubTopic_ToProto(mapCtx *direct.MapContext, in *string) *pb.Channel_PubsubTopic {
	if in == nil {
		return nil
	}
	return &pb.Channel_PubsubTopic{
		PubsubTopic: *in,
	}
}

func EventarcChannelSpec_FromProto(mapCtx *direct.MapContext, in *pb.Channel) *krm.EventarcChannelSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcChannelSpec{}
	// MISSING: Name
	// Provider is a ProviderRef struct in the KRM type, not a string
	if provider := in.GetProvider(); provider != "" {
		out.Provider = &connectorv1.ProviderRef{
			External: provider,
		}
	}
	if in.GetCryptoKeyName() != "" {
		out.KmsKeyRef = &v1beta1.KMSCryptoKeyRef{
			External: in.GetCryptoKeyName(),
		}
	}
	return out
}

func EventarcChannelSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcChannelSpec) *pb.Channel {
	if in == nil {
		return nil
	}
	out := &pb.Channel{}
	// MISSING: Name
	// Extract the string value from the ProviderRef
	if in.Provider != nil {
		out.Provider = in.Provider.External
	}
	if in.KmsKeyRef != nil {
		out.CryptoKeyName = in.KmsKeyRef.External
	}
	return out
}
