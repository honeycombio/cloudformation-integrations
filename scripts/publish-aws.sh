#!/bin/bash

set -e

# if DRYRUN is set to anything, turn it into the awscli switch
[[ -n "${DRYRUN}" ]] && DRYRUN="--dryrun"

# publish the templates to the builds bucket
DEPLOY_ROOT=s3://honeycomb-builds/honeycombio/cloudformation-templates

echo "+++ Uploading templates"
for TEMPLATE in templates/*; do
	aws s3 sync ${DRYRUN} templates/ "${DEPLOY_ROOT}"
done
