import { ShifuServiceClient } from "../../proto/proto/shifu/ShifuServiceClientPb";
import {
  CheckInstallationRequest,
  GetAllAvailableVersionsRequest,
  InstallShifuRequest,
  UninstallShifuRequest,
} from "../../proto/proto/shifu/shifu_pb";

export const client = new ShifuServiceClient("", null, {
  withCredentials: false,
  headers: {
    "Content-Type": "application/grpc-web+proto",
  },
});

export function InstallChecker(): Promise<boolean> {
  return new Promise((resolve, reject) => {
    const request = new CheckInstallationRequest();

    client.checkInstallation(request, {}, (err, response) => {
      if (err) {
        console.error("Failed to check installation:", err);
        reject(err);
        return;
      }
      resolve(response.getInstalled());
    });
  });
}

export function InstallShifu(version: string): Promise<void> {
  return new Promise((resolve, reject) => {
    const request = new InstallShifuRequest();
    request.setVersion(version);

    client.installShifu(request, {}, (err, response) => {
      if (err) {
        console.error("Failed to install Shifu:", err);
        reject(err);
        return;
      }
      if (!response.getSuccess()) {
        reject(new Error(response.getError()));
        return;
      }
      resolve();
    });
  });
}

export function GetAllAvailableVersions(): Promise<string[]> {
  return new Promise((resolve, reject) => {
    const request = new GetAllAvailableVersionsRequest();
    client.getAllAvailableVersions(request, {}, (err, response) => {
      if (err) {
        console.error("Failed to get all available versions:", err);
        reject(err);
        return;
      }
      resolve(response.getVersionsList());
    });
  });
}

export function UninstallShifu(): Promise<void> {
  return new Promise((resolve, reject) => {
    const request = new UninstallShifuRequest();
    client.uninstallShifu(request, {}, (err, response) => {
      if (err) {
        reject(err);
        return;
      }
      resolve();
    });
  });
}
