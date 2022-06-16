// cdk8s-various
package cdk8svarious

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/opencdk8s/cdk8s-various-go/cdk8svarious/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/opencdk8s/cdk8s-various-go/cdk8svarious/internal"
)

type SecurityGroupPolicy interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SecurityGroupPolicy
type jsiiProxy_SecurityGroupPolicy struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SecurityGroupPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewSecurityGroupPolicy(scope constructs.Construct, name *string, opts *SecurityGroupPolicyOptions) SecurityGroupPolicy {
	_init_.Initialize()

	j := jsiiProxy_SecurityGroupPolicy{}

	_jsii_.Create(
		"cdk8s-various.SecurityGroupPolicy",
		[]interface{}{scope, name, opts},
		&j,
	)

	return &j
}

func NewSecurityGroupPolicy_Override(s SecurityGroupPolicy, scope constructs.Construct, name *string, opts *SecurityGroupPolicyOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-various.SecurityGroupPolicy",
		[]interface{}{scope, name, opts},
		s,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func SecurityGroupPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdk8s-various.SecurityGroupPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroupPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SecurityGroupPolicyOptions struct {
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	Selector *map[string]*string `field:"optional" json:"selector" yaml:"selector"`
}

