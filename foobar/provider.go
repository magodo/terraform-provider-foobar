package foobar

import (
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/magodo/terraform-provider-foobar/foobar/internal/provider"
)

func Provider() terraform.ResourceProvider {
	return provider.Provider()
}
