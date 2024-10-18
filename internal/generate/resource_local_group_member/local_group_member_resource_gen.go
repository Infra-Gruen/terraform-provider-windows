// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package resource_local_group_member

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func LocalGroupMemberResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"group_id": schema.StringAttribute{
				Required:            true,
				Description:         "The ID of the local security group you want to add the member to. Changing this forces a new resource to be created.",
				MarkdownDescription: "The ID of the local security group you want to add the member to. Changing this forces a new resource to be created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The ID of this resource.",
				MarkdownDescription: "The ID of this resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"member_id": schema.StringAttribute{
				Required:            true,
				Description:         "The ID of the principal you want to add as a member to the group. Supported object types are local users or groups. Changing this forces a new resource to be created.",
				MarkdownDescription: "The ID of the principal you want to add as a member to the group. Supported object types are local users or groups. Changing this forces a new resource to be created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
		Description:         "Manage group member for local security groups.",
		MarkdownDescription: "Manage group member for local security groups.",
	}
}

type LocalGroupMemberModel struct {
	GroupId  types.String `tfsdk:"group_id"`
	Id       types.String `tfsdk:"id"`
	MemberId types.String `tfsdk:"member_id"`
}
