//go:generate wit-bindgen-wrpc go --out-dir bindings --package github.com/Mattilsynet/map-me-gcp-cloudrunjob/bindings wit

package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	server "github.com/Mattilsynet/map-me-gcp-cloudrunjob/bindings"
	"go.wasmcloud.dev/provider"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// INFO: For local development purposes only, to buypass need to have whole secret manager up and running etc
// Remember to delete your builds such that the par.gz file doesn't contain this jwt secret when pushing to git
var gcpadmin = ``

func run() error {
	// Initialize the provider with callbacks to track linked components
	providerHandler := NewCloudRunJobAdmin()
	p, err := provider.New(
		provider.TargetLinkPut(func(link provider.InterfaceLinkDefinition) error {
			return handleNewTargetLink(&providerHandler, link)
		}),
		provider.TargetLinkDel(func(link provider.InterfaceLinkDefinition) error {
			return handleDelTargetLink(&providerHandler, link)
		}),
		provider.HealthCheck(func() string {
			return handleHealthCheck(&providerHandler)
		}),
		provider.Shutdown(func() error {
			return handleShutdown(&providerHandler)
		}),
	)
	if err != nil {
		return err
	}

	// Store the provider for use in the handlers
	providerHandler.provider = p

	// Setup two channels to await RPC and control interface operations
	providerCh := make(chan error, 1)
	signalCh := make(chan os.Signal, 1)

	// Handle RPC operations
	stopFunc, err := server.Serve(p.RPCClient, &providerHandler)
	if err != nil {
		p.Shutdown()
		return err
	}
	//
	// // Handle control interface operations
	go func() {
		err := p.Start()
		providerCh <- err
	}()
	//
	// Shutdown on SIGINT
	signal.Notify(signalCh, syscall.SIGINT)
	//
	// // Run provider until either a shutdown is requested or a SIGINT is received
	select {
	case err = <-providerCh:
		stopFunc()
		return err
	case <-signalCh:
		p.Shutdown()
		stopFunc()
	}

	return nil
}

// TODO: add validation of link config and secret
func handleNewTargetLink(handler *CloudRunJobAdmin, link provider.InterfaceLinkDefinition) error {
	handler.provider.Logger.Info("Handling new target link", "link", link)
	secret := Secret{
		CloudrunAdminServiceAccountJwt: []byte(link.TargetSecrets["map-me-gcp-cloudrunjob-sa"].String.Reveal()),
	}
	config := Config{
		ProjectId: link.TargetConfig["project_id"],
		Location:  link.TargetConfig["location"],
		Image:     link.TargetConfig["image"],
	}
	if config.ProjectId == "" || config.Location == "" || config.Image == "" {
		handler.provider.Logger.Error("Missing config for target link", "link", link)
		return nil
	}
	// INFO: Local development
	if len(secret.CloudrunAdminServiceAccountJwt) == 0 {
		// INFO: If you want to locally test component without secret manager, then uncoment underneath and comment return nil, add your cloud run admin and act as json inside the gcpadmin variable at the top of this file
		// secret.CloudrunAdminServiceAccountJwt = []byte(gcpadmin)
		// handler.provider.Logger.Info("using local development secret")
		handler.provider.Logger.Error("No secret found for target link", "link", link)
		return nil
	}
	handler.AddTarget(link.SourceID, &secret, &config)
	return nil
}

func handleDelTargetLink(handler *CloudRunJobAdmin, link provider.InterfaceLinkDefinition) error {
	handler.provider.Logger.Info("Handing del source link", "link", link)
	handler.RemoveTarget(link.Target)
	return nil
}

// TODO: Add a check towards google cloud to see if we got a connection as part of the health check
func handleHealthCheck(handler *CloudRunJobAdmin) string {
	return "provider healthy"
}

func handleShutdown(handler *CloudRunJobAdmin) error {
	handler.provider.Logger.Info("Handling shutdown")
	handler.Shutdown()
	return nil
}
