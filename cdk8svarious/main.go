// cdk8s-various
package cdk8svarious

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"cdk8s-various.KongTcpBackend",
		reflect.TypeOf((*KongTcpBackend)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk8s-various.KongTcpIngress",
		reflect.TypeOf((*KongTcpIngress)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KongTcpIngress{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-various.KongTcpOptions",
		reflect.TypeOf((*KongTcpOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk8s-various.KongTcpRule",
		reflect.TypeOf((*KongTcpRule)(nil)).Elem(),
	)
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
	_jsii_.RegisterClass(
		"cdk8s-various.UntypedConfigMap",
		reflect.TypeOf((*UntypedConfigMap)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_UntypedConfigMap{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk8s-various.UntypedConfigMapOptions",
		reflect.TypeOf((*UntypedConfigMapOptions)(nil)).Elem(),
	)
}
