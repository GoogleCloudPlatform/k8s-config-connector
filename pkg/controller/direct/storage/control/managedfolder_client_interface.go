// Copyright 2025 Google LLC
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

package storagecontrol

import (
	"context"

	control "cloud.google.com/go/storage/control/apiv2"
	controlpb "cloud.google.com/go/storage/control/apiv2/controlpb"
	"github.com/googleapis/gax-go/v2"
)

// ManagedFolderAPI is an interface for the StorageManagedFolder GCP client.
type ManagedFolderAPI interface {
	CreateManagedFolder(ctx context.Context, in *controlpb.CreateManagedFolderRequest, opts ...gax.CallOption) (*controlpb.ManagedFolder, error)
	DeleteManagedFolder(ctx context.Context, in *controlpb.DeleteManagedFolderRequest, opts ...gax.CallOption) error
	GetManagedFolder(ctx context.Context, in *controlpb.GetManagedFolderRequest, opts ...gax.CallOption) (*controlpb.ManagedFolder, error)
	ListManagedFolders(ctx context.Context, in *controlpb.ListManagedFoldersRequest, opts ...gax.CallOption) *control.ManagedFolderIterator
}
