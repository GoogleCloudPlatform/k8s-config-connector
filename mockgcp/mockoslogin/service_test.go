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

package mockoslogin_test

import (
	"context"
	"net"
	"testing"

	pb "cloud.google.com/go/oslogin/apiv1/osloginpb"
	commonpb "cloud.google.com/go/oslogin/common/commonpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/mockoslogin"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestServiceCreation(t *testing.T) {
	env := &common.MockEnvironment{}
	st := storage.NewInMemoryStorage()
	svc := mockoslogin.New(env, st)
	if svc == nil {
		t.Fatalf("expected New() to return a non-nil MockService")
	}
	expectedHosts := svc.ExpectedHosts()
	if len(expectedHosts) == 0 || expectedHosts[0] != "oslogin.googleapis.com" {
		t.Errorf("unexpected expected hosts: %v", expectedHosts)
	}

	grpcServer := grpc.NewServer()
	svc.Register(grpcServer)
}

func TestOSLoginOperations(t *testing.T) {
	ctx := context.Background()
	env := &common.MockEnvironment{}
	st := storage.NewInMemoryStorage()
	svc := mockoslogin.New(env, st)

	grpcServer := grpc.NewServer()
	svc.Register(grpcServer)

	const (
		parent = "users/test-user@example.com"
		key    = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIPM4pxpbPpjuBocS6qlW0BHRYgH5xmv/yVrANZR9lc1N test-user@example.com"
	)

	lis := bufconn.Listen(1024 * 1024)
	defer lis.Close()

	go func() {
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			t.Errorf("grpcServer.Serve error: %v", err)
		}
	}()
	defer grpcServer.Stop()

	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewOsLoginServiceClient(conn)

	// 1. Import SSH public key
	importResp, err := client.ImportSshPublicKey(ctx, &pb.ImportSshPublicKeyRequest{
		Parent: parent,
		SshPublicKey: &commonpb.SshPublicKey{
			Key: key,
		},
	})
	if err != nil {
		t.Fatalf("ImportSshPublicKey failed: %v", err)
	}
	if importResp.LoginProfile == nil {
		t.Fatalf("expected non-nil LoginProfile")
	}
	if len(importResp.LoginProfile.SshPublicKeys) != 1 {
		t.Fatalf("expected 1 SSH public key, got %d", len(importResp.LoginProfile.SshPublicKeys))
	}

	var fingerprint string
	for fp := range importResp.LoginProfile.SshPublicKeys {
		fingerprint = fp
	}
	name := parent + "/sshPublicKeys/" + fingerprint

	// 2. Get SSH public key
	getResp, err := client.GetSshPublicKey(ctx, &pb.GetSshPublicKeyRequest{
		Name: name,
	})
	if err != nil {
		t.Fatalf("GetSshPublicKey failed: %v", err)
	}
	if getResp.Fingerprint != fingerprint {
		t.Errorf("expected fingerprint %q, got %q", fingerprint, getResp.Fingerprint)
	}
	if getResp.Key != key {
		t.Errorf("expected key %q, got %q", key, getResp.Key)
	}

	// 3. Update SSH public key
	const newExpiry = int64(1893456000000000)
	updateResp, err := client.UpdateSshPublicKey(ctx, &pb.UpdateSshPublicKeyRequest{
		Name: name,
		SshPublicKey: &commonpb.SshPublicKey{
			ExpirationTimeUsec: newExpiry,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"expiration_time_usec"},
		},
	})
	if err != nil {
		t.Fatalf("UpdateSshPublicKey failed: %v", err)
	}
	if updateResp.ExpirationTimeUsec != newExpiry {
		t.Errorf("expected expiration %d, got %d", newExpiry, updateResp.ExpirationTimeUsec)
	}

	// 4. Get Login Profile
	profileResp, err := client.GetLoginProfile(ctx, &pb.GetLoginProfileRequest{
		Name: parent,
	})
	if err != nil {
		t.Fatalf("GetLoginProfile failed: %v", err)
	}
	if len(profileResp.SshPublicKeys) != 1 {
		t.Errorf("expected 1 SSH key in profile, got %d", len(profileResp.SshPublicKeys))
	}

	// 5. Delete SSH public key
	_, err = client.DeleteSshPublicKey(ctx, &pb.DeleteSshPublicKeyRequest{
		Name: name,
	})
	if err != nil {
		t.Fatalf("DeleteSshPublicKey failed: %v", err)
	}

	// 6. Verify deleted
	_, err = client.GetSshPublicKey(ctx, &pb.GetSshPublicKeyRequest{
		Name: name,
	})
	if err == nil {
		t.Fatalf("expected NotFound error after delete, got nil")
	}

	// 7. Delete Posix Account
	_, err = client.DeletePosixAccount(ctx, &pb.DeletePosixAccountRequest{
		Name: parent + "/projects/test-project",
	})
	if err != nil {
		t.Fatalf("DeletePosixAccount failed: %v", err)
	}
}
