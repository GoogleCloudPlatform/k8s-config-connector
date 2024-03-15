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

package kubecli

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
)

type ClusterOptions struct {
	// Impersonate is the configuration that RESTClient will use for impersonation.
	Impersonate *rest.ImpersonationConfig
}

func (o *ClusterOptions) PopulateDefaults() {

}

func (o *ClusterOptions) AddFlags(cmd *cobra.Command) {
}

type ObjectOptions struct {
	// Kind specifies the kind we want to change.  It will be matched against kind, resource-name, aliases etc.
	Kind string
	// Name is the name of the object we want to change
	Name string
	// Namespace is the namespace of the object we want to change
	Namespace string
}

func (o *ObjectOptions) PopulateDefaults() {

}

func (o *ObjectOptions) AddFlags(cmd *cobra.Command) {
	// cmd.Flags().StringVar(&o.Name, "name", o.Name, "Name of the object to change")
	cmd.Flags().StringVarP(&o.Namespace, "namespace", "n", o.Namespace, "Namespace of the object")
}
