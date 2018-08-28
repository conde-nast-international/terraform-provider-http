package main

import (
	"github.com/conde-nast-international/terraform-provider-http/http"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: http.Provider})
}
