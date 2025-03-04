
VERSION=$(cat version/VERSION)
ALLOWED_CRDS="ComputeAddress,ComputeForwardingRule,ComputeServiceAttachment,RunService,CloudSchedulerJob,IAMServiceAccount,IAMPolicyMember,IAMPartialPolicy,ComputeRegionNetworkEndpointGroup,ComputeBackendService,ComputeURLMap,ComputeTargetHTTPProxy,DNSRecordSet,PubSubTopic,EventarcTrigger"

# add nodeSelector to cnrm-system
mkdir -p _manifest/templates/generated
MANIFEST="operator/channels/packages/configconnector/${VERSION}/cluster/gcp-identity/0-cnrm-system.yaml"
ytt -f $MANIFEST -f ci/patch-manifest.yaml > ${MANIFEST}.tmp
mv ${MANIFEST}.tmp $MANIFEST

pushd fix-manifests &> /dev/null
go mod tidy
go run ./fix-operator.go ${ALLOWED_CRDS}
popd &> /dev/null

mv operator/Dockerfile .
cp -r * ../build
