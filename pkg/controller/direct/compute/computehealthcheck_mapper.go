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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ComputeHealthCheckSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HealthCheck) *krm.ComputeHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeHealthCheckSpec{}
	out.CheckIntervalSec = ConvertInt32ToInt(in.CheckIntervalSec)
	out.Description = in.Description
	out.GRPCHealthCheck = HealthCheckGRPCHealthCheck_v1beta1_FromProto(mapCtx, in.GrpcHealthCheck)
	out.HealthyThreshold = ConvertInt32ToInt(in.HealthyThreshold)
	out.HTTP2HealthCheck = HealthCheckHTTP2HealthCheck_v1beta1_FromProto(mapCtx, in.Http2HealthCheck)
	out.HTTPHealthCheck = HealthCheckHTTPHealthCheck_v1beta1_FromProto(mapCtx, in.HttpHealthCheck)
	out.HTTPSHealthCheck = HealthCheckHTTPSHealthCheck_v1beta1_FromProto(mapCtx, in.HttpsHealthCheck)
	out.LogConfig = HealthCheckLogConfig_v1beta1_FromProto(mapCtx, in.LogConfig)
	out.SSLHealthCheck = HealthCheckSSLHealthCheck_v1beta1_FromProto(mapCtx, in.SslHealthCheck)
	out.TCPHealthCheck = HealthCheckTCPHealthCheck_v1beta1_FromProto(mapCtx, in.TcpHealthCheck)
	out.TimeoutSec = ConvertInt32ToInt(in.TimeoutSec)
	out.UnhealthyThreshold = ConvertInt32ToInt(in.UnhealthyThreshold)
	return out
}

func ComputeHealthCheckSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeHealthCheckSpec) *pb.HealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HealthCheck{}
	out.CheckIntervalSec = ConvertIntToInt32(in.CheckIntervalSec)
	out.Description = in.Description
	out.GrpcHealthCheck = HealthCheckGRPCHealthCheck_v1beta1_ToProto(mapCtx, in.GRPCHealthCheck)
	out.HealthyThreshold = ConvertIntToInt32(in.HealthyThreshold)
	out.Http2HealthCheck = HealthCheckHTTP2HealthCheck_v1beta1_ToProto(mapCtx, in.HTTP2HealthCheck)
	out.HttpHealthCheck = HealthCheckHTTPHealthCheck_v1beta1_ToProto(mapCtx, in.HTTPHealthCheck)
	out.HttpsHealthCheck = HealthCheckHTTPSHealthCheck_v1beta1_ToProto(mapCtx, in.HTTPSHealthCheck)
	out.LogConfig = HealthCheckLogConfig_v1beta1_ToProto(mapCtx, in.LogConfig)
	out.SslHealthCheck = HealthCheckSSLHealthCheck_v1beta1_ToProto(mapCtx, in.SSLHealthCheck)
	out.TcpHealthCheck = HealthCheckTCPHealthCheck_v1beta1_ToProto(mapCtx, in.TCPHealthCheck)
	out.TimeoutSec = ConvertIntToInt32(in.TimeoutSec)
	out.UnhealthyThreshold = ConvertIntToInt32(in.UnhealthyThreshold)
	return out
}

func HealthCheckGRPCHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.GRPCHealthCheck) *krm.HealthCheckGRPCHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckGRPCHealthCheck{}
	out.GRPCServiceName = in.GrpcServiceName
	out.Port = ConvertInt32ToInt(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	return out
}

func HealthCheckGRPCHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckGRPCHealthCheck) *pb.GRPCHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.GRPCHealthCheck{}
	out.GrpcServiceName = in.GRPCServiceName
	out.Port = ConvertIntToInt32(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	return out
}

func HealthCheckHTTP2HealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HTTP2HealthCheck) *krm.HealthCheckHTTP2HealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckHTTP2HealthCheck{}
	out.Host = in.Host
	out.Port = ConvertInt32ToInt(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTP2HealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckHTTP2HealthCheck) *pb.HTTP2HealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HTTP2HealthCheck{}
	out.Host = in.Host
	out.Port = ConvertIntToInt32(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTPHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HTTPHealthCheck) *krm.HealthCheckHTTPHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckHTTPHealthCheck{}
	out.Host = in.Host
	out.Port = ConvertInt32ToInt(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTPHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckHTTPHealthCheck) *pb.HTTPHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HTTPHealthCheck{}
	out.Host = in.Host
	out.Port = ConvertIntToInt32(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTPSHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HTTPSHealthCheck) *krm.HealthCheckHTTPSHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckHTTPSHealthCheck{}
	out.Host = in.Host
	out.Port = ConvertInt32ToInt(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckHTTPSHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckHTTPSHealthCheck) *pb.HTTPSHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HTTPSHealthCheck{}
	out.Host = in.Host
	out.Port = ConvertIntToInt32(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.RequestPath = in.RequestPath
	out.Response = in.Response
	return out
}

func HealthCheckSSLHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.SSLHealthCheck) *krm.HealthCheckSSLHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckSSLHealthCheck{}
	out.Port = ConvertInt32ToInt(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.Request = in.Request
	out.Response = in.Response
	return out
}

func HealthCheckSSLHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckSSLHealthCheck) *pb.SSLHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.SSLHealthCheck{}
	out.Port = ConvertIntToInt32(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.Request = in.Request
	out.Response = in.Response
	return out
}

func HealthCheckTCPHealthCheck_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.TCPHealthCheck) *krm.HealthCheckTCPHealthCheck {
	if in == nil {
		return nil
	}
	out := &krm.HealthCheckTCPHealthCheck{}
	out.Port = ConvertInt32ToInt(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.Request = in.Request
	out.Response = in.Response
	return out
}

func HealthCheckTCPHealthCheck_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.HealthCheckTCPHealthCheck) *pb.TCPHealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.TCPHealthCheck{}
	out.Port = ConvertIntToInt32(in.Port)
	out.PortName = in.PortName
	out.PortSpecification = in.PortSpecification
	out.ProxyHeader = in.ProxyHeader
	out.Request = in.Request
	out.Response = in.Response
	return out
}

func ComputeHealthCheckStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.HealthCheck) *krm.ComputeHealthCheckStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeHealthCheckStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	out.Type = in.Type
	return out
}

func ComputeHealthCheckStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeHealthCheckStatus) *pb.HealthCheck {
	if in == nil {
		return nil
	}
	out := &pb.HealthCheck{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	// Type is output-only/read-only on GCP, but if ToProto needs it, we can set it.
	out.Type = in.Type
	return out
}

func Int64_FromProto(in *int32) *int64 {
	if in == nil {
		return nil
	}
	out := int64(*in)
	return &out
}

func Int64_ToProto(in *int64) *int32 {
	if in == nil {
		return nil
	}
	out := int32(*in)
	return &out
}
