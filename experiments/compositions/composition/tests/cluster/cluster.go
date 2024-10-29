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

package cluster

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/cluster/configcontroller"
	"github.com/GoogleCloudPlatform/k8s-config-connector/experiments/compositions/composition/tests/cluster/kind"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

const (
	CRDManifests          = "../../release/test/crds.yaml"
	FacadeCRDManifests    = "../../release/test/facade_crds.yaml"
	KindOperatorManifests = "../../release/test/kind-operator.yaml"
	CCOperatorManifests   = "../../release/test/cc-operator.yaml"
)

type ClusterUser interface {
	Config() *rest.Config
	Name() string
	RestartWorkloads() error
	WaitForWorkloads() error
	KCCInstalled() bool
	Context() map[string]string
}

type ClusterSet struct {
	available sync.Map
}

var clusterSet ClusterSet

func AddCluster(c ClusterUser) {
	clusterSet.available.Store(c.Name(), c)
}

func RemoveCluster(c ClusterUser) {
	clusterSet.available.Delete(c.Name())
}

func ReserveCluster(t *testing.T) ClusterUser {
	var cluster ClusterUser
	found := false
	for !found {
		clusterSet.available.Range(func(k, v any) bool {
			c, loaded := clusterSet.available.LoadAndDelete(k)
			if !loaded {
				return true
			}
			found = true
			cluster = c.(ClusterUser)
			return false
		})

		if found {
			continue
		}

		t.Logf("\nWaiting for a cluster to become available...")
		time.Sleep(5 * time.Second)
	}

	t.Logf("Reserved cluster %s", cluster.Name())
	return cluster
}

func ReleaseCluster(t *testing.T, cluster ClusterUser) {
	clusterSet.available.Store(cluster.Name(), cluster)
	t.Logf("Released cluster %s", cluster.Name())
}

// Create Kind Clusters
func CreateKindClusters(clusterCount int, images string) {
	var wg sync.WaitGroup
	wg.Add(clusterCount)
	// Start with 1 e2e cluster
	for i := 0; i < clusterCount; i++ {
		go func(index int) {
			name := fmt.Sprintf("composition-e2e-%d", index)
			// kind cluster
			kc := kind.NewCluster(name,
				// that adds these images
				strings.Split(images, ","),
				// and installs these manifests
				[]string{CRDManifests, KindOperatorManifests},
				// and waits for these deployments to be ready
				[]types.NamespacedName{
					{Namespace: "composition-system", Name: "composition-controller-manager"},
				},
			)

			// Bring up the cluster and install the operator
			err := kc.ClusterUp()
			if err != nil {
				log.Fatalf("Error creating kind cluster: %s, %v", name, err)
			}
			// TODO ADD CLUSTER
			AddCluster(kc)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// Create CC Cluster
func CreateCCClusters(clusterCount int, images string) {
	var wg sync.WaitGroup
	if clusterCount > 1 {
		log.Fatalf("clusterCount must be 1 for CC clusters")
		return
	}
	wg.Add(clusterCount)
	// Start with 1 e2e cluster
	for i := 0; i < clusterCount; i++ {
		go func(index int) {
			name := fmt.Sprintf("composition-e2e-%d", index)
			// cc cluster

			// HACK for master CIDR
			// https://stackoverflow.com/questions/55399928/cidr-range-for-master-ipv4-cidr-in-gke-private-cluster
			// You can use any of the following CIDRs
			//   10.0.0.0 – 10.255.255.255,
			//   172.16.0.0 – 172.31.255.255,
			//   192.168.0.0 – 192.168.255.255
			//   except 172.16 and 172.17.
			//   Also this must be /28
			masterCidr := fmt.Sprintf("172.18.%d.0/28", index)

			cc := configcontroller.NewCluster(name, masterCidr,
				// and installs these manifests
				[]string{CRDManifests, CCOperatorManifests},
				// and waits for these deployments to be ready
				[]types.NamespacedName{
					{Namespace: "composition-system", Name: "composition-controller-manager"},
				},
			)

			// Bring up the cluster and install the operator
			err := cc.ClusterUp()
			if err != nil {
				log.Fatalf("Error creating CC cluster: %s, %v", name, err)
			}

			// TODO ADD CLUSTER
			AddCluster(cc)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
