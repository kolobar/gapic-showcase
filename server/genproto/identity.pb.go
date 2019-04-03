// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/showcase/v1alpha3/identity.proto

package genproto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// A user.
type User struct {
	// The resource name of the user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The display_name of the user.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The email address of the user.
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	// The timestamp at which the user was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The latest timestamp at which the user was updated.
	UpdateTime           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b139bdf892a8c744, []int{0}
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

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *User) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

// The request message for the google.showcase.v1alpha3.Identity\CreateUser
// method.
type CreateUserRequest struct {
	// The user to create.
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b139bdf892a8c744, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

// The request message for the google.showcase.v1alpha3.Identity\GetUser
// method.
type GetUserRequest struct {
	// The resource name of the requested user.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b139bdf892a8c744, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for the google.showcase.v1alpha3.Identity\UpdateUser
// method.
type UpdateUserRequest struct {
	// The user to update.
	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// The field mask to determine wich fields are to be updated. If empty, the
	// server will assume all fields are to be updated.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b139bdf892a8c744, []int{3}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateUserRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// The request message for the google.showcase.v1alpha3.Identity\DeleteUser
// method.
type DeleteUserRequest struct {
	// The resource name of the user to delete.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b139bdf892a8c744, []int{4}
}

func (m *DeleteUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteUserRequest.Unmarshal(m, b)
}
func (m *DeleteUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteUserRequest.Marshal(b, m, deterministic)
}
func (m *DeleteUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteUserRequest.Merge(m, src)
}
func (m *DeleteUserRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteUserRequest.Size(m)
}
func (m *DeleteUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteUserRequest proto.InternalMessageInfo

func (m *DeleteUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request message for the google.showcase.v1alpha3.Identity\ListUsers
// method.
type ListUsersRequest struct {
	// The maximum number of users to return. Server may return fewer users
	// than requested. If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The value of google.showcase.v1alpha3.ListUsersResponse.next_page_token
	// returned from the previous call to
	// `google.showcase.v1alpha3.Identity\ListUsers` method.
	PageToken            string   `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersRequest) Reset()         { *m = ListUsersRequest{} }
func (m *ListUsersRequest) String() string { return proto.CompactTextString(m) }
func (*ListUsersRequest) ProtoMessage()    {}
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b139bdf892a8c744, []int{5}
}

func (m *ListUsersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersRequest.Unmarshal(m, b)
}
func (m *ListUsersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersRequest.Marshal(b, m, deterministic)
}
func (m *ListUsersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersRequest.Merge(m, src)
}
func (m *ListUsersRequest) XXX_Size() int {
	return xxx_messageInfo_ListUsersRequest.Size(m)
}
func (m *ListUsersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersRequest proto.InternalMessageInfo

func (m *ListUsersRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListUsersRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The response message for the google.showcase.v1alpha3.Identity\ListUsers
// method.
type ListUsersResponse struct {
	// The list of users.
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in ListUsersRequest.page_token field in the subsequent
	// call to `google.showcase.v1alpha3.Message\ListUsers` method to retrieve the
	// next page of results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListUsersResponse) Reset()         { *m = ListUsersResponse{} }
func (m *ListUsersResponse) String() string { return proto.CompactTextString(m) }
func (*ListUsersResponse) ProtoMessage()    {}
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b139bdf892a8c744, []int{6}
}

func (m *ListUsersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListUsersResponse.Unmarshal(m, b)
}
func (m *ListUsersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListUsersResponse.Marshal(b, m, deterministic)
}
func (m *ListUsersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListUsersResponse.Merge(m, src)
}
func (m *ListUsersResponse) XXX_Size() int {
	return xxx_messageInfo_ListUsersResponse.Size(m)
}
func (m *ListUsersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListUsersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListUsersResponse proto.InternalMessageInfo

func (m *ListUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ListUsersResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "google.showcase.v1alpha3.User")
	proto.RegisterType((*CreateUserRequest)(nil), "google.showcase.v1alpha3.CreateUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "google.showcase.v1alpha3.GetUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "google.showcase.v1alpha3.UpdateUserRequest")
	proto.RegisterType((*DeleteUserRequest)(nil), "google.showcase.v1alpha3.DeleteUserRequest")
	proto.RegisterType((*ListUsersRequest)(nil), "google.showcase.v1alpha3.ListUsersRequest")
	proto.RegisterType((*ListUsersResponse)(nil), "google.showcase.v1alpha3.ListUsersResponse")
}

func init() {
	proto.RegisterFile("google/showcase/v1alpha3/identity.proto", fileDescriptor_b139bdf892a8c744)
}

var fileDescriptor_b139bdf892a8c744 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcb, 0x4e, 0x1b, 0x49,
	0x14, 0x55, 0xfb, 0x31, 0x03, 0xd7, 0x33, 0x30, 0x2e, 0xa1, 0xc1, 0x34, 0x8f, 0xb1, 0x7a, 0x01,
	0x96, 0x99, 0xe9, 0x9e, 0x31, 0x68, 0x22, 0x88, 0x22, 0xc5, 0xe4, 0x81, 0x22, 0x25, 0x08, 0x39,
	0x90, 0x45, 0x36, 0x56, 0xd1, 0x2e, 0xec, 0x12, 0xdd, 0x5d, 0x4d, 0x57, 0x99, 0x04, 0x08, 0x9b,
	0x24, 0x7f, 0x90, 0xfc, 0x49, 0xfe, 0x22, 0xdb, 0xec, 0x58, 0xb1, 0xc8, 0x2a, 0x3f, 0x10, 0x29,
	0xab, 0xa8, 0xaa, 0xba, 0xfd, 0x8c, 0x31, 0x52, 0x56, 0x6d, 0xdf, 0x7b, 0xee, 0xbd, 0xa7, 0x4e,
	0x9d, 0xdb, 0x0d, 0x2b, 0x4d, 0xc6, 0x9a, 0x1e, 0x71, 0x78, 0x8b, 0xbd, 0x70, 0x31, 0x27, 0xce,
	0xc9, 0x7f, 0xd8, 0x0b, 0x5b, 0x78, 0xcd, 0xa1, 0x0d, 0x12, 0x08, 0x2a, 0x4e, 0xed, 0x30, 0x62,
	0x82, 0xa1, 0x82, 0x06, 0xda, 0x09, 0xd0, 0x4e, 0x80, 0xe6, 0x42, 0xdc, 0x02, 0x87, 0xd4, 0xc1,
	0x41, 0xc0, 0x04, 0x16, 0x94, 0x05, 0x5c, 0xd7, 0x99, 0xb3, 0x3d, 0x59, 0xd7, 0xa3, 0x24, 0x10,
	0x71, 0xe2, 0xaf, 0x9e, 0xc4, 0x21, 0x25, 0x5e, 0xa3, 0x7e, 0x40, 0x5a, 0xf8, 0x84, 0xb2, 0x28,
	0x06, 0xcc, 0xf5, 0x00, 0x22, 0xc2, 0x59, 0x3b, 0x72, 0x49, 0x9c, 0x9a, 0x8f, 0x53, 0xea, 0xdf,
	0x41, 0xfb, 0xd0, 0x21, 0x7e, 0x98, 0x30, 0x35, 0x8b, 0x83, 0x49, 0xdd, 0xdd, 0xc7, 0xfc, 0x68,
	0x60, 0x74, 0x07, 0x21, 0xa8, 0x4f, 0xb8, 0xc0, 0x7e, 0xa8, 0x01, 0xd6, 0x57, 0x03, 0x32, 0xfb,
	0x9c, 0x44, 0xa8, 0x04, 0x99, 0x00, 0xfb, 0xa4, 0x60, 0x14, 0x8d, 0xd2, 0xe4, 0xd6, 0xcc, 0x97,
	0x6a, 0x1e, 0xa6, 0xdb, 0x9c, 0x44, 0xdc, 0x39, 0x97, 0x8f, 0x3a, 0x6d, 0x5c, 0xd4, 0x14, 0x02,
	0x2d, 0xc3, 0x6f, 0x0d, 0xca, 0x43, 0x0f, 0x9f, 0xd6, 0x55, 0x45, 0x4a, 0x55, 0xa4, 0xaf, 0xaa,
	0xa9, 0x5a, 0x2e, 0x4e, 0xec, 0x48, 0xdc, 0x1c, 0x64, 0x89, 0x8f, 0xa9, 0x57, 0x48, 0x77, 0x01,
	0x3a, 0x82, 0xee, 0x42, 0xce, 0x8d, 0x08, 0x16, 0xa4, 0x2e, 0xf9, 0x14, 0x32, 0x45, 0xa3, 0x94,
	0xab, 0x98, 0x76, 0x2c, 0x7c, 0x42, 0xd6, 0xde, 0x4b, 0xc8, 0xca, 0xe2, 0x74, 0x0d, 0x74, 0x8d,
	0x8c, 0xca, 0x0e, 0xed, 0xb0, 0xd1, 0xe9, 0x90, 0xbd, 0x61, 0x07, 0x5d, 0x23, 0xa3, 0xd6, 0x36,
	0xe4, 0xef, 0xa9, 0x7e, 0xf2, 0xf8, 0x35, 0x72, 0xdc, 0x26, 0x5c, 0xa0, 0x0a, 0x64, 0xe4, 0x69,
	0x95, 0x0a, 0xb9, 0xca, 0x92, 0x3d, 0xca, 0x0a, 0xb6, 0x2a, 0x52, 0x58, 0xeb, 0x5f, 0x98, 0xda,
	0x26, 0xa2, 0xb7, 0xcb, 0x52, 0x9f, 0x96, 0x70, 0x55, 0x4d, 0x7d, 0xab, 0x66, 0x74, 0x85, 0x8c,
	0x5b, 0x6f, 0x0d, 0xc8, 0xef, 0x2b, 0x26, 0x3f, 0x39, 0x1b, 0xdd, 0xee, 0xc8, 0x20, 0x2f, 0x5d,
	0x5d, 0xc5, 0x8f, 0x64, 0x78, 0x28, 0x7d, 0xf1, 0x04, 0xf3, 0xa3, 0x44, 0x01, 0xf9, 0xdb, 0x5a,
	0x83, 0xfc, 0x7d, 0xe2, 0x91, 0x7e, 0x16, 0xe3, 0xb8, 0xef, 0xc0, 0x1f, 0x8f, 0x29, 0x57, 0xc7,
	0xe5, 0x49, 0xcd, 0x3c, 0x4c, 0x86, 0xb8, 0x49, 0xea, 0x9c, 0x9e, 0xe9, 0xc2, 0x6c, 0x6d, 0x42,
	0x06, 0x9e, 0xd2, 0x33, 0x82, 0x16, 0x01, 0x54, 0x52, 0xb0, 0x23, 0x12, 0x68, 0xb3, 0xd4, 0x14,
	0x7c, 0x4f, 0x06, 0xac, 0x63, 0xc8, 0xf7, 0xf4, 0xe3, 0x21, 0x0b, 0x38, 0x41, 0xeb, 0x90, 0x55,
	0xde, 0x2b, 0x18, 0xc5, 0xf4, 0x0d, 0xb4, 0xd0, 0x60, 0xb4, 0x0c, 0xd3, 0x01, 0x79, 0x29, 0xea,
	0x43, 0xe3, 0x7e, 0x97, 0xe1, 0xdd, 0x64, 0x64, 0xe5, 0x43, 0x16, 0x26, 0x1e, 0xc5, 0x3b, 0x8f,
	0xde, 0x1b, 0x00, 0x5d, 0x1f, 0xa0, 0xd5, 0xd1, 0xa3, 0x86, 0xdc, 0x62, 0x8e, 0xe1, 0x65, 0x6d,
	0x5c, 0x56, 0x17, 0x24, 0x35, 0xbb, 0x77, 0x5f, 0xfe, 0x56, 0x11, 0xb5, 0x07, 0xaf, 0x3f, 0x7d,
	0x7e, 0x97, 0x9a, 0xb1, 0xa6, 0xbb, 0xaf, 0x22, 0x75, 0x90, 0x4d, 0xa3, 0x8c, 0x5e, 0xc1, 0xaf,
	0xb1, 0xa9, 0x50, 0x69, 0xf4, 0x94, 0x7e, 0xdf, 0x8d, 0xe5, 0xb3, 0x72, 0x59, 0x55, 0x97, 0xa8,
	0xe6, 0x9a, 0xa8, 0xd0, 0x9d, 0x7b, 0x2e, 0xc3, 0x77, 0xf4, 0xbe, 0x97, 0x2f, 0xd0, 0x1b, 0x03,
	0xa0, 0x6b, 0xd0, 0xeb, 0x44, 0x19, 0xb2, 0xf1, 0x58, 0x12, 0x25, 0x35, 0xdd, 0xaa, 0x2c, 0xf6,
	0x4c, 0x57, 0x92, 0xf4, 0x51, 0x90, 0x1a, 0x9c, 0x01, 0x74, 0xfd, 0x79, 0x1d, 0x89, 0x21, 0x17,
	0x9b, 0x7f, 0x0e, 0xad, 0xc0, 0x03, 0xf9, 0xde, 0x1c, 0x50, 0xa0, 0x3c, 0x5a, 0x81, 0x0b, 0x98,
	0xec, 0xd8, 0x12, 0x95, 0x47, 0x8f, 0x1e, 0xdc, 0x05, 0x73, 0xf5, 0x46, 0x58, 0xed, 0x73, 0x6b,
	0x56, 0xf1, 0xc8, 0xa3, 0x41, 0x07, 0x98, 0xf9, 0x8f, 0xd5, 0x29, 0x8f, 0xb9, 0xd8, 0x6b, 0x31,
	0x2e, 0x36, 0x6f, 0xad, 0xff, 0xbf, 0xb1, 0xf5, 0x0c, 0x16, 0x5c, 0xe6, 0x8f, 0xec, 0xbe, 0x6b,
	0x3c, 0x5f, 0x6f, 0x52, 0xd1, 0x6a, 0x1f, 0xd8, 0x2e, 0xf3, 0x1d, 0x0d, 0xc3, 0x21, 0xe5, 0x4e,
	0x13, 0x87, 0xd4, 0xfd, 0xa7, 0xf3, 0xd9, 0xe3, 0x24, 0x3a, 0x21, 0x91, 0xd3, 0x24, 0x81, 0xd6,
	0xe6, 0x17, 0xf5, 0x58, 0xfb, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x54, 0xd0, 0xa2, 0x20, 0x07,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IdentityClient is the client API for Identity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IdentityClient interface {
	// Creates a user.
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Retrieves the User with the given uri.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	// Updates a user.
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	// Deletes a user, their profile, and all of their authored messages.
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Lists all users.
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
}

type identityClient struct {
	cc *grpc.ClientConn
}

func NewIdentityClient(cc *grpc.ClientConn) IdentityClient {
	return &identityClient{cc}
}

func (c *identityClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/google.showcase.v1alpha3.Identity/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/google.showcase.v1alpha3.Identity/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/google.showcase.v1alpha3.Identity/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.showcase.v1alpha3.Identity/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/google.showcase.v1alpha3.Identity/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServer is the server API for Identity service.
type IdentityServer interface {
	// Creates a user.
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// Retrieves the User with the given uri.
	GetUser(context.Context, *GetUserRequest) (*User, error)
	// Updates a user.
	UpdateUser(context.Context, *UpdateUserRequest) (*User, error)
	// Deletes a user, their profile, and all of their authored messages.
	DeleteUser(context.Context, *DeleteUserRequest) (*empty.Empty, error)
	// Lists all users.
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
}

// UnimplementedIdentityServer can be embedded to have forward compatible implementations.
type UnimplementedIdentityServer struct {
}

func (*UnimplementedIdentityServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedIdentityServer) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedIdentityServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (*UnimplementedIdentityServer) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (*UnimplementedIdentityServer) ListUsers(ctx context.Context, req *ListUsersRequest) (*ListUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}

func RegisterIdentityServer(s *grpc.Server, srv IdentityServer) {
	s.RegisterService(&_Identity_serviceDesc, srv)
}

func _Identity_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1alpha3.Identity/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1alpha3.Identity/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1alpha3.Identity/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1alpha3.Identity/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Identity_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.showcase.v1alpha3.Identity/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Identity_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.showcase.v1alpha3.Identity",
	HandlerType: (*IdentityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Identity_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Identity_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Identity_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _Identity_DeleteUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _Identity_ListUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/showcase/v1alpha3/identity.proto",
}
