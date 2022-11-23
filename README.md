# Honeycomb CloudFormation Integrations

[![OSS Lifecycle](https://img.shields.io/osslifecycle/honeycombio/cloudformation-integrations)](https://github.com/honeycombio/home/blob/main/honeycomb-oss-lifecycle-and-practices.md)

This repository contains a collection of [CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html) templates for resources in [AWS](https://aws.amazon.com/) to send observability data to [Honeycomb](https://www.honeycomb.io/).

Note: [Terraform modules](https://github.com/honeycombio/terraform-aws-integrations) are also available that support AWS to Honeycomb Integrations.

## How does this work?

![AWS Integrations architecture](docs/overview.png?raw=true)

## Usage

A [quick start template](README.md#quick-start) offers a streamlined path to integrating your AWS environments with Honeycomb.
The quick start template uses all of the per-integration templates below to offer the configuration of many integrations in a single CloudFormation stack.

The quick start template may be suitable for many production purposes, but we encourage you to use per-integration templates in a way that suits your AWS environment.

### Supported Integrations

- [CloudWatch Logs](README.md#cloudwatch-logs)
- [RDS Cloudwatch Logs](README.md#rds-logs)
- [CloudWatch Metrics](README.md#cloudwatch-metrics)
- [Logs from a S3 bucket](README.md#logs-from-a-s3-bucket)
- [Kinesis Firehose Stream to Honeycomb](README.md#kinesis-firehose-stream-to-honeycomb)

### Caveats & Troubleshooting

If the stack creation fails during its initial set up, we recommend a complete deletion and re-creation to ensure the proper creation of all dependencies.

## Quick Start

This 'Quick Start' template allows the configuration of multiple integrations from a single CloudFormation Template.

<a href="https://console.aws.amazon.com/cloudformation/home#/stacks/new?templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/quickstart.yml&stackName=Honeycomb-Integration" target="_blank"><img src="https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png" alt="Launch Stack" /></a>

### Required Inputs

Enter a value for the required input in the UI, or if using the CLI or API, ensure the inclusion of the required input and its value.

<<<<<<< HEAD
- `HoneycombAPIKey`: : Your Honeycomb Team's API Key.
=======
* `HoneycombAPIKey`: : Your Honeycomb Team's API Key.
>>>>>>> 247a6c2fe43b1d598304428bf04ee7c074aa4331

*Note*: All other Parameters are optional, but if no additional parameters are provided, only a S3 Bucket will be created.

## Cloudwatch Logs

This template supports integrating with up to six Cloudwatch Log Groups and shipping them to a Honeycomb dataset.

<a href="https://console.aws.amazon.com/cloudformation/home#/stacks/new?templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/cloudwatch-logs.yml&stackName=Honeycomb-Logs" target="_blank"><img src="https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png" alt="Launch Stack" /></a>

### Required Inputs

Enter a value for each required input in the UI, or if using the CLI or API, ensure the inclusion of each required input and its value.

<<<<<<< HEAD
- `HoneycombAPIKey`: Your Honeycomb Team's API Key.
- `HoneycombDataset`: The target Honeycomb dataset for the Stream to publish to.
- `LogGroupName`: A CloudWatch Log Group name. Additional Log Groups can be added with the `LogGroupNameX` parameters.
- `S3FailureBucketArn`: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
=======
* `HoneycombAPIKey`: Your Honeycomb Team's API Key.
* `HoneycombDataset`: The target Honeycomb dataset for the Stream to publish to.
* `LogGroupName`: A CloudWatch Log Group name. Additional Log Groups can be added with the `LogGroupNameX` parameters.
* `S3FailureBucketArn`: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
>>>>>>> 247a6c2fe43b1d598304428bf04ee7c074aa4331

## Cloudwatch Metrics

This template supports integrating with all metrics flowing to Cloudwatch Metrics and shipping them to a Honeycomb.

<a href="https://console.aws.amazon.com/cloudformation/home#/stacks/new?templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/cloudwatch-metrics.yml&stackName=Honeycomb-CloudMetrics" target="_blank"><img src="https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png" alt="Launch Stack" /></a>

### Required Inputs

Enter a value for each required input in the UI, or if using the CLI or API, ensure the inclusion of each required input and its value.

<<<<<<< HEAD
- `HoneycombAPIKey`: Your Honeycomb Team's API Key.
- `S3FailureBucketArn`: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
=======
* `HoneycombAPIKey`: Your Honeycomb Team's API Key.
* `S3FailureBucketArn`: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
>>>>>>> 247a6c2fe43b1d598304428bf04ee7c074aa4331

## Kinesis Firehose Stream to Honeycomb

This template creates a Kinesis Firehose Stream that streams data received to Honeycomb.

<a href="https://console.aws.amazon.com/cloudformation/home#/stacks/new?templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/kinesis-firehose.yml&stackName=Honeycomb-Kinesis" target="_blank"><img src="https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png" alt="Launch Stack" /></a>

### Required Inputs

Enter a value for each required input in the UI, or if using the CLI or API, ensure the inclusion of each required input and its value.

<<<<<<< HEAD
- `Name`: A Name for this Kinesis Firehose. Must be unique in this Region.
- `HoneycombAPIKey`: Your Honeycomb Team's API Key.
- `HoneycombDataset`: The target Honeycomb dataset for the Stream to publish to.
- `S3FailureBucketArn`: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
=======
* `Name`: A Name for this Kinesis Firehose. Must be unique in this Region.
* `HoneycombAPIKey`: Your Honeycomb Team's API Key.
* `HoneycombDataset`: The target Honeycomb dataset for the Stream to publish to.
* `S3FailureBucketArn`: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
>>>>>>> 247a6c2fe43b1d598304428bf04ee7c074aa4331

### RDS Cloudwatch Logs

This template streams RDS logs from Cloudwatch to a Kinesis Firehose that includes a data transform to **structure** the logs before sending them to Honeycomb.

<a href="https://console.aws.amazon.com/cloudformation/home#/stacks/new?templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/rds-logs.yml&stackName=Honeycomb-RDS-Logs" target="_blank"><img src="https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png" alt="Launch Stack" /></a>

### Required Inputs

Enter a value for each required input in the UI, or if using the CLI or API, ensure the inclusion of each required input and its value.

<<<<<<< HEAD
- `HoneycombAPIKey`: Your Honeycomb Team's API Key.
- `HoneycombDataset`: The target Honeycomb dataset for the Stream to publish to.
- `DBEngineType`: The Engine type of your RDS database. One of `aurora-mysql`, `aurora-postgresql` `mariadb`, `sqlserver`,`mysql`, `oracle`, or `postgresql`.
- `LogGroupName`: A CloudWatch Log Group name for RDS logs. Additional Log Groups can be added with the `LogGroupNameX` parameters.
- `S3FailureBucketArn`: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
=======
* `HoneycombAPIKey`: Your Honeycomb Team's API Key.
* `HoneycombDataset`: The target Honeycomb dataset for the Stream to publish to.
* `DBEngineType`: The Engine type of your RDS database. One of `aurora-mysql`, `aurora-postgresql` `mariadb`, `sqlserver`,`mysql`, `oracle`, or `postgresql`.
* `LogGroupName`: A CloudWatch Log Group name for RDS logs. Additional Log Groups can be added with the `LogGroupNameX` parameters.
* `S3FailureBucketArn`: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
>>>>>>> 247a6c2fe43b1d598304428bf04ee7c074aa4331

## Logs from a S3 bucket

This template supports sending logs flowing to a S3 bucket to Honeycomb.

<a href="https://console.aws.amazon.com/cloudformation/home#/stacks/new?templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/s3-logfile.yml&stackName=Honeycomb-S3-Logs" target="_blank"><img src="https://s3.amazonaws.com/cloudformation-examples/cloudformation-launch-stack.png" alt="Launch Stack" /></a>

### Required Inputs

Enter a value for each required input in the UI, or if using the CLI or API, ensure the inclusion of each required input and its value.

<<<<<<< HEAD
- `HoneycombAPIKey`: Your Honeycomb Team's API Key.
- `HoneycombDataset`: The target Honeycomb dataset for to publish to.
- `ParserType`: The type of log file to parse. Choose one of `alb`, `elb`, `cloudfront`, `keyval`, `json`, `s3-access`, or `vpc-flow`.
- `S3BucketArn`: The ARN of the S3 Bucket storing the logs.
=======
* `HoneycombAPIKey`: Your Honeycomb Team's API Key.
* `HoneycombDataset`: The target Honeycomb dataset for to publish to.
* `ParserType`: The type of log file to parse. Choose one of `alb`, `elb`, `cloudfront`, `keyval`, `json`, `s3-access`, or `vpc-flow`.
* `S3BucketArn`: The ARN of the S3 Bucket storing the logs.
>>>>>>> 247a6c2fe43b1d598304428bf04ee7c074aa4331
