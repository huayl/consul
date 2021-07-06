// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/clusters/dynamic_forward_proxy/v3/cluster.proto

package envoy_extensions_clusters_dynamic_forward_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/common/dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ClusterConfig struct {
	DnsCacheConfig       *v3.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClusterConfig) Reset()         { *m = ClusterConfig{} }
func (m *ClusterConfig) String() string { return proto.CompactTextString(m) }
func (*ClusterConfig) ProtoMessage()    {}
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e2f35129a7471fd, []int{0}
}

func (m *ClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterConfig.Unmarshal(m, b)
}
func (m *ClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterConfig.Marshal(b, m, deterministic)
}
func (m *ClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfig.Merge(m, src)
}
func (m *ClusterConfig) XXX_Size() int {
	return xxx_messageInfo_ClusterConfig.Size(m)
}
func (m *ClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfig proto.InternalMessageInfo

func (m *ClusterConfig) GetDnsCacheConfig() *v3.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterConfig)(nil), "envoy.extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/clusters/dynamic_forward_proxy/v3/cluster.proto", fileDescriptor_5e2f35129a7471fd)
}

var fileDescriptor_5e2f35129a7471fd = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xce, 0x29,
	0x2d, 0x2e, 0x49, 0x2d, 0x2a, 0xd6, 0x4f, 0xa9, 0xcc, 0x4b, 0xcc, 0xcd, 0x4c, 0x8e, 0x4f, 0xcb,
	0x2f, 0x2a, 0x4f, 0x2c, 0x4a, 0x89, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x2f, 0x33, 0x86, 0xa9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x02, 0x9b, 0xa0, 0x87, 0x30, 0x41, 0x0f, 0x66,
	0x82, 0x1e, 0x56, 0x13, 0xf4, 0xca, 0x8c, 0xa5, 0xb0, 0xd8, 0x9a, 0x9f, 0x9b, 0x9b, 0x9f, 0x87,
	0xdb, 0xce, 0x94, 0xbc, 0xe2, 0xf8, 0xe4, 0xc4, 0xe4, 0x8c, 0x54, 0x88, 0xad, 0x52, 0xb2, 0xa5,
	0x29, 0x05, 0x89, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25, 0x89, 0x25, 0x60, 0x13, 0x8a, 0x4b, 0x12,
	0x4b, 0x4a, 0x8b, 0xa1, 0xd2, 0x8a, 0x18, 0xd2, 0x65, 0xa9, 0x45, 0x20, 0x9b, 0x32, 0xf3, 0xd2,
	0xa1, 0x4a, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x84,
	0xd2, 0x39, 0x46, 0x2e, 0x5e, 0x67, 0x88, 0x17, 0x9c, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0x85, 0x4a,
	0xb8, 0x04, 0xe0, 0xf6, 0xc7, 0x27, 0x83, 0xc5, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x1c,
	0xf4, 0x30, 0x7d, 0x0f, 0xf6, 0x09, 0x4e, 0xbf, 0xeb, 0xb9, 0xe4, 0x15, 0x3b, 0x83, 0x0c, 0x82,
	0x98, 0xed, 0xc4, 0xf1, 0xcb, 0x89, 0xb5, 0x8b, 0x91, 0x49, 0x80, 0x31, 0x88, 0x2f, 0x05, 0x45,
	0xc6, 0xca, 0x7d, 0xd6, 0xd1, 0x0e, 0x39, 0x27, 0x68, 0x0c, 0xe9, 0x41, 0xac, 0x85, 0x85, 0x2d,
	0x2e, 0xe3, 0x8d, 0x12, 0x73, 0x0a, 0x32, 0x12, 0xf5, 0x50, 0x9c, 0xef, 0x14, 0xb5, 0xab, 0xe1,
	0xc4, 0x45, 0x36, 0x26, 0x01, 0x26, 0x2e, 0x87, 0xcc, 0x7c, 0x88, 0x83, 0x21, 0xca, 0x49, 0x8f,
	0x39, 0x27, 0x1e, 0xa8, 0xd1, 0x01, 0xa0, 0xa0, 0x0a, 0x60, 0x4c, 0x62, 0x03, 0x87, 0x99, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xd8, 0x96, 0x48, 0x48, 0x02, 0x00, 0x00,
}
