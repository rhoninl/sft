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

export class ListDevicesRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDevicesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ListDevicesRequest): ListDevicesRequest.AsObject;
  static serializeBinaryToWriter(message: ListDevicesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDevicesRequest;
  static deserializeBinaryFromReader(message: ListDevicesRequest, reader: jspb.BinaryReader): ListDevicesRequest;
}

export namespace ListDevicesRequest {
  export type AsObject = {
  }
}

export class ListDevicesResponse extends jspb.Message {
  getDevicesList(): Array<Device>;
  setDevicesList(value: Array<Device>): ListDevicesResponse;
  clearDevicesList(): ListDevicesResponse;
  addDevices(value?: Device, index?: number): Device;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListDevicesResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ListDevicesResponse): ListDevicesResponse.AsObject;
  static serializeBinaryToWriter(message: ListDevicesResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListDevicesResponse;
  static deserializeBinaryFromReader(message: ListDevicesResponse, reader: jspb.BinaryReader): ListDevicesResponse;
}

export namespace ListDevicesResponse {
  export type AsObject = {
    devicesList: Array<Device.AsObject>,
  }
}

export class GetDeviceDetailsRequest extends jspb.Message {
  getName(): string;
  setName(value: string): GetDeviceDetailsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDeviceDetailsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDeviceDetailsRequest): GetDeviceDetailsRequest.AsObject;
  static serializeBinaryToWriter(message: GetDeviceDetailsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDeviceDetailsRequest;
  static deserializeBinaryFromReader(message: GetDeviceDetailsRequest, reader: jspb.BinaryReader): GetDeviceDetailsRequest;
}

export namespace GetDeviceDetailsRequest {
  export type AsObject = {
    name: string,
  }
}

export class GetDeviceDetailsResponse extends jspb.Message {
  getEdgedevice(): Edgedevice | undefined;
  setEdgedevice(value?: Edgedevice): GetDeviceDetailsResponse;
  hasEdgedevice(): boolean;
  clearEdgedevice(): GetDeviceDetailsResponse;

  getApis(): string;
  setApis(value: string): GetDeviceDetailsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDeviceDetailsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDeviceDetailsResponse): GetDeviceDetailsResponse.AsObject;
  static serializeBinaryToWriter(message: GetDeviceDetailsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDeviceDetailsResponse;
  static deserializeBinaryFromReader(message: GetDeviceDetailsResponse, reader: jspb.BinaryReader): GetDeviceDetailsResponse;
}

export namespace GetDeviceDetailsResponse {
  export type AsObject = {
    edgedevice?: Edgedevice.AsObject,
    apis: string,
  }
}

export class Device extends jspb.Message {
  getName(): string;
  setName(value: string): Device;

  getProtocol(): string;
  setProtocol(value: string): Device;

  getAddress(): string;
  setAddress(value: string): Device;

  getStatus(): string;
  setStatus(value: string): Device;

  getAge(): string;
  setAge(value: string): Device;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Device.AsObject;
  static toObject(includeInstance: boolean, msg: Device): Device.AsObject;
  static serializeBinaryToWriter(message: Device, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Device;
  static deserializeBinaryFromReader(message: Device, reader: jspb.BinaryReader): Device;
}

export namespace Device {
  export type AsObject = {
    name: string,
    protocol: string,
    address: string,
    status: string,
    age: string,
  }
}

export class Edgedevice extends jspb.Message {
  getSku(): string;
  setSku(value: string): Edgedevice;

  getConnection(): string;
  setConnection(value: string): Edgedevice;

  getAddress(): string;
  setAddress(value: string): Edgedevice;

  getProtocol(): string;
  setProtocol(value: string): Edgedevice;

  getStatus(): string;
  setStatus(value: string): Edgedevice;

  getAge(): string;
  setAge(value: string): Edgedevice;

  getSetting(): string;
  setSetting(value: string): Edgedevice;

  getGateway(): string;
  setGateway(value: string): Edgedevice;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Edgedevice.AsObject;
  static toObject(includeInstance: boolean, msg: Edgedevice): Edgedevice.AsObject;
  static serializeBinaryToWriter(message: Edgedevice, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Edgedevice;
  static deserializeBinaryFromReader(message: Edgedevice, reader: jspb.BinaryReader): Edgedevice;
}

export namespace Edgedevice {
  export type AsObject = {
    sku: string,
    connection: string,
    address: string,
    protocol: string,
    status: string,
    age: string,
    setting: string,
    gateway: string,
  }
}

export class ForwardPortRequest extends jspb.Message {
  getDeviceName(): string;
  setDeviceName(value: string): ForwardPortRequest;

  getDevicePort(): string;
  setDevicePort(value: string): ForwardPortRequest;

  getLocalPort(): string;
  setLocalPort(value: string): ForwardPortRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ForwardPortRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ForwardPortRequest): ForwardPortRequest.AsObject;
  static serializeBinaryToWriter(message: ForwardPortRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ForwardPortRequest;
  static deserializeBinaryFromReader(message: ForwardPortRequest, reader: jspb.BinaryReader): ForwardPortRequest;
}

export namespace ForwardPortRequest {
  export type AsObject = {
    deviceName: string,
    devicePort: string,
    localPort: string,
  }
}

export class ForwardPortResponse extends jspb.Message {
  getSuccess(): boolean;
  setSuccess(value: boolean): ForwardPortResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ForwardPortResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ForwardPortResponse): ForwardPortResponse.AsObject;
  static serializeBinaryToWriter(message: ForwardPortResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ForwardPortResponse;
  static deserializeBinaryFromReader(message: ForwardPortResponse, reader: jspb.BinaryReader): ForwardPortResponse;
}

export namespace ForwardPortResponse {
  export type AsObject = {
    success: boolean,
  }
}

export class RestartDeviceShifuRequest extends jspb.Message {
  getDeviceName(): string;
  setDeviceName(value: string): RestartDeviceShifuRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RestartDeviceShifuRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RestartDeviceShifuRequest): RestartDeviceShifuRequest.AsObject;
  static serializeBinaryToWriter(message: RestartDeviceShifuRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RestartDeviceShifuRequest;
  static deserializeBinaryFromReader(message: RestartDeviceShifuRequest, reader: jspb.BinaryReader): RestartDeviceShifuRequest;
}

export namespace RestartDeviceShifuRequest {
  export type AsObject = {
    deviceName: string,
  }
}

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class DeleteDeviceShifuRequest extends jspb.Message {
  getDeviceName(): string;
  setDeviceName(value: string): DeleteDeviceShifuRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteDeviceShifuRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteDeviceShifuRequest): DeleteDeviceShifuRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteDeviceShifuRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteDeviceShifuRequest;
  static deserializeBinaryFromReader(message: DeleteDeviceShifuRequest, reader: jspb.BinaryReader): DeleteDeviceShifuRequest;
}

export namespace DeleteDeviceShifuRequest {
  export type AsObject = {
    deviceName: string,
  }
}

export class GetAllContainerNameRequest extends jspb.Message {
  getDeviceName(): string;
  setDeviceName(value: string): GetAllContainerNameRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllContainerNameRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllContainerNameRequest): GetAllContainerNameRequest.AsObject;
  static serializeBinaryToWriter(message: GetAllContainerNameRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllContainerNameRequest;
  static deserializeBinaryFromReader(message: GetAllContainerNameRequest, reader: jspb.BinaryReader): GetAllContainerNameRequest;
}

export namespace GetAllContainerNameRequest {
  export type AsObject = {
    deviceName: string,
  }
}

export class GetAllContainerNameResponse extends jspb.Message {
  getContainerNamesList(): Array<string>;
  setContainerNamesList(value: Array<string>): GetAllContainerNameResponse;
  clearContainerNamesList(): GetAllContainerNameResponse;
  addContainerNames(value: string, index?: number): GetAllContainerNameResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetAllContainerNameResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetAllContainerNameResponse): GetAllContainerNameResponse.AsObject;
  static serializeBinaryToWriter(message: GetAllContainerNameResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetAllContainerNameResponse;
  static deserializeBinaryFromReader(message: GetAllContainerNameResponse, reader: jspb.BinaryReader): GetAllContainerNameResponse;
}

export namespace GetAllContainerNameResponse {
  export type AsObject = {
    containerNamesList: Array<string>,
  }
}

export class GetDeviceShifuLogsRequest extends jspb.Message {
  getDeviceName(): string;
  setDeviceName(value: string): GetDeviceShifuLogsRequest;

  getContainerName(): string;
  setContainerName(value: string): GetDeviceShifuLogsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDeviceShifuLogsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetDeviceShifuLogsRequest): GetDeviceShifuLogsRequest.AsObject;
  static serializeBinaryToWriter(message: GetDeviceShifuLogsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDeviceShifuLogsRequest;
  static deserializeBinaryFromReader(message: GetDeviceShifuLogsRequest, reader: jspb.BinaryReader): GetDeviceShifuLogsRequest;
}

export namespace GetDeviceShifuLogsRequest {
  export type AsObject = {
    deviceName: string,
    containerName: string,
  }
}

export class GetDeviceShifuLogsResponse extends jspb.Message {
  getLog(): string;
  setLog(value: string): GetDeviceShifuLogsResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetDeviceShifuLogsResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetDeviceShifuLogsResponse): GetDeviceShifuLogsResponse.AsObject;
  static serializeBinaryToWriter(message: GetDeviceShifuLogsResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetDeviceShifuLogsResponse;
  static deserializeBinaryFromReader(message: GetDeviceShifuLogsResponse, reader: jspb.BinaryReader): GetDeviceShifuLogsResponse;
}

export namespace GetDeviceShifuLogsResponse {
  export type AsObject = {
    log: string,
  }
}

export class CommandRequest extends jspb.Message {
  getCommand(): string;
  setCommand(value: string): CommandRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CommandRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CommandRequest): CommandRequest.AsObject;
  static serializeBinaryToWriter(message: CommandRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CommandRequest;
  static deserializeBinaryFromReader(message: CommandRequest, reader: jspb.BinaryReader): CommandRequest;
}

export namespace CommandRequest {
  export type AsObject = {
    command: string,
  }
}

export class CommandResponse extends jspb.Message {
  getOutput(): string;
  setOutput(value: string): CommandResponse;

  getIsError(): boolean;
  setIsError(value: boolean): CommandResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CommandResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CommandResponse): CommandResponse.AsObject;
  static serializeBinaryToWriter(message: CommandResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CommandResponse;
  static deserializeBinaryFromReader(message: CommandResponse, reader: jspb.BinaryReader): CommandResponse;
}

export namespace CommandResponse {
  export type AsObject = {
    output: string,
    isError: boolean,
  }
}

export class CompletionRequest extends jspb.Message {
  getPartial(): string;
  setPartial(value: string): CompletionRequest;

  getCurrentDir(): string;
  setCurrentDir(value: string): CompletionRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CompletionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CompletionRequest): CompletionRequest.AsObject;
  static serializeBinaryToWriter(message: CompletionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CompletionRequest;
  static deserializeBinaryFromReader(message: CompletionRequest, reader: jspb.BinaryReader): CompletionRequest;
}

export namespace CompletionRequest {
  export type AsObject = {
    partial: string,
    currentDir: string,
  }
}

export class CompletionResponse extends jspb.Message {
  getCompletionsList(): Array<string>;
  setCompletionsList(value: Array<string>): CompletionResponse;
  clearCompletionsList(): CompletionResponse;
  addCompletions(value: string, index?: number): CompletionResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CompletionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CompletionResponse): CompletionResponse.AsObject;
  static serializeBinaryToWriter(message: CompletionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CompletionResponse;
  static deserializeBinaryFromReader(message: CompletionResponse, reader: jspb.BinaryReader): CompletionResponse;
}

export namespace CompletionResponse {
  export type AsObject = {
    completionsList: Array<string>,
  }
}

export class InstallViaURLRequest extends jspb.Message {
  getUrl(): string;
  setUrl(value: string): InstallViaURLRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InstallViaURLRequest.AsObject;
  static toObject(includeInstance: boolean, msg: InstallViaURLRequest): InstallViaURLRequest.AsObject;
  static serializeBinaryToWriter(message: InstallViaURLRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InstallViaURLRequest;
  static deserializeBinaryFromReader(message: InstallViaURLRequest, reader: jspb.BinaryReader): InstallViaURLRequest;
}

export namespace InstallViaURLRequest {
  export type AsObject = {
    url: string,
  }
}

