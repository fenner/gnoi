// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package os

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

// OSClient is the client API for OS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OSClient interface {
	// Install transfers an OS package into the Target. No concurrent Install RPCs
	// MUST be allowed to the same Target.
	//
	// The OS package file format is platform dependent. The platform MUST
	// validate that the OS package that is supplied is valid and bootable. This
	// SHOULD include a hash check against a known good hash. It is recommended
	// that the hash is embedded in the OS package.
	//
	// The Target manages its own persistent storage, and OS installation process.
	// It stores a set of distinct OS packages, and always proactively frees up
	// space for incoming new OS packages. It is guaranteed that the Target always
	// has enough space for a valid incoming OS package. The currently running OS
	// packages MUST never be removed. The Client MUST expect that the last
	// successfully installed package is available.
	//
	// The Install RPC allows the Client to specify the OS package version. If
	// the Target already has an OS package with the same version then there is no
	// need to transfer the OS package to the Target. If the Target does not have
	// an OS package with the same version, then the OS package is copied.
	//
	// Scenario 1 - When the Target already has the OS package:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [Validated|InstallError]
	//
	//
	// Scenario 2 - When the Target does not have the OS package:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [TransferReady|InstallError]
	//            transfer_content  -->
	//                              ...
	//                              <-- [TransferProgress|InstallError]
	//                              ...
	//                  TransferEnd -->
	//                              <-- [Validated|InstallError]
	//
	// On a dual Supervisor Target, only the Active Supervisor runs this gNOI
	// Service. The Install RPC applies to the Active Supervisor unless
	// InstallRequest->TransferRequest->standby_supervisor is set, in which case
	// it applies to the Standby Supervisor. One Install RPC is required for each
	// Supervisor. The Supervisor order of package installation MUST not be fixed.
	//
	// The Target MUST always attempt to copy the OS package between Supervisors
	// first before accepting the transfer from the Client. The syncing progress
	// is reported to the client with InstallResponse->SyncProgress messages.
	//
	// If a switchover is triggered during the Install RPC, the RPC MUST
	// immediately abort with Error->type->UNEXPECTED_SWITCHOVER.
	//
	// Scenario 3 - When both Supervisors already have the OS package, regardless
	//              of the value in Start.standby_supervisor:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [Validated|InstallError]
	//
	//
	// Scenario 4 - When one of the Supervisors already has the OS package but the
	//              other Supervisor is the target of the Install:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [SyncProgress|InstallError]
	//                              ...
	//                              <-- [Validated|InstallError]
	//
	//
	// Scenario 5 - When neither of the two Supervisors has the OS package:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [TransferReady|InstallError]
	//            transfer_content  -->
	//                              ...
	//                              <-- [TransferProgress|InstallError]
	//                              ...
	//                  TransferEnd -->
	//                              <-- [Validated|InstallError]
	//
	Install(ctx context.Context, opts ...grpc.CallOption) (OS_InstallClient, error)
	// Activate sets the requested OS version as the version which is used at the
	// next reboot, and reboots the Target if the 'activate_and_reboot' flag is
	// set. When booting the requested OS version fails, the Target recovers by
	// booting the previously running OS package.
	Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error)
	// Verify checks the running OS version. This RPC may be called multiple times
	// while the Target boots, until successful.
	Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error)
}

type oSClient struct {
	cc grpc.ClientConnInterface
}

func NewOSClient(cc grpc.ClientConnInterface) OSClient {
	return &oSClient{cc}
}

func (c *oSClient) Install(ctx context.Context, opts ...grpc.CallOption) (OS_InstallClient, error) {
	stream, err := c.cc.NewStream(ctx, &OS_ServiceDesc.Streams[0], "/gnoi.os.OS/Install", opts...)
	if err != nil {
		return nil, err
	}
	x := &oSInstallClient{stream}
	return x, nil
}

type OS_InstallClient interface {
	Send(*InstallRequest) error
	Recv() (*InstallResponse, error)
	grpc.ClientStream
}

type oSInstallClient struct {
	grpc.ClientStream
}

func (x *oSInstallClient) Send(m *InstallRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *oSInstallClient) Recv() (*InstallResponse, error) {
	m := new(InstallResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *oSClient) Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := c.cc.Invoke(ctx, "/gnoi.os.OS/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oSClient) Verify(ctx context.Context, in *VerifyRequest, opts ...grpc.CallOption) (*VerifyResponse, error) {
	out := new(VerifyResponse)
	err := c.cc.Invoke(ctx, "/gnoi.os.OS/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OSServer is the server API for OS service.
// All implementations must embed UnimplementedOSServer
// for forward compatibility
type OSServer interface {
	// Install transfers an OS package into the Target. No concurrent Install RPCs
	// MUST be allowed to the same Target.
	//
	// The OS package file format is platform dependent. The platform MUST
	// validate that the OS package that is supplied is valid and bootable. This
	// SHOULD include a hash check against a known good hash. It is recommended
	// that the hash is embedded in the OS package.
	//
	// The Target manages its own persistent storage, and OS installation process.
	// It stores a set of distinct OS packages, and always proactively frees up
	// space for incoming new OS packages. It is guaranteed that the Target always
	// has enough space for a valid incoming OS package. The currently running OS
	// packages MUST never be removed. The Client MUST expect that the last
	// successfully installed package is available.
	//
	// The Install RPC allows the Client to specify the OS package version. If
	// the Target already has an OS package with the same version then there is no
	// need to transfer the OS package to the Target. If the Target does not have
	// an OS package with the same version, then the OS package is copied.
	//
	// Scenario 1 - When the Target already has the OS package:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [Validated|InstallError]
	//
	//
	// Scenario 2 - When the Target does not have the OS package:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [TransferReady|InstallError]
	//            transfer_content  -->
	//                              ...
	//                              <-- [TransferProgress|InstallError]
	//                              ...
	//                  TransferEnd -->
	//                              <-- [Validated|InstallError]
	//
	// On a dual Supervisor Target, only the Active Supervisor runs this gNOI
	// Service. The Install RPC applies to the Active Supervisor unless
	// InstallRequest->TransferRequest->standby_supervisor is set, in which case
	// it applies to the Standby Supervisor. One Install RPC is required for each
	// Supervisor. The Supervisor order of package installation MUST not be fixed.
	//
	// The Target MUST always attempt to copy the OS package between Supervisors
	// first before accepting the transfer from the Client. The syncing progress
	// is reported to the client with InstallResponse->SyncProgress messages.
	//
	// If a switchover is triggered during the Install RPC, the RPC MUST
	// immediately abort with Error->type->UNEXPECTED_SWITCHOVER.
	//
	// Scenario 3 - When both Supervisors already have the OS package, regardless
	//              of the value in Start.standby_supervisor:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [Validated|InstallError]
	//
	//
	// Scenario 4 - When one of the Supervisors already has the OS package but the
	//              other Supervisor is the target of the Install:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [SyncProgress|InstallError]
	//                              ...
	//                              <-- [Validated|InstallError]
	//
	//
	// Scenario 5 - When neither of the two Supervisors has the OS package:
	//
	//         Client :--------------|--------------> Target
	//              TransferRequest -->
	//                              <-- [TransferReady|InstallError]
	//            transfer_content  -->
	//                              ...
	//                              <-- [TransferProgress|InstallError]
	//                              ...
	//                  TransferEnd -->
	//                              <-- [Validated|InstallError]
	//
	Install(OS_InstallServer) error
	// Activate sets the requested OS version as the version which is used at the
	// next reboot, and reboots the Target if the 'activate_and_reboot' flag is
	// set. When booting the requested OS version fails, the Target recovers by
	// booting the previously running OS package.
	Activate(context.Context, *ActivateRequest) (*ActivateResponse, error)
	// Verify checks the running OS version. This RPC may be called multiple times
	// while the Target boots, until successful.
	Verify(context.Context, *VerifyRequest) (*VerifyResponse, error)
	mustEmbedUnimplementedOSServer()
}

// UnimplementedOSServer must be embedded to have forward compatible implementations.
type UnimplementedOSServer struct {
}

func (UnimplementedOSServer) Install(OS_InstallServer) error {
	return status.Errorf(codes.Unimplemented, "method Install not implemented")
}
func (UnimplementedOSServer) Activate(context.Context, *ActivateRequest) (*ActivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedOSServer) Verify(context.Context, *VerifyRequest) (*VerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedOSServer) mustEmbedUnimplementedOSServer() {}

// UnsafeOSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OSServer will
// result in compilation errors.
type UnsafeOSServer interface {
	mustEmbedUnimplementedOSServer()
}

func RegisterOSServer(s grpc.ServiceRegistrar, srv OSServer) {
	s.RegisterService(&OS_ServiceDesc, srv)
}

func _OS_Install_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OSServer).Install(&oSInstallServer{stream})
}

type OS_InstallServer interface {
	Send(*InstallResponse) error
	Recv() (*InstallRequest, error)
	grpc.ServerStream
}

type oSInstallServer struct {
	grpc.ServerStream
}

func (x *oSInstallServer) Send(m *InstallResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *oSInstallServer) Recv() (*InstallRequest, error) {
	m := new(InstallRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _OS_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.os.OS/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServer).Activate(ctx, req.(*ActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OS_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OSServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gnoi.os.OS/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OSServer).Verify(ctx, req.(*VerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OS_ServiceDesc is the grpc.ServiceDesc for OS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.os.OS",
	HandlerType: (*OSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _OS_Activate_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _OS_Verify_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Install",
			Handler:       _OS_Install_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "os/os.proto",
}
