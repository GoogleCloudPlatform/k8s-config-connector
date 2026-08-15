package utils

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	addonsv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

type mockCommonObject struct {
	addonsv1alpha1.CommonObject
	name string
}

func (m *mockCommonObject) ComponentName() string {
	return m.name
}

func TestGetCommonName(t *testing.T) {
	type args struct {
		instance runtime.Object
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "CommonObject instance",
			args: args{
				instance: &mockCommonObject{name: "test-component"},
			},
			want:    "test-component",
			wantErr: false,
		},
		{
			name: "Unstructured instance",
			args: args{
				instance: &unstructured.Unstructured{
					Object: map[string]interface{}{
						"kind": "TestKind",
					},
				},
			},
			want:    "testkind",
			wantErr: false,
		},
		{
			name: "Invalid instance",
			args: args{
				instance: &runtime.Unknown{},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommonName(tt.args.instance)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommonName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCommonName() = %v, want %v", got, tt.want)
			}
		})
	}
}
