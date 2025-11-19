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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
)

type GKNSummary struct {
	// number of resouces successfully reconciled by TF reconciler
	tfCount int
	// number of resouces successfully reconciled by direct reconciler
	directCount int
	// number of resouces successfully reconciled by DCL reconciler
	dclCount int
	// number of resouces successfully reconciled by custom reconciler
	customCount int
}
type GKN struct {
	Group     string
	Kind      string
	Namespace string
}

type PreviewSummary struct {
	// namespace -> group -> kind -> GKNSummary
	summary  map[string]map[string]map[string]*GKNSummary
	reports  map[GKNN]map[k8s.ReconcilerType][]string
	gknCount map[GKN]int
}

func (p *PreviewSummary) addResult(gknn GKNN, r ReconcileResult) {
	if _, ok := p.summary[gknn.Namespace]; !ok {
		p.summary[gknn.Namespace] = make(map[string]map[string]*GKNSummary)
	}
	if _, ok := p.summary[gknn.Namespace][gknn.Group]; !ok {
		p.summary[gknn.Namespace][gknn.Group] = make(map[string]*GKNSummary)
	}
	var summary *GKNSummary
	if _, ok := p.summary[gknn.Namespace][gknn.Group][gknn.Kind]; !ok {
		summary = &GKNSummary{}
	} else {
		summary = p.summary[gknn.Namespace][gknn.Group][gknn.Kind]
	}
	if r.good {
		switch r.controllerType {
		case k8s.ReconcilerTypeTerraform:
			summary.tfCount++
		case k8s.ReconcilerTypeDCL:
			summary.dclCount++
		case k8s.ReconcilerTypeDirect:
			summary.directCount++
		default:
			summary.customCount++
		}
	}
	p.summary[gknn.Namespace][gknn.Group][gknn.Kind] = summary
}

type ReconcileResult struct {
	controllerType k8s.ReconcilerType
	good           bool
	detail         []string
}

func (r *Recorder) newPreviewSummary() *PreviewSummary {
	summary := &PreviewSummary{
		summary:  make(map[string]map[string]map[string]*GKNSummary),
		reports:  make(map[GKNN]map[k8s.ReconcilerType][]string),
		gknCount: make(map[GKN]int),
	}

	for gknn, info := range r.objects {
		summary.gknCount[GKN{
			Group:     gknn.Group,
			Kind:      gknn.Kind,
			Namespace: gknn.Namespace,
		}]++

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
				fmt.Fprintf(f, "  reconcileStart %+v\n", event.object)

			case EventTypeReconcileEnd:
				fmt.Fprintf(f, "  reconcileEnd %+v\n", event.object)

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
	fmt.Fprintln(w, "Namespace\tGroup\tKind\tTF\tDirect\tDCL\tCustom\tTotal")
	for ns := range summary.summary {
		for group := range summary.summary[ns] {
			for kind := range summary.summary[ns][group] {
				gs := summary.summary[ns][group][kind]
				fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%d\t%d\t%d\t%d\n", ns, group, kind, gs.tfCount, gs.directCount, gs.dclCount, gs.customCount, summary.gknCount[GKN{Group: group, Kind: kind, Namespace: ns}])
			}
		}
	}
	w.Flush()
	if err = f.Close(); err != nil {
		return fmt.Errorf("error closing file %q: %w", filename, err)
	}
	return nil

	// Reconciled all resources and no issue detected.
	// if r.RemainResourcesCount == 0 && summary.totalBad == 0 {
	// 	return nil
	// }
	// detailFileName := fmt.Sprintf("%s-detail", filename)
	// detailFile, err := os.Create(detailFileName)
	// if err != nil {
	// 	return fmt.Errorf("error creating file %q: %w", detailFileName, err)
	// }
	// if r.RemainResourcesCount > 0 {
	// 	fmt.Fprintln(detailFile, "Resources that has not fully reconciled:")
	// 	for gknn := range r.ReconciledResources {
	// 		if !r.GKNNDoneReconcile(gknn) {
	// 			fmt.Fprintf(detailFile, "Group: %s, Kind: %s, Namespace: %s, Name: %s\n", gknn.Group, gknn.Kind, gknn.Namespace, gknn.Name)
	// 		}
	// 	}
	// }
	// if summary.totalBad > 0 {
	// 	for ns := range summary.badGKNN {
	// 		for group := range summary.badGKNN[ns] {
	// 			for kind := range summary.badGKNN[ns][group] {
	// 				for _, name := range summary.badGKNN[ns][group][kind] {
	// 					fmt.Fprintln(detailFile, "-----------------------------------------------------------------")
	// 					fmt.Fprintf(detailFile, "Group: %s, Kind: %s, Namespace: %s, Name: %s\n", group, kind, ns, name)
	// 					for _, s := range summary.reports[GKNN{Group: group, Kind: kind, Namespace: ns, Name: name}] {
	// 						fmt.Fprintln(detailFile, s)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// return detailFile.Close()
}
