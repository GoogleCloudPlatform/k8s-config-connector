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
	betapb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/beta/compute_beta_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

// Server implements the gRPC interface for Image.
type ImageServer struct{}

// ProtoToImageGuestOSFeatureTypeEnum converts a ImageGuestOSFeatureTypeEnum enum from its proto representation.
func ProtoToComputeBetaImageGuestOSFeatureTypeEnum(e betapb.ComputeBetaImageGuestOSFeatureTypeEnum) *beta.ImageGuestOSFeatureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageGuestOSFeatureTypeEnum_name[int32(e)]; ok {
		e := beta.ImageGuestOSFeatureTypeEnum(n[len("ComputeBetaImageGuestOSFeatureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageRawDiskContainerTypeEnum converts a ImageRawDiskContainerTypeEnum enum from its proto representation.
func ProtoToComputeBetaImageRawDiskContainerTypeEnum(e betapb.ComputeBetaImageRawDiskContainerTypeEnum) *beta.ImageRawDiskContainerTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageRawDiskContainerTypeEnum_name[int32(e)]; ok {
		e := beta.ImageRawDiskContainerTypeEnum(n[len("ComputeBetaImageRawDiskContainerTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageShieldedInstanceInitialStatePkFileTypeEnum converts a ImageShieldedInstanceInitialStatePkFileTypeEnum enum from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum(e betapb.ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum) *beta.ImageShieldedInstanceInitialStatePkFileTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum_name[int32(e)]; ok {
		e := beta.ImageShieldedInstanceInitialStatePkFileTypeEnum(n[len("ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageShieldedInstanceInitialStateKekFileTypeEnum converts a ImageShieldedInstanceInitialStateKekFileTypeEnum enum from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum(e betapb.ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum) *beta.ImageShieldedInstanceInitialStateKekFileTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum_name[int32(e)]; ok {
		e := beta.ImageShieldedInstanceInitialStateKekFileTypeEnum(n[len("ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageShieldedInstanceInitialStateDbFileTypeEnum converts a ImageShieldedInstanceInitialStateDbFileTypeEnum enum from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum(e betapb.ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum) *beta.ImageShieldedInstanceInitialStateDbFileTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum_name[int32(e)]; ok {
		e := beta.ImageShieldedInstanceInitialStateDbFileTypeEnum(n[len("ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageShieldedInstanceInitialStateDbxFileTypeEnum converts a ImageShieldedInstanceInitialStateDbxFileTypeEnum enum from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum(e betapb.ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum) *beta.ImageShieldedInstanceInitialStateDbxFileTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum_name[int32(e)]; ok {
		e := beta.ImageShieldedInstanceInitialStateDbxFileTypeEnum(n[len("ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageSourceTypeEnum converts a ImageSourceTypeEnum enum from its proto representation.
func ProtoToComputeBetaImageSourceTypeEnum(e betapb.ComputeBetaImageSourceTypeEnum) *beta.ImageSourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageSourceTypeEnum_name[int32(e)]; ok {
		e := beta.ImageSourceTypeEnum(n[len("ComputeBetaImageSourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageStatusEnum converts a ImageStatusEnum enum from its proto representation.
func ProtoToComputeBetaImageStatusEnum(e betapb.ComputeBetaImageStatusEnum) *beta.ImageStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageStatusEnum_name[int32(e)]; ok {
		e := beta.ImageStatusEnum(n[len("ComputeBetaImageStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageDeprecatedStateEnum converts a ImageDeprecatedStateEnum enum from its proto representation.
func ProtoToComputeBetaImageDeprecatedStateEnum(e betapb.ComputeBetaImageDeprecatedStateEnum) *beta.ImageDeprecatedStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := betapb.ComputeBetaImageDeprecatedStateEnum_name[int32(e)]; ok {
		e := beta.ImageDeprecatedStateEnum(n[len("ComputeBetaImageDeprecatedStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageGuestOSFeature converts a ImageGuestOSFeature resource from its proto representation.
func ProtoToComputeBetaImageGuestOSFeature(p *betapb.ComputeBetaImageGuestOSFeature) *beta.ImageGuestOSFeature {
	if p == nil {
		return nil
	}
	obj := &beta.ImageGuestOSFeature{
		Type: ProtoToComputeBetaImageGuestOSFeatureTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToImageImageEncryptionKey converts a ImageImageEncryptionKey resource from its proto representation.
func ProtoToComputeBetaImageImageEncryptionKey(p *betapb.ComputeBetaImageImageEncryptionKey) *beta.ImageImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.ImageImageEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToImageRawDisk converts a ImageRawDisk resource from its proto representation.
func ProtoToComputeBetaImageRawDisk(p *betapb.ComputeBetaImageRawDisk) *beta.ImageRawDisk {
	if p == nil {
		return nil
	}
	obj := &beta.ImageRawDisk{
		Source:        dcl.StringOrNil(p.Source),
		Sha1Checksum:  dcl.StringOrNil(p.Sha1Checksum),
		ContainerType: ProtoToComputeBetaImageRawDiskContainerTypeEnum(p.GetContainerType()),
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialState converts a ImageShieldedInstanceInitialState resource from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialState(p *betapb.ComputeBetaImageShieldedInstanceInitialState) *beta.ImageShieldedInstanceInitialState {
	if p == nil {
		return nil
	}
	obj := &beta.ImageShieldedInstanceInitialState{
		Pk: ProtoToComputeBetaImageShieldedInstanceInitialStatePk(p.GetPk()),
	}
	for _, r := range p.GetKek() {
		obj.Kek = append(obj.Kek, *ProtoToComputeBetaImageShieldedInstanceInitialStateKek(r))
	}
	for _, r := range p.GetDb() {
		obj.Db = append(obj.Db, *ProtoToComputeBetaImageShieldedInstanceInitialStateDb(r))
	}
	for _, r := range p.GetDbx() {
		obj.Dbx = append(obj.Dbx, *ProtoToComputeBetaImageShieldedInstanceInitialStateDbx(r))
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialStatePk converts a ImageShieldedInstanceInitialStatePk resource from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialStatePk(p *betapb.ComputeBetaImageShieldedInstanceInitialStatePk) *beta.ImageShieldedInstanceInitialStatePk {
	if p == nil {
		return nil
	}
	obj := &beta.ImageShieldedInstanceInitialStatePk{
		Content:  dcl.StringOrNil(p.Content),
		FileType: ProtoToComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum(p.GetFileType()),
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialStateKek converts a ImageShieldedInstanceInitialStateKek resource from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialStateKek(p *betapb.ComputeBetaImageShieldedInstanceInitialStateKek) *beta.ImageShieldedInstanceInitialStateKek {
	if p == nil {
		return nil
	}
	obj := &beta.ImageShieldedInstanceInitialStateKek{
		Content:  dcl.StringOrNil(p.Content),
		FileType: ProtoToComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum(p.GetFileType()),
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialStateDb converts a ImageShieldedInstanceInitialStateDb resource from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialStateDb(p *betapb.ComputeBetaImageShieldedInstanceInitialStateDb) *beta.ImageShieldedInstanceInitialStateDb {
	if p == nil {
		return nil
	}
	obj := &beta.ImageShieldedInstanceInitialStateDb{
		Content:  dcl.StringOrNil(p.Content),
		FileType: ProtoToComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum(p.GetFileType()),
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialStateDbx converts a ImageShieldedInstanceInitialStateDbx resource from its proto representation.
func ProtoToComputeBetaImageShieldedInstanceInitialStateDbx(p *betapb.ComputeBetaImageShieldedInstanceInitialStateDbx) *beta.ImageShieldedInstanceInitialStateDbx {
	if p == nil {
		return nil
	}
	obj := &beta.ImageShieldedInstanceInitialStateDbx{
		Content:  dcl.StringOrNil(p.Content),
		FileType: ProtoToComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum(p.GetFileType()),
	}
	return obj
}

// ProtoToImageSourceDiskEncryptionKey converts a ImageSourceDiskEncryptionKey resource from its proto representation.
func ProtoToComputeBetaImageSourceDiskEncryptionKey(p *betapb.ComputeBetaImageSourceDiskEncryptionKey) *beta.ImageSourceDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.ImageSourceDiskEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToImageSourceImageEncryptionKey converts a ImageSourceImageEncryptionKey resource from its proto representation.
func ProtoToComputeBetaImageSourceImageEncryptionKey(p *betapb.ComputeBetaImageSourceImageEncryptionKey) *beta.ImageSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.ImageSourceImageEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToImageSourceSnapshotEncryptionKey converts a ImageSourceSnapshotEncryptionKey resource from its proto representation.
func ProtoToComputeBetaImageSourceSnapshotEncryptionKey(p *betapb.ComputeBetaImageSourceSnapshotEncryptionKey) *beta.ImageSourceSnapshotEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &beta.ImageSourceSnapshotEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToImageDeprecated converts a ImageDeprecated resource from its proto representation.
func ProtoToComputeBetaImageDeprecated(p *betapb.ComputeBetaImageDeprecated) *beta.ImageDeprecated {
	if p == nil {
		return nil
	}
	obj := &beta.ImageDeprecated{
		State:       ProtoToComputeBetaImageDeprecatedStateEnum(p.GetState()),
		Replacement: dcl.StringOrNil(p.Replacement),
		Deprecated:  dcl.StringOrNil(p.Deprecated),
		Obsolete:    dcl.StringOrNil(p.Obsolete),
		Deleted:     dcl.StringOrNil(p.Deleted),
	}
	return obj
}

// ProtoToImage converts a Image resource from its proto representation.
func ProtoToImage(p *betapb.ComputeBetaImage) *beta.Image {
	obj := &beta.Image{
		ArchiveSizeBytes:             dcl.Int64OrNil(p.ArchiveSizeBytes),
		Description:                  dcl.StringOrNil(p.Description),
		DiskSizeGb:                   dcl.Int64OrNil(p.DiskSizeGb),
		Family:                       dcl.StringOrNil(p.Family),
		ImageEncryptionKey:           ProtoToComputeBetaImageImageEncryptionKey(p.GetImageEncryptionKey()),
		Name:                         dcl.StringOrNil(p.Name),
		RawDisk:                      ProtoToComputeBetaImageRawDisk(p.GetRawDisk()),
		ShieldedInstanceInitialState: ProtoToComputeBetaImageShieldedInstanceInitialState(p.GetShieldedInstanceInitialState()),
		SelfLink:                     dcl.StringOrNil(p.SelfLink),
		SourceDisk:                   dcl.StringOrNil(p.SourceDisk),
		SourceDiskEncryptionKey:      ProtoToComputeBetaImageSourceDiskEncryptionKey(p.GetSourceDiskEncryptionKey()),
		SourceDiskId:                 dcl.StringOrNil(p.SourceDiskId),
		SourceImage:                  dcl.StringOrNil(p.SourceImage),
		SourceImageEncryptionKey:     ProtoToComputeBetaImageSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
		SourceImageId:                dcl.StringOrNil(p.SourceImageId),
		SourceSnapshot:               dcl.StringOrNil(p.SourceSnapshot),
		SourceSnapshotEncryptionKey:  ProtoToComputeBetaImageSourceSnapshotEncryptionKey(p.GetSourceSnapshotEncryptionKey()),
		SourceSnapshotId:             dcl.StringOrNil(p.SourceSnapshotId),
		SourceType:                   ProtoToComputeBetaImageSourceTypeEnum(p.GetSourceType()),
		Status:                       ProtoToComputeBetaImageStatusEnum(p.GetStatus()),
		Deprecated:                   ProtoToComputeBetaImageDeprecated(p.GetDeprecated()),
		Project:                      dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetGuestOsFeature() {
		obj.GuestOSFeature = append(obj.GuestOSFeature, *ProtoToComputeBetaImageGuestOSFeature(r))
	}
	for _, r := range p.GetLicense() {
		obj.License = append(obj.License, r)
	}
	for _, r := range p.GetStorageLocation() {
		obj.StorageLocation = append(obj.StorageLocation, r)
	}
	return obj
}

// ImageGuestOSFeatureTypeEnumToProto converts a ImageGuestOSFeatureTypeEnum enum to its proto representation.
func ComputeBetaImageGuestOSFeatureTypeEnumToProto(e *beta.ImageGuestOSFeatureTypeEnum) betapb.ComputeBetaImageGuestOSFeatureTypeEnum {
	if e == nil {
		return betapb.ComputeBetaImageGuestOSFeatureTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageGuestOSFeatureTypeEnum_value["ImageGuestOSFeatureTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageGuestOSFeatureTypeEnum(v)
	}
	return betapb.ComputeBetaImageGuestOSFeatureTypeEnum(0)
}

// ImageRawDiskContainerTypeEnumToProto converts a ImageRawDiskContainerTypeEnum enum to its proto representation.
func ComputeBetaImageRawDiskContainerTypeEnumToProto(e *beta.ImageRawDiskContainerTypeEnum) betapb.ComputeBetaImageRawDiskContainerTypeEnum {
	if e == nil {
		return betapb.ComputeBetaImageRawDiskContainerTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageRawDiskContainerTypeEnum_value["ImageRawDiskContainerTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageRawDiskContainerTypeEnum(v)
	}
	return betapb.ComputeBetaImageRawDiskContainerTypeEnum(0)
}

// ImageShieldedInstanceInitialStatePkFileTypeEnumToProto converts a ImageShieldedInstanceInitialStatePkFileTypeEnum enum to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnumToProto(e *beta.ImageShieldedInstanceInitialStatePkFileTypeEnum) betapb.ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum {
	if e == nil {
		return betapb.ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum_value["ImageShieldedInstanceInitialStatePkFileTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum(v)
	}
	return betapb.ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnum(0)
}

// ImageShieldedInstanceInitialStateKekFileTypeEnumToProto converts a ImageShieldedInstanceInitialStateKekFileTypeEnum enum to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnumToProto(e *beta.ImageShieldedInstanceInitialStateKekFileTypeEnum) betapb.ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum {
	if e == nil {
		return betapb.ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum_value["ImageShieldedInstanceInitialStateKekFileTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum(v)
	}
	return betapb.ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnum(0)
}

// ImageShieldedInstanceInitialStateDbFileTypeEnumToProto converts a ImageShieldedInstanceInitialStateDbFileTypeEnum enum to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnumToProto(e *beta.ImageShieldedInstanceInitialStateDbFileTypeEnum) betapb.ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum {
	if e == nil {
		return betapb.ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum_value["ImageShieldedInstanceInitialStateDbFileTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum(v)
	}
	return betapb.ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnum(0)
}

// ImageShieldedInstanceInitialStateDbxFileTypeEnumToProto converts a ImageShieldedInstanceInitialStateDbxFileTypeEnum enum to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnumToProto(e *beta.ImageShieldedInstanceInitialStateDbxFileTypeEnum) betapb.ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum {
	if e == nil {
		return betapb.ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum_value["ImageShieldedInstanceInitialStateDbxFileTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum(v)
	}
	return betapb.ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnum(0)
}

// ImageSourceTypeEnumToProto converts a ImageSourceTypeEnum enum to its proto representation.
func ComputeBetaImageSourceTypeEnumToProto(e *beta.ImageSourceTypeEnum) betapb.ComputeBetaImageSourceTypeEnum {
	if e == nil {
		return betapb.ComputeBetaImageSourceTypeEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageSourceTypeEnum_value["ImageSourceTypeEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageSourceTypeEnum(v)
	}
	return betapb.ComputeBetaImageSourceTypeEnum(0)
}

// ImageStatusEnumToProto converts a ImageStatusEnum enum to its proto representation.
func ComputeBetaImageStatusEnumToProto(e *beta.ImageStatusEnum) betapb.ComputeBetaImageStatusEnum {
	if e == nil {
		return betapb.ComputeBetaImageStatusEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageStatusEnum_value["ImageStatusEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageStatusEnum(v)
	}
	return betapb.ComputeBetaImageStatusEnum(0)
}

// ImageDeprecatedStateEnumToProto converts a ImageDeprecatedStateEnum enum to its proto representation.
func ComputeBetaImageDeprecatedStateEnumToProto(e *beta.ImageDeprecatedStateEnum) betapb.ComputeBetaImageDeprecatedStateEnum {
	if e == nil {
		return betapb.ComputeBetaImageDeprecatedStateEnum(0)
	}
	if v, ok := betapb.ComputeBetaImageDeprecatedStateEnum_value["ImageDeprecatedStateEnum"+string(*e)]; ok {
		return betapb.ComputeBetaImageDeprecatedStateEnum(v)
	}
	return betapb.ComputeBetaImageDeprecatedStateEnum(0)
}

// ImageGuestOSFeatureToProto converts a ImageGuestOSFeature resource to its proto representation.
func ComputeBetaImageGuestOSFeatureToProto(o *beta.ImageGuestOSFeature) *betapb.ComputeBetaImageGuestOSFeature {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageGuestOSFeature{
		Type: ComputeBetaImageGuestOSFeatureTypeEnumToProto(o.Type),
	}
	return p
}

// ImageImageEncryptionKeyToProto converts a ImageImageEncryptionKey resource to its proto representation.
func ComputeBetaImageImageEncryptionKeyToProto(o *beta.ImageImageEncryptionKey) *betapb.ComputeBetaImageImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageImageEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// ImageRawDiskToProto converts a ImageRawDisk resource to its proto representation.
func ComputeBetaImageRawDiskToProto(o *beta.ImageRawDisk) *betapb.ComputeBetaImageRawDisk {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageRawDisk{
		Source:        dcl.ValueOrEmptyString(o.Source),
		Sha1Checksum:  dcl.ValueOrEmptyString(o.Sha1Checksum),
		ContainerType: ComputeBetaImageRawDiskContainerTypeEnumToProto(o.ContainerType),
	}
	return p
}

// ImageShieldedInstanceInitialStateToProto converts a ImageShieldedInstanceInitialState resource to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStateToProto(o *beta.ImageShieldedInstanceInitialState) *betapb.ComputeBetaImageShieldedInstanceInitialState {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageShieldedInstanceInitialState{
		Pk: ComputeBetaImageShieldedInstanceInitialStatePkToProto(o.Pk),
	}
	for _, r := range o.Kek {
		p.Kek = append(p.Kek, ComputeBetaImageShieldedInstanceInitialStateKekToProto(&r))
	}
	for _, r := range o.Db {
		p.Db = append(p.Db, ComputeBetaImageShieldedInstanceInitialStateDbToProto(&r))
	}
	for _, r := range o.Dbx {
		p.Dbx = append(p.Dbx, ComputeBetaImageShieldedInstanceInitialStateDbxToProto(&r))
	}
	return p
}

// ImageShieldedInstanceInitialStatePkToProto converts a ImageShieldedInstanceInitialStatePk resource to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStatePkToProto(o *beta.ImageShieldedInstanceInitialStatePk) *betapb.ComputeBetaImageShieldedInstanceInitialStatePk {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageShieldedInstanceInitialStatePk{
		Content:  dcl.ValueOrEmptyString(o.Content),
		FileType: ComputeBetaImageShieldedInstanceInitialStatePkFileTypeEnumToProto(o.FileType),
	}
	return p
}

// ImageShieldedInstanceInitialStateKekToProto converts a ImageShieldedInstanceInitialStateKek resource to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStateKekToProto(o *beta.ImageShieldedInstanceInitialStateKek) *betapb.ComputeBetaImageShieldedInstanceInitialStateKek {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageShieldedInstanceInitialStateKek{
		Content:  dcl.ValueOrEmptyString(o.Content),
		FileType: ComputeBetaImageShieldedInstanceInitialStateKekFileTypeEnumToProto(o.FileType),
	}
	return p
}

// ImageShieldedInstanceInitialStateDbToProto converts a ImageShieldedInstanceInitialStateDb resource to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStateDbToProto(o *beta.ImageShieldedInstanceInitialStateDb) *betapb.ComputeBetaImageShieldedInstanceInitialStateDb {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageShieldedInstanceInitialStateDb{
		Content:  dcl.ValueOrEmptyString(o.Content),
		FileType: ComputeBetaImageShieldedInstanceInitialStateDbFileTypeEnumToProto(o.FileType),
	}
	return p
}

// ImageShieldedInstanceInitialStateDbxToProto converts a ImageShieldedInstanceInitialStateDbx resource to its proto representation.
func ComputeBetaImageShieldedInstanceInitialStateDbxToProto(o *beta.ImageShieldedInstanceInitialStateDbx) *betapb.ComputeBetaImageShieldedInstanceInitialStateDbx {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageShieldedInstanceInitialStateDbx{
		Content:  dcl.ValueOrEmptyString(o.Content),
		FileType: ComputeBetaImageShieldedInstanceInitialStateDbxFileTypeEnumToProto(o.FileType),
	}
	return p
}

// ImageSourceDiskEncryptionKeyToProto converts a ImageSourceDiskEncryptionKey resource to its proto representation.
func ComputeBetaImageSourceDiskEncryptionKeyToProto(o *beta.ImageSourceDiskEncryptionKey) *betapb.ComputeBetaImageSourceDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageSourceDiskEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// ImageSourceImageEncryptionKeyToProto converts a ImageSourceImageEncryptionKey resource to its proto representation.
func ComputeBetaImageSourceImageEncryptionKeyToProto(o *beta.ImageSourceImageEncryptionKey) *betapb.ComputeBetaImageSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageSourceImageEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// ImageSourceSnapshotEncryptionKeyToProto converts a ImageSourceSnapshotEncryptionKey resource to its proto representation.
func ComputeBetaImageSourceSnapshotEncryptionKeyToProto(o *beta.ImageSourceSnapshotEncryptionKey) *betapb.ComputeBetaImageSourceSnapshotEncryptionKey {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageSourceSnapshotEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// ImageDeprecatedToProto converts a ImageDeprecated resource to its proto representation.
func ComputeBetaImageDeprecatedToProto(o *beta.ImageDeprecated) *betapb.ComputeBetaImageDeprecated {
	if o == nil {
		return nil
	}
	p := &betapb.ComputeBetaImageDeprecated{
		State:       ComputeBetaImageDeprecatedStateEnumToProto(o.State),
		Replacement: dcl.ValueOrEmptyString(o.Replacement),
		Deprecated:  dcl.ValueOrEmptyString(o.Deprecated),
		Obsolete:    dcl.ValueOrEmptyString(o.Obsolete),
		Deleted:     dcl.ValueOrEmptyString(o.Deleted),
	}
	return p
}

// ImageToProto converts a Image resource to its proto representation.
func ImageToProto(resource *beta.Image) *betapb.ComputeBetaImage {
	p := &betapb.ComputeBetaImage{
		ArchiveSizeBytes:             dcl.ValueOrEmptyInt64(resource.ArchiveSizeBytes),
		Description:                  dcl.ValueOrEmptyString(resource.Description),
		DiskSizeGb:                   dcl.ValueOrEmptyInt64(resource.DiskSizeGb),
		Family:                       dcl.ValueOrEmptyString(resource.Family),
		ImageEncryptionKey:           ComputeBetaImageImageEncryptionKeyToProto(resource.ImageEncryptionKey),
		Name:                         dcl.ValueOrEmptyString(resource.Name),
		RawDisk:                      ComputeBetaImageRawDiskToProto(resource.RawDisk),
		ShieldedInstanceInitialState: ComputeBetaImageShieldedInstanceInitialStateToProto(resource.ShieldedInstanceInitialState),
		SelfLink:                     dcl.ValueOrEmptyString(resource.SelfLink),
		SourceDisk:                   dcl.ValueOrEmptyString(resource.SourceDisk),
		SourceDiskEncryptionKey:      ComputeBetaImageSourceDiskEncryptionKeyToProto(resource.SourceDiskEncryptionKey),
		SourceDiskId:                 dcl.ValueOrEmptyString(resource.SourceDiskId),
		SourceImage:                  dcl.ValueOrEmptyString(resource.SourceImage),
		SourceImageEncryptionKey:     ComputeBetaImageSourceImageEncryptionKeyToProto(resource.SourceImageEncryptionKey),
		SourceImageId:                dcl.ValueOrEmptyString(resource.SourceImageId),
		SourceSnapshot:               dcl.ValueOrEmptyString(resource.SourceSnapshot),
		SourceSnapshotEncryptionKey:  ComputeBetaImageSourceSnapshotEncryptionKeyToProto(resource.SourceSnapshotEncryptionKey),
		SourceSnapshotId:             dcl.ValueOrEmptyString(resource.SourceSnapshotId),
		SourceType:                   ComputeBetaImageSourceTypeEnumToProto(resource.SourceType),
		Status:                       ComputeBetaImageStatusEnumToProto(resource.Status),
		Deprecated:                   ComputeBetaImageDeprecatedToProto(resource.Deprecated),
		Project:                      dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.GuestOSFeature {
		p.GuestOsFeature = append(p.GuestOsFeature, ComputeBetaImageGuestOSFeatureToProto(&r))
	}
	for _, r := range resource.License {
		p.License = append(p.License, r)
	}
	for _, r := range resource.StorageLocation {
		p.StorageLocation = append(p.StorageLocation, r)
	}

	return p
}

// ApplyImage handles the gRPC request by passing it to the underlying Image Apply() method.
func (s *ImageServer) applyImage(ctx context.Context, c *beta.Client, request *betapb.ApplyComputeBetaImageRequest) (*betapb.ComputeBetaImage, error) {
	p := ProtoToImage(request.GetResource())
	res, err := c.ApplyImage(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ImageToProto(res)
	return r, nil
}

// ApplyImage handles the gRPC request by passing it to the underlying Image Apply() method.
func (s *ImageServer) ApplyComputeBetaImage(ctx context.Context, request *betapb.ApplyComputeBetaImageRequest) (*betapb.ComputeBetaImage, error) {
	cl, err := createConfigImage(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyImage(ctx, cl, request)
}

// DeleteImage handles the gRPC request by passing it to the underlying Image Delete() method.
func (s *ImageServer) DeleteComputeBetaImage(ctx context.Context, request *betapb.DeleteComputeBetaImageRequest) (*emptypb.Empty, error) {

	cl, err := createConfigImage(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteImage(ctx, ProtoToImage(request.GetResource()))

}

// ListComputeBetaImage handles the gRPC request by passing it to the underlying ImageList() method.
func (s *ImageServer) ListComputeBetaImage(ctx context.Context, request *betapb.ListComputeBetaImageRequest) (*betapb.ListComputeBetaImageResponse, error) {
	cl, err := createConfigImage(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListImage(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*betapb.ComputeBetaImage
	for _, r := range resources.Items {
		rp := ImageToProto(r)
		protos = append(protos, rp)
	}
	return &betapb.ListComputeBetaImageResponse{Items: protos}, nil
}

func createConfigImage(ctx context.Context, service_account_file string) (*beta.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return beta.NewClient(conf), nil
}
