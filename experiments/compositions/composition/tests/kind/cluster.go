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

package kind

import (
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"testing"
	"time"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

const (
	ReuseIfPresent    bool = true
	RecreateIfPresent bool = false
)

type kindCluster struct {
	name   string
	config *rest.Config
}

type KindClusterReader interface {
	Config() *rest.Config
	Name() string
}

type KindCluster interface {
	Create() error
	Config() *rest.Config
	Delete() error
	Exists() (bool, error)
	Name() string
	LoadImage(string) error
}

type KindClusterSet struct {
	available sync.Map
}

var kindClusterSet KindClusterSet

func ReserveCluster(t *testing.T) KindClusterReader {
	var cluster KindClusterReader
	found := false
	for !found {
		kindClusterSet.available.Range(func(k, v any) bool {
			v, loaded := kindClusterSet.available.LoadAndDelete(k)
			if !loaded {
				return true
			}
			found = true
			cluster = v.(KindClusterReader)
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

func ReleaseCluster(t *testing.T, cluster KindClusterReader) {
	kindClusterSet.available.Store(cluster.Name(), cluster)
	t.Logf("Released cluster %s", cluster.Name())
}

func verifyKindIsInstalled() error {
	_, err := exec.LookPath("kind")
	if err != nil {
		return err
	}
	return nil
}

// NewKindCluster - return a cluster setup object
func NewKindCluster(name string) (KindCluster, error) {
	err := verifyKindIsInstalled()
	if err != nil {
		return nil, err
	}
	cluster := &kindCluster{
		name: name,
	}
	return cluster, nil
}

// Wait for all clusters to become ready
func (c *kindCluster) Create() error {
	err := c.Delete()
	if err != nil {
		return err
	}
	clusterConfig, err := c.kindClusterDefinition()
	if err != nil {
		return err
	}
	defer os.Remove(clusterConfig)

	c.config, err = c.createCluster(clusterConfig)
	if err != nil {
		return err
	}

	kindClusterSet.available.Store(c.Name(), c)
	return nil
}

// Config return rest.Config
func (c *kindCluster) Config() *rest.Config {
	return c.config
}

// Name return name
func (c *kindCluster) Name() string {
	return c.name
}

func (c *kindCluster) String() string {
	return c.name
}

// LoadImage loads a docker image into the cluster
func (c *kindCluster) LoadImage(image string) error {
	err := exec.Command("kind", "load", "docker-image", image, "--name", c.name).Run()
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes all clusters
func (c *kindCluster) Delete() error {
	kindClusterSet.available.Delete(c.Name())
	err := exec.Command("kind", "delete", "cluster", "--name", c.name).Run()
	if err != nil {
		return err
	}
	return nil
}

// Exists checks if the cluster exists
func (c *kindCluster) Exists() (bool, error) {
	output, err := exec.Command("kind", "get", "clusters").CombinedOutput()
	if err != nil {
		return false, err
	}

	for _, cluster := range strings.Split(string(output), "\n") {
		if cluster == c.name {
			return true, nil
		}
	}
	return false, nil
}

func (c *kindCluster) kindClusterDefinition() (string, error) {
	ipAddress, err := c.getHostIPAddress()
	if err != nil {
		return "", err
	}
	clusterConfigFile, err := os.CreateTemp("", "kind-cluster.yaml")
	if err != nil {
		return "", err
	}
	defer clusterConfigFile.Close()
	kindClusterConfig := `kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  # Allow connections to the API Sever with the host IP address
  apiServerAddress: "` + ipAddress + `"`
	bytes := []byte(kindClusterConfig)
	_, err = clusterConfigFile.Write(bytes)
	return clusterConfigFile.Name(), err
}

func (c *kindCluster) createCluster(clusterConfig string) (*rest.Config, error) {
	_, err := exec.Command("kind", "create", "cluster", "--name", c.name, "--config", clusterConfig).CombinedOutput()
	if err != nil {
		return nil, err
	}

	kubeConfigFile, err := os.CreateTemp("", "kubeconfig.yaml")
	if err != nil {
		return nil, err
	}
	defer os.Remove(kubeConfigFile.Name())
	content, err := exec.Command("kind", "get", "kubeconfig", "--name", c.name).CombinedOutput()
	if err != nil {
		return nil, err
	}
	bytes := []byte(content)
	_, err = kubeConfigFile.Write(bytes)
	kubeConfigFile.Close()

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigFile.Name()},
		&clientcmd.ConfigOverrides{
			ClusterInfo: clientcmdapi.Cluster{
				Server: "",
			},
			CurrentContext: "",
		}).ClientConfig()
}

func (c *kindCluster) getHostIPAddress() (string, error) {
	// Try getting host ip by creating a connection object and reading the localaddr
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
