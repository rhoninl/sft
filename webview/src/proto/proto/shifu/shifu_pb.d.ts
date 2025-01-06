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

  getSetting(): DeviceSettings | undefined;
  setSetting(value?: DeviceSettings): Edgedevice;
  hasSetting(): boolean;
  clearSetting(): Edgedevice;

  getGateway(): GatewaySettings | undefined;
  setGateway(value?: GatewaySettings): Edgedevice;
  hasGateway(): boolean;
  clearGateway(): Edgedevice;

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
    setting?: DeviceSettings.AsObject,
    gateway?: GatewaySettings.AsObject,
  }

  export enum GatewayCase { 
    _GATEWAY_NOT_SET = 0,
    GATEWAY = 8,
  }
}

export class DeviceSettings extends jspb.Message {
  getMqtt(): MQTTSettings | undefined;
  setMqtt(value?: MQTTSettings): DeviceSettings;
  hasMqtt(): boolean;
  clearMqtt(): DeviceSettings;

  getOpcua(): OPCUASettings | undefined;
  setOpcua(value?: OPCUASettings): DeviceSettings;
  hasOpcua(): boolean;
  clearOpcua(): DeviceSettings;

  getSettingsCase(): DeviceSettings.SettingsCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeviceSettings.AsObject;
  static toObject(includeInstance: boolean, msg: DeviceSettings): DeviceSettings.AsObject;
  static serializeBinaryToWriter(message: DeviceSettings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeviceSettings;
  static deserializeBinaryFromReader(message: DeviceSettings, reader: jspb.BinaryReader): DeviceSettings;
}

export namespace DeviceSettings {
  export type AsObject = {
    mqtt?: MQTTSettings.AsObject,
    opcua?: OPCUASettings.AsObject,
  }

  export enum SettingsCase { 
    SETTINGS_NOT_SET = 0,
    MQTT = 1,
    OPCUA = 2,
  }
}

export class GatewaySettings extends jspb.Message {
  getProtocol(): string;
  setProtocol(value: string): GatewaySettings;

  getLwm2m(): LwM2MSettings | undefined;
  setLwm2m(value?: LwM2MSettings): GatewaySettings;
  hasLwm2m(): boolean;
  clearLwm2m(): GatewaySettings;

  getSettingsCase(): GatewaySettings.SettingsCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GatewaySettings.AsObject;
  static toObject(includeInstance: boolean, msg: GatewaySettings): GatewaySettings.AsObject;
  static serializeBinaryToWriter(message: GatewaySettings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GatewaySettings;
  static deserializeBinaryFromReader(message: GatewaySettings, reader: jspb.BinaryReader): GatewaySettings;
}

export namespace GatewaySettings {
  export type AsObject = {
    protocol: string,
    lwm2m?: LwM2MSettings.AsObject,
  }

  export enum SettingsCase { 
    SETTINGS_NOT_SET = 0,
    LWM2M = 2,
  }
}

export class LwM2MSettings extends jspb.Message {
  getEndpointname(): string;
  setEndpointname(value: string): LwM2MSettings;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): LwM2MSettings.AsObject;
  static toObject(includeInstance: boolean, msg: LwM2MSettings): LwM2MSettings.AsObject;
  static serializeBinaryToWriter(message: LwM2MSettings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): LwM2MSettings;
  static deserializeBinaryFromReader(message: LwM2MSettings, reader: jspb.BinaryReader): LwM2MSettings;
}

export namespace LwM2MSettings {
  export type AsObject = {
    endpointname: string,
  }
}

export class MQTTSettings extends jspb.Message {
  getMqtttopic(): string;
  setMqtttopic(value: string): MQTTSettings;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MQTTSettings.AsObject;
  static toObject(includeInstance: boolean, msg: MQTTSettings): MQTTSettings.AsObject;
  static serializeBinaryToWriter(message: MQTTSettings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MQTTSettings;
  static deserializeBinaryFromReader(message: MQTTSettings, reader: jspb.BinaryReader): MQTTSettings;
}

export namespace MQTTSettings {
  export type AsObject = {
    mqtttopic: string,
  }
}

export class OPCUASettings extends jspb.Message {
  getOpcuanodeid(): string;
  setOpcuanodeid(value: string): OPCUASettings;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OPCUASettings.AsObject;
  static toObject(includeInstance: boolean, msg: OPCUASettings): OPCUASettings.AsObject;
  static serializeBinaryToWriter(message: OPCUASettings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OPCUASettings;
  static deserializeBinaryFromReader(message: OPCUASettings, reader: jspb.BinaryReader): OPCUASettings;
}

export namespace OPCUASettings {
  export type AsObject = {
    opcuanodeid: string,
  }
}

