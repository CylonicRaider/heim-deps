// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package finspaceiface provides an interface to enable mocking the FinSpace User Environment Management service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package finspaceiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/finspace"
)

// FinspaceAPI provides an interface to enable mocking the
// finspace.Finspace service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// FinSpace User Environment Management service.
//	func myFunc(svc finspaceiface.FinspaceAPI) bool {
//	    // Make svc.CreateEnvironment request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := finspace.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockFinspaceClient struct {
//	    finspaceiface.FinspaceAPI
//	}
//	func (m *mockFinspaceClient) CreateEnvironment(input *finspace.CreateEnvironmentInput) (*finspace.CreateEnvironmentOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockFinspaceClient{}
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
type FinspaceAPI interface {
	CreateEnvironment(*finspace.CreateEnvironmentInput) (*finspace.CreateEnvironmentOutput, error)
	CreateEnvironmentWithContext(aws.Context, *finspace.CreateEnvironmentInput, ...request.Option) (*finspace.CreateEnvironmentOutput, error)
	CreateEnvironmentRequest(*finspace.CreateEnvironmentInput) (*request.Request, *finspace.CreateEnvironmentOutput)

	CreateKxChangeset(*finspace.CreateKxChangesetInput) (*finspace.CreateKxChangesetOutput, error)
	CreateKxChangesetWithContext(aws.Context, *finspace.CreateKxChangesetInput, ...request.Option) (*finspace.CreateKxChangesetOutput, error)
	CreateKxChangesetRequest(*finspace.CreateKxChangesetInput) (*request.Request, *finspace.CreateKxChangesetOutput)

	CreateKxCluster(*finspace.CreateKxClusterInput) (*finspace.CreateKxClusterOutput, error)
	CreateKxClusterWithContext(aws.Context, *finspace.CreateKxClusterInput, ...request.Option) (*finspace.CreateKxClusterOutput, error)
	CreateKxClusterRequest(*finspace.CreateKxClusterInput) (*request.Request, *finspace.CreateKxClusterOutput)

	CreateKxDatabase(*finspace.CreateKxDatabaseInput) (*finspace.CreateKxDatabaseOutput, error)
	CreateKxDatabaseWithContext(aws.Context, *finspace.CreateKxDatabaseInput, ...request.Option) (*finspace.CreateKxDatabaseOutput, error)
	CreateKxDatabaseRequest(*finspace.CreateKxDatabaseInput) (*request.Request, *finspace.CreateKxDatabaseOutput)

	CreateKxDataview(*finspace.CreateKxDataviewInput) (*finspace.CreateKxDataviewOutput, error)
	CreateKxDataviewWithContext(aws.Context, *finspace.CreateKxDataviewInput, ...request.Option) (*finspace.CreateKxDataviewOutput, error)
	CreateKxDataviewRequest(*finspace.CreateKxDataviewInput) (*request.Request, *finspace.CreateKxDataviewOutput)

	CreateKxEnvironment(*finspace.CreateKxEnvironmentInput) (*finspace.CreateKxEnvironmentOutput, error)
	CreateKxEnvironmentWithContext(aws.Context, *finspace.CreateKxEnvironmentInput, ...request.Option) (*finspace.CreateKxEnvironmentOutput, error)
	CreateKxEnvironmentRequest(*finspace.CreateKxEnvironmentInput) (*request.Request, *finspace.CreateKxEnvironmentOutput)

	CreateKxScalingGroup(*finspace.CreateKxScalingGroupInput) (*finspace.CreateKxScalingGroupOutput, error)
	CreateKxScalingGroupWithContext(aws.Context, *finspace.CreateKxScalingGroupInput, ...request.Option) (*finspace.CreateKxScalingGroupOutput, error)
	CreateKxScalingGroupRequest(*finspace.CreateKxScalingGroupInput) (*request.Request, *finspace.CreateKxScalingGroupOutput)

	CreateKxUser(*finspace.CreateKxUserInput) (*finspace.CreateKxUserOutput, error)
	CreateKxUserWithContext(aws.Context, *finspace.CreateKxUserInput, ...request.Option) (*finspace.CreateKxUserOutput, error)
	CreateKxUserRequest(*finspace.CreateKxUserInput) (*request.Request, *finspace.CreateKxUserOutput)

	CreateKxVolume(*finspace.CreateKxVolumeInput) (*finspace.CreateKxVolumeOutput, error)
	CreateKxVolumeWithContext(aws.Context, *finspace.CreateKxVolumeInput, ...request.Option) (*finspace.CreateKxVolumeOutput, error)
	CreateKxVolumeRequest(*finspace.CreateKxVolumeInput) (*request.Request, *finspace.CreateKxVolumeOutput)

	DeleteEnvironment(*finspace.DeleteEnvironmentInput) (*finspace.DeleteEnvironmentOutput, error)
	DeleteEnvironmentWithContext(aws.Context, *finspace.DeleteEnvironmentInput, ...request.Option) (*finspace.DeleteEnvironmentOutput, error)
	DeleteEnvironmentRequest(*finspace.DeleteEnvironmentInput) (*request.Request, *finspace.DeleteEnvironmentOutput)

	DeleteKxCluster(*finspace.DeleteKxClusterInput) (*finspace.DeleteKxClusterOutput, error)
	DeleteKxClusterWithContext(aws.Context, *finspace.DeleteKxClusterInput, ...request.Option) (*finspace.DeleteKxClusterOutput, error)
	DeleteKxClusterRequest(*finspace.DeleteKxClusterInput) (*request.Request, *finspace.DeleteKxClusterOutput)

	DeleteKxDatabase(*finspace.DeleteKxDatabaseInput) (*finspace.DeleteKxDatabaseOutput, error)
	DeleteKxDatabaseWithContext(aws.Context, *finspace.DeleteKxDatabaseInput, ...request.Option) (*finspace.DeleteKxDatabaseOutput, error)
	DeleteKxDatabaseRequest(*finspace.DeleteKxDatabaseInput) (*request.Request, *finspace.DeleteKxDatabaseOutput)

	DeleteKxDataview(*finspace.DeleteKxDataviewInput) (*finspace.DeleteKxDataviewOutput, error)
	DeleteKxDataviewWithContext(aws.Context, *finspace.DeleteKxDataviewInput, ...request.Option) (*finspace.DeleteKxDataviewOutput, error)
	DeleteKxDataviewRequest(*finspace.DeleteKxDataviewInput) (*request.Request, *finspace.DeleteKxDataviewOutput)

	DeleteKxEnvironment(*finspace.DeleteKxEnvironmentInput) (*finspace.DeleteKxEnvironmentOutput, error)
	DeleteKxEnvironmentWithContext(aws.Context, *finspace.DeleteKxEnvironmentInput, ...request.Option) (*finspace.DeleteKxEnvironmentOutput, error)
	DeleteKxEnvironmentRequest(*finspace.DeleteKxEnvironmentInput) (*request.Request, *finspace.DeleteKxEnvironmentOutput)

	DeleteKxScalingGroup(*finspace.DeleteKxScalingGroupInput) (*finspace.DeleteKxScalingGroupOutput, error)
	DeleteKxScalingGroupWithContext(aws.Context, *finspace.DeleteKxScalingGroupInput, ...request.Option) (*finspace.DeleteKxScalingGroupOutput, error)
	DeleteKxScalingGroupRequest(*finspace.DeleteKxScalingGroupInput) (*request.Request, *finspace.DeleteKxScalingGroupOutput)

	DeleteKxUser(*finspace.DeleteKxUserInput) (*finspace.DeleteKxUserOutput, error)
	DeleteKxUserWithContext(aws.Context, *finspace.DeleteKxUserInput, ...request.Option) (*finspace.DeleteKxUserOutput, error)
	DeleteKxUserRequest(*finspace.DeleteKxUserInput) (*request.Request, *finspace.DeleteKxUserOutput)

	DeleteKxVolume(*finspace.DeleteKxVolumeInput) (*finspace.DeleteKxVolumeOutput, error)
	DeleteKxVolumeWithContext(aws.Context, *finspace.DeleteKxVolumeInput, ...request.Option) (*finspace.DeleteKxVolumeOutput, error)
	DeleteKxVolumeRequest(*finspace.DeleteKxVolumeInput) (*request.Request, *finspace.DeleteKxVolumeOutput)

	GetEnvironment(*finspace.GetEnvironmentInput) (*finspace.GetEnvironmentOutput, error)
	GetEnvironmentWithContext(aws.Context, *finspace.GetEnvironmentInput, ...request.Option) (*finspace.GetEnvironmentOutput, error)
	GetEnvironmentRequest(*finspace.GetEnvironmentInput) (*request.Request, *finspace.GetEnvironmentOutput)

	GetKxChangeset(*finspace.GetKxChangesetInput) (*finspace.GetKxChangesetOutput, error)
	GetKxChangesetWithContext(aws.Context, *finspace.GetKxChangesetInput, ...request.Option) (*finspace.GetKxChangesetOutput, error)
	GetKxChangesetRequest(*finspace.GetKxChangesetInput) (*request.Request, *finspace.GetKxChangesetOutput)

	GetKxCluster(*finspace.GetKxClusterInput) (*finspace.GetKxClusterOutput, error)
	GetKxClusterWithContext(aws.Context, *finspace.GetKxClusterInput, ...request.Option) (*finspace.GetKxClusterOutput, error)
	GetKxClusterRequest(*finspace.GetKxClusterInput) (*request.Request, *finspace.GetKxClusterOutput)

	GetKxConnectionString(*finspace.GetKxConnectionStringInput) (*finspace.GetKxConnectionStringOutput, error)
	GetKxConnectionStringWithContext(aws.Context, *finspace.GetKxConnectionStringInput, ...request.Option) (*finspace.GetKxConnectionStringOutput, error)
	GetKxConnectionStringRequest(*finspace.GetKxConnectionStringInput) (*request.Request, *finspace.GetKxConnectionStringOutput)

	GetKxDatabase(*finspace.GetKxDatabaseInput) (*finspace.GetKxDatabaseOutput, error)
	GetKxDatabaseWithContext(aws.Context, *finspace.GetKxDatabaseInput, ...request.Option) (*finspace.GetKxDatabaseOutput, error)
	GetKxDatabaseRequest(*finspace.GetKxDatabaseInput) (*request.Request, *finspace.GetKxDatabaseOutput)

	GetKxDataview(*finspace.GetKxDataviewInput) (*finspace.GetKxDataviewOutput, error)
	GetKxDataviewWithContext(aws.Context, *finspace.GetKxDataviewInput, ...request.Option) (*finspace.GetKxDataviewOutput, error)
	GetKxDataviewRequest(*finspace.GetKxDataviewInput) (*request.Request, *finspace.GetKxDataviewOutput)

	GetKxEnvironment(*finspace.GetKxEnvironmentInput) (*finspace.GetKxEnvironmentOutput, error)
	GetKxEnvironmentWithContext(aws.Context, *finspace.GetKxEnvironmentInput, ...request.Option) (*finspace.GetKxEnvironmentOutput, error)
	GetKxEnvironmentRequest(*finspace.GetKxEnvironmentInput) (*request.Request, *finspace.GetKxEnvironmentOutput)

	GetKxScalingGroup(*finspace.GetKxScalingGroupInput) (*finspace.GetKxScalingGroupOutput, error)
	GetKxScalingGroupWithContext(aws.Context, *finspace.GetKxScalingGroupInput, ...request.Option) (*finspace.GetKxScalingGroupOutput, error)
	GetKxScalingGroupRequest(*finspace.GetKxScalingGroupInput) (*request.Request, *finspace.GetKxScalingGroupOutput)

	GetKxUser(*finspace.GetKxUserInput) (*finspace.GetKxUserOutput, error)
	GetKxUserWithContext(aws.Context, *finspace.GetKxUserInput, ...request.Option) (*finspace.GetKxUserOutput, error)
	GetKxUserRequest(*finspace.GetKxUserInput) (*request.Request, *finspace.GetKxUserOutput)

	GetKxVolume(*finspace.GetKxVolumeInput) (*finspace.GetKxVolumeOutput, error)
	GetKxVolumeWithContext(aws.Context, *finspace.GetKxVolumeInput, ...request.Option) (*finspace.GetKxVolumeOutput, error)
	GetKxVolumeRequest(*finspace.GetKxVolumeInput) (*request.Request, *finspace.GetKxVolumeOutput)

	ListEnvironments(*finspace.ListEnvironmentsInput) (*finspace.ListEnvironmentsOutput, error)
	ListEnvironmentsWithContext(aws.Context, *finspace.ListEnvironmentsInput, ...request.Option) (*finspace.ListEnvironmentsOutput, error)
	ListEnvironmentsRequest(*finspace.ListEnvironmentsInput) (*request.Request, *finspace.ListEnvironmentsOutput)

	ListKxChangesets(*finspace.ListKxChangesetsInput) (*finspace.ListKxChangesetsOutput, error)
	ListKxChangesetsWithContext(aws.Context, *finspace.ListKxChangesetsInput, ...request.Option) (*finspace.ListKxChangesetsOutput, error)
	ListKxChangesetsRequest(*finspace.ListKxChangesetsInput) (*request.Request, *finspace.ListKxChangesetsOutput)

	ListKxChangesetsPages(*finspace.ListKxChangesetsInput, func(*finspace.ListKxChangesetsOutput, bool) bool) error
	ListKxChangesetsPagesWithContext(aws.Context, *finspace.ListKxChangesetsInput, func(*finspace.ListKxChangesetsOutput, bool) bool, ...request.Option) error

	ListKxClusterNodes(*finspace.ListKxClusterNodesInput) (*finspace.ListKxClusterNodesOutput, error)
	ListKxClusterNodesWithContext(aws.Context, *finspace.ListKxClusterNodesInput, ...request.Option) (*finspace.ListKxClusterNodesOutput, error)
	ListKxClusterNodesRequest(*finspace.ListKxClusterNodesInput) (*request.Request, *finspace.ListKxClusterNodesOutput)

	ListKxClusterNodesPages(*finspace.ListKxClusterNodesInput, func(*finspace.ListKxClusterNodesOutput, bool) bool) error
	ListKxClusterNodesPagesWithContext(aws.Context, *finspace.ListKxClusterNodesInput, func(*finspace.ListKxClusterNodesOutput, bool) bool, ...request.Option) error

	ListKxClusters(*finspace.ListKxClustersInput) (*finspace.ListKxClustersOutput, error)
	ListKxClustersWithContext(aws.Context, *finspace.ListKxClustersInput, ...request.Option) (*finspace.ListKxClustersOutput, error)
	ListKxClustersRequest(*finspace.ListKxClustersInput) (*request.Request, *finspace.ListKxClustersOutput)

	ListKxDatabases(*finspace.ListKxDatabasesInput) (*finspace.ListKxDatabasesOutput, error)
	ListKxDatabasesWithContext(aws.Context, *finspace.ListKxDatabasesInput, ...request.Option) (*finspace.ListKxDatabasesOutput, error)
	ListKxDatabasesRequest(*finspace.ListKxDatabasesInput) (*request.Request, *finspace.ListKxDatabasesOutput)

	ListKxDatabasesPages(*finspace.ListKxDatabasesInput, func(*finspace.ListKxDatabasesOutput, bool) bool) error
	ListKxDatabasesPagesWithContext(aws.Context, *finspace.ListKxDatabasesInput, func(*finspace.ListKxDatabasesOutput, bool) bool, ...request.Option) error

	ListKxDataviews(*finspace.ListKxDataviewsInput) (*finspace.ListKxDataviewsOutput, error)
	ListKxDataviewsWithContext(aws.Context, *finspace.ListKxDataviewsInput, ...request.Option) (*finspace.ListKxDataviewsOutput, error)
	ListKxDataviewsRequest(*finspace.ListKxDataviewsInput) (*request.Request, *finspace.ListKxDataviewsOutput)

	ListKxDataviewsPages(*finspace.ListKxDataviewsInput, func(*finspace.ListKxDataviewsOutput, bool) bool) error
	ListKxDataviewsPagesWithContext(aws.Context, *finspace.ListKxDataviewsInput, func(*finspace.ListKxDataviewsOutput, bool) bool, ...request.Option) error

	ListKxEnvironments(*finspace.ListKxEnvironmentsInput) (*finspace.ListKxEnvironmentsOutput, error)
	ListKxEnvironmentsWithContext(aws.Context, *finspace.ListKxEnvironmentsInput, ...request.Option) (*finspace.ListKxEnvironmentsOutput, error)
	ListKxEnvironmentsRequest(*finspace.ListKxEnvironmentsInput) (*request.Request, *finspace.ListKxEnvironmentsOutput)

	ListKxEnvironmentsPages(*finspace.ListKxEnvironmentsInput, func(*finspace.ListKxEnvironmentsOutput, bool) bool) error
	ListKxEnvironmentsPagesWithContext(aws.Context, *finspace.ListKxEnvironmentsInput, func(*finspace.ListKxEnvironmentsOutput, bool) bool, ...request.Option) error

	ListKxScalingGroups(*finspace.ListKxScalingGroupsInput) (*finspace.ListKxScalingGroupsOutput, error)
	ListKxScalingGroupsWithContext(aws.Context, *finspace.ListKxScalingGroupsInput, ...request.Option) (*finspace.ListKxScalingGroupsOutput, error)
	ListKxScalingGroupsRequest(*finspace.ListKxScalingGroupsInput) (*request.Request, *finspace.ListKxScalingGroupsOutput)

	ListKxScalingGroupsPages(*finspace.ListKxScalingGroupsInput, func(*finspace.ListKxScalingGroupsOutput, bool) bool) error
	ListKxScalingGroupsPagesWithContext(aws.Context, *finspace.ListKxScalingGroupsInput, func(*finspace.ListKxScalingGroupsOutput, bool) bool, ...request.Option) error

	ListKxUsers(*finspace.ListKxUsersInput) (*finspace.ListKxUsersOutput, error)
	ListKxUsersWithContext(aws.Context, *finspace.ListKxUsersInput, ...request.Option) (*finspace.ListKxUsersOutput, error)
	ListKxUsersRequest(*finspace.ListKxUsersInput) (*request.Request, *finspace.ListKxUsersOutput)

	ListKxVolumes(*finspace.ListKxVolumesInput) (*finspace.ListKxVolumesOutput, error)
	ListKxVolumesWithContext(aws.Context, *finspace.ListKxVolumesInput, ...request.Option) (*finspace.ListKxVolumesOutput, error)
	ListKxVolumesRequest(*finspace.ListKxVolumesInput) (*request.Request, *finspace.ListKxVolumesOutput)

	ListTagsForResource(*finspace.ListTagsForResourceInput) (*finspace.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *finspace.ListTagsForResourceInput, ...request.Option) (*finspace.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*finspace.ListTagsForResourceInput) (*request.Request, *finspace.ListTagsForResourceOutput)

	TagResource(*finspace.TagResourceInput) (*finspace.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *finspace.TagResourceInput, ...request.Option) (*finspace.TagResourceOutput, error)
	TagResourceRequest(*finspace.TagResourceInput) (*request.Request, *finspace.TagResourceOutput)

	UntagResource(*finspace.UntagResourceInput) (*finspace.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *finspace.UntagResourceInput, ...request.Option) (*finspace.UntagResourceOutput, error)
	UntagResourceRequest(*finspace.UntagResourceInput) (*request.Request, *finspace.UntagResourceOutput)

	UpdateEnvironment(*finspace.UpdateEnvironmentInput) (*finspace.UpdateEnvironmentOutput, error)
	UpdateEnvironmentWithContext(aws.Context, *finspace.UpdateEnvironmentInput, ...request.Option) (*finspace.UpdateEnvironmentOutput, error)
	UpdateEnvironmentRequest(*finspace.UpdateEnvironmentInput) (*request.Request, *finspace.UpdateEnvironmentOutput)

	UpdateKxClusterCodeConfiguration(*finspace.UpdateKxClusterCodeConfigurationInput) (*finspace.UpdateKxClusterCodeConfigurationOutput, error)
	UpdateKxClusterCodeConfigurationWithContext(aws.Context, *finspace.UpdateKxClusterCodeConfigurationInput, ...request.Option) (*finspace.UpdateKxClusterCodeConfigurationOutput, error)
	UpdateKxClusterCodeConfigurationRequest(*finspace.UpdateKxClusterCodeConfigurationInput) (*request.Request, *finspace.UpdateKxClusterCodeConfigurationOutput)

	UpdateKxClusterDatabases(*finspace.UpdateKxClusterDatabasesInput) (*finspace.UpdateKxClusterDatabasesOutput, error)
	UpdateKxClusterDatabasesWithContext(aws.Context, *finspace.UpdateKxClusterDatabasesInput, ...request.Option) (*finspace.UpdateKxClusterDatabasesOutput, error)
	UpdateKxClusterDatabasesRequest(*finspace.UpdateKxClusterDatabasesInput) (*request.Request, *finspace.UpdateKxClusterDatabasesOutput)

	UpdateKxDatabase(*finspace.UpdateKxDatabaseInput) (*finspace.UpdateKxDatabaseOutput, error)
	UpdateKxDatabaseWithContext(aws.Context, *finspace.UpdateKxDatabaseInput, ...request.Option) (*finspace.UpdateKxDatabaseOutput, error)
	UpdateKxDatabaseRequest(*finspace.UpdateKxDatabaseInput) (*request.Request, *finspace.UpdateKxDatabaseOutput)

	UpdateKxDataview(*finspace.UpdateKxDataviewInput) (*finspace.UpdateKxDataviewOutput, error)
	UpdateKxDataviewWithContext(aws.Context, *finspace.UpdateKxDataviewInput, ...request.Option) (*finspace.UpdateKxDataviewOutput, error)
	UpdateKxDataviewRequest(*finspace.UpdateKxDataviewInput) (*request.Request, *finspace.UpdateKxDataviewOutput)

	UpdateKxEnvironment(*finspace.UpdateKxEnvironmentInput) (*finspace.UpdateKxEnvironmentOutput, error)
	UpdateKxEnvironmentWithContext(aws.Context, *finspace.UpdateKxEnvironmentInput, ...request.Option) (*finspace.UpdateKxEnvironmentOutput, error)
	UpdateKxEnvironmentRequest(*finspace.UpdateKxEnvironmentInput) (*request.Request, *finspace.UpdateKxEnvironmentOutput)

	UpdateKxEnvironmentNetwork(*finspace.UpdateKxEnvironmentNetworkInput) (*finspace.UpdateKxEnvironmentNetworkOutput, error)
	UpdateKxEnvironmentNetworkWithContext(aws.Context, *finspace.UpdateKxEnvironmentNetworkInput, ...request.Option) (*finspace.UpdateKxEnvironmentNetworkOutput, error)
	UpdateKxEnvironmentNetworkRequest(*finspace.UpdateKxEnvironmentNetworkInput) (*request.Request, *finspace.UpdateKxEnvironmentNetworkOutput)

	UpdateKxUser(*finspace.UpdateKxUserInput) (*finspace.UpdateKxUserOutput, error)
	UpdateKxUserWithContext(aws.Context, *finspace.UpdateKxUserInput, ...request.Option) (*finspace.UpdateKxUserOutput, error)
	UpdateKxUserRequest(*finspace.UpdateKxUserInput) (*request.Request, *finspace.UpdateKxUserOutput)

	UpdateKxVolume(*finspace.UpdateKxVolumeInput) (*finspace.UpdateKxVolumeOutput, error)
	UpdateKxVolumeWithContext(aws.Context, *finspace.UpdateKxVolumeInput, ...request.Option) (*finspace.UpdateKxVolumeOutput, error)
	UpdateKxVolumeRequest(*finspace.UpdateKxVolumeInput) (*request.Request, *finspace.UpdateKxVolumeOutput)
}

var _ FinspaceAPI = (*finspace.Finspace)(nil)
