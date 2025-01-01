import { ShifuServiceClient } from "../../proto/proto/shifu/ShifuServiceClientPb";
import {
  CheckInstallationRequest,
  InstallShifuRequest,
} from "../../proto/proto/shifu/shifu_pb";

const client = new ShifuServiceClient("http://localhost:34550", null, {
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
