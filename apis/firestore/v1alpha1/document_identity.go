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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/firestore/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// DocumentIDURL is the format for the externalRef of a FirestoreDocument.
	DocumentIDURL = v1beta1.DatabaseIDURL + "/documents/{{collection}}/{{document}}"
	ServiceDomain = "firestore.googleapis.com"
)

var _ identity.Identity = &FirestoreDocumentIdentity{}

// FirestoreDocumentIdentity represents the identity of a Firestore Document.
// +k8s:deepcopy-gen=false
type FirestoreDocumentIdentity struct {
	Parent     *v1beta1.FirestoreDatabaseIdentity
	Collection string
	Document   string
}

func (i *FirestoreDocumentIdentity) String() string {
	return i.Parent.String() + "/documents/" + i.Collection + "/" + i.Document
}

func (i *FirestoreDocumentIdentity) FromExternal(ref string) error {
	ref = strings.TrimPrefix(ref, "//firestore.googleapis.com/")

	tokens := strings.Split(ref, "/documents/")
	if len(tokens) != 2 {
		return fmt.Errorf("format of FirestoreDocument external=%q was not known (use %s)", ref, DocumentIDURL)
	}
	i.Parent = &v1beta1.FirestoreDatabaseIdentity{}
	if err := i.Parent.FromExternal(tokens[0]); err != nil {
		return err
	}
	collectionAndDocument := strings.SplitN(tokens[1], "/", 2)
	if len(collectionAndDocument) != 2 {
		return fmt.Errorf("format of FirestoreDocument external=%q was not known (use %s)", ref, DocumentIDURL)
	}
	i.Collection = collectionAndDocument[0]
	if i.Collection == "" {
		return fmt.Errorf("collection was empty in external=%q", ref)
	}
	i.Document = collectionAndDocument[1]
	if i.Document == "" {
		return fmt.Errorf("document was empty in external=%q", ref)
	}
	return nil
}

var _ identity.Resource = &FirestoreDocument{}

func (obj *FirestoreDocument) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &FirestoreDocumentIdentity{}

	// Resolve Parent
	if err := obj.Spec.DatabaseRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.parentRef: %w", err)
	}
	newIdentity.Parent = &v1beta1.FirestoreDatabaseIdentity{}
	if err := newIdentity.Parent.FromExternal(obj.Spec.DatabaseRef.External); err != nil {
		return nil, fmt.Errorf("parsing databaseRef.external=%q: %w", obj.Spec.DatabaseRef.External, err)
	}
	// Get desired ID
	newIdentity.Collection = common.ValueOf(obj.Spec.Collection)
	if newIdentity.Collection == "" {
		return nil, fmt.Errorf("spec.collection must be specified")
	}
	newIdentity.Document = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.Document == "" {
		newIdentity.Document = obj.GetName()
	}
	if newIdentity.Document == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &FirestoreDocumentIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}
