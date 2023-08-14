package cdk8svarious

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-various-go/cdk8svarious/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/opencdk8s/cdk8s-various-go/cdk8svarious/internal"
)

type UntypedConfigMap interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for UntypedConfigMap
type jsiiProxy_UntypedConfigMap struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_UntypedConfigMap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewUntypedConfigMap(scope constructs.Construct, name *string, opts *UntypedConfigMapOptions) UntypedConfigMap {
	_init_.Initialize()

	if err := validateNewUntypedConfigMapParameters(scope, name, opts); err != nil {
		panic(err)
	}
	j := jsiiProxy_UntypedConfigMap{}

	_jsii_.Create(
		"cdk8s-various.UntypedConfigMap",
		[]interface{}{scope, name, opts},
		&j,
	)

	return &j
}

func NewUntypedConfigMap_Override(u UntypedConfigMap, scope constructs.Construct, name *string, opts *UntypedConfigMapOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-various.UntypedConfigMap",
		[]interface{}{scope, name, opts},
		u,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func UntypedConfigMap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateUntypedConfigMap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-various.UntypedConfigMap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UntypedConfigMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

