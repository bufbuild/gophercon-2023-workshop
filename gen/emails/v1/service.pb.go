// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: emails/v1/service.proto

package emailsv1

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

type UserEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	EmailAddress string `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	Verified     bool   `protobuf:"varint,3,opt,name=verified,proto3" json:"verified,omitempty"`
}

func (x *UserEmail) Reset() {
	*x = UserEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emails_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserEmail) ProtoMessage() {}

func (x *UserEmail) ProtoReflect() protoreflect.Message {
	mi := &file_emails_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserEmail.ProtoReflect.Descriptor instead.
func (*UserEmail) Descriptor() ([]byte, []int) {
	return file_emails_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserEmail) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserEmail) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *UserEmail) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

type GetEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetEmailRequest) Reset() {
	*x = GetEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emails_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailRequest) ProtoMessage() {}

func (x *GetEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emails_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailRequest.ProtoReflect.Descriptor instead.
func (*GetEmailRequest) Descriptor() ([]byte, []int) {
	return file_emails_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetEmailRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserEmail *UserEmail `protobuf:"bytes,1,opt,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
}

func (x *GetEmailResponse) Reset() {
	*x = GetEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emails_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmailResponse) ProtoMessage() {}

func (x *GetEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emails_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmailResponse.ProtoReflect.Descriptor instead.
func (*GetEmailResponse) Descriptor() ([]byte, []int) {
	return file_emails_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetEmailResponse) GetUserEmail() *UserEmail {
	if x != nil {
		return x.UserEmail
	}
	return nil
}

type UpdateEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NewAddress string `protobuf:"bytes,2,opt,name=new_address,json=newAddress,proto3" json:"new_address,omitempty"`
}

func (x *UpdateEmailRequest) Reset() {
	*x = UpdateEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emails_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailRequest) ProtoMessage() {}

func (x *UpdateEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_emails_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmailRequest) Descriptor() ([]byte, []int) {
	return file_emails_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateEmailRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateEmailRequest) GetNewAddress() string {
	if x != nil {
		return x.NewAddress
	}
	return ""
}

type UpdateEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEmailResponse) Reset() {
	*x = UpdateEmailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_emails_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmailResponse) ProtoMessage() {}

func (x *UpdateEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_emails_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmailResponse.ProtoReflect.Descriptor instead.
func (*UpdateEmailResponse) Descriptor() ([]byte, []int) {
	return file_emails_v1_service_proto_rawDescGZIP(), []int{4}
}

var File_emails_v1_service_proto protoreflect.FileDescriptor

var file_emails_v1_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x65, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x4e, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x65, 0x77, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa1, 0x01, 0x0a, 0x0c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a,
	0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xa6, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x2d, 0x32, 0x30, 0x32,
	0x33, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x09, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x15, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_emails_v1_service_proto_rawDescOnce sync.Once
	file_emails_v1_service_proto_rawDescData = file_emails_v1_service_proto_rawDesc
)

func file_emails_v1_service_proto_rawDescGZIP() []byte {
	file_emails_v1_service_proto_rawDescOnce.Do(func() {
		file_emails_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_emails_v1_service_proto_rawDescData)
	})
	return file_emails_v1_service_proto_rawDescData
}

var file_emails_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_emails_v1_service_proto_goTypes = []interface{}{
	(*UserEmail)(nil),           // 0: emails.v1.UserEmail
	(*GetEmailRequest)(nil),     // 1: emails.v1.GetEmailRequest
	(*GetEmailResponse)(nil),    // 2: emails.v1.GetEmailResponse
	(*UpdateEmailRequest)(nil),  // 3: emails.v1.UpdateEmailRequest
	(*UpdateEmailResponse)(nil), // 4: emails.v1.UpdateEmailResponse
}
var file_emails_v1_service_proto_depIdxs = []int32{
	0, // 0: emails.v1.GetEmailResponse.user_email:type_name -> emails.v1.UserEmail
	1, // 1: emails.v1.EmailService.GetEmail:input_type -> emails.v1.GetEmailRequest
	3, // 2: emails.v1.EmailService.UpdateEmail:input_type -> emails.v1.UpdateEmailRequest
	2, // 3: emails.v1.EmailService.GetEmail:output_type -> emails.v1.GetEmailResponse
	4, // 4: emails.v1.EmailService.UpdateEmail:output_type -> emails.v1.UpdateEmailResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_emails_v1_service_proto_init() }
func file_emails_v1_service_proto_init() {
	if File_emails_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_emails_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserEmail); i {
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
		file_emails_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailRequest); i {
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
		file_emails_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmailResponse); i {
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
		file_emails_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailRequest); i {
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
		file_emails_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmailResponse); i {
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
			RawDescriptor: file_emails_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_emails_v1_service_proto_goTypes,
		DependencyIndexes: file_emails_v1_service_proto_depIdxs,
		MessageInfos:      file_emails_v1_service_proto_msgTypes,
	}.Build()
	File_emails_v1_service_proto = out.File
	file_emails_v1_service_proto_rawDesc = nil
	file_emails_v1_service_proto_goTypes = nil
	file_emails_v1_service_proto_depIdxs = nil
}