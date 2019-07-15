// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta2/field.proto

package admin

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Represents a single field in the database.
//
// Fields are grouped by their "Collection Group", which represent all
// collections in the database with the same id.
type Field struct {
	// A field name of the form
	// `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/fields/{field_path}`
	//
	// A field path may be a simple field name, e.g. `address` or a path to fields
	// within map_value , e.g. `address.city`,
	// or a special field path. The only valid special field is `*`, which
	// represents any field.
	//
	// Field paths may be quoted using ` (backtick). The only character that needs
	// to be escaped within a quoted field path is the backtick character itself,
	// escaped using a backslash. Special characters in field paths that
	// must be quoted include: `*`, `.`,
	// ``` (backtick), `[`, `]`, as well as any ascii symbolic characters.
	//
	// Examples:
	// (Note: Comments here are written in markdown syntax, so there is an
	//  additional layer of backticks to represent a code block)
	// `\`address.city\`` represents a field named `address.city`, not the map key
	// `city` in the field `address`.
	// `\`*\`` represents a field named `*`, not any field.
	//
	// A special `Field` contains the default indexing settings for all fields.
	// This field's resource name is:
	// `projects/{project_id}/databases/{database_id}/collectionGroups/__default__/fields/*`
	// Indexes defined on this `Field` will be applied to all fields which do not
	// have their own `Field` index configuration.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The index configuration for this field. If unset, field indexing will
	// revert to the configuration defined by the `ancestor_field`. To
	// explicitly remove all indexes for this field, specify an index config
	// with an empty list of indexes.
	IndexConfig          *Field_IndexConfig `protobuf:"bytes,2,opt,name=index_config,json=indexConfig,proto3" json:"index_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Field) Reset()         { *m = Field{} }
func (m *Field) String() string { return proto.CompactTextString(m) }
func (*Field) ProtoMessage()    {}
func (*Field) Descriptor() ([]byte, []int) {
	return fileDescriptor_967ea3483ba729a5, []int{0}
}

func (m *Field) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field.Unmarshal(m, b)
}
func (m *Field) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field.Marshal(b, m, deterministic)
}
func (m *Field) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field.Merge(m, src)
}
func (m *Field) XXX_Size() int {
	return xxx_messageInfo_Field.Size(m)
}
func (m *Field) XXX_DiscardUnknown() {
	xxx_messageInfo_Field.DiscardUnknown(m)
}

var xxx_messageInfo_Field proto.InternalMessageInfo

func (m *Field) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Field) GetIndexConfig() *Field_IndexConfig {
	if m != nil {
		return m.IndexConfig
	}
	return nil
}

// The index configuration for this field.
type Field_IndexConfig struct {
	// The indexes supported for this field.
	Indexes []*Index `protobuf:"bytes,1,rep,name=indexes,proto3" json:"indexes,omitempty"`
	// Output only.
	// When true, the `Field`'s index configuration is set from the
	// configuration specified by the `ancestor_field`.
	// When false, the `Field`'s index configuration is defined explicitly.
	UsesAncestorConfig bool `protobuf:"varint,2,opt,name=uses_ancestor_config,json=usesAncestorConfig,proto3" json:"uses_ancestor_config,omitempty"`
	// Output only.
	// Specifies the resource name of the `Field` from which this field's
	// index configuration is set (when `uses_ancestor_config` is true),
	// or from which it *would* be set if this field had no index configuration
	// (when `uses_ancestor_config` is false).
	AncestorField string `protobuf:"bytes,3,opt,name=ancestor_field,json=ancestorField,proto3" json:"ancestor_field,omitempty"`
	// Output only
	// When true, the `Field`'s index configuration is in the process of being
	// reverted. Once complete, the index config will transition to the same
	// state as the field specified by `ancestor_field`, at which point
	// `uses_ancestor_config` will be `true` and `reverting` will be `false`.
	Reverting            bool     `protobuf:"varint,4,opt,name=reverting,proto3" json:"reverting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Field_IndexConfig) Reset()         { *m = Field_IndexConfig{} }
func (m *Field_IndexConfig) String() string { return proto.CompactTextString(m) }
func (*Field_IndexConfig) ProtoMessage()    {}
func (*Field_IndexConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_967ea3483ba729a5, []int{0, 0}
}

func (m *Field_IndexConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Field_IndexConfig.Unmarshal(m, b)
}
func (m *Field_IndexConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Field_IndexConfig.Marshal(b, m, deterministic)
}
func (m *Field_IndexConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Field_IndexConfig.Merge(m, src)
}
func (m *Field_IndexConfig) XXX_Size() int {
	return xxx_messageInfo_Field_IndexConfig.Size(m)
}
func (m *Field_IndexConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Field_IndexConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Field_IndexConfig proto.InternalMessageInfo

func (m *Field_IndexConfig) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *Field_IndexConfig) GetUsesAncestorConfig() bool {
	if m != nil {
		return m.UsesAncestorConfig
	}
	return false
}

func (m *Field_IndexConfig) GetAncestorField() string {
	if m != nil {
		return m.AncestorField
	}
	return ""
}

func (m *Field_IndexConfig) GetReverting() bool {
	if m != nil {
		return m.Reverting
	}
	return false
}

func init() {
	proto.RegisterType((*Field)(nil), "google.firestore.admin.v1beta2.Field")
	proto.RegisterType((*Field_IndexConfig)(nil), "google.firestore.admin.v1beta2.Field.IndexConfig")
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta2/field.proto", fileDescriptor_967ea3483ba729a5)
}

var fileDescriptor_967ea3483ba729a5 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x99, 0xb4, 0xdf, 0xa7, 0x9d, 0xa8, 0x8b, 0xc1, 0x45, 0x28, 0x45, 0x4a, 0xb1, 0x50,
	0x5c, 0xcc, 0xd8, 0xba, 0x74, 0x21, 0x6d, 0xa0, 0xc5, 0x5d, 0x89, 0xe2, 0xc2, 0x4d, 0x99, 0x36,
	0xd3, 0x61, 0x20, 0x9d, 0x5b, 0x92, 0xb4, 0xf8, 0x3c, 0x82, 0x1b, 0xdf, 0xc0, 0x07, 0xf0, 0xbd,
	0x24, 0x77, 0xd2, 0x3f, 0x1b, 0xcd, 0x2e, 0x33, 0xe7, 0x9c, 0xdf, 0x3d, 0xb9, 0x43, 0x6f, 0x34,
	0x80, 0x4e, 0x94, 0x58, 0x9a, 0x54, 0x65, 0x39, 0xa4, 0x4a, 0xc8, 0x78, 0x65, 0xac, 0xd8, 0xf6,
	0xe7, 0x2a, 0x97, 0x03, 0xb1, 0x34, 0x2a, 0x89, 0xf9, 0x3a, 0x85, 0x1c, 0xd8, 0x95, 0xf3, 0xf2,
	0xbd, 0x97, 0xa3, 0x97, 0x97, 0xde, 0x66, 0x15, 0xcb, 0xd8, 0x58, 0xbd, 0x39, 0x56, 0xb3, 0x55,
	0x7a, 0xe5, 0xda, 0x08, 0x69, 0x2d, 0xe4, 0x32, 0x37, 0x60, 0x33, 0xa7, 0x76, 0xbe, 0x3c, 0xfa,
	0x6f, 0x5c, 0x4c, 0x66, 0x8c, 0xd6, 0xad, 0x5c, 0xa9, 0x80, 0xb4, 0x49, 0xaf, 0x11, 0xe1, 0x37,
	0x7b, 0xa6, 0x67, 0x88, 0x9a, 0x2d, 0xc0, 0x2e, 0x8d, 0x0e, 0xbc, 0x36, 0xe9, 0xf9, 0x83, 0x3e,
	0xff, 0xbb, 0x1e, 0x47, 0x20, 0x7f, 0x2c, 0x92, 0x21, 0x06, 0x23, 0xdf, 0x1c, 0x0e, 0xcd, 0x6f,
	0x42, 0xfd, 0x23, 0x91, 0x3d, 0xd0, 0x13, 0x94, 0x55, 0x16, 0x90, 0x76, 0xad, 0xe7, 0x0f, 0xba,
	0x55, 0x03, 0x30, 0x1d, 0xed, 0x52, 0xec, 0x96, 0x5e, 0x6e, 0x32, 0x95, 0xcd, 0xa4, 0x5d, 0xa0,
	0xfd, 0xb8, 0xee, 0x69, 0xc4, 0x0a, 0x6d, 0x58, 0x4a, 0xe5, 0xc8, 0x2e, 0xbd, 0xd8, 0x9b, 0x71,
	0xf1, 0x41, 0x0d, 0x7f, 0xfb, 0x7c, 0x77, 0xeb, 0x76, 0xd2, 0xa2, 0x8d, 0x54, 0x6d, 0x55, 0x9a,
	0x1b, 0xab, 0x83, 0x3a, 0xd2, 0x0e, 0x17, 0xa3, 0x0f, 0x42, 0x3b, 0x0b, 0x58, 0x55, 0x94, 0x1d,
	0x51, 0x64, 0x4d, 0x8b, 0x75, 0x4f, 0xc9, 0x6b, 0x58, 0xba, 0x35, 0x24, 0xd2, 0x6a, 0x0e, 0xa9,
	0x16, 0x5a, 0x59, 0x7c, 0x0c, 0xe1, 0x24, 0xb9, 0x36, 0xd9, 0x6f, 0x2f, 0x7b, 0x8f, 0xa7, 0x77,
	0xaf, 0x3e, 0x09, 0xc7, 0x4f, 0x9f, 0xde, 0xf5, 0xc4, 0xc1, 0xc2, 0x04, 0x36, 0x31, 0x1f, 0xef,
	0x0b, 0x0c, 0xb1, 0xc0, 0x4b, 0x7f, 0x54, 0x64, 0xe6, 0xff, 0x91, 0x7e, 0xf7, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xd4, 0x92, 0xf0, 0x86, 0x82, 0x02, 0x00, 0x00,
}