import { client } from "./shifu";
import { InstallViaURLRequest } from "../../proto/proto/shifu/shifu_pb";

export function InstallViaURL(url: string): Promise<void> {
  return new Promise((resolve, reject) => {
    const request = new InstallViaURLRequest();
    request.setUrl(url);

    client.installViaURL(request, {}, (err, response) => {
      if (err) {
        console.error("Failed to install via URL:", err);
        reject(err);
        return;
      }
      resolve();
    });
  });
}
