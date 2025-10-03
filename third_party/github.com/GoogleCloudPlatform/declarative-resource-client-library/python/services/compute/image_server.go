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
	computepb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/compute/compute_go_proto"
	emptypb "github.com/GoogleCloudPlatform/declarative-resource-client-library/python/proto/empty_go_proto"
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute"
)

// Server implements the gRPC interface for Image.
type ImageServer struct{}

// ProtoToImageGuestOSFeatureTypeEnum converts a ImageGuestOSFeatureTypeEnum enum from its proto representation.
func ProtoToComputeImageGuestOSFeatureTypeEnum(e computepb.ComputeImageGuestOSFeatureTypeEnum) *compute.ImageGuestOSFeatureTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageGuestOSFeatureTypeEnum_name[int32(e)]; ok {
		e := compute.ImageGuestOSFeatureTypeEnum(n[len("ComputeImageGuestOSFeatureTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageRawDiskContainerTypeEnum converts a ImageRawDiskContainerTypeEnum enum from its proto representation.
func ProtoToComputeImageRawDiskContainerTypeEnum(e computepb.ComputeImageRawDiskContainerTypeEnum) *compute.ImageRawDiskContainerTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageRawDiskContainerTypeEnum_name[int32(e)]; ok {
		e := compute.ImageRawDiskContainerTypeEnum(n[len("ComputeImageRawDiskContainerTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageShieldedInstanceInitialStatePkFileTypeEnum converts a ImageShieldedInstanceInitialStatePkFileTypeEnum enum from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialStatePkFileTypeEnum(e computepb.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum) *compute.ImageShieldedInstanceInitialStatePkFileTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum_name[int32(e)]; ok {
		e := compute.ImageShieldedInstanceInitialStatePkFileTypeEnum(n[len("ComputeImageShieldedInstanceInitialStatePkFileTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageShieldedInstanceInitialStateKekFileTypeEnum converts a ImageShieldedInstanceInitialStateKekFileTypeEnum enum from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialStateKekFileTypeEnum(e computepb.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum) *compute.ImageShieldedInstanceInitialStateKekFileTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum_name[int32(e)]; ok {
		e := compute.ImageShieldedInstanceInitialStateKekFileTypeEnum(n[len("ComputeImageShieldedInstanceInitialStateKekFileTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageShieldedInstanceInitialStateDbFileTypeEnum converts a ImageShieldedInstanceInitialStateDbFileTypeEnum enum from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialStateDbFileTypeEnum(e computepb.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum) *compute.ImageShieldedInstanceInitialStateDbFileTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum_name[int32(e)]; ok {
		e := compute.ImageShieldedInstanceInitialStateDbFileTypeEnum(n[len("ComputeImageShieldedInstanceInitialStateDbFileTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageShieldedInstanceInitialStateDbxFileTypeEnum converts a ImageShieldedInstanceInitialStateDbxFileTypeEnum enum from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialStateDbxFileTypeEnum(e computepb.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum) *compute.ImageShieldedInstanceInitialStateDbxFileTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum_name[int32(e)]; ok {
		e := compute.ImageShieldedInstanceInitialStateDbxFileTypeEnum(n[len("ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageSourceTypeEnum converts a ImageSourceTypeEnum enum from its proto representation.
func ProtoToComputeImageSourceTypeEnum(e computepb.ComputeImageSourceTypeEnum) *compute.ImageSourceTypeEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageSourceTypeEnum_name[int32(e)]; ok {
		e := compute.ImageSourceTypeEnum(n[len("ComputeImageSourceTypeEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageStatusEnum converts a ImageStatusEnum enum from its proto representation.
func ProtoToComputeImageStatusEnum(e computepb.ComputeImageStatusEnum) *compute.ImageStatusEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageStatusEnum_name[int32(e)]; ok {
		e := compute.ImageStatusEnum(n[len("ComputeImageStatusEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageDeprecatedStateEnum converts a ImageDeprecatedStateEnum enum from its proto representation.
func ProtoToComputeImageDeprecatedStateEnum(e computepb.ComputeImageDeprecatedStateEnum) *compute.ImageDeprecatedStateEnum {
	if e == 0 {
		return nil
	}
	if n, ok := computepb.ComputeImageDeprecatedStateEnum_name[int32(e)]; ok {
		e := compute.ImageDeprecatedStateEnum(n[len("ComputeImageDeprecatedStateEnum"):])
		return &e
	}
	return nil
}

// ProtoToImageGuestOSFeature converts a ImageGuestOSFeature resource from its proto representation.
func ProtoToComputeImageGuestOSFeature(p *computepb.ComputeImageGuestOSFeature) *compute.ImageGuestOSFeature {
	if p == nil {
		return nil
	}
	obj := &compute.ImageGuestOSFeature{
		Type: ProtoToComputeImageGuestOSFeatureTypeEnum(p.GetType()),
	}
	return obj
}

// ProtoToImageImageEncryptionKey converts a ImageImageEncryptionKey resource from its proto representation.
func ProtoToComputeImageImageEncryptionKey(p *computepb.ComputeImageImageEncryptionKey) *compute.ImageImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.ImageImageEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToImageRawDisk converts a ImageRawDisk resource from its proto representation.
func ProtoToComputeImageRawDisk(p *computepb.ComputeImageRawDisk) *compute.ImageRawDisk {
	if p == nil {
		return nil
	}
	obj := &compute.ImageRawDisk{
		Source:        dcl.StringOrNil(p.Source),
		Sha1Checksum:  dcl.StringOrNil(p.Sha1Checksum),
		ContainerType: ProtoToComputeImageRawDiskContainerTypeEnum(p.GetContainerType()),
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialState converts a ImageShieldedInstanceInitialState resource from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialState(p *computepb.ComputeImageShieldedInstanceInitialState) *compute.ImageShieldedInstanceInitialState {
	if p == nil {
		return nil
	}
	obj := &compute.ImageShieldedInstanceInitialState{
		Pk: ProtoToComputeImageShieldedInstanceInitialStatePk(p.GetPk()),
	}
	for _, r := range p.GetKek() {
		obj.Kek = append(obj.Kek, *ProtoToComputeImageShieldedInstanceInitialStateKek(r))
	}
	for _, r := range p.GetDb() {
		obj.Db = append(obj.Db, *ProtoToComputeImageShieldedInstanceInitialStateDb(r))
	}
	for _, r := range p.GetDbx() {
		obj.Dbx = append(obj.Dbx, *ProtoToComputeImageShieldedInstanceInitialStateDbx(r))
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialStatePk converts a ImageShieldedInstanceInitialStatePk resource from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialStatePk(p *computepb.ComputeImageShieldedInstanceInitialStatePk) *compute.ImageShieldedInstanceInitialStatePk {
	if p == nil {
		return nil
	}
	obj := &compute.ImageShieldedInstanceInitialStatePk{
		Content:  dcl.StringOrNil(p.Content),
		FileType: ProtoToComputeImageShieldedInstanceInitialStatePkFileTypeEnum(p.GetFileType()),
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialStateKek converts a ImageShieldedInstanceInitialStateKek resource from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialStateKek(p *computepb.ComputeImageShieldedInstanceInitialStateKek) *compute.ImageShieldedInstanceInitialStateKek {
	if p == nil {
		return nil
	}
	obj := &compute.ImageShieldedInstanceInitialStateKek{
		Content:  dcl.StringOrNil(p.Content),
		FileType: ProtoToComputeImageShieldedInstanceInitialStateKekFileTypeEnum(p.GetFileType()),
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialStateDb converts a ImageShieldedInstanceInitialStateDb resource from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialStateDb(p *computepb.ComputeImageShieldedInstanceInitialStateDb) *compute.ImageShieldedInstanceInitialStateDb {
	if p == nil {
		return nil
	}
	obj := &compute.ImageShieldedInstanceInitialStateDb{
		Content:  dcl.StringOrNil(p.Content),
		FileType: ProtoToComputeImageShieldedInstanceInitialStateDbFileTypeEnum(p.GetFileType()),
	}
	return obj
}

// ProtoToImageShieldedInstanceInitialStateDbx converts a ImageShieldedInstanceInitialStateDbx resource from its proto representation.
func ProtoToComputeImageShieldedInstanceInitialStateDbx(p *computepb.ComputeImageShieldedInstanceInitialStateDbx) *compute.ImageShieldedInstanceInitialStateDbx {
	if p == nil {
		return nil
	}
	obj := &compute.ImageShieldedInstanceInitialStateDbx{
		Content:  dcl.StringOrNil(p.Content),
		FileType: ProtoToComputeImageShieldedInstanceInitialStateDbxFileTypeEnum(p.GetFileType()),
	}
	return obj
}

// ProtoToImageSourceDiskEncryptionKey converts a ImageSourceDiskEncryptionKey resource from its proto representation.
func ProtoToComputeImageSourceDiskEncryptionKey(p *computepb.ComputeImageSourceDiskEncryptionKey) *compute.ImageSourceDiskEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.ImageSourceDiskEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToImageSourceImageEncryptionKey converts a ImageSourceImageEncryptionKey resource from its proto representation.
func ProtoToComputeImageSourceImageEncryptionKey(p *computepb.ComputeImageSourceImageEncryptionKey) *compute.ImageSourceImageEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.ImageSourceImageEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToImageSourceSnapshotEncryptionKey converts a ImageSourceSnapshotEncryptionKey resource from its proto representation.
func ProtoToComputeImageSourceSnapshotEncryptionKey(p *computepb.ComputeImageSourceSnapshotEncryptionKey) *compute.ImageSourceSnapshotEncryptionKey {
	if p == nil {
		return nil
	}
	obj := &compute.ImageSourceSnapshotEncryptionKey{
		RawKey:               dcl.StringOrNil(p.RawKey),
		KmsKeyName:           dcl.StringOrNil(p.KmsKeyName),
		Sha256:               dcl.StringOrNil(p.Sha256),
		KmsKeyServiceAccount: dcl.StringOrNil(p.KmsKeyServiceAccount),
	}
	return obj
}

// ProtoToImageDeprecated converts a ImageDeprecated resource from its proto representation.
func ProtoToComputeImageDeprecated(p *computepb.ComputeImageDeprecated) *compute.ImageDeprecated {
	if p == nil {
		return nil
	}
	obj := &compute.ImageDeprecated{
		State:       ProtoToComputeImageDeprecatedStateEnum(p.GetState()),
		Replacement: dcl.StringOrNil(p.Replacement),
		Deprecated:  dcl.StringOrNil(p.Deprecated),
		Obsolete:    dcl.StringOrNil(p.Obsolete),
		Deleted:     dcl.StringOrNil(p.Deleted),
	}
	return obj
}

// ProtoToImage converts a Image resource from its proto representation.
func ProtoToImage(p *computepb.ComputeImage) *compute.Image {
	obj := &compute.Image{
		ArchiveSizeBytes:             dcl.Int64OrNil(p.ArchiveSizeBytes),
		Description:                  dcl.StringOrNil(p.Description),
		DiskSizeGb:                   dcl.Int64OrNil(p.DiskSizeGb),
		Family:                       dcl.StringOrNil(p.Family),
		ImageEncryptionKey:           ProtoToComputeImageImageEncryptionKey(p.GetImageEncryptionKey()),
		Name:                         dcl.StringOrNil(p.Name),
		RawDisk:                      ProtoToComputeImageRawDisk(p.GetRawDisk()),
		ShieldedInstanceInitialState: ProtoToComputeImageShieldedInstanceInitialState(p.GetShieldedInstanceInitialState()),
		SelfLink:                     dcl.StringOrNil(p.SelfLink),
		SourceDisk:                   dcl.StringOrNil(p.SourceDisk),
		SourceDiskEncryptionKey:      ProtoToComputeImageSourceDiskEncryptionKey(p.GetSourceDiskEncryptionKey()),
		SourceDiskId:                 dcl.StringOrNil(p.SourceDiskId),
		SourceImage:                  dcl.StringOrNil(p.SourceImage),
		SourceImageEncryptionKey:     ProtoToComputeImageSourceImageEncryptionKey(p.GetSourceImageEncryptionKey()),
		SourceImageId:                dcl.StringOrNil(p.SourceImageId),
		SourceSnapshot:               dcl.StringOrNil(p.SourceSnapshot),
		SourceSnapshotEncryptionKey:  ProtoToComputeImageSourceSnapshotEncryptionKey(p.GetSourceSnapshotEncryptionKey()),
		SourceSnapshotId:             dcl.StringOrNil(p.SourceSnapshotId),
		SourceType:                   ProtoToComputeImageSourceTypeEnum(p.GetSourceType()),
		Status:                       ProtoToComputeImageStatusEnum(p.GetStatus()),
		Deprecated:                   ProtoToComputeImageDeprecated(p.GetDeprecated()),
		Project:                      dcl.StringOrNil(p.Project),
	}
	for _, r := range p.GetGuestOsFeature() {
		obj.GuestOSFeature = append(obj.GuestOSFeature, *ProtoToComputeImageGuestOSFeature(r))
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
func ComputeImageGuestOSFeatureTypeEnumToProto(e *compute.ImageGuestOSFeatureTypeEnum) computepb.ComputeImageGuestOSFeatureTypeEnum {
	if e == nil {
		return computepb.ComputeImageGuestOSFeatureTypeEnum(0)
	}
	if v, ok := computepb.ComputeImageGuestOSFeatureTypeEnum_value["ImageGuestOSFeatureTypeEnum"+string(*e)]; ok {
		return computepb.ComputeImageGuestOSFeatureTypeEnum(v)
	}
	return computepb.ComputeImageGuestOSFeatureTypeEnum(0)
}

// ImageRawDiskContainerTypeEnumToProto converts a ImageRawDiskContainerTypeEnum enum to its proto representation.
func ComputeImageRawDiskContainerTypeEnumToProto(e *compute.ImageRawDiskContainerTypeEnum) computepb.ComputeImageRawDiskContainerTypeEnum {
	if e == nil {
		return computepb.ComputeImageRawDiskContainerTypeEnum(0)
	}
	if v, ok := computepb.ComputeImageRawDiskContainerTypeEnum_value["ImageRawDiskContainerTypeEnum"+string(*e)]; ok {
		return computepb.ComputeImageRawDiskContainerTypeEnum(v)
	}
	return computepb.ComputeImageRawDiskContainerTypeEnum(0)
}

// ImageShieldedInstanceInitialStatePkFileTypeEnumToProto converts a ImageShieldedInstanceInitialStatePkFileTypeEnum enum to its proto representation.
func ComputeImageShieldedInstanceInitialStatePkFileTypeEnumToProto(e *compute.ImageShieldedInstanceInitialStatePkFileTypeEnum) computepb.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum {
	if e == nil {
		return computepb.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum(0)
	}
	if v, ok := computepb.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum_value["ImageShieldedInstanceInitialStatePkFileTypeEnum"+string(*e)]; ok {
		return computepb.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum(v)
	}
	return computepb.ComputeImageShieldedInstanceInitialStatePkFileTypeEnum(0)
}

// ImageShieldedInstanceInitialStateKekFileTypeEnumToProto converts a ImageShieldedInstanceInitialStateKekFileTypeEnum enum to its proto representation.
func ComputeImageShieldedInstanceInitialStateKekFileTypeEnumToProto(e *compute.ImageShieldedInstanceInitialStateKekFileTypeEnum) computepb.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum {
	if e == nil {
		return computepb.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum(0)
	}
	if v, ok := computepb.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum_value["ImageShieldedInstanceInitialStateKekFileTypeEnum"+string(*e)]; ok {
		return computepb.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum(v)
	}
	return computepb.ComputeImageShieldedInstanceInitialStateKekFileTypeEnum(0)
}

// ImageShieldedInstanceInitialStateDbFileTypeEnumToProto converts a ImageShieldedInstanceInitialStateDbFileTypeEnum enum to its proto representation.
func ComputeImageShieldedInstanceInitialStateDbFileTypeEnumToProto(e *compute.ImageShieldedInstanceInitialStateDbFileTypeEnum) computepb.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum {
	if e == nil {
		return computepb.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum(0)
	}
	if v, ok := computepb.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum_value["ImageShieldedInstanceInitialStateDbFileTypeEnum"+string(*e)]; ok {
		return computepb.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum(v)
	}
	return computepb.ComputeImageShieldedInstanceInitialStateDbFileTypeEnum(0)
}

// ImageShieldedInstanceInitialStateDbxFileTypeEnumToProto converts a ImageShieldedInstanceInitialStateDbxFileTypeEnum enum to its proto representation.
func ComputeImageShieldedInstanceInitialStateDbxFileTypeEnumToProto(e *compute.ImageShieldedInstanceInitialStateDbxFileTypeEnum) computepb.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum {
	if e == nil {
		return computepb.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum(0)
	}
	if v, ok := computepb.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum_value["ImageShieldedInstanceInitialStateDbxFileTypeEnum"+string(*e)]; ok {
		return computepb.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum(v)
	}
	return computepb.ComputeImageShieldedInstanceInitialStateDbxFileTypeEnum(0)
}

// ImageSourceTypeEnumToProto converts a ImageSourceTypeEnum enum to its proto representation.
func ComputeImageSourceTypeEnumToProto(e *compute.ImageSourceTypeEnum) computepb.ComputeImageSourceTypeEnum {
	if e == nil {
		return computepb.ComputeImageSourceTypeEnum(0)
	}
	if v, ok := computepb.ComputeImageSourceTypeEnum_value["ImageSourceTypeEnum"+string(*e)]; ok {
		return computepb.ComputeImageSourceTypeEnum(v)
	}
	return computepb.ComputeImageSourceTypeEnum(0)
}

// ImageStatusEnumToProto converts a ImageStatusEnum enum to its proto representation.
func ComputeImageStatusEnumToProto(e *compute.ImageStatusEnum) computepb.ComputeImageStatusEnum {
	if e == nil {
		return computepb.ComputeImageStatusEnum(0)
	}
	if v, ok := computepb.ComputeImageStatusEnum_value["ImageStatusEnum"+string(*e)]; ok {
		return computepb.ComputeImageStatusEnum(v)
	}
	return computepb.ComputeImageStatusEnum(0)
}

// ImageDeprecatedStateEnumToProto converts a ImageDeprecatedStateEnum enum to its proto representation.
func ComputeImageDeprecatedStateEnumToProto(e *compute.ImageDeprecatedStateEnum) computepb.ComputeImageDeprecatedStateEnum {
	if e == nil {
		return computepb.ComputeImageDeprecatedStateEnum(0)
	}
	if v, ok := computepb.ComputeImageDeprecatedStateEnum_value["ImageDeprecatedStateEnum"+string(*e)]; ok {
		return computepb.ComputeImageDeprecatedStateEnum(v)
	}
	return computepb.ComputeImageDeprecatedStateEnum(0)
}

// ImageGuestOSFeatureToProto converts a ImageGuestOSFeature resource to its proto representation.
func ComputeImageGuestOSFeatureToProto(o *compute.ImageGuestOSFeature) *computepb.ComputeImageGuestOSFeature {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageGuestOSFeature{
		Type: ComputeImageGuestOSFeatureTypeEnumToProto(o.Type),
	}
	return p
}

// ImageImageEncryptionKeyToProto converts a ImageImageEncryptionKey resource to its proto representation.
func ComputeImageImageEncryptionKeyToProto(o *compute.ImageImageEncryptionKey) *computepb.ComputeImageImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageImageEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// ImageRawDiskToProto converts a ImageRawDisk resource to its proto representation.
func ComputeImageRawDiskToProto(o *compute.ImageRawDisk) *computepb.ComputeImageRawDisk {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageRawDisk{
		Source:        dcl.ValueOrEmptyString(o.Source),
		Sha1Checksum:  dcl.ValueOrEmptyString(o.Sha1Checksum),
		ContainerType: ComputeImageRawDiskContainerTypeEnumToProto(o.ContainerType),
	}
	return p
}

// ImageShieldedInstanceInitialStateToProto converts a ImageShieldedInstanceInitialState resource to its proto representation.
func ComputeImageShieldedInstanceInitialStateToProto(o *compute.ImageShieldedInstanceInitialState) *computepb.ComputeImageShieldedInstanceInitialState {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageShieldedInstanceInitialState{
		Pk: ComputeImageShieldedInstanceInitialStatePkToProto(o.Pk),
	}
	for _, r := range o.Kek {
		p.Kek = append(p.Kek, ComputeImageShieldedInstanceInitialStateKekToProto(&r))
	}
	for _, r := range o.Db {
		p.Db = append(p.Db, ComputeImageShieldedInstanceInitialStateDbToProto(&r))
	}
	for _, r := range o.Dbx {
		p.Dbx = append(p.Dbx, ComputeImageShieldedInstanceInitialStateDbxToProto(&r))
	}
	return p
}

// ImageShieldedInstanceInitialStatePkToProto converts a ImageShieldedInstanceInitialStatePk resource to its proto representation.
func ComputeImageShieldedInstanceInitialStatePkToProto(o *compute.ImageShieldedInstanceInitialStatePk) *computepb.ComputeImageShieldedInstanceInitialStatePk {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageShieldedInstanceInitialStatePk{
		Content:  dcl.ValueOrEmptyString(o.Content),
		FileType: ComputeImageShieldedInstanceInitialStatePkFileTypeEnumToProto(o.FileType),
	}
	return p
}

// ImageShieldedInstanceInitialStateKekToProto converts a ImageShieldedInstanceInitialStateKek resource to its proto representation.
func ComputeImageShieldedInstanceInitialStateKekToProto(o *compute.ImageShieldedInstanceInitialStateKek) *computepb.ComputeImageShieldedInstanceInitialStateKek {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageShieldedInstanceInitialStateKek{
		Content:  dcl.ValueOrEmptyString(o.Content),
		FileType: ComputeImageShieldedInstanceInitialStateKekFileTypeEnumToProto(o.FileType),
	}
	return p
}

// ImageShieldedInstanceInitialStateDbToProto converts a ImageShieldedInstanceInitialStateDb resource to its proto representation.
func ComputeImageShieldedInstanceInitialStateDbToProto(o *compute.ImageShieldedInstanceInitialStateDb) *computepb.ComputeImageShieldedInstanceInitialStateDb {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageShieldedInstanceInitialStateDb{
		Content:  dcl.ValueOrEmptyString(o.Content),
		FileType: ComputeImageShieldedInstanceInitialStateDbFileTypeEnumToProto(o.FileType),
	}
	return p
}

// ImageShieldedInstanceInitialStateDbxToProto converts a ImageShieldedInstanceInitialStateDbx resource to its proto representation.
func ComputeImageShieldedInstanceInitialStateDbxToProto(o *compute.ImageShieldedInstanceInitialStateDbx) *computepb.ComputeImageShieldedInstanceInitialStateDbx {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageShieldedInstanceInitialStateDbx{
		Content:  dcl.ValueOrEmptyString(o.Content),
		FileType: ComputeImageShieldedInstanceInitialStateDbxFileTypeEnumToProto(o.FileType),
	}
	return p
}

// ImageSourceDiskEncryptionKeyToProto converts a ImageSourceDiskEncryptionKey resource to its proto representation.
func ComputeImageSourceDiskEncryptionKeyToProto(o *compute.ImageSourceDiskEncryptionKey) *computepb.ComputeImageSourceDiskEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageSourceDiskEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// ImageSourceImageEncryptionKeyToProto converts a ImageSourceImageEncryptionKey resource to its proto representation.
func ComputeImageSourceImageEncryptionKeyToProto(o *compute.ImageSourceImageEncryptionKey) *computepb.ComputeImageSourceImageEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageSourceImageEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// ImageSourceSnapshotEncryptionKeyToProto converts a ImageSourceSnapshotEncryptionKey resource to its proto representation.
func ComputeImageSourceSnapshotEncryptionKeyToProto(o *compute.ImageSourceSnapshotEncryptionKey) *computepb.ComputeImageSourceSnapshotEncryptionKey {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageSourceSnapshotEncryptionKey{
		RawKey:               dcl.ValueOrEmptyString(o.RawKey),
		KmsKeyName:           dcl.ValueOrEmptyString(o.KmsKeyName),
		Sha256:               dcl.ValueOrEmptyString(o.Sha256),
		KmsKeyServiceAccount: dcl.ValueOrEmptyString(o.KmsKeyServiceAccount),
	}
	return p
}

// ImageDeprecatedToProto converts a ImageDeprecated resource to its proto representation.
func ComputeImageDeprecatedToProto(o *compute.ImageDeprecated) *computepb.ComputeImageDeprecated {
	if o == nil {
		return nil
	}
	p := &computepb.ComputeImageDeprecated{
		State:       ComputeImageDeprecatedStateEnumToProto(o.State),
		Replacement: dcl.ValueOrEmptyString(o.Replacement),
		Deprecated:  dcl.ValueOrEmptyString(o.Deprecated),
		Obsolete:    dcl.ValueOrEmptyString(o.Obsolete),
		Deleted:     dcl.ValueOrEmptyString(o.Deleted),
	}
	return p
}

// ImageToProto converts a Image resource to its proto representation.
func ImageToProto(resource *compute.Image) *computepb.ComputeImage {
	p := &computepb.ComputeImage{
		ArchiveSizeBytes:             dcl.ValueOrEmptyInt64(resource.ArchiveSizeBytes),
		Description:                  dcl.ValueOrEmptyString(resource.Description),
		DiskSizeGb:                   dcl.ValueOrEmptyInt64(resource.DiskSizeGb),
		Family:                       dcl.ValueOrEmptyString(resource.Family),
		ImageEncryptionKey:           ComputeImageImageEncryptionKeyToProto(resource.ImageEncryptionKey),
		Name:                         dcl.ValueOrEmptyString(resource.Name),
		RawDisk:                      ComputeImageRawDiskToProto(resource.RawDisk),
		ShieldedInstanceInitialState: ComputeImageShieldedInstanceInitialStateToProto(resource.ShieldedInstanceInitialState),
		SelfLink:                     dcl.ValueOrEmptyString(resource.SelfLink),
		SourceDisk:                   dcl.ValueOrEmptyString(resource.SourceDisk),
		SourceDiskEncryptionKey:      ComputeImageSourceDiskEncryptionKeyToProto(resource.SourceDiskEncryptionKey),
		SourceDiskId:                 dcl.ValueOrEmptyString(resource.SourceDiskId),
		SourceImage:                  dcl.ValueOrEmptyString(resource.SourceImage),
		SourceImageEncryptionKey:     ComputeImageSourceImageEncryptionKeyToProto(resource.SourceImageEncryptionKey),
		SourceImageId:                dcl.ValueOrEmptyString(resource.SourceImageId),
		SourceSnapshot:               dcl.ValueOrEmptyString(resource.SourceSnapshot),
		SourceSnapshotEncryptionKey:  ComputeImageSourceSnapshotEncryptionKeyToProto(resource.SourceSnapshotEncryptionKey),
		SourceSnapshotId:             dcl.ValueOrEmptyString(resource.SourceSnapshotId),
		SourceType:                   ComputeImageSourceTypeEnumToProto(resource.SourceType),
		Status:                       ComputeImageStatusEnumToProto(resource.Status),
		Deprecated:                   ComputeImageDeprecatedToProto(resource.Deprecated),
		Project:                      dcl.ValueOrEmptyString(resource.Project),
	}
	for _, r := range resource.GuestOSFeature {
		p.GuestOsFeature = append(p.GuestOsFeature, ComputeImageGuestOSFeatureToProto(&r))
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
func (s *ImageServer) applyImage(ctx context.Context, c *compute.Client, request *computepb.ApplyComputeImageRequest) (*computepb.ComputeImage, error) {
	p := ProtoToImage(request.GetResource())
	res, err := c.ApplyImage(ctx, p)
	if err != nil {
		return nil, err
	}
	r := ImageToProto(res)
	return r, nil
}

// ApplyImage handles the gRPC request by passing it to the underlying Image Apply() method.
func (s *ImageServer) ApplyComputeImage(ctx context.Context, request *computepb.ApplyComputeImageRequest) (*computepb.ComputeImage, error) {
	cl, err := createConfigImage(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return s.applyImage(ctx, cl, request)
}

// DeleteImage handles the gRPC request by passing it to the underlying Image Delete() method.
func (s *ImageServer) DeleteComputeImage(ctx context.Context, request *computepb.DeleteComputeImageRequest) (*emptypb.Empty, error) {

	cl, err := createConfigImage(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, cl.DeleteImage(ctx, ProtoToImage(request.GetResource()))

}

// ListComputeImage handles the gRPC request by passing it to the underlying ImageList() method.
func (s *ImageServer) ListComputeImage(ctx context.Context, request *computepb.ListComputeImageRequest) (*computepb.ListComputeImageResponse, error) {
	cl, err := createConfigImage(ctx, request.ServiceAccountFile)
	if err != nil {
		return nil, err
	}

	resources, err := cl.ListImage(ctx, request.Project)
	if err != nil {
		return nil, err
	}
	var protos []*computepb.ComputeImage
	for _, r := range resources.Items {
		rp := ImageToProto(r)
		protos = append(protos, rp)
	}
	return &computepb.ListComputeImageResponse{Items: protos}, nil
}

func createConfigImage(ctx context.Context, service_account_file string) (*compute.Client, error) {

	conf := dcl.NewConfig(dcl.WithUserAgent("dcl-test"), dcl.WithCredentialsFile(service_account_file))
	return compute.NewClient(conf), nil
}
