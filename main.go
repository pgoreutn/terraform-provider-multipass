package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-multipass/multipass"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: multipass.Provider})
}
