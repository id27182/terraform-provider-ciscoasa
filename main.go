package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-providers/terraform-provider-ciscoasa/ciscoasa"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ciscoasa.Provider,
	})
}
