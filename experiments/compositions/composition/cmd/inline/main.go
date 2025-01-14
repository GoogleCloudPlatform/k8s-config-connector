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

package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	"tailscale.com/atomicfile"
)

var planGVR schema.GroupVersionResource = schema.GroupVersionResource{
	Group:    "composition.google.com",
	Version:  "v1alpha1",
	Resource: "plans",
}

var compositionGVR schema.GroupVersionResource = schema.GroupVersionResource{
	Group:    "composition.google.com",
	Version:  "v1alpha1",
	Resource: "compositions",
}

var contextGVR schema.GroupVersionResource = schema.GroupVersionResource{
	Group:    "composition.google.com",
	Version:  "v1alpha1",
	Resource: "contexts",
}

type InlineSyncer struct {
	dynamicClient dynamic.Interface
	ctx           context.Context
	gvr           schema.GroupVersionResource
	name          string
	namespace     string
	templateName  string
	expanderName  string
	planName      string
	path          string
	timeout       int
}

func (s *InlineSyncer) useConfig(config *rest.Config) error {
	var err error
	s.dynamicClient, err = dynamic.NewForConfig(config)
	return err
}

func (s *InlineSyncer) getObject(namespace, name string,
	gvr schema.GroupVersionResource) (*unstructured.Unstructured, error) {
	if namespace != "" {
		return s.dynamicClient.Resource(gvr).Namespace(namespace).Get(s.ctx, name, metav1.GetOptions{})
	} else {
		return s.dynamicClient.Resource(gvr).Namespace("default").Get(s.ctx, name, metav1.GetOptions{})
	}

}

func (s *InlineSyncer) CopyToFileSystem() error {
	// TODO (barni@): assuming name is fixed
	contextObj, err := s.getObject(s.namespace, "context", contextGVR)
	if err != nil {
		return fmt.Errorf("failed to get context object, error: %v", err)
	}
	crdv, err := s.getObject(s.namespace, s.name, s.gvr)
	if err != nil {
		return fmt.Errorf("failed to get values object, error: %v", err)
	}
	crdt, err := s.getObject("", s.templateName, compositionGVR)
	if err != nil {
		return fmt.Errorf("failed to get template object, error: %v", err)
	}
	crdp, err := s.getObject(s.namespace, s.planName, planGVR)
	if err != nil {
		return fmt.Errorf("failed to get plan object, error: %v", err)
	}

	// Get crdt.spec.expanders[.name=name].template
	expanders, found, err := unstructured.NestedSlice(crdt.Object, "spec", "expanders")
	if err != nil {
		return fmt.Errorf("composition .spec.expanders was not a slice: %v", err)
	}
	if !found {
		return fmt.Errorf("composition .spec.expanders not found")
	}

	var template string
	for index := range expanders {
		if expander, ok := expanders[index].(map[string]interface{}); ok {
			if name, ok := expander["name"].(string); ok {
				if name != s.expanderName {
					continue
				}
				if tmpl, ok := expander["template"].(string); ok {
					template = tmpl
					break
				}
			}
		}
	}

	if template == "" {
		return fmt.Errorf("template not found for expander: %s", s.expanderName)
	}

	// Get crdp.spec.stages[name]
	stages, found, err := unstructured.NestedMap(crdp.Object, "spec", "stages")
	if err != nil {
		return fmt.Errorf("Plan CR .spec.stages was not a map: %v", err)
	}
	if !found {
		klog.Infof("Input API .spec.stages not found")
		stages = map[string]interface{}{}
	}
	// Get Values for the expander if present in crdv
	jsonString := ""
	if _, ok := stages[s.expanderName]; ok {
		if stage, ok := stages[s.expanderName].(map[string]interface{}); ok {
			if v, ok := stage["values"].(string); ok {
				jsonString = v
				klog.Infof("values present: %s", jsonString)
			}
		}
	}

	expanderValues := map[string]interface{}{}
	if jsonString != "" {
		// Unmarshal the JSON string into the map
		err := json.Unmarshal([]byte(jsonString), &expanderValues)
		if err != nil {
			return fmt.Errorf("Error unmarshalling expander jsonValues: %v", err)
		}
	}

	// Construct the values object from crdv and context
	valuesObj := map[string]interface{}{
		"context":      contextObj.Object,
		s.gvr.Resource: crdv.Object,
		"values":       expanderValues,
	}

	// marshall values
	values, err := json.Marshal(valuesObj)
	if err != nil {
		return fmt.Errorf("failed to marshal values: %v", err)
	}

	// Write values to file
	// https://github.com/golang/go/issues/56174
	err = atomicfile.WriteFile(s.path+"/values", values, 0644)
	if err != nil {
		return fmt.Errorf("failed to write values to file: %v", err)
	}

	// Write template to file
	err = atomicfile.WriteFile(s.path+"/template", []byte(template), 0644)
	if err != nil {
		return fmt.Errorf("failed to write template to file: %v", err)
	}

	return nil
}

func (s *InlineSyncer) CopyFromFileSystem() error {
	crdp, err := s.getObject(s.namespace, s.planName, planGVR)
	if err != nil {
		return fmt.Errorf("failed to get Plan object, error: %v", err)
	}

	// wait for /expanded/expanded to be created and then read it
	for i := 0; i < s.timeout; i++ {
		if _, err := os.Stat(s.path + "/expanded"); err == nil {
			break
		}
		// sleep for 1 second
		time.Sleep(time.Duration(1) * time.Second)
	}

	manifests, err := os.ReadFile(s.path + "/expanded")
	if err != nil {
		return fmt.Errorf("failed to read expanded file, error: %v", err)
	}

	stageObj := map[string]interface{}{
		"manifest": string(manifests),
	}
	// Get crdp.spec.stages[name]
	stages, found, err := unstructured.NestedMap(crdp.Object, "spec", "stages")
	if err != nil {
		return fmt.Errorf("Plan CR .spec.stages was not a map: %v", err)
	}
	if !found {
		klog.Infof("Input API .spec.stages not found")
		stages = map[string]interface{}{}
	}

	// Get Values for the stage if present in crdp
	if _, ok := stages[s.expanderName]; ok {
		if stage, ok := stages[s.expanderName].(map[string]interface{}); ok {
			if v, ok := stage["values"].(string); ok {
				stageObj["values"] = v
			}
		}
	}

	stages[s.expanderName] = stageObj
	err = unstructured.SetNestedMap(crdp.Object, stages, "spec", "stages")
	if err != nil {
		return fmt.Errorf("failed to set .spec.stages field, error: %v", err)
	}

	// Using dynamic client update the crdp
	_, err = s.dynamicClient.Resource(planGVR).Namespace(s.namespace).Update(s.ctx, crdp, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update Plan CR: %v", err)
	}
	return nil
}

func main() {
	group := flag.String("group", "", "")
	version := flag.String("version", "", "")
	resource := flag.String("resource", "", "")
	namespace := flag.String("namespace", "", "values object namespace")
	name := flag.String("name", "", "values object name")
	templateName := flag.String("template", "", "Composition object name")
	planName := flag.String("plan", "", "Plan object name")
	expanderName := flag.String("expander", "", "Composition expander name")
	timeout := flag.Int("timeout", 60, "timeout waiting for expanded file")

	path := flag.String("path", "", "path for manifests")
	stage := flag.String("stage", "", "allowed values: beforeExpansion|afterExpansion")
	flag.Parse()

	if *namespace == "" {
		klog.Fatalf("namespace is required")
	}

	if *name == "" {
		klog.Fatalf("name is required")
	}

	if *path == "" {
		klog.Fatalf("path is required")
	}

	if *resource == "" {
		klog.Fatalf("resource is required")
	}

	if *group == "" {
		klog.Fatalf("group is required")
	}

	if *version == "" {
		klog.Fatalf("version is required")
	}

	if *stage != "beforeExpansion" && *stage != "afterExpansion" {
		klog.Fatalf("allowed stages: beforeExpansion|afterExpansion")
	}

	if *templateName == "" {
		klog.Fatalf("template is required")
	}

	if *expanderName == "" {
		klog.Fatalf("expander is required")
	}

	if *planName == "" {
		klog.Fatalf("plan is required")
	}

	klog.Infof("template: %s, stage: %s, plan:%s", *templateName, *stage, *planName)
	klog.Infof("group: %s, version: %s, resource: %s, namespace: %s, name: %s",
		*group, *version, *resource, *namespace, *name)
	klog.Infof("path: %s", *path)

	syncer := &InlineSyncer{
		ctx: context.Background(),
	}

	syncer.gvr = schema.GroupVersionResource{Group: *group, Version: *version, Resource: *resource}
	syncer.path = *path
	syncer.templateName = *templateName
	syncer.expanderName = *expanderName
	syncer.planName = *planName
	syncer.namespace = *namespace
	syncer.name = *name
	syncer.timeout = *timeout

	// Get dynamic client.

	config := ctrl.GetConfigOrDie()
	err := syncer.useConfig(config)
	if err != nil {
		klog.Fatalf("failed to get dynamic client error: %v", err)
	}

	if *stage == "beforeExpansion" {
		err := syncer.CopyToFileSystem()
		if err != nil {
			klog.Fatalf("failed to copy template to file system, error: %v", err)
		}
	} else {
		err := syncer.CopyFromFileSystem()
		if err != nil {
			klog.Fatalf("failed to copy template from file system, error: %v", err)
		}
	}
}
