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

package secretmanager

import (
	"testing"
	"time"

	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestSecretUpdatePathsTTL(t *testing.T) {
	expires := timestamppb.New(time.Date(2035, 10, 2, 15, 1, 23, 0, time.UTC))
	actual := &secretmanagerpb.Secret{
		Expiration: &secretmanagerpb.Secret_ExpireTime{ExpireTime: expires},
	}

	t.Run("no-op TTL reconciliation", func(t *testing.T) {
		resource := &secretmanagerpb.Secret{
			Expiration: &secretmanagerpb.Secret_Ttl{Ttl: durationpb.New(100 * time.Second)},
		}
		got, err := secretUpdatePaths(resource, actual)
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != 0 {
			t.Fatalf("update mask = %v, want empty", got)
		}
	})

	t.Run("retains explicit expiration changes", func(t *testing.T) {
		resource := &secretmanagerpb.Secret{
			Expiration: &secretmanagerpb.Secret_ExpireTime{
				ExpireTime: timestamppb.New(expires.AsTime().Add(time.Hour)),
			},
		}
		got, err := secretUpdatePaths(resource, actual)
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != 1 || !got.Has("expire_time") {
			t.Fatalf("update mask = %v, want only expire_time", got)
		}
	})

	t.Run("retains unrelated changes", func(t *testing.T) {
		resource := &secretmanagerpb.Secret{
			Expiration: &secretmanagerpb.Secret_Ttl{Ttl: durationpb.New(100 * time.Second)},
			Labels:     map[string]string{"new": "value"},
		}
		got, err := secretUpdatePaths(resource, actual)
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != 1 || !got.Has("labels") {
			t.Fatalf("update mask = %v, want only labels", got)
		}
	})
}
