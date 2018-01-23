// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

/*
Package node is a generated protocol buffer package.

It is generated from these files:
	node.proto

It has these top-level messages:
	NodeInfo
*/
package node

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

// NodeInfo represents a node in the k8s cluster.
// It records an unique node ID to identify the node within the cluster.
// ID determines IPAM for the givne node.
type NodeInfo struct {
	Id        uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress" json:"ip_address,omitempty"`
}

func (m *NodeInfo) Reset()                    { *m = NodeInfo{} }
func (m *NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*NodeInfo) ProtoMessage()               {}
func (*NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NodeInfo) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeInfo) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeInfo)(nil), "node.NodeInfo")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x4f, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x7c, 0xb9, 0x38, 0xfc, 0xf2,
	0x53, 0x52, 0x3d, 0xf3, 0xd2, 0xf2, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18,
	0x35, 0x78, 0x83, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x59, 0x2e, 0xae, 0xcc, 0x82, 0xf8, 0xc4, 0x94,
	0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x66, 0xb0, 0x0c, 0x67, 0x66, 0x81, 0x23, 0x44, 0x20, 0x89,
	0x0d, 0x6c, 0xb6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x15, 0xc8, 0x7f, 0x69, 0x00, 0x00,
	0x00,
}