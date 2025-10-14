import { useState } from "react";
import { AllSystemData } from "../types/system";

interface DashboardProps {
  systemData: AllSystemData;
  onRefresh: () => Promise<void>;
  onRefreshCPU: () => Promise<void>;
  onRefreshGPU: () => Promise<void>;
  onRefreshMemory: () => Promise<void>;
  onRefreshDisk: () => Promise<void>;
  onRefreshSystem: () => Promise<void>;
  onRefreshHardware: () => Promise<void>;
}

// Circular Progress Component
interface CircularProgressProps {
  percentage: number;
  size?: number;
  strokeWidth?: number;
  color?: string;
}

function CircularProgress({ percentage, size = 120, strokeWidth = 8, color = "#3B82F6" }: CircularProgressProps) {
  const radius = (size - strokeWidth) / 2;
  const circumference = radius * 2 * Math.PI;
  const strokeDasharray = circumference;
  const strokeDashoffset = circumference - (percentage / 100) * circumference;

  return (
    <div className="relative inline-flex items-center justify-center" style={{ width: size, height: size }}>
      <svg
        width={size}
        height={size}
        className="transform -rotate-90"
      >
        {/* Background circle */}
        <circle
          cx={size / 2}
          cy={size / 2}
          r={radius}
          stroke="#374151"
          strokeWidth={strokeWidth}
          fill="none"
        />
        {/* Progress circle */}
        <circle
          cx={size / 2}
          cy={size / 2}
          r={radius}
          stroke={color}
          strokeWidth={strokeWidth}
          fill="none"
          strokeDasharray={strokeDasharray}
          strokeDashoffset={strokeDashoffset}
          strokeLinecap="round"
          className="transition-all duration-300 ease-in-out"
        />
      </svg>
      {/* Percentage text */}
      <div className="absolute inset-0 flex items-center justify-center">
        <span className="text-2xl font-bold text-white">
          {percentage.toFixed(1)}%
        </span>
      </div>
    </div>
  );
}

// Usage Card Component
interface UsageCardProps {
  title: string;
  percentage: number;
  details: string;
  icon: React.ReactNode;
  color: string;
}

function UsageCard({ title, percentage, details, icon, color }: UsageCardProps) {
  return (
    <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
      <div className="flex items-center gap-3 mb-4">
        <div className="text-2xl">
          {icon}
        </div>
        <h3 className="text-lg font-semibold text-white">{title}</h3>
      </div>
      
      <div className="flex justify-center mb-4">
        <CircularProgress percentage={percentage} color={color} />
      </div>
      
      <p className="text-sm text-gray-400 text-center">{details}</p>
    </div>
  );
}

export default function Dashboard({ 
  systemData, 
  onRefresh, 
  onRefreshCPU, 
  onRefreshGPU, 
  onRefreshMemory, 
  onRefreshDisk, 
  onRefreshSystem, 
  onRefreshHardware 
}: DashboardProps) {
  const { cpu, gpu, gpus, memory, disk, os, hardware } = systemData;
  const [activeTab, setActiveTab] = useState('overview');
  const [refreshing, setRefreshing] = useState<string | null>(null);

  const handleRefresh = async () => {
    setRefreshing('all');
    try {
      await onRefresh();
    } finally {
      setRefreshing(null);
    }
  };

  const handleRefreshCPU = async () => {
    setRefreshing('cpu');
    try {
      await onRefreshCPU();
    } finally {
      setRefreshing(null);
    }
  };

  const handleRefreshGPU = async () => {
    setRefreshing('gpu');
    try {
      await onRefreshGPU();
    } finally {
      setRefreshing(null);
    }
  };

  const handleRefreshMemory = async () => {
    setRefreshing('memory');
    try {
      await onRefreshMemory();
    } finally {
      setRefreshing(null);
    }
  };

  const handleRefreshDisk = async () => {
    setRefreshing('disk');
    try {
      await onRefreshDisk();
    } finally {
      setRefreshing(null);
    }
  };

  const handleRefreshSystem = async () => {
    setRefreshing('system');
    try {
      await onRefreshSystem();
    } finally {
      setRefreshing(null);
    }
  };

  const handleRefreshHardware = async () => {
    setRefreshing('hardware');
    try {
      await onRefreshHardware();
    } finally {
      setRefreshing(null);
    }
  };

  const handleExportData = () => {
    // Export functionality would go here
    console.log('Export data');
  };

  const tabs = [
    { id: 'overview', label: 'Overview' },
    { id: 'cpu', label: 'CPU' },
    { id: 'gpu', label: 'GPU' },
    { id: 'memory', label: 'Memory' },
    { id: 'disk', label: 'Disk' },
    { id: 'system', label: 'System' },
    { id: 'hardware', label: 'Hardware' }
  ];

  // Extract usage percentages
  const cpuUsage = cpu?.cpu_usage_percentage ? parseFloat(cpu.cpu_usage_percentage.replace('%', '')) : 0;
  const gpuUsage = gpu?.usage_percentage ? parseFloat(gpu.usage_percentage.replace('%', '')) : 0;
  const memoryUsage = memory?.used_percentage ? parseFloat(memory.used_percentage.replace('%', '')) : 0;
  const diskUsage = disk?.used_percentage ? parseFloat(disk.used_percentage.replace('%', '')) : 0;

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header Section */}
      <div className="bg-gray-900 px-6 py-8">
        <div className="max-w-7xl mx-auto">
          {/* Title and Icon */}
          <div className="flex items-center justify-center mb-4">
            <div className="text-4xl mr-4">🖥️</div>
            <div className="text-center">
              <h1 className="text-4xl font-bold text-purple-500 mb-2">System Benchmark</h1>
              <p className="text-gray-400 text-lg">Comprehensive System Information Dashboard</p>
            </div>
          </div>

          {/* Action Buttons */}
          <div className="flex justify-center gap-4 mb-8">
            <button
              onClick={handleRefresh}
              disabled={refreshing === 'all'}
              className={`px-6 py-3 text-white font-medium rounded-lg transition-colors flex items-center gap-2 ${
                refreshing === 'all' 
                  ? 'bg-purple-500 cursor-not-allowed' 
                  : 'bg-purple-600 hover:bg-purple-700'
              }`}
            >
              <svg className={`w-4 h-4 ${refreshing === 'all' ? 'animate-spin' : ''}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              {refreshing === 'all' ? 'Refreshing All...' : 'Refresh All'}
            </button>
            <button
              onClick={handleExportData}
              className="px-6 py-3 bg-gray-700 hover:bg-gray-600 text-white font-medium rounded-lg transition-colors flex items-center gap-2"
            >
              <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z" />
              </svg>
              Export Data
            </button>
          </div>

          {/* Navigation Tabs */}
          <div className="flex justify-center">
            <div className="flex bg-gray-800 rounded-lg p-1">
              {tabs.map((tab) => (
                <button
                  key={tab.id}
                  onClick={() => setActiveTab(tab.id)}
                  className={`px-4 py-2 rounded-md font-medium transition-colors ${
                    activeTab === tab.id
                      ? 'bg-purple-600 text-white'
                      : 'text-gray-400 hover:text-white hover:bg-gray-700'
                  }`}
                >
                  {tab.label}
                </button>
              ))}
            </div>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className="px-6 pb-8">
        <div className="max-w-7xl mx-auto">
          {activeTab === 'overview' && (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
              {/* CPU Usage Card */}
              <UsageCard
                title="CPU Usage"
                percentage={cpuUsage}
                details={cpu?.model || 'CPU information not available'}
                icon={<span className="text-blue-400">💻</span>}
                color="#3B82F6"
              />

              {/* GPU Usage Card */}
              <UsageCard
                title="GPU Usage"
                percentage={gpuUsage}
                details={gpu?.name || 'GPU information not available'}
                icon={<span className="text-purple-400">🎮</span>}
                color={gpuUsage > 0 ? "#8B5CF6" : "#6B7280"}
              />

              {/* Memory Usage Card */}
              <UsageCard
                title="Memory Usage"
                percentage={memoryUsage}
                details={memory ? `${memory.used} / ${memory.total}` : 'Memory information not available'}
                icon={<span className="text-pink-400">🧠</span>}
                color="#EC4899"
              />

              {/* Disk Usage Card */}
              <UsageCard
                title="Disk Usage"
                percentage={diskUsage}
                details={disk ? `${disk.used} / ${disk.total}` : 'Disk information not available'}
                icon={<span className="text-purple-400">💾</span>}
                color="#8B5CF6"
              />
            </div>
          )}

          {activeTab === 'cpu' && (
            <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
              <div className="flex justify-between items-center mb-6">
                <h3 className="text-2xl font-bold text-white">CPU Information</h3>
            <button
              onClick={handleRefreshCPU}
              disabled={refreshing === 'cpu'}
              className={`px-4 py-2 text-white font-medium rounded-lg transition-colors flex items-center gap-2 ${
                refreshing === 'cpu' 
                  ? 'bg-blue-500 cursor-not-allowed' 
                  : 'bg-blue-600 hover:bg-blue-700'
              }`}
            >
              <svg className={`w-4 h-4 ${refreshing === 'cpu' ? 'animate-spin' : ''}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              {refreshing === 'cpu' ? 'Refreshing...' : 'Refresh CPU'}
            </button>
              </div>
              {cpu ? (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Model</p>
                    <p className="font-semibold text-white break-words">{cpu.model}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Cores</p>
                    <p className="font-semibold text-white">{cpu.cores}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Frequency</p>
                    <p className="font-semibold text-white">{cpu.ghz}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Usage</p>
                    <p className="font-semibold text-blue-400">{cpu.cpu_usage_percentage}</p>
                  </div>
                </div>
              ) : (
                <div className="text-center py-8">
                  <p className="text-gray-400 text-lg">No CPU information available</p>
                </div>
              )}
            </div>
          )}

          {activeTab === 'gpu' && (
            <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
              <div className="flex justify-between items-center mb-6">
                <h3 className="text-2xl font-bold text-white">GPU Information</h3>
                <button
                  onClick={handleRefreshGPU}
                  disabled={refreshing === 'gpu'}
                  className={`px-4 py-2 text-white font-medium rounded-lg transition-colors flex items-center gap-2 ${
                    refreshing === 'gpu' 
                      ? 'bg-purple-500 cursor-not-allowed' 
                      : 'bg-purple-600 hover:bg-purple-700'
                  }`}
                >
                  <svg className={`w-4 h-4 ${refreshing === 'gpu' ? 'animate-spin' : ''}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                  {refreshing === 'gpu' ? 'Refreshing...' : 'Refresh GPU'}
                </button>
              </div>
              {gpus && gpus.length > 0 ? (
                <div className="space-y-6">
                  {gpus.map((gpuItem, index) => (
                    <div key={index} className="bg-gray-700/50 rounded-lg p-6 border border-gray-600">
                      <h4 className="text-xl font-semibold text-white mb-4 flex items-center gap-2">
                        <span className="text-purple-400">GPU {index + 1}</span>
                        <span className="text-gray-400 text-sm">• {gpuItem.name}</span>
                      </h4>
                      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                        <div className="bg-gray-700 rounded-lg p-4">
                          <p className="text-sm text-gray-400 mb-1">Name</p>
                          <p className="font-semibold text-white break-words">{gpuItem.name}</p>
                        </div>
                        <div className="bg-gray-700 rounded-lg p-4">
                          <p className="text-sm text-gray-400 mb-1">VRAM Size</p>
                          <p className="font-semibold text-white">{gpuItem.vram_size}</p>
                        </div>
                        <div className="bg-gray-700 rounded-lg p-4">
                          <p className="text-sm text-gray-400 mb-1">Driver</p>
                          <p className="font-semibold text-white break-words">{gpuItem.driver}</p>
                        </div>
                        <div className="bg-gray-700 rounded-lg p-4">
                          <p className="text-sm text-gray-400 mb-1">Usage</p>
                          <p className="font-semibold text-purple-400">{gpuItem.usage_percentage}</p>
                        </div>
                        <div className="bg-gray-700 rounded-lg p-4">
                          <p className="text-sm text-gray-400 mb-1">Clock Speed</p>
                          <p className="font-semibold text-white">{gpuItem.clock_speed}</p>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              ) : (
                <div className="text-center py-8">
                  <p className="text-gray-400 text-lg">No GPU information available</p>
                  <p className="text-gray-500 text-sm mt-2">This might be due to:</p>
                  <ul className="text-gray-500 text-sm mt-2 text-left max-w-md mx-auto">
                    <li>• No dedicated GPU detected</li>
                    <li>• GPU drivers not installed</li>
                    <li>• GPU monitoring tools not available</li>
                  </ul>
                </div>
              )}
            </div>
          )}

          {activeTab === 'memory' && (
            <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
              <div className="flex justify-between items-center mb-6">
                <h3 className="text-2xl font-bold text-white">Memory Information</h3>
                <button
                  onClick={handleRefreshMemory}
                  disabled={refreshing === 'memory'}
                  className={`px-4 py-2 text-white font-medium rounded-lg transition-colors flex items-center gap-2 ${
                    refreshing === 'memory' 
                      ? 'bg-pink-500 cursor-not-allowed' 
                      : 'bg-pink-600 hover:bg-pink-700'
                  }`}
                >
                  <svg className={`w-4 h-4 ${refreshing === 'memory' ? 'animate-spin' : ''}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                  {refreshing === 'memory' ? 'Refreshing...' : 'Refresh Memory'}
                </button>
              </div>
              {memory ? (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Usage</p>
                    <p className="font-semibold text-pink-400 text-xl">{memory.used_percentage}</p>
                    <p className="text-xs text-gray-400 mt-1">{memory.used} / {memory.total}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Available</p>
                    <p className="font-semibold text-white">{memory.available}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Total</p>
                    <p className="font-semibold text-white">{memory.total}</p>
                  </div>
                </div>
              ) : (
                <div className="text-center py-8">
                  <p className="text-gray-400 text-lg">No memory information available</p>
                </div>
              )}
            </div>
          )}

          {activeTab === 'disk' && (
            <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
              <div className="flex justify-between items-center mb-6">
                <h3 className="text-2xl font-bold text-white">Disk Information</h3>
                <button
                  onClick={handleRefreshDisk}
                  disabled={refreshing === 'disk'}
                  className={`px-4 py-2 text-white font-medium rounded-lg transition-colors flex items-center gap-2 ${
                    refreshing === 'disk' 
                      ? 'bg-purple-500 cursor-not-allowed' 
                      : 'bg-purple-600 hover:bg-purple-700'
                  }`}
                >
                  <svg className={`w-4 h-4 ${refreshing === 'disk' ? 'animate-spin' : ''}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                  {refreshing === 'disk' ? 'Refreshing...' : 'Refresh Disk'}
                </button>
              </div>
              {disk ? (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Usage</p>
                    <p className="font-semibold text-purple-400 text-xl">{disk.used_percentage}</p>
                    <p className="text-xs text-gray-400 mt-1">{disk.used} / {disk.total}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Free</p>
                    <p className="font-semibold text-white">{disk.free}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Total</p>
                    <p className="font-semibold text-white">{disk.total}</p>
                  </div>
                </div>
              ) : (
                <div className="text-center py-8">
                  <p className="text-gray-400 text-lg">No disk information available</p>
                </div>
              )}
            </div>
          )}

          {activeTab === 'system' && (
            <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
              <div className="flex justify-between items-center mb-6">
                <h3 className="text-2xl font-bold text-white">System Information</h3>
                <button
                  onClick={handleRefreshSystem}
                  disabled={refreshing === 'system'}
                  className={`px-4 py-2 text-white font-medium rounded-lg transition-colors flex items-center gap-2 ${
                    refreshing === 'system' 
                      ? 'bg-gray-500 cursor-not-allowed' 
                      : 'bg-gray-600 hover:bg-gray-700'
                  }`}
                >
                  <svg className={`w-4 h-4 ${refreshing === 'system' ? 'animate-spin' : ''}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                  {refreshing === 'system' ? 'Refreshing...' : 'Refresh System'}
                </button>
              </div>
              {os ? (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Platform</p>
                    <p className="font-semibold text-white break-words">{os.platform}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Architecture</p>
                    <p className="font-semibold text-white">{os.kernel_arch}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Hostname</p>
                    <p className="font-semibold text-white break-words">{os.hostname}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Kernel Version</p>
                    <p className="font-semibold text-white break-words">{os.kernel_version}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Platform Family</p>
                    <p className="font-semibold text-white">{os.platform_family}</p>
                  </div>
                  <div className="bg-gray-700 rounded-lg p-4">
                    <p className="text-sm text-gray-400 mb-1">Platform Version</p>
                    <p className="font-semibold text-white">{os.platform_version}</p>
                  </div>
                </div>
              ) : (
                <div className="text-center py-8">
                  <p className="text-gray-400 text-lg">No system information available</p>
                </div>
              )}
            </div>
          )}

          {activeTab === 'hardware' && (
            <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
              <div className="flex justify-between items-center mb-6">
                <h3 className="text-2xl font-bold text-white">Network Interfaces</h3>
                <button
                  onClick={handleRefreshHardware}
                  disabled={refreshing === 'hardware'}
                  className={`px-4 py-2 text-white font-medium rounded-lg transition-colors flex items-center gap-2 ${
                    refreshing === 'hardware' 
                      ? 'bg-blue-500 cursor-not-allowed' 
                      : 'bg-blue-600 hover:bg-blue-700'
                  }`}
                >
                  <svg className={`w-4 h-4 ${refreshing === 'hardware' ? 'animate-spin' : ''}`} fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                  {refreshing === 'hardware' ? 'Refreshing...' : 'Refresh Hardware'}
                </button>
              </div>
              {hardware && hardware.length > 0 ? (
                <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
                  {hardware.map((hw, index) => (
                    <div key={index} className="bg-gray-700 rounded-lg p-4 border border-gray-600">
                      <div className="flex items-center gap-2 mb-3">
                        <span className="text-blue-400">🌐</span>
                        <p className="font-semibold text-white">{hw.name}</p>
                      </div>
                      <div className="space-y-2 text-sm">
                        <div className="flex justify-between">
                          <span className="text-gray-400">MAC Address:</span>
                          <span className="font-medium text-white">{hw.hardware_addr}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-gray-400">MTU:</span>
                          <span className="font-medium text-white">{hw.mtu}</span>
                        </div>
                        <div className="flex justify-between">
                          <span className="text-gray-400">Flags:</span>
                          <span className="font-medium text-white truncate max-w-[200px]" title={hw.flags}>
                            {hw.flags}
                          </span>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              ) : (
                <div className="text-center py-8">
                  <p className="text-gray-400 text-lg">No hardware information available</p>
                </div>
              )}
            </div>
          )}

        </div>
      </div>
    </div>
  );
}
