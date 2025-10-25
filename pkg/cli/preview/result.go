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
	// namespace -> group -> kind -> []name
	goodGKNN GKNNList
	badGKNN  GKNNList
	reports  map[GKNN][]string
}

type GKNNList struct {
	count int
	// namespace -> group -> kind -> []name
	storage map[string]map[string]map[string][]string
}

func newGKNNList() *GKNNList {
	return &GKNNList{
		storage: make(map[string]map[string]map[string][]string),
		count:   0,
	}
}

func (l *GKNNList) ensurePath(gknn GKNN) {
	if _, ok := l.storage[gknn.Namespace]; !ok {
		l.storage[gknn.Namespace] = make(map[string]map[string][]string)
	}
	if _, ok := l.storage[gknn.Namespace][gknn.Group]; !ok {
		l.storage[gknn.Namespace][gknn.Group] = make(map[string][]string)
	}
	if _, ok := l.storage[gknn.Namespace][gknn.Group][gknn.Kind]; !ok {
		l.storage[gknn.Namespace][gknn.Group][gknn.Kind] = []string{}
	}
}

func (l *GKNNList) add(gknn GKNN) {
	l.storage[gknn.Namespace][gknn.Group][gknn.Kind] = append(l.storage[gknn.Namespace][gknn.Group][gknn.Kind], gknn.Name)
	l.count++
}

func (r *Recorder) newPreviewSummary() *PreviewSummary {
	summary := &PreviewSummary{
		goodGKNN: *newGKNNList(),
		badGKNN:  *newGKNNList(),
		reports:  make(map[GKNN][]string),
	}

	for gknn, info := range r.objects {
		good, report := ParseEventInfo(info)
		summary.goodGKNN.ensurePath(gknn)
		summary.badGKNN.ensurePath(gknn)
		if good {
			summary.goodGKNN.add(gknn)
		} else {
			summary.badGKNN.add(gknn)
			summary.reports[gknn] = report
		}
	}
	return summary
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
			if strings.Contains(event.gcpAction.url, ":getIamPolicy") {
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
		fmt.Fprintln(f, "Known resources that were not reconciled:")
		for gknn, reconciled := range r.ReconciledResources {
			if !reconciled {
				fmt.Fprintln(f, "-----------------------------------------------------------------")
				fmt.Fprintf(f, "Object not reconciled: %+v\n", gknn)
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
	summary := r.newPreviewSummary()
	fmt.Fprintf(f, "Detect %d good and %d bad objects.\n", summary.goodGKNN.count, summary.badGKNN.count)
	for ns := range summary.badGKNN.storage {
		fmt.Fprintln(f, "-----------------------------------------------------------------")
		fmt.Fprintf(f, "Namespace: %s\n", ns)
		w := tabwriter.NewWriter(f, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "GROUP\tKIND\tGOOD\tBAD")
		for group := range summary.badGKNN.storage[ns] {
			for kind := range summary.badGKNN.storage[ns][group] {
				fmt.Fprintf(w, "%s\t%s\t%d\t%d\n", group, kind, len(summary.goodGKNN.storage[ns][group][kind]), len(summary.badGKNN.storage[ns][group][kind]))
			}
		}
		w.Flush()
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("error closing file %q: %w", filename, err)
	}

	// Reconciled all resources and no issue detected.
	if r.RemainResourcesCount == 0 && summary.badGKNN.count == 0 {
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
	if summary.badGKNN.count > 0 {
		for ns := range summary.badGKNN.storage {
			for group := range summary.badGKNN.storage[ns] {
				for kind := range summary.badGKNN.storage[ns][group] {
					for _, name := range summary.badGKNN.storage[ns][group][kind] {
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
