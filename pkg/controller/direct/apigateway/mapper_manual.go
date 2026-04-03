package apigateway

import (
	"encoding/base64"

	pb "cloud.google.com/go/apigateway/apiv1/apigatewaypb"
	krmapigatewayv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/apigateway/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func APIGatewayAPIConfig_File_v1alpha1_FromProto(mapCtx *direct.MapContext, in *pb.ApiConfig_File) *krmapigatewayv1alpha1.APIGatewayAPIConfig_File {
	if in == nil {
		return nil
	}
	out := &krmapigatewayv1alpha1.APIGatewayAPIConfig_File{}
	out.Path = direct.LazyPtr(in.GetPath())

	if in.GetContents() != nil {
		s := base64.StdEncoding.EncodeToString(in.GetContents())
		out.Contents = &s
	}
	return out
}

func APIGatewayAPIConfig_File_v1alpha1_ToProto(mapCtx *direct.MapContext, in *krmapigatewayv1alpha1.APIGatewayAPIConfig_File) *pb.ApiConfig_File {
	if in == nil {
		return nil
	}
	out := &pb.ApiConfig_File{}
	out.Path = direct.ValueOf(in.Path)
	if in.Contents != nil {
		b, err := base64.StdEncoding.DecodeString(*in.Contents)
		if err != nil {
			mapCtx.Errorf("error decoding base64 contents for file %q: %v", direct.ValueOf(in.Path), err)
		} else {
			out.Contents = b
		}
	}
	return out
}
