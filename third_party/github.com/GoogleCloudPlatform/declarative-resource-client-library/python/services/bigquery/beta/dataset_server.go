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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/beta/bigquery_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/beta"
)

// DatasetServer implements the gRPC interface for Dataset.
type DatasetServer struct{}

// ProtoToDatasetAccess converts a DatasetAccess object from its proto representation.
func ProtoToBigqueryBetaDatasetAccess(p *betapb.BigqueryBetaDatasetAccess) *beta.DatasetAccess {
	if p == nil {
		return nil
	}
	obj := &beta.DatasetAccess{
		Role:         dcl.StringOrNil(p.GetRole()),
		UserByEmail:  dcl.StringOrNil(p.GetUserByEmail()),
		GroupByEmail: dcl.StringOrNil(p.GetGroupByEmail()),
		Domain:       dcl.StringOrNil(p.GetDomain()),
		SpecialGroup: dcl.StringOrNil(p.GetSpecialGroup()),
		IamMember:    dcl.StringOrNil(p.GetIamMember()),
		View:         ProtoToBigqueryBetaDatasetAccessView(p.GetView()),
		Routine:      ProtoToBigqueryBetaDatasetAccessRoutine(p.GetRoutine()),
	}
	return obj
}

// ProtoToDatasetAccessView converts a DatasetAccessView object from its proto representation.
func ProtoToBigqueryBetaDatasetAccessView(p *betapb.BigqueryBetaDatasetAccessView) *beta.DatasetAccessView {
	if p == nil {
		return nil
	}
	obj := &beta.DatasetAccessView{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToDatasetAccessRoutine converts a DatasetAccessRoutine object from its proto representation.
func ProtoToBigqueryBetaDatasetAccessRoutine(p *betapb.BigqueryBetaDatasetAccessRoutine) *beta.DatasetAccessRoutine {
	if p == nil {
		return nil
	}
	obj := &beta.DatasetAccessRoutine{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		RoutineId: dcl.StringOrNil(p.GetRoutineId()),
	}
	return obj
}

// ProtoToDatasetDefaultEncryptionConfiguration converts a DatasetDefaultEncryptionConfiguration object from its proto representation.
func ProtoToBigqueryBetaDatasetDefaultEncryptionConfiguration(p *betapb.BigqueryBetaDatasetDefaultEncryptionConfiguration) *beta.DatasetDefaultEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &beta.DatasetDefaultEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToDataset converts a Dataset resource from its proto representation.
func ProtoToDataset(p *betapb.BigqueryBetaDataset) *beta.Dataset {
	obj := &beta.Dataset{
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
		DefaultEncryptionConfiguration: ProtoToBigqueryBetaDatasetDefaultEncryptionConfiguration(p.GetDefaultEncryptionConfiguration()),
	}
	for _, r := range p.GetAccess() {
		obj.Access = append(obj.Access, *ProtoToBigqueryBetaDatasetAccess(r))
	}
	return obj
}

// DatasetAccessToProto converts a DatasetAccess object to its proto representation.
func BigqueryBetaDatasetAccessToProto(o *beta.DatasetAccess) *betapb.BigqueryBetaDatasetAccess {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaDatasetAccess{}
	p.SetRole(dcl.ValueOrEmptyString(o.Role))
	p.SetUserByEmail(dcl.ValueOrEmptyString(o.UserByEmail))
	p.SetGroupByEmail(dcl.ValueOrEmptyString(o.GroupByEmail))
	p.SetDomain(dcl.ValueOrEmptyString(o.Domain))
	p.SetSpecialGroup(dcl.ValueOrEmptyString(o.SpecialGroup))
	p.SetIamMember(dcl.ValueOrEmptyString(o.IamMember))
	p.SetView(BigqueryBetaDatasetAccessViewToProto(o.View))
	p.SetRoutine(BigqueryBetaDatasetAccessRoutineToProto(o.Routine))
	return p
}

// DatasetAccessViewToProto converts a DatasetAccessView object to its proto representation.
func BigqueryBetaDatasetAccessViewToProto(o *beta.DatasetAccessView) *betapb.BigqueryBetaDatasetAccessView {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaDatasetAccessView{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// DatasetAccessRoutineToProto converts a DatasetAccessRoutine object to its proto representation.
func BigqueryBetaDatasetAccessRoutineToProto(o *beta.DatasetAccessRoutine) *betapb.BigqueryBetaDatasetAccessRoutine {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaDatasetAccessRoutine{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetRoutineId(dcl.ValueOrEmptyString(o.RoutineId))
	return p
}

// DatasetDefaultEncryptionConfigurationToProto converts a DatasetDefaultEncryptionConfiguration object to its proto representation.
func BigqueryBetaDatasetDefaultEncryptionConfigurationToProto(o *beta.DatasetDefaultEncryptionConfiguration) *betapb.BigqueryBetaDatasetDefaultEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &betapb.BigqueryBetaDatasetDefaultEncryptionConfiguration{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// DatasetToProto converts a Dataset resource to its proto representation.
func DatasetToProto(resource *beta.Dataset) *betapb.BigqueryBetaDataset {
	p := &betapb.BigqueryBetaDataset{}
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
	p.SetDefaultEncryptionConfiguration(BigqueryBetaDatasetDefaultEncryptionConfigurationToProto(resource.DefaultEncryptionConfiguration))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sAccess := make([]*betapb.BigqueryBetaDatasetAccess, len(resource.Access))
	for i, r := range resource.Access {
		sAccess[i] = BigqueryBetaDatasetAccessToProto(&r)
	}
	p.SetAccess(sAccess)

	return p
}

// applyDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) applyDataset(ctx context.Context, c *beta.Client, request *betapb.ApplyBigqueryBetaDatasetRequest) (*betapb.BigqueryBetaDataset, error) {
	p := ProtoToDataset(request.GetResource())
	res, err := c.ApplyDataset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DatasetToProto(res)
	return r, nil
}

// applyBigqueryBetaDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) ApplyBigqueryBetaDataset(ctx context.Context, request *betapb.ApplyBigqueryBetaDatasetRequest) (*betapb.BigqueryBetaDataset, error) {
	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDataset(ctx, cl, request)
}

// DeleteDataset handles the gRPC request by passing it to the underlying Dataset Delete() method.
func (s *DatasetServer) DeleteBigqueryBetaDataset(ctx context.Context, request *betapb.DeleteBigqueryBetaDatasetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDataset(ctx, ProtoToDataset(request.GetResource()))

}

// ListBigqueryBetaDataset handles the gRPC request by passing it to the underlying DatasetList() method.
func (s *DatasetServer) ListBigqueryBetaDataset(ctx context.Context, request *betapb.ListBigqueryBetaDatasetRequest) (*betapb.ListBigqueryBetaDatasetResponse, error) {
	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDataset(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*betapb.BigqueryBetaDataset
	for _, r := range resources.Items {
		rp := DatasetToProto(r)
		protos = append(protos, rp)
	}
	p := &betapb.ListBigqueryBetaDatasetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDataset(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
