// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ssmiface provides an interface for the Amazon Simple Systems Management Service.
package ssmiface

import (
	"github.com/awslabs/aws-sdk-go/service/ssm"
)

// SSMAPI is the interface type for ssm.SSM.
type SSMAPI interface {
	CreateAssociation(*ssm.CreateAssociationInput) (*ssm.CreateAssociationOutput, error)

	CreateAssociationBatch(*ssm.CreateAssociationBatchInput) (*ssm.CreateAssociationBatchOutput, error)

	CreateDocument(*ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error)

	DeleteAssociation(*ssm.DeleteAssociationInput) (*ssm.DeleteAssociationOutput, error)

	DeleteDocument(*ssm.DeleteDocumentInput) (*ssm.DeleteDocumentOutput, error)

	DescribeAssociation(*ssm.DescribeAssociationInput) (*ssm.DescribeAssociationOutput, error)

	DescribeDocument(*ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error)

	GetDocument(*ssm.GetDocumentInput) (*ssm.GetDocumentOutput, error)

	ListAssociations(*ssm.ListAssociationsInput) (*ssm.ListAssociationsOutput, error)

	ListDocuments(*ssm.ListDocumentsInput) (*ssm.ListDocumentsOutput, error)

	UpdateAssociationStatus(*ssm.UpdateAssociationStatusInput) (*ssm.UpdateAssociationStatusOutput, error)
}
