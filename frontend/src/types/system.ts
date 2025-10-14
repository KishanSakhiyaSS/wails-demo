// Type definitions for system information based on backend models

export interface CPUInfo {
  cores: number;
  model: string;
  cache_size: string;
  ghz: string;
  cpu_usage_percentage: string;
}

export interface GPUInfo {
  name: string;
  vram_size: string;
  driver: string;
  usage_percentage: string;
  clock_speed: string;
}

export interface OSInfo {
  os: string;
  hostname: string;
  platform: string;
  platform_version: string;
  platform_family: string;
  kernel_version: string;
  kernel_arch: string;
  uptime: number;
}

export interface LocationInfo {
  ip: string;
  hostname: string;
  city: string;
  region: string;
  country: string;
  loc: string;
  org: string;
  postal: string;
  timezone: string;
  readme: string;
}

export interface MemoryInfo {
  total: string;
  used: string;
  free: string;
  available: string;
  used_percentage: string;
}

export interface DiskInfo {
  total: string;
  used: string;
  free: string;
  used_percentage: string;
}

export interface HardwareInfo {
  index: number;
  mtu: number;
  name: string;
  hardware_addr: string;
  flags: string;
}

export interface SystemInfo {
  cpu: CPUInfo;
  gpus: GPUInfo[];
  os: OSInfo;
  location: LocationInfo;
  memory: MemoryInfo;
  disk: DiskInfo;
  hardware: HardwareInfo[];
}

export interface AllSystemData {
  cpu: CPUInfo | null;
  gpu: GPUInfo | null; // Will be the first GPU from the array
  gpus: GPUInfo[] | null; // Full array of GPUs
  os: OSInfo | null;
  location: LocationInfo | null;
  memory: MemoryInfo | null;
  disk: DiskInfo | null;
  hardware: HardwareInfo[] | null;
  usagePercentages: UsagePercentages | null;
}

export interface UsagePercentages {
  cpu_usage: string;
  gpu_usage: string;
  memory_usage: string;
  disk_usage: string;
}

export interface User {
  name: string;
  role: string;
}

