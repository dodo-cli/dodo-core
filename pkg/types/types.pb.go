// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/types/types.proto

package types

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ContainerId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerId) Reset()         { *m = ContainerId{} }
func (m *ContainerId) String() string { return proto.CompactTextString(m) }
func (*ContainerId) ProtoMessage()    {}
func (*ContainerId) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{1}
}

func (m *ContainerId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerId.Unmarshal(m, b)
}
func (m *ContainerId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerId.Marshal(b, m, deterministic)
}
func (m *ContainerId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerId.Merge(m, src)
}
func (m *ContainerId) XXX_Size() int {
	return xxx_messageInfo_ContainerId.Size(m)
}
func (m *ContainerId) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerId.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerId proto.InternalMessageInfo

func (m *ContainerId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Backdrop struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Aliases              []string       `protobuf:"bytes,2,rep,name=aliases,proto3" json:"aliases,omitempty"`
	ContainerName        string         `protobuf:"bytes,3,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	ImageId              string         `protobuf:"bytes,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Entrypoint           *Entrypoint    `protobuf:"bytes,5,opt,name=entrypoint,proto3" json:"entrypoint,omitempty"`
	Environment          []*Environment `protobuf:"bytes,6,rep,name=environment,proto3" json:"environment,omitempty"`
	Volumes              []*Volume      `protobuf:"bytes,7,rep,name=volumes,proto3" json:"volumes,omitempty"`
	Devices              []*Device      `protobuf:"bytes,8,rep,name=devices,proto3" json:"devices,omitempty"`
	Ports                []*Port        `protobuf:"bytes,9,rep,name=ports,proto3" json:"ports,omitempty"`
	User                 string         `protobuf:"bytes,10,opt,name=user,proto3" json:"user,omitempty"`
	WorkingDir           string         `protobuf:"bytes,11,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	Attach               bool           `protobuf:"varint,12,opt,name=attach,proto3" json:"attach,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Backdrop) Reset()         { *m = Backdrop{} }
func (m *Backdrop) String() string { return proto.CompactTextString(m) }
func (*Backdrop) ProtoMessage()    {}
func (*Backdrop) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{2}
}

func (m *Backdrop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backdrop.Unmarshal(m, b)
}
func (m *Backdrop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backdrop.Marshal(b, m, deterministic)
}
func (m *Backdrop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backdrop.Merge(m, src)
}
func (m *Backdrop) XXX_Size() int {
	return xxx_messageInfo_Backdrop.Size(m)
}
func (m *Backdrop) XXX_DiscardUnknown() {
	xxx_messageInfo_Backdrop.DiscardUnknown(m)
}

var xxx_messageInfo_Backdrop proto.InternalMessageInfo

func (m *Backdrop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Backdrop) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Backdrop) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *Backdrop) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *Backdrop) GetEntrypoint() *Entrypoint {
	if m != nil {
		return m.Entrypoint
	}
	return nil
}

func (m *Backdrop) GetEnvironment() []*Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *Backdrop) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Backdrop) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *Backdrop) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Backdrop) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Backdrop) GetWorkingDir() string {
	if m != nil {
		return m.WorkingDir
	}
	return ""
}

func (m *Backdrop) GetAttach() bool {
	if m != nil {
		return m.Attach
	}
	return false
}

type Entrypoint struct {
	Interactive          bool     `protobuf:"varint,1,opt,name=interactive,proto3" json:"interactive,omitempty"`
	Script               string   `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
	Interpreter          []string `protobuf:"bytes,3,rep,name=interpreter,proto3" json:"interpreter,omitempty"`
	Arguments            []string `protobuf:"bytes,4,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entrypoint) Reset()         { *m = Entrypoint{} }
func (m *Entrypoint) String() string { return proto.CompactTextString(m) }
func (*Entrypoint) ProtoMessage()    {}
func (*Entrypoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{3}
}

func (m *Entrypoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entrypoint.Unmarshal(m, b)
}
func (m *Entrypoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entrypoint.Marshal(b, m, deterministic)
}
func (m *Entrypoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entrypoint.Merge(m, src)
}
func (m *Entrypoint) XXX_Size() int {
	return xxx_messageInfo_Entrypoint.Size(m)
}
func (m *Entrypoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Entrypoint.DiscardUnknown(m)
}

var xxx_messageInfo_Entrypoint proto.InternalMessageInfo

func (m *Entrypoint) GetInteractive() bool {
	if m != nil {
		return m.Interactive
	}
	return false
}

func (m *Entrypoint) GetScript() string {
	if m != nil {
		return m.Script
	}
	return ""
}

func (m *Entrypoint) GetInterpreter() []string {
	if m != nil {
		return m.Interpreter
	}
	return nil
}

func (m *Entrypoint) GetArguments() []string {
	if m != nil {
		return m.Arguments
	}
	return nil
}

type Environment struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{4}
}

func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (m *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(m, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Environment) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Volume struct {
	Source               string   `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Readonly             bool     `protobuf:"varint,3,opt,name=readonly,proto3" json:"readonly,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Volume) Reset()         { *m = Volume{} }
func (m *Volume) String() string { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()    {}
func (*Volume) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{5}
}

func (m *Volume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Volume.Unmarshal(m, b)
}
func (m *Volume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Volume.Marshal(b, m, deterministic)
}
func (m *Volume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Volume.Merge(m, src)
}
func (m *Volume) XXX_Size() int {
	return xxx_messageInfo_Volume.Size(m)
}
func (m *Volume) XXX_DiscardUnknown() {
	xxx_messageInfo_Volume.DiscardUnknown(m)
}

var xxx_messageInfo_Volume proto.InternalMessageInfo

func (m *Volume) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Volume) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Volume) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

type Device struct {
	CgroupRule           string   `protobuf:"bytes,1,opt,name=cgroup_rule,json=cgroupRule,proto3" json:"cgroup_rule,omitempty"`
	Source               string   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
	Target               string   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	Permissions          string   `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{6}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetCgroupRule() string {
	if m != nil {
		return m.CgroupRule
	}
	return ""
}

func (m *Device) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Device) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Device) GetPermissions() string {
	if m != nil {
		return m.Permissions
	}
	return ""
}

type Port struct {
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Published            string   `protobuf:"bytes,2,opt,name=published,proto3" json:"published,omitempty"`
	Protocol             string   `protobuf:"bytes,3,opt,name=protocol,proto3" json:"protocol,omitempty"`
	HostIp               string   `protobuf:"bytes,4,opt,name=host_ip,json=hostIp,proto3" json:"host_ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{7}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Port) GetPublished() string {
	if m != nil {
		return m.Published
	}
	return ""
}

func (m *Port) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Port) GetHostIp() string {
	if m != nil {
		return m.HostIp
	}
	return ""
}

type ClientOptions struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Host                 string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	CaFile               string   `protobuf:"bytes,3,opt,name=ca_file,json=caFile,proto3" json:"ca_file,omitempty"`
	CertFile             string   `protobuf:"bytes,4,opt,name=cert_file,json=certFile,proto3" json:"cert_file,omitempty"`
	KeyFile              string   `protobuf:"bytes,5,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientOptions) Reset()         { *m = ClientOptions{} }
func (m *ClientOptions) String() string { return proto.CompactTextString(m) }
func (*ClientOptions) ProtoMessage()    {}
func (*ClientOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_6579d14f41ea6320, []int{8}
}

func (m *ClientOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientOptions.Unmarshal(m, b)
}
func (m *ClientOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientOptions.Marshal(b, m, deterministic)
}
func (m *ClientOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientOptions.Merge(m, src)
}
func (m *ClientOptions) XXX_Size() int {
	return xxx_messageInfo_ClientOptions.Size(m)
}
func (m *ClientOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ClientOptions proto.InternalMessageInfo

func (m *ClientOptions) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ClientOptions) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ClientOptions) GetCaFile() string {
	if m != nil {
		return m.CaFile
	}
	return ""
}

func (m *ClientOptions) GetCertFile() string {
	if m != nil {
		return m.CertFile
	}
	return ""
}

func (m *ClientOptions) GetKeyFile() string {
	if m != nil {
		return m.KeyFile
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "types.Empty")
	proto.RegisterType((*ContainerId)(nil), "types.ContainerId")
	proto.RegisterType((*Backdrop)(nil), "types.Backdrop")
	proto.RegisterType((*Entrypoint)(nil), "types.Entrypoint")
	proto.RegisterType((*Environment)(nil), "types.Environment")
	proto.RegisterType((*Volume)(nil), "types.Volume")
	proto.RegisterType((*Device)(nil), "types.Device")
	proto.RegisterType((*Port)(nil), "types.Port")
	proto.RegisterType((*ClientOptions)(nil), "types.ClientOptions")
}

func init() {
	proto.RegisterFile("pkg/types/types.proto", fileDescriptor_6579d14f41ea6320)
}

var fileDescriptor_6579d14f41ea6320 = []byte{
	// 680 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x4d, 0x6f, 0xd4, 0x48,
	0x10, 0xd5, 0x7c, 0x7b, 0xca, 0x49, 0x36, 0xdb, 0x9b, 0xdd, 0xed, 0x84, 0x20, 0x06, 0x4b, 0x88,
	0xb9, 0x10, 0x44, 0x00, 0x21, 0xae, 0x24, 0x01, 0xe5, 0x02, 0x91, 0x05, 0x5c, 0x47, 0x1d, 0xbb,
	0x32, 0x69, 0x8d, 0xc7, 0xdd, 0x74, 0xb7, 0x07, 0x8d, 0xb8, 0x73, 0xe4, 0x5f, 0xf0, 0x07, 0xf8,
	0x85, 0xa8, 0x3f, 0xec, 0xf1, 0x80, 0xb8, 0x58, 0xfd, 0xde, 0xab, 0xaa, 0x57, 0xae, 0xfe, 0x80,
	0x7f, 0xe5, 0x62, 0xfe, 0xd8, 0xac, 0x25, 0x6a, 0xff, 0x3d, 0x91, 0x4a, 0x18, 0x41, 0x06, 0x0e,
	0x24, 0x23, 0x18, 0x5c, 0x2c, 0xa5, 0x59, 0x27, 0x77, 0x21, 0x3e, 0x13, 0xa5, 0x61, 0xbc, 0x44,
	0x75, 0x99, 0x93, 0x3d, 0xe8, 0xf2, 0x9c, 0x76, 0x26, 0x9d, 0xe9, 0x38, 0xed, 0xf2, 0x3c, 0xf9,
	0xde, 0x83, 0xe8, 0x15, 0xcb, 0x16, 0xb9, 0x12, 0x92, 0x10, 0xe8, 0x97, 0x6c, 0x89, 0x41, 0x76,
	0x6b, 0x42, 0x61, 0xc4, 0x0a, 0xce, 0x34, 0x6a, 0xda, 0x9d, 0xf4, 0xa6, 0xe3, 0xb4, 0x86, 0xe4,
	0x01, 0xec, 0x65, 0x75, 0xe5, 0x99, 0xcb, 0xeb, 0xb9, 0xbc, 0xdd, 0x86, 0x7d, 0x6b, 0x0b, 0x1c,
	0x42, 0xc4, 0x97, 0x6c, 0x8e, 0x33, 0x9e, 0xd3, 0xbe, 0x0b, 0x18, 0x39, 0x7c, 0x99, 0x93, 0x27,
	0x00, 0x58, 0x1a, 0xb5, 0x96, 0x82, 0x97, 0x86, 0x0e, 0x26, 0x9d, 0x69, 0x7c, 0xfa, 0xf7, 0x89,
	0xff, 0x9b, 0x8b, 0x46, 0x48, 0x5b, 0x41, 0xe4, 0x19, 0xc4, 0x58, 0xae, 0xb8, 0x12, 0xe5, 0x12,
	0x4b, 0x43, 0x87, 0x93, 0xde, 0x34, 0x3e, 0x25, 0x4d, 0x4e, 0xa3, 0xa4, 0xed, 0x30, 0xf2, 0x10,
	0x46, 0x2b, 0x51, 0x54, 0x4b, 0xd4, 0x74, 0xe4, 0x32, 0x76, 0x43, 0xc6, 0x47, 0xc7, 0xa6, 0xb5,
	0x6a, 0x03, 0x73, 0x5c, 0xf1, 0x0c, 0x35, 0x8d, 0xb6, 0x02, 0xcf, 0x1d, 0x9b, 0xd6, 0x2a, 0xb9,
	0x0f, 0x03, 0x29, 0x94, 0xd1, 0x74, 0xec, 0xc2, 0xe2, 0x10, 0x76, 0x25, 0x94, 0x49, 0xbd, 0x62,
	0xa7, 0x59, 0x69, 0x54, 0x14, 0xfc, 0x34, 0xed, 0x9a, 0xdc, 0x83, 0xf8, 0xb3, 0x50, 0x0b, 0x5e,
	0xce, 0x67, 0x39, 0x57, 0x34, 0x76, 0x12, 0x04, 0xea, 0x9c, 0x2b, 0xf2, 0x1f, 0x0c, 0x99, 0x31,
	0x2c, 0xbb, 0xa5, 0x3b, 0x93, 0xce, 0x34, 0x4a, 0x03, 0x4a, 0xbe, 0x76, 0x00, 0x36, 0x23, 0x21,
	0x13, 0x88, 0x79, 0x69, 0x50, 0xb1, 0xcc, 0xf0, 0x95, 0xdf, 0xb0, 0x28, 0x6d, 0x53, 0xb6, 0x90,
	0xce, 0x14, 0x97, 0x86, 0x76, 0x9d, 0x49, 0x40, 0x4d, 0xa6, 0x54, 0x68, 0x50, 0xd1, 0x9e, 0xdb,
	0xd3, 0x36, 0x45, 0x8e, 0x61, 0xcc, 0xd4, 0xbc, 0xb2, 0x83, 0xd3, 0xb4, 0xef, 0xf4, 0x0d, 0x91,
	0x3c, 0x87, 0xb8, 0x35, 0x66, 0xb2, 0x0f, 0xbd, 0x05, 0xae, 0xc3, 0x89, 0xb1, 0x4b, 0x72, 0x00,
	0x83, 0x15, 0x2b, 0x2a, 0x0c, 0xbe, 0x1e, 0x24, 0xef, 0x61, 0xe8, 0x67, 0xed, 0x1a, 0x13, 0x95,
	0xca, 0xea, 0x63, 0x16, 0x90, 0xe5, 0x0d, 0x53, 0x73, 0x6c, 0x1a, 0xf6, 0x88, 0x1c, 0x41, 0xa4,
	0x90, 0xe5, 0xa2, 0x2c, 0xd6, 0xee, 0x80, 0x45, 0x69, 0x83, 0x93, 0x2f, 0x30, 0xf4, 0x1b, 0x63,
	0x07, 0x9b, 0xcd, 0x95, 0xa8, 0xe4, 0x4c, 0x55, 0x45, 0x5d, 0x1a, 0x3c, 0x95, 0x56, 0x45, 0xdb,
	0xb6, 0xfb, 0x07, 0xdb, 0xde, 0x96, 0xed, 0x04, 0x62, 0x89, 0x6a, 0xc9, 0xb5, 0xe6, 0xa2, 0xd4,
	0xe1, 0xe4, 0xb6, 0xa9, 0xe4, 0x13, 0xf4, 0xed, 0x76, 0xb7, 0x2a, 0x74, 0xb6, 0x2a, 0x1c, 0xc3,
	0x58, 0x56, 0xd7, 0x05, 0xd7, 0xb7, 0x98, 0x07, 0xd3, 0x0d, 0x61, 0x7f, 0xcb, 0x5d, 0xd8, 0x4c,
	0x14, 0xc1, 0xb9, 0xc1, 0xe4, 0x7f, 0x18, 0xdd, 0x0a, 0x6d, 0x66, 0x5c, 0x06, 0xdf, 0xa1, 0x85,
	0x97, 0x32, 0xf9, 0xd6, 0x81, 0xdd, 0xb3, 0x82, 0x63, 0x69, 0xde, 0x49, 0x63, 0x9b, 0xb0, 0xd7,
	0x73, 0x85, 0xca, 0x36, 0x14, 0xdc, 0x6b, 0x68, 0x8f, 0x9f, 0xcd, 0x0a, 0xce, 0x6e, 0x6d, 0x0b,
	0x67, 0x6c, 0x76, 0xc3, 0x8b, 0xfa, 0xae, 0x0e, 0x33, 0xf6, 0x9a, 0x17, 0x48, 0xee, 0xc0, 0x38,
	0x43, 0x65, 0xbc, 0xe4, 0x3d, 0x23, 0x4b, 0x38, 0xf1, 0x10, 0xa2, 0x05, 0xae, 0xbd, 0x36, 0xf0,
	0x26, 0x0b, 0x5c, 0x5b, 0xe9, 0xf4, 0x87, 0x6d, 0x48, 0x94, 0x37, 0x7c, 0x5e, 0x29, 0x66, 0x3b,
	0x22, 0x2f, 0x61, 0xff, 0x0d, 0x9a, 0xed, 0x26, 0xff, 0x0a, 0xb7, 0xa3, 0x7e, 0x68, 0x8e, 0x0e,
	0x02, 0xb1, 0x1d, 0xf6, 0x02, 0xfe, 0xf9, 0x20, 0x73, 0x66, 0x70, 0xbb, 0xe2, 0x6f, 0xd9, 0xbf,
	0x12, 0xe4, 0x11, 0x8c, 0xaf, 0x94, 0x58, 0x71, 0xff, 0xdf, 0x75, 0xed, 0xcd, 0xab, 0x77, 0xb4,
	0x53, 0x3f, 0x10, 0xf6, 0x49, 0xbc, 0x1e, 0xba, 0x41, 0x3f, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff,
	0x8f, 0xf1, 0xcd, 0x09, 0x42, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigurationClient interface {
	GetClientOptions(ctx context.Context, in *Backdrop, opts ...grpc.CallOption) (*ClientOptions, error)
	UpdateConfiguration(ctx context.Context, in *Backdrop, opts ...grpc.CallOption) (*Backdrop, error)
	Provision(ctx context.Context, in *ContainerId, opts ...grpc.CallOption) (*Empty, error)
}

type configurationClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationClient(cc grpc.ClientConnInterface) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) GetClientOptions(ctx context.Context, in *Backdrop, opts ...grpc.CallOption) (*ClientOptions, error) {
	out := new(ClientOptions)
	err := c.cc.Invoke(ctx, "/types.Configuration/GetClientOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) UpdateConfiguration(ctx context.Context, in *Backdrop, opts ...grpc.CallOption) (*Backdrop, error) {
	out := new(Backdrop)
	err := c.cc.Invoke(ctx, "/types.Configuration/UpdateConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Provision(ctx context.Context, in *ContainerId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/types.Configuration/Provision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
type ConfigurationServer interface {
	GetClientOptions(context.Context, *Backdrop) (*ClientOptions, error)
	UpdateConfiguration(context.Context, *Backdrop) (*Backdrop, error)
	Provision(context.Context, *ContainerId) (*Empty, error)
}

// UnimplementedConfigurationServer can be embedded to have forward compatible implementations.
type UnimplementedConfigurationServer struct {
}

func (*UnimplementedConfigurationServer) GetClientOptions(ctx context.Context, req *Backdrop) (*ClientOptions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientOptions not implemented")
}
func (*UnimplementedConfigurationServer) UpdateConfiguration(ctx context.Context, req *Backdrop) (*Backdrop, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfiguration not implemented")
}
func (*UnimplementedConfigurationServer) Provision(ctx context.Context, req *ContainerId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Provision not implemented")
}

func RegisterConfigurationServer(s *grpc.Server, srv ConfigurationServer) {
	s.RegisterService(&_Configuration_serviceDesc, srv)
}

func _Configuration_GetClientOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Backdrop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetClientOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Configuration/GetClientOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetClientOptions(ctx, req.(*Backdrop))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_UpdateConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Backdrop)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).UpdateConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Configuration/UpdateConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).UpdateConfiguration(ctx, req.(*Backdrop))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Provision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainerId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Provision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Configuration/Provision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Provision(ctx, req.(*ContainerId))
	}
	return interceptor(ctx, in, info, handler)
}

var _Configuration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClientOptions",
			Handler:    _Configuration_GetClientOptions_Handler,
		},
		{
			MethodName: "UpdateConfiguration",
			Handler:    _Configuration_UpdateConfiguration_Handler,
		},
		{
			MethodName: "Provision",
			Handler:    _Configuration_Provision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/types/types.proto",
}
