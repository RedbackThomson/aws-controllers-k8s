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

package hyper_parameter_tuning_job

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
	_ = &svcapitypes.HyperParameterTuningJob{}
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

	resp, respErr := rm.sdkapi.DescribeHyperParameterTuningJobWithContext(ctx, input)
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
	if resp.HyperParameterTuningJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.HyperParameterTuningJobArn)
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
	return r.ko.Spec.HyperParameterTuningJobName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeHyperParameterTuningJobInput, error) {
	res := &svcsdk.DescribeHyperParameterTuningJobInput{}

	if r.ko.Spec.HyperParameterTuningJobName != nil {
		res.SetHyperParameterTuningJobName(*r.ko.Spec.HyperParameterTuningJobName)
	}

	return res, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.ListHyperParameterTuningJobsInput, error) {
	res := &svcsdk.ListHyperParameterTuningJobsInput{}

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

	resp, respErr := rm.sdkapi.CreateHyperParameterTuningJobWithContext(ctx, input)
	if respErr != nil {
		return nil, respErr
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.HyperParameterTuningJobArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.HyperParameterTuningJobArn)
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
) (*svcsdk.CreateHyperParameterTuningJobInput, error) {
	res := &svcsdk.CreateHyperParameterTuningJobInput{}

	if r.ko.Spec.HyperParameterTuningJobConfig != nil {
		f0 := &svcsdk.HyperParameterTuningJobConfig{}
		if r.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective != nil {
			f0f0 := &svcsdk.HyperParameterTuningJobObjective{}
			if r.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName != nil {
				f0f0.SetMetricName(*r.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.MetricName)
			}
			if r.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type != nil {
				f0f0.SetType(*r.ko.Spec.HyperParameterTuningJobConfig.HyperParameterTuningJobObjective.Type)
			}
			f0.SetHyperParameterTuningJobObjective(f0f0)
		}
		if r.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges != nil {
			f0f1 := &svcsdk.ParameterRanges{}
			if r.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges.CategoricalParameterRanges != nil {
				f0f1f0 := []*svcsdk.CategoricalParameterRange{}
				for _, f0f1f0iter := range r.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges.CategoricalParameterRanges {
					f0f1f0elem := &svcsdk.CategoricalParameterRange{}
					if f0f1f0iter.Name != nil {
						f0f1f0elem.SetName(*f0f1f0iter.Name)
					}
					if f0f1f0iter.Values != nil {
						f0f1f0elemf1 := []*string{}
						for _, f0f1f0elemf1iter := range f0f1f0iter.Values {
							var f0f1f0elemf1elem string
							f0f1f0elemf1elem = *f0f1f0elemf1iter
							f0f1f0elemf1 = append(f0f1f0elemf1, &f0f1f0elemf1elem)
						}
						f0f1f0elem.SetValues(f0f1f0elemf1)
					}
					f0f1f0 = append(f0f1f0, f0f1f0elem)
				}
				f0f1.SetCategoricalParameterRanges(f0f1f0)
			}
			if r.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges.ContinuousParameterRanges != nil {
				f0f1f1 := []*svcsdk.ContinuousParameterRange{}
				for _, f0f1f1iter := range r.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges.ContinuousParameterRanges {
					f0f1f1elem := &svcsdk.ContinuousParameterRange{}
					if f0f1f1iter.MaxValue != nil {
						f0f1f1elem.SetMaxValue(*f0f1f1iter.MaxValue)
					}
					if f0f1f1iter.MinValue != nil {
						f0f1f1elem.SetMinValue(*f0f1f1iter.MinValue)
					}
					if f0f1f1iter.Name != nil {
						f0f1f1elem.SetName(*f0f1f1iter.Name)
					}
					if f0f1f1iter.ScalingType != nil {
						f0f1f1elem.SetScalingType(*f0f1f1iter.ScalingType)
					}
					f0f1f1 = append(f0f1f1, f0f1f1elem)
				}
				f0f1.SetContinuousParameterRanges(f0f1f1)
			}
			if r.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges.IntegerParameterRanges != nil {
				f0f1f2 := []*svcsdk.IntegerParameterRange{}
				for _, f0f1f2iter := range r.ko.Spec.HyperParameterTuningJobConfig.ParameterRanges.IntegerParameterRanges {
					f0f1f2elem := &svcsdk.IntegerParameterRange{}
					if f0f1f2iter.MaxValue != nil {
						f0f1f2elem.SetMaxValue(*f0f1f2iter.MaxValue)
					}
					if f0f1f2iter.MinValue != nil {
						f0f1f2elem.SetMinValue(*f0f1f2iter.MinValue)
					}
					if f0f1f2iter.Name != nil {
						f0f1f2elem.SetName(*f0f1f2iter.Name)
					}
					if f0f1f2iter.ScalingType != nil {
						f0f1f2elem.SetScalingType(*f0f1f2iter.ScalingType)
					}
					f0f1f2 = append(f0f1f2, f0f1f2elem)
				}
				f0f1.SetIntegerParameterRanges(f0f1f2)
			}
			f0.SetParameterRanges(f0f1)
		}
		if r.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits != nil {
			f0f2 := &svcsdk.ResourceLimits{}
			if r.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs != nil {
				f0f2.SetMaxNumberOfTrainingJobs(*r.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxNumberOfTrainingJobs)
			}
			if r.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs != nil {
				f0f2.SetMaxParallelTrainingJobs(*r.ko.Spec.HyperParameterTuningJobConfig.ResourceLimits.MaxParallelTrainingJobs)
			}
			f0.SetResourceLimits(f0f2)
		}
		if r.ko.Spec.HyperParameterTuningJobConfig.Strategy != nil {
			f0.SetStrategy(*r.ko.Spec.HyperParameterTuningJobConfig.Strategy)
		}
		if r.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType != nil {
			f0.SetTrainingJobEarlyStoppingType(*r.ko.Spec.HyperParameterTuningJobConfig.TrainingJobEarlyStoppingType)
		}
		if r.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria != nil {
			f0f5 := &svcsdk.TuningJobCompletionCriteria{}
			if r.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue != nil {
				f0f5.SetTargetObjectiveMetricValue(*r.ko.Spec.HyperParameterTuningJobConfig.TuningJobCompletionCriteria.TargetObjectiveMetricValue)
			}
			f0.SetTuningJobCompletionCriteria(f0f5)
		}
		res.SetHyperParameterTuningJobConfig(f0)
	}
	if r.ko.Spec.HyperParameterTuningJobName != nil {
		res.SetHyperParameterTuningJobName(*r.ko.Spec.HyperParameterTuningJobName)
	}
	if r.ko.Spec.Tags != nil {
		f2 := []*svcsdk.Tag{}
		for _, f2iter := range r.ko.Spec.Tags {
			f2elem := &svcsdk.Tag{}
			if f2iter.Key != nil {
				f2elem.SetKey(*f2iter.Key)
			}
			if f2iter.Value != nil {
				f2elem.SetValue(*f2iter.Value)
			}
			f2 = append(f2, f2elem)
		}
		res.SetTags(f2)
	}
	if r.ko.Spec.TrainingJobDefinition != nil {
		f3 := &svcsdk.HyperParameterTrainingJobDefinition{}
		if r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification != nil {
			f3f0 := &svcsdk.HyperParameterAlgorithmSpecification{}
			if r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName != nil {
				f3f0.SetAlgorithmName(*r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.AlgorithmName)
			}
			if r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.MetricDefinitions != nil {
				f3f0f1 := []*svcsdk.MetricDefinition{}
				for _, f3f0f1iter := range r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.MetricDefinitions {
					f3f0f1elem := &svcsdk.MetricDefinition{}
					if f3f0f1iter.Name != nil {
						f3f0f1elem.SetName(*f3f0f1iter.Name)
					}
					if f3f0f1iter.Regex != nil {
						f3f0f1elem.SetRegex(*f3f0f1iter.Regex)
					}
					f3f0f1 = append(f3f0f1, f3f0f1elem)
				}
				f3f0.SetMetricDefinitions(f3f0f1)
			}
			if r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage != nil {
				f3f0.SetTrainingImage(*r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingImage)
			}
			if r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode != nil {
				f3f0.SetTrainingInputMode(*r.ko.Spec.TrainingJobDefinition.AlgorithmSpecification.TrainingInputMode)
			}
			f3.SetAlgorithmSpecification(f3f0)
		}
		if r.ko.Spec.TrainingJobDefinition.CheckpointConfig != nil {
			f3f1 := &svcsdk.CheckpointConfig{}
			if r.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath != nil {
				f3f1.SetLocalPath(*r.ko.Spec.TrainingJobDefinition.CheckpointConfig.LocalPath)
			}
			if r.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI != nil {
				f3f1.SetS3Uri(*r.ko.Spec.TrainingJobDefinition.CheckpointConfig.S3URI)
			}
			f3.SetCheckpointConfig(f3f1)
		}
		if r.ko.Spec.TrainingJobDefinition.DefinitionName != nil {
			f3.SetDefinitionName(*r.ko.Spec.TrainingJobDefinition.DefinitionName)
		}
		if r.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption != nil {
			f3.SetEnableInterContainerTrafficEncryption(*r.ko.Spec.TrainingJobDefinition.EnableInterContainerTrafficEncryption)
		}
		if r.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining != nil {
			f3.SetEnableManagedSpotTraining(*r.ko.Spec.TrainingJobDefinition.EnableManagedSpotTraining)
		}
		if r.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation != nil {
			f3.SetEnableNetworkIsolation(*r.ko.Spec.TrainingJobDefinition.EnableNetworkIsolation)
		}
		if r.ko.Spec.TrainingJobDefinition.HyperParameterRanges != nil {
			f3f6 := &svcsdk.ParameterRanges{}
			if r.ko.Spec.TrainingJobDefinition.HyperParameterRanges.CategoricalParameterRanges != nil {
				f3f6f0 := []*svcsdk.CategoricalParameterRange{}
				for _, f3f6f0iter := range r.ko.Spec.TrainingJobDefinition.HyperParameterRanges.CategoricalParameterRanges {
					f3f6f0elem := &svcsdk.CategoricalParameterRange{}
					if f3f6f0iter.Name != nil {
						f3f6f0elem.SetName(*f3f6f0iter.Name)
					}
					if f3f6f0iter.Values != nil {
						f3f6f0elemf1 := []*string{}
						for _, f3f6f0elemf1iter := range f3f6f0iter.Values {
							var f3f6f0elemf1elem string
							f3f6f0elemf1elem = *f3f6f0elemf1iter
							f3f6f0elemf1 = append(f3f6f0elemf1, &f3f6f0elemf1elem)
						}
						f3f6f0elem.SetValues(f3f6f0elemf1)
					}
					f3f6f0 = append(f3f6f0, f3f6f0elem)
				}
				f3f6.SetCategoricalParameterRanges(f3f6f0)
			}
			if r.ko.Spec.TrainingJobDefinition.HyperParameterRanges.ContinuousParameterRanges != nil {
				f3f6f1 := []*svcsdk.ContinuousParameterRange{}
				for _, f3f6f1iter := range r.ko.Spec.TrainingJobDefinition.HyperParameterRanges.ContinuousParameterRanges {
					f3f6f1elem := &svcsdk.ContinuousParameterRange{}
					if f3f6f1iter.MaxValue != nil {
						f3f6f1elem.SetMaxValue(*f3f6f1iter.MaxValue)
					}
					if f3f6f1iter.MinValue != nil {
						f3f6f1elem.SetMinValue(*f3f6f1iter.MinValue)
					}
					if f3f6f1iter.Name != nil {
						f3f6f1elem.SetName(*f3f6f1iter.Name)
					}
					if f3f6f1iter.ScalingType != nil {
						f3f6f1elem.SetScalingType(*f3f6f1iter.ScalingType)
					}
					f3f6f1 = append(f3f6f1, f3f6f1elem)
				}
				f3f6.SetContinuousParameterRanges(f3f6f1)
			}
			if r.ko.Spec.TrainingJobDefinition.HyperParameterRanges.IntegerParameterRanges != nil {
				f3f6f2 := []*svcsdk.IntegerParameterRange{}
				for _, f3f6f2iter := range r.ko.Spec.TrainingJobDefinition.HyperParameterRanges.IntegerParameterRanges {
					f3f6f2elem := &svcsdk.IntegerParameterRange{}
					if f3f6f2iter.MaxValue != nil {
						f3f6f2elem.SetMaxValue(*f3f6f2iter.MaxValue)
					}
					if f3f6f2iter.MinValue != nil {
						f3f6f2elem.SetMinValue(*f3f6f2iter.MinValue)
					}
					if f3f6f2iter.Name != nil {
						f3f6f2elem.SetName(*f3f6f2iter.Name)
					}
					if f3f6f2iter.ScalingType != nil {
						f3f6f2elem.SetScalingType(*f3f6f2iter.ScalingType)
					}
					f3f6f2 = append(f3f6f2, f3f6f2elem)
				}
				f3f6.SetIntegerParameterRanges(f3f6f2)
			}
			f3.SetHyperParameterRanges(f3f6)
		}
		if r.ko.Spec.TrainingJobDefinition.InputDataConfig != nil {
			f3f7 := []*svcsdk.Channel{}
			for _, f3f7iter := range r.ko.Spec.TrainingJobDefinition.InputDataConfig {
				f3f7elem := &svcsdk.Channel{}
				if f3f7iter.ChannelName != nil {
					f3f7elem.SetChannelName(*f3f7iter.ChannelName)
				}
				if f3f7iter.CompressionType != nil {
					f3f7elem.SetCompressionType(*f3f7iter.CompressionType)
				}
				if f3f7iter.ContentType != nil {
					f3f7elem.SetContentType(*f3f7iter.ContentType)
				}
				if f3f7iter.DataSource != nil {
					f3f7elemf3 := &svcsdk.DataSource{}
					if f3f7iter.DataSource.FileSystemDataSource != nil {
						f3f7elemf3f0 := &svcsdk.FileSystemDataSource{}
						if f3f7iter.DataSource.FileSystemDataSource.DirectoryPath != nil {
							f3f7elemf3f0.SetDirectoryPath(*f3f7iter.DataSource.FileSystemDataSource.DirectoryPath)
						}
						if f3f7iter.DataSource.FileSystemDataSource.FileSystemAccessMode != nil {
							f3f7elemf3f0.SetFileSystemAccessMode(*f3f7iter.DataSource.FileSystemDataSource.FileSystemAccessMode)
						}
						if f3f7iter.DataSource.FileSystemDataSource.FileSystemID != nil {
							f3f7elemf3f0.SetFileSystemId(*f3f7iter.DataSource.FileSystemDataSource.FileSystemID)
						}
						if f3f7iter.DataSource.FileSystemDataSource.FileSystemType != nil {
							f3f7elemf3f0.SetFileSystemType(*f3f7iter.DataSource.FileSystemDataSource.FileSystemType)
						}
						f3f7elemf3.SetFileSystemDataSource(f3f7elemf3f0)
					}
					if f3f7iter.DataSource.S3DataSource != nil {
						f3f7elemf3f1 := &svcsdk.S3DataSource{}
						if f3f7iter.DataSource.S3DataSource.AttributeNames != nil {
							f3f7elemf3f1f0 := []*string{}
							for _, f3f7elemf3f1f0iter := range f3f7iter.DataSource.S3DataSource.AttributeNames {
								var f3f7elemf3f1f0elem string
								f3f7elemf3f1f0elem = *f3f7elemf3f1f0iter
								f3f7elemf3f1f0 = append(f3f7elemf3f1f0, &f3f7elemf3f1f0elem)
							}
							f3f7elemf3f1.SetAttributeNames(f3f7elemf3f1f0)
						}
						if f3f7iter.DataSource.S3DataSource.S3DataDistributionType != nil {
							f3f7elemf3f1.SetS3DataDistributionType(*f3f7iter.DataSource.S3DataSource.S3DataDistributionType)
						}
						if f3f7iter.DataSource.S3DataSource.S3DataType != nil {
							f3f7elemf3f1.SetS3DataType(*f3f7iter.DataSource.S3DataSource.S3DataType)
						}
						if f3f7iter.DataSource.S3DataSource.S3URI != nil {
							f3f7elemf3f1.SetS3Uri(*f3f7iter.DataSource.S3DataSource.S3URI)
						}
						f3f7elemf3.SetS3DataSource(f3f7elemf3f1)
					}
					f3f7elem.SetDataSource(f3f7elemf3)
				}
				if f3f7iter.InputMode != nil {
					f3f7elem.SetInputMode(*f3f7iter.InputMode)
				}
				if f3f7iter.RecordWrapperType != nil {
					f3f7elem.SetRecordWrapperType(*f3f7iter.RecordWrapperType)
				}
				if f3f7iter.ShuffleConfig != nil {
					f3f7elemf6 := &svcsdk.ShuffleConfig{}
					if f3f7iter.ShuffleConfig.Seed != nil {
						f3f7elemf6.SetSeed(*f3f7iter.ShuffleConfig.Seed)
					}
					f3f7elem.SetShuffleConfig(f3f7elemf6)
				}
				f3f7 = append(f3f7, f3f7elem)
			}
			f3.SetInputDataConfig(f3f7)
		}
		if r.ko.Spec.TrainingJobDefinition.OutputDataConfig != nil {
			f3f8 := &svcsdk.OutputDataConfig{}
			if r.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID != nil {
				f3f8.SetKmsKeyId(*r.ko.Spec.TrainingJobDefinition.OutputDataConfig.KMSKeyID)
			}
			if r.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath != nil {
				f3f8.SetS3OutputPath(*r.ko.Spec.TrainingJobDefinition.OutputDataConfig.S3OutputPath)
			}
			f3.SetOutputDataConfig(f3f8)
		}
		if r.ko.Spec.TrainingJobDefinition.ResourceConfig != nil {
			f3f9 := &svcsdk.ResourceConfig{}
			if r.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount != nil {
				f3f9.SetInstanceCount(*r.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceCount)
			}
			if r.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType != nil {
				f3f9.SetInstanceType(*r.ko.Spec.TrainingJobDefinition.ResourceConfig.InstanceType)
			}
			if r.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID != nil {
				f3f9.SetVolumeKmsKeyId(*r.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeKMSKeyID)
			}
			if r.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB != nil {
				f3f9.SetVolumeSizeInGB(*r.ko.Spec.TrainingJobDefinition.ResourceConfig.VolumeSizeInGB)
			}
			f3.SetResourceConfig(f3f9)
		}
		if r.ko.Spec.TrainingJobDefinition.RoleARN != nil {
			f3.SetRoleArn(*r.ko.Spec.TrainingJobDefinition.RoleARN)
		}
		if r.ko.Spec.TrainingJobDefinition.StaticHyperParameters != nil {
			f3f11 := map[string]*string{}
			for f3f11key, f3f11valiter := range r.ko.Spec.TrainingJobDefinition.StaticHyperParameters {
				var f3f11val string
				f3f11val = *f3f11valiter
				f3f11[f3f11key] = &f3f11val
			}
			f3.SetStaticHyperParameters(f3f11)
		}
		if r.ko.Spec.TrainingJobDefinition.StoppingCondition != nil {
			f3f12 := &svcsdk.StoppingCondition{}
			if r.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds != nil {
				f3f12.SetMaxRuntimeInSeconds(*r.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxRuntimeInSeconds)
			}
			if r.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds != nil {
				f3f12.SetMaxWaitTimeInSeconds(*r.ko.Spec.TrainingJobDefinition.StoppingCondition.MaxWaitTimeInSeconds)
			}
			f3.SetStoppingCondition(f3f12)
		}
		if r.ko.Spec.TrainingJobDefinition.TuningObjective != nil {
			f3f13 := &svcsdk.HyperParameterTuningJobObjective{}
			if r.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName != nil {
				f3f13.SetMetricName(*r.ko.Spec.TrainingJobDefinition.TuningObjective.MetricName)
			}
			if r.ko.Spec.TrainingJobDefinition.TuningObjective.Type != nil {
				f3f13.SetType(*r.ko.Spec.TrainingJobDefinition.TuningObjective.Type)
			}
			f3.SetTuningObjective(f3f13)
		}
		if r.ko.Spec.TrainingJobDefinition.VPCConfig != nil {
			f3f14 := &svcsdk.VpcConfig{}
			if r.ko.Spec.TrainingJobDefinition.VPCConfig.SecurityGroupIDs != nil {
				f3f14f0 := []*string{}
				for _, f3f14f0iter := range r.ko.Spec.TrainingJobDefinition.VPCConfig.SecurityGroupIDs {
					var f3f14f0elem string
					f3f14f0elem = *f3f14f0iter
					f3f14f0 = append(f3f14f0, &f3f14f0elem)
				}
				f3f14.SetSecurityGroupIds(f3f14f0)
			}
			if r.ko.Spec.TrainingJobDefinition.VPCConfig.Subnets != nil {
				f3f14f1 := []*string{}
				for _, f3f14f1iter := range r.ko.Spec.TrainingJobDefinition.VPCConfig.Subnets {
					var f3f14f1elem string
					f3f14f1elem = *f3f14f1iter
					f3f14f1 = append(f3f14f1, &f3f14f1elem)
				}
				f3f14.SetSubnets(f3f14f1)
			}
			f3.SetVpcConfig(f3f14)
		}
		res.SetTrainingJobDefinition(f3)
	}
	if r.ko.Spec.TrainingJobDefinitions != nil {
		f4 := []*svcsdk.HyperParameterTrainingJobDefinition{}
		for _, f4iter := range r.ko.Spec.TrainingJobDefinitions {
			f4elem := &svcsdk.HyperParameterTrainingJobDefinition{}
			if f4iter.AlgorithmSpecification != nil {
				f4elemf0 := &svcsdk.HyperParameterAlgorithmSpecification{}
				if f4iter.AlgorithmSpecification.AlgorithmName != nil {
					f4elemf0.SetAlgorithmName(*f4iter.AlgorithmSpecification.AlgorithmName)
				}
				if f4iter.AlgorithmSpecification.MetricDefinitions != nil {
					f4elemf0f1 := []*svcsdk.MetricDefinition{}
					for _, f4elemf0f1iter := range f4iter.AlgorithmSpecification.MetricDefinitions {
						f4elemf0f1elem := &svcsdk.MetricDefinition{}
						if f4elemf0f1iter.Name != nil {
							f4elemf0f1elem.SetName(*f4elemf0f1iter.Name)
						}
						if f4elemf0f1iter.Regex != nil {
							f4elemf0f1elem.SetRegex(*f4elemf0f1iter.Regex)
						}
						f4elemf0f1 = append(f4elemf0f1, f4elemf0f1elem)
					}
					f4elemf0.SetMetricDefinitions(f4elemf0f1)
				}
				if f4iter.AlgorithmSpecification.TrainingImage != nil {
					f4elemf0.SetTrainingImage(*f4iter.AlgorithmSpecification.TrainingImage)
				}
				if f4iter.AlgorithmSpecification.TrainingInputMode != nil {
					f4elemf0.SetTrainingInputMode(*f4iter.AlgorithmSpecification.TrainingInputMode)
				}
				f4elem.SetAlgorithmSpecification(f4elemf0)
			}
			if f4iter.CheckpointConfig != nil {
				f4elemf1 := &svcsdk.CheckpointConfig{}
				if f4iter.CheckpointConfig.LocalPath != nil {
					f4elemf1.SetLocalPath(*f4iter.CheckpointConfig.LocalPath)
				}
				if f4iter.CheckpointConfig.S3URI != nil {
					f4elemf1.SetS3Uri(*f4iter.CheckpointConfig.S3URI)
				}
				f4elem.SetCheckpointConfig(f4elemf1)
			}
			if f4iter.DefinitionName != nil {
				f4elem.SetDefinitionName(*f4iter.DefinitionName)
			}
			if f4iter.EnableInterContainerTrafficEncryption != nil {
				f4elem.SetEnableInterContainerTrafficEncryption(*f4iter.EnableInterContainerTrafficEncryption)
			}
			if f4iter.EnableManagedSpotTraining != nil {
				f4elem.SetEnableManagedSpotTraining(*f4iter.EnableManagedSpotTraining)
			}
			if f4iter.EnableNetworkIsolation != nil {
				f4elem.SetEnableNetworkIsolation(*f4iter.EnableNetworkIsolation)
			}
			if f4iter.HyperParameterRanges != nil {
				f4elemf6 := &svcsdk.ParameterRanges{}
				if f4iter.HyperParameterRanges.CategoricalParameterRanges != nil {
					f4elemf6f0 := []*svcsdk.CategoricalParameterRange{}
					for _, f4elemf6f0iter := range f4iter.HyperParameterRanges.CategoricalParameterRanges {
						f4elemf6f0elem := &svcsdk.CategoricalParameterRange{}
						if f4elemf6f0iter.Name != nil {
							f4elemf6f0elem.SetName(*f4elemf6f0iter.Name)
						}
						if f4elemf6f0iter.Values != nil {
							f4elemf6f0elemf1 := []*string{}
							for _, f4elemf6f0elemf1iter := range f4elemf6f0iter.Values {
								var f4elemf6f0elemf1elem string
								f4elemf6f0elemf1elem = *f4elemf6f0elemf1iter
								f4elemf6f0elemf1 = append(f4elemf6f0elemf1, &f4elemf6f0elemf1elem)
							}
							f4elemf6f0elem.SetValues(f4elemf6f0elemf1)
						}
						f4elemf6f0 = append(f4elemf6f0, f4elemf6f0elem)
					}
					f4elemf6.SetCategoricalParameterRanges(f4elemf6f0)
				}
				if f4iter.HyperParameterRanges.ContinuousParameterRanges != nil {
					f4elemf6f1 := []*svcsdk.ContinuousParameterRange{}
					for _, f4elemf6f1iter := range f4iter.HyperParameterRanges.ContinuousParameterRanges {
						f4elemf6f1elem := &svcsdk.ContinuousParameterRange{}
						if f4elemf6f1iter.MaxValue != nil {
							f4elemf6f1elem.SetMaxValue(*f4elemf6f1iter.MaxValue)
						}
						if f4elemf6f1iter.MinValue != nil {
							f4elemf6f1elem.SetMinValue(*f4elemf6f1iter.MinValue)
						}
						if f4elemf6f1iter.Name != nil {
							f4elemf6f1elem.SetName(*f4elemf6f1iter.Name)
						}
						if f4elemf6f1iter.ScalingType != nil {
							f4elemf6f1elem.SetScalingType(*f4elemf6f1iter.ScalingType)
						}
						f4elemf6f1 = append(f4elemf6f1, f4elemf6f1elem)
					}
					f4elemf6.SetContinuousParameterRanges(f4elemf6f1)
				}
				if f4iter.HyperParameterRanges.IntegerParameterRanges != nil {
					f4elemf6f2 := []*svcsdk.IntegerParameterRange{}
					for _, f4elemf6f2iter := range f4iter.HyperParameterRanges.IntegerParameterRanges {
						f4elemf6f2elem := &svcsdk.IntegerParameterRange{}
						if f4elemf6f2iter.MaxValue != nil {
							f4elemf6f2elem.SetMaxValue(*f4elemf6f2iter.MaxValue)
						}
						if f4elemf6f2iter.MinValue != nil {
							f4elemf6f2elem.SetMinValue(*f4elemf6f2iter.MinValue)
						}
						if f4elemf6f2iter.Name != nil {
							f4elemf6f2elem.SetName(*f4elemf6f2iter.Name)
						}
						if f4elemf6f2iter.ScalingType != nil {
							f4elemf6f2elem.SetScalingType(*f4elemf6f2iter.ScalingType)
						}
						f4elemf6f2 = append(f4elemf6f2, f4elemf6f2elem)
					}
					f4elemf6.SetIntegerParameterRanges(f4elemf6f2)
				}
				f4elem.SetHyperParameterRanges(f4elemf6)
			}
			if f4iter.InputDataConfig != nil {
				f4elemf7 := []*svcsdk.Channel{}
				for _, f4elemf7iter := range f4iter.InputDataConfig {
					f4elemf7elem := &svcsdk.Channel{}
					if f4elemf7iter.ChannelName != nil {
						f4elemf7elem.SetChannelName(*f4elemf7iter.ChannelName)
					}
					if f4elemf7iter.CompressionType != nil {
						f4elemf7elem.SetCompressionType(*f4elemf7iter.CompressionType)
					}
					if f4elemf7iter.ContentType != nil {
						f4elemf7elem.SetContentType(*f4elemf7iter.ContentType)
					}
					if f4elemf7iter.DataSource != nil {
						f4elemf7elemf3 := &svcsdk.DataSource{}
						if f4elemf7iter.DataSource.FileSystemDataSource != nil {
							f4elemf7elemf3f0 := &svcsdk.FileSystemDataSource{}
							if f4elemf7iter.DataSource.FileSystemDataSource.DirectoryPath != nil {
								f4elemf7elemf3f0.SetDirectoryPath(*f4elemf7iter.DataSource.FileSystemDataSource.DirectoryPath)
							}
							if f4elemf7iter.DataSource.FileSystemDataSource.FileSystemAccessMode != nil {
								f4elemf7elemf3f0.SetFileSystemAccessMode(*f4elemf7iter.DataSource.FileSystemDataSource.FileSystemAccessMode)
							}
							if f4elemf7iter.DataSource.FileSystemDataSource.FileSystemID != nil {
								f4elemf7elemf3f0.SetFileSystemId(*f4elemf7iter.DataSource.FileSystemDataSource.FileSystemID)
							}
							if f4elemf7iter.DataSource.FileSystemDataSource.FileSystemType != nil {
								f4elemf7elemf3f0.SetFileSystemType(*f4elemf7iter.DataSource.FileSystemDataSource.FileSystemType)
							}
							f4elemf7elemf3.SetFileSystemDataSource(f4elemf7elemf3f0)
						}
						if f4elemf7iter.DataSource.S3DataSource != nil {
							f4elemf7elemf3f1 := &svcsdk.S3DataSource{}
							if f4elemf7iter.DataSource.S3DataSource.AttributeNames != nil {
								f4elemf7elemf3f1f0 := []*string{}
								for _, f4elemf7elemf3f1f0iter := range f4elemf7iter.DataSource.S3DataSource.AttributeNames {
									var f4elemf7elemf3f1f0elem string
									f4elemf7elemf3f1f0elem = *f4elemf7elemf3f1f0iter
									f4elemf7elemf3f1f0 = append(f4elemf7elemf3f1f0, &f4elemf7elemf3f1f0elem)
								}
								f4elemf7elemf3f1.SetAttributeNames(f4elemf7elemf3f1f0)
							}
							if f4elemf7iter.DataSource.S3DataSource.S3DataDistributionType != nil {
								f4elemf7elemf3f1.SetS3DataDistributionType(*f4elemf7iter.DataSource.S3DataSource.S3DataDistributionType)
							}
							if f4elemf7iter.DataSource.S3DataSource.S3DataType != nil {
								f4elemf7elemf3f1.SetS3DataType(*f4elemf7iter.DataSource.S3DataSource.S3DataType)
							}
							if f4elemf7iter.DataSource.S3DataSource.S3URI != nil {
								f4elemf7elemf3f1.SetS3Uri(*f4elemf7iter.DataSource.S3DataSource.S3URI)
							}
							f4elemf7elemf3.SetS3DataSource(f4elemf7elemf3f1)
						}
						f4elemf7elem.SetDataSource(f4elemf7elemf3)
					}
					if f4elemf7iter.InputMode != nil {
						f4elemf7elem.SetInputMode(*f4elemf7iter.InputMode)
					}
					if f4elemf7iter.RecordWrapperType != nil {
						f4elemf7elem.SetRecordWrapperType(*f4elemf7iter.RecordWrapperType)
					}
					if f4elemf7iter.ShuffleConfig != nil {
						f4elemf7elemf6 := &svcsdk.ShuffleConfig{}
						if f4elemf7iter.ShuffleConfig.Seed != nil {
							f4elemf7elemf6.SetSeed(*f4elemf7iter.ShuffleConfig.Seed)
						}
						f4elemf7elem.SetShuffleConfig(f4elemf7elemf6)
					}
					f4elemf7 = append(f4elemf7, f4elemf7elem)
				}
				f4elem.SetInputDataConfig(f4elemf7)
			}
			if f4iter.OutputDataConfig != nil {
				f4elemf8 := &svcsdk.OutputDataConfig{}
				if f4iter.OutputDataConfig.KMSKeyID != nil {
					f4elemf8.SetKmsKeyId(*f4iter.OutputDataConfig.KMSKeyID)
				}
				if f4iter.OutputDataConfig.S3OutputPath != nil {
					f4elemf8.SetS3OutputPath(*f4iter.OutputDataConfig.S3OutputPath)
				}
				f4elem.SetOutputDataConfig(f4elemf8)
			}
			if f4iter.ResourceConfig != nil {
				f4elemf9 := &svcsdk.ResourceConfig{}
				if f4iter.ResourceConfig.InstanceCount != nil {
					f4elemf9.SetInstanceCount(*f4iter.ResourceConfig.InstanceCount)
				}
				if f4iter.ResourceConfig.InstanceType != nil {
					f4elemf9.SetInstanceType(*f4iter.ResourceConfig.InstanceType)
				}
				if f4iter.ResourceConfig.VolumeKMSKeyID != nil {
					f4elemf9.SetVolumeKmsKeyId(*f4iter.ResourceConfig.VolumeKMSKeyID)
				}
				if f4iter.ResourceConfig.VolumeSizeInGB != nil {
					f4elemf9.SetVolumeSizeInGB(*f4iter.ResourceConfig.VolumeSizeInGB)
				}
				f4elem.SetResourceConfig(f4elemf9)
			}
			if f4iter.RoleARN != nil {
				f4elem.SetRoleArn(*f4iter.RoleARN)
			}
			if f4iter.StaticHyperParameters != nil {
				f4elemf11 := map[string]*string{}
				for f4elemf11key, f4elemf11valiter := range f4iter.StaticHyperParameters {
					var f4elemf11val string
					f4elemf11val = *f4elemf11valiter
					f4elemf11[f4elemf11key] = &f4elemf11val
				}
				f4elem.SetStaticHyperParameters(f4elemf11)
			}
			if f4iter.StoppingCondition != nil {
				f4elemf12 := &svcsdk.StoppingCondition{}
				if f4iter.StoppingCondition.MaxRuntimeInSeconds != nil {
					f4elemf12.SetMaxRuntimeInSeconds(*f4iter.StoppingCondition.MaxRuntimeInSeconds)
				}
				if f4iter.StoppingCondition.MaxWaitTimeInSeconds != nil {
					f4elemf12.SetMaxWaitTimeInSeconds(*f4iter.StoppingCondition.MaxWaitTimeInSeconds)
				}
				f4elem.SetStoppingCondition(f4elemf12)
			}
			if f4iter.TuningObjective != nil {
				f4elemf13 := &svcsdk.HyperParameterTuningJobObjective{}
				if f4iter.TuningObjective.MetricName != nil {
					f4elemf13.SetMetricName(*f4iter.TuningObjective.MetricName)
				}
				if f4iter.TuningObjective.Type != nil {
					f4elemf13.SetType(*f4iter.TuningObjective.Type)
				}
				f4elem.SetTuningObjective(f4elemf13)
			}
			if f4iter.VPCConfig != nil {
				f4elemf14 := &svcsdk.VpcConfig{}
				if f4iter.VPCConfig.SecurityGroupIDs != nil {
					f4elemf14f0 := []*string{}
					for _, f4elemf14f0iter := range f4iter.VPCConfig.SecurityGroupIDs {
						var f4elemf14f0elem string
						f4elemf14f0elem = *f4elemf14f0iter
						f4elemf14f0 = append(f4elemf14f0, &f4elemf14f0elem)
					}
					f4elemf14.SetSecurityGroupIds(f4elemf14f0)
				}
				if f4iter.VPCConfig.Subnets != nil {
					f4elemf14f1 := []*string{}
					for _, f4elemf14f1iter := range f4iter.VPCConfig.Subnets {
						var f4elemf14f1elem string
						f4elemf14f1elem = *f4elemf14f1iter
						f4elemf14f1 = append(f4elemf14f1, &f4elemf14f1elem)
					}
					f4elemf14.SetSubnets(f4elemf14f1)
				}
				f4elem.SetVpcConfig(f4elemf14)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTrainingJobDefinitions(f4)
	}
	if r.ko.Spec.WarmStartConfig != nil {
		f5 := &svcsdk.HyperParameterTuningJobWarmStartConfig{}
		if r.ko.Spec.WarmStartConfig.ParentHyperParameterTuningJobs != nil {
			f5f0 := []*svcsdk.ParentHyperParameterTuningJob{}
			for _, f5f0iter := range r.ko.Spec.WarmStartConfig.ParentHyperParameterTuningJobs {
				f5f0elem := &svcsdk.ParentHyperParameterTuningJob{}
				if f5f0iter.HyperParameterTuningJobName != nil {
					f5f0elem.SetHyperParameterTuningJobName(*f5f0iter.HyperParameterTuningJobName)
				}
				f5f0 = append(f5f0, f5f0elem)
			}
			f5.SetParentHyperParameterTuningJobs(f5f0)
		}
		if r.ko.Spec.WarmStartConfig.WarmStartType != nil {
			f5.SetWarmStartType(*r.ko.Spec.WarmStartConfig.WarmStartType)
		}
		res.SetWarmStartConfig(f5)
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
