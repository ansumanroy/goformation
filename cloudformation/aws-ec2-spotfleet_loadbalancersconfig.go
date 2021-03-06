package cloudformation

// AWSEC2SpotFleet_LoadBalancersConfig AWS CloudFormation Resource (AWS::EC2::SpotFleet.LoadBalancersConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-loadbalancersconfig.html
type AWSEC2SpotFleet_LoadBalancersConfig struct {

	// ClassicLoadBalancersConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-loadbalancersconfig.html#cfn-ec2-spotfleet-loadbalancersconfig-classicloadbalancersconfig
	ClassicLoadBalancersConfig *AWSEC2SpotFleet_ClassicLoadBalancersConfig `json:"ClassicLoadBalancersConfig,omitempty"`

	// TargetGroupsConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-loadbalancersconfig.html#cfn-ec2-spotfleet-loadbalancersconfig-targetgroupsconfig
	TargetGroupsConfig *AWSEC2SpotFleet_TargetGroupsConfig `json:"TargetGroupsConfig,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SpotFleet_LoadBalancersConfig) AWSCloudFormationType() string {
	return "AWS::EC2::SpotFleet.LoadBalancersConfig"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2SpotFleet_LoadBalancersConfig) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
