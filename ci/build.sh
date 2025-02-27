
VERSION=$(cat version/VERSION)
ALLOWED_CRDS="ComputeAddress,ComputeForwardingRule,ComputeServiceAttachment,RunService,CloudSchedulerJob,IAMServiceAccount,IAMPolicyMember,IAMPartialPolicy,ComputeRegionNetworkEndpointGroup,ComputeBackendService,ComputeURLMap,ComputeTargetHTTPProxy,DNSRecordSet,PubSubTopic,EventarcTrigger"

pushd "$(dirname "$0")" &> /dev/null

# add nodeSelector to cnrm-system
MANIFEST="../operator/channels/packages/configconnector/${VERSION}/cluster/gcp-identity/0-cnrm-system.yaml"
ytt -f $MANIFEST -f patch-manifest.yaml > ../_manifest/templates/generated/generated.yaml

cd ../fix-manifests
go mod tidy
go run ./fix-operator.go ${ALLOWED_CRDS}
