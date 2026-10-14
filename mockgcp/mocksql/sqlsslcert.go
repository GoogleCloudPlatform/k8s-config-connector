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

package mocksql

import (
	"context"
	"fmt"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/cloud/sql/v1beta4"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type sqlSslCertsServer struct {
	*MockService
	pb.UnimplementedSqlSslCertsServiceServer
}

const (
	dummyClientCert = "-----BEGIN CERTIFICATE-----\n" +
		"MIIDhTCCAm2gAwIBAgIEQ5YU2jANBgkqhkiG9w0BAQsFADB8MS0wKwYDVQQuEyQ3\n" +
		"MjZlMmU5MC01NTAyLTQ5NjctYmJkZi02MTZiMzBhN2ViZTQxKDAmBgNVBAMTH0Ns\n" +
		"b3VkIFNRTCBDbGllbnQgQ0EgY2xpZW50LW5hbWUxFDASBgNVBAoTC0dvb2dsZSwg\n" +
		"SW5jMQswCQYDVQQGEwJVUzAeFw0yNjEwMTQwMDQ0MDJaFw0zNjEwMTEwMDQ1MDJa\n" +
		"MDkxFDASBgNVBAMTC2NsaWVudC1uYW1lMRQwEgYDVQQKEwtHb29nbGUsIEluYzEL\n" +
		"MAkGA1UEBhMCVVMwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCYudgZ\n" +
		"BsmIPVsLP5axLmzou+LsfH9gmlXOFBnl1oU2KiWfI94t+Zzlyfu7E7jVdZWwGt2g\n" +
		"ah9qtIq3LyVso4B8vbjzwnAlhz94wUrhckFk2bCfMjvBlJxl5GW7lu4lbpq6HV7D\n" +
		"sBQ6fXv8vHCKNNF4XniSUC/8ig6KAieiGD7bkQnCIvAScxptaFDgTobiPTkjKu+/\n" +
		"Tah2cuR/gl0P6d1gO6BKZbq+McbM33Y9eFu0gbQ9Ew3KHff9DFqzyQgREbP6G8QG\n" +
		"T2aT4kcfdFQXvq2xjzsczJJU6QxtQyrnHSyHIr+JMdsSB3+FMEVwGLVJpdz9Zr0I\n" +
		"+MyHw7XWBgStFkqdAgMBAAGjUjBQMAkGA1UdEwQCMAAwQwYDVR0RBDwwOoE4b3Zl\n" +
		"cnNlZXIta2NjLXRlc3RlckBjbnJtLWJhcm5pLTQuaWFtLmdzZXJ2aWNlYWNjb3Vu\n" +
		"dC5jb20wDQYJKoZIhvcNAQELBQADggEBALOErfUQV5Gn27pFr7T/1lKM9mj72flG\n" +
		"UDUaKXfyjwVHf/gdZo2F5Z0atA70aLQhfT5KP6kYdVWgs7WnYSpZnEKxq32Lndxp\n" +
		"eQ/VgOE57JnsPDLVo5dFyjbT2Y8TkS/vBLTcLHH35ocxKmKrHduBbh3/T5uu6Jlz\n" +
		"CWKO6nTFPm/ovGWW4BRAC8NvVNnQsz3csnTfbcurUPFYeDmrBPHzIo/pubpOoj+3\n" +
		"oP3T4EY4n2Ec20gRpv0Kx8J+AFPRR65Fk7TdrMMCczM0/x4MwXRzxB3Un9T+Zo18\n" +
		"3UwfUAscOjZYnxq5Joencwboh1fzdEC1NTxJzCxfFblXnM5dbACVvWY=\n" +
		"-----END CERTIFICATE-----"

	dummyPrivateKey = "-----BEGIN RSA PRIVATE KEY-----\n" +
		"MIIEpQIBAAKCAQEA1i7zmSGopddoujVyKjrKVXKK5LktTZdGsq5dFbihwr4Xrp/7\n" +
		"O5Sy4BH4n80bnmJUlLXRoOGrViMBCrGyBPgBfblE85Kv7esx8RQAFoFWev4Vuveo\n" +
		"6nGV98lPiwrhzy0jdK7qg+mcoJ2A1XIg5mMMXutNcTXzjmEZZunRkcerk8Rg2bX7\n" +
		"vM7GjSNFMU7oE6dseXCgeukm4HZSCb3Nh2+eGu6jrG/QiH/q6ytsVHSVI+LnlBCD\n" +
		"MY52Nrvw3O9UoHwVZ/shXzfTBzXXE1j1xFApRBcXfzAd3Rp+4VWr18r8tzU+LVoQ\n" +
		"YGAaYXxFLnQxXo+TJlaueiThENlLWkPiJrHLnQIDAQABAoIBAQDMo/WZlQBG3Cay\n" +
		"64fV83AI7jTozkkLvoMNC+3iaBMeN3P3I+HuDmhOEL2lKVq/HKJFp+bPuW50EWPY\n" +
		"bOlzN+Zs0kygEMJJJxQDjCF9XzxarVPj3OcmgTpRkqWOaupPgYhD3zAws080YuiK\n" +
		"h84Jcg+KzXWjunGn0vxrSPI0QDueJR2i03tEDBAtMZ0pvAsJ0gmXRdzGOc2uRzDm\n" +
		"fbS3y/JIufClO28OzjJ5AJkbc9XgRDeCDOFY2D375bCg2boPYmP7Iw0HVU3RQhcr\n" +
		"t+US27VQBRJF4cQ2CCyr0ZbdaPn41v+/A/qxF6ZPguyy+KoyQjCqK8iFArRQ48hJ\n" +
		"cR2pFx4hAoGBAP2uXIJAdAemrOunv2CWlUHI2iHj/kJ1AXRMpiT+eF0US9E6tipE\n" +
		"mL63HkUhiAs2nJnPi3RDxP+kAO2Z3anqjm1KCeGj+IYYZMavnkC8EVybv9lDwORy\n" +
		"e2O1bfRc/tGa341KmvXLbp8oVMIYIvKz2cZmHGJ4V4DTq8dTvmqoE4/VAoGBANgk\n" +
		"KWY5MJToZJJ5bV0mc2stmGt/IAZZPlKjVmKOjDyzqHRLAhsmbMyUhhgZtyj0dzSW\n" +
		"ILEeaEJknYRrOB48D6IqkB8VnFJyHUG8l+Za41adqRQNid0S5n50/+eYbjZpYCrA\n" +
		"SGmC2dhPZvRD6tOyEEJF5PZMvqxDcNRilc627HipAoGBAKzqrSQbyvtsIXKAZXLx\n" +
		"McwlnIp9XlLubo9Xr+iHjIPl0chMvN8S4wscxwVYVeNO1nABiI03pJCcugU7XFz2\n" +
		"BR952EJ2AnFlL0w/aR+3Eh6OC7eM927Amlrc0JZAzXESoE8vC3F/uWfDlgK3cRr+\n" +
		"fPM/pxl37i1iGzVDYAhTiQIBAoGAPW25nmXumsOZoc+E945wCywAP7z3mxZOEip9\n" +
		"6LDexnnBDJws0w6OqW4k1kCov6kLIBTy4aPkucniwrm+T0l+n/Y807jOntfz3LT+\n" +
		"7ucx6XIRlbNrVTuD6rjR6j52RFyaikvvyJz50PJwLkgHO3dGC6/VrPKO1mKsdJA4\n" +
		"R3HRr1ECgYEAobNQbQSLrSWZ1cozJbmNgRqqvxDNSEDi8LpXukOAw4pz1km7o3ob\n" +
		"hCy1ksfFzsp5glYqwZd/Bahk64u3mII+rKoYwYLrH2l2aFDmMbdTfQUycpQZyi3+\n" +
		"VtGS1PFoKx9fSFDNHhR5ZhfasQcuKHYfeFfO2/DoOxQkNCI1y4I2huo=\n" +
		"-----END RSA PRIVATE KEY-----"

	dummyServerCaCert = "-----BEGIN CERTIFICATE-----\n" +
		"MIIDcTCCAlmgAwIBAgIBADANBgkqhkiG9w0BAQsFADBwMS0wKwYDVQQuEyQ1Yzcw\n" +
		"MmI0ZS1jMzExLTRjZGQtYWVjOS00N2MzOWM3ZDRhODYxHDAaBgNVBAMTE0Nsb3Vk\n" +
		"IFNRTCBTZXJ2ZXIgQ0ExFDASBgNVBAoTC0dvb2dsZSwgSW5jMQswCQYDVQQGEwJV\n" +
		"UzAeFw0yNjEwMTQwMDQyMzBaFw0zNjEwMTEwMDQzMzBaMHAxLTArBgNVBC4TJDVj\n" +
		"NzAyYjRlLWMzMTEtNGNkZC1hZWM5LTQ3YzM5YzdkNGE4NjEcMBoGA1UEAxMTQ2xv\n" +
		"dWQgU1FMIFNlcnZlciBDQTEUMBIGA1UEChMLR29vZ2xlLCBJbmMxCzAJBgNVBAYT\n" +
		"AlVTMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2YOdMjrFrkDS9vfx\n" +
		"lC88iuxMuaqRaWkZL9iVRXg4MHQ2RnvmtbKjVTOzJQV+adnazBU2jBPGDUl4X2xb\n" +
		"PeTiKTWIAovBg48ilL2zUaHzvjlgphUe3/IRiqXkgMcqDcG3j8iCi61sfEFVx5lL\n" +
		"1cD5Ud9OIMuOZ+kbWT6t9PTN+9Y6LPOmjRGClCieWYc4Rtwh6cIrCHXKTOqJfXRi\n" +
		"RaFfGWaJJxQUfMfULxbBVpuoKKhEQBxvKmC6YzyGPSupSzG3Nhwyz6Amz4xCtOOn\n" +
		"pwq9+8pS0AEfpVEpMZIfm61wzJtwJ2qyiWvxE40KuO/gocJo0Nqw6Go+PL45L3k5\n" +
		"aga+fwIDAQABoxYwFDASBgNVHRMBAf8ECDAGAQH/AgEAMA0GCSqGSIb3DQEBCwUA\n" +
		"A4IBAQAHlNisEnlUqwkqQ9VOaHJ6YqhKMgf3re6JY//gCi9CIA+l9T5h3FGu29DN\n" +
		"Uq+Hw0ohMOZrBER4vbzk3CnemJ26RJfVnIYa8kKKLst25LoDKnqTIi4K2wG/clvN\n" +
		"TYpO2SxZXhIy1i0PeLbFii4Sh44CCvsggSWtO+d5NuCsdbjpTrwvjIJzX3K65IP/\n" +
		"yKIlWs8cDJdB4xCC/IWuvXlDZJ2cOpXw/o4tUfQU4OVXEm6zwNAHSextY3NJGVUY\n" +
		"47LkmvpoAKzt1+AUj3FSE3cC6Z5YGmrqXQJ/lAW2B1jRpXZVRx8P5UcQL/a1UX98\n" +
		"UkHyCUyp43q6jU/hR0L6JNS3esPj\n" +
		"-----END CERTIFICATE-----"
)

type SslCertName struct {
	Project      *projects.ProjectData
	InstanceName string
	Fingerprint  string
}

func (n *SslCertName) String() string {
	return "projects/" + n.Project.ID + "/instances/" + n.InstanceName + "/sslCerts/" + n.Fingerprint
}

func (s *MockService) buildSslCertName(projectID, instanceName, fingerprint string) (*SslCertName, error) {
	project, err := s.projects.GetProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	return &SslCertName{
		Project:      project,
		InstanceName: instanceName,
		Fingerprint:  fingerprint,
	}, nil
}

func (s *sqlSslCertsServer) Get(ctx context.Context, req *pb.SqlSslCertsGetRequest) (*pb.SslCert, error) {
	name, err := s.buildSslCertName(req.GetProject(), req.GetInstance(), req.GetSha1Fingerprint())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.SslCert{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The SSL certificate does not exist.")
		}
		return nil, err
	}

	return obj, nil
}

func (s *sqlSslCertsServer) List(ctx context.Context, req *pb.SqlSslCertsListRequest) (*pb.SslCertsListResponse, error) {
	instanceName, err := s.buildInstanceName(req.GetProject(), req.GetInstance())
	if err != nil {
		return nil, err
	}

	ret := &pb.SslCertsListResponse{
		Kind: "sql#sslCertsList",
	}

	certKind := (&pb.SslCert{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, certKind, storage.ListOptions{}, func(obj proto.Message) error {
		cert := obj.(*pb.SslCert)
		if cert.GetInstance() == instanceName.InstanceName {
			ret.Items = append(ret.Items, cert)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return ret, nil
}

func (s *sqlSslCertsServer) Insert(ctx context.Context, req *pb.SqlSslCertsInsertRequest) (*pb.SslCertsInsertResponse, error) {
	instanceName, err := s.buildInstanceName(req.GetProject(), req.GetInstance())
	if err != nil {
		return nil, err
	}

	instance := &pb.DatabaseInstance{}
	if err := s.storage.Get(ctx, instanceName.String(), instance); err != nil {
		return nil, err
	}

	commonName := ""
	if req.GetBody() != nil {
		commonName = req.GetBody().GetCommonName()
	}

	fingerprint := "760567766a057bb737d50ac1e5798bb9c039dea6"
	serialNumber := "1133909210"

	certName, err := s.buildSslCertName(req.GetProject(), req.GetInstance(), fingerprint)
	if err != nil {
		return nil, err
	}

	certObj := &pb.SslCert{
		Kind:             "sql#sslCert",
		CertSerialNumber: serialNumber,
		Cert:             dummyClientCert,
		CommonName:       commonName,
		CreateTime:       timestamppb.New(time.Date(2026, 10, 14, 0, 44, 2, 750000000, time.UTC)),
		ExpirationTime:   timestamppb.New(time.Date(2036, 10, 11, 0, 45, 2, 750000000, time.UTC)),
		Sha1Fingerprint:  fingerprint,
		Instance:         instanceName.InstanceName,
		SelfLink:         fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s/sslCerts/%s", instanceName.Project.ID, instanceName.InstanceName, fingerprint),
	}

	if err := s.storage.Create(ctx, certName.String(), certObj); err != nil {
		return nil, err
	}

	serverCaCert := proto.Clone(instance.ServerCaCert).(*pb.SslCert)
	if serverCaCert == nil {
		serverCaCert = &pb.SslCert{
			CertSerialNumber: "0",
			CommonName:       "common-name",
			CreateTime:       timestamppb.New(time.Date(2024, 4, 1, 12, 34, 56, 123456000, time.UTC)),
			ExpirationTime:   timestamppb.New(time.Date(2024, 4, 1, 12, 34, 56, 123456000, time.UTC)),
			Sha1Fingerprint:  "12345678",
			Instance:         instanceName.InstanceName,
			Kind:             "sql#sslCert",
		}
	}
	serverCaCert.Cert = dummyServerCaCert
	serverCaCert.SelfLink = fmt.Sprintf("https://sqladmin.googleapis.com/sql/v1beta4/projects/%s/instances/%s/sslCerts/4465455dc1e21f3addca06cd586e6f7a9c146140", instanceName.Project.ID, instanceName.InstanceName)

	op := &pb.Operation{
		TargetProject: instanceName.Project.ID,
		OperationType: pb.Operation_UPDATE,
		Status:        pb.Operation_PENDING,
	}

	lro, err := s.operations.startLRO(ctx, op, certObj, func() (proto.Message, error) {
		return certObj, nil
	})
	if err != nil {
		return nil, err
	}

	ret := &pb.SslCertsInsertResponse{
		Kind:         "sql#sslCertsInsert",
		Operation:    lro,
		ServerCaCert: serverCaCert,
		ClientCert: &pb.SslCertDetail{
			CertInfo:       certObj,
			CertPrivateKey: dummyPrivateKey,
		},
	}

	return ret, nil
}

func (s *sqlSslCertsServer) Delete(ctx context.Context, req *pb.SqlSslCertsDeleteRequest) (*pb.Operation, error) {
	name, err := s.buildSslCertName(req.GetProject(), req.GetInstance(), req.GetSha1Fingerprint())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	deleted := &pb.SslCert{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "The SSL certificate does not exist.")
		}
		return nil, err
	}

	op := &pb.Operation{
		TargetProject: name.Project.ID,
		OperationType: pb.Operation_UPDATE,
		Status:        pb.Operation_PENDING,
	}
	return s.operations.startLRO(ctx, op, deleted, func() (proto.Message, error) {
		return deleted, nil
	})
}
