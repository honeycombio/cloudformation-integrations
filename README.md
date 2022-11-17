# Honeycomb CloudFormation Integrations

[![OSS Lifecycle](https://img.shields.io/osslifecycle/honeycombio/cloudformation-integrations)](https://github.com/honeycombio/home/blob/main/honeycomb-oss-lifecycle-and-practices.md)

This repository contains a collection of CloudFormation templates for resources in [AWS](https://aws.amazon.com/) to send 
observabiity data to [Honeycomb](https://www.honeycomb.io/).

## How does this work?

![AWS Integrations architecture](docs/overview.png?raw=true)


## **Integrations supported are:**

* [CloudWatch Logs](FIX ME)

  * [RDS Logs](FIX ME)


* [CloudWatch Metrics](FIX ME)


* [Logs from a S3 bucket](FIX ME)


## Use

### Cloudwatch Logs

### [Click here](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=cloudwatch-logs&templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/cloudwatch-logs.yml) 
to launch the AWS Cloudformation Console to create the integration stack. 


This stack supports integrating with up to 5 Cloudwatch log groups.

**Required inputs:**

- Stack Name


- Honeycomb API Key: Your Honeycomb Team's API key.


- Honeycomb Dataset: The target Honeycomb dataset for the Stream to publish to.


- Log Group Name: At least 1 Cloudwatch log group name


- S3 Failure Bucket Arn: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.

### Cloudwatch Metrics

### [Click here](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=cloudwatch-metrics&templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/cloudwatch-metrics.yml) 
to launch the AWS Cloudformation Console to create the integration stack.


This stack supports integrating with all metrics flowing to Cloudwatch Metrics. Optional filters are available to include or exclude AWS services - without them all metrics will be turned on.

**Required inputs:**

- Stack Name


- Honeycomb API Key: Your Honeycomb Team's API key.


- Honeycomb Dataset: The target Honeycomb dataset for the Stream to publish to.


- S3 Failure Bucket Arn: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.

**Callout:**

// TODO: ADD SUPPORT FOR THIS OR REMOVE THIS SECTION. IS SUPPORTED BY TF ATM.

Optional input to reduce the metrics from Services being sent to Honeycomb:

- Namespace Include

OR

- Namespace Exclude
