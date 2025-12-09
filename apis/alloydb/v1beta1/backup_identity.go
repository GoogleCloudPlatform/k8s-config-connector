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

package v1beta1

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// BackupIdentityURL is the format for the externalRef of an alloydb Backup.
	// This format is used in the AlloyDB API.
	// See: https://docs.cloud.google.com/asset-inventory/docs/asset-names
	BackupIdentityURL = "projects/{{projectID}}/locations/{{location}}/backups/{{backup}}"
)

var _ identity.Identity = &BackupIdentity{}
var _ identity.Resource = &AlloyDBBackup{}

var parser = regexp.MustCompile(`((//)?alloydb.googleapis.com)?/?projects/(?P<projects>[[:alpha:]]+)/locations/(?P<locations>[[:alpha:]]+)/backups/(?P<backups>[[:alpha:]]+)`)

// BackupIdentity represents the identity of an alloydb backup.
// +k8s:deepcopy-gen=false
type BackupIdentity struct {
	Parent   string
	Location string
	Backup   string
}

func (i *BackupIdentity) String() string {
	return "projects/" + i.Parent + "/locations/" + i.Location + "/backups/" + i.Backup
}

func (i *BackupIdentity) FromExternal(ref string) error {
	// TODO: Should be able to parse https://docs.cloud.google.com/asset-inventory/docs/asset-names
	// But that format is //alloydb.googleapis.com/projects/PROJECT_ID/locations/LOCATION/backups/BACKUP
	// which is not the format used by the service.

	err, identityMap := parseIdentityMap(ref, 3)
	if err != nil {
		return fmt.Errorf("format of backup external=%q was not known (use %s): %w", ref, BackupIdentityURL, err)
	}

	i.Parent = identityMap["projects"]
	i.Location = identityMap["locations"]
	i.Backup = identityMap["backups"]

	return nil
}

func parseIdentityMap(ref string, cnt int) (error, map[string]string) {
	raw := parser.FindStringSubmatch(ref)
	if raw == nil {
		return fmt.Errorf("reference %s did not match expected format", ref), nil
	}
	result := make(map[string]string, cnt)
	for i, name := range parser.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = raw[i]
		}
	}
	if len(result) != cnt {
		return fmt.Errorf("reference %s failed to parse %d values", ref, cnt), nil
	}
	return nil, result
}

func (obj *AlloyDBBackup) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	// Get desired resource ID
	resourceID := common.ValueOf(obj.Spec.ResourceID)

	// Server-generated ID; do not fallback to name
	// if resourceID == "" {
	// 	resourceID = obj.GetName()
	// }

	var specIdentity *BackupIdentity
	if resourceID != "" {
		specIdentity = &BackupIdentity{}
		if !strings.HasPrefix(resourceID, "projects/") {
			resourceID = "projects/" + resourceID
		}
		if err := specIdentity.FromExternal(resourceID); err != nil {
			return nil, fmt.Errorf("cannot parse spec.resourceID=%q: %w", resourceID, err)
		}
	}

	// Validate against the ID stored in status.externalRef
	var statusIdentity *BackupIdentity
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity = &BackupIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
	}

	if specIdentity != nil {
		if statusIdentity != nil && statusIdentity.String() != specIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, specIdentity.String())
		}
		return specIdentity, nil
	}

	if statusIdentity != nil {
		return statusIdentity, nil
	}

	return nil, fmt.Errorf("cannot determine identity: spec.resourceID and status.externalRef are both unset")
}
