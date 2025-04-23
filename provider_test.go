package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"testing"

	mattilsynet__me_gcp_cloudrun_job_admin__types "github.com/Mattilsynet/map-me-gcp-cloudrunjob/bindings/mattilsynet/me_gcp_cloudrun_job_admin/types"
	"github.com/Mattilsynet/mapis/gen/go/managedgcpenvironment/v1"
	metav1 "github.com/Mattilsynet/mapis/gen/go/meta/v1"
	"github.com/nats-io/nats.go"
	sdk "go.wasmcloud.dev/provider"
	wrpc "wrpc.io/go"
	wrpcnats "wrpc.io/go/nats"
)

func TestCloudRunAdmin_Create(t *testing.T) {
	provider := sdk.WasmcloudProvider{}
	provider.Logger = slog.Default()
	type fields struct {
		provider   *sdk.WasmcloudProvider
		linkedFrom map[string]map[string]string
		linkedTo   map[string]map[string]string
	}
	type args struct {
		ctx__    context.Context
		manifest *mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *wrpc.Result[*mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest, mattilsynet__me_gcp_cloudrun_job_admin__types.Error]
		wantErr bool
	}{
		{
			name: "Test Update",
			fields: fields{
				provider:   &provider,
				linkedFrom: nil,
				linkedTo:   nil,
			},
			args: args{ctx__: context.Background(), manifest: getManifest()},
			want: &wrpc.Result[*mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest, mattilsynet__me_gcp_cloudrun_job_admin__types.Error]{
				Ok: nil,
				Err: &mattilsynet__me_gcp_cloudrun_job_admin__types.Error{
					Message:   "",
					ErrorType: mattilsynet__me_gcp_cloudrun_job_admin__types.NewErrorTypeAlreadyExists(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := &CloudRunJobAdmin{
				provider:  tt.fields.provider,
				configMap: make(map[string]*Config),
				secretMap: make(map[string]*Secret),
			}
			// copy your .config/gcloud/default_application_credentials.json into string underneath for this to work, also use both gal and gaal
			applicationDefaultCredentialsJson := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
			file, err := os.ReadFile(applicationDefaultCredentialsJson)
			if err != nil {
				t.Errorf("CloudRunAdmin.Create() error %v", err)
			}
			secret := &Secret{
				CloudrunAdminServiceAccountJwt: file,
			}
			config := &Config{
				ProjectId: "map-ops-dev-c2c8",
				Location:  "europe-north1",
				Image:     "us-docker.pkg.dev/cloudrun/container/job:latest",
			}
			target := "testTarget"
			ctx := populateNatsHeaderIntoCtx(tt.args.ctx__, target)
			cl.AddTarget(target, secret, config)
			get, err := cl.Get(ctx, tt.args.manifest)
			if err != nil {
				t.Errorf("CloudRunAdmin.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if get.Err == nil {

				me := managedgcpenvironment.ManagedGcpEnvironment{}
				me.UnmarshalVT(get.Ok.Bytes)
				log.Println(&me)
			} else {
				log.Println(get.Err.Message)
			}
		})
	}
}

type headerKey struct{}

func populateNatsHeaderIntoCtx(ctx__ context.Context, target string) context.Context {
	natsHeaderWithTarget := nats.Header{}
	natsHeaderWithTarget.Add("target", target)
	ctx := wrpcnats.ContextWithHeader(ctx__, natsHeaderWithTarget)
	return ctx
}

func getManifest() *mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest {
	me := managedgcpenvironment.ManagedGcpEnvironment{}
	me.Metadata = &metav1.ObjectMeta{}
	me.Metadata.Name = "test-job"
	me.Spec = &managedgcpenvironment.ManagedGcpEnvironmentSpec{
		Group:          "test-group",
		MapspaceRef:    "123123",
		ParentFolderId: "123123",
		DnsZoneName:    "test-test",
		TeamArRepoId:   "test",
		BudgetAmount:   "50",
		Email:          "some-channel@hq-slack.com",
	}
	meBytes, err := me.MarshalVT()
	if err != nil {
		return nil
	}
	return &mattilsynet__me_gcp_cloudrun_job_admin__types.ManagedEnvironmentGcpManifest{
		Bytes: meBytes,
	}
}
