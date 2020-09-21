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

package notebook_instance_lifecycle_config

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
	_ = &svcapitypes.NotebookInstanceLifecycleConfig{}
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

	resp, respErr := rm.sdkapi.DescribeNotebookInstanceLifecycleConfigWithContext(ctx, input)
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
	if resp.NotebookInstanceLifecycleConfigArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.NotebookInstanceLifecycleConfigArn)
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
	return r.ko.Spec.NotebookInstanceLifecycleConfigName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeNotebookInstanceLifecycleConfigInput, error) {
	res := &svcsdk.DescribeNotebookInstanceLifecycleConfigInput{}

	if r.ko.Spec.NotebookInstanceLifecycleConfigName != nil {
		res.SetNotebookInstanceLifecycleConfigName(*r.ko.Spec.NotebookInstanceLifecycleConfigName)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.ListNotebookInstanceLifecycleConfigsInput, error) {
	res := &svcsdk.ListNotebookInstanceLifecycleConfigsInput{}

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

	resp, respErr := rm.sdkapi.CreateNotebookInstanceLifecycleConfigWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.NotebookInstanceLifecycleConfigArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.NotebookInstanceLifecycleConfigArn)
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
) (*svcsdk.CreateNotebookInstanceLifecycleConfigInput, error) {
	res := &svcsdk.CreateNotebookInstanceLifecycleConfigInput{}

	if r.ko.Spec.NotebookInstanceLifecycleConfigName != nil {
		res.SetNotebookInstanceLifecycleConfigName(*r.ko.Spec.NotebookInstanceLifecycleConfigName)
	}
	if r.ko.Spec.OnCreate != nil {
		f1 := []*svcsdk.NotebookInstanceLifecycleHook{}
		for _, f1iter := range r.ko.Spec.OnCreate {
			f1elem := &svcsdk.NotebookInstanceLifecycleHook{}
			if f1iter.Content != nil {
				f1elem.SetContent(*f1iter.Content)
			}
			f1 = append(f1, f1elem)
		}
		res.SetOnCreate(f1)
	}
	if r.ko.Spec.OnStart != nil {
		f2 := []*svcsdk.NotebookInstanceLifecycleHook{}
		for _, f2iter := range r.ko.Spec.OnStart {
			f2elem := &svcsdk.NotebookInstanceLifecycleHook{}
			if f2iter.Content != nil {
				f2elem.SetContent(*f2iter.Content)
			}
			f2 = append(f2, f2elem)
		}
		res.SetOnStart(f2)
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

	input, err := rm.newUpdateRequestPayload(desired)
	if err != nil {
		return nil, err
	}

	_, respErr := rm.sdkapi.UpdateNotebookInstanceLifecycleConfigWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	r *resource,
) (*svcsdk.UpdateNotebookInstanceLifecycleConfigInput, error) {
	res := &svcsdk.UpdateNotebookInstanceLifecycleConfigInput{}

	if r.ko.Spec.NotebookInstanceLifecycleConfigName != nil {
		res.SetNotebookInstanceLifecycleConfigName(*r.ko.Spec.NotebookInstanceLifecycleConfigName)
	}
	if r.ko.Spec.OnCreate != nil {
		f1 := []*svcsdk.NotebookInstanceLifecycleHook{}
		for _, f1iter := range r.ko.Spec.OnCreate {
			f1elem := &svcsdk.NotebookInstanceLifecycleHook{}
			if f1iter.Content != nil {
				f1elem.SetContent(*f1iter.Content)
			}
			f1 = append(f1, f1elem)
		}
		res.SetOnCreate(f1)
	}
	if r.ko.Spec.OnStart != nil {
		f2 := []*svcsdk.NotebookInstanceLifecycleHook{}
		for _, f2iter := range r.ko.Spec.OnStart {
			f2elem := &svcsdk.NotebookInstanceLifecycleHook{}
			if f2iter.Content != nil {
				f2elem.SetContent(*f2iter.Content)
			}
			f2 = append(f2, f2elem)
		}
		res.SetOnStart(f2)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) error {
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, respErr := rm.sdkapi.DeleteNotebookInstanceLifecycleConfigWithContext(ctx, input)
	return respErr
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteNotebookInstanceLifecycleConfigInput, error) {
	res := &svcsdk.DeleteNotebookInstanceLifecycleConfigInput{}

	if r.ko.Spec.NotebookInstanceLifecycleConfigName != nil {
		res.SetNotebookInstanceLifecycleConfigName(*r.ko.Spec.NotebookInstanceLifecycleConfigName)
	}

	return res, nil
}
