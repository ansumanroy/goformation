package resources

// AWS::S3::Bucket.RoutingRuleCondition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html
type AWSS3Bucket_RoutingRuleCondition struct {

	// HttpErrorCodeReturnedEquals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition-httperrorcodereturnedequals
	HttpErrorCodeReturnedEquals string `json:"HttpErrorCodeReturnedEquals"`

	// KeyPrefixEquals AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-routingrulecondition.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition-keyprefixequals
	KeyPrefixEquals string `json:"KeyPrefixEquals"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_RoutingRuleCondition) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.RoutingRuleCondition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_RoutingRuleCondition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
