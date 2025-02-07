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


// +kcc:proto=google.analytics.admin.v1beta.DataStream
type DataStream struct {
	// Data specific to web streams. Must be populated if type is
	//  WEB_DATA_STREAM.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.web_stream_data
	WebStreamData *DataStream_WebStreamData `json:"webStreamData,omitempty"`

	// Data specific to Android app streams. Must be populated if type is
	//  ANDROID_APP_DATA_STREAM.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.android_app_stream_data
	AndroidAppStreamData *DataStream_AndroidAppStreamData `json:"androidAppStreamData,omitempty"`

	// Data specific to iOS app streams. Must be populated if type is
	//  IOS_APP_DATA_STREAM.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.ios_app_stream_data
	IosAppStreamData *DataStream_IosAppStreamData `json:"iosAppStreamData,omitempty"`

	// Required. Immutable. The type of this DataStream resource.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.type
	Type *string `json:"type,omitempty"`

	// Human-readable display name for the Data Stream.
	//
	//  Required for web data streams.
	//
	//  The max allowed display name length is 255 UTF-16 code units.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataStream.AndroidAppStreamData
type DataStream_AndroidAppStreamData struct {

	// Immutable. The package name for the app being measured.
	//  Example: "com.example.myandroidapp"
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.AndroidAppStreamData.package_name
	PackageName *string `json:"packageName,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataStream.IosAppStreamData
type DataStream_IosAppStreamData struct {

	// Required. Immutable. The Apple App Store Bundle ID for the app
	//  Example: "com.example.myiosapp"
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.IosAppStreamData.bundle_id
	BundleID *string `json:"bundleID,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataStream.WebStreamData
type DataStream_WebStreamData struct {

	// Domain name of the web app being measured, or empty.
	//  Example: "http://www.google.com", "https://www.google.com"
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.WebStreamData.default_uri
	DefaultURI *string `json:"defaultURI,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataStream
type DataStreamObservedState struct {
	// Data specific to web streams. Must be populated if type is
	//  WEB_DATA_STREAM.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.web_stream_data
	WebStreamData *DataStream_WebStreamDataObservedState `json:"webStreamData,omitempty"`

	// Data specific to Android app streams. Must be populated if type is
	//  ANDROID_APP_DATA_STREAM.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.android_app_stream_data
	AndroidAppStreamData *DataStream_AndroidAppStreamDataObservedState `json:"androidAppStreamData,omitempty"`

	// Data specific to iOS app streams. Must be populated if type is
	//  IOS_APP_DATA_STREAM.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.ios_app_stream_data
	IosAppStreamData *DataStream_IosAppStreamDataObservedState `json:"iosAppStreamData,omitempty"`

	// Output only. Resource name of this Data Stream.
	//  Format: properties/{property_id}/dataStreams/{stream_id}
	//  Example: "properties/1000/dataStreams/2000"
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.name
	Name *string `json:"name,omitempty"`

	// Output only. Time when this stream was originally created.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when stream payload fields were last updated.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataStream.AndroidAppStreamData
type DataStream_AndroidAppStreamDataObservedState struct {
	// Output only. ID of the corresponding Android app in Firebase, if any.
	//  This ID can change if the Android app is deleted and recreated.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.AndroidAppStreamData.firebase_app_id
	FirebaseAppID *string `json:"firebaseAppID,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataStream.IosAppStreamData
type DataStream_IosAppStreamDataObservedState struct {
	// Output only. ID of the corresponding iOS app in Firebase, if any.
	//  This ID can change if the iOS app is deleted and recreated.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.IosAppStreamData.firebase_app_id
	FirebaseAppID *string `json:"firebaseAppID,omitempty"`
}

// +kcc:proto=google.analytics.admin.v1beta.DataStream.WebStreamData
type DataStream_WebStreamDataObservedState struct {
	// Output only. Analytics Measurement ID.
	//
	//  Example: "G-1A2BCD345E"
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.WebStreamData.measurement_id
	MeasurementID *string `json:"measurementID,omitempty"`

	// Output only. ID of the corresponding web app in Firebase, if any.
	//  This ID can change if the web app is deleted and recreated.
	// +kcc:proto:field=google.analytics.admin.v1beta.DataStream.WebStreamData.firebase_app_id
	FirebaseAppID *string `json:"firebaseAppID,omitempty"`
}
