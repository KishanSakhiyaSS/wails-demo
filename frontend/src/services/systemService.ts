import { 
  GetAllSystemInfo,
  GetOSInfo,
  GetCPUInfo,
  GetGPUInfo,
  GetMemoryInfo,
  GetDiskInfo,
  GetLocationInfo,
  GetHardwareInfo,
  GetUsagePercentages
} from "../../wailsjs/go/app/App";
import { SystemInfo, OSInfo, CPUInfo, GPUInfo, MemoryInfo, DiskInfo, LocationInfo, HardwareInfo, UsagePercentages } from "../types/system";

export async function getSystemInfo(): Promise<SystemInfo> {
  const data = await GetAllSystemInfo();
  return data as SystemInfo;
}

export async function getOSInfo(): Promise<OSInfo> {
  const data = await GetOSInfo();
  return data as OSInfo;
}

export async function getCPUInfo(): Promise<CPUInfo> {
  const data = await GetCPUInfo();
  return data as CPUInfo;
}

export async function getGPUInfo(): Promise<GPUInfo> {
  const data = await GetGPUInfo();
  return data as GPUInfo;
}

export async function getMemoryInfo(): Promise<MemoryInfo> {
  const data = await GetMemoryInfo();
  return data as MemoryInfo;
}

export async function getDiskInfo(): Promise<DiskInfo> {
  const data = await GetDiskInfo();
  return data as DiskInfo;
}

export async function getLocationInfo(): Promise<LocationInfo> {
  const data = await GetLocationInfo();
  return data as LocationInfo;
}

export async function getHardwareInfo(): Promise<HardwareInfo[]> {
  const data = await GetHardwareInfo();
  return data as HardwareInfo[];
}

export async function getUsagePercentages(): Promise<UsagePercentages> {
  const data = await GetUsagePercentages();
  return data as UsagePercentages;
}
