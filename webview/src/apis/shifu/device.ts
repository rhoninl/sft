import { client } from "./shifu";
import { Device } from "src/proto/proto/shifu/shifu_pb";
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