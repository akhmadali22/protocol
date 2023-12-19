// Code generated by protoc-gen-psrpc v0.5.1, DO NOT EDIT.
// source: rpc/room.proto

package rpc

import (
	"context"

	"github.com/livekit/psrpc"
	"github.com/livekit/psrpc/pkg/client"
	"github.com/livekit/psrpc/pkg/info"
	"github.com/livekit/psrpc/pkg/rand"
	"github.com/livekit/psrpc/pkg/server"
	"github.com/livekit/psrpc/version"
)
import livekit "github.com/akhmadali22/protocol/livekit"
import livekit4 "github.com/akhmadali22/protocol/livekit"

var _ = version.PsrpcVersion_0_5

// =====================
// Room Client Interface
// =====================

type RoomClient[RoomTopicType ~string] interface {
	DeleteRoom(ctx context.Context, room RoomTopicType, req *livekit4.DeleteRoomRequest, opts ...psrpc.RequestOption) (*livekit4.DeleteRoomResponse, error)

	SendData(ctx context.Context, room RoomTopicType, req *livekit4.SendDataRequest, opts ...psrpc.RequestOption) (*livekit4.SendDataResponse, error)

	UpdateRoomMetadata(ctx context.Context, room RoomTopicType, req *livekit4.UpdateRoomMetadataRequest, opts ...psrpc.RequestOption) (*livekit.Room, error)
}

// =========================
// Room ServerImpl Interface
// =========================

type RoomServerImpl interface {
	DeleteRoom(context.Context, *livekit4.DeleteRoomRequest) (*livekit4.DeleteRoomResponse, error)

	SendData(context.Context, *livekit4.SendDataRequest) (*livekit4.SendDataResponse, error)

	UpdateRoomMetadata(context.Context, *livekit4.UpdateRoomMetadataRequest) (*livekit.Room, error)
}

// =====================
// Room Server Interface
// =====================

type RoomServer[RoomTopicType ~string] interface {
	RegisterDeleteRoomTopic(room RoomTopicType) error
	DeregisterDeleteRoomTopic(room RoomTopicType)
	RegisterSendDataTopic(room RoomTopicType) error
	DeregisterSendDataTopic(room RoomTopicType)
	RegisterUpdateRoomMetadataTopic(room RoomTopicType) error
	DeregisterUpdateRoomMetadataTopic(room RoomTopicType)
	RegisterAllRoomTopics(room RoomTopicType) error
	DeregisterAllRoomTopics(room RoomTopicType)

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// ===========
// Room Client
// ===========

type roomClient[RoomTopicType ~string] struct {
	client *client.RPCClient
}

// NewRoomClient creates a psrpc client that implements the RoomClient interface.
func NewRoomClient[RoomTopicType ~string](bus psrpc.MessageBus, opts ...psrpc.ClientOption) (RoomClient[RoomTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Room",
		ID:   rand.NewClientID(),
	}

	sd.RegisterMethod("DeleteRoom", false, false, true, true)
	sd.RegisterMethod("SendData", false, false, true, true)
	sd.RegisterMethod("UpdateRoomMetadata", false, false, true, true)

	rpcClient, err := client.NewRPCClient(sd, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &roomClient[RoomTopicType]{
		client: rpcClient,
	}, nil
}

func (c *roomClient[RoomTopicType]) DeleteRoom(ctx context.Context, room RoomTopicType, req *livekit4.DeleteRoomRequest, opts ...psrpc.RequestOption) (*livekit4.DeleteRoomResponse, error) {
	return client.RequestSingle[*livekit4.DeleteRoomResponse](ctx, c.client, "DeleteRoom", []string{string(room)}, req, opts...)
}

func (c *roomClient[RoomTopicType]) SendData(ctx context.Context, room RoomTopicType, req *livekit4.SendDataRequest, opts ...psrpc.RequestOption) (*livekit4.SendDataResponse, error) {
	return client.RequestSingle[*livekit4.SendDataResponse](ctx, c.client, "SendData", []string{string(room)}, req, opts...)
}

func (c *roomClient[RoomTopicType]) UpdateRoomMetadata(ctx context.Context, room RoomTopicType, req *livekit4.UpdateRoomMetadataRequest, opts ...psrpc.RequestOption) (*livekit.Room, error) {
	return client.RequestSingle[*livekit.Room](ctx, c.client, "UpdateRoomMetadata", []string{string(room)}, req, opts...)
}

// ===========
// Room Server
// ===========

type roomServer[RoomTopicType ~string] struct {
	svc RoomServerImpl
	rpc *server.RPCServer
}

// NewRoomServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewRoomServer[RoomTopicType ~string](svc RoomServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (RoomServer[RoomTopicType], error) {
	sd := &info.ServiceDefinition{
		Name: "Room",
		ID:   rand.NewServerID(),
	}

	s := server.NewRPCServer(sd, bus, opts...)

	sd.RegisterMethod("DeleteRoom", false, false, true, true)
	sd.RegisterMethod("SendData", false, false, true, true)
	sd.RegisterMethod("UpdateRoomMetadata", false, false, true, true)
	return &roomServer[RoomTopicType]{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *roomServer[RoomTopicType]) RegisterDeleteRoomTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "DeleteRoom", []string{string(room)}, s.svc.DeleteRoom, nil)
}

func (s *roomServer[RoomTopicType]) DeregisterDeleteRoomTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("DeleteRoom", []string{string(room)})
}

func (s *roomServer[RoomTopicType]) RegisterSendDataTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "SendData", []string{string(room)}, s.svc.SendData, nil)
}

func (s *roomServer[RoomTopicType]) DeregisterSendDataTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("SendData", []string{string(room)})
}

func (s *roomServer[RoomTopicType]) RegisterUpdateRoomMetadataTopic(room RoomTopicType) error {
	return server.RegisterHandler(s.rpc, "UpdateRoomMetadata", []string{string(room)}, s.svc.UpdateRoomMetadata, nil)
}

func (s *roomServer[RoomTopicType]) DeregisterUpdateRoomMetadataTopic(room RoomTopicType) {
	s.rpc.DeregisterHandler("UpdateRoomMetadata", []string{string(room)})
}

func (s *roomServer[RoomTopicType]) allRoomTopicRegisterers() server.RegistererSlice {
	return server.RegistererSlice{
		server.NewRegisterer(s.RegisterDeleteRoomTopic, s.DeregisterDeleteRoomTopic),
		server.NewRegisterer(s.RegisterSendDataTopic, s.DeregisterSendDataTopic),
		server.NewRegisterer(s.RegisterUpdateRoomMetadataTopic, s.DeregisterUpdateRoomMetadataTopic),
	}
}

func (s *roomServer[RoomTopicType]) RegisterAllRoomTopics(room RoomTopicType) error {
	return s.allRoomTopicRegisterers().Register(room)
}

func (s *roomServer[RoomTopicType]) DeregisterAllRoomTopics(room RoomTopicType) {
	s.allRoomTopicRegisterers().Deregister(room)
}

func (s *roomServer[RoomTopicType]) Shutdown() {
	s.rpc.Close(false)
}

func (s *roomServer[RoomTopicType]) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor5 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2a, 0x48, 0xd6,
	0x2f, 0xca, 0xcf, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96,
	0xe2, 0xcd, 0x2f, 0x28, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x88, 0x49, 0x89, 0xe4, 0x64, 0x96, 0xa5,
	0x66, 0x67, 0x96, 0xc4, 0xe7, 0xe6, 0xa7, 0xa4, 0xe6, 0xc0, 0x44, 0x85, 0x60, 0xa2, 0x08, 0xdd,
	0x46, 0xf3, 0x99, 0xb8, 0x58, 0x82, 0xf2, 0xf3, 0x73, 0x85, 0x62, 0xb9, 0xb8, 0x5c, 0x52, 0x73,
	0x52, 0x4b, 0x52, 0xc1, 0x3c, 0x29, 0x3d, 0xa8, 0x5a, 0x3d, 0x84, 0x60, 0x50, 0x6a, 0x61, 0x69,
	0x6a, 0x71, 0x89, 0x94, 0x34, 0x56, 0xb9, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x25, 0xb1, 0x4d,
	0x9d, 0x8c, 0x42, 0x02, 0x8c, 0x52, 0x7c, 0x5c, 0x2c, 0x20, 0x5b, 0x84, 0xc0, 0xa4, 0x04, 0xa3,
	0x50, 0x38, 0x17, 0x47, 0x70, 0x6a, 0x5e, 0x8a, 0x4b, 0x62, 0x49, 0xa2, 0x90, 0x04, 0xdc, 0x00,
	0x98, 0x10, 0xcc, 0x68, 0x49, 0x2c, 0x32, 0x04, 0x0c, 0x8e, 0xe7, 0x12, 0x0a, 0x2d, 0x48, 0x49,
	0x84, 0x38, 0xc3, 0x37, 0xb5, 0x24, 0x31, 0x05, 0x64, 0x85, 0x12, 0xdc, 0x20, 0x4c, 0x49, 0x98,
	0x65, 0xbc, 0x70, 0x35, 0x20, 0x59, 0x5c, 0x16, 0x38, 0xe9, 0x44, 0x69, 0xa5, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x43, 0xb5, 0xc0, 0xe9, 0x82, 0xec, 0x74, 0xfd, 0xe2,
	0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xfd, 0xa2, 0x82, 0xe4, 0x24, 0x36, 0x70, 0xb0, 0x1a, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x53, 0xa6, 0x88, 0x04, 0xa6, 0x01, 0x00, 0x00,
}
