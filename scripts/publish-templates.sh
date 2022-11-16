#!/bin/bash

set -e

#local testing
LATEST_COMMIT_SHA="`git rev-parse HEAD`"
FOLDER="testing/$LATEST_COMMIT_SHA"

#PRs or Releases
if [ -n "$1" ] && [ $1 == "prs" ] && [ -n "$2" ]; then
  FOLDER="prs/$2"
elif [ -n "$1" ] && [ $1 == "release" ]; then
  FOLDER="latest"
fi

# if DRYRUN is set to anything, turn it into the awscli switch
[[ -n "${DRYRUN}" ]] && DRYRUN="--dryrun"

# publish the templates to the builds bucket
DEPLOY_ROOT=s3://honeycomb-builds/cloudformation-templates

echo "+++ Uploading templates to ${DEPLOY_ROOT}/$FOLDER/"
aws s3 sync ${DRYRUN} templates/ "${DEPLOY_ROOT}/$FOLDER"
echo "+++ Published templates to ${DEPLOY_ROOT}/$FOLDER/"
