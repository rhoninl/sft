import * as jspb from 'google-protobuf'



export class CheckInstallationRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CheckInstallationRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CheckInstallationRequest): CheckInstallationRequest.AsObject;
  static serializeBinaryToWriter(message: CheckInstallationRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CheckInstallationRequest;
  static deserializeBinaryFromReader(message: CheckInstallationRequest, reader: jspb.BinaryReader): CheckInstallationRequest;
}

export namespace CheckInstallationRequest {
  export type AsObject = {
  }
}

export class CheckInstallationResponse extends jspb.Message {
  getInstalled(): boolean;
  setInstalled(value: boolean): CheckInstallationResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CheckInstallationResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CheckInstallationResponse): CheckInstallationResponse.AsObject;
  static serializeBinaryToWriter(message: CheckInstallationResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CheckInstallationResponse;
  static deserializeBinaryFromReader(message: CheckInstallationResponse, reader: jspb.BinaryReader): CheckInstallationResponse;
}

export namespace CheckInstallationResponse {
  export type AsObject = {
    installed: boolean,
  }
}

export class InstallShifuRequest extends jspb.Message {
  getVersion(): string;
  setVersion(value: string): InstallShifuRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InstallShifuRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InstallShifuRequest): InstallShifuRequest.AsObject;
  static serializeBinaryToWriter(message: InstallShifuRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InstallShifuRequest;
  static deserializeBinaryFromReader(message: InstallShifuRequest, reader: jspb.BinaryReader): InstallShifuRequest;
}

export namespace InstallShifuRequest {
  export type AsObject = {
    version: string,
  }
}

export class InstallShifuResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): InstallShifuResponse;

  getError(): string;
  setError(value: string): InstallShifuResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InstallShifuResponse.AsObject;
  static toObject(includeInstance: boolean, msg: InstallShifuResponse): InstallShifuResponse.AsObject;
  static serializeBinaryToWriter(message: InstallShifuResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InstallShifuResponse;
  static deserializeBinaryFromReader(message: InstallShifuResponse, reader: jspb.BinaryReader): InstallShifuResponse;
}

export namespace InstallShifuResponse {
  export type AsObject = {
    success: boolean,
    error: string,
  }
}

export class GetAllAvailableVersionsRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllAvailableVersionsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllAvailableVersionsRequest): GetAllAvailableVersionsRequest.AsObject;
  static serializeBinaryToWriter(message: GetAllAvailableVersionsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllAvailableVersionsRequest;
  static deserializeBinaryFromReader(message: GetAllAvailableVersionsRequest, reader: jspb.BinaryReader): GetAllAvailableVersionsRequest;
}

export namespace GetAllAvailableVersionsRequest {
  export type AsObject = {
  }
}

export class GetAllAvailableVersionsResponse extends jspb.Message {
  getVersionsList(): Array<string>;
  setVersionsList(value: Array<string>): GetAllAvailableVersionsResponse;
  clearVersionsList(): GetAllAvailableVersionsResponse;
  addVersions(value: string, index?: number): GetAllAvailableVersionsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllAvailableVersionsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllAvailableVersionsResponse): GetAllAvailableVersionsResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllAvailableVersionsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllAvailableVersionsResponse;
  static deserializeBinaryFromReader(message: GetAllAvailableVersionsResponse, reader: jspb.BinaryReader): GetAllAvailableVersionsResponse;
}

export namespace GetAllAvailableVersionsResponse {
  export type AsObject = {
    versionsList: Array<string>,
  }
}

export class UninstallShifuRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UninstallShifuRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UninstallShifuRequest): UninstallShifuRequest.AsObject;
  static serializeBinaryToWriter(message: UninstallShifuRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UninstallShifuRequest;
  static deserializeBinaryFromReader(message: UninstallShifuRequest, reader: jspb.BinaryReader): UninstallShifuRequest;
}

export namespace UninstallShifuRequest {
  export type AsObject = {
  }
}

export class UninstallShifuResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): UninstallShifuResponse;

  getError(): string;
  setError(value: string): UninstallShifuResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UninstallShifuResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UninstallShifuResponse): UninstallShifuResponse.AsObject;
  static serializeBinaryToWriter(message: UninstallShifuResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UninstallShifuResponse;
  static deserializeBinaryFromReader(message: UninstallShifuResponse, reader: jspb.BinaryReader): UninstallShifuResponse;
}

export namespace UninstallShifuResponse {
  export type AsObject = {
    success: boolean,
    error: string,
  }
}

