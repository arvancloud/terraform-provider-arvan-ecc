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

// Server -
type Server struct {
	ID             types.String    `tfsdk:"id"`
	Name           types.String    `tfsdk:"name,omitempty"`
	DiskSize       types.Int64     `tfsdk:"disk_size,omitempty"`
	NetworkIds     []types.String  `tfsdk:"network_ids,omitempty"`
	FlavorId       types.String    `tfsdk:"flavor_id,omitempty"`
	ImageId        types.String    `tfsdk:"image_id,omitempty"`
	SecurityGroups []SecurityGroup `tfsdk:"security_groups,omitempty"`
	SSHKey         types.Bool      `tfsdk:"ssh_key,omitempty"`
	KeyName        types.String    `tfsdk:"key_name,omitempty"`
	Count          types.Int64     `tfsdk:"count,omitempty"`
	CreateType     types.String    `tfsdk:"crate_type,omitempty"`
}

// SecurityGroup -
type SecurityGroup struct {
	// ID of security group
	Id string `tfsdk:"id,omitempty"`
	// name of security group
	Name string `tfsdk:"name,omitempty"`
	// description about security group
	Description string `tfsdk:"description,omitempty"`
	// Real name of security group
	RealName string `tfsdk:"real_name,omitempty"`
	// Security group is read only
	ReadOnly bool `tfsdk:"read_only,omitempty"`
	// Instances with this security group
	// Abraks []interface{} `tfsdk:"abraks,omitempty"`
	// rules of security group
	Rules []SecurityGroupRule `tfsdk:"rules,omitempty"`
	// list of security group tags
	// Tags []Tag `tfsdk:"tags,omitempty"`
}

type SecurityGroupRule struct {
	Id string `tfsdk:"id,omitempty"`
	// group id of security group rule
	GroupId string `tfsdk:"group_id,omitempty"`
	// description about security group rule
	Description string `tfsdk:"description,omitempty"`
	// The remote IP prefix to be associated with this security group rule.
	Ip string `tfsdk:"ip,omitempty"`
	// The maximum port number in the range that is matched by the securitygroup rule.
	PortEnd int32 `tfsdk:"port_end,omitempty"`
	// The minimum port number in the range that is matched by the security group rule.
	PortStart int32 `tfsdk:"port_start,omitempty"`
	// The protocol that is matched by the security group rule
	Protocol string `tfsdk:"protocol,omitempty"`
	// The direction in which the security group rule is applied(ingress or egress)
	Direction string `tfsdk:"direction,omitempty"`
	// IPV4 or IPV6
	EtherType string `tfsdk:"ether_type,omitempty"`
	// Rule creation time
	CreatedAt string `tfsdk:"created_at,omitempty"`
	// Rule update time
	UpdatedAt string `tfsdk:"updated_at,omitempty"`
}
