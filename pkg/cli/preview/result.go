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
)

type PreviewSummary struct {
	totalGood int
	totalBad  int
	// namespace -> group -> kind -> []name
	goodGKNN map[string]map[string]map[string][]string
	badGKNN  map[string]map[string]map[string][]string
	reports  map[GKNN][]string
}

func (r *Recorder) newPreviewSummary() *PreviewSummary {
	summary := &PreviewSummary{
		totalGood: 0,
		totalBad:  0,
		goodGKNN:  make(map[string]map[string]map[string][]string),
		badGKNN:   make(map[string]map[string]map[string][]string),
		reports:   make(map[GKNN][]string),
	}

	for gknn, info := range r.objects {
		ensureGKNNPath(summary.goodGKNN, gknn)
		ensureGKNNPath(summary.badGKNN, gknn)

		good, report := ParseEventInfo(info)
		if good {
			summary.totalGood++
			summary.goodGKNN[gknn.Namespace][gknn.Group][gknn.Kind] = append(summary.goodGKNN[gknn.Namespace][gknn.Group][gknn.Kind], gknn.Name)
		} else {
			summary.totalBad++
			summary.badGKNN[gknn.Namespace][gknn.Group][gknn.Kind] = append(summary.badGKNN[gknn.Namespace][gknn.Group][gknn.Kind], gknn.Name)
			summary.reports[gknn] = report
		}
	}
	return summary
}

func ensureGKNNPath(m map[string]map[string]map[string][]string, gknn GKNN) {
	if _, ok := m[gknn.Namespace]; !ok {
		m[gknn.Namespace] = make(map[string]map[string][]string)
	}
	if _, ok := m[gknn.Namespace][gknn.Group]; !ok {
		m[gknn.Namespace][gknn.Group] = make(map[string][]string)
	}
	if _, ok := m[gknn.Namespace][gknn.Group][gknn.Kind]; !ok {
		m[gknn.Namespace][gknn.Group][gknn.Kind] = []string{}
	}
}

func ParseEventInfo(info *objectInfo) (bool, []string) {
	good := true
	report := []string{}
	for _, event := range info.events {
		switch event.eventType {
		case EventTypeDiff:
			report = append(report, "KRM diff detected:")
			if event.diff != nil && event.diff.Fields != nil {
				diffFields := []string{}
				for _, field := range event.diff.Fields {
					if !reflect.DeepEqual(field.Old, field.New) {
						diffFields = append(diffFields, field.ID)
					}
				}
				report = append(report, fmt.Sprintf("  Diff Fields: %s", diffFields))
			}
		case EventTypeReconcileStart, EventTypeReconcileEnd:
			// fmt.Fprintf(f, "  reconcileStart %+v\n", event.object)

		case EventTypeKubeAction:
			// Ignore kubeaction for now. Mostly status update.

		case EventTypeGCPAction:
			// The method is POST but it is actually a read-only call.
			if strings.Contains(event.gcpAction.url, "getIamPolicy") {
				continue
			}
			good = false
			report = append(report, "GCP action detected:")
			report = append(report, fmt.Sprintf("  Method: %s", event.gcpAction.method))
			report = append(report, fmt.Sprintf("  Url: %s", event.gcpAction.url))
			report = append(report, fmt.Sprintf("  Object: %+v", event.gcpAction.body))
		default:
			good = false
		}
	}
	return good, report
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
		for gknn, reconciled := range r.ReconciledResources {
			if !reconciled {
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
	total := len(r.ReconciledResources)
	reconciled := total - r.RemainResourcesCount
	fmt.Fprintf(f, "Finish reconciled %d out of %d resouces.\n", reconciled, total)
	summary := r.getOrCreateSummary()
	fmt.Fprintf(f, "Detect %d good and %d bad objects.\n", summary.totalGood, summary.totalBad)
	for ns := range summary.badGKNN {
		fmt.Fprintln(f, "-----------------------------------------------------------------")
		fmt.Fprintf(f, "Namespace: %s\n", ns)
		w := tabwriter.NewWriter(f, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "GROUP\tKIND\tGOOD\tBAD")
		for group := range summary.badGKNN[ns] {
			for kind := range summary.badGKNN[ns][group] {
				fmt.Fprintf(w, "%s\t%s\t%d\t%d\n", group, kind, len(summary.goodGKNN[ns][group][kind]), len(summary.badGKNN[ns][group][kind]))
			}
		}
		w.Flush()
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("error closing file %q: %w", filename, err)
	}

	// Reconciled all resources and no issue detected.
	if r.RemainResourcesCount == 0 && summary.totalBad == 0 {
		return nil
	}
	detailFileName := fmt.Sprintf("%s-detail", filename)
	detailFile, err := os.Create(detailFileName)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", detailFileName, err)
	}
	if r.RemainResourcesCount > 0 {
		fmt.Fprintln(detailFile, "Resources that has not fully reconciled:")
		for gknn, reconciled := range r.ReconciledResources {
			if !reconciled {
				fmt.Fprintf(detailFile, "Group: %s, Kind: %s, Namespace: %s, Name: %s\n", gknn.Group, gknn.Kind, gknn.Namespace, gknn.Name)
			}
		}
	}
	if summary.totalBad > 0 {
		for ns := range summary.badGKNN {
			for group := range summary.badGKNN[ns] {
				for kind := range summary.badGKNN[ns][group] {
					for _, name := range summary.badGKNN[ns][group][kind] {
						fmt.Fprintln(detailFile, "-----------------------------------------------------------------")
						fmt.Fprintf(detailFile, "Group: %s, Kind: %s, Namespace: %s, Name: %s\n", group, kind, ns, name)
						for _, s := range summary.reports[GKNN{Group: group, Kind: kind, Namespace: ns, Name: name}] {
							fmt.Fprintln(detailFile, s)
						}
					}
				}
			}
		}
	}
	return detailFile.Close()
}
