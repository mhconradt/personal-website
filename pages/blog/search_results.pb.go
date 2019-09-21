// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search_results.proto

package blog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ArticleSnippet struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Timestamp            int32    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArticleSnippet) Reset()         { *m = ArticleSnippet{} }
func (m *ArticleSnippet) String() string { return proto.CompactTextString(m) }
func (*ArticleSnippet) ProtoMessage()    {}
func (*ArticleSnippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba8301929eaa4346, []int{0}
}

func (m *ArticleSnippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticleSnippet.Unmarshal(m, b)
}
func (m *ArticleSnippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticleSnippet.Marshal(b, m, deterministic)
}
func (m *ArticleSnippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticleSnippet.Merge(m, src)
}
func (m *ArticleSnippet) XXX_Size() int {
	return xxx_messageInfo_ArticleSnippet.Size(m)
}
func (m *ArticleSnippet) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticleSnippet.DiscardUnknown(m)
}

var xxx_messageInfo_ArticleSnippet proto.InternalMessageInfo

func (m *ArticleSnippet) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ArticleSnippet) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ArticleSnippet) GetTimestamp() int32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *ArticleSnippet) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Cursor struct {
	Forward              int32    `protobuf:"varint,1,opt,name=forward,proto3" json:"forward,omitempty"`
	Reverse              int32    `protobuf:"varint,2,opt,name=reverse,proto3" json:"reverse,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Term                 string   `protobuf:"bytes,4,opt,name=term,proto3" json:"term,omitempty"`
	Index                string   `protobuf:"bytes,5,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cursor) Reset()         { *m = Cursor{} }
func (m *Cursor) String() string { return proto.CompactTextString(m) }
func (*Cursor) ProtoMessage()    {}
func (*Cursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba8301929eaa4346, []int{1}
}

func (m *Cursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cursor.Unmarshal(m, b)
}
func (m *Cursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cursor.Marshal(b, m, deterministic)
}
func (m *Cursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cursor.Merge(m, src)
}
func (m *Cursor) XXX_Size() int {
	return xxx_messageInfo_Cursor.Size(m)
}
func (m *Cursor) XXX_DiscardUnknown() {
	xxx_messageInfo_Cursor.DiscardUnknown(m)
}

var xxx_messageInfo_Cursor proto.InternalMessageInfo

func (m *Cursor) GetForward() int32 {
	if m != nil {
		return m.Forward
	}
	return 0
}

func (m *Cursor) GetReverse() int32 {
	if m != nil {
		return m.Reverse
	}
	return 0
}

func (m *Cursor) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Cursor) GetTerm() string {
	if m != nil {
		return m.Term
	}
	return ""
}

func (m *Cursor) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type SearchResults struct {
	Results              []*ArticleSnippet `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Cursor               *Cursor           `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SearchResults) Reset()         { *m = SearchResults{} }
func (m *SearchResults) String() string { return proto.CompactTextString(m) }
func (*SearchResults) ProtoMessage()    {}
func (*SearchResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba8301929eaa4346, []int{2}
}

func (m *SearchResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResults.Unmarshal(m, b)
}
func (m *SearchResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResults.Marshal(b, m, deterministic)
}
func (m *SearchResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResults.Merge(m, src)
}
func (m *SearchResults) XXX_Size() int {
	return xxx_messageInfo_SearchResults.Size(m)
}
func (m *SearchResults) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResults.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResults proto.InternalMessageInfo

func (m *SearchResults) GetResults() []*ArticleSnippet {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *SearchResults) GetCursor() *Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func init() {
	proto.RegisterType((*ArticleSnippet)(nil), "blog.ArticleSnippet")
	proto.RegisterType((*Cursor)(nil), "blog.Cursor")
	proto.RegisterType((*SearchResults)(nil), "blog.SearchResults")
}

func init() { proto.RegisterFile("search_results.proto", fileDescriptor_ba8301929eaa4346) }

var fileDescriptor_ba8301929eaa4346 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x6e, 0x83, 0x30,
	0x10, 0x84, 0x05, 0x01, 0xaa, 0x6c, 0xda, 0x1c, 0x2c, 0x0e, 0x3e, 0xf4, 0x80, 0x50, 0x0f, 0x9c,
	0x38, 0xa4, 0x4f, 0x50, 0xf5, 0x0d, 0x9c, 0x07, 0xa8, 0xf8, 0xd9, 0x36, 0x96, 0x00, 0xa3, 0xf5,
	0xd2, 0x1f, 0xa9, 0x0f, 0x5f, 0xd9, 0x8e, 0x55, 0xe5, 0x36, 0xc3, 0x0e, 0xfa, 0xc6, 0x03, 0xa5,
	0xc5, 0x8e, 0x86, 0xcb, 0x1b, 0xa1, 0xdd, 0x26, 0xb6, 0xed, 0x4a, 0x86, 0x8d, 0xc8, 0xfa, 0xc9,
	0x7c, 0xd4, 0x17, 0x38, 0xbe, 0x10, 0xeb, 0x61, 0xc2, 0xf3, 0xa2, 0xd7, 0x15, 0x59, 0x1c, 0x21,
	0xd5, 0xa3, 0x4c, 0xaa, 0xa4, 0xc9, 0x55, 0xaa, 0x47, 0x51, 0x42, 0xce, 0x9a, 0x27, 0x94, 0x69,
	0x95, 0x34, 0x7b, 0x15, 0x8c, 0x78, 0x84, 0x3d, 0xeb, 0x19, 0x2d, 0x77, 0xf3, 0x2a, 0x77, 0x3e,
	0xfc, 0xff, 0x41, 0x08, 0xc8, 0x7a, 0x33, 0xfe, 0xc8, 0xcc, 0xff, 0xe2, 0x75, 0xfd, 0x0b, 0xc5,
	0xeb, 0x46, 0xd6, 0x90, 0x90, 0x70, 0xf7, 0x6e, 0xe8, 0xab, 0xa3, 0x88, 0x89, 0xd6, 0x5d, 0x08,
	0x3f, 0x91, 0x6c, 0xa0, 0xe5, 0x2a, 0x5a, 0xd7, 0x62, 0x30, 0xdb, 0xc2, 0x57, 0x56, 0x30, 0x8e,
	0xc3, 0x48, 0x73, 0xe4, 0x38, 0xed, 0x92, 0x7a, 0x19, 0xf1, 0x5b, 0xe6, 0xa1, 0xaf, 0x37, 0x35,
	0xc2, 0xc3, 0xd9, 0xaf, 0xa0, 0xc2, 0x08, 0xa2, 0x75, 0x28, 0x2f, 0x65, 0x52, 0xed, 0x9a, 0xc3,
	0xa9, 0x6c, 0xdd, 0x20, 0xed, 0xed, 0x1a, 0x2a, 0x86, 0xc4, 0x13, 0x14, 0x83, 0xaf, 0xef, 0x9b,
	0x1d, 0x4e, 0xf7, 0x21, 0x1e, 0x9e, 0xa4, 0xae, 0xb7, 0xbe, 0xf0, 0xdb, 0x3e, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x56, 0x9d, 0x29, 0x61, 0x73, 0x01, 0x00, 0x00,
}