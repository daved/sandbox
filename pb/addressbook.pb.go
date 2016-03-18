// Code generated by protoc-gen-go.
// source: addressbook.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=pbaddbook.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "pbaddbook.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "pbaddbook.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "pbaddbook.AddressBook")
	proto.RegisterEnum("pbaddbook.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xd1, 0x4b, 0xc3, 0x30,
	0x10, 0xc6, 0x6d, 0x97, 0x05, 0x7b, 0x85, 0xd1, 0x1d, 0x22, 0x45, 0x11, 0x46, 0x9f, 0x26, 0x42,
	0x85, 0x89, 0xe0, 0xab, 0x85, 0x81, 0xa2, 0xb3, 0x23, 0x88, 0x82, 0x6f, 0x0d, 0x3d, 0x70, 0xb8,
	0x36, 0xa1, 0x9d, 0x0f, 0xfe, 0xf3, 0x62, 0x92, 0x86, 0x21, 0x88, 0x6f, 0x5f, 0xbe, 0xfc, 0xf8,
	0xee, 0xee, 0x83, 0x69, 0x55, 0xd7, 0x1d, 0xf5, 0xbd, 0x54, 0xea, 0x23, 0xd7, 0x9d, 0xda, 0x29,
	0x8c, 0xb4, 0x34, 0xa6, 0x35, 0xb2, 0xef, 0x00, 0xf8, 0x9a, 0xba, 0x5e, 0xb5, 0x88, 0xc0, 0xda,
	0xaa, 0xa1, 0x34, 0x98, 0x05, 0xf3, 0x48, 0x38, 0x8d, 0x13, 0x08, 0x37, 0x75, 0x1a, 0x1a, 0x67,
	0x2c, 0x8c, 0xc2, 0x23, 0x18, 0x53, 0x53, 0x6d, 0xb6, 0xe9, 0xc8, 0x41, 0xc3, 0x03, 0xaf, 0x81,
	0xeb, 0x77, 0xd5, 0x52, 0x9f, 0xb2, 0xd9, 0x68, 0x1e, 0x2f, 0xce, 0xf2, 0xfd, 0x80, 0x7c, 0x08,
	0xcf, 0xd7, 0xf6, 0xff, 0xe9, 0xb3, 0x91, 0xd4, 0x09, 0x0f, 0x9f, 0xbc, 0x40, 0xfc, 0xcb, 0xc6,
	0x63, 0xe0, 0xad, 0x53, 0x7e, 0x03, 0xff, 0xc2, 0x4b, 0x60, 0xbb, 0x2f, 0x4d, 0x6e, 0x8b, 0xc9,
	0xe2, 0xf4, 0x9f, 0xec, 0x67, 0x83, 0x08, 0x07, 0x66, 0x17, 0x10, 0xed, 0x2d, 0x04, 0xe0, 0xab,
	0xb2, 0xb8, 0x7f, 0x5c, 0x26, 0x07, 0x78, 0x08, 0xec, 0xae, 0x5c, 0x2d, 0x93, 0xc0, 0xaa, 0xd7,
	0x52, 0x3c, 0x24, 0x61, 0x76, 0x03, 0xf1, 0xed, 0x50, 0x50, 0x61, 0x22, 0xf1, 0xdc, 0x9c, 0x42,
	0x4a, 0x6f, 0x6d, 0x0d, 0xf6, 0x94, 0xe9, 0x9f, 0x71, 0xc2, 0x03, 0x05, 0x7b, 0x0b, 0xb5, 0x94,
	0xdc, 0x55, 0x7a, 0xf5, 0x13, 0x00, 0x00, 0xff, 0xff, 0x54, 0x00, 0x4f, 0x21, 0x67, 0x01, 0x00,
	0x00,
}
