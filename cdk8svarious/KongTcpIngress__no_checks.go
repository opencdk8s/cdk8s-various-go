//go:build no_runtime_type_checking

package cdk8svarious

// Building without runtime type checking enabled, so all the below just return nil

func validateKongTcpIngress_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewKongTcpIngressParameters(scope constructs.Construct, name *string, opts *KongTcpOptions) error {
	return nil
}

