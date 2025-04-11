// Copyright 2024 Google LLC
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

package v1beta1

import (
	"context"
	"fmt"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type BasicCryptoKeyRef_OneOf struct {
	// Default CMEK crypto key. This is for API backward compatibility and cannot be changed.
	*refs.KMSCryptoKeyRef `json:",inline"`
	// A reference to the Autokey `KMSKeyHandle`, which auto generates a crypto key.
	KMSKeyHandleRef *KMSKeyHandleRef `json:"autokeyRef,omitempty"`
}

func ResolveKMSCryptoKeyExternal(ctx context.Context, reader client.Reader, src client.Object, oneof *BasicCryptoKeyRef_OneOf) (string, error) {

	if oneof.KMSCryptoKeyRef != nil && oneof.KMSKeyHandleRef != nil {
		return "", fmt.Errorf("KMSCryptoKey should either use `.external` or `.autokeyRef`")
	}

	// default cryptokey
	if oneof.KMSCryptoKeyRef != nil {
		var err error
		oneof.KMSCryptoKeyRef, err = refs.ResolveKMSCryptoKeyRef(ctx, reader, src, oneof.KMSCryptoKeyRef)
		if err != nil {
			return "", err
		}

		return oneof.KMSCryptoKeyRef.External, nil
	}

	return oneof.KMSKeyHandleRef.NormalizedCryptoKey(ctx, reader, src.GetNamespace())
}
