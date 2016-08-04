package testvariables

import "github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"

const (
	OperationFailsKey = "OPERATION_FAILS"

	GenerateManifestServiceDeploymentFileKey = "GENERATE_MANIFEST_SERVICE_DEPLOYMENT_FILE"
	GenerateManifestPlanFileKey              = "GENERATE_MANIFEST_PLAN_FILE"
	GenerateManifestArbitraryParamsFileKey   = "GENERATE_MANIFEST_ARBITRARY_PARAMS_FILE"
	GenerateManifestPreviousManifestFileKey  = "GENERATE_MANIFEST_PREVIOUS_MANIFEST_FILE"
	GenerateManifestPreviousPlanFileKey      = "GENERATE_MANIFEST_PREVIOUS_PLAN_FILE"

	BindingIdFileKey       = "BINDING_ID_FILE_KEY"
	BindingVmsFileKey      = "BINDING_VMS_FILE_KEY"
	BindingManifestFileKey = "BINDING_MANIFEST_FILE_KEY"
	BindingParamsFileKey   = "BINDING_PARAMS_FILE_KEY"

	DashboardURLInstanceIDKey = "DASHBOARD_URL_INSTANCE_ID_FILE"
	DashboardURLPlanKey       = "DASHBOARD_URL_PLAN_FILE"
	DashboardURLManifestKey   = "DASHBOARD_URL_MANIFEST_FILE"

	DoNotImplementInterfacesKey = "DO_NOT_IMPLEMENT_INTERFACES"

	ErrAppGuidNotProvided   = "no app guid"
	ErrBindingAlreadyExists = "binding already exists"
)

var SuccessfulBinding = serviceadapter.Binding{
	RouteServiceURL: "a route",
	SyslogDrainURL:  "a url",
	Credentials: map[string]interface{}{
		"binding": "this binds",
	},
}
