package manifest

import (
	"testing"

	me_gcp "github.com/Mattilsynet/mapis/gen/go/managedgcpenvironment/v1"
	metav1 "github.com/Mattilsynet/mapis/gen/go/meta/v1"
)

func TestIsChanged(t *testing.T) {
	type args struct {
		meGcp *me_gcp.ManagedGcpEnvironment
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Test with nil Status is false",
			args{
				meGcp: &me_gcp.ManagedGcpEnvironment{
					Metadata: &metav1.ObjectMeta{
						ResourceVersion: "1",
					},
				},
			},
			true,
		},
		{
			"Test with nil Status is false",
			args{
				meGcp: &me_gcp.ManagedGcpEnvironment{
					Metadata: &metav1.ObjectMeta{
						ResourceVersion: "1",
					},
					Status: &me_gcp.ManagedGcpEnvironmentStatus{
						StatusMap: map[string]string{"resource-version": "1"},
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if changed := IsChanged(tt.args.meGcp); changed != tt.wantErr {
				t.Errorf("AddResourceVersion() error = %v, wantErr %v", changed, tt.wantErr)
			}
		})
	}
}
