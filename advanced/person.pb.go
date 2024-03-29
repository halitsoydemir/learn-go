// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

package advanced

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

type SocialFollowers struct {
	Youtube              int32    `protobuf:"varint,1,opt,name=youtube,proto3" json:"youtube,omitempty"`
	Twitter              int32    `protobuf:"varint,2,opt,name=twitter,proto3" json:"twitter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SocialFollowers) Reset()         { *m = SocialFollowers{} }
func (m *SocialFollowers) String() string { return proto.CompactTextString(m) }
func (*SocialFollowers) ProtoMessage()    {}
func (*SocialFollowers) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}

func (m *SocialFollowers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocialFollowers.Unmarshal(m, b)
}
func (m *SocialFollowers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocialFollowers.Marshal(b, m, deterministic)
}
func (m *SocialFollowers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocialFollowers.Merge(m, src)
}
func (m *SocialFollowers) XXX_Size() int {
	return xxx_messageInfo_SocialFollowers.Size(m)
}
func (m *SocialFollowers) XXX_DiscardUnknown() {
	xxx_messageInfo_SocialFollowers.DiscardUnknown(m)
}

var xxx_messageInfo_SocialFollowers proto.InternalMessageInfo

func (m *SocialFollowers) GetYoutube() int32 {
	if m != nil {
		return m.Youtube
	}
	return 0
}

func (m *SocialFollowers) GetTwitter() int32 {
	if m != nil {
		return m.Twitter
	}
	return 0
}

type Person struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32            `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	SocialFollowers      *SocialFollowers `protobuf:"bytes,3,opt,name=socialFollowers,proto3" json:"socialFollowers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetSocialFollowers() *SocialFollowers {
	if m != nil {
		return m.SocialFollowers
	}
	return nil
}

func init() {
	proto.RegisterType((*SocialFollowers)(nil), "advanced.SocialFollowers")
	proto.RegisterType((*Person)(nil), "advanced.Person")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x4c, 0x29, 0x4b, 0xcc, 0x4b,
	0x4e, 0x4d, 0x51, 0x72, 0xe5, 0xe2, 0x0f, 0xce, 0x4f, 0xce, 0x4c, 0xcc, 0x71, 0xcb, 0xcf, 0xc9,
	0xc9, 0x2f, 0x4f, 0x2d, 0x2a, 0x16, 0x92, 0xe0, 0x62, 0xaf, 0xcc, 0x2f, 0x2d, 0x29, 0x4d, 0x4a,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0x71, 0x41, 0x32, 0x25, 0xe5, 0x99, 0x25, 0x25,
	0xa9, 0x45, 0x12, 0x4c, 0x10, 0x19, 0x28, 0x57, 0xa9, 0x98, 0x8b, 0x2d, 0x00, 0x6c, 0x81, 0x90,
	0x10, 0x17, 0x4b, 0x5e, 0x62, 0x2e, 0x44, 0x2b, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc0, 0xc5, 0x9c,
	0x98, 0x9e, 0x0a, 0xd5, 0x03, 0x62, 0x0a, 0x39, 0x73, 0xf1, 0x17, 0xa3, 0x5a, 0x2b, 0xc1, 0xac,
	0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x07, 0x73, 0x9a, 0x1e, 0x9a, 0xbb, 0x82, 0xd0, 0x75, 0x24,
	0xb1, 0x81, 0x3d, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xdf, 0x57, 0x7c, 0xdc, 0x00,
	0x00, 0x00,
}
