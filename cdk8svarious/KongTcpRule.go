package cdk8svarious


type KongTcpRule struct {
	Backend *KongTcpBackend `field:"required" json:"backend" yaml:"backend"`
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

