// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pdfcomposegrpc.proto

package pdfcomposegrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PDFComposer_ConverToPdf_FullMethodName = "/pdfcomposegrpc.PDFComposer/ConverToPdf"
)

// PDFComposerClient is the client API for PDFComposer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PDFComposerClient interface {
	ConverToPdf(ctx context.Context, opts ...grpc.CallOption) (PDFComposer_ConverToPdfClient, error)
}

type pDFComposerClient struct {
	cc grpc.ClientConnInterface
}

func NewPDFComposerClient(cc grpc.ClientConnInterface) PDFComposerClient {
	return &pDFComposerClient{cc}
}

func (c *pDFComposerClient) ConverToPdf(ctx context.Context, opts ...grpc.CallOption) (PDFComposer_ConverToPdfClient, error) {
	stream, err := c.cc.NewStream(ctx, &PDFComposer_ServiceDesc.Streams[0], PDFComposer_ConverToPdf_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pDFComposerConverToPdfClient{stream}
	return x, nil
}

type PDFComposer_ConverToPdfClient interface {
	Send(*ImageList) error
	Recv() (*PDFResponse, error)
	grpc.ClientStream
}

type pDFComposerConverToPdfClient struct {
	grpc.ClientStream
}

func (x *pDFComposerConverToPdfClient) Send(m *ImageList) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pDFComposerConverToPdfClient) Recv() (*PDFResponse, error) {
	m := new(PDFResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PDFComposerServer is the server API for PDFComposer service.
// All implementations must embed UnimplementedPDFComposerServer
// for forward compatibility
type PDFComposerServer interface {
	ConverToPdf(PDFComposer_ConverToPdfServer) error
	mustEmbedUnimplementedPDFComposerServer()
}

// UnimplementedPDFComposerServer must be embedded to have forward compatible implementations.
type UnimplementedPDFComposerServer struct {
}

func (UnimplementedPDFComposerServer) ConverToPdf(PDFComposer_ConverToPdfServer) error {
	return status.Errorf(codes.Unimplemented, "method ConverToPdf not implemented")
}
func (UnimplementedPDFComposerServer) mustEmbedUnimplementedPDFComposerServer() {}

// UnsafePDFComposerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PDFComposerServer will
// result in compilation errors.
type UnsafePDFComposerServer interface {
	mustEmbedUnimplementedPDFComposerServer()
}

func RegisterPDFComposerServer(s grpc.ServiceRegistrar, srv PDFComposerServer) {
	s.RegisterService(&PDFComposer_ServiceDesc, srv)
}

func _PDFComposer_ConverToPdf_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PDFComposerServer).ConverToPdf(&pDFComposerConverToPdfServer{stream})
}

type PDFComposer_ConverToPdfServer interface {
	Send(*PDFResponse) error
	Recv() (*ImageList, error)
	grpc.ServerStream
}

type pDFComposerConverToPdfServer struct {
	grpc.ServerStream
}

func (x *pDFComposerConverToPdfServer) Send(m *PDFResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pDFComposerConverToPdfServer) Recv() (*ImageList, error) {
	m := new(ImageList)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PDFComposer_ServiceDesc is the grpc.ServiceDesc for PDFComposer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PDFComposer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdfcomposegrpc.PDFComposer",
	HandlerType: (*PDFComposerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConverToPdf",
			Handler:       _PDFComposer_ConverToPdf_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pdfcomposegrpc.proto",
}
