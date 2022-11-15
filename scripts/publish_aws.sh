#!/bin/bash
# ships templates to S3 in ~all regions in the "honeycomb-integrations-$REGION" bucket

set -e

REGIONS="us-east-1 us-east-2 us-west-1 us-west-2 ap-south-1 ap-northeast-2 ap-southeast-1 ap-southeast-2 ap-northeast-1 ca-central-1 eu-central-1 eu-west-1 eu-west-2 eu-west-3 sa-east-1"

# if DRYRUN is set to anything, turn it into the awscli switch
[[ -n "${DRYRUN}" ]] && DRYRUN="--dryrun"

echo "+++ Uploading templates"
for REGION in ${REGIONS}; do
	DEPLOY_ROOT=s3://honeycomb-integrations-${REGION}/cloudformation-templates
	aws s3 sync "${DRYRUN}" templates/ "${DEPLOY_ROOT}"
done
