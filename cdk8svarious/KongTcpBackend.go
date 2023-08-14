package cdk8svarious


type KongTcpBackend struct {
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	ServicePort *float64 `field:"required" json:"servicePort" yaml:"servicePort"`
}

