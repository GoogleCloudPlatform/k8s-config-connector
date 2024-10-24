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

package mockmonitoring

import (
	"context"
	"sort"
	"strings"
	"time"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/monitoring/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/genproto/googleapis/api"
	metric "google.golang.org/genproto/googleapis/api/metric"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog/v2"
)

type metricService struct {
	*MockService
	pb.UnimplementedMetricServiceServer
}

func (s *metricService) GetMetricDescriptor(ctx context.Context, req *pb.GetMetricDescriptorRequest) (*metric.MetricDescriptor, error) {
	name, err := s.parseMetricDescriptorName(req.GetName())
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := &metric.MetricDescriptor{}
	if err := s.storage.Get(ctx, fqn, obj); err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, status.Errorf(codes.NotFound, "Could not find descriptor for metric '%s'.", name.MetricID)
		}
		return nil, err
	}

	// returned object does not include launchStage or metadata
	retObj := ProtoClone(obj)
	retObj.LaunchStage = api.LaunchStage_LAUNCH_STAGE_UNSPECIFIED
	retObj.Metadata = nil

	return retObj, nil
}

func populateDefaultsForMetricDescriptor(obj *metric.MetricDescriptor) {
	if obj.MonitoredResourceTypes == nil {
		obj.MonitoredResourceTypes = []string{
			"aws_ec2_instance",
			"aws_lambda_function",
			"aws_sqs_queue",
			"baremetalsolution.googleapis.com/Instance",
			"cloud_composer_environment",
			"cloud_composer_workflow",
			"cloud_dataproc_batch",
			"dataflow_job",
			"gae_instance",
			"gce_instance",
			"generic_node",
			"generic_task",
			"gke_container",
			"global",
			"k8s_cluster",
			"k8s_container",
			"k8s_node",
			"k8s_pod",
			"k8s_service",
			"prometheus_target",
		}
	}
}

func (s *metricService) CreateMetricDescriptor(ctx context.Context, req *pb.CreateMetricDescriptorRequest) (*metric.MetricDescriptor, error) {
	metricDescriptorID := req.GetMetricDescriptor().GetType()
	if metricDescriptorID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "type is required")
	}

	reqName := req.GetName() + "/metricDescriptors/" + metricDescriptorID
	name, err := s.parseMetricDescriptorName(reqName)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	obj := ProtoClone(req.MetricDescriptor)
	obj.Name = fqn

	// Creation is actually async - but without an LRO (!!!)
	go func() {
		ctx := context.Background()
		log := klog.FromContext(ctx)

		time.Sleep(1 * time.Second)

		obj = ProtoClone(obj)

		populateDefaultsForMetricDescriptor(obj)

		// Labels are reordered
		sort.Slice(obj.Labels, func(i, j int) bool {
			return obj.Labels[i].Key < obj.Labels[j].Key
		})

		if err := s.storage.Create(ctx, fqn, obj); err != nil {
			log.Error(err, "error creating metricDescriptor")
		}
	}()

	return obj, nil
}

func (s *metricService) DeleteMetricDescriptor(ctx context.Context, req *pb.DeleteMetricDescriptorRequest) (*empty.Empty, error) {
	name, err := s.parseMetricDescriptorName(req.Name)
	if err != nil {
		return nil, err
	}

	fqn := name.String()

	// Deletion is actually async - but without an LRO (!!!)
	go func() {
		ctx := context.Background()
		log := klog.FromContext(ctx)

		time.Sleep(10 * time.Second)
		deleted := &metric.MetricDescriptor{}
		if err := s.storage.Delete(ctx, fqn, deleted); err != nil {
			log.Error(err, "error deleting metricDescriptor")
		}
	}()

	return &empty.Empty{}, nil
}

type metricDescriptorName struct {
	Project  *projects.ProjectData
	MetricID string
}

func (n *metricDescriptorName) String() string {
	return "projects/" + n.Project.ID + "/metricDescriptors/" + n.MetricID
}

// parseMetricDescriptorName parses a string into a metricDescriptorName.
// The format is:
//
//	projects/[PROJECT_ID_OR_NUMBER]/metricDescriptors/[METRIC_ID]
//
// An example value of `[METRIC_ID]` is
// `"compute.googleapis.com/instance/disk/read_bytes_count"`.
func (s *MockService) parseMetricDescriptorName(name string) (*metricDescriptorName, error) {
	tokens := strings.Split(name, "/")

	// Note: this is unusual - the ID can contain slashes (which are our delimiter)
	if len(tokens) >= 4 && tokens[0] == "projects" && tokens[2] == "metricDescriptors" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &metricDescriptorName{
			Project:  project,
			MetricID: strings.Join(tokens[3:], "/"),
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
