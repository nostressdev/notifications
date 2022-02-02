// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/models.proto

package proto

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type DeviceType int32

const (
	DeviceType_IOS     DeviceType = 0
	DeviceType_ANDROID DeviceType = 1
	DeviceType_WEB     DeviceType = 2
	DeviceType_HUAWEI  DeviceType = 3
	DeviceType_EMAIL   DeviceType = 4 // more platforms can be added, see https://documentation.onesignal.com/reference/add-a-device
)

// Enum value maps for DeviceType.
var (
	DeviceType_name = map[int32]string{
		0: "IOS",
		1: "ANDROID",
		2: "WEB",
		3: "HUAWEI",
		4: "EMAIL",
	}
	DeviceType_value = map[string]int32{
		"IOS":     0,
		"ANDROID": 1,
		"WEB":     2,
		"HUAWEI":  3,
		"EMAIL":   4,
	}
)

func (x DeviceType) Enum() *DeviceType {
	p := new(DeviceType)
	*p = x
	return p
}

func (x DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_models_proto_enumTypes[0].Descriptor()
}

func (DeviceType) Type() protoreflect.EnumType {
	return &file_api_models_proto_enumTypes[0]
}

func (x DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceType.Descriptor instead.
func (DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{0}
}

type DeviceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceType DeviceType `protobuf:"varint,1,opt,name=DeviceType,proto3,enum=notifications.DeviceType" json:"DeviceType,omitempty"`
	Identifier string     `protobuf:"bytes,2,opt,name=Identifier,proto3" json:"Identifier,omitempty"` // push token
	Language   string     `protobuf:"bytes,3,opt,name=Language,proto3" json:"Language,omitempty"`     // more info can be added
}

func (x *DeviceInfo) Reset() {
	*x = DeviceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceInfo) ProtoMessage() {}

func (x *DeviceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceInfo.ProtoReflect.Descriptor instead.
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceInfo) GetDeviceType() DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return DeviceType_IOS
}

func (x *DeviceInfo) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *DeviceInfo) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID   string      `protobuf:"bytes,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	DeviceInfo *DeviceInfo `protobuf:"bytes,2,opt,name=DeviceInfo,proto3" json:"DeviceInfo,omitempty"`
	AccountID  string      `protobuf:"bytes,3,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetDeviceID() string {
	if x != nil {
		return x.DeviceID
	}
	return ""
}

func (x *Device) GetDeviceInfo() *DeviceInfo {
	if x != nil {
		return x.DeviceInfo
	}
	return nil
}

func (x *Device) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string    `protobuf:"bytes,1,opt,name=AccountID,proto3" json:"AccountID,omitempty"`
	Devices   []*Device `protobuf:"bytes,2,rep,name=Devices,proto3" json:"Devices,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *User) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type JSONObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]string `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *JSONObject) Reset() {
	*x = JSONObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JSONObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JSONObject) ProtoMessage() {}

func (x *JSONObject) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JSONObject.ProtoReflect.Descriptor instead.
func (*JSONObject) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{3}
}

func (x *JSONObject) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// for each language
	Title    map[string]string `protobuf:"bytes,1,rep,name=Title,proto3" json:"Title,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body     map[string]string `protobuf:"bytes,2,rep,name=Body,proto3" json:"Body,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ImageURL string            `protobuf:"bytes,3,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	// for each platform
	ClickAction map[string]string      `protobuf:"bytes,4,rep,name=ClickAction,proto3" json:"ClickAction,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data        map[string]*JSONObject `protobuf:"bytes,5,rep,name=Data,proto3" json:"Data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_models_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_api_models_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_api_models_proto_rawDescGZIP(), []int{4}
}

func (x *Notification) GetTitle() map[string]string {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *Notification) GetBody() map[string]string {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *Notification) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *Notification) GetClickAction() map[string]string {
	if x != nil {
		return x.ClickAction
	}
	return nil
}

func (x *Notification) GetData() map[string]*JSONObject {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_models_proto protoreflect.FileDescriptor

var file_api_models_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x0a, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x7d, 0x0a, 0x06, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12,
	0x39, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x5e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x25, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x07, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x07, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x82, 0x01, 0x0a, 0x0a, 0x4a, 0x53, 0x4f,
	0x4e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4a, 0x53, 0x4f, 0x4e, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb5, 0x04,
	0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x04,
	0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x55, 0x52, 0x4c, 0x12, 0x4e, 0x0a, 0x0b, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x38,
	0x0a, 0x0a, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x37, 0x0a, 0x09, 0x42, 0x6f, 0x64, 0x79,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x52, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x4a, 0x53, 0x4f, 0x4e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x42, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x57, 0x45, 0x42,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x55, 0x41, 0x57, 0x45, 0x49, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x04, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_models_proto_rawDescOnce sync.Once
	file_api_models_proto_rawDescData = file_api_models_proto_rawDesc
)

func file_api_models_proto_rawDescGZIP() []byte {
	file_api_models_proto_rawDescOnce.Do(func() {
		file_api_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_models_proto_rawDescData)
	})
	return file_api_models_proto_rawDescData
}

var file_api_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_models_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_models_proto_goTypes = []interface{}{
	(DeviceType)(0),      // 0: notifications.DeviceType
	(*DeviceInfo)(nil),   // 1: notifications.DeviceInfo
	(*Device)(nil),       // 2: notifications.Device
	(*User)(nil),         // 3: notifications.User
	(*JSONObject)(nil),   // 4: notifications.JSONObject
	(*Notification)(nil), // 5: notifications.Notification
	nil,                  // 6: notifications.JSONObject.ValueEntry
	nil,                  // 7: notifications.Notification.TitleEntry
	nil,                  // 8: notifications.Notification.BodyEntry
	nil,                  // 9: notifications.Notification.ClickActionEntry
	nil,                  // 10: notifications.Notification.DataEntry
}
var file_api_models_proto_depIdxs = []int32{
	0,  // 0: notifications.DeviceInfo.DeviceType:type_name -> notifications.DeviceType
	1,  // 1: notifications.Device.DeviceInfo:type_name -> notifications.DeviceInfo
	2,  // 2: notifications.User.Devices:type_name -> notifications.Device
	6,  // 3: notifications.JSONObject.Value:type_name -> notifications.JSONObject.ValueEntry
	7,  // 4: notifications.Notification.Title:type_name -> notifications.Notification.TitleEntry
	8,  // 5: notifications.Notification.Body:type_name -> notifications.Notification.BodyEntry
	9,  // 6: notifications.Notification.ClickAction:type_name -> notifications.Notification.ClickActionEntry
	10, // 7: notifications.Notification.Data:type_name -> notifications.Notification.DataEntry
	4,  // 8: notifications.Notification.DataEntry.value:type_name -> notifications.JSONObject
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_api_models_proto_init() }
func file_api_models_proto_init() {
	if File_api_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceInfo); i {
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
		file_api_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_api_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_api_models_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JSONObject); i {
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
		file_api_models_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_api_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_models_proto_goTypes,
		DependencyIndexes: file_api_models_proto_depIdxs,
		EnumInfos:         file_api_models_proto_enumTypes,
		MessageInfos:      file_api_models_proto_msgTypes,
	}.Build()
	File_api_models_proto = out.File
	file_api_models_proto_rawDesc = nil
	file_api_models_proto_goTypes = nil
	file_api_models_proto_depIdxs = nil
}
