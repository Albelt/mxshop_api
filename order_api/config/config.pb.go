// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.0--rc1
// source: config.proto

package config

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *Server   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Redis    *Redis    `protobuf:"bytes,4,opt,name=redis,proto3" json:"redis,omitempty"`
	Consul   *Consul   `protobuf:"bytes,6,opt,name=consul,proto3" json:"consul,omitempty"`
	GrpcSrvs *GrpcSrvs `protobuf:"bytes,7,opt,name=grpc_srvs,json=grpcSrvs,proto3" json:"grpc_srvs,omitempty" mapstructure:"grpc_srvs"` //@gotags: mapstructure:"grpc_srvs"
	Alipay   *Alipay   `protobuf:"bytes,8,opt,name=alipay,proto3" json:"alipay,omitempty"`
	Tracing  *Tracing  `protobuf:"bytes,9,opt,name=tracing,proto3" json:"tracing,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Config) GetRedis() *Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Config) GetConsul() *Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *Config) GetGrpcSrvs() *GrpcSrvs {
	if x != nil {
		return x.GrpcSrvs
	}
	return nil
}

func (x *Config) GetAlipay() *Alipay {
	if x != nil {
		return x.Alipay
	}
	return nil
}

func (x *Config) GetTracing() *Tracing {
	if x != nil {
		return x.Tracing
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip   string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// 服务名称、标签、出口IP
	Name       string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Tags       []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	ExternalIp string   `protobuf:"bytes,5,opt,name=external_ip,json=externalIp,proto3" json:"external_ip,omitempty" mapstructure:"external_ip"` //@gotags: mapstructure:"external_ip"
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Server) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Server) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Server) GetExternalIp() string {
	if x != nil {
		return x.ExternalIp
	}
	return ""
}

type Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr     string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Db       int64  `protobuf:"varint,3,opt,name=db,proto3" json:"db,omitempty"`
}

func (x *Redis) Reset() {
	*x = Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redis) ProtoMessage() {}

func (x *Redis) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redis.ProtoReflect.Descriptor instead.
func (*Redis) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Redis) GetDb() int64 {
	if x != nil {
		return x.Db
	}
	return 0
}

type Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr        string              `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	HealthCheck *Consul_HealthCheck `protobuf:"bytes,2,opt,name=health_check,json=healthCheck,proto3" json:"health_check,omitempty" mapstructure:"health_check"` //@gotags: mapstructure:"health_check"
}

func (x *Consul) Reset() {
	*x = Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consul) ProtoMessage() {}

func (x *Consul) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consul.ProtoReflect.Descriptor instead.
func (*Consul) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *Consul) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Consul) GetHealthCheck() *Consul_HealthCheck {
	if x != nil {
		return x.HealthCheck
	}
	return nil
}

type GrpcSrvs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderSrvName string `protobuf:"bytes,1,opt,name=order_srv_name,json=orderSrvName,proto3" json:"order_srv_name,omitempty" mapstructure:"order_srv_name"` //@gotags: mapstructure:"order_srv_name"
}

func (x *GrpcSrvs) Reset() {
	*x = GrpcSrvs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrpcSrvs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrpcSrvs) ProtoMessage() {}

func (x *GrpcSrvs) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrpcSrvs.ProtoReflect.Descriptor instead.
func (*GrpcSrvs) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *GrpcSrvs) GetOrderSrvName() string {
	if x != nil {
		return x.OrderSrvName
	}
	return ""
}

type Alipay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Appid        string `protobuf:"bytes,1,opt,name=appid,proto3" json:"appid,omitempty"`
	AliPublicKey string `protobuf:"bytes,2,opt,name=ali_public_key,json=aliPublicKey,proto3" json:"ali_public_key,omitempty" mapstructure:"ali_public_key"` //@gotags: mapstructure:"ali_public_key"
	PublicKey    string `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty" mapstructure:"public_key"`            //@gotags: mapstructure:"public_key"
	PrivateKey   string `protobuf:"bytes,4,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty" mapstructure:"private_key"`         //@gotags: mapstructure:"private_key"
	NotifyUrl    string `protobuf:"bytes,5,opt,name=notify_url,json=notifyUrl,proto3" json:"notify_url,omitempty" mapstructure:"notify_url"`            //@gotags: mapstructure:"notify_url"
	ReturnUrl    string `protobuf:"bytes,6,opt,name=return_url,json=returnUrl,proto3" json:"return_url,omitempty" mapstructure:"return_url"`            //@gotags: mapstructure:"return_url"
}

func (x *Alipay) Reset() {
	*x = Alipay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alipay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alipay) ProtoMessage() {}

func (x *Alipay) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alipay.ProtoReflect.Descriptor instead.
func (*Alipay) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *Alipay) GetAppid() string {
	if x != nil {
		return x.Appid
	}
	return ""
}

func (x *Alipay) GetAliPublicKey() string {
	if x != nil {
		return x.AliPublicKey
	}
	return ""
}

func (x *Alipay) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Alipay) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

func (x *Alipay) GetNotifyUrl() string {
	if x != nil {
		return x.NotifyUrl
	}
	return ""
}

func (x *Alipay) GetReturnUrl() string {
	if x != nil {
		return x.ReturnUrl
	}
	return ""
}

type Tracing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JaegerAddr string `protobuf:"bytes,1,opt,name=jaeger_addr,json=jaegerAddr,proto3" json:"jaeger_addr,omitempty" mapstructure:"jaeger_addr"` //@gotags: mapstructure:"jaeger_addr"
}

func (x *Tracing) Reset() {
	*x = Tracing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tracing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tracing) ProtoMessage() {}

func (x *Tracing) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tracing.ProtoReflect.Descriptor instead.
func (*Tracing) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{6}
}

func (x *Tracing) GetJaegerAddr() string {
	if x != nil {
		return x.JaegerAddr
	}
	return ""
}

type Consul_HealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interval   string `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	Timeout    string `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Deregister string `protobuf:"bytes,3,opt,name=deregister,proto3" json:"deregister,omitempty"`
}

func (x *Consul_HealthCheck) Reset() {
	*x = Consul_HealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consul_HealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consul_HealthCheck) ProtoMessage() {}

func (x *Consul_HealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consul_HealthCheck.ProtoReflect.Descriptor instead.
func (*Consul_HealthCheck) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Consul_HealthCheck) GetInterval() string {
	if x != nil {
		return x.Interval
	}
	return ""
}

func (x *Consul_HealthCheck) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

func (x *Consul_HealthCheck) GetDeregister() string {
	if x != nil {
		return x.Deregister
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5,
	0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x05, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x52, 0x65, 0x64, 0x69,
	0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x26, 0x0a, 0x09, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x73, 0x72, 0x76, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x47,
	0x72, 0x70, 0x63, 0x53, 0x72, 0x76, 0x73, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x53, 0x72, 0x76,
	0x73, 0x12, 0x1f, 0x0a, 0x06, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x52, 0x06, 0x61, 0x6c, 0x69, 0x70,
	0x61, 0x79, 0x12, 0x22, 0x0a, 0x07, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x22, 0x75, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x22, 0x47, 0x0a,
	0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x64, 0x62, 0x22, 0xb9, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x36, 0x0a, 0x0c, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x1a, 0x63, 0x0a,
	0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x22, 0x30, 0x0a, 0x08, 0x47, 0x72, 0x70, 0x63, 0x53, 0x72, 0x76, 0x73, 0x12, 0x24,
	0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x72, 0x76, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x72, 0x76,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x06, 0x41, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x70, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x6c, 0x69, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x6c, 0x69, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x2a, 0x0a, 0x07, 0x54, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x6a, 0x61, 0x65, 0x67, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x61, 0x65, 0x67, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_config_proto_goTypes = []interface{}{
	(*Config)(nil),             // 0: Config
	(*Server)(nil),             // 1: Server
	(*Redis)(nil),              // 2: Redis
	(*Consul)(nil),             // 3: Consul
	(*GrpcSrvs)(nil),           // 4: GrpcSrvs
	(*Alipay)(nil),             // 5: Alipay
	(*Tracing)(nil),            // 6: Tracing
	(*Consul_HealthCheck)(nil), // 7: Consul.HealthCheck
}
var file_config_proto_depIdxs = []int32{
	1, // 0: Config.server:type_name -> Server
	2, // 1: Config.redis:type_name -> Redis
	3, // 2: Config.consul:type_name -> Consul
	4, // 3: Config.grpc_srvs:type_name -> GrpcSrvs
	5, // 4: Config.alipay:type_name -> Alipay
	6, // 5: Config.tracing:type_name -> Tracing
	7, // 6: Consul.health_check:type_name -> Consul.HealthCheck
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redis); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consul); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrpcSrvs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alipay); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tracing); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consul_HealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}