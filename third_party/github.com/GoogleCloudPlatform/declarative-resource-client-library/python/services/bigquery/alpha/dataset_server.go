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
	alphapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/bigquery/alpha/bigquery_alpha_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/bigquery/alpha"
)

// DatasetServer implements the gRPC interface for Dataset.
type DatasetServer struct{}

// ProtoToDatasetAccess converts a DatasetAccess object from its proto representation.
func ProtoToBigqueryAlphaDatasetAccess(p *alphapb.BigqueryAlphaDatasetAccess) *alpha.DatasetAccess {
	if p == nil {
		return nil
	}
	obj := &alpha.DatasetAccess{
		Role:         dcl.StringOrNil(p.GetRole()),
		UserByEmail:  dcl.StringOrNil(p.GetUserByEmail()),
		GroupByEmail: dcl.StringOrNil(p.GetGroupByEmail()),
		Domain:       dcl.StringOrNil(p.GetDomain()),
		SpecialGroup: dcl.StringOrNil(p.GetSpecialGroup()),
		IamMember:    dcl.StringOrNil(p.GetIamMember()),
		View:         ProtoToBigqueryAlphaDatasetAccessView(p.GetView()),
		Routine:      ProtoToBigqueryAlphaDatasetAccessRoutine(p.GetRoutine()),
	}
	return obj
}

// ProtoToDatasetAccessView converts a DatasetAccessView object from its proto representation.
func ProtoToBigqueryAlphaDatasetAccessView(p *alphapb.BigqueryAlphaDatasetAccessView) *alpha.DatasetAccessView {
	if p == nil {
		return nil
	}
	obj := &alpha.DatasetAccessView{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		TableId:   dcl.StringOrNil(p.GetTableId()),
	}
	return obj
}

// ProtoToDatasetAccessRoutine converts a DatasetAccessRoutine object from its proto representation.
func ProtoToBigqueryAlphaDatasetAccessRoutine(p *alphapb.BigqueryAlphaDatasetAccessRoutine) *alpha.DatasetAccessRoutine {
	if p == nil {
		return nil
	}
	obj := &alpha.DatasetAccessRoutine{
		ProjectId: dcl.StringOrNil(p.GetProjectId()),
		DatasetId: dcl.StringOrNil(p.GetDatasetId()),
		RoutineId: dcl.StringOrNil(p.GetRoutineId()),
	}
	return obj
}

// ProtoToDatasetDefaultEncryptionConfiguration converts a DatasetDefaultEncryptionConfiguration object from its proto representation.
func ProtoToBigqueryAlphaDatasetDefaultEncryptionConfiguration(p *alphapb.BigqueryAlphaDatasetDefaultEncryptionConfiguration) *alpha.DatasetDefaultEncryptionConfiguration {
	if p == nil {
		return nil
	}
	obj := &alpha.DatasetDefaultEncryptionConfiguration{
		KmsKeyName: dcl.StringOrNil(p.GetKmsKeyName()),
	}
	return obj
}

// ProtoToDataset converts a Dataset resource from its proto representation.
func ProtoToDataset(p *alphapb.BigqueryAlphaDataset) *alpha.Dataset {
	obj := &alpha.Dataset{
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
		DefaultEncryptionConfiguration: ProtoToBigqueryAlphaDatasetDefaultEncryptionConfiguration(p.GetDefaultEncryptionConfiguration()),
	}
	for _, r := range p.GetAccess() {
		obj.Access = append(obj.Access, *ProtoToBigqueryAlphaDatasetAccess(r))
	}
	return obj
}

// DatasetAccessToProto converts a DatasetAccess object to its proto representation.
func BigqueryAlphaDatasetAccessToProto(o *alpha.DatasetAccess) *alphapb.BigqueryAlphaDatasetAccess {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaDatasetAccess{}
	p.SetRole(dcl.ValueOrEmptyString(o.Role))
	p.SetUserByEmail(dcl.ValueOrEmptyString(o.UserByEmail))
	p.SetGroupByEmail(dcl.ValueOrEmptyString(o.GroupByEmail))
	p.SetDomain(dcl.ValueOrEmptyString(o.Domain))
	p.SetSpecialGroup(dcl.ValueOrEmptyString(o.SpecialGroup))
	p.SetIamMember(dcl.ValueOrEmptyString(o.IamMember))
	p.SetView(BigqueryAlphaDatasetAccessViewToProto(o.View))
	p.SetRoutine(BigqueryAlphaDatasetAccessRoutineToProto(o.Routine))
	return p
}

// DatasetAccessViewToProto converts a DatasetAccessView object to its proto representation.
func BigqueryAlphaDatasetAccessViewToProto(o *alpha.DatasetAccessView) *alphapb.BigqueryAlphaDatasetAccessView {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaDatasetAccessView{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetTableId(dcl.ValueOrEmptyString(o.TableId))
	return p
}

// DatasetAccessRoutineToProto converts a DatasetAccessRoutine object to its proto representation.
func BigqueryAlphaDatasetAccessRoutineToProto(o *alpha.DatasetAccessRoutine) *alphapb.BigqueryAlphaDatasetAccessRoutine {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaDatasetAccessRoutine{}
	p.SetProjectId(dcl.ValueOrEmptyString(o.ProjectId))
	p.SetDatasetId(dcl.ValueOrEmptyString(o.DatasetId))
	p.SetRoutineId(dcl.ValueOrEmptyString(o.RoutineId))
	return p
}

// DatasetDefaultEncryptionConfigurationToProto converts a DatasetDefaultEncryptionConfiguration object to its proto representation.
func BigqueryAlphaDatasetDefaultEncryptionConfigurationToProto(o *alpha.DatasetDefaultEncryptionConfiguration) *alphapb.BigqueryAlphaDatasetDefaultEncryptionConfiguration {
	if o == nil {
		return nil
	}
	p := &alphapb.BigqueryAlphaDatasetDefaultEncryptionConfiguration{}
	p.SetKmsKeyName(dcl.ValueOrEmptyString(o.KmsKeyName))
	return p
}

// DatasetToProto converts a Dataset resource to its proto representation.
func DatasetToProto(resource *alpha.Dataset) *alphapb.BigqueryAlphaDataset {
	p := &alphapb.BigqueryAlphaDataset{}
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
	p.SetDefaultEncryptionConfiguration(BigqueryAlphaDatasetDefaultEncryptionConfigurationToProto(resource.DefaultEncryptionConfiguration))
	mLabels := make(map[string]string, len(resource.Labels))
	for k, r := range resource.Labels {
		mLabels[k] = r
	}
	p.SetLabels(mLabels)
	sAccess := make([]*alphapb.BigqueryAlphaDatasetAccess, len(resource.Access))
	for i, r := range resource.Access {
		sAccess[i] = BigqueryAlphaDatasetAccessToProto(&r)
	}
	p.SetAccess(sAccess)

	return p
}

// applyDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) applyDataset(ctx context.Context, c *alpha.Client, request *alphapb.ApplyBigqueryAlphaDatasetRequest) (*alphapb.BigqueryAlphaDataset, error) {
	p := ProtoToDataset(request.GetResource())
	res, err := c.ApplyDataset(ctx, p)
	if err != nil {
		return nil, err
	}
	r := DatasetToProto(res)
	return r, nil
}

// applyBigqueryAlphaDataset handles the gRPC request by passing it to the underlying Dataset Apply() method.
func (s *DatasetServer) ApplyBigqueryAlphaDataset(ctx context.Context, request *alphapb.ApplyBigqueryAlphaDatasetRequest) (*alphapb.BigqueryAlphaDataset, error) {
	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return s.applyDataset(ctx, cl, request)
}

// DeleteDataset handles the gRPC request by passing it to the underlying Dataset Delete() method.
func (s *DatasetServer) DeleteBigqueryAlphaDataset(ctx context.Context, request *alphapb.DeleteBigqueryAlphaDatasetRequest) (*emptypb.Empty, error) {

	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteDataset(ctx, ProtoToDataset(request.GetResource()))

}

// ListBigqueryAlphaDataset handles the gRPC request by passing it to the underlying DatasetList() method.
func (s *DatasetServer) ListBigqueryAlphaDataset(ctx context.Context, request *alphapb.ListBigqueryAlphaDatasetRequest) (*alphapb.ListBigqueryAlphaDatasetResponse, error) {
	cl, err := createConfigDataset(ctx, request.GetServiceAccountFile())
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListDataset(ctx, request.GetProject())
	if err != nil {
		return nil, err
	}
	var protos []*alphapb.BigqueryAlphaDataset
	for _, r := range resources.Items {
		rp := DatasetToProto(r)
		protos = append(protos, rp)
	}
	p := &alphapb.ListBigqueryAlphaDatasetResponse{}
	p.SetItems(protos)
	return p, nil
}

func createConfigDataset(ctx context.Context, service_account_file string) (*alpha.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return alpha.NewClient(conf), nil
}
