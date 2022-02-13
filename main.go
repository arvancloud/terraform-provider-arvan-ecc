package main

import (
	"context"

	"terraform-provider-arvan-ecc/arvanecc"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func main() {
	tfsdk.Serve(context.Background(), arvanecc.New, tfsdk.ServeOpts{
		Name: "arvanecc",
	})
}
