// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package voiceidiface provides an interface to enable mocking the Amazon Voice ID service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package voiceidiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/voiceid"
)

// VoiceIDAPI provides an interface to enable mocking the
// voiceid.VoiceID service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Voice ID.
//	func myFunc(svc voiceidiface.VoiceIDAPI) bool {
//	    // Make svc.AssociateFraudster request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := voiceid.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockVoiceIDClient struct {
//	    voiceidiface.VoiceIDAPI
//	}
//	func (m *mockVoiceIDClient) AssociateFraudster(input *voiceid.AssociateFraudsterInput) (*voiceid.AssociateFraudsterOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockVoiceIDClient{}
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
type VoiceIDAPI interface {
	AssociateFraudster(*voiceid.AssociateFraudsterInput) (*voiceid.AssociateFraudsterOutput, error)
	AssociateFraudsterWithContext(aws.Context, *voiceid.AssociateFraudsterInput, ...request.Option) (*voiceid.AssociateFraudsterOutput, error)
	AssociateFraudsterRequest(*voiceid.AssociateFraudsterInput) (*request.Request, *voiceid.AssociateFraudsterOutput)

	CreateDomain(*voiceid.CreateDomainInput) (*voiceid.CreateDomainOutput, error)
	CreateDomainWithContext(aws.Context, *voiceid.CreateDomainInput, ...request.Option) (*voiceid.CreateDomainOutput, error)
	CreateDomainRequest(*voiceid.CreateDomainInput) (*request.Request, *voiceid.CreateDomainOutput)

	CreateWatchlist(*voiceid.CreateWatchlistInput) (*voiceid.CreateWatchlistOutput, error)
	CreateWatchlistWithContext(aws.Context, *voiceid.CreateWatchlistInput, ...request.Option) (*voiceid.CreateWatchlistOutput, error)
	CreateWatchlistRequest(*voiceid.CreateWatchlistInput) (*request.Request, *voiceid.CreateWatchlistOutput)

	DeleteDomain(*voiceid.DeleteDomainInput) (*voiceid.DeleteDomainOutput, error)
	DeleteDomainWithContext(aws.Context, *voiceid.DeleteDomainInput, ...request.Option) (*voiceid.DeleteDomainOutput, error)
	DeleteDomainRequest(*voiceid.DeleteDomainInput) (*request.Request, *voiceid.DeleteDomainOutput)

	DeleteFraudster(*voiceid.DeleteFraudsterInput) (*voiceid.DeleteFraudsterOutput, error)
	DeleteFraudsterWithContext(aws.Context, *voiceid.DeleteFraudsterInput, ...request.Option) (*voiceid.DeleteFraudsterOutput, error)
	DeleteFraudsterRequest(*voiceid.DeleteFraudsterInput) (*request.Request, *voiceid.DeleteFraudsterOutput)

	DeleteSpeaker(*voiceid.DeleteSpeakerInput) (*voiceid.DeleteSpeakerOutput, error)
	DeleteSpeakerWithContext(aws.Context, *voiceid.DeleteSpeakerInput, ...request.Option) (*voiceid.DeleteSpeakerOutput, error)
	DeleteSpeakerRequest(*voiceid.DeleteSpeakerInput) (*request.Request, *voiceid.DeleteSpeakerOutput)

	DeleteWatchlist(*voiceid.DeleteWatchlistInput) (*voiceid.DeleteWatchlistOutput, error)
	DeleteWatchlistWithContext(aws.Context, *voiceid.DeleteWatchlistInput, ...request.Option) (*voiceid.DeleteWatchlistOutput, error)
	DeleteWatchlistRequest(*voiceid.DeleteWatchlistInput) (*request.Request, *voiceid.DeleteWatchlistOutput)

	DescribeDomain(*voiceid.DescribeDomainInput) (*voiceid.DescribeDomainOutput, error)
	DescribeDomainWithContext(aws.Context, *voiceid.DescribeDomainInput, ...request.Option) (*voiceid.DescribeDomainOutput, error)
	DescribeDomainRequest(*voiceid.DescribeDomainInput) (*request.Request, *voiceid.DescribeDomainOutput)

	DescribeFraudster(*voiceid.DescribeFraudsterInput) (*voiceid.DescribeFraudsterOutput, error)
	DescribeFraudsterWithContext(aws.Context, *voiceid.DescribeFraudsterInput, ...request.Option) (*voiceid.DescribeFraudsterOutput, error)
	DescribeFraudsterRequest(*voiceid.DescribeFraudsterInput) (*request.Request, *voiceid.DescribeFraudsterOutput)

	DescribeFraudsterRegistrationJob(*voiceid.DescribeFraudsterRegistrationJobInput) (*voiceid.DescribeFraudsterRegistrationJobOutput, error)
	DescribeFraudsterRegistrationJobWithContext(aws.Context, *voiceid.DescribeFraudsterRegistrationJobInput, ...request.Option) (*voiceid.DescribeFraudsterRegistrationJobOutput, error)
	DescribeFraudsterRegistrationJobRequest(*voiceid.DescribeFraudsterRegistrationJobInput) (*request.Request, *voiceid.DescribeFraudsterRegistrationJobOutput)

	DescribeSpeaker(*voiceid.DescribeSpeakerInput) (*voiceid.DescribeSpeakerOutput, error)
	DescribeSpeakerWithContext(aws.Context, *voiceid.DescribeSpeakerInput, ...request.Option) (*voiceid.DescribeSpeakerOutput, error)
	DescribeSpeakerRequest(*voiceid.DescribeSpeakerInput) (*request.Request, *voiceid.DescribeSpeakerOutput)

	DescribeSpeakerEnrollmentJob(*voiceid.DescribeSpeakerEnrollmentJobInput) (*voiceid.DescribeSpeakerEnrollmentJobOutput, error)
	DescribeSpeakerEnrollmentJobWithContext(aws.Context, *voiceid.DescribeSpeakerEnrollmentJobInput, ...request.Option) (*voiceid.DescribeSpeakerEnrollmentJobOutput, error)
	DescribeSpeakerEnrollmentJobRequest(*voiceid.DescribeSpeakerEnrollmentJobInput) (*request.Request, *voiceid.DescribeSpeakerEnrollmentJobOutput)

	DescribeWatchlist(*voiceid.DescribeWatchlistInput) (*voiceid.DescribeWatchlistOutput, error)
	DescribeWatchlistWithContext(aws.Context, *voiceid.DescribeWatchlistInput, ...request.Option) (*voiceid.DescribeWatchlistOutput, error)
	DescribeWatchlistRequest(*voiceid.DescribeWatchlistInput) (*request.Request, *voiceid.DescribeWatchlistOutput)

	DisassociateFraudster(*voiceid.DisassociateFraudsterInput) (*voiceid.DisassociateFraudsterOutput, error)
	DisassociateFraudsterWithContext(aws.Context, *voiceid.DisassociateFraudsterInput, ...request.Option) (*voiceid.DisassociateFraudsterOutput, error)
	DisassociateFraudsterRequest(*voiceid.DisassociateFraudsterInput) (*request.Request, *voiceid.DisassociateFraudsterOutput)

	EvaluateSession(*voiceid.EvaluateSessionInput) (*voiceid.EvaluateSessionOutput, error)
	EvaluateSessionWithContext(aws.Context, *voiceid.EvaluateSessionInput, ...request.Option) (*voiceid.EvaluateSessionOutput, error)
	EvaluateSessionRequest(*voiceid.EvaluateSessionInput) (*request.Request, *voiceid.EvaluateSessionOutput)

	ListDomains(*voiceid.ListDomainsInput) (*voiceid.ListDomainsOutput, error)
	ListDomainsWithContext(aws.Context, *voiceid.ListDomainsInput, ...request.Option) (*voiceid.ListDomainsOutput, error)
	ListDomainsRequest(*voiceid.ListDomainsInput) (*request.Request, *voiceid.ListDomainsOutput)

	ListDomainsPages(*voiceid.ListDomainsInput, func(*voiceid.ListDomainsOutput, bool) bool) error
	ListDomainsPagesWithContext(aws.Context, *voiceid.ListDomainsInput, func(*voiceid.ListDomainsOutput, bool) bool, ...request.Option) error

	ListFraudsterRegistrationJobs(*voiceid.ListFraudsterRegistrationJobsInput) (*voiceid.ListFraudsterRegistrationJobsOutput, error)
	ListFraudsterRegistrationJobsWithContext(aws.Context, *voiceid.ListFraudsterRegistrationJobsInput, ...request.Option) (*voiceid.ListFraudsterRegistrationJobsOutput, error)
	ListFraudsterRegistrationJobsRequest(*voiceid.ListFraudsterRegistrationJobsInput) (*request.Request, *voiceid.ListFraudsterRegistrationJobsOutput)

	ListFraudsterRegistrationJobsPages(*voiceid.ListFraudsterRegistrationJobsInput, func(*voiceid.ListFraudsterRegistrationJobsOutput, bool) bool) error
	ListFraudsterRegistrationJobsPagesWithContext(aws.Context, *voiceid.ListFraudsterRegistrationJobsInput, func(*voiceid.ListFraudsterRegistrationJobsOutput, bool) bool, ...request.Option) error

	ListFraudsters(*voiceid.ListFraudstersInput) (*voiceid.ListFraudstersOutput, error)
	ListFraudstersWithContext(aws.Context, *voiceid.ListFraudstersInput, ...request.Option) (*voiceid.ListFraudstersOutput, error)
	ListFraudstersRequest(*voiceid.ListFraudstersInput) (*request.Request, *voiceid.ListFraudstersOutput)

	ListFraudstersPages(*voiceid.ListFraudstersInput, func(*voiceid.ListFraudstersOutput, bool) bool) error
	ListFraudstersPagesWithContext(aws.Context, *voiceid.ListFraudstersInput, func(*voiceid.ListFraudstersOutput, bool) bool, ...request.Option) error

	ListSpeakerEnrollmentJobs(*voiceid.ListSpeakerEnrollmentJobsInput) (*voiceid.ListSpeakerEnrollmentJobsOutput, error)
	ListSpeakerEnrollmentJobsWithContext(aws.Context, *voiceid.ListSpeakerEnrollmentJobsInput, ...request.Option) (*voiceid.ListSpeakerEnrollmentJobsOutput, error)
	ListSpeakerEnrollmentJobsRequest(*voiceid.ListSpeakerEnrollmentJobsInput) (*request.Request, *voiceid.ListSpeakerEnrollmentJobsOutput)

	ListSpeakerEnrollmentJobsPages(*voiceid.ListSpeakerEnrollmentJobsInput, func(*voiceid.ListSpeakerEnrollmentJobsOutput, bool) bool) error
	ListSpeakerEnrollmentJobsPagesWithContext(aws.Context, *voiceid.ListSpeakerEnrollmentJobsInput, func(*voiceid.ListSpeakerEnrollmentJobsOutput, bool) bool, ...request.Option) error

	ListSpeakers(*voiceid.ListSpeakersInput) (*voiceid.ListSpeakersOutput, error)
	ListSpeakersWithContext(aws.Context, *voiceid.ListSpeakersInput, ...request.Option) (*voiceid.ListSpeakersOutput, error)
	ListSpeakersRequest(*voiceid.ListSpeakersInput) (*request.Request, *voiceid.ListSpeakersOutput)

	ListSpeakersPages(*voiceid.ListSpeakersInput, func(*voiceid.ListSpeakersOutput, bool) bool) error
	ListSpeakersPagesWithContext(aws.Context, *voiceid.ListSpeakersInput, func(*voiceid.ListSpeakersOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*voiceid.ListTagsForResourceInput) (*voiceid.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *voiceid.ListTagsForResourceInput, ...request.Option) (*voiceid.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*voiceid.ListTagsForResourceInput) (*request.Request, *voiceid.ListTagsForResourceOutput)

	ListWatchlists(*voiceid.ListWatchlistsInput) (*voiceid.ListWatchlistsOutput, error)
	ListWatchlistsWithContext(aws.Context, *voiceid.ListWatchlistsInput, ...request.Option) (*voiceid.ListWatchlistsOutput, error)
	ListWatchlistsRequest(*voiceid.ListWatchlistsInput) (*request.Request, *voiceid.ListWatchlistsOutput)

	ListWatchlistsPages(*voiceid.ListWatchlistsInput, func(*voiceid.ListWatchlistsOutput, bool) bool) error
	ListWatchlistsPagesWithContext(aws.Context, *voiceid.ListWatchlistsInput, func(*voiceid.ListWatchlistsOutput, bool) bool, ...request.Option) error

	OptOutSpeaker(*voiceid.OptOutSpeakerInput) (*voiceid.OptOutSpeakerOutput, error)
	OptOutSpeakerWithContext(aws.Context, *voiceid.OptOutSpeakerInput, ...request.Option) (*voiceid.OptOutSpeakerOutput, error)
	OptOutSpeakerRequest(*voiceid.OptOutSpeakerInput) (*request.Request, *voiceid.OptOutSpeakerOutput)

	StartFraudsterRegistrationJob(*voiceid.StartFraudsterRegistrationJobInput) (*voiceid.StartFraudsterRegistrationJobOutput, error)
	StartFraudsterRegistrationJobWithContext(aws.Context, *voiceid.StartFraudsterRegistrationJobInput, ...request.Option) (*voiceid.StartFraudsterRegistrationJobOutput, error)
	StartFraudsterRegistrationJobRequest(*voiceid.StartFraudsterRegistrationJobInput) (*request.Request, *voiceid.StartFraudsterRegistrationJobOutput)

	StartSpeakerEnrollmentJob(*voiceid.StartSpeakerEnrollmentJobInput) (*voiceid.StartSpeakerEnrollmentJobOutput, error)
	StartSpeakerEnrollmentJobWithContext(aws.Context, *voiceid.StartSpeakerEnrollmentJobInput, ...request.Option) (*voiceid.StartSpeakerEnrollmentJobOutput, error)
	StartSpeakerEnrollmentJobRequest(*voiceid.StartSpeakerEnrollmentJobInput) (*request.Request, *voiceid.StartSpeakerEnrollmentJobOutput)

	TagResource(*voiceid.TagResourceInput) (*voiceid.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *voiceid.TagResourceInput, ...request.Option) (*voiceid.TagResourceOutput, error)
	TagResourceRequest(*voiceid.TagResourceInput) (*request.Request, *voiceid.TagResourceOutput)

	UntagResource(*voiceid.UntagResourceInput) (*voiceid.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *voiceid.UntagResourceInput, ...request.Option) (*voiceid.UntagResourceOutput, error)
	UntagResourceRequest(*voiceid.UntagResourceInput) (*request.Request, *voiceid.UntagResourceOutput)

	UpdateDomain(*voiceid.UpdateDomainInput) (*voiceid.UpdateDomainOutput, error)
	UpdateDomainWithContext(aws.Context, *voiceid.UpdateDomainInput, ...request.Option) (*voiceid.UpdateDomainOutput, error)
	UpdateDomainRequest(*voiceid.UpdateDomainInput) (*request.Request, *voiceid.UpdateDomainOutput)

	UpdateWatchlist(*voiceid.UpdateWatchlistInput) (*voiceid.UpdateWatchlistOutput, error)
	UpdateWatchlistWithContext(aws.Context, *voiceid.UpdateWatchlistInput, ...request.Option) (*voiceid.UpdateWatchlistOutput, error)
	UpdateWatchlistRequest(*voiceid.UpdateWatchlistInput) (*request.Request, *voiceid.UpdateWatchlistOutput)
}

var _ VoiceIDAPI = (*voiceid.VoiceID)(nil)
