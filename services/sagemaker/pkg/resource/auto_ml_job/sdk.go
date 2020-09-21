// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package auto_ml_job

import (
	"context"

	ackv1alpha1 "github.com/aws/aws-controllers-k8s/apis/core/v1alpha1"
	ackcompare "github.com/aws/aws-controllers-k8s/pkg/compare"
	ackerr "github.com/aws/aws-controllers-k8s/pkg/errors"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws/aws-controllers-k8s/services/sagemaker/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.AutoMLJob{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.DescribeAutoMLJobWithContext(ctx, input)
	if respErr != nil {
		if awsErr, ok := ackerr.AWSError(respErr); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.AutoMLJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.AutoMLJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required by not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.AutoMLJobName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeAutoMLJobInput, error) {
	res := &svcsdk.DescribeAutoMLJobInput{}

	if r.ko.Spec.AutoMLJobName != nil {
		res.SetAutoMLJobName(*r.ko.Spec.AutoMLJobName)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.ListAutoMLJobsInput, error) {
	res := &svcsdk.ListAutoMLJobsInput{}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a new resource with any fields in the Status field filled in
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	r *resource,
) (*resource, error) {
	input, err := rm.newCreateRequestPayload(r)
	if err != nil {
		return nil, err
	}

	resp, respErr := rm.sdkapi.CreateAutoMLJobWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.AutoMLJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.AutoMLJobArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	ko.Status.Conditions = []*ackv1alpha1.Condition{}
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	r *resource,
) (*svcsdk.CreateAutoMLJobInput, error) {
	res := &svcsdk.CreateAutoMLJobInput{}

	if r.ko.Spec.AutoMLJobConfig != nil {
		f0 := &svcsdk.AutoMLJobConfig{}
		if r.ko.Spec.AutoMLJobConfig.CompletionCriteria != nil {
			f0f0 := &svcsdk.AutoMLJobCompletionCriteria{}
			if r.ko.Spec.AutoMLJobConfig.CompletionCriteria.MaxAutoMLJobRuntimeInSeconds != nil {
				f0f0.SetMaxAutoMLJobRuntimeInSeconds(*r.ko.Spec.AutoMLJobConfig.CompletionCriteria.MaxAutoMLJobRuntimeInSeconds)
			}
			if r.ko.Spec.AutoMLJobConfig.CompletionCriteria.MaxCandidates != nil {
				f0f0.SetMaxCandidates(*r.ko.Spec.AutoMLJobConfig.CompletionCriteria.MaxCandidates)
			}
			if r.ko.Spec.AutoMLJobConfig.CompletionCriteria.MaxRuntimePerTrainingJobInSeconds != nil {
				f0f0.SetMaxRuntimePerTrainingJobInSeconds(*r.ko.Spec.AutoMLJobConfig.CompletionCriteria.MaxRuntimePerTrainingJobInSeconds)
			}
			f0.SetCompletionCriteria(f0f0)
		}
		if r.ko.Spec.AutoMLJobConfig.SecurityConfig != nil {
			f0f1 := &svcsdk.AutoMLSecurityConfig{}
			if r.ko.Spec.AutoMLJobConfig.SecurityConfig.EnableInterContainerTrafficEncryption != nil {
				f0f1.SetEnableInterContainerTrafficEncryption(*r.ko.Spec.AutoMLJobConfig.SecurityConfig.EnableInterContainerTrafficEncryption)
			}
			if r.ko.Spec.AutoMLJobConfig.SecurityConfig.VolumeKMSKeyID != nil {
				f0f1.SetVolumeKmsKeyId(*r.ko.Spec.AutoMLJobConfig.SecurityConfig.VolumeKMSKeyID)
			}
			if r.ko.Spec.AutoMLJobConfig.SecurityConfig.VPCConfig != nil {
				f0f1f2 := &svcsdk.VpcConfig{}
				if r.ko.Spec.AutoMLJobConfig.SecurityConfig.VPCConfig.SecurityGroupIDs != nil {
					f0f1f2f0 := []*string{}
					for _, f0f1f2f0iter := range r.ko.Spec.AutoMLJobConfig.SecurityConfig.VPCConfig.SecurityGroupIDs {
						var f0f1f2f0elem string
						f0f1f2f0elem = *f0f1f2f0iter
						f0f1f2f0 = append(f0f1f2f0, &f0f1f2f0elem)
					}
					f0f1f2.SetSecurityGroupIds(f0f1f2f0)
				}
				if r.ko.Spec.AutoMLJobConfig.SecurityConfig.VPCConfig.Subnets != nil {
					f0f1f2f1 := []*string{}
					for _, f0f1f2f1iter := range r.ko.Spec.AutoMLJobConfig.SecurityConfig.VPCConfig.Subnets {
						var f0f1f2f1elem string
						f0f1f2f1elem = *f0f1f2f1iter
						f0f1f2f1 = append(f0f1f2f1, &f0f1f2f1elem)
					}
					f0f1f2.SetSubnets(f0f1f2f1)
				}
				f0f1.SetVpcConfig(f0f1f2)
			}
			f0.SetSecurityConfig(f0f1)
		}
		res.SetAutoMLJobConfig(f0)
	}
	if r.ko.Spec.AutoMLJobName != nil {
		res.SetAutoMLJobName(*r.ko.Spec.AutoMLJobName)
	}
	if r.ko.Spec.AutoMLJobObjective != nil {
		f2 := &svcsdk.AutoMLJobObjective{}
		if r.ko.Spec.AutoMLJobObjective.MetricName != nil {
			f2.SetMetricName(*r.ko.Spec.AutoMLJobObjective.MetricName)
		}
		res.SetAutoMLJobObjective(f2)
	}
	if r.ko.Spec.GenerateCandidateDefinitionsOnly != nil {
		res.SetGenerateCandidateDefinitionsOnly(*r.ko.Spec.GenerateCandidateDefinitionsOnly)
	}
	if r.ko.Spec.InputDataConfig != nil {
		f4 := []*svcsdk.AutoMLChannel{}
		for _, f4iter := range r.ko.Spec.InputDataConfig {
			f4elem := &svcsdk.AutoMLChannel{}
			if f4iter.CompressionType != nil {
				f4elem.SetCompressionType(*f4iter.CompressionType)
			}
			if f4iter.DataSource != nil {
				f4elemf1 := &svcsdk.AutoMLDataSource{}
				if f4iter.DataSource.S3DataSource != nil {
					f4elemf1f0 := &svcsdk.AutoMLS3DataSource{}
					if f4iter.DataSource.S3DataSource.S3DataType != nil {
						f4elemf1f0.SetS3DataType(*f4iter.DataSource.S3DataSource.S3DataType)
					}
					if f4iter.DataSource.S3DataSource.S3URI != nil {
						f4elemf1f0.SetS3Uri(*f4iter.DataSource.S3DataSource.S3URI)
					}
					f4elemf1.SetS3DataSource(f4elemf1f0)
				}
				f4elem.SetDataSource(f4elemf1)
			}
			if f4iter.TargetAttributeName != nil {
				f4elem.SetTargetAttributeName(*f4iter.TargetAttributeName)
			}
			f4 = append(f4, f4elem)
		}
		res.SetInputDataConfig(f4)
	}
	if r.ko.Spec.OutputDataConfig != nil {
		f5 := &svcsdk.AutoMLOutputDataConfig{}
		if r.ko.Spec.OutputDataConfig.KMSKeyID != nil {
			f5.SetKmsKeyId(*r.ko.Spec.OutputDataConfig.KMSKeyID)
		}
		if r.ko.Spec.OutputDataConfig.S3OutputPath != nil {
			f5.SetS3OutputPath(*r.ko.Spec.OutputDataConfig.S3OutputPath)
		}
		res.SetOutputDataConfig(f5)
	}
	if r.ko.Spec.ProblemType != nil {
		res.SetProblemType(*r.ko.Spec.ProblemType)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}
	if r.ko.Spec.Tags != nil {
		f8 := []*svcsdk.Tag{}
		for _, f8iter := range r.ko.Spec.Tags {
			f8elem := &svcsdk.Tag{}
			if f8iter.Key != nil {
				f8elem.SetKey(*f8iter.Key)
			}
			if f8iter.Value != nil {
				f8elem.SetValue(*f8iter.Value)
			}
			f8 = append(f8, f8elem)
		}
		res.SetTags(f8)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	diffReporter *ackcompare.Reporter,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	// TODO(jaypipes): Figure this out...
	return nil

}
