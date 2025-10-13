import { GetSystemInfo } from "../wailsjs/wailsjs/go/handlers/SystemHandler";

export async function getSystemInfo() {
  return await GetSystemInfo();
}
