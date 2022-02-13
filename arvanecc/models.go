package arvanecc

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SSHKey -
type SSHKey struct {
	Region      types.String `tfsdk:"region"`
	Name        types.String `tfsdk:"name"`
	PublicKey   types.String `tfsdk:"public_key"`
	CreatedAt   types.String `tfsdk:"created_at"`
	Fingerprint types.String `tfsdk:"fingerprint"`
}

// // OrderItem -
// type OrderItem struct {
// 	Coffee   Coffee `tfsdk:"coffee"`
// 	Quantity int    `tfsdk:"quantity"`
// }

// // Coffee -
// // This Coffee struct is for Order.Items[].Coffee which does not have an
// // ingredients field in the schema defined by plugin framework. Since the
// // resource schema must match the struct exactly (extra field will return an
// // error). This struct has Ingredients commented out.
// type Coffee struct {
// 	ID          int          `tfsdk:"id"`
// 	Name        types.String `tfsdk:"name"`
// 	Teaser      types.String `tfsdk:"teaser"`
// 	Description types.String `tfsdk:"description"`
// 	Price       types.Number `tfsdk:"price"`
// 	Image       types.String `tfsdk:"image"`
// 	// Ingredients []Ingredient   `tfsdk:"ingredients"`
// }

// // Ingredient -
// // type Ingredient struct {
// // 	ID       int    `tfsdk:"ingredient_id"`
// // 	Name     string `tfsdk:"name"`
// // 	Quantity int    `tfsdk:"quantity"`
// // 	Unit     string `tfsdk:"unit"`
// // }
