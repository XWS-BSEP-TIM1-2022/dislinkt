// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: user/userModel.proto

package user

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname     string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Gender      bool   `protobuf:"varint,6,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthDate   string `protobuf:"bytes,7,opt,name=birthDate,proto3" json:"birthDate,omitempty"`
	Username    string `protobuf:"bytes,8,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	Bio         string `protobuf:"bytes,10,opt,name=bio,proto3" json:"bio,omitempty"`
	Skills      string `protobuf:"bytes,11,opt,name=skills,proto3" json:"skills,omitempty"`
	Interests   string `protobuf:"bytes,12,opt,name=interests,proto3" json:"interests,omitempty"`
	Private     bool   `protobuf:"varint,13,opt,name=private,proto3" json:"private,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_userModel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_userModel_proto_msgTypes[0]
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
	return file_user_userModel_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetGender() bool {
	if x != nil {
		return x.Gender
	}
	return false
}

func (x *User) GetBirthDate() string {
	if x != nil {
		return x.BirthDate
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetSkills() string {
	if x != nil {
		return x.Skills
	}
	return ""
}

func (x *User) GetInterests() string {
	if x != nil {
		return x.Interests
	}
	return ""
}

func (x *User) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

var File_user_userModel_proto protoreflect.FileDescriptor

var file_user_userModel_proto_rawDesc = []byte{
	0x0a, 0x14, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xcc, 0x02, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x6f,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x62, 0x69, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x58, 0x57, 0x53, 0x2d, 0x42, 0x53,
	0x45, 0x50, 0x2d, 0x54, 0x49, 0x4d, 0x31, 0x2d, 0x32, 0x30, 0x32, 0x32, 0x2f, 0x64, 0x69, 0x73,
	0x6c, 0x69, 0x6e, 0x6b, 0x74, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_userModel_proto_rawDescOnce sync.Once
	file_user_userModel_proto_rawDescData = file_user_userModel_proto_rawDesc
)

func file_user_userModel_proto_rawDescGZIP() []byte {
	file_user_userModel_proto_rawDescOnce.Do(func() {
		file_user_userModel_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_userModel_proto_rawDescData)
	})
	return file_user_userModel_proto_rawDescData
}

var file_user_userModel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_userModel_proto_goTypes = []interface{}{
	(*User)(nil), // 0: user.User
}
var file_user_userModel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_userModel_proto_init() }
func file_user_userModel_proto_init() {
	if File_user_userModel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_userModel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_userModel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_userModel_proto_goTypes,
		DependencyIndexes: file_user_userModel_proto_depIdxs,
		MessageInfos:      file_user_userModel_proto_msgTypes,
	}.Build()
	File_user_userModel_proto = out.File
	file_user_userModel_proto_rawDesc = nil
	file_user_userModel_proto_goTypes = nil
	file_user_userModel_proto_depIdxs = nil
}
