package mapcloudrunjobadminprovider

import (
	"errors"

	admin "github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/gen/mattilsynet/me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin"
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

func Get(manifest *admin.ManagedEnvironmentGcpManifest) (*admin.ManagedEnvironmentGcpManifest, error) {
	adminResult := admin.Get(*manifest)
	return FromCmResultToIdomaticGo(adminResult)
}

func Update(manifest *admin.ManagedEnvironmentGcpManifest) (*admin.ManagedEnvironmentGcpManifest, error) {
	adminResult := admin.Update(*manifest)
	return FromCmResultToIdomaticGo(adminResult)
}

func Delete(manifest *admin.ManagedEnvironmentGcpManifest) (*admin.ManagedEnvironmentGcpManifest, error) {
	adminResult := admin.Delete(*manifest)
	return FromCmResultToIdomaticGo(adminResult)
}

func FromCmResultToIdomaticGo(adminResult cm.Result[admin.ErrorShape, admin.ManagedEnvironmentGcpManifest, admin.Error]) (*admin.ManagedEnvironmentGcpManifest, error) {
	if adminResult.Err() != nil {
		return nil, errors.New(adminResult.Err().Message)
	}
	return adminResult.OK(), nil
}
