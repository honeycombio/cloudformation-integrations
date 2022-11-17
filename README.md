# Honeycomb CloudFormation Integrations

[![OSS Lifecycle](https://img.shields.io/osslifecycle/honeycombio/cloudformation-integrations)](https://github.com/honeycombio/home/blob/main/honeycomb-oss-lifecycle-and-practices.md)

This repository contains a collection of [CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html) templates for resources in [AWS](https://aws.amazon.com/) to send 
observabiity data to [Honeycomb](https://www.honeycomb.io/).

## How does this work?

![AWS Integrations architecture](docs/overview.png?raw=true)

Note: [Terraform modules](https://github.com/honeycombio/terraform-aws-integrations) are also available that support AWS to Honeycomb Integrations.

## Use

### **Integrations supported are:**

* [CloudWatch Logs](README.md#cloudwatch-logs)

  * [RDS Cloudwatch Logs](README.md#rds-logs)


* [CloudWatch Metrics](README.md#cloudwatch-metrics)


* [Logs from a S3 bucket](FIX ME)


* [Kinesis Firehose Stream to Honeycomb](README.md#kinesis-firehose-stream-to-honeycomb)

### Caveats & Troubleshooting

* If on initial set up the stack create fails - we recommend a complete delete and recreate to ensure all dependencies properly get created.


## Cloudwatch Logs

### [Click here](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=cloudwatch-logs&templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/cloudwatch-logs.yml) 
to launch the AWS Cloudformation Console to create the integration stack. 


This stack supports integrating with up to 5 Cloudwatch log groups.

**Required inputs:**

```
- Stack Name


- Honeycomb API Key: Your Honeycomb Team's API key.


- Honeycomb Dataset: The target Honeycomb dataset for the Stream to publish to.


- Log Group Name: At least 1 Cloudwatch log group name


- S3 Failure Bucket Arn: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
```

## Cloudwatch Metrics

### [Click here](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=cloudwatch-metrics&templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/cloudwatch-metrics.yml) 
to launch the AWS Cloudformation Console to create the integration stack.


This stack supports integrating with all metrics flowing to Cloudwatch Metrics. Optional filters are available to include or exclude AWS services - without them all metrics will be turned on.

**Required inputs:**
```
- Stack Name


- Honeycomb API Key: Your Honeycomb Team's API key.


- Honeycomb Dataset: The target Honeycomb dataset for the Stream to publish to.


- S3 Failure Bucket Arn: The ARN of the S3 Bucket that will store any events that failed to be sent to Honeycomb.
```

**Callout:**

// TODO: ADD SUPPORT FOR THIS OR REMOVE THIS SECTION. IS SUPPORTED BY TF ATM.

Optional input to reduce the metrics from Services being sent to Honeycomb:

```
- Namespace Include

OR (cannot have both)

- Namespace Exclude
```

## Kinesis Firehose stream to Honeycomb

### [Click here](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=kinesis-firehose-stream&templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/kinesis-firehose.yml)
to launch the AWS Cloudformation Console to create the integration stack.


This stack supports creating and managing a Kinesis Firehose that streams data received to Honeycomb.

**Required inputs:**
```
- Stack Name


- Name: A name for this Kinesis Firehose Delivery Stream


- Honeycomb API Key: Your Honeycomb Team's API key.


- Honeycomb Dataset: The target Honeycomb dataset for the Stream to publish to.


- S3 Failure Bucket Arn: The ARN of the S3 Bucket that will store any events that failed to be sent to Honeycomb.
```

## RDS Cloudwatch Logs

### [Click here](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=rds-logs&templateURL=https://honeycomb-builds.s3.amazonaws.com/cloudformation-templates/latest/rds-logs.yml)
to launch the AWS Cloudformation Console to create the integration stack.


This stack supports integrating RDS logs from Cloudwatch to a Kinesis Firehose that includes a data transform to **structure** the logs before sending them to Honeycomb.

**Required inputs:**
```
- Stack Name


- DB Engine Type: Select the Engine type of your RDS database.


- Honeycomb API Key: Your Honeycomb Team's API key.


- Honeycomb Dataset: The target Honeycomb dataset for the Stream to publish to.


- Log Group Name: At least 1 Cloudwatch log group name to which the DB is sending logs


- S3 Failure Bucket Arn: The ARN of the S3 Bucket that will store any logs that failed to be sent to Honeycomb.
```
