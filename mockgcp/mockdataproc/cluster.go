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

// +tool:mockgcp-support
// proto.service: google.cloud.dataproc.v1.ClusterController
// proto.message: google.cloud.dataproc.v1.Cluster

package mockdataproc

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/pkg/storage"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func (s *clusterControllerServer) GetCluster(ctx context.Context, req *pb.GetClusterRequest) (*pb.Cluster, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, req.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Not found: Cluster %s", fqn)
		}
		return nil, err
	}

	return obj, nil
}

func (s *clusterControllerServer) ListClusters(ctx context.Context, req *pb.ListClustersRequest) (*pb.ListClustersResponse, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, "")
	if err != nil {
		return nil, err
	}

	findPrefix := name.String()

	var clusters []*pb.Cluster

	findKind := (&pb.Cluster{}).ProtoReflect().Descriptor()
	if err := s.storage.List(ctx, findKind, storage.ListOptions{Prefix: findPrefix}, func(obj proto.Message) error {
		cluster := obj.(*pb.Cluster)
		clusters = append(clusters, cluster)
		return nil
	}); err != nil {
		return nil, err
	}

	return &pb.ListClustersResponse{
		Clusters: clusters,
	}, nil

}

func (s *clusterControllerServer) CreateCluster(ctx context.Context, req *pb.CreateClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, req.Cluster.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := proto.Clone(req.GetCluster()).(*pb.Cluster)
	obj.ProjectId = name.Project.ID
	obj.ClusterName = name.ClusterName
	s.setStatus(obj, pb.ClusterStatus_CREATING)

	s.populateDefaultsForCluster(obj, name)

	if err := s.storage.Create(ctx, fqn, obj); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		ClusterName:   name.ClusterName,
		ClusterUuid:   obj.ClusterUuid,
		OperationType: "CREATE",
		Description:   "Create cluster with 2 workers",
		Status: &pb.ClusterOperationStatus{
			InnerState:     "PENDING",
			State:          pb.ClusterOperationStatus_PENDING,
			StateStartTime: timestamppb.New(now),
		},

		Warnings: []string{
			"The firewall rules for specified network or subnetwork would allow ingress traffic from 0.0.0.0/0, which could be a security risk.",
			"The specified custom staging bucket '" + obj.GetConfig().GetConfigBucket() + "' is not using uniform bucket level access IAM configuration. It is recommended to update bucket to enable the same. See https://cloud.google.com/storage/docs/uniform-bucket-level-access.",
			"No image specified. Using the default image version. It is recommended to select a specific image version in production, as the default image version may change at any time.",
		},
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		lroMetadata.Status.InnerState = "DONE"
		lroMetadata.Status.State = pb.ClusterOperationStatus_DONE

		lroMetadata.StatusHistory = []*pb.ClusterOperationStatus{
			{
				State:          pb.ClusterOperationStatus_PENDING,
				StateStartTime: timestamppb.New(now),
			},
			{
				State:          pb.ClusterOperationStatus_RUNNING,
				StateStartTime: timestamppb.Now(),
			},
		}
		updated, err := mutateObject(ctx, s.storage, fqn, func(obj *pb.Cluster) error {
			s.setStatus(obj, pb.ClusterStatus_RUNNING)

			obj.Config.EndpointConfig = &pb.EndpointConfig{}
			if obj.Config.GceClusterConfig.InternalIpOnly == nil {
				obj.Config.GceClusterConfig.InternalIpOnly = PtrTo(true)
			}
			obj.Config.GceClusterConfig.NetworkUri = "https://www.googleapis.com/compute/v1/projects/" + name.Project.ID + "/global/networks/default"
			obj.Config.GceClusterConfig.ServiceAccountScopes = []string{"https://www.googleapis.com/auth/cloud-platform"}
			obj.Config.MasterConfig.DiskConfig.BootDiskSizeGb = 1000

			if obj.Config.GceClusterConfig.ShieldedInstanceConfig == nil {
				obj.Config.GceClusterConfig.ShieldedInstanceConfig = &pb.ShieldedInstanceConfig{}
			}
			obj.Config.GceClusterConfig.ShieldedInstanceConfig.EnableIntegrityMonitoring = PtrTo(true)
			obj.Config.GceClusterConfig.ShieldedInstanceConfig.EnableSecureBoot = PtrTo(true)
			obj.Config.GceClusterConfig.ShieldedInstanceConfig.EnableVtpm = PtrTo(true)

			obj.Config.GceClusterConfig.ZoneUri = "https://www.googleapis.com/compute/v1/projects/" + name.Project.ID + "/zones/us-central1-c"

			s.populateLabels(obj, name)
			obj.Labels["goog-drz-dataproc-uuid"] = "cluster-" + obj.ClusterUuid

			return nil
		})
		if err != nil {
			return nil, err
		}

		// Not all fields are returned in the LRO
		ret := proto.Clone(updated).(*pb.Cluster)
		ret.Status = nil
		ret.StatusHistory = nil
		ret.Config.WorkerConfig.InstanceNames = nil
		ret.Config.MasterConfig.InstanceNames = nil
		return ret, nil
	})
}

func populateMachineType(name *clusterName, machineType string, defaultMachineType string) string {
	if machineType == "" {
		machineType = "n2-standard-4"
	}

	// Ensure fully qualified
	zone := name.Region + "-c"
	tokens := strings.Split(machineType, "/")
	if len(tokens) == 1 {
		machineType = "https://www.googleapis.com/compute/v1/projects/" + name.Project.ID + "/zones/" + zone + "/machineTypes/" + machineType
	}
	return machineType
}

func (s *clusterControllerServer) populateLabels(obj *pb.Cluster, name *clusterName) {
	if obj.Labels == nil {
		obj.Labels = make(map[string]string)
	}
	obj.Labels["goog-dataproc-autozone"] = "enabled"
	obj.Labels["goog-dataproc-cluster-name"] = name.ClusterName
	obj.Labels["goog-dataproc-cluster-uuid"] = obj.ClusterUuid
	obj.Labels["goog-dataproc-location"] = name.Region
}

func (s *clusterControllerServer) populateInstanceGroupConfig(config *pb.InstanceGroupConfig, name *clusterName, isMaster bool) *pb.InstanceGroupConfig {
	if config == nil {
		config = &pb.InstanceGroupConfig{}
	}
	if config.DiskConfig == nil {
		config.DiskConfig = &pb.DiskConfig{}
	}

	if isMaster {
		if config.NumInstances == 0 {
			config.NumInstances = 1
		}

	} else {
		if config.NumInstances == 0 {
			config.NumInstances = 2
		}
		if config.DiskConfig.BootDiskSizeGb == 0 {
			config.DiskConfig.BootDiskSizeGb = 1000
		}
	}

	if config.DiskConfig.BootDiskType == "" {
		config.DiskConfig.BootDiskType = "pd-standard"
	}
	if config.ImageUri == "" {
		config.ImageUri = "https://www.googleapis.com/compute/v1/projects/cloud-dataproc/global/images/dataproc-2-2-deb12-20250212-155100-rc01"
	}

	config.MachineTypeUri = populateMachineType(name, config.MachineTypeUri, "n2-standard-4")

	if config.MinCpuPlatform == "" {
		config.MinCpuPlatform = "AUTOMATIC"
	}
	if config.Preemptibility == pb.InstanceGroupConfig_PREEMPTIBILITY_UNSPECIFIED {
		config.Preemptibility = pb.InstanceGroupConfig_NON_PREEMPTIBLE
	}

	if isMaster {
		config.InstanceNames = []string{name.ClusterName + "-m"}
	} else {
		instanceNames := []string{}
		for i := int32(0); i < config.NumInstances; i++ {
			s := fmt.Sprintf("%s-w-%d", name.ClusterName, i)
			instanceNames = append(instanceNames, s)
		}
		config.InstanceNames = instanceNames
	}
	return config
}

func (s *clusterControllerServer) populateDefaultsForCluster(obj *pb.Cluster, name *clusterName) {
	if obj.ClusterUuid == "" {
		obj.ClusterUuid = fmt.Sprintf("%x", time.Now().UnixNano())
	}
	if obj.Config == nil {
		obj.Config = &pb.ClusterConfig{}
	}

	if obj.Config.TempBucket == "" {
		obj.Config.TempBucket = fmt.Sprintf("dataproc-temp-%s-%d-xxxxxxxx", name.Region, name.Project.Number)
	}
	if obj.Config.ConfigBucket == "" {
		obj.Config.ConfigBucket = fmt.Sprintf("dataproc-staging-%s-%d-xxxxxxxx", name.Region, name.Project.Number)
	}

	obj.Config.MasterConfig = s.populateInstanceGroupConfig(obj.Config.MasterConfig, name, true)

	obj.Config.WorkerConfig = s.populateInstanceGroupConfig(obj.Config.WorkerConfig, name, false)

	s.populateSoftwareConfig(obj)
}

func (s *clusterControllerServer) UpdateCluster(ctx context.Context, req *pb.UpdateClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, req.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	obj := &pb.Cluster{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		return nil, err
	}

	description := ""

	updated := proto.Clone(obj).(*pb.Cluster)
	for _, field := range req.GetUpdateMask().GetPaths() {
		switch field {
		case "config.worker_config.num_instances":
			description = "Add 1 workers."
			updated.Config.WorkerConfig.NumInstances = req.GetCluster().GetConfig().GetWorkerConfig().GetNumInstances()
		case "labels":
			updated.Labels = req.GetCluster().GetLabels()
			s.populateLabels(updated, name)
		default:
			return nil, fmt.Errorf("updateMask %q not supported by mockgcp", field)
		}
	}

	s.setStatus(updated, pb.ClusterStatus_UPDATING)

	if err := s.storage.Update(ctx, fqn, updated); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		ClusterName:   name.ClusterName,
		ClusterUuid:   string(obj.ClusterUuid),
		Description:   description,
		OperationType: "UPDATE",
		Status: &pb.ClusterOperationStatus{
			InnerState:     "PENDING",
			State:          pb.ClusterOperationStatus_PENDING,
			StateStartTime: timestamppb.New(now),
		},
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		// Intermediate states appear in statusHistory
		s.setOperationStatus(lroMetadata, pb.ClusterOperationStatus_RUNNING)
		s.setOperationStatus(lroMetadata, pb.ClusterOperationStatus_DONE)

		if err := s.storage.Get(ctx, fqn, obj); err != nil {
			return nil, err
		}

		updated, err := mutateObject(ctx, s.storage, fqn, func(obj *pb.Cluster) error {
			s.setStatus(obj, pb.ClusterStatus_RUNNING)

			// Metrics are pretty delayed; simulate this by only populating them here
			s.populateMetrics(obj)

			s.populateDefaultsForCluster(obj, name)

			return nil
		})
		if err != nil {
			return nil, err
		}

		// Not all fields are returned in the LRO
		ret := proto.Clone(updated).(*pb.Cluster)
		ret.Status = nil
		ret.StatusHistory = nil
		ret.Config.WorkerConfig.InstanceNames = nil
		ret.Config.MasterConfig.InstanceNames = nil
		ret.Metrics = nil
		return ret, nil
	})
}

func (s *clusterControllerServer) setStatus(obj *pb.Cluster, state pb.ClusterStatus_State) {
	if obj.Status != nil {
		obj.StatusHistory = append(obj.StatusHistory, obj.Status)
	}
	now := time.Now()
	obj.Status = &pb.ClusterStatus{
		State:          state,
		StateStartTime: timestamppb.New(now),
	}

}

func (s *clusterControllerServer) setOperationStatus(obj *pb.ClusterOperationMetadata, state pb.ClusterOperationStatus_State) {
	if obj.Status != nil {
		obj.Status.InnerState = ""
		obj.StatusHistory = append(obj.StatusHistory, obj.Status)
	}
	now := time.Now()
	obj.Status = &pb.ClusterOperationStatus{
		InnerState:     state.String(),
		State:          state,
		StateStartTime: timestamppb.New(now),
	}
}

func (s *clusterControllerServer) populateMetrics(obj *pb.Cluster) {
	if obj.Metrics == nil {
		obj.Metrics = &pb.ClusterMetrics{}
	}
	obj.Metrics.HdfsMetrics = map[string]int64{
		"dfs-blocks-corrupt":                    0,
		"dfs-blocks-default-replication-factor": 2,
		"dfs-blocks-missing":                    0,
		"dfs-blocks-missing-repl-one":           0,
		"dfs-blocks-pending-deletion":           0,
		"dfs-blocks-under-replication":          0,
		"dfs-capacity-present":                  1995884572672,
		"dfs-capacity-remaining":                1995884523520,
		"dfs-capacity-total":                    2113237753856,
		"dfs-capacity-used":                     49152,
		"dfs-nodes-decommissioned":              0,
		"dfs-nodes-decommissioning":             0,
		"dfs-nodes-running":                     int64(obj.GetConfig().GetWorkerConfig().GetNumInstances()),
	}

	obj.Metrics.YarnMetrics = map[string]int64{
		"yarn-apps-completed":        0,
		"yarn-apps-failed":           0,
		"yarn-apps-killed":           0,
		"yarn-apps-pending":          0,
		"yarn-apps-running":          0,
		"yarn-apps-submitted":        0,
		"yarn-containers-allocated":  0,
		"yarn-containers-pending":    0,
		"yarn-containers-reserved":   0,
		"yarn-memory-mb-allocated":   0,
		"yarn-memory-mb-available":   27088,
		"yarn-memory-mb-pending":     0,
		"yarn-memory-mb-reserved":    0,
		"yarn-memory-mb-total":       27088,
		"yarn-nodes-active":          2,
		"yarn-nodes-decommissioned":  0,
		"yarn-nodes-decommissioning": 0,
		"yarn-nodes-lost":            0,
		"yarn-nodes-new":             0,
		"yarn-nodes-rebooted":        0,
		"yarn-nodes-shutdown":        0,
		"yarn-nodes-unhealthy":       0,
		"yarn-vcores-allocated":      0,
		"yarn-vcores-available":      8,
		"yarn-vcores-pending":        0,
		"yarn-vcores-reserved":       0,
		"yarn-vcores-total":          8,
	}
}

func (s *clusterControllerServer) populateSoftwareConfig(obj *pb.Cluster) {
	if obj.Config.SoftwareConfig == nil {
		obj.Config.SoftwareConfig = &pb.SoftwareConfig{}
	}

	if obj.Config.SoftwareConfig.ImageVersion == "" {
		obj.Config.SoftwareConfig.ImageVersion = "2.2.47-debian12"
	}

	if obj.Config.SoftwareConfig.Properties == nil {
		// Match some of the properties

		// Would be better to look this up in a list of GCE machine types, but we don't have it yet
		masterMemory := 4 * 1024
		if strings.HasSuffix(obj.GetConfig().GetMasterConfig().GetMachineTypeUri(), "n2-standard-2") {
			masterMemory = 8 * 1024
		}
		if strings.HasSuffix(obj.GetConfig().GetMasterConfig().GetMachineTypeUri(), "n2-standard-4") {
			masterMemory = 16 * 1024
		}
		hadoopJobHistoryHeapSize := min(4000, masterMemory/4)
		sparkDaemonMemory := min(4000, masterMemory/4)
		yarnResourceManagerHeapSize := min(4000, masterMemory/4)
		yarmTimelineserverHeapSize := min(4000, masterMemory/4)
		sparkDriverMemory := masterMemory / 4
		sparkDriverMaxResultSize := masterMemory / 8

		obj.Config.SoftwareConfig.Properties = map[string]string{
			"capacity-scheduler:yarn.scheduler.capacity.resource-calculator":          "org.apache.hadoop.yarn.util.resource.DominantResourceCalculator",
			"capacity-scheduler:yarn.scheduler.capacity.root.default.ordering-policy": "fair",
			"core:fs.gs.block.size":                                              "134217728",
			"core:fs.gs.metadata.cache.enable":                                   "false",
			"core:hadoop.ssl.enabled.protocols":                                  "TLSv1,TLSv1.1,TLSv1.2",
			"distcp:mapreduce.map.java.opts":                                     "-Xmx768m",
			"distcp:mapreduce.map.memory.mb":                                     "1024",
			"distcp:mapreduce.reduce.java.opts":                                  "-Xmx768m",
			"distcp:mapreduce.reduce.memory.mb":                                  "1024",
			"hadoop-env:HADOOP_DATANODE_OPTS":                                    "-Xmx512m",
			"hdfs:dfs.datanode.address":                                          "0.0.0.0:9866",
			"hdfs:dfs.datanode.http.address":                                     "0.0.0.0:9864",
			"hdfs:dfs.datanode.https.address":                                    "0.0.0.0:9865",
			"hdfs:dfs.datanode.ipc.address":                                      "0.0.0.0:9867",
			"hdfs:dfs.namenode.handler.count":                                    "20",
			"hdfs:dfs.namenode.http-address":                                     "0.0.0.0:9870",
			"hdfs:dfs.namenode.https-address":                                    "0.0.0.0:9871",
			"hdfs:dfs.namenode.lifeline.rpc-address":                             obj.ClusterName + "-m:8050",
			"hdfs:dfs.namenode.secondary.http-address":                           "0.0.0.0:9868",
			"hdfs:dfs.namenode.secondary.https-address":                          "0.0.0.0:9869",
			"hdfs:dfs.namenode.service.handler.count":                            "10",
			"hdfs:dfs.namenode.servicerpc-address":                               obj.ClusterName + "-m:8051",
			"mapred-env:HADOOP_JOB_HISTORYSERVER_HEAPSIZE":                       strconv.Itoa(hadoopJobHistoryHeapSize),
			"mapred:mapreduce.job.maps":                                          "21",
			"mapred:mapreduce.job.reduce.slowstart.completedmaps":                "0.95",
			"mapred:mapreduce.job.reduces":                                       "7",
			"mapred:mapreduce.jobhistory.recovery.store.class":                   "org.apache.hadoop.mapreduce.v2.hs.HistoryServerLeveldbStateStoreService",
			"mapred:mapreduce.map.cpu.vcores":                                    "1",
			"mapred:mapreduce.map.java.opts":                                     "-Xmx2708m",
			"mapred:mapreduce.map.memory.mb":                                     "3386",
			"mapred:mapreduce.reduce.cpu.vcores":                                 "1",
			"mapred:mapreduce.reduce.java.opts":                                  "-Xmx2708m",
			"mapred:mapreduce.reduce.memory.mb":                                  "3386",
			"mapred:mapreduce.task.io.sort.mb":                                   "256",
			"mapred:yarn.app.mapreduce.am.command-opts":                          "-Xmx2708m",
			"mapred:yarn.app.mapreduce.am.resource.cpu-vcores":                   "1",
			"mapred:yarn.app.mapreduce.am.resource.mb":                           "3386",
			"spark-env:SPARK_DAEMON_MEMORY":                                      strconv.Itoa(sparkDaemonMemory) + "m",
			"spark:spark.driver.maxResultSize":                                   strconv.Itoa(sparkDriverMaxResultSize) + "m",
			"spark:spark.driver.memory":                                          strconv.Itoa(sparkDriverMemory) + "m",
			"spark:spark.executor.cores":                                         "2",
			"spark:spark.executor.instances":                                     "2",
			"spark:spark.executor.memory":                                        "6157m",
			"spark:spark.executorEnv.OPENBLAS_NUM_THREADS":                       "1",
			"spark:spark.plugins.defaultList":                                    "com.google.cloud.dataproc.DataprocSparkPlugin",
			"spark:spark.scheduler.mode":                                         "FAIR",
			"spark:spark.sql.cbo.enabled":                                        "true",
			"spark:spark.sql.optimizer.runtime.bloomFilter.join.pattern.enabled": "true",
			"spark:spark.ui.port":                                                "0",
			"spark:spark.yarn.am.memory":                                         "640m",
			"yarn-env:YARN_NODEMANAGER_HEAPSIZE":                                 "1638",
			"yarn-env:YARN_RESOURCEMANAGER_HEAPSIZE":                             strconv.Itoa(yarnResourceManagerHeapSize),
			"yarn-env:YARN_TIMELINESERVER_HEAPSIZE":                              strconv.Itoa(yarmTimelineserverHeapSize),
			"yarn:yarn.nodemanager.address":                                      "0.0.0.0:8026",
			"yarn:yarn.nodemanager.resource.cpu-vcores":                          "4",
			"yarn:yarn.nodemanager.resource.memory-mb":                           "13544",
			"yarn:yarn.resourcemanager.decommissioning-nodes-watcher.decommission-if-no-shuffle-data": "true",
			"yarn:yarn.resourcemanager.nodemanager-graceful-decommission-timeout-secs":                "86400",
			"yarn:yarn.scheduler.maximum-allocation-mb":                                               "13544",
			"yarn:yarn.scheduler.minimum-allocation-mb":                                               "1",
		}
	}
}

func (s *clusterControllerServer) DeleteCluster(ctx context.Context, req *pb.DeleteClusterRequest) (*longrunningpb.Operation, error) {
	name, err := s.buildClusterName(req.ProjectId, req.Region, req.ClusterName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	now := time.Now()

	deleted := &pb.Cluster{}
	if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
		return nil, err
	}

	lroPrefix := fmt.Sprintf("projects/%s/regions/%s", name.Project.ID, name.Region)
	lroMetadata := &pb.ClusterOperationMetadata{
		ClusterName:   name.ClusterName,
		ClusterUuid:   string(deleted.ClusterUuid),
		Description:   "Delete cluster",
		OperationType: "DELETE",
		Status: &pb.ClusterOperationStatus{
			InnerState:     "PENDING",
			State:          pb.ClusterOperationStatus_PENDING,
			StateStartTime: timestamppb.New(now),
		},
	}
	return s.operations.StartLRO(ctx, lroPrefix, lroMetadata, func() (proto.Message, error) {
		// Intermediate states appear in statusHistory
		s.setOperationStatus(lroMetadata, pb.ClusterOperationStatus_RUNNING)
		s.setOperationStatus(lroMetadata, pb.ClusterOperationStatus_DONE)
		return &emptypb.Empty{}, nil
	})
}

type clusterName struct {
	Project     *projects.ProjectData
	Region      string
	ClusterName string
}

func (n *clusterName) String() string {
	return fmt.Sprintf("projects/%s/regions/%s/clusters/%s", n.Project.ID, n.Region, n.ClusterName)
}

// parseClusterName parses a string into an clusterName.
// The expected form is `projects/*/regions/*/clusters/*`.
func (s *MockService) parseClusterName(name string) (*clusterName, error) {

	tokens := strings.Split(name, "/")
	if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "regions" && tokens[4] == "clusters" {
		return s.buildClusterName(tokens[1], tokens[3], tokens[5])
	}

	return nil, status.Errorf(codes.InvalidArgument, "invalid name %q", name)
}

// buildClusterName builds a clusterName from the components.
func (s *MockService) buildClusterName(projectName, region, cluster string) (*clusterName, error) {

	project, err := s.Projects.GetProjectByID(projectName)
	if err != nil {
		return nil, err
	}

	return &clusterName{
		Project:     project,
		Region:      region,
		ClusterName: cluster,
	}, nil
}

// mutateObject updates the object; it gets the object by fqn, calls mutator, then updates the object
func mutateObject[T proto.Message](ctx context.Context, storage storage.Storage, fqn string, mutator func(obj T) error) (T, error) {
	var nilT T

	typeT := reflect.TypeOf(nilT)
	obj := reflect.New(typeT.Elem()).Interface().(T)
	if err := storage.Get(ctx, fqn, obj); err != nil {
		return nilT, err
	}

	if err := mutator(obj); err != nil {
		return nilT, err
	}

	if err := storage.Update(ctx, fqn, obj); err != nil {
		return nilT, err
	}

	return obj, nil
}
