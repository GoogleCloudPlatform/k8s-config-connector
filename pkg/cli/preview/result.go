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
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/resourceconfig"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/structuredreporting"
	"k8s.io/apimachinery/pkg/runtime/schema"
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
	if diffs == nil {
		return ""
	}
	return strings.Join(diffs.FieldIDs(), ",")
}

// RecorderReconciledResults is the result of reconciling all GKNN objects recorded by the recorder.
type RecorderReconciledResults struct {
	results   map[GKNN]*GKNNReconciledResult
	badResult []*GKNNReconciledResult
	badCount  int
	goodCount int
}

func (r *RecorderReconciledResults) AddResult(gknn GKNN, result *GKNNReconciledResult) {
	if r.results == nil {
		r.results = make(map[GKNN]*GKNNReconciledResult)
	}
	r.results[gknn] = result
}

func (r *Recorder) GenerateRecorderReconciledResults() *RecorderReconciledResults {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	recorderReconciledResults := &RecorderReconciledResults{
		results: make(map[GKNN]*GKNNReconciledResult),
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
			recorderReconciledResults.badResult = append(recorderReconciledResults.badResult, result)
			recorderReconciledResults.badCount++
		} else {
			recorderReconciledResults.goodCount++
		}
	}
	return recorderReconciledResults
}

func (r *RecorderReconciledResults) CombinedSummaryReport(summaryFile string, altResult *RecorderReconciledResults, altExpectedMap map[schema.GroupKind]k8s.ReconcilerType) error {
	var combinedBadResult []*GKNNReconciledResult
	if len(r.badResult) > 0 {
		combinedBadResult = append(combinedBadResult, r.badResult...)
	}
	if altResult != nil && len(altResult.badResult) > 0 {
		combinedBadResult = append(combinedBadResult, altResult.badResult...)
	}

	defer func() {
		for _, result := range combinedBadResult {
			klog.V(0).Info("\"PreviewResult\" ", result.FormatGKNNReconciledResult())
		}
	}()

	f, err := os.Create(summaryFile)
	if err != nil {
		return fmt.Errorf("error creating file %q: %w", summaryFile, err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			klog.ErrorS(err, "failed to close file", "file", summaryFile)
		}
	}()

	fmt.Fprintf(f, "Detected %d good and %d bad objects in default run\n", r.goodCount, r.badCount)
	if altResult != nil {
		fmt.Fprintf(f, "Detected %d good and %d bad objects in alternative run\n", altResult.goodCount, altResult.badCount)
	}
	w := tabwriter.NewWriter(f, 0, 0, 3, ' ', 0)
	defer func() {
		if err := w.Flush(); err != nil {
			klog.ErrorS(err, "error flushing summary report", "file", summaryFile)
		}
	}()
	fmt.Fprintln(w, "GROUP\tKIND\tNAME\tDEFAULT-CONTROLLER\tDEFAULT-RESULT\tDEFAULT-DIFFS\tALTERNATIVE-CONTROLLER\tALTERNATIVE-RESULT\tALTERNATIVE-DIFFS")
	type resultPair struct {
		def  *GKNNReconciledResult
		alt  *GKNNReconciledResult
		gknn GKNN
	}

	type combinedResult struct {
		results map[GKNN]*resultPair
	}

	combined := &combinedResult{
		results: make(map[GKNN]*resultPair),
	}

	addPair := func(gknn GKNN, result *GKNNReconciledResult, isAlt bool) {
		pair, ok := combined.results[gknn]
		if !ok {
			pair = &resultPair{gknn: gknn}
			combined.results[gknn] = pair
		}
		if isAlt {
			pair.alt = result
		} else {
			pair.def = result
		}
	}

	for gknn, result := range r.results {
		addPair(gknn, result, false)
	}

	if altResult != nil {
		for gknn, result := range altResult.results {
			addPair(gknn, result, true)
		}
	}

	// Sort results by GKNN for stable output.
	var sortedGKNNs []GKNN
	for gknn := range combined.results {
		sortedGKNNs = append(sortedGKNNs, gknn)
	}
	sort.Slice(sortedGKNNs, func(i, j int) bool {
		if sortedGKNNs[i].Group != sortedGKNNs[j].Group {
			return sortedGKNNs[i].Group < sortedGKNNs[j].Group
		}
		if sortedGKNNs[i].Kind != sortedGKNNs[j].Kind {
			return sortedGKNNs[i].Kind < sortedGKNNs[j].Kind
		}
		if sortedGKNNs[i].Namespace != sortedGKNNs[j].Namespace {
			return sortedGKNNs[i].Namespace < sortedGKNNs[j].Namespace
		}
		return sortedGKNNs[i].Name < sortedGKNNs[j].Name
	})

	for _, gknn := range sortedGKNNs {
		pair := combined.results[gknn]
		defCtrl, defStatus := "N/A", "N/A"
		altCtrl, altStatus := "N/A", "N/A"

		// Determine if the resource class supports an alternative controller using the pre-computed static map
		gk := schema.GroupKind{Group: pair.gknn.Group, Kind: pair.gknn.Kind}
		altExpected, hasAlternative := altExpectedMap[gk]

		// Track stringified diff representations explicitly. If empty or unprocessed, default to "N/A"
		defDiffs := "N/A"
		altDiffs := "N/A"

		// Safely load the details resulting from the default preview execution run
		if pair.def != nil {
			defCtrl = string(pair.def.ControllerType)
			defStatus = formatReconciledStatus(pair.def)
			defDiffs = FormatFieldIDs(pair.def.Diffs)
		}

		// Evaluate alternative results logic.
		// If no alternative target exists for this group kind, omit standard reporting.
		// If one exists and the alternate run successfully populated a result bearing the expected controller type, assign its tracked specifics.
		// Otherwise, the alternative controller implicitly failed to execute and should distinctly be marked as "Missing".
		if !hasAlternative {
			altCtrl = "N/A"
			altStatus = "N/A"
		} else if pair.alt != nil && pair.alt.ControllerType == altExpected {
			altCtrl = string(pair.alt.ControllerType)
			altStatus = formatReconciledStatus(pair.alt)
			altDiffs = FormatFieldIDs(pair.alt.Diffs)
		} else {
			altCtrl = string(altExpected)
			altStatus = "Missing"
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", pair.gknn.Group, pair.gknn.Kind, pair.gknn.Name, defCtrl, defStatus, defDiffs, altCtrl, altStatus, altDiffs)
	}

	if len(combinedBadResult) > 0 {
		if err := r.BadResultReport(summaryFile+"-detail", combinedBadResult); err != nil {
			return fmt.Errorf("error creating bad result detail report: %w", err)
		}
	}

	return nil
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
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.reconcileTrackerMutex.Lock()
	defer r.reconcileTrackerMutex.Unlock()

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

func GetAlternativeControllerExpectedMap(configMap resourceconfig.ResourcesControllerMap) map[schema.GroupKind]k8s.ReconcilerType {
	altExpectedMap := make(map[schema.GroupKind]k8s.ReconcilerType)
	for gk, config := range configMap {
		for _, ctrl := range config.SupportedControllers {
			if ctrl != config.DefaultController {
				altExpectedMap[gk] = ctrl
				break
			}
		}
	}
	return altExpectedMap
}
