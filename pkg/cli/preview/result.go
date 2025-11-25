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

package preview

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type GKNSummary struct {
	// number of resouces successfully reconciled by TF reconciler
	tfCount *int
	// number of resouces successfully reconciled by direct reconciler
	directCount *int
	// number of resouces successfully reconciled by DCL reconciler
	dclCount *int
	// number of resouces successfully reconciled by custom reconciler
	customCount *int
}

func (s GKNSummary) PrettyString(total int) string {
	tf := "NA"
	tfBad := "NA"
	if s.tfCount != nil {	
		tf = fmt.Sprintf("%d", *s.tfCount)
		tfBad = fmt.Sprintf("%d", total-*s.tfCount)
	}
	direct := "NA"
	directBad := "NA"
	if s.directCount != nil {
		direct = fmt.Sprintf("%d", *s.directCount)
		directBad = fmt.Sprintf("%d", total-*s.directCount)
	}
	dcl := "NA"
	dclBad := "NA"
	if s.dclCount != nil {
		dcl = fmt.Sprintf("%d", *s.dclCount)
		dclBad = fmt.Sprintf("%d", total-*s.dclCount)
	}
	custom := "NA"
	customBad := "NA"
	if s.customCount != nil {
		custom = fmt.Sprintf("%d", *s.customCount)
		customBad = fmt.Sprintf("%d", total-*s.customCount)
	}
	return fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", tf, tfBad, direct, directBad, dcl, dclBad, custom, customBad)
}
type GKN struct {
	Group     string
	Kind      string
	Namespace string
}

type PreviewSummary struct {
	// namespace -> group -> kind -> GKNSummary
	// Count the number of resources successfully reconciled by each reconciler type
	summary  map[string]map[string]map[string]*GKNSummary
	// GKNN -> controller type -> []string
	// reports store the GKNN and controller type that has issues
	reports  map[GKNN]map[k8s.ReconcilerType][]string
	// namespace -> group -> kind -> count
	// Count the total number of resources in each namespace, group, kind
	gknCount map[string]map[string]map[string]int
}

// add result to the summary
func (p *PreviewSummary) addResult(gknn GKNN, r ReconcileResult) {
	if p.summary[gknn.Namespace] == nil {
		p.summary[gknn.Namespace] = make(map[string]map[string]*GKNSummary)
	}
	if p.summary[gknn.Namespace][gknn.Group] == nil {
		p.summary[gknn.Namespace][gknn.Group] = make(map[string]*GKNSummary)
	}
	summary := p.summary[gknn.Namespace][gknn.Group][gknn.Kind]
	if summary == nil {
		summary = &GKNSummary{}
		p.summary[gknn.Namespace][gknn.Group][gknn.Kind] = summary
	}

	var countPtr **int
	switch r.controllerType {
	case k8s.ReconcilerTypeTerraform:
		countPtr = &summary.tfCount
	case k8s.ReconcilerTypeDCL:
		countPtr = &summary.dclCount
	case k8s.ReconcilerTypeDirect:
		countPtr = &summary.directCount
	default:
		countPtr = &summary.customCount
	}

	if r.good {
		inc(countPtr)
	} else {
		if p.reports[gknn] == nil {
			p.reports[gknn] = make(map[k8s.ReconcilerType][]string)
		}
		// Ensure the counter is initialized even if we don't increment it (so it shows up as 0 instead of NA)
		if *countPtr == nil {
			*countPtr = new(int)
		}
		p.reports[gknn][r.controllerType] = append(p.reports[gknn][r.controllerType], r.detail...)
	}
}

// increment the counter if it is not nil
// if it is nil, initialize it to 1
func inc(ptr **int) {
	if *ptr == nil {
		*ptr = new(int)
	}
	**ptr++
}

type ReconcileResult struct {
	controllerType k8s.ReconcilerType
	good           bool
	detail         []string
}

func (r *Recorder) newPreviewSummary() *PreviewSummary {
	summary := &PreviewSummary{
		// ns -> group -> kind -> GKNSummary
		summary:  make(map[string]map[string]map[string]*GKNSummary),
		reports:  make(map[GKNN]map[k8s.ReconcilerType][]string),
		// ns -> group -> kind -> count
		gknCount: make(map[string]map[string]map[string]int),
	}

	for gkn, count := range r.gknCount {
		if summary.gknCount[gkn.Namespace] == nil {
			summary.gknCount[gkn.Namespace] = make(map[string]map[string]int)
		}
		if summary.gknCount[gkn.Namespace][gkn.Group] == nil {
			summary.gknCount[gkn.Namespace][gkn.Group] = make(map[string]int)
		}
		summary.gknCount[gkn.Namespace][gkn.Group][gkn.Kind] = count
	}

	for gknn, info := range r.objects {
		results := ParseEventInfo(info)
		for _, result := range results {
			summary.addResult(gknn, result)
		}
	}
	return summary
}

func ParseEventInfo(info *objectInfo) []ReconcileResult {
	good := true
	detail := []string{}
	var t k8s.ReconcilerType
	result := []ReconcileResult{}
	for _, event := range info.events {
		switch event.eventType {
		case EventTypeDiff:
			detail = append(detail, "KRM diff detected:")
			if event.diff != nil && event.diff.Fields != nil {
				diffFields := []string{}
				for _, field := range event.diff.Fields {
					if !reflect.DeepEqual(field.Old, field.New) {
						diffFields = append(diffFields, field.ID)
					}
				}
				detail = append(detail, fmt.Sprintf("  Diff Fields: %s", diffFields))
			}
		// reset recocile result and set reconciler type
		case EventTypeReconcileStart:
			good = true
			detail = []string{}
			t = event.reconcilerType
		case EventTypeReconcileEnd:
			t = event.reconcilerType
			result = append(result, ReconcileResult{
				good:           good,
				detail:         detail,
				controllerType: t,
			})
		case EventTypeKubeAction:
			// Ignore kubeaction for now. Mostly status update.

		case EventTypeGCPAction:
			// The method is POST but it is actually a read-only call.
			if strings.Contains(event.gcpAction.url, "getIamPolicy") {
				continue
			}
			good = false
			detail = append(detail, "GCP action detected:")
			detail = append(detail, fmt.Sprintf("  Method: %s", event.gcpAction.method))
			detail = append(detail, fmt.Sprintf("  Url: %s", event.gcpAction.url))
			detail = append(detail, fmt.Sprintf("  Object: %+v", event.gcpAction.body))
		default:
			good = false
		}
	}
	if len(info.events) > 0 && info.events[len(info.events)-1].eventType != EventTypeReconcileEnd {
		result = append(result, ReconcileResult{
			good:           false,
			detail:         detail,
			controllerType: t,
		})
	}
	return result
}

// ExportObjectsEvent writes all captured GKNN and its event to filename.
func (r *Recorder) ExportDetailObjectsEvent(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", filename, err)
	}
	fmt.Fprintf(f, "Total number of resources: %d\n", len(r.objects))
	fmt.Fprintf(f, "Number of resources that have not been fully reconciled: %d\n", r.RemainResourcesCount)
	if r.RemainResourcesCount != 0 {
		fmt.Fprintln(f, "Known Resource that was not reconciled")
		for gknn := range r.ReconciledResources {
			if !r.GKNNDoneReconcile(gknn) {
				fmt.Fprintln(f, "-----------------------------------------------------------------")
				fmt.Fprintf(f, "Not reconciled object %+v\n", gknn)
			}
		}
	}
	for gknn, info := range r.objects {
		fmt.Fprintln(f, "-----------------------------------------------------------------")
		fmt.Fprintf(f, "object %+v\n", gknn)
		for _, event := range info.events {
			switch event.eventType {
			case EventTypeDiff:
				fmt.Fprintf(f, "  diff %+v\n", event.diff)

			case EventTypeReconcileStart:
				fmt.Fprintf(f, "  Reconcile Start, reconcilerType: %s\n", event.reconcilerType)

			case EventTypeReconcileEnd:
				fmt.Fprintf(f, "  Reconcile End, reconcilerType: %s\n", event.reconcilerType)

			case EventTypeKubeAction:
				fmt.Fprintf(f, "  kubeAction %+v\n", event.kubeAction)

			case EventTypeGCPAction:
				fmt.Fprintf(f, "  gcpAction %+v\n", event.gcpAction)

			default:
				fmt.Fprintf(f, "  unknown event: %+v\n", event)
			}
		}
	}
	return f.Close()
}

func (r *Recorder) SummaryReport(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", filename, err)
	}
	summary := r.newPreviewSummary()
	w := tabwriter.NewWriter(f, 0, 0, 3, ' ', 0)
	controllerType := resourceconfig.LoadConfig()
	fmt.Fprintln(w, "Namespace\tGroup\tKind\tTotal\tDefault Controller Type\tTF-Good\tTF-Bad\tDirect-Good\tDirect-Bad\tDCL-Good\tDCL-Bad\tCustom-Good\tCustom-Bad")
	for ns := range summary.gknCount {
		for group := range summary.gknCount[ns] {
			for kind := range summary.gknCount[ns][group] {
				total := summary.gknCount[ns][group][kind]
				s := summary.summary[ns][group][kind]
				gvk := schema.GroupVersionKind{
					Group:   group,
					Kind:    kind,
				}
				controllerConfig, err := controllerType.GetControllersForGVK(gvk)
				if err != nil {
					return fmt.Errorf("error getting controllers for GVK %v: %w", gvk, err)
				}
				fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\t%s\n", ns, group, kind, total, controllerConfig.DefaultController, s.PrettyString(total))
			}
		}
	}
	w.Flush()
	if err = f.Close(); err != nil {
		return fmt.Errorf("error closing file %q: %w", filename, err)
	}
	if len(summary.reports) == 0 && r.RemainResourcesCount == 0 {
		return nil
	}
	detailFileName := fmt.Sprintf("%s-detail", filename)
	detailFile, err := os.Create(detailFileName)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", detailFileName, err)
	}
	if r.RemainResourcesCount > 0 {
		fmt.Fprintln(detailFile, "Resources that has not fully reconciled:")
		for gknn := range r.ReconciledResources {
			for controllerType, done := range r.ReconciledResources[gknn] {
				if done {
					continue
				}
				fmt.Fprintln(detailFile, "-----------------------------------------------------------------")
				fmt.Fprintf(detailFile, "Group: %s, Kind: %s, Namespace: %s, Name: %s, Controller Type: %s\n", gknn.Group, gknn.Kind, gknn.Namespace, gknn.Name, controllerType)
			}
		}
	}
	if len(summary.reports) > 0 {
		fmt.Fprintln(detailFile, "Resources that has issues:")
		for gknn := range summary.reports {
			for controllerType := range summary.reports[gknn] {
				fmt.Fprintln(detailFile, "-----------------------------------------------------------------")
				fmt.Fprintf(detailFile, "Group: %s, Kind: %s, Namespace: %s, Name: %s, Controller Type: %s\n", gknn.Group, gknn.Kind, gknn.Namespace, gknn.Name, controllerType)
				for _, s := range summary.reports[gknn][controllerType] {
					fmt.Fprintln(detailFile, s)
				}
			}
		}
	}
	return detailFile.Close()
}
