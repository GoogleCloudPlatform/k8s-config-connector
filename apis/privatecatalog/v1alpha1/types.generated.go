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


// +kcc:proto=google.cloud.privatecatalog.v1beta1.AssetReference
type AssetReference struct {

	// The version of the source used for this asset reference.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.GcsSource
type GcsSource struct {
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.GitSource
type GitSource struct {
	// Location of the Git repo to build.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.GitSource.repo
	Repo *string `json:"repo,omitempty"`

	// Directory, relative to the source root, in which to run the build.
	//
	//  This must be a relative path. If a step's `dir` is specified and is an
	//  absolute path, this value is ignored for that step's execution.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.GitSource.dir
	Dir *string `json:"dir,omitempty"`

	// The revision commit to use.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.GitSource.commit
	Commit *string `json:"commit,omitempty"`

	// The revision branch to use.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.GitSource.branch
	Branch *string `json:"branch,omitempty"`

	// The revision tag to use.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.GitSource.tag
	Tag *string `json:"tag,omitempty"`
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.Inputs
type Inputs struct {
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.Product
type Product struct {
}

// +kcc:proto=google.longrunning.Operation
type Operation struct {
	// The server-assigned name, which is only unique within the same service that
	//  originally returns it. If you use the default HTTP mapping, the
	//  `name` should be a resource name ending with `operations/{unique_id}`.
	// +kcc:proto:field=google.longrunning.Operation.name
	Name *string `json:"name,omitempty"`

	// Service-specific metadata associated with the operation.  It typically
	//  contains progress information and common metadata such as create time.
	//  Some services might not provide such metadata.  Any method that returns a
	//  long-running operation should document the metadata type, if any.
	// +kcc:proto:field=google.longrunning.Operation.metadata
	Metadata *Any `json:"metadata,omitempty"`

	// If the value is `false`, it means the operation is still in progress.
	//  If `true`, the operation is completed, and either `error` or `response` is
	//  available.
	// +kcc:proto:field=google.longrunning.Operation.done
	Done *bool `json:"done,omitempty"`

	// The error result of the operation in case of failure or cancellation.
	// +kcc:proto:field=google.longrunning.Operation.error
	Error *Status `json:"error,omitempty"`

	// The normal, successful response of the operation.  If the original
	//  method returns no data on success, such as `Delete`, the response is
	//  `google.protobuf.Empty`.  If the original method is standard
	//  `Get`/`Create`/`Update`, the response should be the resource.  For other
	//  methods, the response should have the type `XxxResponse`, where `Xxx`
	//  is the original method name.  For example, if the original method name
	//  is `TakeSnapshot()`, the inferred response type is
	//  `TakeSnapshotResponse`.
	// +kcc:proto:field=google.longrunning.Operation.response
	Response *Any `json:"response,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.AssetReference
type AssetReferenceObservedState struct {
	// Output only. A unique identifier among asset references in a product.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.id
	ID *string `json:"id,omitempty"`

	// Output only. The human-readable description of the referenced asset. Maximum 256
	//  characters in length.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.description
	Description *string `json:"description,omitempty"`

	// Output only. The definition of input parameters to hydrate the asset template.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.inputs
	Inputs *Inputs `json:"inputs,omitempty"`

	// Output only. The current state of the asset reference.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.validation_status
	ValidationStatus *string `json:"validationStatus,omitempty"`

	// Output only. The validation process metadata.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.validation_operation
	ValidationOperation *Operation `json:"validationOperation,omitempty"`

	// Output only. The asset resource name if an asset is hosted by Private Catalog.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.asset
	Asset *string `json:"asset,omitempty"`

	// Output only. The cloud storage object path.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.gcs_path
	GcsPath *string `json:"gcsPath,omitempty"`

	// Output only. The git source.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.git_source
	GitSource *GitSource `json:"gitSource,omitempty"`

	// Output only. The cloud storage source.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.gcs_source
	GcsSource *GcsSource `json:"gcsSource,omitempty"`

	// Output only. The creation timestamp of the asset reference.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update timestamp of the asset reference.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.AssetReference.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.GcsSource
type GcsSourceObservedState struct {
	// Output only. the cloud storage object path.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.GcsSource.gcs_path
	GcsPath *string `json:"gcsPath,omitempty"`

	// Output only. Generation of the object, which is set when the content of an object starts
	//  being written.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.GcsSource.generation
	Generation *int64 `json:"generation,omitempty"`

	// Output only. The time when the object metadata was last changed.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.GcsSource.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.Inputs
type InputsObservedState struct {
	// Output only. The JSON schema defining the inputs and their formats.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Inputs.parameters
	Parameters map[string]string `json:"parameters,omitempty"`
}

// +kcc:proto=google.cloud.privatecatalog.v1beta1.Product
type ProductObservedState struct {
	// Output only. The resource name of the target product, in the format of
	//  `products/[a-z][-a-z0-9]*[a-z0-9]'.
	//
	//  A unique identifier for the product under a catalog.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Product.name
	Name *string `json:"name,omitempty"`

	// Output only. The type of the product asset. It can be one of the following values:
	//
	//  * `google.deploymentmanager.Template`
	//  * `google.cloudprivatecatalog.ListingOnly`
	//  * `google.cloudprivatecatalog.Terraform`
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Product.asset_type
	AssetType *string `json:"assetType,omitempty"`

	// Required. Output only. The display metadata to describe the product. The JSON schema of the
	//  metadata differs by [Product.asset_type][google.cloud.privatecatalog.v1beta1.Product.asset_type].
	//  When the type is `google.deploymentmanager.Template`, the schema is as
	//  follows:
	//
	//  ```
	//  "$schema": http://json-schema.org/draft-04/schema#
	//  type: object
	//  properties:
	//    name:
	//      type: string
	//      minLength: 1
	//      maxLength: 64
	//    description:
	//      type: string
	//      minLength: 1
	//      maxLength: 2048
	//    tagline:
	//      type: string
	//      minLength: 1
	//      maxLength: 100
	//    support_info:
	//      type: string
	//      minLength: 1
	//      maxLength: 2048
	//    creator:
	//      type: string
	//      minLength: 1
	//      maxLength: 100
	//    documentations:
	//      type: array
	//      items:
	//        type: object
	//        properties:
	//          url:
	//            type: string
	//            pattern:
	//            "^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]"
	//          title:
	//            type: string
	//            minLength: 1
	//            maxLength: 64
	//          description:
	//            type: string
	//            minLength: 1
	//            maxLength: 2048
	//  required:
	//  - name
	//  - description
	//  additionalProperties: false
	//
	//  ```
	//
	//  When the asset type is `google.cloudprivatecatalog.ListingOnly`, the schema
	//  is as follows:
	//
	//  ```
	//  "$schema": http://json-schema.org/draft-04/schema#
	//  type: object
	//  properties:
	//    name:
	//      type: string
	//      minLength: 1
	//      maxLength: 64
	//    description:
	//      type: string
	//      minLength: 1
	//      maxLength: 2048
	//    tagline:
	//      type: string
	//      minLength: 1
	//      maxLength: 100
	//    support_info:
	//      type: string
	//      minLength: 1
	//      maxLength: 2048
	//    creator:
	//      type: string
	//      minLength: 1
	//      maxLength: 100
	//    documentations:
	//      type: array
	//      items:
	//        type: object
	//        properties:
	//          url:
	//            type: string
	//            pattern:
	//            "^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]"
	//          title:
	//            type: string
	//            minLength: 1
	//            maxLength: 64
	//          description:
	//            type: string
	//            minLength: 1
	//            maxLength: 2048
	//    signup_url:
	//      type: string
	//      pattern:
	//      "^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]"
	//  required:
	//  - name
	//  - description
	//  - signup_url
	//  additionalProperties: false
	//
	//  ```
	//
	//  When the asset type is `google.cloudprivatecatalog.Terraform`, the schema
	//  is as follows:
	//
	//  ```
	//  "$schema": http://json-schema.org/draft-04/schema#
	//  type: object
	//  properties:
	//    name:
	//      type: string
	//      minLength: 1
	//      maxLength: 64
	//    description:
	//      type: string
	//      minLength: 1
	//      maxLength: 2048
	//    tagline:
	//      type: string
	//      minLength: 1
	//      maxLength: 100
	//    support_info:
	//      type: string
	//      minLength: 1
	//      maxLength: 2048
	//    creator:
	//      type: string
	//      minLength: 1
	//      maxLength: 100
	//    documentations:
	//      type: array
	//      items:
	//        type: object
	//        properties:
	//          url:
	//            type: string
	//            pattern:
	//            "^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]"
	//          title:
	//            type: string
	//            minLength: 1
	//            maxLength: 64
	//          description:
	//            type: string
	//            minLength: 1
	//            maxLength: 2048
	//  required:
	//  - name
	//  - description
	//  additionalProperties: true
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Product.display_metadata
	DisplayMetadata map[string]string `json:"displayMetadata,omitempty"`

	// Output only. The icon URI of the product.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Product.icon_uri
	IconURI *string `json:"iconURI,omitempty"`

	// Output only. A collection of assets referred by a product.
	//  This field is set for Terraform Products only.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Product.asset_references
	AssetReferences []AssetReference `json:"assetReferences,omitempty"`

	// Output only. The time when the product was created.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Product.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the product was last updated.
	// +kcc:proto:field=google.cloud.privatecatalog.v1beta1.Product.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
