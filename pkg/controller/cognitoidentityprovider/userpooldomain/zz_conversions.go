/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package userpooldomain

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	svcapitypes "github.com/crossplane/provider-aws/apis/cognitoidentityprovider/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeUserPoolDomainInput returns input for read
// operation.
func GenerateDescribeUserPoolDomainInput(cr *svcapitypes.UserPoolDomain) *svcsdk.DescribeUserPoolDomainInput {
	res := &svcsdk.DescribeUserPoolDomainInput{}

	return res
}

// GenerateUserPoolDomain returns the current state in the form of *svcapitypes.UserPoolDomain.
func GenerateUserPoolDomain(resp *svcsdk.DescribeUserPoolDomainOutput) *svcapitypes.UserPoolDomain {
	cr := &svcapitypes.UserPoolDomain{}

	if resp.DomainDescription.CustomDomainConfig != nil {
		f2 := &svcapitypes.CustomDomainConfigType{}
		if resp.DomainDescription.CustomDomainConfig.CertificateArn != nil {
			f2.CertificateARN = resp.DomainDescription.CustomDomainConfig.CertificateArn
		}
		cr.Spec.ForProvider.CustomDomainConfig = f2
	} else {
		cr.Spec.ForProvider.CustomDomainConfig = nil
	}

	return cr
}

// GenerateCreateUserPoolDomainInput returns a create input.
func GenerateCreateUserPoolDomainInput(cr *svcapitypes.UserPoolDomain) *svcsdk.CreateUserPoolDomainInput {
	res := &svcsdk.CreateUserPoolDomainInput{}

	if cr.Spec.ForProvider.CustomDomainConfig != nil {
		f0 := &svcsdk.CustomDomainConfigType{}
		if cr.Spec.ForProvider.CustomDomainConfig.CertificateARN != nil {
			f0.SetCertificateArn(*cr.Spec.ForProvider.CustomDomainConfig.CertificateARN)
		}
		res.SetCustomDomainConfig(f0)
	}

	return res
}

// GenerateUpdateUserPoolDomainInput returns an update input.
func GenerateUpdateUserPoolDomainInput(cr *svcapitypes.UserPoolDomain) *svcsdk.UpdateUserPoolDomainInput {
	res := &svcsdk.UpdateUserPoolDomainInput{}

	if cr.Spec.ForProvider.CustomDomainConfig != nil {
		f0 := &svcsdk.CustomDomainConfigType{}
		if cr.Spec.ForProvider.CustomDomainConfig.CertificateARN != nil {
			f0.SetCertificateArn(*cr.Spec.ForProvider.CustomDomainConfig.CertificateARN)
		}
		res.SetCustomDomainConfig(f0)
	}

	return res
}

// GenerateDeleteUserPoolDomainInput returns a deletion input.
func GenerateDeleteUserPoolDomainInput(cr *svcapitypes.UserPoolDomain) *svcsdk.DeleteUserPoolDomainInput {
	res := &svcsdk.DeleteUserPoolDomainInput{}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ResourceNotFoundException"
}
