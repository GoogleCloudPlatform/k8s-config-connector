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

package controller

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var jobTemplate string = `apiVersion: batch/v1
kind: Job
metadata:
  name: {{.Name}}
  namespace: default
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
        image: {{.ImageRegistry}}/manifests-inline:latest
        imagePullPolicy: Always
        args: ["--template", "{{.CompositionName}}", "--plan", "{{.PlanName}}", "--expander", "{{.ExpanderName}}", "--group", "{{.InputAPIGroup}}", "--version", "{{.InputAPIVersion}}", "--resource", "{{.InputAPIResource}}", "--name", "{{.InputAPIName}}", "--namespace", "{{.InputAPINamespace}}", "--path", "/expanded", "--stage", "afterExpansion"]
        volumeMounts:
        - name: expanded
          mountPath: /expanded
      - name: expand
        image: {{.ImageRegistry}}/expander-jinja2:latest
        imagePullPolicy: Always
        args: ["/inputs/template", "/inputs/values", "--format=json", "-o", "/expanded/expanded"]
        volumeMounts:
        - name: inputs
          mountPath: /inputs
        - name: expanded
          mountPath: /expanded
      initContainers:
      - name: copyin
        image: {{.ImageRegistry}}/manifests-inline:latest
        imagePullPolicy: Always
        args: ["--template", "{{.CompositionName}}", "--plan", "{{.PlanName}}", "--expander", "{{.ExpanderName}}", "--group", "{{.InputAPIGroup}}", "--version", "{{.InputAPIVersion}}", "--resource", "{{.InputAPIResource}}", "--name", "{{.InputAPIName}}", "--namespace", "{{.InputAPINamespace}}", "--path", "/inputs", "--stage", "beforeExpansion"]
        volumeMounts:
        - name: inputs
          mountPath: /inputs`

type JobFactory struct {
	InputAPIGroup     string
	InputAPIResource  string
	InputAPIName      string
	InputAPINamespace string
	InputAPIVersion   string
	CompositionName   string
	ExpanderName      string
	ImageRegistry     string
	PlanName          string
	Name              string
	logger            logr.Logger
	ctx               context.Context
	client            client.Client
	objects           []unstructured.Unstructured
	timeout           time.Duration
}

func NewJobFactory(ctx context.Context, logger logr.Logger,
	r *ExpanderReconciler,
	cr *unstructured.Unstructured, expanderName string,
	planName string, imageRegistry string) *JobFactory {
	return &JobFactory{
		InputAPIGroup:     r.InputGVK.Group,
		InputAPIVersion:   r.InputGVK.Version,
		InputAPIResource:  r.Resource,
		InputAPIName:      cr.GetName(),
		InputAPINamespace: cr.GetNamespace(),
		CompositionName:   r.Composition.Name,
		ExpanderName:      expanderName,
		PlanName:          planName,
		ImageRegistry:     imageRegistry,
		Name:              r.Composition.Name + "-" + cr.GetName() + "-" + expanderName,
		logger:            logger,
		ctx:               ctx,
		client:            r.Client,
	}
}

func (f *JobFactory) parseObjectTemplate(name string, objTemplate string) ([]byte, error) {
	var manifests bytes.Buffer

	tmpl, err := template.New(name).Parse(objTemplate)
	if err != nil {
		return nil, err
	}
	// Execute template tmpl and write to a string
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
				"namespace": "default", // TODO(barni@) FIX THIS
			},
		},
	}
}

func (j *JobFactory) getClusterRole() unstructured.Unstructured {
	return unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRole",
			"metadata": map[string]interface{}{
				"name": j.Name,
			},
			"rules": []interface{}{
				map[string]interface{}{
					"apiGroups": []interface{}{"composition.google.com"},
					"resources": []interface{}{"compositions", "contexts"},
					"verbs":     []interface{}{"get"},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{"composition.google.com"},
					"resources": []interface{}{"plans"},
					"verbs":     []interface{}{"get", "patch", "update"},
				},
				map[string]interface{}{
					"apiGroups": []interface{}{j.InputAPIGroup},
					"resources": []interface{}{j.InputAPIResource, j.InputAPIResource + "/status"},
					"verbs":     []interface{}{"get", "patch", "update"},
				},
			},
		},
	}
}

func (j *JobFactory) getClusterRoleBinding() unstructured.Unstructured {
	return unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "rbac.authorization.k8s.io/v1",
			"kind":       "ClusterRoleBinding",
			"metadata": map[string]interface{}{
				"name": j.Name,
			},
			"roleRef": map[string]interface{}{
				"apiGroup": "rbac.authorization.k8s.io",
				"kind":     "ClusterRole",
				"name":     j.Name,
			},
			"subjects": []interface{}{
				map[string]interface{}{
					"kind":      "ServiceAccount",
					"name":      j.Name,
					"namespace": "default", //TODO(barni@) FIX THIS
				},
			},
		},
	}
}

func (j *JobFactory) getJob() (unstructured.Unstructured, error) {
	return j.getUnstructuredFromTemplate("job", jobTemplate)
}

func (j *JobFactory) Create() error {
	j.objects = []unstructured.Unstructured{j.getServiceAccount(), j.getClusterRole(), j.getClusterRoleBinding()}
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
			if health == "Complete" {
				return true, nil
			}
			if health == "Failed" {
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
		case "Complete":
			if c.Status == corev1.ConditionTrue {
				logger.Info("Completed")
				return "Complete", nil
			}
		case "Failed":
			if c.Status == corev1.ConditionTrue {
				logger.Info("Failed")
				return "Failed", nil
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
