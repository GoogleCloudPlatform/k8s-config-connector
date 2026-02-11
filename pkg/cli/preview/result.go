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
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/klog/v2"
)

type ReconcileStatus int

// ReconcileStatus is the status of a GKNN object after being reconciled.
const (
	ReconcileStatusUnknown ReconcileStatus = iota
	ReconcileStatusHealthy
	ReconcileStatusUnhealthy
)

// reconcileStatusMap is a map of ReconcileStatus to string.
var reconcileStatusMap = map[ReconcileStatus]string{
	ReconcileStatusHealthy:   "RECONCILE_STATUS_HEALTHY",
	ReconcileStatusUnhealthy: "RECONCILE_STATUS_UNHEALTHY",
}

func (s ReconcileStatus) String() string {
	return reconcileStatusMap[s]
}

// GKNNReconciledResult is the result of reconciling a GKNN object with a specific controller type.
type GKNNReconciledResult struct {
	GKNN            GKNN
	ControllerType  k8s.ReconcilerType
	ReconcileStatus ReconcileStatus
	Diffs           *structuredreporting.Diff
	GCPActions      []*gcpAction
}

// FormatGKNNReconciledResult formats the GKNNReconciledResult into a string.
func (r *GKNNReconciledResult) FormatGKNNReconciledResult() string {
	return fmt.Sprintf("ns=\"%s\" name=\"%s\" group=\"%s\" kind=\"%s\" controller_type=\"%s\" diffs=\"%s\" reconcile_status=\"%s\"", r.GKNN.Namespace, r.GKNN.Name, r.GKNN.Group, r.GKNN.Kind, r.ControllerType, FormatFieldIDs(r.Diffs), r.ReconcileStatus.String())
}

func FormatFieldIDs(diffs *structuredreporting.Diff) string {
	return strings.Join(diffs.FieldIDs(), ",")
}

// RecorderReconciledResults is the result of reconciling all GKNN objects recorded by the recorder.
type RecorderReconciledResults struct {
	results   map[string]map[string]map[string][]*GKNNReconciledResult
	badCount  int
	goodCount int
}

func (r *RecorderReconciledResults) AddResult(gknn GKNN, result *GKNNReconciledResult) {
	if _, ok := r.results[gknn.Group]; !ok {
		r.results[gknn.Group] = make(map[string]map[string][]*GKNNReconciledResult)
	}
	if _, ok := r.results[gknn.Group][gknn.Kind]; !ok {
		r.results[gknn.Group][gknn.Kind] = make(map[string][]*GKNNReconciledResult)
	}
	if _, ok := r.results[gknn.Group][gknn.Kind][gknn.Namespace]; !ok {
		r.results[gknn.Group][gknn.Kind][gknn.Namespace] = []*GKNNReconciledResult{}
	}
	r.results[gknn.Group][gknn.Kind][gknn.Namespace] = append(r.results[gknn.Group][gknn.Kind][gknn.Namespace], result)
}

func (r *Recorder) GenerateRecorderReconciledResults() *RecorderReconciledResults {
	recorderReconciledResults := &RecorderReconciledResults{
		results: make(map[string]map[string]map[string][]*GKNNReconciledResult),
	}

	for gknn := range r.objects {
		result := &GKNNReconciledResult{
			GKNN:            gknn,
			Diffs:           &structuredreporting.Diff{},
			ReconcileStatus: ReconcileStatusHealthy,
			GCPActions:      []*gcpAction{},
		}
		for _, event := range r.objects[gknn].events {
			switch event.eventType {
			case EventTypeDiff:
				result.Diffs.AddDiff(event.diff)
			case EventTypeReconcileStart:
				result.ControllerType = event.reconcilerType
			case EventTypeReconcileEnd:
				result.ControllerType = event.reconcilerType
			case EventTypeKubeAction:
				// Ignore for now
			case EventTypeGCPAction:
				// The method is POST but it is actually a read-only call.
				if strings.Contains(event.gcpAction.URL, "getIamPolicy") {
					continue
				}
				result.GCPActions = append(result.GCPActions, event.gcpAction)
				result.ReconcileStatus = ReconcileStatusUnhealthy
			default:
				// Ignore for now
			}
		}
		recorderReconciledResults.AddResult(gknn, result)
		if result.ReconcileStatus == ReconcileStatusUnhealthy {
			recorderReconciledResults.badCount++
		} else {
			recorderReconciledResults.goodCount++
		}
	}
	return recorderReconciledResults
}

func (r *RecorderReconciledResults) SummaryReport(summaryFile string) error {
	f, err := os.Create(summaryFile)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", summaryFile, err)
	}
	badResult := []*GKNNReconciledResult{}
	fmt.Fprintf(f, "Detected %d good and %d bad objects\n", r.goodCount, r.badCount)
	w := tabwriter.NewWriter(f, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "GROUP\tKIND\tNAMESPACE\tNAME\tController Type\tReconcile Status\tDiff Fields")
	for group := range r.results {
		for kind := range r.results[group] {
			for namespace := range r.results[group][kind] {
				for _, result := range r.results[group][kind][namespace] {
					reconcileStatus := formatReconciledStatus(result)
					if result.ReconcileStatus == ReconcileStatusUnhealthy {
						klog.V(0).Info("PreviewResult ", result.FormatGKNNReconciledResult())
						badResult = append(badResult, result)
					}
					fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n", group, kind, namespace, result.GKNN.Name, result.ControllerType, reconcileStatus, FormatFieldIDs(result.Diffs))
				}
			}
		}
	}
	if len(badResult) > 0 {
		if err := r.BadResultReport(summaryFile+"-detail", badResult); err != nil {
			return fmt.Errorf("error creating bad result detail report: %w", err)
		}
	}

	if err := w.Flush(); err != nil {
		return fmt.Errorf("error flushing summary report: %w", err)
	}
	return f.Close()
}

func (r *RecorderReconciledResults) BadResultReport(badResultFile string, badResult []*GKNNReconciledResult) error {
	f, err := os.Create(badResultFile)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", badResultFile, err)
	}
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(badResult); err != nil {
		return fmt.Errorf("error encoding bad results to json: %w", err)
	}
	return f.Close()
}

func formatReconciledStatus(result *GKNNReconciledResult) string {
	if result == nil {
		return "N/A"
	}
	if result.ReconcileStatus == ReconcileStatusUnhealthy {
		return "UNHEALTHY"
	}
	return "HEALTHY"
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
