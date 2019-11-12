package multipass

// Package multipass serves as an entrypoint, returning the list of available
// resources for the plugin.

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"os"
)

func init() {
	// Terraform is already adding the timestamp for us
	log.SetFlags(log.Lshortfile)
	log.SetPrefix(fmt.Sprintf("pid-%d-", os.Getpid()))
}

// Provider returns a resource provider for multipass.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"multipass_vm": resourceItem(),
		},
	}
}
