// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: ocr/ocr.proto

package ocr

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OCRService_OCR_FullMethodName = "/ocr.OCRService/OCR"
)

// OCRServiceClient is the client API for OCRService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OCRServiceClient interface {
	OCR(ctx context.Context, in *OCRRequest, opts ...grpc.CallOption) (*OCRResponse, error)
}

type oCRServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOCRServiceClient(cc grpc.ClientConnInterface) OCRServiceClient {
	return &oCRServiceClient{cc}
}

func (c *oCRServiceClient) OCR(ctx context.Context, in *OCRRequest, opts ...grpc.CallOption) (*OCRResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OCRResponse)
	err := c.cc.Invoke(ctx, OCRService_OCR_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OCRServiceServer is the server API for OCRService service.
// All implementations must embed UnimplementedOCRServiceServer
// for forward compatibility.
type OCRServiceServer interface {
	OCR(context.Context, *OCRRequest) (*OCRResponse, error)
	mustEmbedUnimplementedOCRServiceServer()
}

// UnimplementedOCRServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOCRServiceServer struct{}

func (UnimplementedOCRServiceServer) OCR(context.Context, *OCRRequest) (*OCRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OCR not implemented")
}
func (UnimplementedOCRServiceServer) mustEmbedUnimplementedOCRServiceServer() {}
func (UnimplementedOCRServiceServer) testEmbeddedByValue()                    {}

// UnsafeOCRServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OCRServiceServer will
// result in compilation errors.
type UnsafeOCRServiceServer interface {
	mustEmbedUnimplementedOCRServiceServer()
}

func RegisterOCRServiceServer(s grpc.ServiceRegistrar, srv OCRServiceServer) {
	// If the following call pancis, it indicates UnimplementedOCRServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OCRService_ServiceDesc, srv)
}

func _OCRService_OCR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OCRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCRServiceServer).OCR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OCRService_OCR_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCRServiceServer).OCR(ctx, req.(*OCRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OCRService_ServiceDesc is the grpc.ServiceDesc for OCRService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OCRService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocr.OCRService",
	HandlerType: (*OCRServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OCR",
			Handler:    _OCRService_OCR_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocr/ocr.proto",
}
