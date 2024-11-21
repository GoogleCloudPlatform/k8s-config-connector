package mockalloydb

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/alloydb/v1beta"
)

func constructOperationMetadata(target, verb string) *pb.OperationMetadata {
	now := timestamppb.Now()
	return &pb.OperationMetadata{
		Target:                target,
		CreateTime:            now,
		ApiVersion:            "v1beta",
		RequestedCancellation: false,
		Verb:                  verb,
	}
}
