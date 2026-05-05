package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/alloydb/v1beta"
)

func main() {
	obj := &pb.Cluster{
		DataplexConfig: &pb.Cluster_DataplexConfig{
			Enabled: new(bool),
		},
	}
	*obj.DataplexConfig.Enabled = true
	
	m := &jsonpb.Marshaler{}
	str, _ := m.MarshalToString(obj)
	fmt.Println("Serialized:", str)
}
