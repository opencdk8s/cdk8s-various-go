package cdk8svarious

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type UntypedConfigMapOptions struct {
	Data *string `field:"required" json:"data" yaml:"data"`
	Metadata *cdk8s.ApiObjectMetadata `field:"required" json:"metadata" yaml:"metadata"`
}

