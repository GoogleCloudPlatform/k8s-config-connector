package v1_test

import (
	"testing"

	v1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/connector/v1"
)

func TestProviderIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		external string
		want     v1.ProviderIdentity
		wantErr  bool
	}{
		{
			name:     "valid external",
			external: "projects/my-project/locations/us-central1/providers/my-provider",
			want: v1.ProviderIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Provider: "my-provider",
			},
			wantErr: false,
		},
		{
			name:     "invalid external format",
			external: "invalid/format",
			wantErr:  true,
		},
		{
			name:     "with scheme",
			external: "https://connectors.googleapis.com/projects/my-project/locations/us-central1/providers/my-provider",
			want: v1.ProviderIdentity{
				Project:  "my-project",
				Location: "us-central1",
				Provider: "my-provider",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i v1.ProviderIdentity
			err := i.FromExternal(tt.external)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProviderIdentity.FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && i != tt.want {
				t.Errorf("ProviderIdentity.FromExternal() = %v, want %v", i, tt.want)
			}
		})
	}
}

func TestProviderIdentity_String(t *testing.T) {
	i := v1.ProviderIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Provider: "my-provider",
	}
	want := "projects/my-project/locations/us-central1/providers/my-provider"
	if got := i.String(); got != want {
		t.Errorf("ProviderIdentity.String() = %v, want %v", got, want)
	}
}

func TestProviderIdentity_Parent(t *testing.T) {
	i := v1.ProviderIdentity{
		Project:  "my-project",
		Location: "us-central1",
		Provider: "my-provider",
	}
	parent := i.Parent()
	if parent.ProjectID != "my-project" {
		t.Errorf("ProviderIdentity.Parent().ProjectID = %v, want %v", parent.ProjectID, "my-project")
	}
	if parent.Location != "us-central1" {
		t.Errorf("ProviderIdentity.Parent().Location = %v, want %v", parent.Location, "us-central1")
	}
}
