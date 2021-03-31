// Code generated by protoc-gen-yarpc. DO NOT EDIT.
// source: storage.proto

/*
Package storage is a generated protocol buffer package.

It is generated from these files:
	storage.proto
	predicate.proto

It has these top-level messages:
	ReadRequest
	Aggregate
	Tag
	ReadResponse
	CapabilitiesResponse
	HintsResponse
	TimestampRange
	Node
	Predicate
*/
package storage

import (
	context "context"

	yarpc "github.com/EnnioRC/yarpc"
)

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/EnnioRC/yarpc/yarpcproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ yarpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the yarpc package it is being compiled against.
const _ = yarpc.SupportPackageIsVersion1

// Client API for Storage service

type StorageClient interface {
	// Read performs a read operation using the given ReadRequest
	Read(ctx context.Context, in *ReadRequest) (Storage_ReadClient, error)
	// Capabilities returns a map of keys and values identifying the capabilities supported by the storage engine
	Capabilities(ctx context.Context, in *google_protobuf1.Empty) (*CapabilitiesResponse, error)
	Hints(ctx context.Context, in *google_protobuf1.Empty) (*HintsResponse, error)
}

type storageClient struct {
	cc *yarpc.ClientConn
}

func NewStorageClient(cc *yarpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) Read(ctx context.Context, in *ReadRequest) (Storage_ReadClient, error) {
	stream, err := yarpc.NewClientStream(ctx, &_Storage_serviceDesc.Streams[0], c.cc, 0x0000)
	if err != nil {
		return nil, err
	}
	x := &storageReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	return x, nil
}

type Storage_ReadClient interface {
	Recv() (*ReadResponse, error)
	yarpc.ClientStream
}

type storageReadClient struct {
	yarpc.ClientStream
}

func (x *storageReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storageClient) Capabilities(ctx context.Context, in *google_protobuf1.Empty) (*CapabilitiesResponse, error) {
	out := new(CapabilitiesResponse)
	err := yarpc.Invoke(ctx, 0x0001, in, out, c.cc)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) Hints(ctx context.Context, in *google_protobuf1.Empty) (*HintsResponse, error) {
	out := new(HintsResponse)
	err := yarpc.Invoke(ctx, 0x0002, in, out, c.cc)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Storage service

type StorageServer interface {
	// Read performs a read operation using the given ReadRequest
	Read(*ReadRequest, Storage_ReadServer) error
	// Capabilities returns a map of keys and values identifying the capabilities supported by the storage engine
	Capabilities(context.Context, *google_protobuf1.Empty) (*CapabilitiesResponse, error)
	Hints(context.Context, *google_protobuf1.Empty) (*HintsResponse, error)
}

func RegisterStorageServer(s *yarpc.Server, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_Read_Handler(srv interface{}, stream yarpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StorageServer).Read(m, &storageReadServer{stream})
}

type Storage_ReadServer interface {
	Send(*ReadResponse) error
	yarpc.ServerStream
}

type storageReadServer struct {
	yarpc.ServerStream
}

func (x *storageReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Storage_Capabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(StorageServer).Capabilities(ctx, in)
}

func _Storage_Hints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	return srv.(StorageServer).Hints(ctx, in)
}

var _Storage_serviceDesc = yarpc.ServiceDesc{
	ServiceName: "com.github.EnnioRC.influxdb.services.storage.Storage",
	Index:       0,
	HandlerType: (*StorageServer)(nil),
	Methods: []yarpc.MethodDesc{
		{
			MethodName: "Capabilities",
			Index:      1,
			Handler:    _Storage_Capabilities_Handler,
		},
		{
			MethodName: "Hints",
			Index:      2,
			Handler:    _Storage_Hints_Handler,
		},
	},
	Streams: []yarpc.StreamDesc{
		{
			StreamName:    "Read",
			Index:         0,
			Handler:       _Storage_Read_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "storage.proto",
}
