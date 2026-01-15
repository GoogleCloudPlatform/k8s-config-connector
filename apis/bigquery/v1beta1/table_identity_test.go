package v1beta1

import (
	"testing"
)

func TestTableIdentity_FromExternal(t *testing.T) {
	tests := []struct {
		name     string
		ref      string
		wantErr  bool
		want     *TableIdentity
	}{
		{
			name: "valid reference",
			ref:  "projects/my-project/datasets/my-dataset/tables/my-table",
			want: &TableIdentity{
				Project: "my-project",
				Dataset: "my-dataset",
				Table:   "my-table",
			},
		},
		{
			name:    "invalid reference format",
			ref:     "invalid/format",
			wantErr: true,
		},
		{
			name: "full url",
			ref:  "https://bigquery.googleapis.com/projects/my-project/datasets/my-dataset/tables/my-table",
			want: &TableIdentity{
				Project: "my-project",
				Dataset: "my-dataset",
				Table:   "my-table",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &TableIdentity{}
			err := i.FromExternal(tt.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromExternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if i.Project != tt.want.Project {
					t.Errorf("Project = %v, want %v", i.Project, tt.want.Project)
				}
				if i.Dataset != tt.want.Dataset {
					t.Errorf("Dataset = %v, want %v", i.Dataset, tt.want.Dataset)
				}
				if i.Table != tt.want.Table {
					t.Errorf("Table = %v, want %v", i.Table, tt.want.Table)
				}
			}
		})
	}
}
