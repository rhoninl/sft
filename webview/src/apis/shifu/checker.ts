import axios from "axios";

export function InstallChecker(): Promise<boolean> {
  return axios
    .get("/api/shifu/checker")
    .then((response) => response.data.installed)
    .catch((error) => {
      console.error("Failed to check installation:", error);
      return false;
    });
}

export function InstallShifu(version: string): Promise<void> {
  return axios
    .get(`/api/shifu/install`)
    .then((response) => response.data)
    .catch((error) => {
      console.error("Failed to install Shifu:", error);
    });
}
