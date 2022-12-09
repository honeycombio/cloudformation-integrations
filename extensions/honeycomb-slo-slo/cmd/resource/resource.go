package resource

import (
	"context"
	"fmt"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	c, err := newHnyClient(&req)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	data := &hnyclient.SLO{
		Name:             *currentModel.Name,
		TimePeriodDays:   *currentModel.TimePeriod,
		TargetPerMillion: floatToTPM(*currentModel.TargetPercentage),
		SLI:              hnyclient.SLIRef{Alias: *currentModel.SLI},
	}
	if currentModel.Description != nil {
		data.Description = *currentModel.Description
	}

	mySLO, err := c.SLOs.Create(context.Background(), *currentModel.Dataset, data)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("unable to create SLO: %v", err)
	}

	currentModel.ID = &mySLO.ID
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	c, err := newHnyClient(&req)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	slo, err := c.SLOs.Get(context.Background(), *currentModel.Dataset, *currentModel.ID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("unable to fetch SLO: %v", err)
	}

	currentModel.Name = &slo.Name
	currentModel.Description = &slo.Description
	currentModel.SLI = &slo.SLI.Alias
	currentModel.TargetPercentage = hnyclient.ToPtr(tpmToFloat(slo.TargetPerMillion))
	currentModel.TimePeriod = &slo.TimePeriodDays

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	c, err := newHnyClient(&req)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
	slo := &hnyclient.SLO{
		ID:               *currentModel.ID,
		Name:             *currentModel.Name,
		TargetPerMillion: floatToTPM(*currentModel.TargetPercentage),
		TimePeriodDays:   *currentModel.TimePeriod,
		SLI:              hnyclient.SLIRef{Alias: *currentModel.SLI},
	}
	if currentModel.Description != nil {
		slo.Description = *currentModel.Description
	}

	_, err = c.SLOs.Update(context.Background(), *currentModel.Dataset, slo)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("unable to update SLO: %v", err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	c, err := newHnyClient(&req)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	err = c.SLOs.Delete(context.Background(), *currentModel.Dataset, *currentModel.ID)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("unable to delete SLO: %v", err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModel:   currentModel,
	}, nil
}

type ResourceTypeConfig struct {
	HoneycombConfig struct {
		APIKey string `json:"ApiKey"`
		APIUrl string `json:"ApiUrl"`
	} `json:"HoneycombConfig"`
}

func newHnyClient(req *handler.Request) (*hnyclient.Client, error) {
	var tc ResourceTypeConfig
	err := req.UnmarshalTypeConfig(&tc)
	if err != nil {
		return nil, fmt.Errorf("unable to parse resource type configuration: %v", err)
	}

	c, err := hnyclient.NewClient(&hnyclient.Config{APIKey: tc.HoneycombConfig.APIKey, APIUrl: tc.HoneycombConfig.APIUrl})
	if err != nil {
		return nil, fmt.Errorf("unable to initilize client: %v", err)
	}

	return c, nil
}

// converts a floating point percentage to a 'Target Per Million' SLO value
func floatToTPM(f float64) int {
	return int(f * 10000)
}

// converts a SLO 'Target Per Million' value to a floating point percentage
func tpmToFloat(t int) float64 {
	return float64(t) / 10000
}
