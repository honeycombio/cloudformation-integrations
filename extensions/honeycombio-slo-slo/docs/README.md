# Honeycomb::SLO::SLO

Honeycomb SLOs allows you to define and monitor Service Level Objectives (SLOs) for your organization.

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "Type" : "Honeycomb::SLO::SLO",
    "Properties" : {
        "<a href="#name" title="Name">Name</a>" : <i>String</i>,
        "<a href="#dataset" title="Dataset">Dataset</a>" : <i>String</i>,
        "<a href="#description" title="Description">Description</a>" : <i>String</i>,
        "<a href="#timeperiod" title="TimePeriod">TimePeriod</a>" : <i>Integer</i>,
        "<a href="#targetpercentage" title="TargetPercentage">TargetPercentage</a>" : <i>Double</i>,
        "<a href="#sli" title="SLI">SLI</a>" : <i>String</i>
    }
}
</pre>

### YAML

<pre>
Type: Honeycomb::SLO::SLO
Properties:
    <a href="#name" title="Name">Name</a>: <i>String</i>
    <a href="#dataset" title="Dataset">Dataset</a>: <i>String</i>
    <a href="#description" title="Description">Description</a>: <i>String</i>
    <a href="#timeperiod" title="TimePeriod">TimePeriod</a>: <i>Integer</i>
    <a href="#targetpercentage" title="TargetPercentage">TargetPercentage</a>: <i>Double</i>
    <a href="#sli" title="SLI">SLI</a>: <i>String</i>
</pre>

## Properties

#### Name

The name of the SLO.

_Required_: Yes

_Type_: String

_Minimum Length_: <code>1</code>

_Maximum Length_: <code>120</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Dataset

The dataset this SLO is created in. Must be the same dataset as the SLI.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### Description

A description of the SLO.

_Required_: No

_Type_: String

_Maximum Length_: <code>1023</code>

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TimePeriod

The time period, in days, over which your SLO will be evaluated.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### TargetPercentage

The percentage of qualified events that you expect to succeed during the time period.

_Required_: Yes

_Type_: Double

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### SLI

The alias of the Derived Column that will be used as the SLI to indicate event success.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return Values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, Ref returns the ID.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).

#### ID

The ID of the SLO.

