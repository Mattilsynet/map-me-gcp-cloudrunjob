//go:generate go run github.com/bytecodealliance/wasm-tools-go/cmd/wit-bindgen-go generate --world component --out gen ./wit
package main

import (
	"log/slog"

	megcpcloudrunjobadmin "github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/gen/mattilsynet/me-gcp-cloudrun-job-admin/me-gcp-cloudrun-job-admin"
	managedenvironment "github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/pkg/managed-environment"
	"github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/pkg/manifest"
	idiomaticgo "github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/pkg/map-me-gcp-cloudrunjob"
	"github.com/Mattilsynet/map-me-gcp-cloudrunjob/component/pkg/nats"
	"go.wasmcloud.dev/component/log/wasilog"
)

const (
	GET    = "map.get"
	UPDATE = "map.update"
	DELETE = "map.delete"
)

var (
	conn   *nats.Conn
	logger *slog.Logger
)

func init() {
	logger = wasilog.ContextLogger("someComponent")
	logger.Info("Initializing component")
	conn = nats.NewConn()
	logger = wasilog.ContextLogger("map-gcp-cloudrun-component")
	conn.RegisterSubscription(SubscriptionHandler)
}

func SubscriptionHandler(natsMsg *nats.Msg) {
	data := natsMsg.Data
	managedGcpEnvironment, err := managedenvironment.ToManagedEnvironment(data)
	if err != nil {
		logger.Info("failed to convert data to managedGcpEnvironment", "error", err)
		return
	}
	manifest, err := manifest.ToWitManifest(managedGcpEnvironment)
	if err != nil {
		logger.Info("failed to convert data to managedGcpEnvironment", "error", err)
		return
	}
	var returnedManifest *megcpcloudrunjobadmin.ManagedEnvironmentGcpManifest = nil
	switch natsMsg.Subject {
	case GET:
		returnedManifest, err = idiomaticgo.Get(manifest)
		if err != nil {
			logger.Error("failed to convert data to managedGcpEnvironment", "error", err)
			return
		}
	case UPDATE:
		returnedManifest, err = idiomaticgo.Update(manifest)
		if err != nil {
			logger.Error("failed to convert data to managedGcpEnvironment", "error", err)
			return
		}
	case DELETE:
		returnedManifest, err = idiomaticgo.Delete(manifest)
		if err != nil {
			logger.Error("failed to convert data to managedGcpEnvironment", "error", err)
			return
		}
	default:
		returnedManifest = nil
		logger.Error("failed to recognize nats subject", "error", err, "nats subject got", natsMsg.Subject, "expected ONEOF", "map.get, map.update, map.delete")

	}
	logger.Info("successfully fetched manifest", "manifest", returnedManifest)
}

func main() {}
