import { ClientReadableStream } from "grpc-web";
import { client } from "./shifu";
import {
  Device,
  GetDeviceDetailsResponse,
  GetDeviceDetailsRequest,
  ForwardPortResponse,
  ForwardPortRequest,
  RestartDeviceShifuRequest,
  Empty,
  DeleteDeviceShifuRequest,
  GetAllContainerNameRequest,
  GetAllContainerNameResponse,
  GetDeviceShifuLogsRequest,
  GetDeviceShifuLogsResponse,
} from "src/proto/proto/shifu/shifu_pb";
import { ListDevicesRequest } from "src/proto/proto/shifu/shifu_pb";

export function ListDevices(): Promise<Device[]> {
  return new Promise((resolve, reject) => {
    const request = new ListDevicesRequest();
    client.listDevices(request, {}, (err, response) => {
      if (err) {
        reject(err);
        return;
      }
      resolve(response.getDevicesList());
    });
  });
}

export function GetDeviceDetails(
  name: string
): Promise<GetDeviceDetailsResponse> {
  const request = new GetDeviceDetailsRequest();
  request.setName(name);
  return client.getDeviceDetails(request, {});
}

export function ForwardPort(
  deviceName: string,
  devicePort: string,
  localPort: string
): { promise: Promise<boolean>; cancel: () => void } {
  const stream = client.forwardPort(
    new ForwardPortRequest()
      .setDeviceName(deviceName)
      .setDevicePort(devicePort)
      .setLocalPort(localPort),
    {}
  );

  const promise = new Promise<boolean>((resolve, reject) => {
    stream.on("data", (response: ForwardPortResponse) => {
      resolve(response.getSuccess());
    });

    stream.on("error", (err) => {
      reject(err);
    });

    stream.on("end", () => {
      resolve(false);
    });
  });

  return {
    promise,
    cancel: () => {
      stream.cancel();
    },
  };
}

export function RestartDevice(deviceName: string): Promise<Empty> {
  const request = new RestartDeviceShifuRequest();
  request.setDeviceName(deviceName);
  return client.restartDeviceShifu(request, {});
}

export function DeleteDevice(deviceName: string): Promise<Empty> {
  const request = new DeleteDeviceShifuRequest();
  request.setDeviceName(deviceName);
  return client.deleteDeviceShifu(request, {});
}

export function GetAllContainerName(
  deviceName: string
): Promise<GetAllContainerNameResponse> {
  const request = new GetAllContainerNameRequest();
  request.setDeviceName(deviceName);
  return client.getAllContainerName(request, {});
}

export function GetDeviceShifuLogs(
  deviceName: string,
  containerName: string
): {
  stream: ClientReadableStream<GetDeviceShifuLogsResponse>;
  cancel: () => void;
} {
  const request = new GetDeviceShifuLogsRequest();
  request.setDeviceName(deviceName);
  request.setContainerName(containerName);
  const stream = client.getDeviceShifuLogs(request, {});

  return {
    stream,
    cancel: () => {
      stream.cancel();
    },
  };
}
