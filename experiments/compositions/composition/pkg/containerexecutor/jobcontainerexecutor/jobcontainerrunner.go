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

package jobcontainerexecutor

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/safetext/yamltemplate"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ConditionTypeComplete string = "Complete"
	ConditionTypeFailed   string = "Failed"
)

var jobTemplate string = `apiVersion: batch/v1
kind: Job
metadata:
  name: {{.Name}}
  namespace: {{.CompositionNamespace}}
  labels:
    compositionname: {{.CompositionName}}
    name: {{.InputAPIName}}
    namespace: {{.InputAPINamespace}}
spec:
  backoffLimit: 2 # retry twice at most
  ttlSecondsAfterFinished: 300 # delete after 60 seconds
  template:
    metadata:
      name: expander
      labels:
        compositionname: {{.CompositionName}}
        name: {{.InputAPIName}}
        namespace: {{.InputAPINamespace}}
    spec:
      restartPolicy: Never
      serviceAccountName: {{.Name}}
      terminationGracePeriodSeconds: 10
      volumes:
      - name: inputs
        emptyDir: {}
      - name: expanded
        emptyDir: {}
      containers:
      - name: copyout
        image: {{.ImageRegistry}}/manifests-inline:v0.0.1
        args:
        - "--template"
        - "{{.CompositionName}}"
        - "--plan"
        - "{{.PlanName}}"
        - "--expander"
        - "{{.ExpanderName}}"
        - "--group"
        - "{{.InputAPIGroup}}"
        - "--version"
        - "{{.InputAPIVersion}}"
        - "--resource"
        - "{{.InputAPIResource}}"
        - "--name"
        - "{{.InputAPIName}}"
        - "--namespace"
        - "{{.InputAPINamespace}}"
        - "--path"
        - "/expanded"
        - "--stage"
        - "afterExpansion"
        volumeMounts:
        - name: expanded
          mountPath: /expanded
      - name: expand
        image: {{.ExpanderImage}}
        args: ["/inputs/template", "/inputs/values", "--format=json", "-o", "/expanded/expanded"]
        volumeMounts:
        - name: inputs
          mountPath: /inputs
        - name: expanded
          mountPath: /expanded
      initContainers:
      - name: copyin
        image: {{.ImageRegistry}}/manifests-inline:v0.0.1
        args:
        - "--template"
        - "{{.CompositionName}}"
        - "--plan"
        - "{{.PlanName}}"
        - "--expander"
        - "{{.ExpanderName}}"
        - "--group"
        - "{{.InputAPIGroup}}"
        - "--version"
        - "{{.InputAPIVersion}}"
        - "--resource"
        - "{{.InputAPIResource}}"
        - "--name"
        - "{{.InputAPIName}}"
        - "--namespace"
        - "{{.InputAPINamespace}}"
        - "--path"
        - "/inputs"
        - "--stage"
        - "beforeExpansion"
        volumeMounts:
        - name: inputs
          mountPath: /inputs`

type JobFactory struct {
	InputAPIGroup        string
	InputAPIResource     string
	InputAPIName         string
	InputAPINamespace    string
	InputAPIVersion      string
	CompositionName      string
	CompositionNamespace string
	ExpanderImage        string
	ExpanderName         string
	ImageRegistry        string
	PlanName             string
	Name                 string
	logger               logr.Logger
	ctx                  context.Context
	client               client.Client
	objects              []unstructured.Unstructured
	timeout              time.Duration
}

func NewJobFactory(ctx context.Context, logger logr.Logger, client client.Client,
	inputGVK schema.GroupVersionKind, inputGVR schema.GroupVersionResource,
	compositionName string, compositionNamespace string,
	cr *unstructured.Unstructured, expanderName string, expanderImage string,
	planName string, imageRegistry string) *JobFactory {
	return &JobFactory{
		InputAPIGroup:        inputGVK.Group,
		InputAPIVersion:      inputGVK.Version,
		InputAPIResource:     inputGVR.Resource,
		InputAPIName:         cr.GetName(),
		InputAPINamespace:    cr.GetNamespace(),
		CompositionName:      compositionName,
		CompositionNamespace: compositionNamespace,
		ExpanderName:         expanderName,
		ExpanderImage:        expanderImage,
		PlanName:             planName,
		ImageRegistry:        imageRegistry,
		Name:                 compositionName + "-" + cr.GetName() + "-" + expanderName,
		logger:               logger,
		ctx:                  ctx,
		client:               client,
	}
}

func (f *JobFactory) parseObjectTemplate(name string, objTemplate string) ([]byte, error) {
	var manifests bytes.Buffer

	tmpl, err := yamltemplate.New(name).Parse(objTemplate)
	if err != nil {
		return nil, err
	}
	// Execute yamltemplate tmpl and write to a string
	err = tmpl.Execute(&manifests, *f)
	if err != nil {
		return nil, err
	}
	return manifests.Bytes(), nil
}

func (j *JobFactory) getUnstructuredFromTemplate(name, objTemplate string) (unstructured.Unstructured, error) {
	out := unstructured.Unstructured{}
	logger := j.logger.WithName(name)
	raw, err := j.parseObjectTemplate(name, objTemplate)
	if err != nil {
		logger.Error(err, "Unable to parse manifest")
		return out, err
	}
	raw = bytes.TrimSpace(raw)
	if err := k8syaml.Unmarshal(raw, &out); err != nil {
		logger.Error(err, "Unable to marshall")
		return out, err
	}

	return out, nil
}

func (f *JobFactory) GetManifests() ([]byte, error) {
	return f.parseObjectTemplate("all", jobTemplate)
}

// Get a Service account kubernetes object for the Job
func (j *JobFactory) getServiceAccount() unstructured.Unstructured {
	return unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ServiceAccount",
			"metadata": map[string]interface{}{
				"name":      j.Name,
				"namespace": j.CompositionNamespace,
			},
		},
	}
}

func (j *JobFactory) getCompositionRole() unstructured.Unstructured {
	return unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "composition-" + j.Name,
				"namespace": j.CompositionNamespace,
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{"composition.google.com"},
					"resources": []interface{}{"compositions"},
					"verbs":     []interface{}{"get"},
				},
			},
		},
	}
}

func (j *JobFactory) getInputAPIRule() unstructured.Unstructured {
	return unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "Role",
			"metadata": map[string]interface{}{
				"name":      "inputapi-" + j.Name,
				"namespace": j.InputAPINamespace,
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{"composition.google.com"},
					"resources": []interface{}{"contexts"},
					"verbs":     []interface{}{"get"},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{j.InputAPIGroup},
					"resources": []interface{}{j.InputAPIResource, j.InputAPIResource + "/status"},
					"verbs":     []interface{}{"get", "patch", "update"},
				},
				// TODO (barney-s@) : Move this to composition namespace
				map[string]interface{}{
					"apiGroups": []interface{}{"composition.google.com"},
					"resources": []interface{}{"plans"},
					"verbs":     []interface{}{"get", "patch", "update"},
				},
			},
		},
	}
}

func (j *JobFactory) getCompositionRoleBinding() unstructured.Unstructured {
	return unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "composition-" + j.Name,
				"namespace": j.CompositionNamespace,
			},
			"roleRef": map[string]interface{}{
				"apiGroup":  "rbac.authorization.k8s.io",
				"kind":      "Role",
				"name":      "composition-" + j.Name,
				"namespace": j.CompositionNamespace,
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      j.Name,
					"namespace": j.CompositionNamespace,
				},
			},
		},
	}
}

func (j *JobFactory) getInputAPIRoleBinding() unstructured.Unstructured {
	return unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "RoleBinding",
			"metadata": map[string]interface{}{
				"name":      "inputapi-" + j.Name,
				"namespace": j.InputAPINamespace,
			},
			"roleRef": map[string]interface{}{
				"apiGroup":  "rbac.authorization.k8s.io",
				"kind":      "Role",
				"name":      "inputapi-" + j.Name,
				"namespace": j.InputAPINamespace,
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      j.Name,
					"namespace": j.CompositionNamespace,
				},
			},
		},
	}
}

func (j *JobFactory) getJob() (unstructured.Unstructured, error) {
	return j.getUnstructuredFromTemplate("job", jobTemplate)
}

func (j *JobFactory) Create() error {
	j.objects = []unstructured.Unstructured{
		j.getServiceAccount(),
		j.getCompositionRole(),
		j.getInputAPIRule(),
		j.getCompositionRoleBinding(),
		j.getInputAPIRoleBinding(),
	}
	job, err := j.getJob()
	if err != nil {
		j.logger.Error(err, "failed to marshall a job object")
		return err
	}
	j.objects = append(j.objects, job)

	for _, obj := range j.objects {
		objStr := fmt.Sprintf("%s/%s", obj.GroupVersionKind().Group, obj.GetName())
		if err := j.client.Create(j.ctx, &obj); err != nil {
			if apierrors.IsAlreadyExists(err) {
				j.logger.Info("Object already exists: " + objStr)
				continue
			}
			j.logger.Error(err, "Failed to create Object: "+objStr)
			return err
		}
		j.logger.Info("Created Object: " + objStr)
	}
	return nil
}

func (j *JobFactory) CleanUp() {
	for _, obj := range j.objects {
		objStr := fmt.Sprintf("%s/%s", obj.GroupVersionKind().Kind, obj.GetName())
		// delete with propagation true
		propagationPolicy := metav1.DeletePropagationBackground
		if err := j.client.Delete(j.ctx, &obj,
			&client.DeleteOptions{
				PropagationPolicy: &propagationPolicy,
			}); err != nil {
			j.logger.Error(err, "Failed to delete object: "+objStr)
		}
		j.logger.Info("Deleted Object: " + objStr)
	}
}

func (j *JobFactory) Wait() (bool, error) {
	// Not elegant !!
	job := j.objects[len(j.objects)-1]
	j.timeout = time.Second * 60

	logger := j.logger.WithName("Job").WithName(job.GetName())
	nn := types.NamespacedName{Name: job.GetName(), Namespace: job.GetNamespace()}

	ctx, cancel := context.WithTimeout(j.ctx, j.timeout)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			if err := j.client.Get(j.ctx, nn, &job); err != nil {
				logger.Error(err, "failed to get object")
				time.Sleep(5 * time.Second)
				continue
			}
			health, err := j.jobStatus(&job)
			if err != nil {
				logger.Error(err, "failed computing job status")
				time.Sleep(5 * time.Second)
				continue
			}
			if health == ConditionTypeComplete {
				return true, nil
			}
			if health == ConditionTypeFailed {
				return false, nil
			}
			time.Sleep(5 * time.Second)
		}
	}
}

func (j *JobFactory) jobStatus(u *unstructured.Unstructured) (string, error) {
	logger := j.logger.WithName("Job").WithName(u.GetName())
	obj := u.UnstructuredContent()

	succeeded := status.GetIntField(obj, ".status.succeeded", 0)
	active := status.GetIntField(obj, ".status.active", 0)
	failed := status.GetIntField(obj, ".status.failed", 0)
	starttime := status.GetStringField(obj, ".status.startTime", "")

	// Conditions
	// https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/job/utils.go#L24
	objc, err := status.GetObjectWithConditions(obj)
	if err != nil {
		return "", err
	}
	for _, c := range objc.Status.Conditions {
		switch c.Type {
		case ConditionTypeComplete:
			if c.Status == corev1.ConditionTrue {
				logger.Info("Completed")
				return ConditionTypeComplete, nil
			}
		case ConditionTypeFailed:
			if c.Status == corev1.ConditionTrue {
				logger.Info("Failed")
				return ConditionTypeFailed, nil
			}
		}
	}

	// replicas
	if starttime == "" {
		logger.Info("Not Started")
		return "NotStarted", nil
	}

	logstr := fmt.Sprintf("In Progress. success:%d, active: %d, failed: %d", succeeded, active, failed)
	logger.Info(logstr)
	return "InProgress", nil
}
