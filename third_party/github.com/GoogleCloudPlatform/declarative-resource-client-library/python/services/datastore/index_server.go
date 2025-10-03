// Copyright 2021 Google LLC. All Rights Reserved.
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
package server

import (
	"context"

	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	datastorepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/datastore/datastore_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/datastore"
)

// Server implements the gRPC interface for Index.
type IndexServer struct{}

// ProtoToIndexAncestorEnum converts a IndexAncestorEnum enum from its proto representation.
func ProtoToDatastoreIndexAncestorEnum(e datastorepb.DatastoreIndexAncestorEnum) *datastore.IndexAncestorEnum {
	if e == 0 {
		return nil
	}
	if n, ok := datastorepb.DatastoreIndexAncestorEnum_name[int32(e)]; ok {
		e := datastore.IndexAncestorEnum(n[len("DatastoreIndexAncestorEnum"):])
		return &e
	}
	return nil
}

// ProtoToIndexPropertiesDirectionEnum converts a IndexPropertiesDirectionEnum enum from its proto representation.
func ProtoToDatastoreIndexPropertiesDirectionEnum(e datastorepb.DatastoreIndexPropertiesDirectionEnum) *datastore.IndexPropertiesDirectionEnum {
	if e == 0 {
		return nil
	}
	if n, ok := datastorepb.DatastoreIndexPropertiesDirectionEnum_name[int32(e)]; ok {
		e := datastore.IndexPropertiesDirectionEnum(n[len("DatastoreIndexPropertiesDirectionEnum"):])
		return &e
	}
	return nil
}

// ProtoToIndexStateEnum converts a IndexStateEnum enum from its proto representation.
func ProtoToDatastoreIndexStateEnum(e datastorepb.DatastoreIndexStateEnum) *datastore.IndexStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := datastorepb.DatastoreIndexStateEnum_name[int32(e)]; ok {
		e := datastore.IndexStateEnum(n[len("DatastoreIndexStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToIndexProperties converts a IndexProperties resource from its proto representation.
func ProtoToDatastoreIndexProperties(p *datastorepb.DatastoreIndexProperties) *datastore.IndexProperties {
	if p == nil {
		return nil
	}
	obj := &datastore.IndexProperties{
		Name:      dcl.StringOrNil(p.Name),
		Direction: ProtoToDatastoreIndexPropertiesDirectionEnum(p.GetDirection()),
	}
	return obj
}

// ProtoToIndex converts a Index resource from its proto representation.
func ProtoToIndex(p *datastorepb.DatastoreIndex) *datastore.Index {
	obj := &datastore.Index{
		Ancestor: ProtoToDatastoreIndexAncestorEnum(p.GetAncestor()),
		IndexId:  dcl.StringOrNil(p.IndexId),
		Kind:     dcl.StringOrNil(p.Kind),
		Project:  dcl.StringOrNil(p.Project),
		State:    ProtoToDatastoreIndexStateEnum(p.GetState()),
	}
	for _, r := range p.GetProperties() {
		obj.Properties = append(obj.Properties, *ProtoToDatastoreIndexProperties(r))
	}
	return obj
}

// IndexAncestorEnumToProto converts a IndexAncestorEnum enum to its proto representation.
func DatastoreIndexAncestorEnumToProto(e *datastore.IndexAncestorEnum) datastorepb.DatastoreIndexAncestorEnum {
	if e == nil {
		return datastorepb.DatastoreIndexAncestorEnum(0)
	}
	if v, ok := datastorepb.DatastoreIndexAncestorEnum_value["IndexAncestorEnum"+string(*e)]; ok {
		return datastorepb.DatastoreIndexAncestorEnum(v)
	}
	return datastorepb.DatastoreIndexAncestorEnum(0)
}

// IndexPropertiesDirectionEnumToProto converts a IndexPropertiesDirectionEnum enum to its proto representation.
func DatastoreIndexPropertiesDirectionEnumToProto(e *datastore.IndexPropertiesDirectionEnum) datastorepb.DatastoreIndexPropertiesDirectionEnum {
	if e == nil {
		return datastorepb.DatastoreIndexPropertiesDirectionEnum(0)
	}
	if v, ok := datastorepb.DatastoreIndexPropertiesDirectionEnum_value["IndexPropertiesDirectionEnum"+string(*e)]; ok {
		return datastorepb.DatastoreIndexPropertiesDirectionEnum(v)
	}
	return datastorepb.DatastoreIndexPropertiesDirectionEnum(0)
}

// IndexStateEnumToProto converts a IndexStateEnum enum to its proto representation.
func DatastoreIndexStateEnumToProto(e *datastore.IndexStateEnum) datastorepb.DatastoreIndexStateEnum {
	if e == nil {
		return datastorepb.DatastoreIndexStateEnum(0)
	}
	if v, ok := datastorepb.DatastoreIndexStateEnum_value["IndexStateEnum"+string(*e)]; ok {
		return datastorepb.DatastoreIndexStateEnum(v)
	}
	return datastorepb.DatastoreIndexStateEnum(0)
}

// IndexPropertiesToProto converts a IndexProperties resource to its proto representation.
func DatastoreIndexPropertiesToProto(o *datastore.IndexProperties) *datastorepb.DatastoreIndexProperties {
	if o == nil {
		return nil
	}
	p := &datastorepb.DatastoreIndexProperties{
		Name:      dcl.ValueOrEmptyString(o.Name),
		Direction: DatastoreIndexPropertiesDirectionEnumToProto(o.Direction),
	}
	return p
}

// IndexToProto converts a Index resource to its proto representation.
func IndexToProto(resource *datastore.Index) *datastorepb.DatastoreIndex {
	p := &datastorepb.DatastoreIndex{
		Ancestor: DatastoreIndexAncestorEnumToProto(resource.Ancestor),
		IndexId:  dcl.ValueOrEmptyString(resource.IndexId),
		Kind:     dcl.ValueOrEmptyString(resource.Kind),
		Project:  dcl.ValueOrEmptyString(resource.Project),
		State:    DatastoreIndexStateEnumToProto(resource.State),
	}
	for _, r := range resource.Properties {
		p.Properties = append(p.Properties, DatastoreIndexPropertiesToProto(&r))
	}

	return p
}

// ApplyIndex handles the gRPC request by passing it to the underlying Index Apply() method.
func (s *IndexServer) applyIndex(ctx context.Context, c *datastore.Client, request *datastorepb.ApplyDatastoreIndexRequest) (*datastorepb.DatastoreIndex, error) {
	p := ProtoToIndex(request.GetResource())
	res, err := c.ApplyIndex(ctx, p)
	if err != nil {
		return nil, err
	}
	r := IndexToProto(res)
	return r, nil
}

// ApplyIndex handles the gRPC request by passing it to the underlying Index Apply() method.
func (s *IndexServer) ApplyDatastoreIndex(ctx context.Context, request *datastorepb.ApplyDatastoreIndexRequest) (*datastorepb.DatastoreIndex, error) {
	cl, err := createConfigIndex(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyIndex(ctx, cl, request)
}

// DeleteIndex handles the gRPC request by passing it to the underlying Index Delete() method.
func (s *IndexServer) DeleteDatastoreIndex(ctx context.Context, request *datastorepb.DeleteDatastoreIndexRequest) (*emptypb.Empty, error) {

	cl, err := createConfigIndex(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteIndex(ctx, ProtoToIndex(request.GetResource()))

}

// ListDatastoreIndex handles the gRPC request by passing it to the underlying IndexList() method.
func (s *IndexServer) ListDatastoreIndex(ctx context.Context, request *datastorepb.ListDatastoreIndexRequest) (*datastorepb.ListDatastoreIndexResponse, error) {
	cl, err := createConfigIndex(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListIndex(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*datastorepb.DatastoreIndex
	for _, r := range resources.Items {
		rp := IndexToProto(r)
		protos = append(protos, rp)
	}
	return &datastorepb.ListDatastoreIndexResponse{Items: protos}, nil
}

func createConfigIndex(ctx context.Context, service_account_file string) (*datastore.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return datastore.NewClient(conf), nil
}
