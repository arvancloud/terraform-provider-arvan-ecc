module terraform-provider-arvan-ecc

go 1.16

replace github.com/arvancloud/ecc-go-client => /Users/samsambabadi/projects/arvan/terraform/ecc-go-client

require (
	github.com/arvancloud/ecc-go-client v0.0.0-00010101000000-000000000000
	github.com/hashicorp/terraform-plugin-framework v0.5.0
	github.com/hashicorp/terraform-plugin-go v0.4.0
)
