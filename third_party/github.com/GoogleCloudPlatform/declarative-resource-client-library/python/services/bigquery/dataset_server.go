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
package server

import (
	"context"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	bigquerypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/bigquery_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery"
)

// DatasetServer implements the gRPC interface for Dataset.
type DatasetServer struct{}

// ProtoToDatasetAccess converts a DatasetAccess object from its proto representation.
func ProtoToBigqueryDatasetAccess(p *bigquerypb.BigqueryDatasetAccess) *bigquery.DatasetAccess {
	if p == nil {
		return nil
	}
	obj := &bigquery.DatasetAccess{
		Role:         dcl.StringOrNil(p.GetRole()),
		UserByEmail:  dcl.StringOrNil(p.GetUserByEmail()),
		GroupByEmail: dcl.StringOrNil(p.GetGroupByEmail()),
		Domain:       dcl.StringOrNil(p.GetDomain()),
		SpecialGroup: dcl.StringOrNil(p.GetSpecialGroup()),
		IamMember:    dcl.StringOrNil(p.GetIamMember()),
		View:         ProtoToBigqueryDatasetAccessView(p.GetView()),
		Routine:      ProtoToBigqueryDatasetAccessRoutine(p.GetRoutine()),
	}
	return obj
}

// ProtoToDatasetAccessView converts a DatasetAccessView object from its proto representation.
func ProtoToBigqueryDatasetAccessView(p *bigquerypb.BigqueryDatasetAccessView) *bigquery.DatasetAccessView {
	if p == nil {
		return nil
	}
	obj := &bigquery.DatasetAccessView{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToDatasetAccessRoutine converts a DatasetAccessRoutine object from its proto representation.
func ProtoToBigqueryDatasetAccessRoutine(p *bigquerypb.BigqueryDatasetAccessRoutine) *bigquery.DatasetAccessRoutine {
	if p == nil {
		return nil
	}
	obj := &bigquery.DatasetAccessRoutine{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		RoutineId: dcl.StringOrNil(p.GetRoutineId()),
	}
	return obj
}

// ProtoToDatasetDefaultEncryptionConfiguration converts a DatasetDefaultEncryptionConfiguration object from its proto representation.
func ProtoToBigqueryDatasetDefaultEncryptionConfiguration(p *bigquerypb.BigqueryDatasetDefaultEncryptionConfiguration) *bigquery.DatasetDefaultEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &bigquery.DatasetDefaultEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToDataset converts a Dataset resource from its proto representation.
func ProtoToDataset(p *bigquerypb.BigqueryDataset) *bigquery.Dataset {
	obj := &bigquery.Dataset{
		Etag:                           dcl.StringOrNil(p.GetEtag()),
		Id:                             dcl.StringOrNil(p.GetId()),
		SelfLink:                       dcl.StringOrNil(p.GetSelfLink()),
		Name:                           dcl.StringOrNil(p.GetName()),
		Project:                        dcl.StringOrNil(p.GetProject()),
		FriendlyName:                   dcl.StringOrNil(p.GetFriendlyName()),
		Description:                    dcl.StringOrNil(p.GetDescription()),
		DefaultTableExpirationMs:       dcl.StringOrNil(p.GetDefaultTableExpirationMs()),
		DefaultPartitionExpirationMs:   dcl.StringOrNil(p.GetDefaultPartitionExpirationMs()),
		CreationTime:                   dcl.Int64OrNil(p.GetCreationTime()),
		LastModifiedTime:               dcl.Int64OrNil(p.GetLastModifiedTime()),
		Location:                       dcl.StringOrNil(p.GetLocation()),
		Published:                      dcl.Bool(p.GetPublished()),
		DefaultEncryptionConfiguration: ProtoToBigqueryDatasetDefaultEncryptionConfiguration(p.GetDefaultEncryptionConfiguration()),
	}
	for _, r := range p.GetAccess() {
		obj.Access = append(obj.Access, *ProtoToBigqueryDatasetAccess(r))
	}
	return obj
}

// DatasetAccessToProto converts a DatasetAccess object to its proto representation.
func BigqueryDatasetAccessToProto(o *bigquery.DatasetAccess) *bigquerypb.BigqueryDatasetAccess {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryDatasetAccess{}
	p.SetRole(dcl.ValueOrEmptyString(o.Role))
	p.SetUserByEmail(dcl.ValueOrEmptyString(o.UserByEmail))
	p.SetGroupByEmail(dcl.ValueOrEmptyString(o.GroupByEmail))
	p.SetDomain(dcl.ValueOrEmptyString(o.Domain))
	p.SetSpecialGroup(dcl.ValueOrEmptyString(o.SpecialGroup))
	p.SetIamMember(dcl.ValueOrEmptyString(o.IamMember))
	p.SetView(BigqueryDatasetAccessViewToProto(o.View))
	p.SetRoutine(BigqueryDatasetAccessRoutineToProto(o.Routine))
	return p
}

// DatasetAccessViewToProto converts a DatasetAccessView object to its proto representation.
func BigqueryDatasetAccessViewToProto(o *bigquery.DatasetAccessView) *bigquerypb.BigqueryDatasetAccessView {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryDatasetAccessView{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// DatasetAccessRoutineToProto converts a DatasetAccessRoutine object to its proto representation.
func BigqueryDatasetAccessRoutineToProto(o *bigquery.DatasetAccessRoutine) *bigquerypb.BigqueryDatasetAccessRoutine {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryDatasetAccessRoutine{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetRoutineId(dcl.ValueOrEmptyString(o.RoutineId))
	return p
}

// DatasetDefaultEncryptionConfigurationToProto converts a DatasetDefaultEncryptionConfiguration object to its proto representation.
func BigqueryDatasetDefaultEncryptionConfigurationToProto(o *bigquery.DatasetDefaultEncryptionConfiguration) *bigquerypb.BigqueryDatasetDefaultEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &bigquerypb.BigqueryDatasetDefaultEncryptionConfiguration{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// DatasetToProto converts a Dataset resource to its proto representation.
func DatasetToProto(resource *bigquery.Dataset) *bigquerypb.BigqueryDataset {
	p := &bigquerypb.BigqueryDataset{}
	p.SetEtag(dcl.ValueOrEmptyString(resource.Etag))
	p.SetId(dcl.ValueOrEmptyString(resource.Id))
	p.SetSelfLink(dcl.ValueOrEmptyString(resource.SelfLink))
	p.SetName(dcl.ValueOrEmptyString(resource.Name))
	p.SetProject(dcl.ValueOrEmptyString(resource.Project))
	p.SetFriendlyName(dcl.ValueOrEmptyString(resource.FriendlyName))
	p.SetDescription(dcl.ValueOrEmptyString(resource.Description))
	p.SetDefaultTableExpirationMs(dcl.ValueOrEmptyString(resource.DefaultTableExpirationMs))
	p.SetDefaultPartitionExpirationMs(dcl.ValueOrEmptyString(resource.DefaultPartitionExpirationMs))
	p.SetCreationTime(dcl.ValueOrEmptyInt64(resource.CreationTime))
	p.SetLastModifiedTime(dcl.ValueOrEmptyInt64(resource.LastModifiedTime))
	p.SetLocation(dcl.ValueOrEmptyString(resource.Location))
	p.SetPublished(dcl.ValueOrEmptyBool(resource.Published))
	p.SetDefaultEncryptionConfiguration(BigqueryDatasetDefaultEncryptionConfigurationToProto(resource.DefaultEncryptionConfiguration))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sAccess := make([]*bigquerypb.BigqueryDatasetAccess, len(resource.Access))
	for i, r := range resource.Access {
		sAccess[i] = BigqueryDatasetAccessToProto(&r)
	}
	p.SetAccess(sAccess)

	return p
}

// applyDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) applyDataset(ctx context.Context, c *bigquery.Client, request *bigquerypb.ApplyBigqueryDatasetRequest) (*bigquerypb.BigqueryDataset, error) {
	p := ProtoToDataset(request.GetResource())
	res, err := c.ApplyDataset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DatasetToProto(res)
	return r, nil
}

// applyBigqueryDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) ApplyBigqueryDataset(ctx context.Context, request *bigquerypb.ApplyBigqueryDatasetRequest) (*bigquerypb.BigqueryDataset, error) {
	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDataset(ctx, cl, request)
}

// DeleteDataset handles the gRPC request by passing it to the underlying Dataset Delete() method.
func (s *DatasetServer) DeleteBigqueryDataset(ctx context.Context, request *bigquerypb.DeleteBigqueryDatasetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDataset(ctx, ProtoToDataset(request.GetResource()))

}

// ListBigqueryDataset handles the gRPC request by passing it to the underlying DatasetList() method.
func (s *DatasetServer) ListBigqueryDataset(ctx context.Context, request *bigquerypb.ListBigqueryDatasetRequest) (*bigquerypb.ListBigqueryDatasetResponse, error) {
	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDataset(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*bigquerypb.BigqueryDataset
	for _, r := range resources.Items {
		rp := DatasetToProto(r)
		protos = append(protos, rp)
	}
	p := &bigquerypb.ListBigqueryDatasetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDataset(ctx context.Context, service_account_file string) (*bigquery.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return bigquery.NewClient(conf), nil
}
