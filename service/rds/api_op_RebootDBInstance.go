// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You might need to reboot your DB instance, usually for maintenance reasons. For
// example, if you make certain modifications, or if you change the DB parameter
// group associated with the DB instance, you must reboot the instance for the
// changes to take effect. Rebooting a DB instance restarts the database engine
// service. Rebooting a DB instance results in a momentary outage, during which the
// DB instance status is set to rebooting. For more information about rebooting,
// see Rebooting a DB Instance
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_RebootInstance.html)
// in the Amazon RDS User Guide. This command doesn't apply to RDS Custom. If your
// DB instance is part of a Multi-AZ DB cluster, you can reboot the DB cluster with
// the RebootDBCluster operation.
func (c *Client) RebootDBInstance(ctx context.Context, params *RebootDBInstanceInput, optFns ...func(*Options)) (*RebootDBInstanceOutput, error) {
	if params == nil {
		params = &RebootDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootDBInstance", params, optFns, c.addOperationRebootDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RebootDBInstanceInput struct {

	// The DB instance identifier. This parameter is stored as a lowercase string.
	// Constraints:
	//
	// * Must match the identifier of an existing DBInstance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// A value that indicates whether the reboot is conducted through a Multi-AZ
	// failover. Constraint: You can't enable force failover if the instance isn't
	// configured for Multi-AZ.
	ForceFailover *bool

	noSmithyDocumentSerde
}

type RebootDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the operations CreateDBInstance,
	// CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances,
	// ModifyDBInstance, PromoteReadReplica, RebootDBInstance,
	// RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3,
	// RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRebootDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRebootDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRebootDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRebootDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRebootDBInstance(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opRebootDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RebootDBInstance",
	}
}
