package cdk8svarious


type SecurityGroupPolicyOptions struct {
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	Selector *map[string]*string `field:"optional" json:"selector" yaml:"selector"`
}

