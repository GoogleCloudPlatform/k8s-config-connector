// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package sql

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func TestSqlUser_findSqlUser(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name            string
		databaseVersion string
		userName        string
		host            string
		users           []*sqladmin.User
		wantName        string
		wantType        string
	}{
		{
			name:            "SQL Server ENTRAID_USER preserves full email UPN",
			databaseVersion: "SQLSERVER_2022_STANDARD",
			userName:        "user@tenant.onmicrosoft.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "sqlserver", Type: "BUILT_IN"},
				{Name: "user@tenant.onmicrosoft.com", Type: "ENTRAID_USER"},
			},
			wantName: "user@tenant.onmicrosoft.com",
			wantType: "ENTRAID_USER",
		},
		{
			name:            "SQL Server ENTRAID_USER does not collide with same-prefix BUILT_IN login",
			databaseVersion: "SQLSERVER_2022_STANDARD",
			userName:        "sqlserver@tenant.onmicrosoft.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "sqlserver", Type: "BUILT_IN"},
				{Name: "sqlserver@tenant.onmicrosoft.com", Type: "ENTRAID_USER"},
			},
			wantName: "sqlserver@tenant.onmicrosoft.com",
			wantType: "ENTRAID_USER",
		},
		{
			name:            "SQL Server ENTRAID_USER missing from list returns nil",
			databaseVersion: "SQLSERVER_2022_STANDARD",
			userName:        "missing@tenant.onmicrosoft.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "sqlserver", Type: "BUILT_IN"},
			},
			wantName: "",
		},
		{
			name:            "SQL Server BUILT_IN user matches exact username",
			databaseVersion: "SQLSERVER_2022_STANDARD",
			userName:        "appuser",
			host:            "",
			users: []*sqladmin.User{
				{Name: "appuser", Type: "BUILT_IN"},
				{Name: "sqlserver", Type: "BUILT_IN"},
			},
			wantName: "appuser",
			wantType: "BUILT_IN",
		},
		{
			name:            "MySQL BUILT_IN user matches username and host",
			databaseVersion: "MYSQL_8_0",
			userName:        "admin",
			host:            "gmail.com",
			users: []*sqladmin.User{
				{Name: "admin", Host: "google.com", Type: "BUILT_IN"},
				{Name: "admin", Host: "gmail.com", Type: "BUILT_IN"},
			},
			wantName: "admin",
			wantType: "BUILT_IN",
		},
		{
			name:            "MySQL CLOUD_IAM_USER truncates domain at @",
			databaseVersion: "MYSQL_8_0",
			userName:        "admin@hashicorptest.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "admin", Host: "%", Type: "CLOUD_IAM_USER"},
			},
			wantName: "admin",
			wantType: "CLOUD_IAM_USER",
		},
		{
			name:            "MySQL CLOUD_IAM_SERVICE_ACCOUNT truncates domain at @",
			databaseVersion: "MYSQL_8_0",
			userName:        "tf-test-sa@tf-test-project.iam.gserviceaccount.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "tf-test-sa", Host: "%", Type: "CLOUD_IAM_SERVICE_ACCOUNT"},
			},
			wantName: "tf-test-sa",
			wantType: "CLOUD_IAM_SERVICE_ACCOUNT",
		},
		{
			name:            "MySQL CLOUD_IAM_GROUP preserves @ when preceded by other user types",
			databaseVersion: "MYSQL_8_0",
			userName:        "db-readers@example.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "root", Host: "%", Type: "BUILT_IN"},
				{Name: "app-sa", Host: "%", Type: "CLOUD_IAM_SERVICE_ACCOUNT"},
				{Name: "db-readers@example.com", Host: "%", Type: "CLOUD_IAM_GROUP"},
			},
			wantName: "db-readers@example.com",
			wantType: "CLOUD_IAM_GROUP",
		},
		{
			name:            "MySQL CLOUD_IAM_GROUP preserves @ and lowercases domain",
			databaseVersion: "MYSQL_8_0",
			userName:        "iam-group-auth-test-group@GOOGLE.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "iam-group-auth-test-group@google.com", Host: "%", Type: "CLOUD_IAM_GROUP"},
			},
			wantName: "iam-group-auth-test-group@google.com",
			wantType: "CLOUD_IAM_GROUP",
		},
		{
			name:            "PostgreSQL CLOUD_IAM_USER preserves full email",
			databaseVersion: "POSTGRES_15",
			userName:        "admin@hashicorptest.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "postgres", Type: "BUILT_IN"},
				{Name: "admin@hashicorptest.com", Type: "CLOUD_IAM_USER"},
			},
			wantName: "admin@hashicorptest.com",
			wantType: "CLOUD_IAM_USER",
		},
		{
			name:            "PostgreSQL CLOUD_IAM_GROUP preserves exact name",
			databaseVersion: "POSTGRES_15",
			userName:        "iam-group-auth-test-group@GOOGLE.com",
			host:            "",
			users: []*sqladmin.User{
				{Name: "iam-group-auth-test-group@GOOGLE.com", Type: "CLOUD_IAM_GROUP"},
			},
			wantName: "iam-group-auth-test-group@GOOGLE.com",
			wantType: "CLOUD_IAM_GROUP",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := findSqlUser(tc.databaseVersion, tc.userName, tc.host, tc.users)
			if tc.wantName == "" {
				if got != nil {
					t.Fatalf("findSqlUser() = %+v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Fatalf("findSqlUser() = nil, want user %q (%s)", tc.wantName, tc.wantType)
			}
			if got.Name != tc.wantName || got.Type != tc.wantType {
				t.Fatalf("findSqlUser() = {Name: %q, Type: %q}, want {Name: %q, Type: %q}", got.Name, got.Type, tc.wantName, tc.wantType)
			}
		})
	}
}

func TestSqlUser_diffSuppressIamUserName(t *testing.T) {
	t.Parallel()

	res := ResourceSqlUser()

	cases := []struct {
		name     string
		userType string
		host     string
		oldName  string
		newName  string
		want     bool
	}{
		{
			name:     "MySQL CLOUD_IAM_USER suppresses stripped domain",
			userType: "CLOUD_IAM_USER",
			host:     "%",
			oldName:  "alice",
			newName:  "alice@example.com",
			want:     true,
		},
		{
			name:     "MySQL CLOUD_IAM_GROUP suppresses lowercase domain when host is %",
			userType: "CLOUD_IAM_GROUP",
			host:     "%",
			oldName:  "db-readers@example.com",
			newName:  "db-readers@EXAMPLE.COM",
			want:     true,
		},
		{
			name:     "MySQL CLOUD_IAM_GROUP does not suppress different group domain",
			userType: "CLOUD_IAM_GROUP",
			host:     "%",
			oldName:  "db-readers@example.com",
			newName:  "db-readers@other.com",
			want:     false,
		},
		{
			name:     "Postgres CLOUD_IAM_GROUP does not suppress uppercase domain when host is empty",
			userType: "CLOUD_IAM_GROUP",
			host:     "",
			oldName:  "db-readers@example.com",
			newName:  "db-readers@EXAMPLE.COM",
			want:     false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			d := res.Data(&terraform.InstanceState{
				Attributes: map[string]string{
					"type": tc.userType,
					"host": tc.host,
				},
			})
			got := diffSuppressIamUserName("name", tc.oldName, tc.newName, d)
			if got != tc.want {
				t.Fatalf("diffSuppressIamUserName(%q, %q) = %v, want %v", tc.oldName, tc.newName, got, tc.want)
			}
		})
	}
}

func TestSqlUser_mysqlCloudIamGroupLifecycle(t *testing.T) {
	t.Parallel()

	const (
		project      = "tf-test-project"
		instanceName = "tf-test-mysql"
		groupEmail   = "db-readers@example.com"
	)

	var deletedName, deletedHost string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == fmt.Sprintf("/sql/v1beta4/projects/%s/instances/%s/users", project, instanceName):
			_ = json.NewEncoder(w).Encode(&sqladmin.Operation{
				Name:   "op-insert",
				Status: "DONE",
			})
		case r.Method == http.MethodGet && r.URL.Path == fmt.Sprintf("/sql/v1beta4/projects/%s/operations/op-insert", project):
			_ = json.NewEncoder(w).Encode(&sqladmin.Operation{
				Name:   "op-insert",
				Status: "DONE",
			})
		case r.Method == http.MethodGet && r.URL.Path == fmt.Sprintf("/sql/v1beta4/projects/%s/operations/op-delete", project):
			_ = json.NewEncoder(w).Encode(&sqladmin.Operation{
				Name:   "op-delete",
				Status: "DONE",
			})
		case r.Method == http.MethodGet && r.URL.Path == fmt.Sprintf("/sql/v1beta4/projects/%s/instances/%s", project, instanceName):
			_ = json.NewEncoder(w).Encode(&sqladmin.DatabaseInstance{
				Name:            instanceName,
				DatabaseVersion: "MYSQL_8_0",
			})
		case r.Method == http.MethodGet && r.URL.Path == fmt.Sprintf("/sql/v1beta4/projects/%s/instances/%s/users", project, instanceName):
			_ = json.NewEncoder(w).Encode(&sqladmin.UsersListResponse{
				Items: []*sqladmin.User{
					{
						Name:     "root",
						Host:     "%",
						Instance: instanceName,
						Type:     "BUILT_IN",
					},
					{
						Name:     "example-app",
						Host:     "%",
						Instance: instanceName,
						Type:     "CLOUD_IAM_SERVICE_ACCOUNT",
					},
					{
						Name:     groupEmail,
						Host:     "%",
						Instance: instanceName,
						Type:     "CLOUD_IAM_GROUP",
					},
				},
			})
		case r.Method == http.MethodDelete && r.URL.Path == fmt.Sprintf("/sql/v1beta4/projects/%s/instances/%s/users", project, instanceName):
			deletedName = r.URL.Query().Get("name")
			deletedHost = r.URL.Query().Get("host")
			_ = json.NewEncoder(w).Encode(&sqladmin.Operation{
				Name:   "op-delete",
				Status: "DONE",
			})
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	cfg := &transport_tpg.Config{
		Context:      context.Background(),
		Client:       server.Client(),
		Project:      project,
		PollInterval: time.Millisecond,
		SQLBasePath:  server.URL + "/",
	}

	res := ResourceSqlUser()
	d, err := schema.InternalMap(res.Schema).Data(&terraform.InstanceState{}, &terraform.InstanceDiff{
		Attributes: map[string]*terraform.ResourceAttrDiff{
			"instance": {New: instanceName},
			"name":     {New: groupEmail},
			"type":     {New: "CLOUD_IAM_GROUP"},
		},
	})
	if err != nil {
		t.Fatalf("Data init failed: %v", err)
	}

	if err := res.Create(d, cfg); err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	wantID := fmt.Sprintf("%s/%%/%s", groupEmail, instanceName)
	if d.Id() != wantID {
		t.Fatalf("Create/Read produced ID %q, want %q", d.Id(), wantID)
	}
	if got := d.Get("name").(string); got != groupEmail {
		t.Fatalf("Read produced name %q, want %q", got, groupEmail)
	}
	if got := d.Get("host").(string); got != "%" {
		t.Fatalf("Read produced host %q, want %q", got, "%")
	}
	if got := d.Get("type").(string); got != "CLOUD_IAM_GROUP" {
		t.Fatalf("Read produced type %q, want %q", got, "CLOUD_IAM_GROUP")
	}

	// Simulate KCC's FetchLiveState (Importer + Read) on subsequent reconciles.
	importData := res.Data(nil)
	importData.SetId(fmt.Sprintf("%s/%s/%s", project, instanceName, groupEmail))
	imported, err := res.Importer.State(importData, cfg)
	if err != nil {
		t.Fatalf("Importer failed: %v", err)
	}
	if err := res.Read(imported[0], cfg); err != nil {
		t.Fatalf("Post-import Read failed: %v", err)
	}
	if imported[0].Id() != wantID {
		t.Fatalf("Post-import Read produced ID %q, want %q", imported[0].Id(), wantID)
	}

	if err := res.Delete(d, cfg); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if deletedName != groupEmail || deletedHost != "%" {
		t.Fatalf("Delete sent name=%q host=%q, want name=%q host=%%", deletedName, deletedHost, groupEmail)
	}
}
