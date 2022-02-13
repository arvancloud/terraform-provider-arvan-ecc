package arvanecc

import (
	"context"

	"github.com/arvancloud/ecc-go-client/resource"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type resourceSSHKeyType struct{}

// SSHKey Resource schema
func (r resourceSSHKeyType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"region": {
				Type:     types.StringType,
				Required: true,
				Computed: false,
			},
			"name": {
				Type:     types.StringType,
				Required: true,
				Computed: false,
			},
			"public_key": {
				Type:      types.StringType,
				Required:  true,
				Computed:  false,
				Sensitive: true,
			},
			"created_at": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
			"fingerprint": {
				Type:     types.StringType,
				Optional: true,
				Computed: true,
			},
		},
	}, nil
}

// New resource instance
func (r resourceSSHKeyType) NewResource(_ context.Context, p tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	return resourceSSHKey{
		p: *(p.(*provider)),
	}, nil
}

type resourceSSHKey struct {
	p provider
}

// Create a new resource
func (r resourceSSHKey) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Arvan Provider not configured",
			"The Arvan provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	// Retrieve values from sshkey
	var sshkey SSHKey
	diags := req.Plan.Get(ctx, &sshkey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, res, err := r.p.client.SSHKeysApi.Delete(context.TODO(), sshkey.Region.Value, sshkey.Name.Value)

	if err != nil && res.StatusCode != 404 {
		resp.Diagnostics.AddError(
			"Error on delete SSH Key before create it.",
			"Try to delete \"\""+sshkey.Name.Value+" SSH Key and request get error with message: "+err.Error()+"\n"+res.Status,
		)
		return
	}

	var body resource.SSHKey
	body.Name = sshkey.Name.Value
	body.PublicKey = sshkey.PublicKey.Value

	resBody, res, err := r.p.client.SSHKeysApi.Create(context.TODO(), sshkey.Region.Value, body)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error on create SSH Key",
			"Try to create \"\""+sshkey.Name.Value+" SSH Key and request get error with message: "+err.Error()+"\n"+res.Status,
		)
		return
	}

	sshKeyResult := transformResponseToSSHkey(resBody)
	sshKeyResult.Region.Value = sshkey.Region.Value

	diags = resp.State.Set(ctx, sshKeyResult)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information
func (r resourceSSHKey) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Arvan Provider not configured",
			"The Arvan provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	// Get current state
	var state SSHKey
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resBody, res, err := r.p.client.SSHKeysApi.Get(context.TODO(), state.Region.Value, state.Name.Value)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error on read SSH Key",
			"Try to read \"\""+state.Name.Value+" SSH Key and request get error with message: "+err.Error()+"\n"+res.Status,
		)
		return
	}

	sshkey := transformResponseToSSHkey(resBody)
	sshkey.Region.Value = state.Region.Value

	// Set state
	diags = resp.State.Set(ctx, &sshkey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update resource
func (r resourceSSHKey) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Arvan Provider not configured",
			"The Arvan provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	// Retrieve values from sshkey
	var sshkey SSHKey
	diags := req.Plan.Get(ctx, &sshkey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, res, err := r.p.client.SSHKeysApi.Delete(context.TODO(), sshkey.Region.Value, sshkey.Name.Value)

	if err != nil && res.StatusCode != 404 {
		resp.Diagnostics.AddError(
			"Error on delete SSH Key before create it.",
			"Try to delete \"\""+sshkey.Name.Value+" SSH Key and request get error with message: "+err.Error()+"\n"+res.Status,
		)
		return
	}

	var body resource.SSHKey
	body.Name = sshkey.Name.Value
	body.PublicKey = sshkey.PublicKey.Value

	resBody, res, err := r.p.client.SSHKeysApi.Create(context.TODO(), sshkey.Region.Value, body)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error on create SSH Key",
			"Try to create \"\""+sshkey.Name.Value+" SSH Key and request get error with message: "+err.Error()+"\n"+res.Status,
		)
		return
	}

	sshKeyResult := transformResponseToSSHkey(resBody)
	sshKeyResult.Region.Value = sshkey.Region.Value

	diags = resp.State.Set(ctx, sshKeyResult)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete resource
func (r resourceSSHKey) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	if !r.p.configured {
		resp.Diagnostics.AddError(
			"Arvan Provider not configured",
			"The Arvan provider hasn't been configured before apply, likely because it depends on an unknown value from another resource. This leads to weird stuff happening, so we'd prefer if you didn't do that. Thanks!",
		)
		return
	}

	// Retrieve values from sshkey
	var sshkey SSHKey
	diags := req.State.Get(ctx, &sshkey)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, res, err := r.p.client.SSHKeysApi.Delete(context.TODO(), sshkey.Region.Value, sshkey.Name.Value)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error on delete SSH Key",
			"Try to delete \"\""+sshkey.Name.Value+" SSH Key and request get error with message: "+err.Error()+"\n"+res.Status,
		)
		return
	}
}

// Import resource
func (r resourceSSHKey) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	// Save the import identifier in the id attribute
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("name"), req, resp)
}

func transformResponseToSSHkey(body interface{}) SSHKey {

	data := body.(map[string]map[string]string)

	var sshKey = SSHKey{
		Name:        types.String{Value: data["data"]["name"]},
		PublicKey:   types.String{Value: data["data"]["public_key"]},
		CreatedAt:   types.String{Value: data["data"]["created_at"]},
		Fingerprint: types.String{Value: data["data"]["fingerprint"]},
	}

	return sshKey
}
