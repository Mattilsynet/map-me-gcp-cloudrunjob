// Code generated by wit-bindgen-go. DO NOT EDIT.

package megcpcloudrunjobadmin

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// This file contains wasmimport and wasmexport declarations for "mattilsynet:me-gcp-cloudrun-job-admin@0.1.0".

//go:wasmimport mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0 update
//go:noescape
func wasmimport_Update(manifest0 *uint8, manifest1 uint32, result *cm.Result[ErrorShape, ManagedEnvironmentGcpManifest, Error])

//go:wasmimport mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0 get
//go:noescape
func wasmimport_Get(manifest0 *uint8, manifest1 uint32, result *cm.Result[ErrorShape, ManagedEnvironmentGcpManifest, Error])

//go:wasmimport mattilsynet:me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin@0.1.0 delete
//go:noescape
func wasmimport_Delete(manifest0 *uint8, manifest1 uint32, result *cm.Result[ErrorShape, ManagedEnvironmentGcpManifest, Error])
