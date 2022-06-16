package cdk8svarious

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk8s-various.SecurityGroupPolicy",
		reflect.TypeOf((*SecurityGroupPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityGroupPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-various.SecurityGroupPolicyOptions",
		reflect.TypeOf((*SecurityGroupPolicyOptions)(nil)).Elem(),
	)
}
