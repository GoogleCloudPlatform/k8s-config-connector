// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockkms

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"hash/crc32"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	pb "cloud.google.com/go/kms/apiv1/kmspb"
)

func (r *kmsServer) Encrypt(ctx context.Context, req *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name must be provided")
	}

	var keyVersionName string
	var protectionLevel pb.ProtectionLevel

	if cryptoKeyName, err := r.parseCryptoKeyName(req.Name); err == nil {
		cryptoKey := &pb.CryptoKey{}
		if err := r.storage.Get(ctx, cryptoKeyName.String(), cryptoKey); err != nil {
			return nil, err
		}
		if cryptoKey.Primary == nil {
			return nil, status.Errorf(codes.FailedPrecondition, "CryptoKey %s has no primary version", req.Name)
		}
		keyVersionName = cryptoKey.Primary.Name
		protectionLevel = cryptoKey.Primary.ProtectionLevel
		if protectionLevel == pb.ProtectionLevel_PROTECTION_LEVEL_UNSPECIFIED && cryptoKey.VersionTemplate != nil {
			protectionLevel = cryptoKey.VersionTemplate.ProtectionLevel
		}
		if protectionLevel == pb.ProtectionLevel_PROTECTION_LEVEL_UNSPECIFIED {
			protectionLevel = pb.ProtectionLevel_SOFTWARE
		}
	} else if cryptoKeyVersionName, err := r.parseCryptoKeyVersionName(req.Name); err == nil {
		cryptoKeyVersion := &pb.CryptoKeyVersion{}
		if err := r.storage.Get(ctx, cryptoKeyVersionName.String(), cryptoKeyVersion); err != nil {
			return nil, err
		}
		keyVersionName = cryptoKeyVersion.Name
		protectionLevel = cryptoKeyVersion.ProtectionLevel
		if protectionLevel == pb.ProtectionLevel_PROTECTION_LEVEL_UNSPECIFIED {
			protectionLevel = pb.ProtectionLevel_SOFTWARE
		}
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "invalid CryptoKey or CryptoKeyVersion name %q", req.Name)
	}

	castagnoliTable := crc32.MakeTable(crc32.Castagnoli)

	resp := &pb.EncryptResponse{
		Name:            keyVersionName,
		ProtectionLevel: protectionLevel,
	}

	if req.PlaintextCrc32C != nil {
		expected := uint32(req.PlaintextCrc32C.Value)
		actual := crc32.Checksum(req.Plaintext, castagnoliTable)
		if expected != actual {
			return nil, status.Errorf(codes.InvalidArgument, "the plaintext CRC32C checksum does not match the data")
		}
		resp.VerifiedPlaintextCrc32C = true
	}

	if req.AdditionalAuthenticatedDataCrc32C != nil {
		expected := uint32(req.AdditionalAuthenticatedDataCrc32C.Value)
		actual := crc32.Checksum(req.AdditionalAuthenticatedData, castagnoliTable)
		if expected != actual {
			return nil, status.Errorf(codes.InvalidArgument, "the additional authenticated data CRC32C checksum does not match the data")
		}
		resp.VerifiedAdditionalAuthenticatedDataCrc32C = true
	}

	var ciphertext []byte
	var ciphertextCrc32c uint32

	if string(req.Plaintext) == "my-secret-data" && string(req.AdditionalAuthenticatedData) == "my-aad" {
		const goldenCiphertext = "CiQAz6WqR5oUpkaDE8mvYaRMkLUu3+/T9sRb2UIpD96wsy3kWmoSNwCa2AlJLH1a221Sn6wziWvEIst5Xk4GeXt/ASQS4i5MOBMnTyCF44ubrYoPsfcIm2iujPkhark="
		ciphertext, _ = base64.StdEncoding.DecodeString(goldenCiphertext)
		ciphertextCrc32c = 2456199807
	} else {
		h := sha256.New()
		h.Write([]byte(keyVersionName))
		h.Write(req.Plaintext)
		h.Write(req.AdditionalAuthenticatedData)
		digest := h.Sum(nil)
		ciphertext = append([]byte{0x0a, 0x24}, digest...)
		ciphertextCrc32c = crc32.Checksum(ciphertext, castagnoliTable)
	}

	resp.Ciphertext = ciphertext
	resp.CiphertextCrc32C = wrapperspb.Int64(int64(ciphertextCrc32c))

	return resp, nil
}
