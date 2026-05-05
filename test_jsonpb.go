package main

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/alloydb/v1beta"
)

func main() {
	js := `{"dataplexConfig":{"enabled":true}}`
	obj := &pb.Cluster{}
	err := jsonpb.Unmarshal(bytes.NewReader([]byte(js)), obj)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Parsed: %+v\n", obj.DataplexConfig)
}
