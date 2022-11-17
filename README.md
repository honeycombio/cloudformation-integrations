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

[Click here](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=honeycomb-cloudwatch-integration&templateURL=https://s3.amazonaws.com/honeycomb-builds/honeycombio/integrations-for-aws/LATEST/templates/cloudwatch-logs-json.yml) to launch the AWS Cloudformation Console to create the integration stack. 


You will need one stack per Cloudwatch Log Group. 

Required inputs:

- Stack Name
- Cloudwatch Log Group Name
- Your honeycomb write key (optionally encrypted)
- Target honeycomb dataset

Optional inputs:

- Sample rate
- The ID of the AWS Key Management Service key used to encrypt your write key. If your write key is not encrypted, do not set a value here

### Cloudwatch Metrics

[Click here](https://console.aws.amazon.com/cloudformation/home#/stacks/new?stackName=honeycomb-cloudwatch-integration&templateURL=https://s3.amazonaws.com/honeycomb-builds/honeycombio/integrations-for-aws/LATEST/templates/cloudwatch-logs-json.yml) to launch the AWS Cloudformation Console to create the integration stack.


You will need one stack per Cloudwatch Log Group.

Required inputs:

- Stack Name
- Cloudwatch Log Group Name
- Your honeycomb write key (optionally encrypted)
- Target honeycomb dataset

Optional inputs:

- Sample rate
- The ID of the AWS Key Management Service key used to encrypt your write key. If your write key is not encrypted, do not set a value here
