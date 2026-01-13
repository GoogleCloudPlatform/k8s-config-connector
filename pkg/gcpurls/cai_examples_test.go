package gcpurls

import (
	"testing"
)

// AlloyDBBackupIdentity corresponds to alloydb.googleapis.com/Backup
// Format: //alloydb.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/backups/{{BACKUP}}
type AlloyDBBackupIdentity struct {
	Project  string
	Location string
	Backup   string
}

// ApigeeInstanceIdentity corresponds to apigee.googleapis.com/Instance
// Format: //apigee.googleapis.com/organizations/{{ORGANIZATION_ID}}/instances/{{INSTANCE_ID}}
type ApigeeInstanceIdentity struct {
	OrganizationID string
	InstanceID     string
}

// AppEngineServiceIdentity corresponds to appengine.googleapis.com/Service
// Format: //appengine.googleapis.com/apps/{{APP}}/services/{{SERVICE}}
type AppEngineServiceIdentity struct {
	App     string
	Service string
}

// K8sDeploymentIdentity corresponds to apps.k8s.io/Deployment in GKE
// Format: //container.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/clusters/{{CLUSTER}}/k8s/namespaces/{{NAMESPACE}}/apps/deployments/{{DEPLOYMENT}}
type K8sDeploymentIdentity struct {
	Project    string
	Location   string
	Cluster    string
	Namespace  string
	Deployment string
}

func TestCAIExamples(t *testing.T) {
	// 1. AlloyDB Backup
	alloyDBTmpl := Template[AlloyDBBackupIdentity]("projects/{project}/locations/{location}/backups/{backup}")
	alloyDBURL := "//alloydb.googleapis.com/projects/my-project/locations/us-central1/backups/my-backup"

	alloyDBId, err := alloyDBTmpl.Parse(alloyDBURL)
	if err != nil {
		t.Fatalf("Failed to parse AlloyDB URL: %v", err)
	}
	expectedAlloyDB := AlloyDBBackupIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Backup:   "my-backup",
	}
	if *alloyDBId != expectedAlloyDB {
		t.Errorf("AlloyDB match failed: got %+v, expected %+v", *alloyDBId, expectedAlloyDB)
	}

	// Reverse
	s, err := alloyDBTmpl.ToString(*alloyDBId)
	if err != nil {
		t.Errorf("Failed to ToString AlloyDB: %v", err)
	}
	if s != "projects/my-project/locations/us-central1/backups/my-backup" {
		t.Errorf("AlloyDB ToString mismatch: got %q", s)
	}

	// 2. Apigee Instance
	apigeeTmpl := Template[ApigeeInstanceIdentity]("organizations/{OrganizationID}/instances/{InstanceID}")
	apigeeURL := "//apigee.googleapis.com/organizations/my-org/instances/my-instance"

	apigeeId, err := apigeeTmpl.Parse(apigeeURL)
	if err != nil {
		t.Fatalf("Failed to parse Apigee URL: %v", err)
	}
	expectedApigee := ApigeeInstanceIdentity{
		OrganizationID: "my-org",
		InstanceID:     "my-instance",
	}
	if *apigeeId != expectedApigee {
		t.Errorf("Apigee match failed: got %+v, expected %+v", *apigeeId, expectedApigee)
	}

	// 3. AppEngine Service
	appEngineTmpl := Template[AppEngineServiceIdentity]("apps/{app}/services/{service}")
	appEngineURL := "//appengine.googleapis.com/apps/my-app/services/default"

	appEngineId, err := appEngineTmpl.Parse(appEngineURL)
	if err != nil {
		t.Fatalf("Failed to parse AppEngine URL: %v", err)
	}
	expectedAppEngine := AppEngineServiceIdentity{
		App:     "my-app",
		Service: "default",
	}
	if *appEngineId != expectedAppEngine {
		t.Errorf("AppEngine match failed: got %+v, expected %+v", *appEngineId, expectedAppEngine)
	}

	// 4. K8s Deployment
	k8sTmpl := Template[K8sDeploymentIdentity]("projects/{project}/locations/{location}/clusters/{cluster}/k8s/namespaces/{namespace}/apps/deployments/{deployment}")
	k8sURL := "//container.googleapis.com/projects/p1/locations/l1/clusters/c1/k8s/namespaces/ns1/apps/deployments/d1"

	k8sId, err := k8sTmpl.Parse(k8sURL)
	if err != nil {
		t.Fatalf("Failed to parse K8s URL: %v", err)
	}
	expectedK8s := K8sDeploymentIdentity{
		Project:    "p1",
		Location:   "l1",
		Cluster:    "c1",
		Namespace:  "ns1",
		Deployment: "d1",
	}
	if *k8sId != expectedK8s {
		t.Errorf("K8s match failed: got %+v, expected %+v", *k8sId, expectedK8s)
	}
}
