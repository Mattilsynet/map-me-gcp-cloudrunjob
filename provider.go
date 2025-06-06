package main

import (
	"context"
	"strings"

	r "cloud.google.com/go/run/apiv2"
	runpb "cloud.google.com/go/run/apiv2/runpb"
	"github.com/Mattilsynet/map-me-gcp-cloudrunjob/bindings/exports/mattilsynet/me_gcp_cloudrun_job_admin/me_gcp_cloudrun_job_admin"
	"github.com/Mattilsynet/map-me-gcp-cloudrunjob/bindings/mattilsynet/me_gcp_cloudrun_job_admin/types"
	"github.com/Mattilsynet/mapis/gen/go/managedgcpenvironment/v1"
	sdk "go.wasmcloud.dev/provider"
	"google.golang.org/api/option"
	wrpc "wrpc.io/go"
	wrpcnats "wrpc.io/go/nats"
)

const (
	HEALTHY = true
)

type Config struct {
	ProjectId                      string `json:"project_id"`
	Location                       string `json:"location"`
	Image                          string `json:"image"`
	CloudrunAdminServiceAccountJwt string `json:"gcp_sa_jwt"`
}
type Secret struct {
	CloudrunAdminServiceAccountJwt []byte `json:"gcp_sa_jwt"`
}

// TODO:
// 1. refactor out all the cloud run admin sdk to pkg
// 2. centralize the DRY code in each of update, delete and get
// 3. extract config and secret structs to pkg
type CloudRunJobAdmin struct {
	provider  *sdk.WasmcloudProvider
	configMap map[string]*Config
	secretMap map[string]*Secret
}

func NewCloudRunJobAdmin() CloudRunJobAdmin {
	return CloudRunJobAdmin{
		configMap: make(map[string]*Config),
		secretMap: make(map[string]*Secret),
	}
}

func (cl *CloudRunJobAdmin) AddTarget(target string, secret *Secret, config *Config) {
	cl.configMap[target] = config
	cl.secretMap[target] = secret
}

func (cl *CloudRunJobAdmin) RemoveTarget(target string) {
	delete(cl.configMap, target)
	delete(cl.secretMap, target)
}

func (cl *CloudRunJobAdmin) Shutdown() {
	for target := range cl.configMap {
		cl.RemoveTarget(target)
	}
	cl.configMap = nil
	cl.secretMap = nil
}

// TODO:
// add validation rules in query / command api towards protobuf incomming
func CrjEnvsFrom(me *managedgcpenvironment.ManagedGcpEnvironment) []*runpb.EnvVar {
	spec := me.Spec
	envVars := []*runpb.EnvVar{
		{
			Name:   "name",
			Values: &runpb.EnvVar_Value{Value: me.Metadata.Name},
		},
		{
			Name:   "group",
			Values: &runpb.EnvVar_Value{Value: spec.Group},
		},
		{
			Name:   "parent_folder_id",
			Values: &runpb.EnvVar_Value{Value: spec.ParentFolderId},
		},
		{
			Name:   "team_ar_repo_id",
			Values: &runpb.EnvVar_Value{Value: spec.TeamArRepoId},
		},
		{
			Name:   "budget_amount",
			Values: &runpb.EnvVar_Value{Value: spec.BudgetAmount},
		},
		{
			Name:   "dns_zone_name",
			Values: &runpb.EnvVar_Value{Value: spec.DnsZoneName},
		},
		{
			Name:   "slack_channel_email",
			Values: &runpb.EnvVar_Value{Value: spec.Email},
		},
	}
	return envVars
}

func (cl *CloudRunJobAdmin) Update(ctx__ context.Context, manifest *types.ManagedEnvironmentGcpManifest) (*wrpc.Result[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest, types.Error], error) {
	isLinked, target := cl.isLinkedWith(ctx__)
	if !isLinked {
		unauthorized := types.Error{
			ErrorType: types.NewErrorTypeUnauthorized(),
			Message:   "Unauthorized, target not linked",
		}
		cl.provider.Logger.Error("target", "Unauthorized, target not linked", target)
		return wrpc.Err[types.ManagedEnvironmentGcpManifest](unauthorized), nil
	}
	config := cl.configMap[target]
	secret := cl.secretMap[target]
	me := managedgcpenvironment.ManagedGcpEnvironment{}
	me.UnmarshalVT(manifest.Bytes)
	jwtOpt := option.WithCredentialsJSON(secret.CloudrunAdminServiceAccountJwt)
	svc, err := r.NewJobsClient(ctx__, jwtOpt)
	if err != nil {
		cl.provider.Logger.Error("error creating job", "err", err)
		return nil, err
	}
	gcpProjectToPutCrj := config.ProjectId
	imageToUse := config.Image
	envVars := CrjEnvsFrom(&me)
	// teams seed project
	// TODO: move underneath to pkg
	// Format: projects/{project}/locations/{location}/jobs/{jobId}
	jobId := "projects/" + gcpProjectToPutCrj + "/locations/" + config.Location + "/jobs/" + me.Metadata.Name
	updateReq := runpb.UpdateJobRequest{
		// INFO: AllowMissing will create the job if it doesn't exist
		AllowMissing: true,
		Job: &runpb.Job{
			Name: jobId,
			Template: &runpb.ExecutionTemplate{
				Template: &runpb.TaskTemplate{
					Containers: []*runpb.Container{
						{
							Image: imageToUse,
							Env:   envVars,
						},
					},
				},
			},
		},
	}
	_, err = svc.UpdateJob(ctx__, &updateReq)
	if err != nil {
		unknownErr := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
		return wrpc.Err[types.ManagedEnvironmentGcpManifest](unknownErr), err
	}
	updatedManifestBytes, err := me.MarshalVT()
	if err != nil {
		errExists := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
		return wrpc.Err[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest](errExists), nil
	}
	manifest.Bytes = updatedManifestBytes
	return wrpc.Ok[types.Error](*manifest), nil
}

func (cl *CloudRunJobAdmin) Delete(ctx__ context.Context, manifest *types.ManagedEnvironmentGcpManifest) (*wrpc.Result[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest, types.Error], error) {
	isLinked, target := cl.isLinkedWith(ctx__)
	if !isLinked {
		unauthorized := types.Error{
			ErrorType: types.NewErrorTypeUnauthorized(),
			Message:   "Unauthorized, target not linked",
		}
		cl.provider.Logger.Error("target", "Unauthorized, target not linked", target)
		return wrpc.Err[types.ManagedEnvironmentGcpManifest](unauthorized), nil
	}
	config := cl.configMap[target]
	secret := cl.secretMap[target]
	me := managedgcpenvironment.ManagedGcpEnvironment{}
	me.UnmarshalVT(manifest.Bytes)
	jwtOpt := option.WithCredentialsJSON(secret.CloudrunAdminServiceAccountJwt)
	svc, err := r.NewJobsClient(ctx__, jwtOpt)
	if err != nil {
		cl.provider.Logger.Error("error creating job", "err", err)
		return nil, err
	}
	gcpProjectToPutCrj := config.ProjectId
	jobId := "projects/" + gcpProjectToPutCrj + "/locations/" + config.Location + "/jobs/" + me.Metadata.Name
	deleteReq := runpb.DeleteJobRequest{
		Name: jobId,
	}
	_, err = svc.DeleteJob(ctx__, &deleteReq)
	if err != nil {
		unknownErr := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
		return wrpc.Err[types.ManagedEnvironmentGcpManifest](unknownErr), err
	}
	updatedManifestBytes, err := me.MarshalVT()
	if err != nil {
		errExists := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
		return wrpc.Err[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest](errExists), nil
	}
	manifest.Bytes = updatedManifestBytes
	return wrpc.Ok[types.Error](*manifest), nil
}

func (cl *CloudRunJobAdmin) Get(ctx__ context.Context, manifest *types.ManagedEnvironmentGcpManifest) (*wrpc.Result[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest, types.Error], error) {
	isLinked, target := cl.isLinkedWith(ctx__)
	if !isLinked {
		unauthorized := types.Error{
			ErrorType: types.NewErrorTypeUnauthorized(),
			Message:   "Unauthorized, target not linked",
		}
		cl.provider.Logger.Error("target", "Unauthorized, target not linked", target)
		return wrpc.Err[types.ManagedEnvironmentGcpManifest](unauthorized), nil
	}
	config := cl.configMap[target]
	secret := cl.secretMap[target]
	me := managedgcpenvironment.ManagedGcpEnvironment{}
	me.UnmarshalVT(manifest.Bytes)
	jwtOpt := option.WithCredentialsJSON(secret.CloudrunAdminServiceAccountJwt)
	svc, err := r.NewJobsClient(ctx__, jwtOpt)
	lastExecutionStatus := "UNKNOWN"
	if err != nil {
		updatedManifestbytes, innerErr := GetMeWithStatusAsBytes(&me, lastExecutionStatus)
		if innerErr != nil {
			errUnknown := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
			return &wrpc.Result[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest, types.Error]{Ok: manifest, Err: &errUnknown}, nil
		}
		manifest.Bytes = updatedManifestbytes
		cl.provider.Logger.Error("error getting job", "err", err)
		errUnknown := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
		return &wrpc.Result[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest, types.Error]{Ok: manifest, Err: &errUnknown}, err
	}
	gcpProjectToPutCrj := config.ProjectId
	jobId := "projects/" + gcpProjectToPutCrj + "/locations/" + config.Location + "/jobs/" + me.Metadata.Name
	getReq := runpb.GetJobRequest{
		Name: jobId,
	}
	job, err := svc.GetJob(ctx__, &getReq)
	if err != nil {
		updatedManifestbytes, innerErr := GetMeWithStatusAsBytes(&me, lastExecutionStatus)
		if innerErr != nil {
			errExists := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
			return wrpc.Err[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest](errExists), nil
		}
		manifest.Bytes = updatedManifestbytes
		unknownErr := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
		if strings.Contains(strings.ToLower(err.Error()), "notfound") {
			unknownErr.Message = "Job not found"
			return &wrpc.Result[types.ManagedEnvironmentGcpManifest, types.Error]{Ok: manifest, Err: &unknownErr}, nil
		}
		return &wrpc.Result[types.ManagedEnvironmentGcpManifest, types.Error]{Ok: manifest, Err: &unknownErr}, err
	}

	lastExecutionStatus = "NOT_RUN_YET"
	if job != nil && job.LatestCreatedExecution != nil {
		lastExecutionStatus = job.LatestCreatedExecution.CompletionStatus.String()
	}
	updatedManifestbytes, err := GetMeWithStatusAsBytes(&me, lastExecutionStatus)
	if err != nil {
		errExists := types.Error{Message: err.Error(), ErrorType: types.NewErrorTypeUnknown()}
		return wrpc.Err[me_gcp_cloudrun_job_admin.ManagedEnvironmentGcpManifest](errExists), nil
	}
	manifest.Bytes = updatedManifestbytes
	return wrpc.Ok[types.Error](*manifest), nil
}

func (cl *CloudRunJobAdmin) isLinkedWith(ctx context.Context) (bool, string) {
	header, ok := wrpcnats.HeaderFromContext(ctx)
	if !ok {
		return false, ""
	}
	target := header.Get("source-id")
	if cl.configMap[target] == nil || cl.secretMap[target] == nil {
		return false, ""
	}
	return true, target
}

func GetMeWithStatusAsBytes(managedEnvironment *managedgcpenvironment.ManagedGcpEnvironment, lastExecutionStatus string) ([]byte, error) {
	managedEnvironment.Status = &managedgcpenvironment.ManagedGcpEnvironmentStatus{
		StatusMap: map[string]string{
			"Status": lastExecutionStatus,
		},
	}
	return managedEnvironment.MarshalVT()
}
