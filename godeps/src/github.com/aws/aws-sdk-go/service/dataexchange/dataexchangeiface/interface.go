// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package dataexchangeiface provides an interface to enable mocking the AWS Data Exchange service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package dataexchangeiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dataexchange"
)

// DataExchangeAPI provides an interface to enable mocking the
// dataexchange.DataExchange service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS Data Exchange.
//	func myFunc(svc dataexchangeiface.DataExchangeAPI) bool {
//	    // Make svc.CancelJob request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := dataexchange.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockDataExchangeClient struct {
//	    dataexchangeiface.DataExchangeAPI
//	}
//	func (m *mockDataExchangeClient) CancelJob(input *dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockDataExchangeClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DataExchangeAPI interface {
	CancelJob(*dataexchange.CancelJobInput) (*dataexchange.CancelJobOutput, error)
	CancelJobWithContext(aws.Context, *dataexchange.CancelJobInput, ...request.Option) (*dataexchange.CancelJobOutput, error)
	CancelJobRequest(*dataexchange.CancelJobInput) (*request.Request, *dataexchange.CancelJobOutput)

	CreateDataSet(*dataexchange.CreateDataSetInput) (*dataexchange.CreateDataSetOutput, error)
	CreateDataSetWithContext(aws.Context, *dataexchange.CreateDataSetInput, ...request.Option) (*dataexchange.CreateDataSetOutput, error)
	CreateDataSetRequest(*dataexchange.CreateDataSetInput) (*request.Request, *dataexchange.CreateDataSetOutput)

	CreateEventAction(*dataexchange.CreateEventActionInput) (*dataexchange.CreateEventActionOutput, error)
	CreateEventActionWithContext(aws.Context, *dataexchange.CreateEventActionInput, ...request.Option) (*dataexchange.CreateEventActionOutput, error)
	CreateEventActionRequest(*dataexchange.CreateEventActionInput) (*request.Request, *dataexchange.CreateEventActionOutput)

	CreateJob(*dataexchange.CreateJobInput) (*dataexchange.CreateJobOutput, error)
	CreateJobWithContext(aws.Context, *dataexchange.CreateJobInput, ...request.Option) (*dataexchange.CreateJobOutput, error)
	CreateJobRequest(*dataexchange.CreateJobInput) (*request.Request, *dataexchange.CreateJobOutput)

	CreateRevision(*dataexchange.CreateRevisionInput) (*dataexchange.CreateRevisionOutput, error)
	CreateRevisionWithContext(aws.Context, *dataexchange.CreateRevisionInput, ...request.Option) (*dataexchange.CreateRevisionOutput, error)
	CreateRevisionRequest(*dataexchange.CreateRevisionInput) (*request.Request, *dataexchange.CreateRevisionOutput)

	DeleteAsset(*dataexchange.DeleteAssetInput) (*dataexchange.DeleteAssetOutput, error)
	DeleteAssetWithContext(aws.Context, *dataexchange.DeleteAssetInput, ...request.Option) (*dataexchange.DeleteAssetOutput, error)
	DeleteAssetRequest(*dataexchange.DeleteAssetInput) (*request.Request, *dataexchange.DeleteAssetOutput)

	DeleteDataSet(*dataexchange.DeleteDataSetInput) (*dataexchange.DeleteDataSetOutput, error)
	DeleteDataSetWithContext(aws.Context, *dataexchange.DeleteDataSetInput, ...request.Option) (*dataexchange.DeleteDataSetOutput, error)
	DeleteDataSetRequest(*dataexchange.DeleteDataSetInput) (*request.Request, *dataexchange.DeleteDataSetOutput)

	DeleteEventAction(*dataexchange.DeleteEventActionInput) (*dataexchange.DeleteEventActionOutput, error)
	DeleteEventActionWithContext(aws.Context, *dataexchange.DeleteEventActionInput, ...request.Option) (*dataexchange.DeleteEventActionOutput, error)
	DeleteEventActionRequest(*dataexchange.DeleteEventActionInput) (*request.Request, *dataexchange.DeleteEventActionOutput)

	DeleteRevision(*dataexchange.DeleteRevisionInput) (*dataexchange.DeleteRevisionOutput, error)
	DeleteRevisionWithContext(aws.Context, *dataexchange.DeleteRevisionInput, ...request.Option) (*dataexchange.DeleteRevisionOutput, error)
	DeleteRevisionRequest(*dataexchange.DeleteRevisionInput) (*request.Request, *dataexchange.DeleteRevisionOutput)

	GetAsset(*dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error)
	GetAssetWithContext(aws.Context, *dataexchange.GetAssetInput, ...request.Option) (*dataexchange.GetAssetOutput, error)
	GetAssetRequest(*dataexchange.GetAssetInput) (*request.Request, *dataexchange.GetAssetOutput)

	GetDataSet(*dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error)
	GetDataSetWithContext(aws.Context, *dataexchange.GetDataSetInput, ...request.Option) (*dataexchange.GetDataSetOutput, error)
	GetDataSetRequest(*dataexchange.GetDataSetInput) (*request.Request, *dataexchange.GetDataSetOutput)

	GetEventAction(*dataexchange.GetEventActionInput) (*dataexchange.GetEventActionOutput, error)
	GetEventActionWithContext(aws.Context, *dataexchange.GetEventActionInput, ...request.Option) (*dataexchange.GetEventActionOutput, error)
	GetEventActionRequest(*dataexchange.GetEventActionInput) (*request.Request, *dataexchange.GetEventActionOutput)

	GetJob(*dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error)
	GetJobWithContext(aws.Context, *dataexchange.GetJobInput, ...request.Option) (*dataexchange.GetJobOutput, error)
	GetJobRequest(*dataexchange.GetJobInput) (*request.Request, *dataexchange.GetJobOutput)

	GetRevision(*dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error)
	GetRevisionWithContext(aws.Context, *dataexchange.GetRevisionInput, ...request.Option) (*dataexchange.GetRevisionOutput, error)
	GetRevisionRequest(*dataexchange.GetRevisionInput) (*request.Request, *dataexchange.GetRevisionOutput)

	ListDataSetRevisions(*dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error)
	ListDataSetRevisionsWithContext(aws.Context, *dataexchange.ListDataSetRevisionsInput, ...request.Option) (*dataexchange.ListDataSetRevisionsOutput, error)
	ListDataSetRevisionsRequest(*dataexchange.ListDataSetRevisionsInput) (*request.Request, *dataexchange.ListDataSetRevisionsOutput)

	ListDataSetRevisionsPages(*dataexchange.ListDataSetRevisionsInput, func(*dataexchange.ListDataSetRevisionsOutput, bool) bool) error
	ListDataSetRevisionsPagesWithContext(aws.Context, *dataexchange.ListDataSetRevisionsInput, func(*dataexchange.ListDataSetRevisionsOutput, bool) bool, ...request.Option) error

	ListDataSets(*dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error)
	ListDataSetsWithContext(aws.Context, *dataexchange.ListDataSetsInput, ...request.Option) (*dataexchange.ListDataSetsOutput, error)
	ListDataSetsRequest(*dataexchange.ListDataSetsInput) (*request.Request, *dataexchange.ListDataSetsOutput)

	ListDataSetsPages(*dataexchange.ListDataSetsInput, func(*dataexchange.ListDataSetsOutput, bool) bool) error
	ListDataSetsPagesWithContext(aws.Context, *dataexchange.ListDataSetsInput, func(*dataexchange.ListDataSetsOutput, bool) bool, ...request.Option) error

	ListEventActions(*dataexchange.ListEventActionsInput) (*dataexchange.ListEventActionsOutput, error)
	ListEventActionsWithContext(aws.Context, *dataexchange.ListEventActionsInput, ...request.Option) (*dataexchange.ListEventActionsOutput, error)
	ListEventActionsRequest(*dataexchange.ListEventActionsInput) (*request.Request, *dataexchange.ListEventActionsOutput)

	ListEventActionsPages(*dataexchange.ListEventActionsInput, func(*dataexchange.ListEventActionsOutput, bool) bool) error
	ListEventActionsPagesWithContext(aws.Context, *dataexchange.ListEventActionsInput, func(*dataexchange.ListEventActionsOutput, bool) bool, ...request.Option) error

	ListJobs(*dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *dataexchange.ListJobsInput, ...request.Option) (*dataexchange.ListJobsOutput, error)
	ListJobsRequest(*dataexchange.ListJobsInput) (*request.Request, *dataexchange.ListJobsOutput)

	ListJobsPages(*dataexchange.ListJobsInput, func(*dataexchange.ListJobsOutput, bool) bool) error
	ListJobsPagesWithContext(aws.Context, *dataexchange.ListJobsInput, func(*dataexchange.ListJobsOutput, bool) bool, ...request.Option) error

	ListRevisionAssets(*dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error)
	ListRevisionAssetsWithContext(aws.Context, *dataexchange.ListRevisionAssetsInput, ...request.Option) (*dataexchange.ListRevisionAssetsOutput, error)
	ListRevisionAssetsRequest(*dataexchange.ListRevisionAssetsInput) (*request.Request, *dataexchange.ListRevisionAssetsOutput)

	ListRevisionAssetsPages(*dataexchange.ListRevisionAssetsInput, func(*dataexchange.ListRevisionAssetsOutput, bool) bool) error
	ListRevisionAssetsPagesWithContext(aws.Context, *dataexchange.ListRevisionAssetsInput, func(*dataexchange.ListRevisionAssetsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *dataexchange.ListTagsForResourceInput, ...request.Option) (*dataexchange.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*dataexchange.ListTagsForResourceInput) (*request.Request, *dataexchange.ListTagsForResourceOutput)

	RevokeRevision(*dataexchange.RevokeRevisionInput) (*dataexchange.RevokeRevisionOutput, error)
	RevokeRevisionWithContext(aws.Context, *dataexchange.RevokeRevisionInput, ...request.Option) (*dataexchange.RevokeRevisionOutput, error)
	RevokeRevisionRequest(*dataexchange.RevokeRevisionInput) (*request.Request, *dataexchange.RevokeRevisionOutput)

	SendApiAsset(*dataexchange.SendApiAssetInput) (*dataexchange.SendApiAssetOutput, error)
	SendApiAssetWithContext(aws.Context, *dataexchange.SendApiAssetInput, ...request.Option) (*dataexchange.SendApiAssetOutput, error)
	SendApiAssetRequest(*dataexchange.SendApiAssetInput) (*request.Request, *dataexchange.SendApiAssetOutput)

	SendDataSetNotification(*dataexchange.SendDataSetNotificationInput) (*dataexchange.SendDataSetNotificationOutput, error)
	SendDataSetNotificationWithContext(aws.Context, *dataexchange.SendDataSetNotificationInput, ...request.Option) (*dataexchange.SendDataSetNotificationOutput, error)
	SendDataSetNotificationRequest(*dataexchange.SendDataSetNotificationInput) (*request.Request, *dataexchange.SendDataSetNotificationOutput)

	StartJob(*dataexchange.StartJobInput) (*dataexchange.StartJobOutput, error)
	StartJobWithContext(aws.Context, *dataexchange.StartJobInput, ...request.Option) (*dataexchange.StartJobOutput, error)
	StartJobRequest(*dataexchange.StartJobInput) (*request.Request, *dataexchange.StartJobOutput)

	TagResource(*dataexchange.TagResourceInput) (*dataexchange.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *dataexchange.TagResourceInput, ...request.Option) (*dataexchange.TagResourceOutput, error)
	TagResourceRequest(*dataexchange.TagResourceInput) (*request.Request, *dataexchange.TagResourceOutput)

	UntagResource(*dataexchange.UntagResourceInput) (*dataexchange.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *dataexchange.UntagResourceInput, ...request.Option) (*dataexchange.UntagResourceOutput, error)
	UntagResourceRequest(*dataexchange.UntagResourceInput) (*request.Request, *dataexchange.UntagResourceOutput)

	UpdateAsset(*dataexchange.UpdateAssetInput) (*dataexchange.UpdateAssetOutput, error)
	UpdateAssetWithContext(aws.Context, *dataexchange.UpdateAssetInput, ...request.Option) (*dataexchange.UpdateAssetOutput, error)
	UpdateAssetRequest(*dataexchange.UpdateAssetInput) (*request.Request, *dataexchange.UpdateAssetOutput)

	UpdateDataSet(*dataexchange.UpdateDataSetInput) (*dataexchange.UpdateDataSetOutput, error)
	UpdateDataSetWithContext(aws.Context, *dataexchange.UpdateDataSetInput, ...request.Option) (*dataexchange.UpdateDataSetOutput, error)
	UpdateDataSetRequest(*dataexchange.UpdateDataSetInput) (*request.Request, *dataexchange.UpdateDataSetOutput)

	UpdateEventAction(*dataexchange.UpdateEventActionInput) (*dataexchange.UpdateEventActionOutput, error)
	UpdateEventActionWithContext(aws.Context, *dataexchange.UpdateEventActionInput, ...request.Option) (*dataexchange.UpdateEventActionOutput, error)
	UpdateEventActionRequest(*dataexchange.UpdateEventActionInput) (*request.Request, *dataexchange.UpdateEventActionOutput)

	UpdateRevision(*dataexchange.UpdateRevisionInput) (*dataexchange.UpdateRevisionOutput, error)
	UpdateRevisionWithContext(aws.Context, *dataexchange.UpdateRevisionInput, ...request.Option) (*dataexchange.UpdateRevisionOutput, error)
	UpdateRevisionRequest(*dataexchange.UpdateRevisionInput) (*request.Request, *dataexchange.UpdateRevisionOutput)
}

var _ DataExchangeAPI = (*dataexchange.DataExchange)(nil)
