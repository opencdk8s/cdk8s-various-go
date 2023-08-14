//go:build no_runtime_type_checking

package cdk8svarious

// Building without runtime type checking enabled, so all the below just return nil

func validateSecurityGroupPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSecurityGroupPolicyParameters(scope constructs.Construct, name *string, opts *SecurityGroupPolicyOptions) error {
	return nil
}

