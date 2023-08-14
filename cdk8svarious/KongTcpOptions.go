package cdk8svarious


type KongTcpOptions struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Rules *[]*KongTcpRule `field:"optional" json:"rules" yaml:"rules"`
}

