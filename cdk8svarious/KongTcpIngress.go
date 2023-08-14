package cdk8svarious

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-various-go/cdk8svarious/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/opencdk8s/cdk8s-various-go/cdk8svarious/internal"
)

type KongTcpIngress interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for KongTcpIngress
type jsiiProxy_KongTcpIngress struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_KongTcpIngress) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewKongTcpIngress(scope constructs.Construct, name *string, opts *KongTcpOptions) KongTcpIngress {
	_init_.Initialize()

	if err := validateNewKongTcpIngressParameters(scope, name, opts); err != nil {
		panic(err)
	}
	j := jsiiProxy_KongTcpIngress{}

	_jsii_.Create(
		"cdk8s-various.KongTcpIngress",
		[]interface{}{scope, name, opts},
		&j,
	)

	return &j
}

func NewKongTcpIngress_Override(k KongTcpIngress, scope constructs.Construct, name *string, opts *KongTcpOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-various.KongTcpIngress",
		[]interface{}{scope, name, opts},
		k,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func KongTcpIngress_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKongTcpIngress_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-various.KongTcpIngress",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KongTcpIngress) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

