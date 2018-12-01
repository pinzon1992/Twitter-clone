// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userpb/user.proto

package userpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	Follows              []int32  `protobuf:"varint,6,rep,packed,name=Follows,proto3" json:"Follows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetFollows() []int32 {
	if m != nil {
		return m.Follows
	}
	return nil
}

type UserListFields struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	UserType             string   `protobuf:"bytes,4,opt,name=UserType,proto3" json:"UserType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListFields) Reset()         { *m = UserListFields{} }
func (m *UserListFields) String() string { return proto.CompactTextString(m) }
func (*UserListFields) ProtoMessage()    {}
func (*UserListFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{1}
}

func (m *UserListFields) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListFields.Unmarshal(m, b)
}
func (m *UserListFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListFields.Marshal(b, m, deterministic)
}
func (m *UserListFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListFields.Merge(m, src)
}
func (m *UserListFields) XXX_Size() int {
	return xxx_messageInfo_UserListFields.Size(m)
}
func (m *UserListFields) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListFields.DiscardUnknown(m)
}

var xxx_messageInfo_UserListFields proto.InternalMessageInfo

func (m *UserListFields) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserListFields) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserListFields) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserListFields) GetUserType() string {
	if m != nil {
		return m.UserType
	}
	return ""
}

type UserList struct {
	List                 []*UserListFields `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{2}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetList() []*UserListFields {
	if m != nil {
		return m.List
	}
	return nil
}

type Login struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Login) Reset()         { *m = Login{} }
func (m *Login) String() string { return proto.CompactTextString(m) }
func (*Login) ProtoMessage()    {}
func (*Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{3}
}

func (m *Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Login.Unmarshal(m, b)
}
func (m *Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Login.Marshal(b, m, deterministic)
}
func (m *Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Login.Merge(m, src)
}
func (m *Login) XXX_Size() int {
	return xxx_messageInfo_Login.Size(m)
}
func (m *Login) XXX_DiscardUnknown() {
	xxx_messageInfo_Login.DiscardUnknown(m)
}

var xxx_messageInfo_Login proto.InternalMessageInfo

func (m *Login) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type AddUserParameters struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddUserParameters) Reset()         { *m = AddUserParameters{} }
func (m *AddUserParameters) String() string { return proto.CompactTextString(m) }
func (*AddUserParameters) ProtoMessage()    {}
func (*AddUserParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{4}
}

func (m *AddUserParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddUserParameters.Unmarshal(m, b)
}
func (m *AddUserParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddUserParameters.Marshal(b, m, deterministic)
}
func (m *AddUserParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddUserParameters.Merge(m, src)
}
func (m *AddUserParameters) XXX_Size() int {
	return xxx_messageInfo_AddUserParameters.Size(m)
}
func (m *AddUserParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_AddUserParameters.DiscardUnknown(m)
}

var xxx_messageInfo_AddUserParameters proto.InternalMessageInfo

func (m *AddUserParameters) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *AddUserParameters) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *AddUserParameters) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AddUserParameters) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginDetails struct {
	Email                string   `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginDetails) Reset()         { *m = LoginDetails{} }
func (m *LoginDetails) String() string { return proto.CompactTextString(m) }
func (*LoginDetails) ProtoMessage()    {}
func (*LoginDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{5}
}

func (m *LoginDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginDetails.Unmarshal(m, b)
}
func (m *LoginDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginDetails.Marshal(b, m, deterministic)
}
func (m *LoginDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginDetails.Merge(m, src)
}
func (m *LoginDetails) XXX_Size() int {
	return xxx_messageInfo_LoginDetails.Size(m)
}
func (m *LoginDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginDetails.DiscardUnknown(m)
}

var xxx_messageInfo_LoginDetails proto.InternalMessageInfo

func (m *LoginDetails) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginDetails) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type UserId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserId) Reset()         { *m = UserId{} }
func (m *UserId) String() string { return proto.CompactTextString(m) }
func (*UserId) ProtoMessage()    {}
func (*UserId) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{6}
}

func (m *UserId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserId.Unmarshal(m, b)
}
func (m *UserId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserId.Marshal(b, m, deterministic)
}
func (m *UserId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserId.Merge(m, src)
}
func (m *UserId) XXX_Size() int {
	return xxx_messageInfo_UserId.Size(m)
}
func (m *UserId) XXX_DiscardUnknown() {
	xxx_messageInfo_UserId.DiscardUnknown(m)
}

var xxx_messageInfo_UserId proto.InternalMessageInfo

func (m *UserId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type FollowerParameters struct {
	UserId               int32    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FollowerId           int32    `protobuf:"varint,2,opt,name=FollowerId,proto3" json:"FollowerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FollowerParameters) Reset()         { *m = FollowerParameters{} }
func (m *FollowerParameters) String() string { return proto.CompactTextString(m) }
func (*FollowerParameters) ProtoMessage()    {}
func (*FollowerParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{7}
}

func (m *FollowerParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FollowerParameters.Unmarshal(m, b)
}
func (m *FollowerParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FollowerParameters.Marshal(b, m, deterministic)
}
func (m *FollowerParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FollowerParameters.Merge(m, src)
}
func (m *FollowerParameters) XXX_Size() int {
	return xxx_messageInfo_FollowerParameters.Size(m)
}
func (m *FollowerParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_FollowerParameters.DiscardUnknown(m)
}

var xxx_messageInfo_FollowerParameters proto.InternalMessageInfo

func (m *FollowerParameters) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *FollowerParameters) GetFollowerId() int32 {
	if m != nil {
		return m.FollowerId
	}
	return 0
}

type Status struct {
	ResponseStatus       bool     `protobuf:"varint,1,opt,name=ResponseStatus,proto3" json:"ResponseStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eaba19047509177, []int{8}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetResponseStatus() bool {
	if m != nil {
		return m.ResponseStatus
	}
	return false
}

func init() {
	proto.RegisterType((*User)(nil), "userpb.User")
	proto.RegisterType((*UserListFields)(nil), "userpb.UserListFields")
	proto.RegisterType((*UserList)(nil), "userpb.UserList")
	proto.RegisterType((*Login)(nil), "userpb.Login")
	proto.RegisterType((*AddUserParameters)(nil), "userpb.AddUserParameters")
	proto.RegisterType((*LoginDetails)(nil), "userpb.LoginDetails")
	proto.RegisterType((*UserId)(nil), "userpb.UserId")
	proto.RegisterType((*FollowerParameters)(nil), "userpb.FollowerParameters")
	proto.RegisterType((*Status)(nil), "userpb.Status")
}

func init() { proto.RegisterFile("userpb/user.proto", fileDescriptor_2eaba19047509177) }

var fileDescriptor_2eaba19047509177 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x6e, 0x92, 0x26, 0x76, 0xcf, 0xd6, 0xe0, 0x1e, 0x4a, 0x89, 0x41, 0xa4, 0xcc, 0x85, 0x14,
	0x85, 0xba, 0xac, 0xa0, 0x20, 0x22, 0xee, 0xa2, 0x95, 0x40, 0x91, 0x25, 0x75, 0x1f, 0x20, 0xeb,
	0x8c, 0x25, 0x90, 0x36, 0x25, 0x67, 0xba, 0x4b, 0xaf, 0x7c, 0x11, 0x1f, 0xd2, 0x47, 0x90, 0xc9,
	0x64, 0xda, 0x24, 0xa5, 0x5e, 0xc8, 0x5e, 0x4d, 0xcf, 0xcf, 0xf7, 0xcd, 0x77, 0xbe, 0x33, 0x0d,
	0x9c, 0x6d, 0x48, 0x14, 0xeb, 0xdb, 0xd7, 0xea, 0x98, 0xac, 0x8b, 0x5c, 0xe6, 0xe8, 0xe9, 0x14,
	0xfb, 0x6d, 0x41, 0xf7, 0x86, 0x44, 0x81, 0x3e, 0xd8, 0x11, 0x0f, 0xac, 0x91, 0x35, 0x76, 0x63,
	0x3b, 0xe2, 0xf8, 0x0c, 0x4e, 0xa6, 0x69, 0x41, 0xf2, 0x5b, 0xb2, 0x14, 0x81, 0x3d, 0xb2, 0xc6,
	0x27, 0xf1, 0x3e, 0x81, 0x21, 0xf4, 0x66, 0x49, 0x55, 0x74, 0xca, 0xe2, 0x2e, 0xc6, 0x01, 0xb8,
	0x5f, 0x96, 0x49, 0x9a, 0x05, 0xdd, 0xb2, 0xa0, 0x03, 0x85, 0xb8, 0x4e, 0x88, 0xee, 0xf3, 0x82,
	0x07, 0xae, 0x46, 0x98, 0x18, 0x03, 0x78, 0x34, 0xcd, 0xb3, 0x2c, 0xbf, 0xa7, 0xc0, 0x1b, 0x39,
	0x63, 0x37, 0x36, 0x21, 0xbb, 0x03, 0x5f, 0xa9, 0x9b, 0xa5, 0x24, 0xa7, 0xa9, 0xc8, 0x38, 0x3d,
	0xa0, 0xce, 0x10, 0x7a, 0x8a, 0xfb, 0xfb, 0x76, 0x2d, 0x2a, 0xa9, 0xbb, 0x98, 0xbd, 0xd5, 0x35,
	0x75, 0x2f, 0xbe, 0x84, 0xae, 0x3a, 0x03, 0x6b, 0xe4, 0x8c, 0x4f, 0x2f, 0x86, 0x13, 0xed, 0xdc,
	0xa4, 0xa9, 0x2b, 0x2e, 0x7b, 0xd8, 0x2b, 0x70, 0x67, 0xf9, 0x22, 0x5d, 0x21, 0x03, 0x57, 0x35,
	0x50, 0x85, 0xea, 0xd7, 0x51, 0xb1, 0x2e, 0xb1, 0x5f, 0x70, 0x76, 0xc9, 0xb9, 0xfa, 0x7d, 0x9d,
	0x14, 0xc9, 0x52, 0x48, 0x51, 0x50, 0x73, 0x1e, 0xeb, 0x5f, 0xf3, 0xd8, 0xc7, 0x7c, 0x77, 0x8e,
	0xf9, 0xde, 0x6d, 0xfa, 0xce, 0x3e, 0x41, 0xbf, 0x54, 0xfb, 0x59, 0xc8, 0x24, 0xcd, 0x68, 0xcf,
	0x60, 0x1d, 0x63, 0xb0, 0x5b, 0x0c, 0x01, 0x78, 0x4a, 0x7f, 0xc4, 0xdb, 0x7b, 0x61, 0x33, 0x40,
	0xbd, 0xc4, 0xc6, 0x74, 0x43, 0xd3, 0x5f, 0x75, 0x1a, 0xf4, 0x73, 0x00, 0xd3, 0x1d, 0xe9, 0x5b,
	0xdc, 0xb8, 0x96, 0x61, 0xe7, 0xe0, 0xcd, 0x65, 0x22, 0x37, 0x84, 0x2f, 0xc0, 0x8f, 0x05, 0xad,
	0xf3, 0x15, 0x09, 0x9d, 0x29, 0x99, 0x7a, 0x71, 0x2b, 0x7b, 0xf1, 0xc7, 0x86, 0x53, 0x45, 0x3e,
	0x17, 0xc5, 0x5d, 0xfa, 0x43, 0xe0, 0x39, 0x38, 0x97, 0x9c, 0xe3, 0x53, 0xb3, 0x88, 0x03, 0xe7,
	0xc3, 0xc6, 0x8e, 0x58, 0x07, 0x3f, 0xc2, 0xf0, 0xab, 0x90, 0x2a, 0xb8, 0xda, 0x96, 0x4e, 0xec,
	0xde, 0xeb, 0xc0, 0x74, 0xd6, 0xdd, 0x3b, 0xc0, 0xbf, 0x37, 0x33, 0x95, 0xff, 0xaf, 0xd0, 0x54,
	0x0f, 0x5d, 0x09, 0x7d, 0x53, 0xd3, 0xda, 0x59, 0x07, 0x3f, 0x40, 0xff, 0x66, 0xf5, 0xf3, 0xff,
	0xd1, 0x4a, 0xb9, 0x69, 0x9d, 0x6f, 0x16, 0x0b, 0x41, 0x32, 0xcd, 0x57, 0x84, 0x7e, 0x5d, 0x63,
	0xc4, 0xc3, 0x27, 0xed, 0xd7, 0xcc, 0x3a, 0xf8, 0x0e, 0x06, 0xd5, 0xdc, 0x86, 0x81, 0xae, 0xb6,
	0x6a, 0xc3, 0x2d, 0xec, 0xe3, 0x86, 0x0b, 0xac, 0x73, 0xeb, 0x95, 0x9f, 0x96, 0x37, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x68, 0xb7, 0x08, 0xdf, 0x6f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Add(ctx context.Context, in *AddUserParameters, opts ...grpc.CallOption) (*User, error)
	GetUserByEmailPassword(ctx context.Context, in *LoginDetails, opts ...grpc.CallOption) (*User, error)
	FollowUser(ctx context.Context, in *FollowerParameters, opts ...grpc.CallOption) (*Status, error)
	UnfollowUser(ctx context.Context, in *FollowerParameters, opts ...grpc.CallOption) (*Status, error)
	GetFollowerSuggestions(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserList, error)
	GetUserFollowersById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Login, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Add(ctx context.Context, in *AddUserParameters, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userpb.UserService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByEmailPassword(ctx context.Context, in *LoginDetails, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userpb.UserService/GetUserByEmailPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FollowUser(ctx context.Context, in *FollowerParameters, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/userpb.UserService/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UnfollowUser(ctx context.Context, in *FollowerParameters, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/userpb.UserService/UnfollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetFollowerSuggestions(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, "/userpb.UserService/GetFollowerSuggestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserFollowersById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*Login, error) {
	out := new(Login)
	err := c.cc.Invoke(ctx, "/userpb.UserService/GetUserFollowersById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Add(context.Context, *AddUserParameters) (*User, error)
	GetUserByEmailPassword(context.Context, *LoginDetails) (*User, error)
	FollowUser(context.Context, *FollowerParameters) (*Status, error)
	UnfollowUser(context.Context, *FollowerParameters) (*Status, error)
	GetFollowerSuggestions(context.Context, *UserId) (*UserList, error)
	GetUserFollowersById(context.Context, *UserId) (*Login, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Add(ctx, req.(*AddUserParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByEmailPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByEmailPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/GetUserByEmailPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByEmailPassword(ctx, req.(*LoginDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FollowUser(ctx, req.(*FollowerParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UnfollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UnfollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/UnfollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UnfollowUser(ctx, req.(*FollowerParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetFollowerSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetFollowerSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/GetFollowerSuggestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetFollowerSuggestions(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserFollowersById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserFollowersById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/GetUserFollowersById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserFollowersById(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "userpb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _UserService_Add_Handler,
		},
		{
			MethodName: "GetUserByEmailPassword",
			Handler:    _UserService_GetUserByEmailPassword_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _UserService_FollowUser_Handler,
		},
		{
			MethodName: "UnfollowUser",
			Handler:    _UserService_UnfollowUser_Handler,
		},
		{
			MethodName: "GetFollowerSuggestions",
			Handler:    _UserService_GetFollowerSuggestions_Handler,
		},
		{
			MethodName: "GetUserFollowersById",
			Handler:    _UserService_GetUserFollowersById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userpb/user.proto",
}