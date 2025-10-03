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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// +kcc:proto=google.api.apikeys.v2.AndroidApplication
type AndroidApplication struct {
	// The SHA1 fingerprint of the application. For example, both sha1 formats are
	//  acceptable : DA:39:A3:EE:5E:6B:4B:0D:32:55:BF:EF:95:60:18:90:AF:D8:07:09 or
	//  DA39A3EE5E6B4B0D3255BFEF95601890AFD80709.
	//  Output format is the latter.
	// +required
	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty"`

	// The package name of the application.
	// +required
	PackageName *string `json:"packageName,omitempty"`
}

// +kcc:proto=google.api.apikeys.v2.AndroidKeyRestrictions
type AndroidKeyRestrictions struct {
	// A list of Android applications that are allowed to make API calls with this key.
	// +required
	AllowedApplications []AndroidApplication `json:"allowedApplications,omitempty"`
}

// +kcc:proto=google.api.apikeys.v2.ApiTarget
type ApiTarget struct {
	// The service for this restriction. It should be the canonical
	//  service name, for example: `translate.googleapis.com`.
	//  You can use [`gcloud services list`](/sdk/gcloud/reference/services/list)
	//  to get a list of services that are enabled in the project.
	// +required
	Service *string `json:"service,omitempty"`

	// Optional. List of one or more methods that can be called.
	//  If empty, all methods for the service are allowed. A wildcard
	//  (*) can be used as the last symbol.
	//  Valid examples:
	//    `google.cloud.translate.v2.TranslateService.GetSupportedLanguage`
	//    `TranslateText`
	//    `Get*`
	//    `translate.googleapis.com.Get*`
	Methods []string `json:"methods,omitempty"`
}

// +kcc:proto=google.api.apikeys.v2.BrowserKeyRestrictions
type BrowserKeyRestrictions struct {
	// A list of regular expressions for the referrer URLs that are allowed
	//  to make API calls with this key.
	// +required
	AllowedReferrers []string `json:"allowedReferrers,omitempty"`
}

// +kcc:proto=google.api.apikeys.v2.IosKeyRestrictions
type IosKeyRestrictions struct {
	// A list of bundle IDs that are allowed when making API calls with this key.
	// +required
	AllowedBundleIds []string `json:"allowedBundleIds,omitempty"`
}

// +kcc:spec:proto=google.api.apikeys.v2.Key
type APIKeySpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1beta1.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Human-readable display name of this key that you can modify.
	//  The maximum length is 63 characters.
	DisplayName *string `json:"displayName,omitempty"`

	/*NOTYET
	// Annotations is an unstructured key-value map stored with a policy that
	//  may be set by external tools to store and retrieve arbitrary metadata.
	//  They are not queryable and should be preserved when modifying objects.
	Annotations []Key_AnnotationsEntry `json:"annotations,omitempty"`
	*/

	// Key restrictions.
	Restrictions *Restrictions `json:"restrictions,omitempty"`
}

// +kcc:observedstate:proto=google.api.apikeys.v2.Key
type APIKeyObservedState struct {
	/*NOTYET
	// Output only. The resource name of the key.
	//  The `name` has the form:
	//  `projects/<PROJECT_NUMBER>/locations/global/keys/<KEY_ID>`.
	//  For example:
	//  `projects/123456867718/locations/global/keys/b7ff1f9f-8275-410a-94dd-3855ee9b5dd2`
	//
	//  NOTE: Key is a global resource; hence the only supported value for
	//  location is `global`.
	Name *string `json:"name,omitempty"`
	*/

	// Output only. Unique id in UUID4 format.
	Uid *string `json:"uid,omitempty"`

	/*NOTYET
	// Output only. An encrypted and signed value held by this key.
	//  This field can be accessed only through the `GetKeyString` method.
	KeyString *string `json:"keyString,omitempty"`
	*/

	/*NOTYET
	// Output only. A timestamp identifying the time this key was originally
	//  created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. A timestamp identifying the time this key was last
	//  updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A timestamp when this key was deleted. If the resource is not
	//  deleted, this must be empty.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. A checksum computed by the server based on the current value
	//  of the Key resource. This may be sent on update and delete requests to
	//  ensure the client has an up-to-date value before proceeding. See
	//  https://google.aip.dev/154.
	Etag *string `json:"etag,omitempty"`
	*/
}

// +kcc:proto=google.api.apikeys.v2.Key.AnnotationsEntry
type Key_AnnotationsEntry struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.api.apikeys.v2.Restrictions
type Restrictions struct {
	// The HTTP referrers (websites) that are allowed to use the key.
	BrowserKeyRestrictions *BrowserKeyRestrictions `json:"browserKeyRestrictions,omitempty"`

	// The IP addresses of callers that are allowed to use the key.
	ServerKeyRestrictions *ServerKeyRestrictions `json:"serverKeyRestrictions,omitempty"`

	// The Android apps that are allowed to use the key.
	AndroidKeyRestrictions *AndroidKeyRestrictions `json:"androidKeyRestrictions,omitempty"`

	// The iOS apps that are allowed to use the key.
	IosKeyRestrictions *IosKeyRestrictions `json:"iosKeyRestrictions,omitempty"`

	// A restriction for a specific service and optionally one or
	//  more specific methods. Requests are allowed if they
	//  match any of these restrictions. If no restrictions are
	//  specified, all targets are allowed.
	ApiTargets []ApiTarget `json:"apiTargets,omitempty"`
}

// +kcc:proto=google.api.apikeys.v2.ServerKeyRestrictions
type ServerKeyRestrictions struct {
	// A list of the caller IP addresses that are allowed to make API calls
	//  with this key.
	// +required
	AllowedIps []string `json:"allowedIps,omitempty"`
}
