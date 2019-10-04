// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateCampaignInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the campaign.
	//
	// CampaignArn is a required field
	CampaignArn *string `locationName:"campaignArn" type:"string" required:"true"`

	// Specifies the requested minimum provisioned transactions (recommendations)
	// per second that Amazon Personalize will support.
	MinProvisionedTPS *int64 `locationName:"minProvisionedTPS" min:"1" type:"integer"`

	// The ARN of a new solution version to deploy.
	SolutionVersionArn *string `locationName:"solutionVersionArn" type:"string"`
}

// String returns the string representation
func (s UpdateCampaignInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCampaignInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCampaignInput"}

	if s.CampaignArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignArn"))
	}
	if s.MinProvisionedTPS != nil && *s.MinProvisionedTPS < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MinProvisionedTPS", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateCampaignOutput struct {
	_ struct{} `type:"structure"`

	// The same campaign ARN as given in the request.
	CampaignArn *string `locationName:"campaignArn" type:"string"`
}

// String returns the string representation
func (s UpdateCampaignOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateCampaign = "UpdateCampaign"

// UpdateCampaignRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Updates a campaign by either deploying a new solution or changing the value
// of the campaign's minProvisionedTPS parameter.
//
// To update a campaign, the campaign status must be ACTIVE or CREATE FAILED.
// Check the campaign status using the DescribeCampaign API.
//
// You must wait until the status of the updated campaign is ACTIVE before asking
// the campaign for recommendations.
//
// For more information on campaigns, see CreateCampaign.
//
//    // Example sending a request using UpdateCampaignRequest.
//    req := client.UpdateCampaignRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/UpdateCampaign
func (c *Client) UpdateCampaignRequest(input *UpdateCampaignInput) UpdateCampaignRequest {
	op := &aws.Operation{
		Name:       opUpdateCampaign,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateCampaignInput{}
	}

	req := c.newRequest(op, input, &UpdateCampaignOutput{})
	return UpdateCampaignRequest{Request: req, Input: input, Copy: c.UpdateCampaignRequest}
}

// UpdateCampaignRequest is the request type for the
// UpdateCampaign API operation.
type UpdateCampaignRequest struct {
	*aws.Request
	Input *UpdateCampaignInput
	Copy  func(*UpdateCampaignInput) UpdateCampaignRequest
}

// Send marshals and sends the UpdateCampaign API request.
func (r UpdateCampaignRequest) Send(ctx context.Context) (*UpdateCampaignResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateCampaignResponse{
		UpdateCampaignOutput: r.Request.Data.(*UpdateCampaignOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateCampaignResponse is the response type for the
// UpdateCampaign API operation.
type UpdateCampaignResponse struct {
	*UpdateCampaignOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCampaign request.
func (r *UpdateCampaignResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}