// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package connector provides the bridge between the MMv2 go library and other languages.
package main

import (
	"C"
	"unsafe"

	glog "github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"

	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	connectorpb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/connector_go_proto"
)

var unaryCall = UnaryCall

var errFailedToMarshal = func() []byte {
	b, err := proto.Marshal(connectorpb.UnaryCallResponse_builder{
		Status: &statuspb.Status{
			Code:    int32(codes.Internal),
			Message: "cannot marshal error message; see server logs",
		},
	}.Build())
	if err != nil {
		glog.Exitf("Could not initialize errFailedToMarshal: %v", err)
	}
	return b
}()

// Initialize exposes the C interface for the InitializeServer() method.
//
//export Initialize
func Initialize(request []byte) (unsafe.Pointer, int) {
	protoRequest := &connectorpb.InitializeRequest{}
	err := proto.Unmarshal(request, protoRequest)
	if err != nil {
		return initializeError(err)
	}

	protoResponse := InitializeServer()
	response, err := proto.Marshal(protoResponse)
	if err != nil {
		return initializeError(err)
	}

	return C.CBytes(response), len(response)
}

func initializeError(err error) (unsafe.Pointer, int) {
	b, err := proto.Marshal(connectorpb.InitializeResponse_builder{
		Status: &statuspb.Status{
			Code:    int32(codes.Internal),
			Message: err.Error(),
		},
	}.Build())
	if err != nil {
		return C.CBytes(errFailedToMarshal), len(errFailedToMarshal)
	}
	return C.CBytes(b), len(b)
}

// Call exposes the C interface for the UnaryCall() method.
//
//export Call
func Call(request []byte) (unsafe.Pointer, int) {
	protoRequest := &connectorpb.UnaryCallRequest{}
	err := proto.Unmarshal(request, protoRequest)
	if err != nil {
		return callError(err)
	}

	protoResponse := unaryCall(protoRequest)
	response, err := proto.Marshal(protoResponse)
	if err != nil {
		return callError(err)
	}

	return C.CBytes(response), len(response)
}

func callError(err error) (unsafe.Pointer, int) {
	b, err := proto.Marshal(connectorpb.UnaryCallResponse_builder{
		Status: &statuspb.Status{
			Code:    int32(codes.Internal),
			Message: err.Error(),
		},
	}.Build())
	if err != nil {
		return C.CBytes(errFailedToMarshal), len(errFailedToMarshal)
	}
	return C.CBytes(b), len(b)
}

// We need this for external cgo, it's not used here, though.
func main() {}
