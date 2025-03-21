// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: base/v1/v10.proto

package base

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterVehicleRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Either chassis number or vin must be provided.
	ChassisNumber string `protobuf:"bytes,1,opt,name=chassisNumber,proto3" json:"chassisNumber,omitempty"`
	Vin           string `protobuf:"bytes,2,opt,name=vin,proto3" json:"vin,omitempty"`
	Make          string `protobuf:"bytes,3,opt,name=make,proto3" json:"make,omitempty"`
	Model         string `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty"`
	Year          int32  `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
	Kilometers    int32  `protobuf:"varint,6,opt,name=kilometers,proto3" json:"kilometers,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterVehicleRequest) Reset() {
	*x = RegisterVehicleRequest{}
	mi := &file_base_v1_v10_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterVehicleRequest) ProtoMessage() {}

func (x *RegisterVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_v1_v10_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterVehicleRequest.ProtoReflect.Descriptor instead.
func (*RegisterVehicleRequest) Descriptor() ([]byte, []int) {
	return file_base_v1_v10_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterVehicleRequest) GetChassisNumber() string {
	if x != nil {
		return x.ChassisNumber
	}
	return ""
}

func (x *RegisterVehicleRequest) GetVin() string {
	if x != nil {
		return x.Vin
	}
	return ""
}

func (x *RegisterVehicleRequest) GetMake() string {
	if x != nil {
		return x.Make
	}
	return ""
}

func (x *RegisterVehicleRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *RegisterVehicleRequest) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *RegisterVehicleRequest) GetKilometers() int32 {
	if x != nil {
		return x.Kilometers
	}
	return 0
}

// Request message for the Do rpc.
type DoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DoRequest) Reset() {
	*x = DoRequest{}
	mi := &file_base_v1_v10_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoRequest) ProtoMessage() {}

func (x *DoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_base_v1_v10_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoRequest.ProtoReflect.Descriptor instead.
func (*DoRequest) Descriptor() ([]byte, []int) {
	return file_base_v1_v10_proto_rawDescGZIP(), []int{1}
}

func (x *DoRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// Response message for the Do rpc.
type DoResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          string                 `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DoResponse) Reset() {
	*x = DoResponse{}
	mi := &file_base_v1_v10_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoResponse) ProtoMessage() {}

func (x *DoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_base_v1_v10_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoResponse.ProtoReflect.Descriptor instead.
func (*DoResponse) Descriptor() ([]byte, []int) {
	return file_base_v1_v10_proto_rawDescGZIP(), []int{2}
}

func (x *DoResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_base_v1_v10_proto protoreflect.FileDescriptor

var file_base_v1_v10_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x31, 0x30, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x76, 0x31, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xae, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x73, 0x73, 0x69, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x76, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6d, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x69, 0x6c, 0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6b, 0x69, 0x6c, 0x6f, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x22, 0x1f, 0x0a, 0x09, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x20, 0x0a, 0x0a, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xee, 0x02, 0x0a, 0x03, 0x56, 0x31, 0x30, 0x12, 0x56, 0x0a, 0x02,
	0x44, 0x6f, 0x12, 0x1b, 0x2e, 0x76, 0x31, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x76, 0x31, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x76,
	0x31, 0x3a, 0x64, 0x6f, 0x12, 0x75, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x28, 0x2e, 0x76, 0x31, 0x30, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x30, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x97, 0x01, 0x92, 0x41,
	0x93, 0x01, 0x12, 0x30, 0x28, 0x41, 0x4c, 0x50, 0x48, 0x41, 0x29, 0x20, 0x56, 0x31, 0x30, 0x20,
	0x41, 0x50, 0x49, 0x2e, 0x20, 0x42, 0x61, 0x73, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x3a, 0x20, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x61,
	0x6c, 0x2e, 0x61, 0x69, 0x1a, 0x5f, 0x0a, 0x24, 0x53, 0x65, 0x65, 0x20, 0x68, 0x65, 0x72, 0x65,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x12, 0x37, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x30, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x65, 0x65, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76,
	0x31, 0x30, 0x2f, 0x76, 0x31, 0x42, 0xbc, 0x05, 0x92, 0x41, 0x96, 0x05, 0x12, 0x81, 0x05, 0x0a,
	0x11, 0x56, 0x31, 0x30, 0x20, 0x41, 0x50, 0x49, 0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0xe4, 0x04, 0x44, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x20, 0x41, 0x49, 0x27, 0x73,
	0x20, 0x2a, 0x2a, 0x56, 0x31, 0x30, 0x2a, 0x2a, 0x20, 0x41, 0x50, 0x49, 0x20, 0x69, 0x73, 0x20,
	0x61, 0x20, 0x52, 0x45, 0x53, 0x54, 0x66, 0x75, 0x6c, 0x20, 0x41, 0x50, 0x49, 0x20, 0x74, 0x68,
	0x61, 0x74, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x62, 0x65, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x65, 0x64, 0x20, 0x62, 0x79, 0x20, 0x61, 0x6e, 0x20, 0x48, 0x54, 0x54, 0x50, 0x20, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x20, 0x73, 0x75, 0x63, 0x68, 0x20, 0x61, 0x73, 0x20, 0x60, 0x63, 0x75,
	0x72, 0x6c, 0x60, 0x2c, 0x20, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x79, 0x20, 0x48, 0x54, 0x54, 0x50,
	0x20, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x20, 0x77, 0x68, 0x69, 0x63, 0x68, 0x20, 0x69,
	0x73, 0x20, 0x70, 0x61, 0x72, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6d, 0x6f, 0x73, 0x74, 0x20, 0x6d,
	0x6f, 0x64, 0x65, 0x72, 0x6e, 0x20, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e,
	0x67, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x20, 0x54, 0x68, 0x69,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x20, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x20,
	0x69, 0x73, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x5b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x20,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2d,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x29, 0x20, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64,
	0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x69, 0x73, 0x20, 0x5b, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x5d, 0x28, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x2d,
	0x61, 0x69, 0x2f, 0x76, 0x31, 0x30, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x29, 0x2e, 0x0a, 0x0a,
	0x53, 0x6f, 0x6d, 0x65, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2c, 0x20,
	0x65, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c, 0x6c, 0x79, 0x20, 0x74, 0x68, 0x6f, 0x73, 0x65,
	0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x20, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x20, 0x6f, 0x66, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2c,
	0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x20,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x3b, 0x20, 0x6e, 0x65, 0x77, 0x6c, 0x69,
	0x6e, 0x65, 0x2d, 0x73, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x65, 0x64, 0x20, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x20, 0x6f, 0x66, 0x20, 0xe2, 0x80, 0x9c, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0xe2, 0x80, 0x9d, 0x2e, 0x20, 0x45, 0x61, 0x63, 0x68, 0x20, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x20,
	0x69, 0x73, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x20, 0x74,
	0x68, 0x61, 0x74, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x20,
	0x65, 0x69, 0x74, 0x68, 0x65, 0x72, 0x20, 0x61, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x20, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x20, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x20, 0x4f, 0x6e, 0x6c, 0x79, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x6c, 0x61, 0x73, 0x74, 0x20, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x20, 0x77, 0x69, 0x6c, 0x6c, 0x20,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x2c, 0x20, 0x69, 0x66, 0x20, 0x61, 0x6e, 0x79, 0x2e, 0x32, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x2e, 0x61, 0x69, 0x2a,
	0x01, 0x02, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x61, 0x6c, 0x2d, 0x61, 0x69, 0x2f, 0x76, 0x31, 0x30, 0x2d, 0x67, 0x6f, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_base_v1_v10_proto_rawDescOnce sync.Once
	file_base_v1_v10_proto_rawDescData []byte
)

func file_base_v1_v10_proto_rawDescGZIP() []byte {
	file_base_v1_v10_proto_rawDescOnce.Do(func() {
		file_base_v1_v10_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_base_v1_v10_proto_rawDesc), len(file_base_v1_v10_proto_rawDesc)))
	})
	return file_base_v1_v10_proto_rawDescData
}

var file_base_v1_v10_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_base_v1_v10_proto_goTypes = []any{
	(*RegisterVehicleRequest)(nil), // 0: v10proto.base.v1.RegisterVehicleRequest
	(*DoRequest)(nil),              // 1: v10proto.base.v1.DoRequest
	(*DoResponse)(nil),             // 2: v10proto.base.v1.DoResponse
	(*emptypb.Empty)(nil),          // 3: google.protobuf.Empty
}
var file_base_v1_v10_proto_depIdxs = []int32{
	1, // 0: v10proto.base.v1.V10.Do:input_type -> v10proto.base.v1.DoRequest
	0, // 1: v10proto.base.v1.V10.RegisterVehicle:input_type -> v10proto.base.v1.RegisterVehicleRequest
	2, // 2: v10proto.base.v1.V10.Do:output_type -> v10proto.base.v1.DoResponse
	3, // 3: v10proto.base.v1.V10.RegisterVehicle:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_base_v1_v10_proto_init() }
func file_base_v1_v10_proto_init() {
	if File_base_v1_v10_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_base_v1_v10_proto_rawDesc), len(file_base_v1_v10_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_base_v1_v10_proto_goTypes,
		DependencyIndexes: file_base_v1_v10_proto_depIdxs,
		MessageInfos:      file_base_v1_v10_proto_msgTypes,
	}.Build()
	File_base_v1_v10_proto = out.File
	file_base_v1_v10_proto_goTypes = nil
	file_base_v1_v10_proto_depIdxs = nil
}
