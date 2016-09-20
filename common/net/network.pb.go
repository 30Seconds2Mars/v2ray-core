// Code generated by protoc-gen-go.
// source: v2ray.com/core/common/net/network.proto
// DO NOT EDIT!

package net

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Network int32

const (
	Network_Unknown   Network = 0
	Network_RawTCP    Network = 1
	Network_TCP       Network = 2
	Network_UDP       Network = 3
	Network_KCP       Network = 4
	Network_WebSocket Network = 5
)

var Network_name = map[int32]string{
	0: "Unknown",
	1: "RawTCP",
	2: "TCP",
	3: "UDP",
	4: "KCP",
	5: "WebSocket",
}
var Network_value = map[string]int32{
	"Unknown":   0,
	"RawTCP":    1,
	"TCP":       2,
	"UDP":       3,
	"KCP":       4,
	"WebSocket": 5,
}

func (x Network) String() string {
	return proto.EnumName(Network_name, int32(x))
}
func (Network) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterEnum("com.v2ray.core.common.net.Network", Network_name, Network_value)
}

func init() { proto.RegisterFile("v2ray.com/core/common/net/network.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2f, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0xcf, 0x4b, 0x2d, 0x01, 0xe1, 0xf2, 0xfc, 0xa2, 0x6c, 0xbd, 0x82, 0xa2, 0xfc, 0x92,
	0x7c, 0x21, 0xc9, 0xe4, 0xfc, 0x5c, 0x3d, 0x98, 0xe2, 0xa2, 0x54, 0x3d, 0x88, 0x42, 0xbd, 0xbc,
	0xd4, 0x12, 0x2d, 0x1f, 0x2e, 0x76, 0x3f, 0x88, 0x5a, 0x21, 0x6e, 0x2e, 0xf6, 0xd0, 0xbc, 0xec,
	0xbc, 0xfc, 0xf2, 0x3c, 0x01, 0x06, 0x21, 0x2e, 0x2e, 0xb6, 0xa0, 0xc4, 0xf2, 0x10, 0xe7, 0x00,
	0x01, 0x46, 0x21, 0x76, 0x2e, 0x66, 0x10, 0x83, 0x09, 0xc4, 0x08, 0x75, 0x09, 0x10, 0x60, 0x06,
	0x31, 0xbc, 0x9d, 0x03, 0x04, 0x58, 0x84, 0x78, 0xb9, 0x38, 0xc3, 0x53, 0x93, 0x82, 0xf3, 0x93,
	0xb3, 0x53, 0x4b, 0x04, 0x58, 0x9d, 0x58, 0xa3, 0x98, 0xf3, 0x52, 0x4b, 0x92, 0xd8, 0xc0, 0xd6,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x83, 0xa9, 0x75, 0xb3, 0xa1, 0x00, 0x00, 0x00,
}
