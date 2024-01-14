// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package local_resources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func LocalUserResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_expires": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Define when the local user account expires (UTC). If not specified, the user account never expires.<br>The string time format is the following: `yyyy-MM-dd hh:mm:ss` (see [go time package](https://pkg.go.dev/time#pkg-constants) `DateTime`).",
				MarkdownDescription: "Define when the local user account expires (UTC). If not specified, the user account never expires.<br>The string time format is the following: `yyyy-MM-dd hh:mm:ss` (see [go time package](https://pkg.go.dev/time#pkg-constants) `DateTime`).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Define a description for the local user. The maximum length is 48 characters.",
				MarkdownDescription: "Define a description for the local user. The maximum length is 48 characters.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 48),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "(Default: `true`)<br>Define whether the local user is enabled.",
				MarkdownDescription: "(Default: `true`)<br>Define whether the local user is enabled.",
				Default:             booldefault.StaticBool(true),
			},
			"full_name": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Define the full name of the local user. The full name differs from the user name of the user account.",
				MarkdownDescription: "Define the full name of the local user. The full name differs from the user name of the user account.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed:            true,
				Description:         "The ID of the retrieved local security group. This is the same as the SID.",
				MarkdownDescription: "The ID of the retrieved local security group. This is the same as the SID.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_login": schema.StringAttribute{
				Computed:            true,
				Description:         "The last login time of the local user.",
				MarkdownDescription: "The last login time of the local user.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				Description:         "Define the name for the local user. A user name can contain up to 20 uppercase characters or lowercase characters. A user name can't contain the following characters: `\"`, `/`, `\\`, `[`, `]`, `:`, `;`, `|`, `=`, `,`, `+`, `*`, `?`, `<`, `>`, `@`",
				MarkdownDescription: "Define the name for the local user. A user name can contain up to 20 uppercase characters or lowercase characters. A user name can't contain the following characters: `\"`, `/`, `\\`, `[`, `]`, `:`, `;`, `|`, `=`, `,`, `+`, `*`, `?`, `<`, `>`, `@`",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 20),
				},
			},
			"password": schema.StringAttribute{
				Optional:            true,
				Sensitive:           true,
				Description:         "Define a password for the local user. A password can contain up to 127 characters.",
				MarkdownDescription: "Define a password for the local user. A password can contain up to 127 characters.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 127),
				},
			},
			"password_changeable_date": schema.StringAttribute{
				Computed:            true,
				Description:         "The password changeable date of the local user.",
				MarkdownDescription: "The password changeable date of the local user.",
			},
			"password_expires": schema.StringAttribute{
				Computed:            true,
				Description:         "The time when the password of the local user expires.",
				MarkdownDescription: "The time when the password of the local user expires.",
			},
			"password_last_set": schema.StringAttribute{
				Computed:            true,
				Description:         "The last time when the password was set for the local user.",
				MarkdownDescription: "The last time when the password was set for the local user.",
			},
			"password_never_expires": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "(Default: `true`)<br>Define whether the password of the local user.",
				MarkdownDescription: "(Default: `true`)<br>Define whether the password of the local user.",
				Default:             booldefault.StaticBool(true),
			},
			"password_required": schema.BoolAttribute{
				Computed:            true,
				Description:         "If true a password is required login with the local user.",
				MarkdownDescription: "If true a password is required login with the local user.",
			},
			"sid": schema.StringAttribute{
				Computed:            true,
				Description:         "The security ID of the local user.",
				MarkdownDescription: "The security ID of the local user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"user_may_change_password": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "(Default: `true`)<br>Define whether the local user can change it's own password.",
				MarkdownDescription: "(Default: `true`)<br>Define whether the local user can change it's own password.",
				Default:             booldefault.StaticBool(true),
			},
		},
	}
}

type LocalUserModel struct {
	AccountExpires         types.String `tfsdk:"account_expires"`
	Description            types.String `tfsdk:"description"`
	Enabled                types.Bool   `tfsdk:"enabled"`
	FullName               types.String `tfsdk:"full_name"`
	Id                     types.String `tfsdk:"id"`
	LastLogin              types.String `tfsdk:"last_login"`
	Name                   types.String `tfsdk:"name"`
	Password               types.String `tfsdk:"password"`
	PasswordChangeableDate types.String `tfsdk:"password_changeable_date"`
	PasswordExpires        types.String `tfsdk:"password_expires"`
	PasswordLastSet        types.String `tfsdk:"password_last_set"`
	PasswordNeverExpires   types.Bool   `tfsdk:"password_never_expires"`
	PasswordRequired       types.Bool   `tfsdk:"password_required"`
	Sid                    types.String `tfsdk:"sid"`
	UserMayChangePassword  types.Bool   `tfsdk:"user_may_change_password"`
}
