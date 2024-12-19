// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_sshkey

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func SshkeyResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"created_at": schema.Int64Attribute{
				Required:            true,
				Description:         "SSH key createdAt",
				MarkdownDescription: "SSH key createdAt",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "SSH key name",
				MarkdownDescription: "SSH key name",
			},
			"public_key": schema.StringAttribute{
				Required:            true,
				Description:         "Public SSH key contents",
				MarkdownDescription: "Public SSH key contents",
			},
			"uuid": schema.StringAttribute{
				Required:            true,
				Description:         "SSH key UUID as specified in RFC 4122",
				MarkdownDescription: "SSH key UUID as specified in RFC 4122",
			},
		},
	}
}

type SshkeyModel struct {
	CreatedAt types.Int64  `tfsdk:"created_at"`
	Name      types.String `tfsdk:"name"`
	PublicKey types.String `tfsdk:"public_key"`
	Uuid      types.String `tfsdk:"uuid"`
}
