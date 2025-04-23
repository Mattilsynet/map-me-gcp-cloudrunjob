package manifest

import (
	admin "github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/gen/mattilsynet/me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin"
	me_gcp "github.com/Mattilsynet/mapis/gen/go/managedgcpenvironment/v1"
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

func ToWitManifest(managedGcpEnvironment *me_gcp.ManagedGcpEnvironment) (*admin.ManagedEnvironmentGcpManifest, error) {
	bytes, err := managedGcpEnvironment.MarshalVT()
	if err != nil {
		return nil, err
	}
	m := admin.ManagedEnvironmentGcpManifest{
		Bytes: cm.ToList(bytes),
	}
	return &m, nil
}

func FromWitManifest(manifest *admin.ManagedEnvironmentGcpManifest) (*me_gcp.ManagedGcpEnvironment, error) {
	managedGcpEnvironmentBytes := manifest.Bytes.Slice()
	managedGcpEnvironment := &me_gcp.ManagedGcpEnvironment{}
	err := managedGcpEnvironment.UnmarshalVT(managedGcpEnvironmentBytes)
	if err != nil {
		return nil, err
	}
	return managedGcpEnvironment, nil
}
