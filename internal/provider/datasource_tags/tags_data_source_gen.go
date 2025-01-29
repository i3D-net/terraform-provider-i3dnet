// Code generated by terraform-plugin-framework-generator DO NOT EDIT.

package datasource_tags

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func TagsDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"tags": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"resources": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"count": schema.Int64Attribute{
									Computed:            true,
									Description:         "The total number of resources that use this tag",
									MarkdownDescription: "The total number of resources that use this tag",
								},
								"flex_metal_servers": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"count": schema.Int64Attribute{
											Computed:            true,
											Description:         "The amount of resources of this type that use this tag.",
											MarkdownDescription: "The amount of resources of this type that use this tag.",
										},
									},
									CustomType: FlexMetalServersType{
										ObjectType: types.ObjectType{
											AttrTypes: FlexMetalServersValue{}.AttributeTypes(ctx),
										},
									},
									Computed:            true,
									Description:         "A summary of tag usage for the FlexMetalServer resource type",
									MarkdownDescription: "A summary of tag usage for the FlexMetalServer resource type",
								},
							},
							CustomType: ResourcesType{
								ObjectType: types.ObjectType{
									AttrTypes: ResourcesValue{}.AttributeTypes(ctx),
								},
							},
							Computed:            true,
							Description:         "A list of resources that use this tag",
							MarkdownDescription: "A list of resources that use this tag",
						},
						"tag": schema.StringAttribute{
							Computed:            true,
							Description:         "The tag name",
							MarkdownDescription: "The tag name",
						},
					},
					CustomType: TagsType{
						ObjectType: types.ObjectType{
							AttrTypes: TagsValue{}.AttributeTypes(ctx),
						},
					},
				},
				Computed: true,
			},
		},
	}
}

type TagsModel struct {
	Tags types.Set `tfsdk:"tags"`
}

var _ basetypes.ObjectTypable = TagsType{}

type TagsType struct {
	basetypes.ObjectType
}

func (t TagsType) Equal(o attr.Type) bool {
	other, ok := o.(TagsType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t TagsType) String() string {
	return "TagsType"
}

func (t TagsType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	resourcesAttribute, ok := attributes["resources"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`resources is missing from object`)

		return nil, diags
	}

	resourcesVal, ok := resourcesAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`resources expected to be basetypes.ObjectValue, was: %T`, resourcesAttribute))
	}

	tagAttribute, ok := attributes["tag"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tag is missing from object`)

		return nil, diags
	}

	tagVal, ok := tagAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tag expected to be basetypes.StringValue, was: %T`, tagAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return TagsValue{
		Resources: resourcesVal,
		Tag:       tagVal,
		state:     attr.ValueStateKnown,
	}, diags
}

func NewTagsValueNull() TagsValue {
	return TagsValue{
		state: attr.ValueStateNull,
	}
}

func NewTagsValueUnknown() TagsValue {
	return TagsValue{
		state: attr.ValueStateUnknown,
	}
}

func NewTagsValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (TagsValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing TagsValue Attribute Value",
				"While creating a TagsValue value, a missing attribute value was detected. "+
					"A TagsValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("TagsValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid TagsValue Attribute Type",
				"While creating a TagsValue value, an invalid attribute value was detected. "+
					"A TagsValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("TagsValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("TagsValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra TagsValue Attribute Value",
				"While creating a TagsValue value, an extra attribute value was detected. "+
					"A TagsValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra TagsValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewTagsValueUnknown(), diags
	}

	resourcesAttribute, ok := attributes["resources"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`resources is missing from object`)

		return NewTagsValueUnknown(), diags
	}

	resourcesVal, ok := resourcesAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`resources expected to be basetypes.ObjectValue, was: %T`, resourcesAttribute))
	}

	tagAttribute, ok := attributes["tag"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`tag is missing from object`)

		return NewTagsValueUnknown(), diags
	}

	tagVal, ok := tagAttribute.(basetypes.StringValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`tag expected to be basetypes.StringValue, was: %T`, tagAttribute))
	}

	if diags.HasError() {
		return NewTagsValueUnknown(), diags
	}

	return TagsValue{
		Resources: resourcesVal,
		Tag:       tagVal,
		state:     attr.ValueStateKnown,
	}, diags
}

func NewTagsValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) TagsValue {
	object, diags := NewTagsValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewTagsValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t TagsType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewTagsValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewTagsValueUnknown(), nil
	}

	if in.IsNull() {
		return NewTagsValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewTagsValueMust(TagsValue{}.AttributeTypes(ctx), attributes), nil
}

func (t TagsType) ValueType(ctx context.Context) attr.Value {
	return TagsValue{}
}

var _ basetypes.ObjectValuable = TagsValue{}

type TagsValue struct {
	Resources basetypes.ObjectValue `tfsdk:"resources"`
	Tag       basetypes.StringValue `tfsdk:"tag"`
	state     attr.ValueState
}

func (v TagsValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["resources"] = basetypes.ObjectType{
		AttrTypes: ResourcesValue{}.AttributeTypes(ctx),
	}.TerraformType(ctx)
	attrTypes["tag"] = basetypes.StringType{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Resources.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["resources"] = val

		val, err = v.Tag.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["tag"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v TagsValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v TagsValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v TagsValue) String() string {
	return "TagsValue"
}

func (v TagsValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var resources basetypes.ObjectValue

	if v.Resources.IsNull() {
		resources = types.ObjectNull(
			ResourcesValue{}.AttributeTypes(ctx),
		)
	}

	if v.Resources.IsUnknown() {
		resources = types.ObjectUnknown(
			ResourcesValue{}.AttributeTypes(ctx),
		)
	}

	if !v.Resources.IsNull() && !v.Resources.IsUnknown() {
		resources = types.ObjectValueMust(
			ResourcesValue{}.AttributeTypes(ctx),
			v.Resources.Attributes(),
		)
	}

	attributeTypes := map[string]attr.Type{
		"resources": basetypes.ObjectType{
			AttrTypes: ResourcesValue{}.AttributeTypes(ctx),
		},
		"tag": basetypes.StringType{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"resources": resources,
			"tag":       v.Tag,
		})

	return objVal, diags
}

func (v TagsValue) Equal(o attr.Value) bool {
	other, ok := o.(TagsValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Resources.Equal(other.Resources) {
		return false
	}

	if !v.Tag.Equal(other.Tag) {
		return false
	}

	return true
}

func (v TagsValue) Type(ctx context.Context) attr.Type {
	return TagsType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v TagsValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"resources": basetypes.ObjectType{
			AttrTypes: ResourcesValue{}.AttributeTypes(ctx),
		},
		"tag": basetypes.StringType{},
	}
}

var _ basetypes.ObjectTypable = ResourcesType{}

type ResourcesType struct {
	basetypes.ObjectType
}

func (t ResourcesType) Equal(o attr.Type) bool {
	other, ok := o.(ResourcesType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t ResourcesType) String() string {
	return "ResourcesType"
}

func (t ResourcesType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	countAttribute, ok := attributes["count"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`count is missing from object`)

		return nil, diags
	}

	countVal, ok := countAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`count expected to be basetypes.Int64Value, was: %T`, countAttribute))
	}

	flexMetalServersAttribute, ok := attributes["flex_metal_servers"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`flex_metal_servers is missing from object`)

		return nil, diags
	}

	flexMetalServersVal, ok := flexMetalServersAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`flex_metal_servers expected to be basetypes.ObjectValue, was: %T`, flexMetalServersAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return ResourcesValue{
		Count:            countVal,
		FlexMetalServers: flexMetalServersVal,
		state:            attr.ValueStateKnown,
	}, diags
}

func NewResourcesValueNull() ResourcesValue {
	return ResourcesValue{
		state: attr.ValueStateNull,
	}
}

func NewResourcesValueUnknown() ResourcesValue {
	return ResourcesValue{
		state: attr.ValueStateUnknown,
	}
}

func NewResourcesValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (ResourcesValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing ResourcesValue Attribute Value",
				"While creating a ResourcesValue value, a missing attribute value was detected. "+
					"A ResourcesValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ResourcesValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid ResourcesValue Attribute Type",
				"While creating a ResourcesValue value, an invalid attribute value was detected. "+
					"A ResourcesValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("ResourcesValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("ResourcesValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra ResourcesValue Attribute Value",
				"While creating a ResourcesValue value, an extra attribute value was detected. "+
					"A ResourcesValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra ResourcesValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewResourcesValueUnknown(), diags
	}

	countAttribute, ok := attributes["count"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`count is missing from object`)

		return NewResourcesValueUnknown(), diags
	}

	countVal, ok := countAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`count expected to be basetypes.Int64Value, was: %T`, countAttribute))
	}

	flexMetalServersAttribute, ok := attributes["flex_metal_servers"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`flex_metal_servers is missing from object`)

		return NewResourcesValueUnknown(), diags
	}

	flexMetalServersVal, ok := flexMetalServersAttribute.(basetypes.ObjectValue)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`flex_metal_servers expected to be basetypes.ObjectValue, was: %T`, flexMetalServersAttribute))
	}

	if diags.HasError() {
		return NewResourcesValueUnknown(), diags
	}

	return ResourcesValue{
		Count:            countVal,
		FlexMetalServers: flexMetalServersVal,
		state:            attr.ValueStateKnown,
	}, diags
}

func NewResourcesValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) ResourcesValue {
	object, diags := NewResourcesValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewResourcesValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t ResourcesType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewResourcesValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewResourcesValueUnknown(), nil
	}

	if in.IsNull() {
		return NewResourcesValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewResourcesValueMust(ResourcesValue{}.AttributeTypes(ctx), attributes), nil
}

func (t ResourcesType) ValueType(ctx context.Context) attr.Value {
	return ResourcesValue{}
}

var _ basetypes.ObjectValuable = ResourcesValue{}

type ResourcesValue struct {
	Count            basetypes.Int64Value  `tfsdk:"count"`
	FlexMetalServers basetypes.ObjectValue `tfsdk:"flex_metal_servers"`
	state            attr.ValueState
}

func (v ResourcesValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 2)

	var val tftypes.Value
	var err error

	attrTypes["count"] = basetypes.Int64Type{}.TerraformType(ctx)
	attrTypes["flex_metal_servers"] = basetypes.ObjectType{
		AttrTypes: FlexMetalServersValue{}.AttributeTypes(ctx),
	}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 2)

		val, err = v.Count.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["count"] = val

		val, err = v.FlexMetalServers.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["flex_metal_servers"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v ResourcesValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v ResourcesValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v ResourcesValue) String() string {
	return "ResourcesValue"
}

func (v ResourcesValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	var flexMetalServers basetypes.ObjectValue

	if v.FlexMetalServers.IsNull() {
		flexMetalServers = types.ObjectNull(
			FlexMetalServersValue{}.AttributeTypes(ctx),
		)
	}

	if v.FlexMetalServers.IsUnknown() {
		flexMetalServers = types.ObjectUnknown(
			FlexMetalServersValue{}.AttributeTypes(ctx),
		)
	}

	if !v.FlexMetalServers.IsNull() && !v.FlexMetalServers.IsUnknown() {
		flexMetalServers = types.ObjectValueMust(
			FlexMetalServersValue{}.AttributeTypes(ctx),
			v.FlexMetalServers.Attributes(),
		)
	}

	attributeTypes := map[string]attr.Type{
		"count": basetypes.Int64Type{},
		"flex_metal_servers": basetypes.ObjectType{
			AttrTypes: FlexMetalServersValue{}.AttributeTypes(ctx),
		},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"count":              v.Count,
			"flex_metal_servers": flexMetalServers,
		})

	return objVal, diags
}

func (v ResourcesValue) Equal(o attr.Value) bool {
	other, ok := o.(ResourcesValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Count.Equal(other.Count) {
		return false
	}

	if !v.FlexMetalServers.Equal(other.FlexMetalServers) {
		return false
	}

	return true
}

func (v ResourcesValue) Type(ctx context.Context) attr.Type {
	return ResourcesType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v ResourcesValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"count": basetypes.Int64Type{},
		"flex_metal_servers": basetypes.ObjectType{
			AttrTypes: FlexMetalServersValue{}.AttributeTypes(ctx),
		},
	}
}

var _ basetypes.ObjectTypable = FlexMetalServersType{}

type FlexMetalServersType struct {
	basetypes.ObjectType
}

func (t FlexMetalServersType) Equal(o attr.Type) bool {
	other, ok := o.(FlexMetalServersType)

	if !ok {
		return false
	}

	return t.ObjectType.Equal(other.ObjectType)
}

func (t FlexMetalServersType) String() string {
	return "FlexMetalServersType"
}

func (t FlexMetalServersType) ValueFromObject(ctx context.Context, in basetypes.ObjectValue) (basetypes.ObjectValuable, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributes := in.Attributes()

	countAttribute, ok := attributes["count"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`count is missing from object`)

		return nil, diags
	}

	countVal, ok := countAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`count expected to be basetypes.Int64Value, was: %T`, countAttribute))
	}

	if diags.HasError() {
		return nil, diags
	}

	return FlexMetalServersValue{
		Count: countVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewFlexMetalServersValueNull() FlexMetalServersValue {
	return FlexMetalServersValue{
		state: attr.ValueStateNull,
	}
}

func NewFlexMetalServersValueUnknown() FlexMetalServersValue {
	return FlexMetalServersValue{
		state: attr.ValueStateUnknown,
	}
}

func NewFlexMetalServersValue(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) (FlexMetalServersValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Reference: https://github.com/hashicorp/terraform-plugin-framework/issues/521
	ctx := context.Background()

	for name, attributeType := range attributeTypes {
		attribute, ok := attributes[name]

		if !ok {
			diags.AddError(
				"Missing FlexMetalServersValue Attribute Value",
				"While creating a FlexMetalServersValue value, a missing attribute value was detected. "+
					"A FlexMetalServersValue must contain values for all attributes, even if null or unknown. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("FlexMetalServersValue Attribute Name (%s) Expected Type: %s", name, attributeType.String()),
			)

			continue
		}

		if !attributeType.Equal(attribute.Type(ctx)) {
			diags.AddError(
				"Invalid FlexMetalServersValue Attribute Type",
				"While creating a FlexMetalServersValue value, an invalid attribute value was detected. "+
					"A FlexMetalServersValue must use a matching attribute type for the value. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("FlexMetalServersValue Attribute Name (%s) Expected Type: %s\n", name, attributeType.String())+
					fmt.Sprintf("FlexMetalServersValue Attribute Name (%s) Given Type: %s", name, attribute.Type(ctx)),
			)
		}
	}

	for name := range attributes {
		_, ok := attributeTypes[name]

		if !ok {
			diags.AddError(
				"Extra FlexMetalServersValue Attribute Value",
				"While creating a FlexMetalServersValue value, an extra attribute value was detected. "+
					"A FlexMetalServersValue must not contain values beyond the expected attribute types. "+
					"This is always an issue with the provider and should be reported to the provider developers.\n\n"+
					fmt.Sprintf("Extra FlexMetalServersValue Attribute Name: %s", name),
			)
		}
	}

	if diags.HasError() {
		return NewFlexMetalServersValueUnknown(), diags
	}

	countAttribute, ok := attributes["count"]

	if !ok {
		diags.AddError(
			"Attribute Missing",
			`count is missing from object`)

		return NewFlexMetalServersValueUnknown(), diags
	}

	countVal, ok := countAttribute.(basetypes.Int64Value)

	if !ok {
		diags.AddError(
			"Attribute Wrong Type",
			fmt.Sprintf(`count expected to be basetypes.Int64Value, was: %T`, countAttribute))
	}

	if diags.HasError() {
		return NewFlexMetalServersValueUnknown(), diags
	}

	return FlexMetalServersValue{
		Count: countVal,
		state: attr.ValueStateKnown,
	}, diags
}

func NewFlexMetalServersValueMust(attributeTypes map[string]attr.Type, attributes map[string]attr.Value) FlexMetalServersValue {
	object, diags := NewFlexMetalServersValue(attributeTypes, attributes)

	if diags.HasError() {
		// This could potentially be added to the diag package.
		diagsStrings := make([]string, 0, len(diags))

		for _, diagnostic := range diags {
			diagsStrings = append(diagsStrings, fmt.Sprintf(
				"%s | %s | %s",
				diagnostic.Severity(),
				diagnostic.Summary(),
				diagnostic.Detail()))
		}

		panic("NewFlexMetalServersValueMust received error(s): " + strings.Join(diagsStrings, "\n"))
	}

	return object
}

func (t FlexMetalServersType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	if in.Type() == nil {
		return NewFlexMetalServersValueNull(), nil
	}

	if !in.Type().Equal(t.TerraformType(ctx)) {
		return nil, fmt.Errorf("expected %s, got %s", t.TerraformType(ctx), in.Type())
	}

	if !in.IsKnown() {
		return NewFlexMetalServersValueUnknown(), nil
	}

	if in.IsNull() {
		return NewFlexMetalServersValueNull(), nil
	}

	attributes := map[string]attr.Value{}

	val := map[string]tftypes.Value{}

	err := in.As(&val)

	if err != nil {
		return nil, err
	}

	for k, v := range val {
		a, err := t.AttrTypes[k].ValueFromTerraform(ctx, v)

		if err != nil {
			return nil, err
		}

		attributes[k] = a
	}

	return NewFlexMetalServersValueMust(FlexMetalServersValue{}.AttributeTypes(ctx), attributes), nil
}

func (t FlexMetalServersType) ValueType(ctx context.Context) attr.Value {
	return FlexMetalServersValue{}
}

var _ basetypes.ObjectValuable = FlexMetalServersValue{}

type FlexMetalServersValue struct {
	Count basetypes.Int64Value `tfsdk:"count"`
	state attr.ValueState
}

func (v FlexMetalServersValue) ToTerraformValue(ctx context.Context) (tftypes.Value, error) {
	attrTypes := make(map[string]tftypes.Type, 1)

	var val tftypes.Value
	var err error

	attrTypes["count"] = basetypes.Int64Type{}.TerraformType(ctx)

	objectType := tftypes.Object{AttributeTypes: attrTypes}

	switch v.state {
	case attr.ValueStateKnown:
		vals := make(map[string]tftypes.Value, 1)

		val, err = v.Count.ToTerraformValue(ctx)

		if err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		vals["count"] = val

		if err := tftypes.ValidateValue(objectType, vals); err != nil {
			return tftypes.NewValue(objectType, tftypes.UnknownValue), err
		}

		return tftypes.NewValue(objectType, vals), nil
	case attr.ValueStateNull:
		return tftypes.NewValue(objectType, nil), nil
	case attr.ValueStateUnknown:
		return tftypes.NewValue(objectType, tftypes.UnknownValue), nil
	default:
		panic(fmt.Sprintf("unhandled Object state in ToTerraformValue: %s", v.state))
	}
}

func (v FlexMetalServersValue) IsNull() bool {
	return v.state == attr.ValueStateNull
}

func (v FlexMetalServersValue) IsUnknown() bool {
	return v.state == attr.ValueStateUnknown
}

func (v FlexMetalServersValue) String() string {
	return "FlexMetalServersValue"
}

func (v FlexMetalServersValue) ToObjectValue(ctx context.Context) (basetypes.ObjectValue, diag.Diagnostics) {
	var diags diag.Diagnostics

	attributeTypes := map[string]attr.Type{
		"count": basetypes.Int64Type{},
	}

	if v.IsNull() {
		return types.ObjectNull(attributeTypes), diags
	}

	if v.IsUnknown() {
		return types.ObjectUnknown(attributeTypes), diags
	}

	objVal, diags := types.ObjectValue(
		attributeTypes,
		map[string]attr.Value{
			"count": v.Count,
		})

	return objVal, diags
}

func (v FlexMetalServersValue) Equal(o attr.Value) bool {
	other, ok := o.(FlexMetalServersValue)

	if !ok {
		return false
	}

	if v.state != other.state {
		return false
	}

	if v.state != attr.ValueStateKnown {
		return true
	}

	if !v.Count.Equal(other.Count) {
		return false
	}

	return true
}

func (v FlexMetalServersValue) Type(ctx context.Context) attr.Type {
	return FlexMetalServersType{
		basetypes.ObjectType{
			AttrTypes: v.AttributeTypes(ctx),
		},
	}
}

func (v FlexMetalServersValue) AttributeTypes(ctx context.Context) map[string]attr.Type {
	return map[string]attr.Type{
		"count": basetypes.Int64Type{},
	}
}
