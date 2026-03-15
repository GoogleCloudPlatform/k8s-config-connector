// Copyright 2022 Google LLC
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

package mockprivateca

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/security/privateca/v1"
)

func (s *PrivateCAV1) GetCertificateAuthority(ctx context.Context, req *pb.GetCertificateAuthorityRequest) (*pb.CertificateAuthority, error) {
	name, err := s.parseCertificateAuthorityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.CertificateAuthority{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Resource '%s' was not found", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *PrivateCAV1) CreateCertificateAuthority(ctx context.Context, req *pb.CreateCertificateAuthorityRequest) (*longrunning.Operation, error) {
	reqName := req.Parent + "/certificateAuthorities/" + req.CertificateAuthorityId
	name, err := s.parseCertificateAuthorityName(reqName)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	fqn := name.String()

	obj := proto.Clone(req.CertificateAuthority).(*pb.CertificateAuthority)
	obj.Name = fqn

	obj.CreateTime = timestamppb.New(now)
	obj.UpdateTime = timestamppb.New(now)

	// Populate Output-only fields
	obj.PemCaCertificates = []string{"-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n"}
	caDesc := &pb.CertificateDescription{
		CertFingerprint: &pb.CertificateDescription_CertificateFingerprint{
			Sha256Hash: fmt.Sprintf("0123456789abcdef0123456789abcdef0123456789abcdef0123456789%s", name.CertificateAuthorityID),
		},
		AuthorityKeyId: &pb.CertificateDescription_KeyId{
			KeyId: "58ff0120decc0d87caa30eb45fef39e38133e733",
		},
		SubjectKeyId: &pb.CertificateDescription_KeyId{
			KeyId: "58ff0120decc0d87caa30eb45fef39e38133e733",
		},
	}
	if obj.Config != nil && obj.Config.SubjectConfig != nil {
		caDesc.SubjectDescription = &pb.CertificateDescription_SubjectDescription{
			Subject:         proto.Clone(obj.Config.SubjectConfig.Subject).(*pb.Subject),
			SubjectAltName:  proto.Clone(obj.Config.SubjectConfig.SubjectAltName).(*pb.SubjectAltNames),
			HexSerialNumber: "0123456789abcdef",
			Lifetime:        obj.Lifetime,
			NotBeforeTime:   obj.CreateTime,
			NotAfterTime:    timestamppb.New(now.Add(time.Duration(obj.Lifetime.Seconds) * time.Second)),
		}
	}
	if obj.Config != nil {
		caDesc.X509Description = proto.Clone(obj.Config.X509Config).(*pb.X509Parameters)
		decodedKey, _ := base64.StdEncoding.DecodeString("LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQ0lqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FnOEFNSUlDQ2dLQ0FnRUFyeGt3dVBoREZTNlc1eEgvMW1MVApTYXhzMTNONnpGYlRXSUY4aG4vTlk1cXlJYmFpd0FYdEVOeU53NkhSSmd4R2c0WElkRXlMVGpCK1VNVVVLYU9OCnd0WFJTUW9CeGR2VitKeWRlL05jTUUzM3QyT3d1UDBRVzY2UnRLak52b2R5dzRLTHphVmp6T1hPY0YwV0NNYy8KRjV4cU5uUHpOalVncUlBanozSHkrejROWmNnT0lnM3dVdEJRRlNTUm16TmtzdjMwejhLQjdXUXJkaElWb3JuOQpuWVFPeW11eTZPT0dLM3FIUXRZYk1MYU9oREY1VnJ3amozblUxeStQMzdRR1kwdG5KS3VYbGNwR3hJN0tkWTI5CjRJM1F6K0JHemI0Wll0WE1uTzNOZFZVaTRteG14VTBMbVE3VlJkdHpkTGN6cTJDUEoyV2JjTTZkUldmVERoNisKYzdDQys3K1ZBdzlxeW1OSnFXN0wxb2JNSkNuTHpwcHBPVWg0RjNXU1V0SXVLcUZ2alo1eFMzUXBFSGJsRFoybgpUaTY1Tm4yR2JzeVYrc2FxTkpOdUdmVEpvUzRJOGFoakljY2hDaXUzRDMxNm5MczlENmMwckRLNmxsbUpHdFRLCmZsSVEyQkZmY2FtV2VlSXlNU1dIK3Uza2lKTTY2YWF1NVJrNlFXWWlKMTRhaWswNmsvK1pMRnNMazUvbVV3eVEKTWFUMXNPUGQ1Z3FSeUsrdGgyVXNkb1p1dE5PZFZYMWdRNjk4cXdVZk1oTmpkSzNLZUkzNWQvY2xuR2F1UGVPSgo2ZlgySy9WN1hTblQrcGRjRTExZjNFU0FZVEIybnJJSXgzK3NjYXdZalREd0Qrd2JQZ0p4ZG4wc1ppOTNtV2FKCkFPdTV4QjBOSStsYXhZT2tPZHhrYklVQ0F3RUFBUT09Ci0tLS0tRU5EIFBVQkxJQyBLRVktLS0tLQo=")
		caDesc.PublicKey = &pb.PublicKey{
			Key:    decodedKey,
			Format: pb.PublicKey_PEM,
		}
	}
	obj.CaCertificateDescriptions = []*pb.CertificateDescription{caDesc}

	// Fetch CAPool to check publishing options
	caPool := &pb.CaPool{}
	caPoolName := &caPoolName{
		Project:    name.Project,
		Location:   name.Location,
		CAPoolName: name.CAPoolName,
	}
	if err := s.storage.Get(ctx, caPoolName.String(), caPool); err != nil {
		return nil, err
	}

	obj.Tier = caPool.Tier
	obj.State = pb.CertificateAuthority_STAGED

	if caPool.GetPublishingOptions().GetPublishCrl() {
		obj.AccessUrls = &pb.CertificateAuthority_AccessUrls{
			CrlAccessUrls: []string{
				fmt.Sprintf("http://privateca-content-00000000-0000-0000-0000-000000000000.storage.googleapis.com/%s/crl", name.CertificateAuthorityID),
			},
		}
	}

	// service seems to remove "zero" values
	pruneKU := func(ku *pb.KeyUsage) {
		if ku != nil && proto.Equal(ku.ExtendedKeyUsage, &pb.KeyUsage_ExtendedKeyUsageOptions{}) {
			ku.ExtendedKeyUsage = nil
		}
	}
	pruneKU(obj.GetConfig().GetX509Config().GetKeyUsage())
	for _, caDesc := range obj.CaCertificateDescriptions {
		pruneKU(caDesc.GetX509Description().GetKeyUsage())
	}

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "create",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *PrivateCAV1) DeleteCertificateAuthority(ctx context.Context, req *pb.DeleteCertificateAuthorityRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateAuthorityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	oldObj := &pb.CertificateAuthority{}
	if err := s.storage.Delete(ctx, fqn, oldObj); err != nil {
		return nil, err
	}

	// gcloud expects these to be set on the response
	oldObj.State = pb.CertificateAuthority_DELETED
	oldObj.DeleteTime = timestamppb.New(now)
	oldObj.ExpireTime = timestamppb.New(now.Add(30 * 24 * time.Hour))

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "delete",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return oldObj, nil
	})
}

func (s *PrivateCAV1) EnableCertificateAuthority(ctx context.Context, req *pb.EnableCertificateAuthorityRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateAuthorityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := &pb.CertificateAuthority{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.CertificateAuthority_ENABLED

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "enable",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *PrivateCAV1) DisableCertificateAuthority(ctx context.Context, req *pb.DisableCertificateAuthorityRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateAuthorityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := &pb.CertificateAuthority{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.CertificateAuthority_DISABLED

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "disable",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *PrivateCAV1) UpdateCertificateAuthority(ctx context.Context, req *pb.UpdateCertificateAuthorityRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateAuthorityName(req.GetCertificateAuthority().GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := &pb.CertificateAuthority{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	// TODO: FieldMask
	proto.Merge(obj, req.GetCertificateAuthority())

	// service seems to remove "zero" values
	pruneKU := func(ku *pb.KeyUsage) {
		if ku != nil && proto.Equal(ku.ExtendedKeyUsage, &pb.KeyUsage_ExtendedKeyUsageOptions{}) {
			ku.ExtendedKeyUsage = nil
		}
	}
	pruneKU(obj.GetConfig().GetX509Config().GetKeyUsage())
	for _, caDesc := range obj.CaCertificateDescriptions {
		pruneKU(caDesc.GetX509Description().GetKeyUsage())
	}

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "update",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

func (s *PrivateCAV1) ActivateCertificateAuthority(ctx context.Context, req *pb.ActivateCertificateAuthorityRequest) (*longrunning.Operation, error) {
	name, err := s.parseCertificateAuthorityName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := &pb.CertificateAuthority{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	obj.State = pb.CertificateAuthority_ENABLED

	if err := s.storage.Update(ctx, fqn, obj); err != nil {
		return nil, err
	}

	opMetadata := &pb.OperationMetadata{
		ApiVersion:            "v1",
		CreateTime:            timestamppb.New(now),
		Verb:                  "activate",
		RequestedCancellation: false,
		Target:                fqn,
	}
	opPrefix := fmt.Sprintf("projects/%s/locations/%s", name.Project.ID, name.Location)
	return s.operations.StartLRO(ctx, opPrefix, opMetadata, func() (proto.Message, error) {
		opMetadata.EndTime = timestamppb.Now()
		return obj, nil
	})
}

type certificateAuthorityName struct {
	caPoolName
	CertificateAuthorityID string
}

func (n *certificateAuthorityName) String() string {
	return "projects/" + n.Project.ID + "/locations/" + n.Location + "/caPools/" + n.CAPoolName + "/certificateAuthorities/" + n.CertificateAuthorityID
}

// parseCertificateAuthorityName parses a string into a certificateAuthorityName.
// The expected form is projects/<projectID>/locations/<region>/caPools/<capoolName>/certificateAuthorities/<caName>
func (s *MockService) parseCertificateAuthorityName(name string) (*certificateAuthorityName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 8 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "caPools" && tokens[6] == "certificateAuthorities" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &certificateAuthorityName{
			caPoolName: caPoolName{
				Project:    project,
				Location:   tokens[3],
				CAPoolName: tokens[5],
			},
			CertificateAuthorityID: tokens[7],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
