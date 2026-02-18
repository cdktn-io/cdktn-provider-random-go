// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bytes

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BytesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The number of bytes requested. The minimum value for length is 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/random/3.8.1/docs/resources/bytes#length Bytes#length}
	Length *float64 `field:"required" json:"length" yaml:"length"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource.
	//
	// See [the main provider documentation](../index.html) for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/random/3.8.1/docs/resources/bytes#keepers Bytes#keepers}
	Keepers *map[string]*string `field:"optional" json:"keepers" yaml:"keepers"`
}

