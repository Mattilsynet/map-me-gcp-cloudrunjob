package main

import (
	"log"
	"testing"
	"time"

	me_gcp "github.com/Mattilsynet/mapis/gen/go/managedgcpenvironment/v1"
	metadata "github.com/Mattilsynet/mapis/gen/go/meta/v1"
	"github.com/nats-io/nats.go"
)

func TestSuper(t *testing.T) {
	me := &me_gcp.ManagedGcpEnvironment{}
	spec := me_gcp.ManagedGcpEnvironmentSpec{}
	oMeta := &metadata.ObjectMeta{}
	oMeta.Name = "test-job"
	me.Metadata = oMeta
	spec.BudgetAmount = "100"
	spec.DnsZoneName = "DZone"
	spec.TeamArRepoId = "Super-repo"
	spec.Email = "superduper@super-mail.com"
	spec.Group = "group2"
	spec.MapspaceRef = "map-ops-dev-c2c8"
	spec.ParentFolderId = "123123123"
	me.Spec = &spec
	meAsBytes, err := me.MarshalVT()
	if err != nil {
		t.Fail()
	}
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Println("Error connecting to NATS server:", err)
		return
	}
	err = nc.Publish("map.delete", meAsBytes)
	if err != nil {
		log.Println("Error publishing message:", err)
	}
	time.Sleep(1 * time.Second)
}
