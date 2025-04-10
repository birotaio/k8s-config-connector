
VERSION=$(cat version/VERSION)

# add nodeSelector to cnrm-system
mkdir -p _manifest/templates/generated
MANIFEST="operator/channels/packages/configconnector/${VERSION}/cluster/gcp-identity/0-cnrm-system.yaml"
ytt -f $MANIFEST -f ci/patch-manifest.yaml > ${MANIFEST}.tmp
mv ${MANIFEST}.tmp $MANIFEST


mkdir -p _manifest/templates/crds
mkdir -p _manifest/templates/operator
rm -f _manifest/templates/crds/*.yaml


CRDS=(
  "computeforwardingrules.compute"
  "computeaddresses.compute"
  "computeforwardingrules.compute"
  "computeserviceattachments.compute"
  "runservices.run"
  "cloudschedulerjobs.cloudscheduler"
  "iamserviceaccounts.iam"
  "iampolicymembers.iam"
  "iampartialpolicies.iam"
  "computeregionnetworkendpointgroups.compute"
  "computebackendservices.compute"
  "computeurlmaps.compute"
  "computetargethttpproxies.compute"
  "dnsrecordsets.dns"
  "pubsubtopics.pubsub"
  "eventarctriggers.eventarc"
)

for crd in "${CRDS[@]}"; do
  cp config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_${crd}.cnrm.cloud.google.com.yaml _manifest/templates/crds
done

cp operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectors.yaml _manifest/templates/crds/
cp operator/config/crd/bases/core.cnrm.cloud.google.com_configconnectorcontexts.yaml _manifest/templates/crds/


make build-rbac-manifests
cp operator/config/manager/manager.yaml _manifest/templates/operator/
cp config/installbundle/release-manifests/rbac.yaml _manifest/templates/operator/

mv operator/Dockerfile .
cp -r * ../build
