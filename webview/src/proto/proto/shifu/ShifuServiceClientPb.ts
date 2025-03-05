/**
 * @fileoverview gRPC-Web generated client stub for shifu
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.29.3
// source: proto/shifu/shifu.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_shifu_shifu_pb from '../../proto/shifu/shifu_pb'; // proto import: "proto/shifu/shifu.proto"


export class ShifuServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorGetAllAvailableVersions = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/GetAllAvailableVersions',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.GetAllAvailableVersionsRequest,
    proto_shifu_shifu_pb.GetAllAvailableVersionsResponse,
    (request: proto_shifu_shifu_pb.GetAllAvailableVersionsRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.GetAllAvailableVersionsResponse.deserializeBinary
  );

  getAllAvailableVersions(
    request: proto_shifu_shifu_pb.GetAllAvailableVersionsRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.GetAllAvailableVersionsResponse>;

  getAllAvailableVersions(
    request: proto_shifu_shifu_pb.GetAllAvailableVersionsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.GetAllAvailableVersionsResponse) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.GetAllAvailableVersionsResponse>;

  getAllAvailableVersions(
    request: proto_shifu_shifu_pb.GetAllAvailableVersionsRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.GetAllAvailableVersionsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/GetAllAvailableVersions',
        request,
        metadata || {},
        this.methodDescriptorGetAllAvailableVersions,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/GetAllAvailableVersions',
    request,
    metadata || {},
    this.methodDescriptorGetAllAvailableVersions);
  }

  methodDescriptorCheckInstallation = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/CheckInstallation',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.CheckInstallationRequest,
    proto_shifu_shifu_pb.CheckInstallationResponse,
    (request: proto_shifu_shifu_pb.CheckInstallationRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.CheckInstallationResponse.deserializeBinary
  );

  checkInstallation(
    request: proto_shifu_shifu_pb.CheckInstallationRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.CheckInstallationResponse>;

  checkInstallation(
    request: proto_shifu_shifu_pb.CheckInstallationRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.CheckInstallationResponse) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.CheckInstallationResponse>;

  checkInstallation(
    request: proto_shifu_shifu_pb.CheckInstallationRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.CheckInstallationResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/CheckInstallation',
        request,
        metadata || {},
        this.methodDescriptorCheckInstallation,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/CheckInstallation',
    request,
    metadata || {},
    this.methodDescriptorCheckInstallation);
  }

  methodDescriptorInstallShifu = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/InstallShifu',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.InstallShifuRequest,
    proto_shifu_shifu_pb.InstallShifuResponse,
    (request: proto_shifu_shifu_pb.InstallShifuRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.InstallShifuResponse.deserializeBinary
  );

  installShifu(
    request: proto_shifu_shifu_pb.InstallShifuRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.InstallShifuResponse>;

  installShifu(
    request: proto_shifu_shifu_pb.InstallShifuRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.InstallShifuResponse) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.InstallShifuResponse>;

  installShifu(
    request: proto_shifu_shifu_pb.InstallShifuRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.InstallShifuResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/InstallShifu',
        request,
        metadata || {},
        this.methodDescriptorInstallShifu,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/InstallShifu',
    request,
    metadata || {},
    this.methodDescriptorInstallShifu);
  }

  methodDescriptorUninstallShifu = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/UninstallShifu',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.UninstallShifuRequest,
    proto_shifu_shifu_pb.UninstallShifuResponse,
    (request: proto_shifu_shifu_pb.UninstallShifuRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.UninstallShifuResponse.deserializeBinary
  );

  uninstallShifu(
    request: proto_shifu_shifu_pb.UninstallShifuRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.UninstallShifuResponse>;

  uninstallShifu(
    request: proto_shifu_shifu_pb.UninstallShifuRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.UninstallShifuResponse) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.UninstallShifuResponse>;

  uninstallShifu(
    request: proto_shifu_shifu_pb.UninstallShifuRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.UninstallShifuResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/UninstallShifu',
        request,
        metadata || {},
        this.methodDescriptorUninstallShifu,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/UninstallShifu',
    request,
    metadata || {},
    this.methodDescriptorUninstallShifu);
  }

  methodDescriptorListDevices = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/ListDevices',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.ListDevicesRequest,
    proto_shifu_shifu_pb.ListDevicesResponse,
    (request: proto_shifu_shifu_pb.ListDevicesRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.ListDevicesResponse.deserializeBinary
  );

  listDevices(
    request: proto_shifu_shifu_pb.ListDevicesRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.ListDevicesResponse>;

  listDevices(
    request: proto_shifu_shifu_pb.ListDevicesRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.ListDevicesResponse) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.ListDevicesResponse>;

  listDevices(
    request: proto_shifu_shifu_pb.ListDevicesRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.ListDevicesResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/ListDevices',
        request,
        metadata || {},
        this.methodDescriptorListDevices,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/ListDevices',
    request,
    metadata || {},
    this.methodDescriptorListDevices);
  }

  methodDescriptorGetDeviceDetails = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/GetDeviceDetails',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.GetDeviceDetailsRequest,
    proto_shifu_shifu_pb.GetDeviceDetailsResponse,
    (request: proto_shifu_shifu_pb.GetDeviceDetailsRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.GetDeviceDetailsResponse.deserializeBinary
  );

  getDeviceDetails(
    request: proto_shifu_shifu_pb.GetDeviceDetailsRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.GetDeviceDetailsResponse>;

  getDeviceDetails(
    request: proto_shifu_shifu_pb.GetDeviceDetailsRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.GetDeviceDetailsResponse) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.GetDeviceDetailsResponse>;

  getDeviceDetails(
    request: proto_shifu_shifu_pb.GetDeviceDetailsRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.GetDeviceDetailsResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/GetDeviceDetails',
        request,
        metadata || {},
        this.methodDescriptorGetDeviceDetails,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/GetDeviceDetails',
    request,
    metadata || {},
    this.methodDescriptorGetDeviceDetails);
  }

  methodDescriptorForwardPort = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/ForwardPort',
    grpcWeb.MethodType.SERVER_STREAMING,
    proto_shifu_shifu_pb.ForwardPortRequest,
    proto_shifu_shifu_pb.ForwardPortResponse,
    (request: proto_shifu_shifu_pb.ForwardPortRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.ForwardPortResponse.deserializeBinary
  );

  forwardPort(
    request: proto_shifu_shifu_pb.ForwardPortRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.ForwardPortResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/shifu.ShifuService/ForwardPort',
      request,
      metadata || {},
      this.methodDescriptorForwardPort);
  }

  methodDescriptorRestartDeviceShifu = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/RestartDeviceShifu',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.RestartDeviceShifuRequest,
    proto_shifu_shifu_pb.Empty,
    (request: proto_shifu_shifu_pb.RestartDeviceShifuRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.Empty.deserializeBinary
  );

  restartDeviceShifu(
    request: proto_shifu_shifu_pb.RestartDeviceShifuRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.Empty>;

  restartDeviceShifu(
    request: proto_shifu_shifu_pb.RestartDeviceShifuRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.Empty) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.Empty>;

  restartDeviceShifu(
    request: proto_shifu_shifu_pb.RestartDeviceShifuRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/RestartDeviceShifu',
        request,
        metadata || {},
        this.methodDescriptorRestartDeviceShifu,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/RestartDeviceShifu',
    request,
    metadata || {},
    this.methodDescriptorRestartDeviceShifu);
  }

  methodDescriptorDeleteDeviceShifu = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/DeleteDeviceShifu',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.DeleteDeviceShifuRequest,
    proto_shifu_shifu_pb.Empty,
    (request: proto_shifu_shifu_pb.DeleteDeviceShifuRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.Empty.deserializeBinary
  );

  deleteDeviceShifu(
    request: proto_shifu_shifu_pb.DeleteDeviceShifuRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.Empty>;

  deleteDeviceShifu(
    request: proto_shifu_shifu_pb.DeleteDeviceShifuRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.Empty) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.Empty>;

  deleteDeviceShifu(
    request: proto_shifu_shifu_pb.DeleteDeviceShifuRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/DeleteDeviceShifu',
        request,
        metadata || {},
        this.methodDescriptorDeleteDeviceShifu,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/DeleteDeviceShifu',
    request,
    metadata || {},
    this.methodDescriptorDeleteDeviceShifu);
  }

  methodDescriptorGetAllContainerName = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/GetAllContainerName',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.GetAllContainerNameRequest,
    proto_shifu_shifu_pb.GetAllContainerNameResponse,
    (request: proto_shifu_shifu_pb.GetAllContainerNameRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.GetAllContainerNameResponse.deserializeBinary
  );

  getAllContainerName(
    request: proto_shifu_shifu_pb.GetAllContainerNameRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.GetAllContainerNameResponse>;

  getAllContainerName(
    request: proto_shifu_shifu_pb.GetAllContainerNameRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.GetAllContainerNameResponse) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.GetAllContainerNameResponse>;

  getAllContainerName(
    request: proto_shifu_shifu_pb.GetAllContainerNameRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.GetAllContainerNameResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/GetAllContainerName',
        request,
        metadata || {},
        this.methodDescriptorGetAllContainerName,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/GetAllContainerName',
    request,
    metadata || {},
    this.methodDescriptorGetAllContainerName);
  }

  methodDescriptorGetDeviceShifuLogs = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/GetDeviceShifuLogs',
    grpcWeb.MethodType.SERVER_STREAMING,
    proto_shifu_shifu_pb.GetDeviceShifuLogsRequest,
    proto_shifu_shifu_pb.GetDeviceShifuLogsResponse,
    (request: proto_shifu_shifu_pb.GetDeviceShifuLogsRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.GetDeviceShifuLogsResponse.deserializeBinary
  );

  getDeviceShifuLogs(
    request: proto_shifu_shifu_pb.GetDeviceShifuLogsRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.GetDeviceShifuLogsResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/shifu.ShifuService/GetDeviceShifuLogs',
      request,
      metadata || {},
      this.methodDescriptorGetDeviceShifuLogs);
  }

  methodDescriptorExecuteCommand = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/ExecuteCommand',
    grpcWeb.MethodType.SERVER_STREAMING,
    proto_shifu_shifu_pb.CommandRequest,
    proto_shifu_shifu_pb.CommandResponse,
    (request: proto_shifu_shifu_pb.CommandRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.CommandResponse.deserializeBinary
  );

  executeCommand(
    request: proto_shifu_shifu_pb.CommandRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.CommandResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/shifu.ShifuService/ExecuteCommand',
      request,
      metadata || {},
      this.methodDescriptorExecuteCommand);
  }

  methodDescriptorGetCompletions = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/GetCompletions',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.CompletionRequest,
    proto_shifu_shifu_pb.CompletionResponse,
    (request: proto_shifu_shifu_pb.CompletionRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.CompletionResponse.deserializeBinary
  );

  getCompletions(
    request: proto_shifu_shifu_pb.CompletionRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.CompletionResponse>;

  getCompletions(
    request: proto_shifu_shifu_pb.CompletionRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.CompletionResponse) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.CompletionResponse>;

  getCompletions(
    request: proto_shifu_shifu_pb.CompletionRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.CompletionResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/GetCompletions',
        request,
        metadata || {},
        this.methodDescriptorGetCompletions,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/GetCompletions',
    request,
    metadata || {},
    this.methodDescriptorGetCompletions);
  }

  methodDescriptorInstallViaURL = new grpcWeb.MethodDescriptor(
    '/shifu.ShifuService/InstallViaURL',
    grpcWeb.MethodType.UNARY,
    proto_shifu_shifu_pb.InstallViaURLRequest,
    proto_shifu_shifu_pb.Empty,
    (request: proto_shifu_shifu_pb.InstallViaURLRequest) => {
      return request.serializeBinary();
    },
    proto_shifu_shifu_pb.Empty.deserializeBinary
  );

  installViaURL(
    request: proto_shifu_shifu_pb.InstallViaURLRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_shifu_shifu_pb.Empty>;

  installViaURL(
    request: proto_shifu_shifu_pb.InstallViaURLRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.Empty) => void): grpcWeb.ClientReadableStream<proto_shifu_shifu_pb.Empty>;

  installViaURL(
    request: proto_shifu_shifu_pb.InstallViaURLRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_shifu_shifu_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/shifu.ShifuService/InstallViaURL',
        request,
        metadata || {},
        this.methodDescriptorInstallViaURL,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/shifu.ShifuService/InstallViaURL',
    request,
    metadata || {},
    this.methodDescriptorInstallViaURL);
  }

}

