// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/teams.proto

package v2

import (
	context "context"
	fmt "fmt"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/api"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/teams.proto", fileDescriptor_16a78463b445c3d8)
}

var fileDescriptor_16a78463b445c3d8 = []byte{
	// 1013 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0x4d, 0x6f, 0x1b, 0x45,
	0x18, 0xd6, 0xa6, 0xa9, 0x9b, 0x0e, 0x90, 0x26, 0x83, 0x69, 0x8c, 0x55, 0x54, 0x77, 0x54, 0xda,
	0x62, 0xb2, 0x5e, 0xb2, 0xf1, 0x47, 0x6c, 0x0e, 0x60, 0x40, 0x54, 0x20, 0x4a, 0x45, 0x9b, 0x5e,
	0xa8, 0x10, 0x1a, 0xef, 0x8e, 0x37, 0x5b, 0xbc, 0x3b, 0x93, 0x9d, 0x71, 0x42, 0x14, 0x45, 0x42,
	0x41, 0x44, 0xc2, 0x17, 0xa4, 0xf4, 0x0a, 0xe2, 0x44, 0xc5, 0x1f, 0xf0, 0x99, 0x23, 0x3f, 0xa0,
	0xe2, 0xc0, 0x9d, 0x3f, 0xc0, 0x01, 0x71, 0x42, 0x42, 0xb3, 0x9b, 0xec, 0x47, 0xd7, 0x6b, 0x6f,
	0xae, 0x3d, 0xad, 0x76, 0xe6, 0x79, 0x67, 0xde, 0xe7, 0x79, 0x9f, 0x77, 0x66, 0x80, 0x6e, 0x50,
	0x87, 0x51, 0x97, 0xb8, 0x82, 0x6b, 0x78, 0x28, 0xa8, 0x83, 0x05, 0x51, 0x2d, 0x2c, 0xc8, 0x2e,
	0xde, 0xd3, 0x30, 0xb3, 0x35, 0x1b, 0x3b, 0xda, 0x8e, 0xae, 0x09, 0x82, 0x1d, 0x5e, 0x63, 0x1e,
	0x15, 0x14, 0x96, 0x8c, 0x2d, 0xd2, 0xaf, 0x9d, 0xa2, 0x6b, 0x98, 0xd9, 0x35, 0x1b, 0x3b, 0xb5,
	0x1d, 0xbd, 0x7c, 0xc5, 0xa2, 0xd4, 0x1a, 0x10, 0x3f, 0x10, 0xbb, 0x2e, 0x15, 0x58, 0xd8, 0xd4,
	0x3d, 0x89, 0x2b, 0xaf, 0xfa, 0x1f, 0x43, 0xb5, 0x88, 0xab, 0xf2, 0x5d, 0x6c, 0x59, 0xc4, 0xd3,
	0x28, 0xf3, 0x11, 0x13, 0xd0, 0x9d, 0x9c, 0x99, 0x79, 0x64, 0x7b, 0x48, 0xb8, 0x88, 0x67, 0x58,
	0x7e, 0x3b, 0x77, 0x2c, 0x67, 0xd4, 0xe5, 0x24, 0x11, 0xfc, 0xee, 0xc4, 0x60, 0x8f, 0x19, 0x5a,
	0x2c, 0x7f, 0x46, 0x07, 0xb6, 0xb1, 0x97, 0x41, 0xf4, 0x2c, 0x2b, 0xc8, 0x4c, 0x52, 0x2b, 0xe8,
	0xff, 0x14, 0xc1, 0xf9, 0x4d, 0x99, 0x13, 0x7c, 0x3a, 0x07, 0xc0, 0xfb, 0x1e, 0xc1, 0x82, 0xc8,
	0x7f, 0x78, 0xb3, 0x96, 0x25, 0x7e, 0x2d, 0x42, 0xdd, 0x23, 0xdb, 0xe5, 0x5b, 0xf9, 0x80, 0x9c,
	0xa1, 0xbf, 0x95, 0xe3, 0xee, 0x0f, 0x0a, 0x38, 0xef, 0x6b, 0xf0, 0xe8, 0x3b, 0x05, 0x2c, 0x7e,
	0xad, 0x1a, 0xd4, 0x24, 0x2a, 0xc7, 0x0e, 0x1b, 0x10, 0x0e, 0xb9, 0xbe, 0x0d, 0x68, 0xd5, 0x01,
	0x8b, 0x60, 0x7e, 0x80, 0x5d, 0x0b, 0x16, 0xca, 0xf3, 0x1f, 0xdf, 0xbf, 0xfb, 0x29, 0x78, 0x08,
	0x0a, 0x9c, 0x0e, 0x3d, 0x83, 0xc0, 0xcf, 0xca, 0x77, 0xf7, 0x91, 0x6d, 0xa2, 0x4e, 0x05, 0x09,
	0xc2, 0x85, 0x6a, 0x9b, 0x68, 0xb5, 0x82, 0x5c, 0xec, 0x10, 0x39, 0x74, 0x67, 0xaf, 0xb2, 0x49,
	0xb8, 0xa8, 0xc8, 0x1d, 0xe5, 0x38, 0xf3, 0xe8, 0x23, 0x62, 0x08, 0x8e, 0x3a, 0x95, 0x87, 0xa7,
	0x3f, 0x6b, 0xb1, 0x09, 0x1d, 0x7d, 0x71, 0x70, 0x38, 0x2e, 0xbd, 0x08, 0x00, 0x1e, 0x8a, 0xad,
	0x8e, 0x9f, 0xd7, 0xe1, 0xb8, 0xb4, 0x00, 0x0b, 0x86, 0x9f, 0xf9, 0x68, 0x5c, 0x7a, 0x01, 0x5c,
	0xb4, 0xb1, 0x13, 0x4c, 0x8d, 0xc6, 0x25, 0x08, 0x97, 0xc2, 0xdf, 0x4e, 0x00, 0x3a, 0x7c, 0xfa,
	0xd7, 0xe3, 0xb9, 0x15, 0x04, 0x65, 0x89, 0x78, 0xc2, 0xc5, 0x1d, 0xa5, 0x0a, 0x7f, 0x53, 0xc0,
	0xc5, 0x4f, 0x6c, 0x2e, 0x02, 0x8d, 0x6f, 0x64, 0x6b, 0x15, 0x82, 0xa4, 0xa6, 0x37, 0x73, 0xe1,
	0x38, 0x43, 0xf8, 0xb8, 0x7b, 0xe1, 0x44, 0xd0, 0x09, 0x44, 0x0a, 0x70, 0xde, 0x23, 0xd8, 0x4c,
	0xd3, 0x58, 0x82, 0x8b, 0x11, 0x8d, 0x81, 0xcd, 0x85, 0x4f, 0xa2, 0x08, 0x27, 0x90, 0x80, 0xbf,
	0x2b, 0xe0, 0xc2, 0x6d, 0xe2, 0xef, 0x09, 0xaf, 0x67, 0xe7, 0x75, 0x02, 0x91, 0xd9, 0xbf, 0x9e,
	0x03, 0xc5, 0x19, 0xa2, 0x89, 0xdc, 0x97, 0xc1, 0xa5, 0x28, 0xf7, 0xce, 0xbe, 0x6d, 0x1e, 0x24,
	0x08, 0x2c, 0x81, 0x58, 0xc6, 0x72, 0x76, 0x34, 0x2e, 0x5d, 0x82, 0x2f, 0x45, 0x63, 0x16, 0x09,
	0x48, 0xbc, 0x0a, 0x57, 0xd2, 0x24, 0x34, 0x19, 0x01, 0xff, 0x9c, 0x03, 0xe0, 0x01, 0x33, 0x73,
	0x38, 0x3c, 0x42, 0xcd, 0x70, 0x78, 0x1c, 0xc8, 0x19, 0xfa, 0x57, 0x39, 0xee, 0x1e, 0x85, 0x0e,
	0x3f, 0x48, 0x19, 0xfc, 0x2b, 0xdd, 0x06, 0x56, 0x95, 0xa4, 0x0c, 0x7e, 0x3f, 0x34, 0xf8, 0x47,
	0xe5, 0xdb, 0xfb, 0x71, 0x3f, 0x07, 0x1b, 0x9c, 0xcd, 0xd6, 0x2e, 0xd9, 0x0d, 0x8c, 0x3d, 0x51,
	0xd3, 0x05, 0x58, 0x18, 0xfa, 0x8b, 0x66, 0xa8, 0x9a, 0xb0, 0x78, 0x80, 0xf4, 0x85, 0xbd, 0x52,
	0xce, 0x12, 0x56, 0xfa, 0xfc, 0x0f, 0x05, 0x80, 0x0f, 0xc8, 0x80, 0xcc, 0xd6, 0x36, 0x42, 0xcd,
	0xd0, 0x36, 0x0e, 0xe4, 0x0c, 0x89, 0xd9, 0x76, 0x59, 0x80, 0x05, 0xd3, 0x0f, 0xca, 0x43, 0x2d,
	0x40, 0x06, 0x9e, 0xa9, 0x66, 0x7a, 0xe6, 0x3f, 0x05, 0x2c, 0x9f, 0x98, 0xf6, 0x0e, 0x71, 0x7a,
	0xc4, 0xe3, 0x5b, 0x36, 0x83, 0xb5, 0x99, 0x0e, 0x8f, 0xc0, 0x92, 0xa5, 0x76, 0x26, 0x3c, 0x67,
	0xe8, 0x48, 0x49, 0xb0, 0x5d, 0x01, 0xaf, 0x3c, 0xc3, 0xb6, 0x33, 0xe4, 0xc4, 0x4b, 0xf6, 0xf8,
	0x65, 0x50, 0x4c, 0x32, 0x0e, 0x30, 0xa3, 0x71, 0xa9, 0x08, 0xe1, 0xe9, 0xcc, 0x03, 0x39, 0x14,
	0xb5, 0xfc, 0x55, 0xf8, 0x5a, 0x06, 0x73, 0xcd, 0x0f, 0x86, 0x8f, 0xcf, 0x81, 0xc5, 0xae, 0x69,
	0xc6, 0x52, 0x84, 0x6f, 0x66, 0x93, 0x49, 0x22, 0x25, 0xf3, 0xd5, 0xfc, 0x60, 0xce, 0xd0, 0x2f,
	0x73, 0xc7, 0xdd, 0x27, 0x61, 0xff, 0xfc, 0x98, 0xbe, 0x21, 0xbe, 0x57, 0xf4, 0x23, 0x05, 0x7c,
	0xab, 0x54, 0xbf, 0x51, 0x52, 0x4d, 0xe4, 0x86, 0x4d, 0x64, 0x96, 0x7b, 0xfb, 0x4e, 0x28, 0xe8,
	0x97, 0xb6, 0xc9, 0x65, 0xa7, 0x34, 0xf4, 0x16, 0x31, 0xdb, 0xcd, 0xbe, 0xaa, 0x13, 0xa3, 0xa7,
	0xd6, 0xfb, 0x1b, 0x7d, 0x15, 0xf7, 0xcc, 0x96, 0xfa, 0x56, 0xaf, 0xdf, 0x6c, 0xac, 0xad, 0xd5,
	0x1b, 0x6d, 0x6c, 0xc8, 0x2e, 0x6a, 0x6f, 0xb4, 0x8c, 0x8d, 0x7a, 0xab, 0xa1, 0x36, 0x5a, 0xf5,
	0x96, 0x5a, 0xef, 0xb7, 0x7b, 0x2a, 0x6e, 0x35, 0x9b, 0xaa, 0xb1, 0xbe, 0xde, 0xea, 0xb7, 0xd6,
	0xdb, 0xcd, 0x06, 0x26, 0xd3, 0xfb, 0x2b, 0xbc, 0x3d, 0x26, 0x99, 0xf0, 0x32, 0x2c, 0x26, 0x8b,
	0x11, 0xbb, 0x46, 0x6e, 0xa0, 0x6b, 0x53, 0xcb, 0xd1, 0xc1, 0xa6, 0x29, 0xbb, 0xed, 0xe7, 0x73,
	0x60, 0xf9, 0x1e, 0x71, 0xe8, 0x0e, 0x89, 0x17, 0x66, 0x8a, 0x2b, 0x53, 0xe0, 0x19, 0xae, 0x9c,
	0x80, 0xe7, 0x0c, 0xfd, 0xfa, 0x7c, 0x94, 0x67, 0xea, 0x19, 0x91, 0x2a, 0x4f, 0xec, 0x9c, 0x78,
	0x03, 0x5d, 0x9f, 0x5e, 0x1e, 0xcf, 0x97, 0x4d, 0x56, 0xe8, 0xa7, 0xb9, 0xf0, 0xdc, 0xe0, 0x1f,
	0x52, 0x2f, 0x50, 0x31, 0xc7, 0xb9, 0x11, 0x81, 0xf3, 0x9d, 0x1b, 0x71, 0x3c, 0x67, 0xe8, 0x49,
	0xf2, 0xdc, 0x40, 0xa0, 0xe2, 0x2b, 0x10, 0x64, 0x97, 0x14, 0xfd, 0x60, 0xc2, 0x33, 0xe1, 0x1a,
	0xb8, 0x2a, 0xc9, 0x4f, 0x81, 0x8f, 0xc6, 0xa5, 0x97, 0xe1, 0xf2, 0x29, 0x68, 0x33, 0x71, 0xf5,
	0x56, 0xe1, 0xad, 0x84, 0x3c, 0xfe, 0x32, 0xda, 0x33, 0xcb, 0x04, 0xa2, 0xbd, 0xd7, 0xfd, 0xfc,
	0x1d, 0xcb, 0x16, 0x5b, 0xc3, 0x5e, 0xcd, 0xa0, 0x8e, 0x26, 0x59, 0x86, 0x0f, 0x58, 0x2d, 0xdf,
	0x9b, 0xba, 0x57, 0xf0, 0x5f, 0xb0, 0xeb, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x65, 0x80, 0x86,
	0x15, 0x5a, 0x0c, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TeamsClient is the client API for Teams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamsClient interface {
	//
	//Creates a local team
	//
	//Creates a local team that is used to group local users as members of IAM policies.
	//
	//Authorization Action:
	//```
	//iam:teams:create
	//```
	CreateTeam(ctx context.Context, in *request.CreateTeamReq, opts ...grpc.CallOption) (*response.CreateTeamResp, error)
	//
	//Lists all local teams
	//
	//Lists all local teams.
	//
	//Authorization Action:
	//```
	//iam:teams:list
	//```
	ListTeams(ctx context.Context, in *request.ListTeamsReq, opts ...grpc.CallOption) (*response.ListTeamsResp, error)
	//
	//Get a team
	//
	//Returns the details for a team.
	//
	//Authorization Action:
	//```
	//iam:teams:get
	//```
	GetTeam(ctx context.Context, in *request.GetTeamReq, opts ...grpc.CallOption) (*response.GetTeamResp, error)
	//
	//Updates a local team
	//
	//Updates a local team.
	//
	//Authorization Action:
	//```
	//iam:teams:update
	//```
	UpdateTeam(ctx context.Context, in *request.UpdateTeamReq, opts ...grpc.CallOption) (*response.UpdateTeamResp, error)
	//
	//Deletes a local team
	//
	//Deletes a local team and removes it from any policies.
	//
	//Authorization Action:
	//```
	//iam:teams:delete
	//```
	DeleteTeam(ctx context.Context, in *request.DeleteTeamReq, opts ...grpc.CallOption) (*response.DeleteTeamResp, error)
	//
	//Gets local team membership
	//
	//Lists all users of a local team. Users are listed by their membership_id.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:list
	//```
	GetTeamMembership(ctx context.Context, in *request.GetTeamMembershipReq, opts ...grpc.CallOption) (*response.GetTeamMembershipResp, error)
	//
	//Adds local team membership
	//
	//Adds a list of users to a local team. Users are added by their membership_id.
	//The request currently does not validate that membership_id maps to a real user.
	//
	//The membership_id for users can be found via GET /apis/apis/iam/v2/users/<user_id>.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:create
	//```
	AddTeamMembers(ctx context.Context, in *request.AddTeamMembersReq, opts ...grpc.CallOption) (*response.AddTeamMembersResp, error)
	//
	//Removes local team membership
	//
	//Removes a list of users from a local team. Users are removed by their membership_id.
	//The request currently does not validate that membership_id maps to a real user.
	//
	//The membership_id for users can be found via GET /apis/apis/iam/v2/users/<user_id>.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:delete
	//```
	RemoveTeamMembers(ctx context.Context, in *request.RemoveTeamMembersReq, opts ...grpc.CallOption) (*response.RemoveTeamMembersResp, error)
	//
	//Gets team membership for a user
	//
	//Lists all local teams for a specific user. You must use their membership_id in the request URL.
	//
	//Authorization Action:
	//```
	//iam:userTeams:get
	//```
	GetTeamsForMember(ctx context.Context, in *request.GetTeamsForMemberReq, opts ...grpc.CallOption) (*response.GetTeamsForMemberResp, error)
}

type teamsClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamsClient(cc grpc.ClientConnInterface) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) CreateTeam(ctx context.Context, in *request.CreateTeamReq, opts ...grpc.CallOption) (*response.CreateTeamResp, error) {
	out := new(response.CreateTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/CreateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) ListTeams(ctx context.Context, in *request.ListTeamsReq, opts ...grpc.CallOption) (*response.ListTeamsResp, error) {
	out := new(response.ListTeamsResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/ListTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeam(ctx context.Context, in *request.GetTeamReq, opts ...grpc.CallOption) (*response.GetTeamResp, error) {
	out := new(response.GetTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/GetTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) UpdateTeam(ctx context.Context, in *request.UpdateTeamReq, opts ...grpc.CallOption) (*response.UpdateTeamResp, error) {
	out := new(response.UpdateTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) DeleteTeam(ctx context.Context, in *request.DeleteTeamReq, opts ...grpc.CallOption) (*response.DeleteTeamResp, error) {
	out := new(response.DeleteTeamResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/DeleteTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeamMembership(ctx context.Context, in *request.GetTeamMembershipReq, opts ...grpc.CallOption) (*response.GetTeamMembershipResp, error) {
	out := new(response.GetTeamMembershipResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/GetTeamMembership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) AddTeamMembers(ctx context.Context, in *request.AddTeamMembersReq, opts ...grpc.CallOption) (*response.AddTeamMembersResp, error) {
	out := new(response.AddTeamMembersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/AddTeamMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) RemoveTeamMembers(ctx context.Context, in *request.RemoveTeamMembersReq, opts ...grpc.CallOption) (*response.RemoveTeamMembersResp, error) {
	out := new(response.RemoveTeamMembersResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/RemoveTeamMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeamsForMember(ctx context.Context, in *request.GetTeamsForMemberReq, opts ...grpc.CallOption) (*response.GetTeamsForMemberResp, error) {
	out := new(response.GetTeamsForMemberResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Teams/GetTeamsForMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamsServer is the server API for Teams service.
type TeamsServer interface {
	//
	//Creates a local team
	//
	//Creates a local team that is used to group local users as members of IAM policies.
	//
	//Authorization Action:
	//```
	//iam:teams:create
	//```
	CreateTeam(context.Context, *request.CreateTeamReq) (*response.CreateTeamResp, error)
	//
	//Lists all local teams
	//
	//Lists all local teams.
	//
	//Authorization Action:
	//```
	//iam:teams:list
	//```
	ListTeams(context.Context, *request.ListTeamsReq) (*response.ListTeamsResp, error)
	//
	//Get a team
	//
	//Returns the details for a team.
	//
	//Authorization Action:
	//```
	//iam:teams:get
	//```
	GetTeam(context.Context, *request.GetTeamReq) (*response.GetTeamResp, error)
	//
	//Updates a local team
	//
	//Updates a local team.
	//
	//Authorization Action:
	//```
	//iam:teams:update
	//```
	UpdateTeam(context.Context, *request.UpdateTeamReq) (*response.UpdateTeamResp, error)
	//
	//Deletes a local team
	//
	//Deletes a local team and removes it from any policies.
	//
	//Authorization Action:
	//```
	//iam:teams:delete
	//```
	DeleteTeam(context.Context, *request.DeleteTeamReq) (*response.DeleteTeamResp, error)
	//
	//Gets local team membership
	//
	//Lists all users of a local team. Users are listed by their membership_id.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:list
	//```
	GetTeamMembership(context.Context, *request.GetTeamMembershipReq) (*response.GetTeamMembershipResp, error)
	//
	//Adds local team membership
	//
	//Adds a list of users to a local team. Users are added by their membership_id.
	//The request currently does not validate that membership_id maps to a real user.
	//
	//The membership_id for users can be found via GET /apis/apis/iam/v2/users/<user_id>.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:create
	//```
	AddTeamMembers(context.Context, *request.AddTeamMembersReq) (*response.AddTeamMembersResp, error)
	//
	//Removes local team membership
	//
	//Removes a list of users from a local team. Users are removed by their membership_id.
	//The request currently does not validate that membership_id maps to a real user.
	//
	//The membership_id for users can be found via GET /apis/apis/iam/v2/users/<user_id>.
	//
	//Authorization Action:
	//```
	//iam:teamUsers:delete
	//```
	RemoveTeamMembers(context.Context, *request.RemoveTeamMembersReq) (*response.RemoveTeamMembersResp, error)
	//
	//Gets team membership for a user
	//
	//Lists all local teams for a specific user. You must use their membership_id in the request URL.
	//
	//Authorization Action:
	//```
	//iam:userTeams:get
	//```
	GetTeamsForMember(context.Context, *request.GetTeamsForMemberReq) (*response.GetTeamsForMemberResp, error)
}

// UnimplementedTeamsServer can be embedded to have forward compatible implementations.
type UnimplementedTeamsServer struct {
}

func (*UnimplementedTeamsServer) CreateTeam(ctx context.Context, req *request.CreateTeamReq) (*response.CreateTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTeam not implemented")
}
func (*UnimplementedTeamsServer) ListTeams(ctx context.Context, req *request.ListTeamsReq) (*response.ListTeamsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeams not implemented")
}
func (*UnimplementedTeamsServer) GetTeam(ctx context.Context, req *request.GetTeamReq) (*response.GetTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (*UnimplementedTeamsServer) UpdateTeam(ctx context.Context, req *request.UpdateTeamReq) (*response.UpdateTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (*UnimplementedTeamsServer) DeleteTeam(ctx context.Context, req *request.DeleteTeamReq) (*response.DeleteTeamResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTeam not implemented")
}
func (*UnimplementedTeamsServer) GetTeamMembership(ctx context.Context, req *request.GetTeamMembershipReq) (*response.GetTeamMembershipResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamMembership not implemented")
}
func (*UnimplementedTeamsServer) AddTeamMembers(ctx context.Context, req *request.AddTeamMembersReq) (*response.AddTeamMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeamMembers not implemented")
}
func (*UnimplementedTeamsServer) RemoveTeamMembers(ctx context.Context, req *request.RemoveTeamMembersReq) (*response.RemoveTeamMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTeamMembers not implemented")
}
func (*UnimplementedTeamsServer) GetTeamsForMember(ctx context.Context, req *request.GetTeamsForMemberReq) (*response.GetTeamsForMemberResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamsForMember not implemented")
}

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_CreateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).CreateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/CreateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).CreateTeam(ctx, req.(*request.CreateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_ListTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListTeamsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).ListTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/ListTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).ListTeams(ctx, req.(*request.ListTeamsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeam(ctx, req.(*request.GetTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).UpdateTeam(ctx, req.(*request.UpdateTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_DeleteTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteTeamReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).DeleteTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/DeleteTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).DeleteTeam(ctx, req.(*request.DeleteTeamReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeamMembership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamMembershipReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeamMembership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/GetTeamMembership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeamMembership(ctx, req.(*request.GetTeamMembershipReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_AddTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.AddTeamMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).AddTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/AddTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).AddTeamMembers(ctx, req.(*request.AddTeamMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_RemoveTeamMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.RemoveTeamMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).RemoveTeamMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/RemoveTeamMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).RemoveTeamMembers(ctx, req.(*request.RemoveTeamMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeamsForMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTeamsForMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeamsForMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Teams/GetTeamsForMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeamsForMember(ctx, req.(*request.GetTeamsForMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTeam",
			Handler:    _Teams_CreateTeam_Handler,
		},
		{
			MethodName: "ListTeams",
			Handler:    _Teams_ListTeams_Handler,
		},
		{
			MethodName: "GetTeam",
			Handler:    _Teams_GetTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Teams_UpdateTeam_Handler,
		},
		{
			MethodName: "DeleteTeam",
			Handler:    _Teams_DeleteTeam_Handler,
		},
		{
			MethodName: "GetTeamMembership",
			Handler:    _Teams_GetTeamMembership_Handler,
		},
		{
			MethodName: "AddTeamMembers",
			Handler:    _Teams_AddTeamMembers_Handler,
		},
		{
			MethodName: "RemoveTeamMembers",
			Handler:    _Teams_RemoveTeamMembers_Handler,
		},
		{
			MethodName: "GetTeamsForMember",
			Handler:    _Teams_GetTeamsForMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "components/automate-gateway/api/iam/v2/teams.proto",
}
