/**
 * @fileoverview gRPC-Web generated client stub for shifu
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v5.28.0-rc1
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

}
