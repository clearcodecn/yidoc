package swaggos

import (
	"reflect"

	"github.com/go-openapi/spec"
)

type documentTag struct {
	tag       reflect.StructTag
	attribute *Attribute
}

func newTags(tag reflect.StructTag) *documentTag {
	dt := new(documentTag)
	dt.tag = tag
	a := &Attribute{}
	a.parseTag(tag)
	dt.attribute = a
	return dt
}

func (t *documentTag) name() string {
	if t.attribute.Model != "" {
		return t.attribute.Model
	}
	if t.attribute.JSON != "" {
		return t.attribute.JSON
	}
	return ""
}

func (t *documentTag) ignore() bool {
	return t.attribute.Ignore
}

func (t *documentTag) jsonTag() string {
	return t.tag.Get("json")
}

func (t *documentTag) required() bool {
	return t.attribute.Required
}

func (t *documentTag) Attribute() *Attribute {
	return t.attribute
}

func (t *documentTag) jsonName() string {
	return t.attribute.JSON
}

func (t *documentTag) mergeSchema(schema spec.Schema) spec.Schema {
	schema.Description = t.attribute.Description
	schema.Example = t.attribute.Example
	schema.Nullable = t.attribute.Nullable
	schema.Format = t.attribute.Format
	schema.Title = t.attribute.Title
	schema.MaxLength = t.attribute.MaxLength
	schema.MinLength = t.attribute.MinLength
	schema.Pattern = t.attribute.Pattern
	schema.Maximum = t.attribute.Maximum
	schema.Minimum = t.attribute.Minimum
	schema.MaxItems = t.attribute.MaxItems
	schema.MinItems = t.attribute.MinItems
	schema.Enum = t.attribute.Enum
	schema.Default = t.attribute.Default
	return schema
}
