// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"github.com/oodle-ai/opentelemetry-collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetMapResourceAttr sets provided value as "map.resource.attr" attribute.
func (rb *ResourceBuilder) SetMapResourceAttr(val map[string]any) {
	if rb.config.MapResourceAttr.Enabled {
		rb.res.Attributes().PutEmptyMap("map.resource.attr").FromRaw(val)
	}
}

// SetOptionalResourceAttr sets provided value as "optional.resource.attr" attribute.
func (rb *ResourceBuilder) SetOptionalResourceAttr(val string) {
	if rb.config.OptionalResourceAttr.Enabled {
		rb.res.Attributes().PutStr("optional.resource.attr", val)
	}
}

// SetSliceResourceAttr sets provided value as "slice.resource.attr" attribute.
func (rb *ResourceBuilder) SetSliceResourceAttr(val []any) {
	if rb.config.SliceResourceAttr.Enabled {
		rb.res.Attributes().PutEmptySlice("slice.resource.attr").FromRaw(val)
	}
}

// SetStringEnumResourceAttrOne sets "string.enum.resource.attr=one" attribute.
func (rb *ResourceBuilder) SetStringEnumResourceAttrOne() {
	if rb.config.StringEnumResourceAttr.Enabled {
		rb.res.Attributes().PutStr("string.enum.resource.attr", "one")
	}
}

// SetStringEnumResourceAttrTwo sets "string.enum.resource.attr=two" attribute.
func (rb *ResourceBuilder) SetStringEnumResourceAttrTwo() {
	if rb.config.StringEnumResourceAttr.Enabled {
		rb.res.Attributes().PutStr("string.enum.resource.attr", "two")
	}
}

// SetStringResourceAttr sets provided value as "string.resource.attr" attribute.
func (rb *ResourceBuilder) SetStringResourceAttr(val string) {
	if rb.config.StringResourceAttr.Enabled {
		rb.res.Attributes().PutStr("string.resource.attr", val)
	}
}

// SetStringResourceAttrDisableWarning sets provided value as "string.resource.attr_disable_warning" attribute.
func (rb *ResourceBuilder) SetStringResourceAttrDisableWarning(val string) {
	if rb.config.StringResourceAttrDisableWarning.Enabled {
		rb.res.Attributes().PutStr("string.resource.attr_disable_warning", val)
	}
}

// SetStringResourceAttrRemoveWarning sets provided value as "string.resource.attr_remove_warning" attribute.
func (rb *ResourceBuilder) SetStringResourceAttrRemoveWarning(val string) {
	if rb.config.StringResourceAttrRemoveWarning.Enabled {
		rb.res.Attributes().PutStr("string.resource.attr_remove_warning", val)
	}
}

// SetStringResourceAttrToBeRemoved sets provided value as "string.resource.attr_to_be_removed" attribute.
func (rb *ResourceBuilder) SetStringResourceAttrToBeRemoved(val string) {
	if rb.config.StringResourceAttrToBeRemoved.Enabled {
		rb.res.Attributes().PutStr("string.resource.attr_to_be_removed", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
