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

