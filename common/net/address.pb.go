// Code generated by protoc-gen-go.
// source: v2ray.com/core/common/net/address.proto
// DO NOT EDIT!

/*
Package net is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/common/net/address.proto
	v2ray.com/core/common/net/destination.proto
	v2ray.com/core/common/net/network.proto
	v2ray.com/core/common/net/port.proto

It has these top-level messages:
	AddressPB
	PortRange
*/
package net

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddressPB struct {
	// Types that are valid to be assigned to Address:
	//	*AddressPB_Ip
	//	*AddressPB_Domain
	Address isAddressPB_Address `protobuf_oneof:"address"`
}

func (m *AddressPB) Reset()                    { *m = AddressPB{} }
func (m *AddressPB) String() string            { return proto.CompactTextString(m) }
func (*AddressPB) ProtoMessage()               {}
func (*AddressPB) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isAddressPB_Address interface {
	isAddressPB_Address()
}

type AddressPB_Ip struct {
	Ip []byte `protobuf:"bytes,1,opt,name=ip,proto3,oneof"`
}
type AddressPB_Domain struct {
	Domain string `protobuf:"bytes,2,opt,name=domain,oneof"`
}

func (*AddressPB_Ip) isAddressPB_Address()     {}
func (*AddressPB_Domain) isAddressPB_Address() {}

func (m *AddressPB) GetAddress() isAddressPB_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AddressPB) GetIp() []byte {
	if x, ok := m.GetAddress().(*AddressPB_Ip); ok {
		return x.Ip
	}
	return nil
}

func (m *AddressPB) GetDomain() string {
	if x, ok := m.GetAddress().(*AddressPB_Domain); ok {
		return x.Domain
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AddressPB) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AddressPB_OneofMarshaler, _AddressPB_OneofUnmarshaler, _AddressPB_OneofSizer, []interface{}{
		(*AddressPB_Ip)(nil),
		(*AddressPB_Domain)(nil),
	}
}

func _AddressPB_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AddressPB)
	// address
	switch x := m.Address.(type) {
	case *AddressPB_Ip:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.Ip)
	case *AddressPB_Domain:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Domain)
	case nil:
	default:
		return fmt.Errorf("AddressPB.Address has unexpected type %T", x)
	}
	return nil
}

func _AddressPB_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AddressPB)
	switch tag {
	case 1: // address.ip
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Address = &AddressPB_Ip{x}
		return true, err
	case 2: // address.domain
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Address = &AddressPB_Domain{x}
		return true, err
	default:
		return false, nil
	}
}

func _AddressPB_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AddressPB)
	// address
	switch x := m.Address.(type) {
	case *AddressPB_Ip:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Ip)))
		n += len(x.Ip)
	case *AddressPB_Domain:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Domain)))
		n += len(x.Domain)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AddressPB)(nil), "com.v2ray.core.common.net.AddressPB")
}

func init() { proto.RegisterFile("v2ray.com/core/common/net/address.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0xcf, 0x4b, 0x2d, 0xd1, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4c, 0xce, 0xcf, 0xd5, 0x83, 0x29, 0x2e, 0x4a, 0xd5, 0x83, 0x28,
	0xd4, 0xcb, 0x4b, 0x2d, 0x51, 0x72, 0xe2, 0xe2, 0x74, 0x84, 0xa8, 0x0d, 0x70, 0x12, 0x12, 0xe0,
	0x62, 0xca, 0x2c, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0xf1, 0x60, 0x08, 0x62, 0xca, 0x2c, 0x10,
	0x92, 0xe0, 0x62, 0x4b, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x93, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0xf4,
	0x60, 0x08, 0x82, 0xf2, 0x9d, 0x38, 0xb9, 0xd8, 0xa1, 0x96, 0x38, 0xb1, 0x46, 0x31, 0xe7, 0xa5,
	0x96, 0x24, 0xb1, 0x81, 0x2d, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x04, 0x63, 0x74,
	0x97, 0x00, 0x00, 0x00,
}
