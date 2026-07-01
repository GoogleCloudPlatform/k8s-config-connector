// Copyright 2026 Google LLC
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

package bigtable

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	gcp "cloud.google.com/go/bigtable"
	pb "cloud.google.com/go/bigtable/admin/apiv2/adminpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigtable/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/tags"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(krm.BigtableGCPolicyGVK, NewGCPolicyModel)
}

func NewGCPolicyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &modelGCPolicy{config: *config}, nil
}

var _ directbase.Model = &modelGCPolicy{}

type modelGCPolicy struct {
	config config.ControllerConfig
}

func (m *modelGCPolicy) client(ctx context.Context, project, instance string) (*gcp.AdminClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.GRPCClientOptions()
	if err != nil {
		return nil, fmt.Errorf("building BigtableGCPolicy client options: %w", err)
	}
	gcpClient, err := gcp.NewAdminClient(ctx, project, instance, opts...)
	if err != nil {
		return nil, fmt.Errorf("building BigtableGCPolicy admin client: %w", err)
	}
	return gcpClient, err
}

func (m *modelGCPolicy) AdapterForObject(ctx context.Context, op *directbase.AdapterForObjectOperation) (directbase.Adapter, error) {
	u := op.GetUnstructured()
	reader := op.Reader
	obj := &krm.BigtableGCPolicy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	identity, err := obj.GetIdentity(ctx, reader)
	if err != nil {
		return nil, err
	}
	id, ok := identity.(*krm.BigtableGCPolicyIdentity)
	if !ok {
		return nil, fmt.Errorf("unexpected identity type: %T", identity)
	}

	gcpClient, err := m.client(ctx, id.Project, id.Instance)
	if err != nil {
		return nil, err
	}

	return &GCPolicyAdapter{
		id:        id,
		gcpClient: gcpClient,
		desired:   obj,
		reader:    reader,
	}, nil
}

func (m *modelGCPolicy) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type GCPolicyAdapter struct {
	id        *krm.BigtableGCPolicyIdentity
	gcpClient *gcp.AdminClient
	desired   *krm.BigtableGCPolicy
	reader    client.Reader
	actual    *pb.GcRule
}

var _ directbase.Adapter = &GCPolicyAdapter{}

func (a *GCPolicyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}

func (a *GCPolicyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("getting BigtableGCPolicy", "name", a.id)

	ti, err := a.gcpClient.TableInfo(ctx, a.id.Table)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NotFound") {
			return false, nil
		}
		return false, fmt.Errorf("getting BigtableGCPolicy %q: %w", a.id, err)
	}

	for _, fi := range ti.FamilyInfos {
		if fi.Name == a.id.ColumnFamily {
			if fi.FullGCPolicy == nil || fi.FullGCPolicy.String() == "" {
				return false, nil
			}
			pbRule, err := convertGCPolicyToProto(fi.FullGCPolicy)
			if err != nil {
				return false, fmt.Errorf("converting GCPolicy to proto: %w", err)
			}
			a.actual = pbRule
			return true, nil
		}
	}

	return false, nil
}

func (a *GCPolicyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("creating BigtableGCPolicy", "name", a.id)

	desiredPolicy, err := a.getDesiredGCPolicy()
	if err != nil {
		return err
	}

	err = a.gcpClient.SetGCPolicy(ctx, a.id.Table, a.id.ColumnFamily, desiredPolicy)
	if err != nil {
		return fmt.Errorf("creating BigtableGCPolicy %s: %w", a.id, err)
	}

	log.V(2).Info("successfully created BigtableGCPolicy", "name", a.id)
	return a.updateStatus(ctx, createOp)
}

func (a *GCPolicyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	log := klog.FromContext(ctx)
	log.V(2).Info("updating BigtableGCPolicy", "name", a.id)

	desiredPolicy, err := a.getDesiredGCPolicy()
	if err != nil {
		return err
	}

	desiredPb, err := convertGCPolicyToProto(desiredPolicy)
	if err != nil {
		return err
	}

	diffs, _, err := compareGcRule(ctx, a.actual, desiredPb)
	if err != nil {
		return err
	}

	if !diffs.HasDiff() {
		log.V(2).Info("BigtableGCPolicy is up to date", "name", a.id)
		return nil
	}

	structuredreporting.ReportDiff(ctx, diffs)

	err = a.gcpClient.SetGCPolicy(ctx, a.id.Table, a.id.ColumnFamily, desiredPolicy)
	if err != nil {
		return fmt.Errorf("updating BigtableGCPolicy %s: %w", a.id, err)
	}

	log.V(2).Info("successfully updated BigtableGCPolicy", "name", a.id)
	return a.updateStatus(ctx, updateOp)
}

func (a *GCPolicyAdapter) Delete(ctx context.Context, _ *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx)
	log.V(2).Info("deleting BigtableGCPolicy", "name", a.id)

	if a.desired.Spec.DeletionPolicy != nil && *a.desired.Spec.DeletionPolicy == "ABANDON" {
		log.V(2).Info("BigtableGCPolicy is abandoned, skipping delete on GCP", "name", a.id)
		return true, nil
	}

	err := a.gcpClient.SetGCPolicy(ctx, a.id.Table, a.id.ColumnFamily, gcp.NoGcPolicy())
	if err != nil {
		if direct.IsNotFound(err) {
			log.V(2).Info("BigtableGCPolicy table/columnFamily not found, assuming already deleted", "name", a.id)
			return true, nil
		}
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "NotFound") {
			log.V(2).Info("BigtableGCPolicy table/columnFamily not found, assuming already deleted", "name", a.id)
			return true, nil
		}
		return false, fmt.Errorf("deleting BigtableGCPolicy %s: %w", a.id, err)
	}

	log.V(2).Info("successfully deleted BigtableGCPolicy", "name", a.id)
	return true, nil
}

func (a *GCPolicyAdapter) updateStatus(ctx context.Context, op directbase.Operation) error {
	status := &krm.BigtableGCPolicyStatus{}
	return op.UpdateStatus(ctx, status, nil)
}

func (a *GCPolicyAdapter) getDesiredGCPolicy() (gcp.GCPolicy, error) {
	spec := a.desired.Spec

	if spec.GcRules != nil && *spec.GcRules != "" {
		var topLevelPolicy map[string]interface{}
		if err := json.Unmarshal([]byte(*spec.GcRules), &topLevelPolicy); err != nil {
			return nil, fmt.Errorf("unmarshaling gcRules: %w", err)
		}
		return getGCPolicyFromJSON(topLevelPolicy, true)
	}

	var policies []gcp.GCPolicy

	if len(spec.MaxAge) > 0 {
		var d time.Duration
		var err error
		if spec.MaxAge[0].Duration != nil {
			d, err = ParseDuration(*spec.MaxAge[0].Duration)
			if err != nil {
				return nil, fmt.Errorf("parsing maxAge duration %q: %w", *spec.MaxAge[0].Duration, err)
			}
		} else if spec.MaxAge[0].Days != nil {
			d = time.Hour * 24 * time.Duration(*spec.MaxAge[0].Days)
		}
		policies = append(policies, gcp.MaxAgePolicy(d))
	}

	if len(spec.MaxVersion) > 0 {
		policies = append(policies, gcp.MaxVersionsPolicy(int(spec.MaxVersion[0].Number)))
	}

	if len(policies) == 0 {
		return gcp.NoGcPolicy(), nil
	}

	if len(policies) > 1 {
		if spec.Mode == nil || *spec.Mode == "" {
			return nil, fmt.Errorf("if multiple policies are set, mode can't be empty")
		}
		mode := strings.ToUpper(*spec.Mode)
		switch mode {
		case "UNION":
			return gcp.UnionPolicy(policies...), nil
		case "INTERSECTION":
			return gcp.IntersectionPolicy(policies...), nil
		default:
			return nil, fmt.Errorf("unknown mode: %q", *spec.Mode)
		}
	}

	return policies[0], nil
}

func getGCPolicyFromJSON(inputPolicy map[string]interface{}, isTopLevel bool) (gcp.GCPolicy, error) {
	policy := []gcp.GCPolicy{}

	rulesObj, ok := inputPolicy["rules"]
	if isTopLevel && !ok {
		return nil, fmt.Errorf("invalid nested policy, need `rules`")
	}

	if ok {
		rules, arrOk := rulesObj.([]interface{})
		if !arrOk {
			return nil, fmt.Errorf("`rules` must be array")
		}
		for _, p := range rules {
			childPolicy, mapOk := p.(map[string]interface{})
			if !mapOk {
				return nil, fmt.Errorf("child rule must be a map")
			}

			if childPolicy["max_age"] != nil {
				maxAge := childPolicy["max_age"].(string)
				duration, err := ParseDuration(maxAge)
				if err != nil {
					return nil, fmt.Errorf("invalid duration string: %v", maxAge)
				}
				policy = append(policy, gcp.MaxAgePolicy(duration))
			}

			if childPolicy["max_version"] != nil {
				var version int
				switch v := childPolicy["max_version"].(type) {
				case float64:
					version = int(v)
				case int:
					version = v
				case int64:
					version = int(v)
				case string:
					parsed, err := strconv.Atoi(v)
					if err != nil {
						return nil, fmt.Errorf("invalid max_version string: %v", v)
					}
					version = parsed
				case json.Number:
					parsed, _ := v.Int64()
					version = int(parsed)
				default:
					return nil, fmt.Errorf("invalid type for max_version: %T", childPolicy["max_version"])
				}
				policy = append(policy, gcp.MaxVersionsPolicy(version))
			}

			if childPolicy["mode"] != nil {
				n, err := getGCPolicyFromJSON(childPolicy, false)
				if err != nil {
					return nil, err
				}
				policy = append(policy, n)
			}
		}
	}

	mode, hasMode := inputPolicy["mode"].(string)
	if hasMode {
		switch strings.ToLower(mode) {
		case "union":
			return gcp.UnionPolicy(policy...), nil
		case "intersection":
			return gcp.IntersectionPolicy(policy...), nil
		default:
			if len(policy) > 0 {
				return policy[0], nil
			}
		}
	} else if len(policy) > 0 {
		return policy[0], nil
	}

	return gcp.NoGcPolicy(), nil
}

func convertGCPolicyToProto(gc gcp.GCPolicy) (*pb.GcRule, error) {
	if gc == nil {
		return nil, nil
	}
	switch gcp.GetPolicyType(gc) {
	case gcp.PolicyMaxAge:
		durationString := gc.(gcp.MaxAgeGCPolicy).GetDurationString()
		d, err := ParseDuration(durationString)
		if err != nil {
			return nil, fmt.Errorf("parsing max_age duration string %q: %w", durationString, err)
		}
		return &pb.GcRule{
			Rule: &pb.GcRule_MaxAge{
				MaxAge: durationpb.New(d),
			},
		}, nil
	case gcp.PolicyMaxVersion:
		v := int32(gc.(gcp.MaxVersionsGCPolicy))
		return &pb.GcRule{
			Rule: &pb.GcRule_MaxNumVersions{
				MaxNumVersions: v,
			},
		}, nil
	case gcp.PolicyUnion:
		var rules []*pb.GcRule
		for _, child := range gc.(gcp.UnionGCPolicy).Children {
			childRule, err := convertGCPolicyToProto(child)
			if err != nil {
				return nil, err
			}
			rules = append(rules, childRule)
		}
		return &pb.GcRule{
			Rule: &pb.GcRule_Union_{
				Union: &pb.GcRule_Union{
					Rules: rules,
				},
			},
		}, nil
	case gcp.PolicyIntersection:
		var rules []*pb.GcRule
		for _, child := range gc.(gcp.IntersectionGCPolicy).Children {
			childRule, err := convertGCPolicyToProto(child)
			if err != nil {
				return nil, err
			}
			rules = append(rules, childRule)
		}
		return &pb.GcRule{
			Rule: &pb.GcRule_Intersection_{
				Intersection: &pb.GcRule_Intersection{
					Rules: rules,
				},
			},
		}, nil
	}
	return nil, nil
}

func compareGcRule(ctx context.Context, actual, desired *pb.GcRule) (*structuredreporting.Diff, *fieldmaskpb.FieldMask, error) {
	if actual == nil {
		clonedDesired := proto.CloneOf(desired)
		diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), nil)
		return diffs, updateMask, err
	}
	clonedDesired := proto.CloneOf(desired)
	clonedActual := proto.CloneOf(actual)
	diffs, updateMask, err := tags.DiffForTopLevelFields(ctx, clonedDesired.ProtoReflect(), clonedActual.ProtoReflect())
	if err != nil {
		return nil, nil, err
	}
	return diffs, updateMask, nil
}
