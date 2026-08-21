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

const (
	// ConfigConnectorAllowedName is the only recognized name for the ConfigConnector (CC) object.
	// This follows the pattern where a singleton object must have the name that matches the resource name.
	ConfigConnectorAllowedName = "configconnector.core.cnrm.cloud.google.com"

	// ConfigConnectorContextAllowedName is the only recognized name for the ConfigConnectorContext (CCC) object.
	// This follows the pattern where a singleton object must have the name that matches the resource name.
	ConfigConnectorContextAllowedName = "configconnectorcontext.core.cnrm.cloud.google.com"

	// ConfigConnectorContextNamespaceLabel is a label added to objects owned by the CCC.
	// The value of the label is the name fo the namespace.
	ConfigConnectorContextNamespaceLabel = "configconnectorcontext.cnrm.cloud.google.com/namespace"
)
