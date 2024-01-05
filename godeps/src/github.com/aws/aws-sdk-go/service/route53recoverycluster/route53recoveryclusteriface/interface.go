// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package route53recoveryclusteriface provides an interface to enable mocking the Route53 Recovery Cluster service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package route53recoveryclusteriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53recoverycluster"
)

// Route53RecoveryClusterAPI provides an interface to enable mocking the
// route53recoverycluster.Route53RecoveryCluster service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Route53 Recovery Cluster.
//	func myFunc(svc route53recoveryclusteriface.Route53RecoveryClusterAPI) bool {
//	    // Make svc.GetRoutingControlState request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := route53recoverycluster.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockRoute53RecoveryClusterClient struct {
//	    route53recoveryclusteriface.Route53RecoveryClusterAPI
//	}
//	func (m *mockRoute53RecoveryClusterClient) GetRoutingControlState(input *route53recoverycluster.GetRoutingControlStateInput) (*route53recoverycluster.GetRoutingControlStateOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockRoute53RecoveryClusterClient{}
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
type Route53RecoveryClusterAPI interface {
	GetRoutingControlState(*route53recoverycluster.GetRoutingControlStateInput) (*route53recoverycluster.GetRoutingControlStateOutput, error)
	GetRoutingControlStateWithContext(aws.Context, *route53recoverycluster.GetRoutingControlStateInput, ...request.Option) (*route53recoverycluster.GetRoutingControlStateOutput, error)
	GetRoutingControlStateRequest(*route53recoverycluster.GetRoutingControlStateInput) (*request.Request, *route53recoverycluster.GetRoutingControlStateOutput)

	ListRoutingControls(*route53recoverycluster.ListRoutingControlsInput) (*route53recoverycluster.ListRoutingControlsOutput, error)
	ListRoutingControlsWithContext(aws.Context, *route53recoverycluster.ListRoutingControlsInput, ...request.Option) (*route53recoverycluster.ListRoutingControlsOutput, error)
	ListRoutingControlsRequest(*route53recoverycluster.ListRoutingControlsInput) (*request.Request, *route53recoverycluster.ListRoutingControlsOutput)

	ListRoutingControlsPages(*route53recoverycluster.ListRoutingControlsInput, func(*route53recoverycluster.ListRoutingControlsOutput, bool) bool) error
	ListRoutingControlsPagesWithContext(aws.Context, *route53recoverycluster.ListRoutingControlsInput, func(*route53recoverycluster.ListRoutingControlsOutput, bool) bool, ...request.Option) error

	UpdateRoutingControlState(*route53recoverycluster.UpdateRoutingControlStateInput) (*route53recoverycluster.UpdateRoutingControlStateOutput, error)
	UpdateRoutingControlStateWithContext(aws.Context, *route53recoverycluster.UpdateRoutingControlStateInput, ...request.Option) (*route53recoverycluster.UpdateRoutingControlStateOutput, error)
	UpdateRoutingControlStateRequest(*route53recoverycluster.UpdateRoutingControlStateInput) (*request.Request, *route53recoverycluster.UpdateRoutingControlStateOutput)

	UpdateRoutingControlStates(*route53recoverycluster.UpdateRoutingControlStatesInput) (*route53recoverycluster.UpdateRoutingControlStatesOutput, error)
	UpdateRoutingControlStatesWithContext(aws.Context, *route53recoverycluster.UpdateRoutingControlStatesInput, ...request.Option) (*route53recoverycluster.UpdateRoutingControlStatesOutput, error)
	UpdateRoutingControlStatesRequest(*route53recoverycluster.UpdateRoutingControlStatesInput) (*request.Request, *route53recoverycluster.UpdateRoutingControlStatesOutput)
}

var _ Route53RecoveryClusterAPI = (*route53recoverycluster.Route53RecoveryCluster)(nil)
