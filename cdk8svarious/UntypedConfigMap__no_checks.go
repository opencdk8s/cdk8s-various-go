//go:build no_runtime_type_checking

package cdk8svarious

// Building without runtime type checking enabled, so all the below just return nil

func validateUntypedConfigMap_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewUntypedConfigMapParameters(scope constructs.Construct, name *string, opts *UntypedConfigMapOptions) error {
	return nil
}

